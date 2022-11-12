package state

import (
	"bytes"
	"fmt"
	"math/big"

	log "github.com/sirupsen/logrus"
	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/theta/crypto"
	"github.com/thetatoken/theta/ledger/types"
	"github.com/thetatoken/theta/rlp"
	"github.com/thetatoken/theta/store/database"
	score "github.com/thetatoken/thetasubchain/core"
	streestore "github.com/thetatoken/thetasubchain/store/treestore"
)

var logger *log.Entry = log.WithFields(log.Fields{"prefix": "ledger"})

//
// ------------------------- StoreView -------------------------
//

type StoreView struct {
	height uint64 // block height
	store  *streestore.TreeStore

	coinbaseTransactinProcessed             bool
	subchainValidatorSetTransactinProcessed bool
	slashIntents                            []types.SlashIntent
	refund                                  uint64                 // Gas refund during smart contract execution
	logs                                    []*types.Log           // Temporary store of events during smart contract execution
	balanceChanges                          []*types.BalanceChange // Temporary store of balance changes during smart contract execution
}

// NewStoreView creates an instance of the StoreView
func NewStoreView(height uint64, root common.Hash, db database.Database) *StoreView {
	store := streestore.NewTreeStore(root, db)
	if store == nil {
		return nil
	}

	sv := &StoreView{
		height:       height,
		store:        store,
		slashIntents: []types.SlashIntent{},
		refund:       0,
	}
	return sv
}

// Copy returns a copy of the StoreView
func (sv *StoreView) Copy() (*StoreView, error) {
	copiedStore, err := sv.store.Copy()
	if err != nil {
		return nil, err
	}
	copiedStoreView := &StoreView{
		height:       sv.height,
		store:        copiedStore,
		slashIntents: []types.SlashIntent{},
		refund:       0,
	}
	return copiedStoreView, nil
}

// GetDB returns the underlying database.
func (sv *StoreView) GetDB() database.Database {
	return sv.store.GetDB()
}

// Hash returns the root hash of the tree store
func (sv *StoreView) Hash() common.Hash {
	return sv.store.Hash()
}

// Height returns the block height corresponding to the stored state
func (sv *StoreView) Height() uint64 {
	return sv.height
}

// IncrementHeight increments the block height by 1
func (sv *StoreView) IncrementHeight() {
	sv.height++
}

// Save saves the StoreView to the persistent storage, and return the root hash
func (sv *StoreView) Save() common.Hash {
	rootHash, err := sv.store.Commit()

	logger.Debugf("Commit to data store, height: %v, rootHash: %v", sv.height+1, rootHash.Hex())

	if err != nil {
		log.Panicf("Failed to save the StoreView: %v", err)
	}
	return rootHash
}

// Get returns the value corresponding to the key
func (sv *StoreView) Get(key common.Bytes) common.Bytes {
	value := sv.store.Get(key)
	return value
}

// Traverse traverses the trie and calls cb callback func on every key/value pair
// with key having prefix
func (sv *StoreView) Traverse(prefix common.Bytes, cb func(k, v common.Bytes) bool) bool {
	return sv.store.Traverse(prefix, cb)
}

func (sv *StoreView) ProveValidatorSet(vspKey []byte, vsp *score.ValidatorSetProof) error {
	return sv.store.ProveValidatorSet(vspKey, vsp)
}

// Delete removes the value corresponding to the key
func (sv *StoreView) Delete(key common.Bytes) {
	sv.store.Delete(key)
}

// Set returns the value corresponding to the key
func (sv *StoreView) Set(key common.Bytes, value common.Bytes) {
	sv.store.Set(key, value)
}

// AddSlashIntent adds slashIntent
func (sv *StoreView) AddSlashIntent(slashIntent types.SlashIntent) {
	sv.slashIntents = append(sv.slashIntents, slashIntent)
}

// GetSlashIntents retrieves all the slashIntents
func (sv *StoreView) GetSlashIntents() []types.SlashIntent {
	return sv.slashIntents
}

// ClearSlashIntents clears all the slashIntents
func (sv *StoreView) ClearSlashIntents() {
	sv.slashIntents = []types.SlashIntent{}
}

