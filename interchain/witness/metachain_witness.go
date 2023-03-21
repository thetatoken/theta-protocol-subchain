package witness

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	// "github.com/thetatoken/theta/crypto"

	scom "github.com/thetatoken/thetasubchain/common"
	score "github.com/thetatoken/thetasubchain/core"
	scta "github.com/thetatoken/thetasubchain/interchain/contracts/accessors"
	siu "github.com/thetatoken/thetasubchain/interchain/utils"

	// "github.com/thetatoken/thetasubchain/eth/abi/bind"
	//"github.com/ethereum/go-ethereum/common"
	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/theta/store"
	"github.com/thetatoken/theta/store/database"

	ec "github.com/thetatoken/thetasubchain/eth/ethclient"
	//ec "github.com/ethereum/go-ethereum/ethclient"
)

var logger *log.Entry = log.WithFields(log.Fields{"prefix": "witness"})

type MetachainWitness struct {
	updateTicker   *time.Ticker
	updateInterval int
	witnessState   *metachainWitnessState

	queryTopics string
	// The mainchain
	mainchainID                   *big.Int
	mainchainEthRpcUrl            string
	mainchainEthRpcClient         *ec.Client
	witnessedDynasty              *big.Int
	chainRegistrarOnMainchain     *scta.ChainRegistrarOnMainchain // the ChainRegistrarOnMainchain contract deployed on the mainchain
	mainchainTFuelTokenBankAddr   common.Address
	mainchainTFuelTokenBank       *scta.TFuelTokenBank // the TFuelTokenBank contract deployed on the mainchain
	mainchainTNT20TokenBankAddr   common.Address
	mainchainTNT20TokenBank       *scta.TNT20TokenBank // the TNT20TokenBank contract deployed on the mainchain
	mainchainTNT721TokenBankAddr  common.Address
	mainchainTNT721TokenBank      *scta.TNT721TokenBank // the TNT721TokenBank contract deployed on the mainchain
	mainchainTNT1155TokenBankAddr common.Address
	mainchainTNT1155TokenBank     *scta.TNT1155TokenBank // the TNT1155TokenBank contract deployed on the mainchain

	mainchainBlockHeight       *big.Int
	lastQueryedMainChainHeight *big.Int

	// The subchain
	subchainID                   *big.Int
	subchainEthRpcUrl            string
	subchainEthRpcClient         *ec.Client
	subchainBlockHeight          *big.Int
	subchainTFuelTokenBankAddr   common.Address
	subchainTFuelTokenBank       *scta.TFuelTokenBank // the TFuelTokenBank contract deployed on the subchain
	subchainTNT20TokenBankAddr   common.Address
	subchainTNT20TokenBank       *scta.TNT20TokenBank // the TNT20TokenBank contract deployed on the subchain
	subchainTNT721TokenBankAddr  common.Address
	subchainTNT721TokenBank      *scta.TNT721TokenBank
	subchainTNT1155TokenBankAddr common.Address
	subchainTNT1155TokenBank     *scta.TNT1155TokenBank
	// Validator set
	cacheMutex        *sync.Mutex // mutex to for validatorSetCache concurrent write protection
	validatorSetCache map[string]*score.ValidatorSet

	// Inter-chain messaging
	interChainEventCache *siu.InterChainEventCache

	// Life cycle
	wg     *sync.WaitGroup
	ctx    context.Context
	cancel context.CancelFunc
}

