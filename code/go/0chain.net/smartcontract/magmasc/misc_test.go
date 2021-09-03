package magmasc

import (
	"errors"
	"io/ioutil"
	"math/rand"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/0chain/gorocksdb"
	zbm "github.com/0chain/gosdk/zmagmacore/magmasc"
	"github.com/stretchr/testify/require"

	"0chain.net/chaincore/block"
	chain "0chain.net/chaincore/chain/state"
	"0chain.net/chaincore/state"
	"0chain.net/chaincore/transaction"
	store "0chain.net/core/ememorystore"
	"0chain.net/core/encryption"
	"0chain.net/core/util"
)

func createStateContextAndDB(sciDbDir, logDir, dbDir string, txn *transaction.Transaction, t require.TestingT) (
	*gorocksdb.TransactionDB, chain.StateContextI) {

	var (
		sci = createSCI(sciDbDir, logDir, txn, t)
		db  = createDB(dbDir, t)
	)
	sci.GetState().SetRoot(getRoot(db, t))

	return db, sci
}

// createSCI creates state.StateContextI with only util.NewMerklePatriciaTrie initialized,
// and provided transaction.
//
// For util.NewMerklePatriciaTrie util.PNodeDB is used.
func createSCI(dbDir, logDir string, txn *transaction.Transaction, t require.TestingT) chain.StateContextI {
	pNodeDB, err := util.NewPNodeDB(dbDir, logDir)
	require.NoError(t, err)

	return chain.NewStateContext(
		&block.Block{},
		util.NewMerklePatriciaTrie(pNodeDB, 1),
		&state.Deserializer{},
		txn,
		func(*block.Block) []string { return []string{} },
		func() *block.Block { return &block.Block{} },
		func() *block.MagicBlock { return &block.MagicBlock{} },
		func() encryption.SignatureScheme { return &encryption.BLS0ChainScheme{} },
	)
}

// createDB opens gorocksdb.TransactionDB  on provided path.
func createDB(path string, t require.TestingT) *gorocksdb.TransactionDB {
	db, err := store.CreateDB(path)
	require.NoError(t, err)

	store.AddPool(storeName, db)

	return db
}

func countSessions(sciDbDir, logDir, dbDir string, t require.TestingT) (act, inact int) {
	var (
		sc = NewMagmaSmartContract()

		db, sci = createStateContextAndDB(sciDbDir, logDir, dbDir, nil, t)
	)
	sc.SetDB(db)

	handl := sc.RestHandlers["/acknowledgmentAccepted"]
	for i := 0; ; i++ {
		val := url.Values{}
		val.Set("id", getSessionName(i, true))
		output, err := handl(nil, val, sci)
		if err != nil && errors.Is(err, util.ErrValueNotPresent) {
			break
		}
		require.NoError(t, err)

		outputAckn := output.(*zbm.Acknowledgment)
		if outputAckn.Billing.CompletedAt == 0 {
			act++
		} else {
			inact++
		}
	}

	for i := 0; ; i++ {
		val := url.Values{}
		val.Set("id", getSessionName(i, false))
		output, err := handl(nil, val, sci)
		if err != nil && errors.Is(err, util.ErrValueNotPresent) {
			break
		}
		require.NoError(t, err)

		outputAckn := output.(*zbm.Acknowledgment)
		if outputAckn.Billing.CompletedAt == 0 {
			act++
		} else {
			inact++
		}
	}

	closeSciAndDB(sci, db, t)

	return act, inact
}

func countNodes(sciDbDir, logDir, dbDir string, t require.TestingT) (consumers, providers int) {
	var (
		sc = NewMagmaSmartContract()

		db, sci = createStateContextAndDB(sciDbDir, logDir, dbDir, nil, t)
	)
	sc.SetDB(db)

	output, err := sc.allConsumers(nil, nil, sci)
	require.NoError(t, err)

	cList := output.([]*zbm.Consumer)
	consumers = len(cList)

	output, err = sc.allProviders(nil, nil, sci)
	require.NoError(t, err)

	pList := output.([]*zbm.Provider)
	providers = len(pList)

	closeSciAndDB(sci, db, t)
	return consumers, providers
}

const (
	sessionActPrefix   = "act_"
	sessionInactPrefix = "inact_"
	sessionName        = "session_"
)

func getSessionName(num int, active bool) string {
	prefix := ""
	if active {
		prefix = sessionActPrefix
	} else {
		prefix = sessionInactPrefix
	}

	return prefix + sessionName + strconv.Itoa(num)
}

const (
	rootKey = "root"
)

func saveRoot(root []byte, db *gorocksdb.TransactionDB, t require.TestingT) {
	conn := store.GetTransaction(db)
	err := conn.Conn.Put([]byte(rootKey), root)
	require.NoError(t, err)

	err = conn.Commit()
	require.NoError(t, err)
}

func getRoot(db *gorocksdb.TransactionDB, t require.TestingT) []byte {
	conn := store.GetTransaction(db)
	sl, err := conn.Conn.Get(conn.ReadOptions, []byte(rootKey))
	require.NoError(t, err)

	err = conn.Commit()
	require.NoError(t, err)
	return sl.Data()
}

func copyDir(source, destination string, t require.TestingT) {
	var err = filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		var relPath = strings.Replace(path, source, "", 1)
		if relPath == "" {
			return nil
		}
		if info.IsDir() {
			return os.Mkdir(filepath.Join(destination, relPath), 0755)
		} else {
			var (
				data, err = ioutil.ReadFile(filepath.Join(source, relPath))
			)
			if err != nil {
				return err
			}
			return ioutil.WriteFile(filepath.Join(destination, relPath), data, 0777)
		}
	})
	require.NoError(t, err)
}

func closeSciAndDB(sci chain.StateContextI, db *gorocksdb.TransactionDB, t require.TestingT) {
	saveRoot(sci.GetState().GetRoot(), db, t)

	db.Close()

	pNodeDB, ok := sci.GetState().GetNodeDB().(*util.PNodeDB)
	require.True(t, ok, "Must be pNodeDB type")
	pNodeDB.Close()
}

func randString(n int) string {
	rand.Seed(time.Now().UnixNano())

	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
