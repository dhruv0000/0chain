package state

import (
	"github.com/0chain/gorocksdb"

	store "0chain.net/core/ememorystore"
)

const (
	rootKey = "root"
)

func saveRoot(root []byte, db *gorocksdb.TransactionDB) error {
	conn := store.GetTransaction(db)
	if err := conn.Conn.Put([]byte(rootKey), root); err != nil {
		_ = conn.Conn.Rollback()
		return err
	}
	return conn.Commit()
}

func getRoot(db *gorocksdb.TransactionDB) ([]byte, error) {
	conn := store.GetTransaction(db)
	sl, err := conn.Conn.Get(conn.ReadOptions, []byte(rootKey))
	if err != nil {
		_ = conn.Conn.Rollback()
		return nil, err
	}

	return sl.Data(), conn.Commit()
}