// NewMetachainWitness creates a new MetachainWitness
func NewMetachainWitness(db database.Database, updateInterval int, interChainEventCache *siu.InterChainEventCache) *MetachainWitness {
	mainchainEthRpcURL := viper.GetString(scom.CfgMainchainEthRpcURL)
	mainchainEthRpcClient, err := ec.Dial(mainchainEthRpcURL)
	if err != nil {
		logger.Fatalf("the ETH client failed to connect to the mainchain ETH RPC: %v\n", err)
	}
	mainchainID, err := mainchainEthRpcClient.ChainID(context.Background())
	if err != nil {
		logger.Fatalf("failed to get the chainID of the mainchain, is the mainchain RPC API service running? error: %v\n", err)
	}
	chainRegistrarOnMainchainAddr := common.HexToAddress(viper.GetString(scom.CfgChainRegistrarOnMainchainContractAddress))
	chainRegistrarOnMainchain, err := scta.NewChainRegistrarOnMainchain(chainRegistrarOnMainchainAddr, mainchainEthRpcClient)
	if err != nil {
		logger.Fatalf("failed to create ChainRegistrarOnMainchain contract: %v\n", err)
	}
	mainchainTFuelTokenBankAddr := common.HexToAddress(viper.GetString(scom.CfgMainchainTFuelTokenBankContractAddress))
	mainchainTFuelTokenBank, err := scta.NewTFuelTokenBank(mainchainTFuelTokenBankAddr, mainchainEthRpcClient)
	if err != nil {
		logger.Fatalf("failed to create MainchainTFuelTokenBank contract: %v\n", err)
	}
	mainchainTNT20TokenBankAddr := common.HexToAddress(viper.GetString(scom.CfgMainchainTNT20TokenBankContractAddress))
	mainchainTNT20TokenBank, err := scta.NewTNT20TokenBank(mainchainTNT20TokenBankAddr, mainchainEthRpcClient)
	if err != nil {
		logger.Fatalf("failed to create MainchainTNT20TokenBank contract: %v\n", err)
	}
	mainchainTNT721TokenBankAddr := common.HexToAddress(viper.GetString(scom.CfgMainchainTNT721TokenBankContractAddress))
	mainchainTNT721TokenBank, err := scta.NewTNT721TokenBank(mainchainTNT721TokenBankAddr, mainchainEthRpcClient)
	if err != nil {
		logger.Fatalf("failed to create MainchainTNT20TokenBank contract: %v\n", err)
	}
	mainchainTNT1155TokenBankAddr := common.HexToAddress(viper.GetString(scom.CfgMainchainTNT1155TokenBankContractAddress))
	mainchainTNT1155TokenBank, err := scta.NewTNT1155TokenBank(mainchainTNT1155TokenBankAddr, mainchainEthRpcClient)
	if err != nil {
		logger.Fatalf("failed to create MainchainTNT20TokenBank contract %v\n", err)
	}

	subchainID := big.NewInt(viper.GetInt64(scom.CfgSubchainID))
	subchainEthRpcURL := viper.GetString(scom.CfgSubchainEthRpcURL)
	subchainEthRpcClient, err := ec.Dial(subchainEthRpcURL)
	if err != nil {
		logger.Fatalf("the ETH client failed to connect to the subchain ETH RPC: %v\n", err)
	}

	witnessState := newMetachainWitnessState(db)
	validatorSet := make(map[string]*score.ValidatorSet)

	var queryTopics string
	for _, eventTopicString := range siu.EventSelectors {
		queryTopics = queryTopics + ",\"" + eventTopicString + "\""
	}

	mw := &MetachainWitness{
		updateInterval: updateInterval,
		witnessState:   witnessState,
		queryTopics:    queryTopics[1:],

		mainchainID:                   mainchainID,
		mainchainEthRpcUrl:            mainchainEthRpcURL,
		mainchainEthRpcClient:         mainchainEthRpcClient,
		witnessedDynasty:              big.NewInt(0),
		chainRegistrarOnMainchain:     chainRegistrarOnMainchain,
		mainchainTFuelTokenBankAddr:   mainchainTFuelTokenBankAddr,
		mainchainTFuelTokenBank:       mainchainTFuelTokenBank,
		mainchainTNT20TokenBankAddr:   mainchainTNT20TokenBankAddr,
		mainchainTNT20TokenBank:       mainchainTNT20TokenBank,
		mainchainTNT721TokenBankAddr:  mainchainTNT721TokenBankAddr,
		mainchainTNT721TokenBank:      mainchainTNT721TokenBank,
		mainchainTNT1155TokenBankAddr: mainchainTNT1155TokenBankAddr,
		mainchainTNT1155TokenBank:     mainchainTNT1155TokenBank,
		mainchainBlockHeight:          nil,
		lastQueryedMainChainHeight:    big.NewInt(0),

		subchainID:           subchainID,
		subchainEthRpcUrl:    subchainEthRpcURL,
		subchainEthRpcClient: subchainEthRpcClient,
		subchainBlockHeight:  nil,
		cacheMutex:           &sync.Mutex{},
		validatorSetCache:    validatorSet,
		interChainEventCache: interChainEventCache,

		wg: &sync.WaitGroup{},
	}
	return mw
}

