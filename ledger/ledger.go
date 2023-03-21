package ledger

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"sync"
	"time"

	"github.com/spf13/viper"
	"github.com/thetatoken/theta/store"
	"github.com/thetatoken/theta/store/kvstore"

	log "github.com/sirupsen/logrus"

	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/theta/common/result"
	"github.com/thetatoken/theta/crypto"
	"github.com/thetatoken/theta/ledger/types"
	"github.com/thetatoken/theta/store/database"

	sbc "github.com/thetatoken/thetasubchain/blockchain"
	scom "github.com/thetatoken/thetasubchain/common"
	score "github.com/thetatoken/thetasubchain/core"
	"github.com/thetatoken/thetasubchain/interchain/witness"
	sexec "github.com/thetatoken/thetasubchain/ledger/execution"
	slst "github.com/thetatoken/thetasubchain/ledger/state"
	stypes "github.com/thetatoken/thetasubchain/ledger/types"
	smp "github.com/thetatoken/thetasubchain/mempool"
)

var logger *log.Entry = log.WithFields(log.Fields{"prefix": "ledger"})

var _ score.Ledger = (*Ledger)(nil)

//
// Ledger implements the score.Ledger interface
//
type Ledger struct {
	db           database.Database
	chain        *sbc.Chain
	consensus    score.ConsensusEngine
	valMgr       score.ValidatorManager
	mempool      *smp.Mempool
	currentBlock *score.Block

	mu       *sync.RWMutex // Lock for accessing ledger state.
	state    *slst.LedgerState
	executor *sexec.Executor

	metachainWitness witness.ChainWitness
}

// NewLedger creates an instance of Ledger
func NewLedger(chainID string, db database.Database, tagger slst.Tagger, chain *sbc.Chain, consensus score.ConsensusEngine,
	valMgr score.ValidatorManager, mempool *smp.Mempool, metachainWitness witness.ChainWitness) *Ledger {
	state := slst.NewLedgerState(chainID, db, tagger)
	ledger := &Ledger{
		db:               db,
		chain:            chain,
		consensus:        consensus,
		valMgr:           valMgr,
		mempool:          mempool,
		mu:               &sync.RWMutex{},
		state:            state,
		metachainWitness: metachainWitness,
	}
	executor := sexec.NewExecutor(db, chain, state, consensus, valMgr, ledger, metachainWitness)
	ledger.SetExecutor(executor)

	return ledger
}

// SetExecutor sets the executor for the ledger
func (ledger *Ledger) SetExecutor(executor *sexec.Executor) {
	ledger.executor = executor
}

// State returns the state of the ledger
func (ledger *Ledger) State() *slst.LedgerState {
	return ledger.state
}

// GetCurrentBlock returns the block currently being processed
func (ledger *Ledger) GetCurrentBlock() *score.Block {
	return ledger.currentBlock
}

// GetDynasty returns the current dyansty
func (ledger *Ledger) GetDynasty() *big.Int {
	return ledger.state.Finalized().GetDynasty()
}

// GetScreenedSnapshot returns a snapshot of screened ledger state to query about accounts, etc.
func (ledger *Ledger) GetScreenedSnapshot() (*slst.StoreView, error) {
	ledger.mu.Lock()
	defer ledger.mu.Unlock()

	return ledger.state.Screened().Copy()
}

// GetDeliveredSnapshot returns a snapshot of delivered ledger state to query about accounts, etc.
func (ledger *Ledger) GetDeliveredSnapshot() (*slst.StoreView, error) {
	ledger.mu.Lock()
	defer ledger.mu.Unlock()

	return ledger.state.Delivered().Copy()
}

// GetFinalizedSnapshot returns a snapshot of finalized ledger state to query about accounts, etc.
func (ledger *Ledger) GetFinalizedSnapshot() (*slst.StoreView, error) {
	ledger.mu.Lock()
	defer ledger.mu.Unlock()

	return ledger.state.Finalized().Copy()
}