// CoinbaseTransactinProcessed returns whether the coinbase transaction for the current block has been processed
func (sv *StoreView) CoinbaseTransactinProcessed() bool {
	return sv.coinbaseTransactinProcessed
}

// SetCoinbaseTransactionProcessed sets whether the coinbase transaction for the current block has been processed
func (sv *StoreView) SetCoinbaseTransactionProcessed(processed bool) {
	sv.coinbaseTransactinProcessed = processed
}

// SubchainValidatorSetTransactionProcessed returns whether the validator set update transaction for the current block has been processed
func (sv *StoreView) SubchainValidatorSetTransactionProcessed() bool {
	return sv.subchainValidatorSetTransactinProcessed
}

// SetSubchainValidatorSetTransactionProcessed sets whether the validator set update transaction for the current block has been processed
func (sv *StoreView) SetSubchainValidatorSetTransactionProcessed(processed bool) {
	sv.subchainValidatorSetTransactinProcessed = processed
}

// GetAccount returns an account.
func (sv *StoreView) GetAccount(addr common.Address) *types.Account {
	data := sv.Get(AccountKey(addr))
	if data == nil || len(data) == 0 {
		return nil
	}
	acc := &types.Account{}
	err := types.FromBytes(data, acc)
	if err != nil {
		log.Panicf("Error reading account %X error: %v",
			data, err.Error())
	}
	return acc
}

// // SetAccount sets an account.
// func (sv *StoreView) SetAccount(addr common.Address, acc *types.Account) {
// 	accBytes, err := types.ToBytes(acc)
// 	if err != nil {
// 		log.Panicf("Error writing account %v error: %v",
// 			acc, err.Error())
// 	}
// 	sv.Set(AccountKey(addr), accBytes)
// }

// SetAccount sets an account.
func (sv *StoreView) SetAccount(addr common.Address, acc *types.Account) {
	sv.setAccount(addr, acc, true)
}

func (sv *StoreView) setAccountWithoutStateTreeRefCountUpdate(addr common.Address, acc *types.Account) {
	sv.setAccount(addr, acc, false)
}

func (sv *StoreView) setAccount(addr common.Address, acc *types.Account, updateRefCountForAccountStateTree bool) {
	accBytes, err := types.ToBytes(acc)
	if err != nil {
		log.Panicf("Error writing account %v error: %v",
			acc, err.Error())
	}
	sv.Set(AccountKey(addr), accBytes)

	if !updateRefCountForAccountStateTree {
		return
	}

	if (acc == nil || acc.Root == common.Hash{}) || (acc.Root == score.EmptyRootHash) {
		return
	}

	tree := sv.getAccountStorage(acc)
	_, err = tree.Commit() // update the reference count of the account state trie root
	if err != nil {
		log.Panic(err)
	}
}

// DeleteAccount deletes an account.
func (sv *StoreView) DeleteAccount(addr common.Address) {
	sv.Delete(AccountKey(addr))
}

// GetDynasty gets the dynasty associated with the view
func (sv *StoreView) GetDynasty() *big.Int {
	data := sv.Get(CurrentValidatorSetKey())
	if data == nil || len(data) == 0 {
		return nil
	}
	vs := &score.ValidatorSet{}
	err := types.FromBytes(data, vs)
	if err != nil {
		log.Panicf("Error reading validator set %X, error: %v",
			data, err.Error())
	}
	dynasty := vs.Dynasty() // retrieve the dynasty from the validator set
	return dynasty
}

// GetValidatorSet gets the validator set.
func (sv *StoreView) GetValidatorSet() *score.ValidatorSet {
	data := sv.Get(CurrentValidatorSetKey())
	if data == nil || len(data) == 0 {
		return nil
	}
	vs := &score.ValidatorSet{}
	err := types.FromBytes(data, vs)
	if err != nil {
		log.Panicf("Error reading validator set %X, error: %v",
			data, err.Error())
	}
	return vs
}