func (mw *MetachainWitness) Start(ctx context.Context) {
	c, cancel := context.WithCancel(ctx)
	mw.ctx = c
	mw.cancel = cancel

	mw.wg.Add(1)
	go mw.mainloop(ctx)
}

func (mw *MetachainWitness) Stop() {
	if mw.updateTicker != nil {
		mw.updateTicker.Stop()
	}
	mw.cancel()
}

func (mw *MetachainWitness) Wait() {
	mw.wg.Wait()
}

func (mw *MetachainWitness) SetSubchainTokenBanks(ledger score.Ledger) {
	var err error
	subchainTFuelTokenBankAddr := ledger.GetTokenBankContractAddress(score.CrossChainTokenTypeTFuel)
	if subchainTFuelTokenBankAddr == nil {
		logger.Fatalf("failed to obtain SubchainTFuelTokenBank contract address\n")
	}
	mw.subchainTFuelTokenBankAddr = *subchainTFuelTokenBankAddr
	mw.subchainTFuelTokenBank, err = scta.NewTFuelTokenBank(*subchainTFuelTokenBankAddr, mw.subchainEthRpcClient)
	if err != nil {
		logger.Fatalf("failed to set the SubchainTFuelTokenBank contract: %v\n", err)
	}

	subchainTNT20TokenBankAddr := ledger.GetTokenBankContractAddress(score.CrossChainTokenTypeTNT20)
	if subchainTNT20TokenBankAddr == nil {
		logger.Fatalf("failed to obtain SubchainTNT20TokenBank contract address\n")
	}
	mw.subchainTNT20TokenBankAddr = *subchainTNT20TokenBankAddr
	mw.subchainTNT20TokenBank, err = scta.NewTNT20TokenBank(*subchainTNT20TokenBankAddr, mw.subchainEthRpcClient)
	if err != nil {
		logger.Fatalf("failed to set the SubchainTNT20TokenBankAddr contract: %v\n", err)
	}

	subchainTNT721TokenBankAddr := ledger.GetTokenBankContractAddress(score.CrossChainTokenTypeTNT721)
	if subchainTNT721TokenBankAddr == nil {
		logger.Fatalf("failed to obtain SubchainTNT721TokenBank contract address\n")
	}
	mw.subchainTNT721TokenBankAddr = *subchainTNT721TokenBankAddr
	mw.subchainTNT721TokenBank, err = scta.NewTNT721TokenBank(*subchainTNT721TokenBankAddr, mw.subchainEthRpcClient)
	if err != nil {
		logger.Fatalf("failed to set the SubchainTNT20TokenBankAddr contract: %v\n", err)
	}
	subchainTNT1155TokenBankAddr := ledger.GetTokenBankContractAddress(score.CrossChainTokenTypeTNT1155)
	if subchainTNT721TokenBankAddr == nil || err != nil {
		logger.Fatalf("failed to obtain SubchainTNT721TokenBank contract address: %v\n", err)
	}
	mw.subchainTNT1155TokenBankAddr = *subchainTNT1155TokenBankAddr
	mw.subchainTNT1155TokenBank, err = scta.NewTNT1155TokenBank(*subchainTNT1155TokenBankAddr, mw.subchainEthRpcClient)
	if err != nil {
		logger.Fatalf("failed to set the SubchainTNT20TokenBankAddr contract: %v\n", err)
	}
}

// TODO: make sure the block number returned by the client.BlockNumber() call is the lastest *finalized* block number
func (mw *MetachainWitness) GetMainchainBlockHeight() (*big.Int, error) {
	if mw.mainchainBlockHeight == nil {
		return nil, fmt.Errorf("mainchain block height not been updated yet")
	}
	return mw.mainchainBlockHeight, nil
}

func (mw *MetachainWitness) GetSubchainBlockHeight() (*big.Int, error) {
	if mw.subchainBlockHeight == nil {
		return nil, fmt.Errorf("mainchain block height not been updated yet")
	}
	return mw.subchainBlockHeight, nil
}

func (mw *MetachainWitness) GetValidatorSetByDynasty(dynasty *big.Int) (*score.ValidatorSet, error) {
	validatorSet, ok := mw.validatorSetCache[dynasty.String()]
	if ok && validatorSet != nil && validatorSet.Dynasty() == dynasty {
		return validatorSet, nil
	}

	var err error
	validatorSet, err = mw.updateValidatorSetCache(dynasty) // cache lazy update
	if err != nil {
		return nil, err
	}

	return validatorSet, nil
}