// GetFinalizedValidatorSet returns the validator set of the latest DIRECTLY finalized block
func (ledger *Ledger) GetFinalizedValidatorSet(blockHash common.Hash, isNext bool) (*score.ValidatorSet, error) {
	db := ledger.state.DB()
	store := kvstore.NewKVStore(db)

	var i int
	if isNext {
		i = 1
	} else {
		i = 2
	}
	for ; ; i-- {
		block, err := findBlock(store, blockHash)
		if err != nil {
			logger.Errorf("Failed to find block for VS: %v, err: %v", blockHash.Hex(), err)
			return nil, err
		}
		if block == nil {
			return nil, fmt.Errorf("Block is nil for hash %v", blockHash.Hex())
		}

		// Grandparent or root block.
		if i == 0 || block.HCC.BlockHash.IsEmpty() || block.Status.IsTrusted() {
			stateRoot := block.BlockHeader.StateHash
			storeView := slst.NewStoreView(block.Height, stateRoot, db)
			if storeView == nil {
				logger.WithFields(log.Fields{
					"block.Hash":                  blockHash.Hex(),
					"block.Height":                block.Height,
					"block.HCC.BlockHash":         block.HCC.BlockHash.Hex(),
					"block.BlockHeader.StateHash": block.BlockHeader.StateHash.Hex(),
					"block.Status.IsTrusted()":    block.Status.IsTrusted(),
				}).Panic("Failed to load state for validator pool")
			}
			vs := storeView.GetValidatorSet()
			return vs, nil
		}
		blockHash = block.HCC.BlockHash
	}

	return nil, fmt.Errorf("Failed to find a directly finalized ancestor block for %v", blockHash)
}

func (ledger *Ledger) GetTokenBankContractAddress(tokenType score.CrossChainTokenType) *common.Address {
	// storeView := ledger.state.Finalized()
	db := ledger.state.DB()
	store := kvstore.NewKVStore(db)
	blockHash := ledger.chain.Root().Hash()
	block, err := findBlock(store, blockHash)
	if err != nil {
		logger.Fatalf("Failed to find block for last processed nonce: %v, err: %v", blockHash.Hex(), err) // should not happen
	}
	if block == nil {
		logger.Fatalf("block is nil for hash %v", blockHash.Hex()) // should not happend
	}

	stateRoot := block.BlockHeader.StateHash
	storeView := slst.NewStoreView(block.Height, stateRoot, db)
	switch tokenType {
	case score.CrossChainTokenTypeTFuel:
		return storeView.GetTFuelTokenBankContractAddress()
	case score.CrossChainTokenTypeTNT20:
		return storeView.GetTNT20TokenBankContractAddress()
	case score.CrossChainTokenTypeTNT721:
		return storeView.GetTNT721TokenBankContractAddress()
	case score.CrossChainTokenTypeTNT1155:
		return storeView.GetTNT1155TokenBankContractAddress()
	default:
		return nil
	}
}

func findBlock(store store.Store, blockHash common.Hash) (*score.ExtendedBlock, error) {
	var block score.ExtendedBlock
	err := store.Get(blockHash[:], &block)
	if err != nil {
		return nil, err
	}
	return &block, nil
}

// ScreenTxUnsafe screens the given transaction without locking.
func (ledger *Ledger) ScreenTxUnsafe(rawTx common.Bytes) (res result.Result) {
	var tx types.Tx
	tx, err := stypes.TxFromBytes(rawTx)
	if err != nil {
		return result.Error("Error decoding tx: %v", err)
	}

	_, res = ledger.executor.ScreenTx(tx)
	return res
}

// ScreenTx screens the given transaction
func (ledger *Ledger) ScreenTx(rawTx common.Bytes) (txInfo *score.TxInfo, res result.Result) {
	var tx types.Tx
	tx, err := stypes.TxFromBytes(rawTx)
	if err != nil {
		return nil, result.Error("Error decoding tx: %v", err)
	}

	if ledger.shouldSkipCheckTx(tx) {
		return nil, result.Error("Unauthorized transaction, should skip").
			WithErrorCode(result.CodeUnauthorizedTx)
	}

	ledger.mu.RLock()
	defer ledger.mu.RUnlock()

	_, res = ledger.executor.ScreenTx(tx)
	if res.IsError() {
		return nil, res
	}

	txInfo, res = ledger.executor.GetTxInfo(tx)
	if res.IsError() {
		return nil, res
	}

	return txInfo, res
}