// UpdateValidatorSet updates the validator set.
func (sv *StoreView) UpdateValidatorSet(chainID *big.Int, vs *score.ValidatorSet) {
	logger.Debugf("Update validator set, chainID: %v, valset: %v", chainID, vs)

	vsBytes, err := types.ToBytes(vs)
	if err != nil {
		log.Panicf("Error writing validator set %v, error: %v",
			vs, err.Error())
	}
	sv.Set(CurrentValidatorSetKey(), vsBytes)
	sv.Set(ValidatorSetForChainDuringDynastyKey(chainID, vs.Dynasty()), vsBytes)
}

func (sv *StoreView) GetValidatorSetForChainDuringDynasty(chainID *big.Int, dynasty *big.Int) *score.ValidatorSet {
	data := sv.Get(ValidatorSetForChainDuringDynastyKey(chainID, dynasty))
	if data == nil || len(data) == 0 {
		return nil
	}
	vs := &score.ValidatorSet{}
	err := types.FromBytes(data, vs)
	if err != nil {
		log.Panicf("Error reading validator set %X, error: %v",
			data, err.Error())
	}
	return vs
}

// GetChainRegistrarContractAddress gets chain registrar contract address
func (sv *StoreView) GetChainRegistrarContractAddress() *common.Address {
	data := sv.Get(ChainRegistrarContractAddressKey())
	if len(data) == 0 {
		return nil
	}
	crca := &common.Address{}
	err := types.FromBytes(data, crca)
	if err != nil {
		log.Panicf("Error reading chain registrar contract address %X, error: %v",
			data, err.Error())
	}
	return crca
}

// GetTFuelTokenBankContractAddress gets the TFuel token bank contract address.
func (sv *StoreView) GetTFuelTokenBankContractAddress() *common.Address {
	data := sv.Get(TFuelTokenBankContractAddressKey())
	if len(data) == 0 {
		return nil
	}
	tbca := &common.Address{}
	err := types.FromBytes(data, tbca)
	if err != nil {
		log.Panicf("Error reading TFuel token bank contract address %X, error: %v",
			data, err.Error())
	}
	return tbca
}

// GetTNT20TokenBankContractAddress gets the TFuel token bank contract address.
func (sv *StoreView) GetTNT20TokenBankContractAddress() *common.Address {
	data := sv.Get(TNT20TokenBankContractAddressKey())
	if len(data) == 0 {
		return nil
	}
	tbca := &common.Address{}
	err := types.FromBytes(data, tbca)
	if err != nil {
		log.Panicf("Error reading TNT20 token bank contract address %X, error: %v",
			data, err.Error())
	}
	return tbca
}

// GetTNT721TokenBankContractAddress gets the TFuel token bank contract address.
func (sv *StoreView) GetTNT721TokenBankContractAddress() *common.Address {
	data := sv.Get(TNT721TokenBankContractAddressKey())
	if len(data) == 0 {
		return nil
	}
	tbca := &common.Address{}
	err := types.FromBytes(data, tbca)
	if err != nil {
		log.Panicf("Error reading TNT721 token bank contract address %X, error: %v",
			data, err.Error())
	}
	return tbca
}

// GetTNT1155TokenBankContractAddress gets the TNT1155 token bank contract address.
func (sv *StoreView) GetTNT1155TokenBankContractAddress() *common.Address {
	data := sv.Get(TNT1155TokenBankContractAddressKey())
	if len(data) == 0 {
		return nil
	}
	tbca := &common.Address{}
	err := types.FromBytes(data, tbca)
	if err != nil {
		log.Panicf("Error reading TNT1155 token bank contract address %X, error: %v",
			data, err.Error())
	}
	return tbca
}

// GetBalanceCheckerContractAddress gets the TNT1155 token bank contract address.
func (sv *StoreView) GetBalanceCheckerContractAddress() *common.Address {
	data := sv.Get(BalanceCheckerContractAddressKey())
	if len(data) == 0 {
		return nil
	}
	bbca := &common.Address{}
	err := types.FromBytes(data, bbca)
	if err != nil {
		log.Panicf("Error reading balance checker contract address %X, error: %v",
			data, err.Error())
	}
	return bbca
}

// GetValidatorSetUpdateTxHeightList gets the heights of blocks that contain stake related transactions
func (sv *StoreView) GetValidatorSetUpdateTxHeightList() *types.HeightList {
	data := sv.Get(ValidatorSetUpdateTxHeightListKey())
	if data == nil || len(data) == 0 {
		return nil
	}

	hl := &types.HeightList{}
	err := types.FromBytes(data, hl)
	if err != nil {
		log.Panicf("Error reading height list %X, error: %v",
			data, err.Error())
	}
	return hl
}