func (mw *MetachainWitness) mainloop(ctx context.Context) {
	mw.updateTicker = time.NewTicker(time.Duration(mw.updateInterval) * time.Millisecond)
	for {
		select {
		case <-ctx.Done():
			return
		case <-mw.updateTicker.C:
			mw.update()
		}
	}
}

func (mw *MetachainWitness) update() {
	// Mainchain
	mw.updateMainchainBlockHeight()
	dynasty := scom.CalculateDynasty(mw.mainchainBlockHeight)
	if mw.witnessedDynasty == nil || dynasty.Cmp(mw.witnessedDynasty) > 0 { // needs to update the cache
		mw.updateValidatorSetCache(dynasty)
		mw.witnessedDynasty = dynasty
	}
	mw.collectInterChainMessageEventsOnMainchain()

	// Subchain
	mw.collectInterChainMessageEventsOnSubchain()
	mw.updateSubchainBlockHeight()
}

func (mw *MetachainWitness) updateMainchainBlockHeight() {
	mbh, err := mw.mainchainEthRpcClient.BlockNumber(context.Background())
	if err != nil {
		logger.Warnf("failed to get the mainchain block height %v\n", err)
		return
	}
	mw.mainchainBlockHeight = big.NewInt(int64(mbh))
}

func (mw *MetachainWitness) updateSubchainBlockHeight() {
	sbh, err := mw.subchainEthRpcClient.BlockNumber(context.Background())
	if err != nil {
		return
	}
	mw.subchainBlockHeight = big.NewInt(int64(sbh))
}

func (mw *MetachainWitness) collectInterChainMessageEventsOnMainchain() {
	mw.collectInterChainMessageEventsOnChain(mw.mainchainID, mw.mainchainEthRpcUrl, mw.mainchainTFuelTokenBankAddr, mw.mainchainTNT20TokenBankAddr, mw.mainchainTNT721TokenBankAddr, mw.mainchainTNT1155TokenBankAddr)
}

func (mw *MetachainWitness) collectInterChainMessageEventsOnSubchain() {
	mw.collectInterChainMessageEventsOnChain(mw.subchainID, mw.subchainEthRpcUrl, mw.subchainTFuelTokenBankAddr, mw.subchainTNT20TokenBankAddr, mw.subchainTNT721TokenBankAddr, mw.subchainTNT1155TokenBankAddr)
}

func (mw *MetachainWitness) collectInterChainMessageEventsOnChain(queriedChainID *big.Int, ethRpcUrl string,
	tfuelTokenBankAddr common.Address, tnt20TokenBankAddr common.Address, tnt721TokenBankAddr common.Address, tnt1155TokenBankAddr common.Address) {
	// mw.getBlockScanStartingHeight(queriedChainID) // testing code

	fromBlock, err := mw.witnessState.getLastQueryedHeightForType(queriedChainID)
	if err == store.ErrKeyNotFound {
		logger.Infof("Seting the initial block scanning height for chain %v", queriedChainID.String())
		if queriedChainID.Cmp(mw.mainchainID) == 0 && viper.GetInt(scom.CfgSubchainMainchainWitenessStartScanHeight) >= 0 {
			// if the start scanning height is specified in the config, use the config value
			fromBlock = big.NewInt(int64(viper.GetInt(scom.CfgSubchainMainchainWitenessStartScanHeight)))
			logger.Infof("Reading the initial scanning height for chain %v from configs: %v", queriedChainID.String(), fromBlock.String())
		} else if queriedChainID.Cmp(mw.subchainID) == 0 && viper.GetInt(scom.CfgSubchainSubchainWitenessStartScanHeight) >= 0 {
			// if the start scanning height is specified in the config, use the config value
			fromBlock = big.NewInt(int64(viper.GetInt(scom.CfgSubchainSubchainWitenessStartScanHeight)))
			logger.Infof("Reading the initial scanning height for chain %v from configs: %v", queriedChainID.String(), fromBlock.String())
		} else {
			// Otherwise, figure out the starting height from the TokenBanks
			fromBlock = mw.getBlockScanStartingHeight(queriedChainID) // set the proper fromBlock for the code-start scenario, i.e, bootstrapping a new validator
		}

		mw.witnessState.setLastQueryedHeightForType(queriedChainID, fromBlock)
		logger.Infof("Initial block scanning height for chain %v set to: %v", queriedChainID.String(), fromBlock.String())
	} else if err != nil {
		logger.Warnf("failed to get the last queryed height %v\n", err)
	}
	toBlock := mw.calculateToBlock(fromBlock, queriedChainID)
	logger.Infof("Query inter-chain message events from block height %v to %v on chain %v", fromBlock.String(), toBlock.String(), queriedChainID.String())
	events := siu.QueryInterChainEventLog(queriedChainID, fromBlock, toBlock, tfuelTokenBankAddr, tnt20TokenBankAddr, tnt721TokenBankAddr, tnt1155TokenBankAddr, mw.queryTopics, ethRpcUrl)
	err = mw.interChainEventCache.InsertList(events)
	if err != nil { // should not happen
		logger.Panicf("failed to insert events into cache")
	}
	mw.witnessState.setLastQueryedHeightForType(queriedChainID, toBlock)
}