// ProposeBlockTxs collects and executes a list of transactions, which will be used to assemble the next blockl
// It also clears these transactions from the mempool.
func (ledger *Ledger) ProposeBlockTxs(block *score.Block, validatorMajorityInTheSameDynasty bool) (stateRootHash common.Hash, blockRawTxs []common.Bytes, res result.Result) {
	// Must always acquire locks in following order to avoid deadlock: mempool, ledger.
	// Otherwise, could cause deadlock since mempool.InsertTransaction() also first acquires the mempool, and then the ledger lock
	logger.Debugf("ProposeBlockTxs: Propose block transactions, block.height = %v", block.Height)
	start := time.Now()

	ledger.mempool.Lock()
	defer ledger.mempool.Unlock()

	ledger.mu.Lock()
	defer ledger.mu.Unlock()

	ledger.currentBlock = block
	defer func() { ledger.currentBlock = nil }()

	view := ledger.state.Checked()

	logger.Debugf("ProposeBlockTxs: Start adding block transactions, block.height = %v", block.Height)
	preparationTime := time.Since(start)
	start = time.Now()

	// Add special transactions
	rawTxCandidates := []common.Bytes{}
	ledger.addSpecialTransactions(block, view, &rawTxCandidates, validatorMajorityInTheSameDynasty)

	// Add regular transactions submitted by the clients
	regularRawTxs := ledger.mempool.ReapUnsafe(score.MaxNumRegularTxsPerBlock)
	for _, regularRawTx := range regularRawTxs {
		rawTxCandidates = append(rawTxCandidates, regularRawTx)
		logger.Debugf("regular raw tx %v added to block", regularRawTx)
	}

	logger.Debugf("ProposeBlockTxs: block transactions added, block.height = %v", block.Height)
	addTxsTime := time.Since(start)
	start = time.Now()

	blockRawTxs = []common.Bytes{}
	for _, rawTxCandidate := range rawTxCandidates {
		tx, err := stypes.TxFromBytes(rawTxCandidate)
		if err != nil {
			continue
		}

		_, res := ledger.executor.CheckTx(tx)
		if res.IsError() {
			logger.Errorf("Transaction check failed: errMsg = %v, tx = %v", res.Message, tx)
			continue
		}
		blockRawTxs = append(blockRawTxs, rawTxCandidate)
	}

	logger.Debugf("ProposeBlockTxs: block transactions executed, block.height = %v", block.Height)
	execTxsTime := time.Since(start)
	start = time.Now()

	stateRootHash = view.Hash()

	logger.Debugf("ProposeBlockTxs: delay update handled, block.height = %v", block.Height)
	handleDelayedUpdateTime := time.Since(start)

	logger.Debugf("ProposeBlockTxs: Done, block.height = %v, preparationTime = %v, addTxsTime = %v, execTxsTime = %v, handleDelayedUpdateTime = %v",
		block.Height, preparationTime, addTxsTime, execTxsTime, handleDelayedUpdateTime)

	return stateRootHash, blockRawTxs, result.OK
}

// ApplyBlockTxs applies the given block transactions. If any of the transactions failed, it returns
// an error immediately. If all the transactions execute successfully, it then validates the state
// root hash. If the states root hash matches the expected value, it clears the transactions from the mempool
func (ledger *Ledger) ApplyBlockTxs(block *score.Block) result.Result {
	// Must always acquire locks in following order to avoid deadlock: mempool, ledger.
	// Otherwise, could cause deadlock since mempool.InsertTransaction() also first acquires the mempool, and then the ledger lock
	logger.Debugf("ApplyBlockTxs: Apply block transactions, block.height = %v", block.Height)

	ledger.mu.Lock()
	defer ledger.mu.Unlock()

	ledger.currentBlock = block
	defer func() { ledger.currentBlock = nil }()

	blockRawTxs := ledger.currentBlock.Txs
	expectedStateRoot := ledger.currentBlock.StateHash

	view := ledger.state.Delivered()

	// currHeight := view.Height()
	// currStateRoot := view.Hash()
	extParentBlock, err := ledger.chain.FindBlock(block.Parent)
	if extParentBlock == nil || err != nil {
		panic(fmt.Sprintf("Failed to find the parent block: %v, err: %v", block.Parent.Hex(), err))
	}
	parentBlock := extParentBlock.Block
	logger.Debugf("ApplyBlockTxs: Start applying block transactions, block.height = %v", block.Height)

	hasValidatorUpdate := false
	txProcessTime := []time.Duration{}
	for _, rawTx := range blockRawTxs {
		start := time.Now()
		tx, err := stypes.TxFromBytes(rawTx)
		if err != nil {
			//ledger.resetState(currHeight, currStateRoot)
			ledger.resetState(parentBlock)
			return result.Error("Failed to parse transaction: %v", hex.EncodeToString(rawTx))
		}
		if dtx, ok := tx.(*types.DepositStakeTx); ok && dtx.Purpose == score.StakeForValidator {
			hasValidatorUpdate = true
		} else if wtx, ok := tx.(*types.WithdrawStakeTx); ok && wtx.Purpose == score.StakeForValidator {
			hasValidatorUpdate = true
		}
		_, res := ledger.executor.ExecuteTx(tx)
		if res.IsError() || res.IsUndecided() {
			//ledger.resetState(currHeight, currStateRoot)
			ledger.resetState(parentBlock)
			return res
		}
		txProcessTime = append(txProcessTime, time.Since(start))
	}

	logger.Debugf("ApplyBlockTxs: Finish applying block transactions, block.height=%v, txProcessTime=%v", block.Height, txProcessTime)

	start := time.Now()
	handleDelayedUpdateTime := time.Since(start)

	newStateRoot := view.Hash()
	if newStateRoot != expectedStateRoot {
		//ledger.resetState(currHeight, currStateRoot)
		ledger.resetState(parentBlock)
		return result.Error("State root mismatch! root: %v, exptected: %v",
			hex.EncodeToString(newStateRoot[:]),
			hex.EncodeToString(expectedStateRoot[:]))
	}

	start = time.Now()
	ledger.state.Commit() // commit to persistent storage
	commitTime := time.Since(start)

	logger.Debugf("ApplyBlockTxs: Committed state change, block.height = %v", block.Height)

	go func() {
		ledger.mempool.Lock()
		defer ledger.mempool.Unlock()

		ledger.mempool.UpdateUnsafe(blockRawTxs) // clear txs from the mempool
	}()

	logger.Debugf("ApplyBlockTxs: Cleared mempool transactions, block.height = %v", block.Height)

	logger.Debugf("ApplyBlockTxs: Done, block.height = %v, txProcessTime = %v, handleDelayedUpdateTime = %v, commitTime = %v",
		block.Height, txProcessTime, handleDelayedUpdateTime, commitTime)

	return result.OKWith(result.Info{"hasValidatorUpdate": hasValidatorUpdate})
}