// UpdateValidatorSetUpdateTxHeightList updates the heights of blocks that contain stake related transactions
func (sv *StoreView) UpdateValidatorSetUpdateTxHeightList(hl *types.HeightList) {
	hlBytes, err := types.ToBytes(hl)
	if err != nil {
		log.Panicf("Error writing height list %v, error: %v",
			hl, err.Error())
	}
	sv.Set(ValidatorSetUpdateTxHeightListKey(), hlBytes)
}

type StakeWithHolder struct {
	Holder common.Address
	Stake  score.Stake
}

func (sv *StoreView) GetStore() *streestore.TreeStore {
	return sv.store
}

func (sv *StoreView) ResetLogs() {
	sv.logs = []*types.Log{}
}

func (sv *StoreView) PopLogs() []*types.Log {
	ret := sv.logs
	sv.ResetLogs()
	return ret
}

func (sv *StoreView) ResetBalanceChanges() {
	sv.balanceChanges = []*types.BalanceChange{}
}

func (sv *StoreView) PopBalanceChanges() []*types.BalanceChange {
	ret := sv.balanceChanges
	sv.ResetBalanceChanges()
	return ret
}

//
// ---------- Implement svm.StateDB interface -----------
//

func (sv *StoreView) CreateAccount(addr common.Address) {
	account := types.NewAccount(addr)
	sv.SetAccount(addr, account)
}

func (sv *StoreView) GetOrCreateAccount(addr common.Address) *types.Account {
	account := sv.GetAccount(addr)
	if account != nil {
		return account
	}
	return types.NewAccount(addr)
}

func (sv *StoreView) CreateAccountWithPreviousBalance(addr common.Address) {
	account := types.NewAccount(addr)

	existingAccount := sv.GetAccount(addr)
	if existingAccount != nil { // only copy over the account balance, reset other fields including the account sequence
		account.Balance = existingAccount.Balance.NoNil()
	}

	sv.SetAccount(addr, account)
}

func (sv *StoreView) SubBalance(addr common.Address, amount *big.Int) {
	if amount.Sign() == 0 {
		return
	}
	account := sv.GetAccount(addr)
	if account == nil {
		panic(fmt.Sprintf("Account for %v does not exist!", addr))
	}
	account.Balance = account.Balance.NoNil()
	account.Balance.TFuelWei.Sub(account.Balance.TFuelWei, amount)
	sv.SetAccount(addr, account)

	sv.addBalanceChange(&types.BalanceChange{
		Address:    addr,
		TokenType:  1,
		IsNegative: true,
		Delta:      new(big.Int).Set(amount),
	})
}

func (sv *StoreView) AddBalance(addr common.Address, amount *big.Int) {
	if amount.Sign() == 0 {
		return
	}
	account := sv.GetOrCreateAccount(addr)
	account.Balance = account.Balance.NoNil()
	account.Balance.TFuelWei.Add(account.Balance.TFuelWei, amount)
	sv.SetAccount(addr, account)

	sv.addBalanceChange(&types.BalanceChange{
		Address:    addr,
		TokenType:  1,
		IsNegative: false,
		Delta:      new(big.Int).Set(amount),
	})
}

func (sv *StoreView) GetBalance(addr common.Address) *big.Int {
	return sv.GetOrCreateAccount(addr).Balance.TFuelWei
}

func (sv *StoreView) SubThetaBalance(addr common.Address, amount *big.Int) {
	if amount.Sign() == 0 {
		return
	}
	account := sv.GetAccount(addr)
	if account == nil {
		panic(fmt.Sprintf("Account for %v does not exist!", addr))
	}
	account.Balance = account.Balance.NoNil()
	account.Balance.ThetaWei.Sub(account.Balance.ThetaWei, amount)
	sv.SetAccount(addr, account)

	sv.addBalanceChange(&types.BalanceChange{
		Address:    addr,
		TokenType:  0,
		IsNegative: true,
		Delta:      new(big.Int).Set(amount),
	})
}