func (mw *MetachainWitness) getBlockScanStartingHeight(queriedChainID *big.Int) *big.Int {
	updateHeight := big.NewInt(0).Set(common.BigMaxUint64)

	eventTypes := []score.InterChainMessageEventType{
		score.IMCEventTypeCrossChainTokenLockTFuel,
		score.IMCEventTypeCrossChainTokenLockTNT20,
		score.IMCEventTypeCrossChainTokenLockTNT721,
		score.IMCEventTypeCrossChainVoucherBurnTFuel,
		score.IMCEventTypeCrossChainVoucherBurnTNT20,
		score.IMCEventTypeCrossChainVoucherBurnTNT721,
		score.IMCEventTypeCrossChainVoucherBurnTNT1155,
	}

	for _, eventType := range eventTypes {
		var height *big.Int

		if queriedChainID.Cmp(mw.mainchainID) == 0 {
			height = mw.getMainchainMaxProcessedNonceEventHeight(eventType)
		} else if queriedChainID.Cmp(mw.subchainID) == 0 {
			height = mw.getSubchainMaxProcessedNonceEventHeight(eventType)
		} else {
			logger.Panicf("invalid queriedChainID") // should not happen
		}

		if height.Cmp(updateHeight) < 0 {
			updateHeight.Set(height)
		}
	}

	startHeight := big.NewInt(0).Set(updateHeight)
	margin := big.NewInt(5)
	if startHeight.Cmp(margin) > 0 {
		startHeight = big.NewInt(0).Sub(startHeight, margin)
	} else {
		startHeight = common.Big0
	}

	logger.Infof("Block scanning starting height for chain %v: %v (nonce updateHeight: %v)", queriedChainID, startHeight, updateHeight)

	return startHeight
}