// ApplyBlockTxsForChainCorrection applies all block's txs and re-calculate root hash
func (ledger *Ledger) ApplyBlockTxsForChainCorrection(block *score.Block) (common.Hash, result.Result) {
	ledger.mempool.Lock()
	defer ledger.mempool.Unlock()

	ledger.mu.Lock()
	defer ledger.mu.Unlock()

	ledger.currentBlock = block
	defer func() { ledger.currentBlock = nil }()

	blockRawTxs := ledger.currentBlock.Txs

	view := ledger.state.Delivered()

	//currHeight := view.Height()
	//currStateRoot := view.Hash()
	extParentBlock, err := ledger.chain.FindBlock(block.Parent)
	if extParentBlock == nil || err != nil {
		panic(fmt.Sprintf("Failed to find the parent block: %v, err: %v", block.Parent.Hex(), err))
	}
	parentBlock := extParentBlock.Block

	hasValidatorUpdate := false
	for _, rawTx := range blockRawTxs {
		tx, err := stypes.TxFromBytes(rawTx)
		if err != nil {
			//ledger.resetState(currHeight, currStateRoot)
			ledger.resetState(parentBlock)
			return common.Hash{}, result.Error("Failed to parse transaction: %v", hex.EncodeToString(rawTx))
		}
		if dtx, ok := tx.(*types.DepositStakeTx); ok && dtx.Purpose == score.StakeForValidator {
			hasValidatorUpdate = true
		} else if wtx, ok := tx.(*types.WithdrawStakeTx); ok && wtx.Purpose == score.StakeForValidator {
			hasValidatorUpdate = true
		}
		_, res := ledger.executor.ExecuteTx(tx)
		if res.IsError() || res.IsUndecided() {
			//ledger.resetState(currHeight, currStateRoot)
			ledger.resetState(parentBlock)
			return common.Hash{}, res
		}
	}

	ledger.state.Commit() // commit to persistent storage

	return view.Hash(), result.OKWith(result.Info{"hasValidatorUpdate": hasValidatorUpdate})
}