func (sv *StoreView) AddThetaBalance(addr common.Address, amount *big.Int) {
	if amount.Sign() == 0 {
		return
	}
	account := sv.GetOrCreateAccount(addr)
	account.Balance = account.Balance.NoNil()
	account.Balance.ThetaWei.Add(account.Balance.ThetaWei, amount)
	sv.SetAccount(addr, account)

	sv.addBalanceChange(&types.BalanceChange{
		Address:    addr,
		TokenType:  0,
		IsNegative: false,
		Delta:      new(big.Int).Set(amount),
	})
}

// GetThetaBalance returns the ThetaWei balance of the given address
func (sv *StoreView) GetThetaBalance(addr common.Address) *big.Int {
	return sv.GetOrCreateAccount(addr).Balance.ThetaWei
}

// GetThetaStake returns the total amount of ThetaWei the address staked to validators
func (sv *StoreView) GetThetaStake(addr common.Address) *big.Int {
	totalStake := big.NewInt(0) // subchain does not support native Theta, hence we always return 0
	return totalStake
}

func (sv *StoreView) GetNonce(addr common.Address) uint64 {
	return sv.GetOrCreateAccount(addr).Sequence
}

func (sv *StoreView) SetNonce(addr common.Address, nonce uint64) {
	account := sv.GetOrCreateAccount(addr)
	account.Sequence = nonce
	sv.SetAccount(addr, account)
}

func (sv *StoreView) GetCodeHash(addr common.Address) common.Hash {
	account := sv.GetAccount(addr)
	if account == nil {
		return common.Hash{}
	}
	return account.CodeHash
}

func (sv *StoreView) GetCode(addr common.Address) []byte {
	account := sv.GetAccount(addr)
	if account == nil {
		return nil
	}
	if (account.CodeHash == types.EmptyCodeHash) || (account.CodeHash == score.SuicidedCodeHash) {
		return nil
	}
	return sv.GetCodeByHash(account.CodeHash)
}

func (sv *StoreView) GetCodeByHash(codeHash common.Hash) []byte {
	if codeHash == score.SuicidedCodeHash {
		return nil
	}
	codeKey := CodeKey(codeHash[:])
	return sv.Get(codeKey)
}

func (sv *StoreView) SetCode(addr common.Address, code []byte) {
	account := sv.GetOrCreateAccount(addr)
	codeHash := crypto.Keccak256Hash(code)
	account.CodeHash = codeHash
	sv.Set(CodeKey(account.CodeHash[:]), code)
	sv.SetAccount(addr, account)
}

func (sv *StoreView) GetCodeSize(addr common.Address) int {
	return len(sv.GetCode(addr))
}

func (sv *StoreView) AddRefund(gas uint64) {
	sv.refund += gas
}

func (sv *StoreView) SubRefund(gas uint64) {
	if gas > sv.refund {
		log.Panic("Refund counter below zero")
	}
	sv.refund -= gas
}

func (sv *StoreView) GetRefund() uint64 {
	return sv.refund
}

func (sv *StoreView) ResetRefund() {
	sv.refund = 0
}

func (sv *StoreView) GetBlockHeight() uint64 {
	blockHeight := sv.height + 1
	return blockHeight
}

func (sv *StoreView) GetCommittedState(addr common.Address, key common.Hash) common.Hash {
	return sv.GetState(addr, key)
}

func (sv *StoreView) getAccountStorage(account *types.Account) *streestore.TreeStore {
	return streestore.NewTreeStore(account.Root, sv.store.GetDB())
}

func (sv *StoreView) GetState(addr common.Address, key common.Hash) common.Hash {
	account := sv.GetAccount(addr)
	if account == nil {
		return common.Hash{}
	}
	logger.Debugf("StoreView.GetState, address: %v, account.root: %v, key: %v", addr, account.Root.Hex(), key.Hex())

	enc, err := sv.getAccountStorage(account).TryGet(key[:])
	if err != nil {
		log.Panic(err)
	}
	if len(enc) > 0 {
		_, content, _, err := rlp.Split(enc)
		if err != nil {
			log.Panic(err)
		}
		return common.BytesToHash(content)
	}
	return common.Hash{}
}

