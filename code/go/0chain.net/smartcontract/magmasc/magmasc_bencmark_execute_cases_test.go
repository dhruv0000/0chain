package magmasc

import (
	"github.com/0chain/gorocksdb"
	"github.com/0chain/gosdk/zmagmacore/magmasc"
	"github.com/stretchr/testify/require"

	chain "0chain.net/chaincore/chain/state"
	"0chain.net/chaincore/state"
	"0chain.net/chaincore/transaction"
	"0chain.net/core/util"
)

type (
	// benchExecuteCase represents interface that provides all needed data for
	// executing MagmaSmartContract.Execute.
	benchExecuteCase interface {
		txn() *transaction.Transaction
		funcName() string
		input() []byte
		setupDatabaseAndState(dbDir, stateDBDir, stateLogDir string, t require.TestingT) (*gorocksdb.TransactionDB, chain.StateContextI)
		refresh()
	}
)

// loadBenchExecuteCases loads all benchExecuteCase for each function of smart contract.
func loadBenchExecuteCases() []benchExecuteCase {
	return []benchExecuteCase{
		newConsumerRegisterEC(),
		newProviderRegisterEC(),
		newProviderEC(),
		newConsumerUpdateEC(),
		newProviderSessionInitEC(),
		newConsumerSessionStartEC(),
		newProviderDataUsageEC(),
		newConsumerSessionStopEC(),
	}
}

type (
	// consumerRegisterEC represents simple case for executing consumerRegister.
	consumerRegisterEC struct {
		consumer *magmasc.Consumer
	}
)

var (
	// Ensure consumerRegisterEC implements benchExecuteCase.
	_ benchExecuteCase = (*consumerRegisterEC)(nil)
)

func newConsumerRegisterEC() *consumerRegisterEC {
	return &consumerRegisterEC{
		consumer: mockConsumer(),
	}
}

func (r *consumerRegisterEC) txn() *transaction.Transaction {
	return &transaction.Transaction{
		ClientID: r.consumer.ID,
	}
}

func (r *consumerRegisterEC) funcName() string {
	return consumerRegister
}

func (r *consumerRegisterEC) input() []byte {
	return r.consumer.Encode()
}

// setupDatabaseAndState returns empty database and state context.
func (r *consumerRegisterEC) setupDatabaseAndState(
	dbDir, stateDBDir, stateLogDir string, t require.TestingT) (*gorocksdb.TransactionDB, chain.StateContextI) {

	return createStateContextAndDB(stateDBDir, stateLogDir, dbDir, nil, t)
}

func (r *consumerRegisterEC) refresh() {
	r.consumer = mockConsumer()
}

type (
	// providerRegisterEC represents simple case for executing providerRegister.
	providerRegisterEC struct {
		provider *magmasc.Provider
	}
)

var (
	// Ensure providerRegisterEC implements benchExecuteCase.
	_ benchExecuteCase = (*providerRegisterEC)(nil)
)

func newProviderRegisterEC() *providerRegisterEC {
	return &providerRegisterEC{
		provider: mockProvider(),
	}
}

func (r *providerRegisterEC) txn() *transaction.Transaction {
	return &transaction.Transaction{
		ClientID: r.provider.ID,
	}
}

func (r *providerRegisterEC) funcName() string {
	return providerRegister
}

func (r *providerRegisterEC) input() []byte {
	return r.provider.Encode()
}

// setupDatabaseAndState returns default database and state context.
func (r *providerRegisterEC) setupDatabaseAndState(
	dbDir, stateDBDir, stateLogDir string, t require.TestingT) (*gorocksdb.TransactionDB, chain.StateContextI) {

	return createStateContextAndDB(stateDBDir, stateLogDir, dbDir, nil, t)
}

func (r *providerRegisterEC) refresh() {
	r.provider = mockProvider()
}

type (
	// providerUpdateEC represents a simple case for executing providerUpdate.
	providerUpdateEC struct {
		provider *magmasc.Provider
	}
)

var (
	// Ensure providerUpdateEC implements benchExecuteCase.
	_ benchExecuteCase = (*providerUpdateEC)(nil)
)

func newProviderEC() *providerUpdateEC {
	return &providerUpdateEC{
		provider: mockProvider(),
	}
}

func (r *providerUpdateEC) txn() *transaction.Transaction {
	return &transaction.Transaction{
		ClientID: r.provider.ID,
	}
}

func (r *providerUpdateEC) funcName() string {
	return providerUpdate
}

func (r *providerUpdateEC) input() []byte {
	return r.provider.Encode()
}