// PruneState attempts to prune the state up to the targetEndHeight
func (ledger *Ledger) PruneState(targetEndHeight uint64) error {
	// Permanently disabled
	return nil

	// var processedHeight uint64
	// db := ledger.State().DB()
	// kvStore := kvstore.NewKVStore(db)
	// err := kvStore.Get(state.StatePruningProgressKey(), &processedHeight)
	// if err != nil {
	// 	processedHeight = ledger.chain.Root().Height
	// }

	// pruneInterval := uint64(viper.GetInt(common.CfgStorageStatePruningInterval))
	// maxHeightsToPrune := 3 * pruneInterval // prune too many heights at once could cause hang, should catchup gradually
	// endHeight := processedHeight + maxHeightsToPrune
	// if endHeight > targetEndHeight {
	// 	endHeight = targetEndHeight
	// }

	// startHeight := processedHeight + 1
	// if endHeight < startHeight {
	// 	errMsg := fmt.Sprintf("endHeight (%v) < startHeight (%v)", endHeight, startHeight)
	// 	logger.Warnf(errMsg)
	// 	return fmt.Errorf(errMsg)
	// }

	// lastFinalizedBlock := ledger.consensus.GetLastFinalizedBlock()
	// if endHeight >= lastFinalizedBlock.Height {
	// 	errMsg := fmt.Sprintf("Can't prune at height >= %v yet", lastFinalizedBlock.Height)
	// 	logger.Warnf(errMsg)
	// 	return fmt.Errorf(errMsg)
	// }

	// // Need to save the progress before pruning -- in case the program exits during pruning (e.g. Ctrl+C),
	// // the states that are already pruned do not get pruned again
	// kvStore.Put(state.StatePruningProgressKey(), endHeight)

	// err = ledger.pruneStateForRange(startHeight, endHeight)
	// if err != nil {
	// 	logger.Warnf("Unable to pruning state: %v", err)
	// 	return err
	// }

	// return nil
}

// pruneStateForRange prunes states from startHeight to endHeight (inclusive for both end)
func (ledger *Ledger) pruneStateForRange(startHeight, endHeight uint64) error {
	logger.Infof("Prune state from height %v to %v", startHeight, endHeight)

	db := ledger.State().DB()
	consensus := ledger.consensus
	chain := ledger.chain
	lastFinalizedBlock := consensus.GetLastFinalizedBlock()

	sv := slst.NewStoreView(lastFinalizedBlock.Height, lastFinalizedBlock.BlockHeader.StateHash, db)

	stateHashMap := make(map[string]bool)
	kvStore := kvstore.NewKVStore(db)
	hl := sv.GetValidatorSetUpdateTxHeightList().Heights
	for _, height := range hl {
		// check kvstore first
		blockTrio := &score.SnapshotBlockTrio{}
		blockTrioKey := []byte(score.BlockTrioStoreKeyPrefix + strconv.FormatUint(height, 10))
		err := kvStore.Get(blockTrioKey, blockTrio)
		if err == nil {
			stateHashMap[blockTrio.First.Header.StateHash.String()] = true
			continue
		}

		if height == score.GenesisBlockHeight {
			blocks := chain.FindBlocksByHeight(score.GenesisBlockHeight)
			genesisBlock := blocks[0]
			stateHashMap[genesisBlock.StateHash.String()] = true
		} else {
			blocks := chain.FindBlocksByHeight(height)
			for _, block := range blocks {
				if block.Status.IsDirectlyFinalized() {
					stateHashMap[block.StateHash.String()] = true
					break
				}
			}
		}
	}

	for height := endHeight; height >= startHeight && height > 0; height-- {
		if common.IsCheckPointHeight(height+1) && viper.GetBool(common.CfgStorageStatePruningSkipCheckpoints) {
			logger.Infof("Skip pruning checkpoint blocks")
			continue // preserve checkpoint states
		}
		blocks := chain.FindBlocksByHeight(height)
		for idx, block := range blocks {
			if _, ok := stateHashMap[block.StateHash.String()]; !ok {
				if block.HasValidatorUpdate {
					continue
				}

				if block.Status.IsPending() || block.Status.IsInvalid() || block.Status.IsTrusted() {
					continue // This could happen if the block is stored in the chain but its
					// txs were not processed (e.g. an invalid block). In such cases the block
					// is stored in the chain, but its state trie is not saved
				}

				logger.Infof("Prune state, idx: %v, height: %v, StateHash: %v", idx, height, block.StateHash.Hex())
				_, err := db.Get(block.StateHash[:])
				if err != nil {
					logger.Errorf("StateRoot %v not found, skip pruning", block.StateHash.Hex())
					continue
				}

				sv := slst.NewStoreView(height, block.StateHash, db)
				err = sv.Prune()
				if err != nil {
					return fmt.Errorf("Failed to prune storeview at height %v, %v", height, err)
				}
			}
		}
	}

	logger.Infof("Prune state from height %v to %v completed", startHeight, endHeight)

	return nil
}