func (mw *MetachainWitness) getMainchainMaxProcessedNonceEventHeight(icmeType score.InterChainMessageEventType) *big.Int {
	var maxProcessedNonce *big.Int
	var eventHeight *big.Int
	var err error

	// For mainchain -> subchain asset transfers, the "max processed nonce" (for each event type) is recorded on the subchain side.
	// Yet the height for the corresponding event is recorded on the mainchain. Hence we get the "max processed nonce" from the subchain
	// and use it to lookup the event height on the mainchain.
	switch icmeType {
	case score.IMCEventTypeCrossChainTokenLockTFuel:
		maxProcessedNonce, err = mw.subchainTFuelTokenBank.GetMaxProcessedTokenLockNonce(nil, mw.mainchainID)
		if err != nil {
			break
		}
		eventHeight, err = mw.mainchainTFuelTokenBank.GetTokenLockEventHeight(nil, mw.subchainID, maxProcessedNonce)
	case score.IMCEventTypeCrossChainTokenLockTNT20:
		maxProcessedNonce, err = mw.subchainTNT20TokenBank.GetMaxProcessedTokenLockNonce(nil, mw.mainchainID)
		if err != nil {
			break
		}
		eventHeight, err = mw.mainchainTNT20TokenBank.GetTokenLockEventHeight(nil, mw.subchainID, maxProcessedNonce)
	case score.IMCEventTypeCrossChainTokenLockTNT721:
		maxProcessedNonce, err = mw.subchainTNT721TokenBank.GetMaxProcessedTokenLockNonce(nil, mw.mainchainID)
		if err != nil {
			break
		}
		eventHeight, err = mw.mainchainTNT721TokenBank.GetTokenLockEventHeight(nil, mw.subchainID, maxProcessedNonce)
	case score.IMCEventTypeCrossChainTokenLockTNT1155:
		maxProcessedNonce, err = mw.subchainTNT1155TokenBank.GetMaxProcessedTokenLockNonce(nil, mw.mainchainID)
		if err != nil {
			break
		}
		eventHeight, err = mw.mainchainTNT721TokenBank.GetTokenLockEventHeight(nil, mw.subchainID, maxProcessedNonce)

	case score.IMCEventTypeCrossChainVoucherBurnTFuel:
		// Note: TFuelVoucherBurn is not allowed on the Mainchain so it is safe to return the latest block height
		maxProcessedNonce = common.Big0
		h, _ := mw.mainchainEthRpcClient.BlockNumber(context.Background())
		eventHeight = big.NewInt(int64(h))
	case score.IMCEventTypeCrossChainVoucherBurnTNT20:
		maxProcessedNonce, err = mw.subchainTNT20TokenBank.GetMaxProcessedVoucherBurnNonce(nil, mw.mainchainID)
		if err != nil {
			break
		}
		eventHeight, err = mw.mainchainTNT20TokenBank.GetVoucherBurnEventHeight(nil, mw.subchainID, maxProcessedNonce)
	case score.IMCEventTypeCrossChainVoucherBurnTNT721:
		maxProcessedNonce, err = mw.subchainTNT721TokenBank.GetMaxProcessedVoucherBurnNonce(nil, mw.mainchainID)
		if err != nil {
			break
		}
		eventHeight, err = mw.mainchainTNT721TokenBank.GetVoucherBurnEventHeight(nil, mw.subchainID, maxProcessedNonce)
	case score.IMCEventTypeCrossChainVoucherBurnTNT1155:
		maxProcessedNonce, err = mw.subchainTNT1155TokenBank.GetMaxProcessedVoucherBurnNonce(nil, mw.mainchainID)
		if err != nil {
			break
		}
		eventHeight, err = mw.mainchainTNT1155TokenBank.GetVoucherBurnEventHeight(nil, mw.subchainID, maxProcessedNonce)

	default:
		logger.Panicf("invalid event type: %v", icmeType) // should not happen
	}
	if err != nil {
		logger.Warnf("Failed to get the update height for max processed nonce on the main chain for event type %v: %v", icmeType, err)
	}

	if maxProcessedNonce == nil || eventHeight == nil {
		eventHeight = big.NewInt(0)
	} else if maxProcessedNonce.Cmp(common.Big0) == 0 {
		// no event of the current icmeType has ever been processed, hence it is safe to scan from the current block height
		h, _ := mw.mainchainEthRpcClient.BlockNumber(context.Background())
		eventHeight = big.NewInt(int64(h))
	}

	logger.Infof("On the Mainchain, event type: %v, max processed nonce: %v, update height: %v", icmeType, maxProcessedNonce, eventHeight)

	return eventHeight
}