// setupDatabaseAndState creates database and state context with providerUpdateEC.provider stored.
func (r *providerUpdateEC) setupDatabaseAndState(
	dbDir, stateDBDir, stateLogDir string, t require.TestingT) (*gorocksdb.TransactionDB, chain.StateContextI) {

	var (
		db, sci = createStateContextAndDB(stateDBDir, stateLogDir, dbDir, nil, t)
	)

	list, err := providersFetch(AllProvidersKey, db)
	require.NoError(t, err)

	err = list.write(magmasc.Address, r.provider, db, sci)
	require.NoError(t, err)

	return db, sci
}

func (r *providerUpdateEC) refresh() {
	r.provider = mockProvider()
}

type (
	// consumerUpdateEC represents a simple case for executing consumerUpdate.
	consumerUpdateEC struct {
		consumer *magmasc.Consumer
	}
)

var (
	// Ensure consumerUpdateEC implements benchExecuteCase.
	_ benchExecuteCase = (*consumerUpdateEC)(nil)
)

func newConsumerUpdateEC() *consumerUpdateEC {
	return &consumerUpdateEC{
		consumer: mockConsumer(),
	}
}

func (r *consumerUpdateEC) txn() *transaction.Transaction {
	return &transaction.Transaction{
		ClientID: r.consumer.ID,
	}
}

func (r *consumerUpdateEC) funcName() string {
	return consumerUpdate
}

func (r *consumerUpdateEC) input() []byte {
	return r.consumer.Encode()
}

// setupDatabaseAndState creates database and state context with consumerUpdateEC.consumer stored.
func (r *consumerUpdateEC) setupDatabaseAndState(
	dbDir, stateDBDir, stateLogDir string, t require.TestingT) (*gorocksdb.TransactionDB, chain.StateContextI) {

	var (
		db, sci = createStateContextAndDB(stateDBDir, stateLogDir, dbDir, nil, t)
	)

	list, err := consumersFetch(AllConsumersKey, db)
	require.NoError(t, err)

	err = list.write(magmasc.Address, r.consumer, db, sci)
	require.NoError(t, err)

	return db, sci
}

func (r *consumerUpdateEC) refresh() {
	r.consumer = mockConsumer()
}

type (
	// it a simple case for executing providerSessionInit.
	providerSessionInitEC struct {
		ackn *magmasc.Acknowledgment
	}
)

var (
	// Ensure providerSessionInitEC implements benchExecuteCase.
	_ benchExecuteCase = (*providerSessionInitEC)(nil)
)

func newProviderSessionInitEC() *providerSessionInitEC {
	return &providerSessionInitEC{
		ackn: mockAcknowledgment(),
	}
}

func (p providerSessionInitEC) txn() *transaction.Transaction {
	return &transaction.Transaction{
		ClientID: p.ackn.Provider.ID,
	}
}

func (p providerSessionInitEC) funcName() string {
	return providerSessionInit
}

func (p providerSessionInitEC) input() []byte {
	return p.ackn.Encode()
}

// setupDatabaseAndState creates database and state context with stored provider and consumer.
func (p providerSessionInitEC) setupDatabaseAndState(
	dbDir, stateDBDir, stateLogDir string, t require.TestingT) (*gorocksdb.TransactionDB, chain.StateContextI) {

	var (
		db, sci = createStateContextAndDB(stateDBDir, stateLogDir, dbDir, nil, t)
	)

	pList, err := providersFetch(AllProvidersKey, db)
	require.NoError(t, err)

	err = pList.write(magmasc.Address, p.ackn.Provider, db, sci)
	require.NoError(t, err)

	cList, err := consumersFetch(AllConsumersKey, db)
	require.NoError(t, err)

	err = cList.write(magmasc.Address, p.ackn.Consumer, db, sci)
	require.NoError(t, err)

	return db, sci
}

func (p providerSessionInitEC) refresh() {
	p.ackn = mockAcknowledgment()
}

type (
	// it a simple case for executing consumerSessionStart.
	consumerSessionStartEC struct {
		ackn *magmasc.Acknowledgment
	}
)

var (
	// Ensure consumerSessionStartEC implements benchExecuteCase.
	_ benchExecuteCase = (*consumerSessionStartEC)(nil)
)

func newConsumerSessionStartEC() *consumerSessionStartEC {
	return &consumerSessionStartEC{
		ackn: mockAcknowledgment(),
	}
}

func (c consumerSessionStartEC) txn() *transaction.Transaction {
	return &transaction.Transaction{
		ClientID: c.ackn.Consumer.ID,
	}
}

func (c consumerSessionStartEC) funcName() string {
	return consumerSessionStart
}

func (c consumerSessionStartEC) input() []byte {
	return c.ackn.Encode()
}