// ResetState sets the ledger state with the designated root
//func (ledger *Ledger) ResetState(height uint64, rootHash common.Hash) result.Result {
func (ledger *Ledger) ResetState(block *score.Block) result.Result {
	ledger.mu.Lock()
	defer ledger.mu.Unlock()

	//return ledger.resetState(height, rootHash)
	return ledger.resetState(block)
}

// FinalizeState sets the ledger state with the finalized root
func (ledger *Ledger) FinalizeState(height uint64, rootHash common.Hash) result.Result {
	ledger.mu.Lock()
	defer ledger.mu.Unlock()

	res := ledger.state.Finalize(height, rootHash)
	if res.IsError() {
		return result.Error("Failed to finalize state root: %v", hex.EncodeToString(rootHash[:]))
	}
	return result.OK
}

// resetState sets the ledger state with the designated root
//func (ledger *Ledger) resetState(height uint64, rootHash common.Hash) result.Result
func (ledger *Ledger) resetState(block *score.Block) result.Result {
	height := block.Height
	rootHash := block.StateHash
	logger.Debugf("Reseting state to height %v, hash %v\n", height, rootHash.Hex())

	//res := ledger.state.ResetState(height, rootHash)
	res := ledger.state.ResetState(block)
	if res.IsError() {
		return result.Error("Failed to set state root: %v", hex.EncodeToString(rootHash[:]))
	}
	return result.OK
}

// CheckTx() should skip all the transactions that can only be initiated by the validators
// i.e., if a regular user submits a coinbaseTx or slashTx, it should be skipped so it will not
// get into the mempool
func (ledger *Ledger) shouldSkipCheckTx(tx types.Tx) bool {
	switch tx.(type) {
	case *types.CoinbaseTx:
		return true
	case *types.SlashTx:
		return true
	default:
		return false
	}
}

// addSpecialTransactions adds special transactions (e.g. coinbase transaction, slash transaction) to the block
func (ledger *Ledger) addSpecialTransactions(block *score.Block, view *slst.StoreView, rawTxs *[]common.Bytes, validatorMajorityInTheSameDynasty bool) {
	if block == nil {
		logger.Warnf("addSpecialTransactions: block is nil")
		return
	}

	// Note 1: Should call GetNextProposer() with the hash of the parent block, since the current block is not fully assembled yet
	// Note 2: GetNextProposer() should use the current epoch instead of the parent's epoch, since different epochs might have different proposers
	// Note 3: Similarly, should call GetNextValidatorSet() on the hash of the parent block
	parentBlkHash := block.Parent
	proposer := ledger.valMgr.GetNextProposer(parentBlkHash, block.Epoch)
	currentValidatorSet := ledger.valMgr.GetNextValidatorSet(parentBlkHash)

	// ------- Add coinbase transaction ------- //
	ledger.addCoinbaseTx(view, &proposer, currentValidatorSet, rawTxs)

	// ------- Add subchain validator set update transaction ------- //

	// Here we add the subchain validator set update tx regardless whether the validator set has changed, since
	// the token bank contracts need to query the validator set of each dynasty
	enteringNewDynasty, newDynasty, newValidatorSet := ledger.getNewDynastyAndValidatorSet(view)
	if enteringNewDynasty && validatorMajorityInTheSameDynasty {
		ledger.addSubchainValidatorSetUpdateTx(view, &proposer, newDynasty, newValidatorSet, rawTxs)
	}
	logger.Debugf("Checking whether to add subchain validator update transactions: validatorMajorityInTheSameDynasty: %v, enteringNewDynasty: %v, newDynasty: %v, newValidatorSet: %v",
		validatorMajorityInTheSameDynasty, enteringNewDynasty, newDynasty, newValidatorSet)

	// hasValidatorSetUpdate, newDynasty, newValidatorSet := ledger.hasSubchainValidatorSetUpdate(currentValidatorSet, view)
	// logger.Debugf("Add special transactions: validatorMajorityInTheSameDynasty: %v, hasValidatorSetUpdate: %v, newDynasty: %v, newValidatorSet: %v",
	// 	validatorMajorityInTheSameDynasty, hasValidatorSetUpdate, newDynasty, newValidatorSet)
	// if validatorMajorityInTheSameDynasty && hasValidatorSetUpdate {
	// 	ledger.addSubchainValidatorSetUpdateTx(view, &proposer, newDynasty, newValidatorSet, rawTxs)
	// }
}