func (mw *MetachainWitness) getSubchainMaxProcessedNonceEventHeight(icmeType score.InterChainMessageEventType) *big.Int {
	var maxProcessedNonce *big.Int
	var eventHeight *big.Int
	var err error

	// For subchain -> mainchain asset transfers, the "max processed nonce" (for each event type) is recorded on the mainchain side.
	// Yet the height for the corresponding event is recorded on the subchain. Hence we get the "max processed nonce" from the mainchain
	// and use it to lookup the event height on the subchain.
	switch icmeType { // Note: TFuelTokenLock is not allowed on a Subchain so this event is not processed below

	case score.IMCEventTypeCrossChainTokenLockTFuel:
		// Note: TFuelTokenLock is not allowed on a Subchain so it is safe to return the latest block height
		maxProcessedNonce = common.Big0
		h, _ := mw.subchainEthRpcClient.BlockNumber(context.Background())
		eventHeight = big.NewInt(int64(h))
	case score.IMCEventTypeCrossChainTokenLockTNT20:
		maxProcessedNonce, err = mw.mainchainTNT20TokenBank.GetMaxProcessedTokenLockNonce(nil, mw.subchainID)
		if err != nil {
			break
		}
		eventHeight, err = mw.subchainTNT20TokenBank.GetTokenLockEventHeight(nil, mw.mainchainID, maxProcessedNonce)
	case score.IMCEventTypeCrossChainTokenLockTNT721:
		maxProcessedNonce, err = mw.mainchainTNT721TokenBank.GetMaxProcessedTokenLockNonce(nil, mw.subchainID)
		if err != nil {
			break
		}
		eventHeight, err = mw.subchainTNT721TokenBank.GetTokenLockEventHeight(nil, mw.mainchainID, maxProcessedNonce)
	case score.IMCEventTypeCrossChainTokenLockTNT1155:
		maxProcessedNonce, err = mw.mainchainTNT1155TokenBank.GetMaxProcessedTokenLockNonce(nil, mw.subchainID)
		if err != nil {
			break
		}
		eventHeight, err = mw.subchainTNT1155TokenBank.GetTokenLockEventHeight(nil, mw.mainchainID, maxProcessedNonce)

	case score.IMCEventTypeCrossChainVoucherBurnTFuel:
		maxProcessedNonce, err = mw.mainchainTFuelTokenBank.GetMaxProcessedVoucherBurnNonce(nil, mw.subchainID)
		if err != nil {
			break
		}
		eventHeight, err = mw.subchainTFuelTokenBank.GetVoucherBurnEventHeight(nil, mw.mainchainID, maxProcessedNonce)
	case score.IMCEventTypeCrossChainVoucherBurnTNT20:
		maxProcessedNonce, err = mw.mainchainTNT20TokenBank.GetMaxProcessedVoucherBurnNonce(nil, mw.subchainID)
		if err != nil {
			break
		}
		eventHeight, err = mw.subchainTNT20TokenBank.GetVoucherBurnEventHeight(nil, mw.mainchainID, maxProcessedNonce)
	case score.IMCEventTypeCrossChainVoucherBurnTNT721:
		maxProcessedNonce, err = mw.mainchainTNT721TokenBank.GetMaxProcessedVoucherBurnNonce(nil, mw.subchainID)
		if err != nil {
			break
		}
		eventHeight, err = mw.subchainTNT721TokenBank.GetVoucherBurnEventHeight(nil, mw.mainchainID, maxProcessedNonce)
	case score.IMCEventTypeCrossChainVoucherBurnTNT1155:
		maxProcessedNonce, err = mw.mainchainTNT1155TokenBank.GetMaxProcessedVoucherBurnNonce(nil, mw.subchainID)
		if err != nil {
			break
		}
		eventHeight, err = mw.subchainTNT1155TokenBank.GetVoucherBurnEventHeight(nil, mw.mainchainID, maxProcessedNonce)

	default:
		logger.Panicf("invalid event type: %v", icmeType) // should not happen
	}
	if err != nil {
		logger.Warnf("Failed to get the update height for max processed nonce on subchain %v for event type %v: %v", mw.subchainID, icmeType, err)
	}

	if maxProcessedNonce == nil || eventHeight == nil {
		eventHeight = big.NewInt(0)
	} else if maxProcessedNonce.Cmp(common.Big0) == 0 {
		// no event of the current icmeType has ever been processed, hence it is safe to scan from the current block height
		h, _ := mw.subchainEthRpcClient.BlockNumber(context.Background())
		eventHeight = big.NewInt(int64(h))
	}

	logger.Infof("On the Subchain, event type: %v, max processed nonce: %v, update height: %v", icmeType, maxProcessedNonce, eventHeight)

	return eventHeight
}

func (mw *MetachainWitness) calculateToBlock(fromBlock *big.Int, queriedChainID *big.Int) *big.Int {
	var toBlock *big.Int
	var err error
	if queriedChainID == mw.mainchainID {
		toBlock, err = mw.GetMainchainBlockHeight()
	} else {
		toBlock, err = mw.GetSubchainBlockHeight()
	}
	if err != nil {
		return fromBlock
	}
	maxBlockRange := int64(300) // block range query allows at most 5000 blocks, here we intentionally use a much smaller range to limit cpu/mem resource usage
	minBlockGap := int64(2)     // tentative, to ensure the chain has enough time to finalize the event
	if new(big.Int).Sub(toBlock, fromBlock).Cmp(big.NewInt(maxBlockRange)) > 0 {
		// catch-up phase, gap is over maxBlockRange，catch-up at full speed
		toBlock = new(big.Int).Add(fromBlock, big.NewInt(maxBlockRange))
	} else {
		// steady phase, gap is between minBlockGap and maxBlockRange
		toBlock = new(big.Int).Sub(toBlock, big.NewInt(minBlockGap))
	}
	return toBlock
}