func (sv *StoreView) SetState(addr common.Address, key, val common.Hash) {
	account := sv.GetAccount(addr)
	if account == nil {
		account = types.NewAccount(addr)
	}
	tree := sv.getAccountStorage(account)
	if (val == common.Hash{}) {
		tree.TryDelete(key[:])
		root, err := tree.Commit()
		if err != nil {
			log.Panic(err)
		}
		account.Root = root
		sv.setAccountWithoutStateTreeRefCountUpdate(addr, account) // The ref counts of the state tree already got updated above
		logger.Debugf("StoreView.SetState, address: %v, account.root: %v, key: %v, val: %v", addr.Hex(), root.Hex(), key.Hex(), val.Hex())
		return
	}
	// Encoding []byte cannot fail, ok to ignore the error.
	v, _ := rlp.EncodeToBytes(bytes.TrimLeft(val[:], "\x00"))
	tree.TryUpdate(key[:], v)
	root, err := tree.Commit()
	if err != nil {
		log.Panic(err)
	}

	account.Root = root
	sv.setAccountWithoutStateTreeRefCountUpdate(addr, account) // The ref counts of the state tree already got updated above

	logger.Debugf("StoreView.SetState, address: %v, account.root: %v, key: %v, val: %v", addr.Hex(), root.Hex(), key.Hex(), val.Hex())
}

func (sv *StoreView) Suicide(addr common.Address) bool {
	account := sv.GetAccount(addr)
	if account == nil {
		return false
	}
	account.CodeHash = score.SuicidedCodeHash

	sv.addBalanceChange(&types.BalanceChange{
		Address:    addr,
		TokenType:  1,
		IsNegative: true,
		Delta:      new(big.Int).Set(account.Balance.TFuelWei),
	})

	account.Balance.TFuelWei = big.NewInt(0)
	sv.SetAccount(addr, account)
	return true
}

func (sv *StoreView) HasSuicided(addr common.Address) bool {
	account := sv.GetAccount(addr)
	if account == nil {
		return true
	}
	hasSuicided := (account.CodeHash == score.SuicidedCodeHash)
	return hasSuicided
}

// Exist reports whether the given account exists in state.
// Notably this should also return true for suicided accounts.
func (sv *StoreView) Exist(addr common.Address) bool {
	account := sv.GetAccount(addr)
	return account != nil
}

// Empty returns whether the given account is empty. Empty
// is defined according to EIP161 (balance = nonce = code = 0).
func (sv *StoreView) Empty(addr common.Address) bool {
	account := sv.GetAccount(addr)
	if account == nil {
		return true
	}
	return account.Sequence == 0 &&
		account.CodeHash == types.EmptyCodeHash &&
		account.Balance.IsZero()
}

func (sv *StoreView) RevertToSnapshot(root common.Hash) {
	var err error
	sv.store, err = sv.store.Revert(root) // revert to one of the previous roots
	if err != nil {
		log.Panic(err)
	}
}

func (sv *StoreView) Snapshot() common.Hash {
	sv.store.Trie.Commit(nil) // Needs to commit to the in-memory trie DB
	return sv.store.Hash()
}

func (sv *StoreView) Prune() error {
	err := sv.store.Prune(func(node []byte) bool {
		account := &types.Account{}
		err := types.FromBytes(node, account)
		if err != nil {
			return false
		}
		if (account.Root == (common.Hash{})) || (account.Root == score.EmptyRootHash) {
			return false
		}
		storage := sv.getAccountStorage(account)
		logger.Debugf("StoreView.Prune, address: %v, account.root: %v", account.Address, account.Root.Hex())

		err = storage.Prune(nil)
		if err != nil {
			logger.Errorf("Failed to prune storage for account %v", account)
			return false
		}
		return true
	})
	if err != nil {
		return fmt.Errorf("Failed to prune store view, %v", err)
	}
	return nil
}

func (sv *StoreView) AddLog(l *types.Log) {
	sv.logs = append(sv.logs, l)
}

func (sv *StoreView) addBalanceChange(bc *types.BalanceChange) {
	sv.balanceChanges = append(sv.balanceChanges, bc)
}