// setupDatabaseAndState creates database and state context with stored consumerSessionStartEC.ackn,
// and non empty consumer's wallet.
func (c consumerSessionStartEC) setupDatabaseAndState(
	dbDir, stateDBDir, stateLogDir string, t require.TestingT) (*gorocksdb.TransactionDB, chain.StateContextI) {

	var (
		txn = transaction.Transaction{
			ClientID:   c.ackn.Consumer.ID,
			ToClientID: c.ackn.Consumer.ID,
		}
		db, sci = createStateContextAndDB(stateDBDir, stateLogDir, dbDir, &txn, t)
	)

	_, err := sci.InsertTrieNode(nodeUID(magmasc.Address, acknowledgment, c.ackn.SessionID), c.ackn)
	require.NoError(t, err)

	_, err = sci.GetState().Insert(
		util.Path(c.ackn.Consumer.ID),
		mockState(state.Balance(c.ackn.Terms.GetAmount()), t),
	)
	require.NoError(t, err)

	return db, sci
}

func (c consumerSessionStartEC) refresh() {
	c.ackn = mockAcknowledgment()
}

type (
	// providerDataUsageEC represents a simple case for executing providerDataUsage.
	providerDataUsageEC struct {
		ackn *magmasc.Acknowledgment
	}
)

var (
	// Ensure providerDataUsageEC implements benchExecuteCase.
	_ benchExecuteCase = (*providerDataUsageEC)(nil)
)

func newProviderDataUsageEC() *providerDataUsageEC {
	return &providerDataUsageEC{
		ackn: mockAcknowledgment(),
	}
}

func (p providerDataUsageEC) txn() *transaction.Transaction {
	return &transaction.Transaction{
		ClientID: p.ackn.Provider.ID,
	}
}

func (p providerDataUsageEC) funcName() string {
	return providerDataUsage
}

func (p providerDataUsageEC) input() []byte {
	du := p.ackn.Billing.DataUsage
	du.UploadBytes++
	du.DownloadBytes++
	du.SessionTime++

	return du.Encode()
}

// setupDatabaseAndState creates database and state context with stored consumerSessionStartEC.ackn,
// and provider.
func (p providerDataUsageEC) setupDatabaseAndState(
	dbDir, stateDBDir, stateLogDir string, t require.TestingT) (*gorocksdb.TransactionDB, chain.StateContextI) {

	var (
		db, sci = createStateContextAndDB(stateDBDir, stateLogDir, dbDir, nil, t)
		tpMock  = mockTokenPool().TokenPool
	)

	p.ackn.TokenPool = &tpMock
	_, err := sci.InsertTrieNode(nodeUID(magmasc.Address, acknowledgment, p.ackn.SessionID), p.ackn)
	require.NoError(t, err)

	list, err := providersFetch(AllProvidersKey, db)
	require.NoError(t, err)

	err = list.write(magmasc.Address, p.ackn.Provider, db, sci)
	require.NoError(t, err)

	return db, sci
}

func (p providerDataUsageEC) refresh() {
	p.ackn = mockAcknowledgment()
}

type (
	// consumerSessionStopEC represents a simple case for executing consumerSessionStop.
	consumerSessionStopEC struct {
		ackn *magmasc.Acknowledgment
	}
)

var (
	// Ensure consumerSessionStopEC implements benchExecuteCase.
	_ benchExecuteCase = (*consumerSessionStopEC)(nil)
)

func newConsumerSessionStopEC() *consumerSessionStopEC {
	return &consumerSessionStopEC{
		ackn: mockAcknowledgment(),
	}
}

func (c consumerSessionStopEC) txn() *transaction.Transaction {
	return &transaction.Transaction{
		ClientID:   c.ackn.Consumer.ID,
		ToClientID: c.ackn.Consumer.ID,
	}
}

func (c consumerSessionStopEC) funcName() string {
	return consumerSessionStop
}

func (c consumerSessionStopEC) input() []byte {
	return c.ackn.Encode()
}

// setupDatabaseAndState creates database and state context with stored consumerSessionStartEC.ackn.
func (c consumerSessionStopEC) setupDatabaseAndState(
	dbDir, stateDBDir, stateLogDir string, t require.TestingT) (*gorocksdb.TransactionDB, chain.StateContextI) {

	var (
		txn = transaction.Transaction{
			ClientID:   c.ackn.Consumer.ID,
			ToClientID: c.ackn.Consumer.ID,
		}

		db, sci = createStateContextAndDB(stateDBDir, stateLogDir, dbDir, &txn, t)
	)

	c.ackn.TokenPool = &mockTokenPool().TokenPool
	_, err := sci.InsertTrieNode(nodeUID(magmasc.Address, acknowledgment, c.ackn.SessionID), c.ackn)
	require.NoError(t, err)

	return db, sci
}

func (c consumerSessionStopEC) refresh() {
	c.ackn = mockAcknowledgment()
}