func (mw *MetachainWitness) updateValidatorSetCache(dynasty *big.Int) (*score.ValidatorSet, error) {
	mw.cacheMutex.Lock()
	defer mw.cacheMutex.Unlock()

	// queryBlockHeight := big.NewInt(1).Mul(dynasty, big.NewInt(1).SetInt64(scom.NumMainchainBlocksPerDynasty))
	// queryBlockHeight = big.NewInt(0).Add(queryBlockHeight, big.NewInt(1)) // increment by one to make sure the query block height falls into the dynasty
	// vs, err := mw.chainRegistrarOnMainchain.GetValidatorSet(nil, mw.subchainID, queryBlockHeight)

	// adjustedDynasty := dynasty
	// registrationMainchainHeight, ok, err := mw.chainRegistrarOnMainchain.GetSubchainRegistrationHeight(nil, mw.subchainID)
	// if !ok {
	// 	errMsg := fmt.Sprintf("subchain with ID %v does not exist", mw.subchainID)
	// 	logger.Warnf("updateValidatorSetCache: %v", errMsg)
	// 	return nil, fmt.Errorf(errMsg)
	// }

	// if err != nil {
	// 	logger.Warnf("updateValidatorSetCache: %v", err)
	// 	return nil, err
	// }

	// registrationDynasty := scom.CalculateDynasty(registrationMainchainHeight)
	// if adjustedDynasty.Cmp(registrationDynasty) == 0 {
	// 	// Special handling for the initial dynasty query on the **Mainchain**:
	// 	// The initial set of validators hardcoded in the genesis snapshot should be
	// 	// in charge during the initial dyansty. However, if we directly query the
	// 	// initial dynasty, we would obtain an empty set since the validators registered
	// 	// on the mainchain takes charge only when the next dynasty begins. Hence, for
	// 	// the initial dyansty, we query the next dynasty instead, which should return
	// 	// a validator set that matches with the validator set hardcoded in the
	// 	// genesis snapshot.
	// 	adjustedDynasty = big.NewInt(0).Add(adjustedDynasty, big.NewInt(1))

	// 	logger.Debugf("Subchain %v registration main chain height: %v, registration dynasty: %v, dynasty: %v, adjustedDynasty: %v",
	// 		mw.subchainID, registrationMainchainHeight, registrationDynasty, dynasty, adjustedDynasty)
	// }

	// vs, err := mw.chainRegistrarOnMainchain.GetValidatorSet(nil, mw.subchainID, adjustedDynasty)

	vs, err := mw.chainRegistrarOnMainchain.GetValidatorSet(nil, mw.subchainID, dynasty)
	validatorAddrs := vs.Validators
	validatorStakes := vs.ShareAmounts

	if err != nil {
		return nil, err
	}

	if len(validatorAddrs) != len(validatorStakes) {
		return nil, fmt.Errorf("the length of validatorAddrs and validatorStakes are not equal")
	}

	validatorSet := score.NewValidatorSet(dynasty)
	for i := 0; i < len(validatorAddrs); i++ {
		validator := score.NewValidator(validatorAddrs[i].Hex(), validatorStakes[i])
		validatorSet.AddValidator(validator)
	}

	mw.validatorSetCache[dynasty.String()] = validatorSet

	logger.Debugf("updateValidatorSetCache, dynasty = %v, validatorSet = %v", dynasty.String(), validatorSet.String())

	return validatorSet, nil
}

func (mw *MetachainWitness) GetInterChainEventCache() *siu.InterChainEventCache {
	return mw.interChainEventCache
}

func (mw *MetachainWitness) GetSubchainRegistrationHeight() (*big.Int, error) {
	registrationMainchainHeight, ok, err := mw.chainRegistrarOnMainchain.GetSubchainRegistrationHeight(nil, mw.subchainID)
	if !ok {
		errMsg := fmt.Sprintf("subchain with ID %v does not exist", mw.subchainID)
		logger.Warnf("GetSubchainRegistrationHeight: %v", errMsg)
		return nil, fmt.Errorf(errMsg)
	}

	if err != nil {
		logger.Warnf("GetSubchainRegistrationHeight: %v", err)
		return nil, err
	}

	return registrationMainchainHeight, nil
}