func (ledger *Ledger) getNewDynastyAndValidatorSet(view *slst.StoreView) (enteringNewDynasty bool, newDynasty *big.Int, newValidatorSet *score.ValidatorSet) {
	// Note that here we get the "current" dynasty from the view, even though the block containing the
	// validator set update is NOT finalized yet (typically needs two blocks). Otherwise, if we instead
	// retrieve the dynasty from the "finalized" validator set, the code could issue validator set update
	// txs for two consecutive blocks, where the second tx will be rejected.
	currentDynasty := view.GetDynasty()
	mainchainBlockHeight, err := ledger.metachainWitness.GetMainchainBlockHeight()
	if err != nil {
		logger.Warn("Failed to get mainchain block number when checking validator set updates, err: %v", err)
		return false, nil, nil
	}

	registrationMainchainHeight, err := ledger.metachainWitness.GetSubchainRegistrationHeight()
	if err != nil {
		logger.Warnf("Failed to get subchain registration height: %v", err)
		return false, nil, nil
	}

	registrationDynasty := scom.CalculateDynasty(registrationMainchainHeight)
	witnessedDynasty := scom.CalculateDynasty(mainchainBlockHeight)

	logger.Debugf("currentDynasty: %v, witnessedDynasty: %v, egistrationDynasty: %v, registrationMainchainHeight: %v", currentDynasty, witnessedDynasty, registrationDynasty, registrationMainchainHeight)

	if witnessedDynasty.Cmp(registrationDynasty) == 0 {
		// Specicial case: When the node is in the dynasty where the subchain was registered. The subchain should use the ValidatorSet
		//                 specified in the snapshot (in stead of querying the main chain). Yet, the node needs to update the dynasty
		//                 in the ViewStore to witnessedDynasty if witnessedDynasty is not 0

		if witnessedDynasty.Cmp(common.Big0) != 0 && currentDynasty.Cmp(common.Big0) == 0 { // the first validator set update
			// update the dynast to witnessedDynasty, and continue to use the validator set specified in the snapshot
			return true, witnessedDynasty, view.GetValidatorSet()
		} else {
			// witnessedDynasty is 0, so no update is needed, continue to use the ValidatorSet loaded from the snapshot
			return false, nil, nil
		}
	}

	witnessedValidatorSet, err := ledger.metachainWitness.GetValidatorSetByDynasty(witnessedDynasty)
	if err != nil {
		logger.Warnf("Failed to get validator set by dynasty %v when checking validator set updates, err: %v", witnessedDynasty, err)
		return false, nil, nil
	}

	if witnessedDynasty.Cmp(currentDynasty) <= 0 {
		return false, nil, nil
	}

	// at this point: witnessedDynasty >= currentDynasty + 1, we are entering a new dynasty

	validatorSetInView := view.GetValidatorSet()

	logger.Debugf("block height: %v", view.GetBlockHeight())
	logger.Debugf("witnessedValidatorSet: %v", witnessedValidatorSet)
	//logger.Debugf("currentValidatorSet  : %v", currentValidatorSet)
	logger.Debugf("validatorSetInView   : %v", validatorSetInView)
	logger.Debugf("currentDynasty       : %v", currentDynasty)
	logger.Debugf("witnessedDynasty     : %v", witnessedDynasty)

	return true, witnessedDynasty, witnessedValidatorSet
}

// Not used
func (ledger *Ledger) hasSubchainValidatorSetUpdate(currentValidatorSet *score.ValidatorSet, view *slst.StoreView) (bool, *big.Int, *score.ValidatorSet) {
	currentDynasty := currentValidatorSet.Dynasty()
	mainchainBlockHeight, err := ledger.metachainWitness.GetMainchainBlockHeight()
	if err != nil {
		logger.Warn("Failed to get mainchain block number when checking validator set updates, err: %v", err)
		return false, nil, nil
	}

	witnessedDynasty := scom.CalculateDynasty(mainchainBlockHeight)
	if witnessedDynasty.Cmp(currentDynasty) <= 0 {
		return false, nil, nil
	}
	// at this point: witnessedDynasty >= currentDynasty + 1

	witnessedValidatorSet, err := ledger.metachainWitness.GetValidatorSetByDynasty(witnessedDynasty)
	if err != nil {
		logger.Warnf("Failed to get validator set by dynasty %v when checking validator set updates, err: %v", witnessedDynasty, err)
		return false, nil, nil
	}

	if currentValidatorSet.Equals(witnessedValidatorSet) {
		return false, nil, nil
	}

	// Need to compare the validator set in the view against the witnessed validator set. This is
	// because `currentValidatorSet` refers to the validator set of the last finalized block, after
	// which there might have been a block containing the subchainValidatorSetUpdateTx. In this case,
	// we do not need to add the subchainValidatorSetUpdateTx again to the block.
	validatorSetInView := view.GetValidatorSet()
	if validatorSetInView.Equals(witnessedValidatorSet) {
		return false, nil, nil
	}

	logger.Debugf("block height: %v", view.GetBlockHeight())
	logger.Debugf("witnessedValidatorSet: %v", witnessedValidatorSet)
	logger.Debugf("currentValidatorSet  : %v", currentValidatorSet)
	logger.Debugf("validatorSetInView   : %v", validatorSetInView)

	return true, witnessedDynasty, witnessedValidatorSet
}

// addCoinbaseTx adds a Coinbase transaction
func (ledger *Ledger) addCoinbaseTx(view *slst.StoreView, proposer *score.Validator,
	validatorSet *score.ValidatorSet, rawTxs *[]common.Bytes) {
	proposerAddress := proposer.Address
	proposerTxIn := types.TxInput{
		Address: proposerAddress,
	}

	coinbaseTxOutputs := []types.TxOutput{}
	coinbaseTx := &types.CoinbaseTx{
		Proposer:    proposerTxIn,
		Outputs:     coinbaseTxOutputs,
		BlockHeight: ledger.state.Height(),
	}

	signature, err := ledger.signTransaction(coinbaseTx)
	if err != nil {
		logger.Errorf("Failed to add coinbase transaction: %v", err)
		return
	}
	coinbaseTx.SetSignature(proposerAddress, signature)
	coinbaseTxBytes, err := stypes.TxToBytes(coinbaseTx)
	if err != nil {
		logger.Errorf("Failed to add coinbase transaction: %v", err)
		return
	}

	*rawTxs = append(*rawTxs, coinbaseTxBytes)
	logger.Debugf("Added coinbase transction: tx: %v, bytes: %v", coinbaseTx, hex.EncodeToString(coinbaseTxBytes))
}

// addSubchainValidatorSetUpdateTx adds a validator update transaction
func (ledger *Ledger) addSubchainValidatorSetUpdateTx(view *slst.StoreView, proposer *score.Validator,
	newDynasty *big.Int, newValidatorSet *score.ValidatorSet, rawTxs *[]common.Bytes) {
	proposerAccount := view.GetAccount(proposer.Address)
	if proposerAccount == nil {
		// should not happen, since the the validator set update tx shouuld create the propser account if it does not exist
		logger.Fatalf("Failed to get proposer account: %v", proposer.Address)
	}

	proposerAddress := proposer.Address
	proposerTxIn := types.TxInput{
		Address:  proposerAddress,
		Sequence: proposerAccount.Sequence + 1,
	}

	subchainValidatorSetUpdateTx := &stypes.SubchainValidatorSetUpdateTx{
		Proposer:   proposerTxIn,
		Dynasty:    newDynasty,
		Validators: newValidatorSet.Validators(),
	}

	signature, err := ledger.signTransaction(subchainValidatorSetUpdateTx)
	if err != nil {
		logger.Fatalf("Failed to add subchain validator set update transaction: %v", err) // do not expect this to happen
		return
	}
	subchainValidatorSetUpdateTx.SetSignature(proposerAddress, signature)
	subchainValidatorSetUpdateTxBytes, err := stypes.TxToBytes(subchainValidatorSetUpdateTx)
	if err != nil {
		logger.Fatalf("Failed to serialize subchain validator set update transaction: %v", err) // do not expect this to happen
		return
	}

	*rawTxs = append(*rawTxs, subchainValidatorSetUpdateTxBytes)
	logger.Infof("Added subchain validator set update transction: tx: %v", subchainValidatorSetUpdateTx)
	logger.Debugf("Subchain validator set update transction bytes: %v", hex.EncodeToString(subchainValidatorSetUpdateTxBytes))
}

// signTransaction signs the given transaction
func (ledger *Ledger) signTransaction(tx types.Tx) (*crypto.Signature, error) {
	chainID := ledger.state.GetChainID()
	signBytes := tx.SignBytes(chainID)
	signature, err := ledger.consensus.PrivateKey().Sign(signBytes)
	if err != nil {
		return nil, err
	}
	return signature, nil
}
