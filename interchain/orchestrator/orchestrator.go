package orchestrator

import (
	"context"
	"math/big"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/thetatoken/theta/crypto"
	ts "github.com/thetatoken/theta/store"
	"github.com/thetatoken/theta/store/database"
	"github.com/thetatoken/thetasubchain/eth/abi/bind"
	siu "github.com/thetatoken/thetasubchain/interchain/utils"
	"github.com/thetatoken/thetasubchain/interchain/witness"

	scom "github.com/thetatoken/thetasubchain/common"
	score "github.com/thetatoken/thetasubchain/core"
	scta "github.com/thetatoken/thetasubchain/interchain/contracts/accessors"

	"github.com/thetatoken/theta/common"
	ec "github.com/thetatoken/thetasubchain/eth/ethclient"
)

var logger *log.Entry = log.WithFields(log.Fields{"prefix": "orchestrator"})

type Orchestrator struct {
	updateInterval          int
	privateKey              *crypto.PrivateKey
	ledger                  score.Ledger
	eventProcessingTicker   *time.Ticker
	metachainWitness        witness.ChainWitness
	eventProcessedTimestamp map[string]time.Time

	// The mainchain
	mainchainID                 *big.Int
	mainchainEthRpcURL          string
	mainchainEthRpcClient       *ec.Client
	mainchainTFuelTokenBankAddr common.Address
	mainchainTFuelTokenBank     *scta.TFuelTokenBank
	mainchainTNT20TokenBankAddr common.Address
	mainchainTNT20TokenBank     *scta.TNT20TokenBank

	// The subchain
	subchainID                 *big.Int
	subchainEthRpcURL          string
	subchainEthRpcClient       *ec.Client
	subchainTFuelTokenBankAddr common.Address
	subchainTFuelTokenBank     *scta.TFuelTokenBank
	subchainTNT20TokenBankAddr common.Address
	subchainTNT20TokenBank     *scta.TNT20TokenBank

	// Inter-chain messaging
	interChainEventCache *siu.InterChainEventCache

	// Life cycle
	wg     *sync.WaitGroup
	ctx    context.Context
	cancel context.CancelFunc
}

// NewOrchestrator creates a new Orchestrator
func NewOrchestrator(db database.Database, updateInterval int, interChainEventCache *siu.InterChainEventCache,
	metachainWitness witness.ChainWitness, privateKey *crypto.PrivateKey) *Orchestrator {

	mainchainEthRpcURL := viper.GetString(scom.CfgMainchainEthRpcURL)
	mainchainEthRpcClient, err := ec.Dial(mainchainEthRpcURL)
	if err != nil {
		logger.Fatalf("the ETH client failed to connect to the mainchain ETH RPC %v\n", err)
	}
	mainchainID, err := mainchainEthRpcClient.ChainID(context.Background())
	if err != nil {
		logger.Fatalf("failed to get the chainID of the mainchain %v\n", err)
	}
	mainchainTFuelTokenBankAddr := common.HexToAddress(viper.GetString(scom.CfgMainchainTFuelTokenBankContractAddress))
	mainchainTFuelTokenBank, err := scta.NewTFuelTokenBank(mainchainTFuelTokenBankAddr, mainchainEthRpcClient)
	if err != nil {
		logger.Fatalf("failed to create MainchainTFuelTokenBank contract %v\n", err)
	}
	mainchainTNT20TokenBankAddr := common.HexToAddress(viper.GetString(scom.CfgMainchainTNT20TokenBankContractAddress))
	mainchainTNT20TokenBank, err := scta.NewTNT20TokenBank(mainchainTNT20TokenBankAddr, mainchainEthRpcClient)
	if err != nil {
		logger.Fatalf("failed to create MainchainTNT20TokenBank contract %v\n", err)
	}

	subchainID := big.NewInt(viper.GetInt64(scom.CfgSubchainID))
	subchainEthRpcURL := viper.GetString(scom.CfgSubchainEthRpcURL)
	subchainEthRpcClient, err := ec.Dial(subchainEthRpcURL)
	if err != nil {
		logger.Fatalf("the ETH client failed to connect to the subchain ETH RPC %v\n", err)
	}

	oc := &Orchestrator{
		updateInterval:   updateInterval,
		privateKey:       privateKey,
		metachainWitness: metachainWitness,

		mainchainID:                 mainchainID,
		mainchainEthRpcURL:          mainchainEthRpcURL,
		mainchainEthRpcClient:       mainchainEthRpcClient,
		mainchainTFuelTokenBankAddr: mainchainTFuelTokenBankAddr,
		mainchainTFuelTokenBank:     mainchainTFuelTokenBank,
		mainchainTNT20TokenBankAddr: mainchainTNT20TokenBankAddr,
		mainchainTNT20TokenBank:     mainchainTNT20TokenBank,

		subchainID:           subchainID,
		subchainEthRpcURL:    subchainEthRpcURL,
		subchainEthRpcClient: subchainEthRpcClient,

		interChainEventCache: interChainEventCache,

		wg: &sync.WaitGroup{},
	}
	return oc
}

func (oc *Orchestrator) Start(ctx context.Context) {
	c, cancel := context.WithCancel(ctx)
	oc.ctx = c
	oc.cancel = cancel

	oc.wg.Add(1)
	go oc.mainloop(ctx)
	logger.Info("Metachain orchestrator started")
}

func (oc *Orchestrator) Stop() {
	if oc.eventProcessingTicker != nil {
		oc.eventProcessingTicker.Stop()
	}
	oc.cancel()
	logger.Info("Metachain orchestrator stopped")
}

func (oc *Orchestrator) Wait() {
	oc.wg.Wait()
}

func (oc *Orchestrator) SetLedgerAndSubchainTokenBanks(ledger score.Ledger) {
	oc.ledger = ledger

	subchainTFuelTokenBankAddr, err := ledger.GetTokenBankContractAddress(score.CrossChainTokenTypeTFuel)
	if subchainTFuelTokenBankAddr == nil || err != nil {
		logger.Fatalf("failed to obtain SubchainTFuelTokenBank contract address: %v\n", err)
	}
	oc.subchainTFuelTokenBankAddr = *subchainTFuelTokenBankAddr
	oc.subchainTFuelTokenBank, err = scta.NewTFuelTokenBank(*subchainTFuelTokenBankAddr, oc.subchainEthRpcClient)
	if err != nil {
		logger.Fatalf("failed to set the SubchainTFuelTokenBank contract: %v\n", err)
	}

	subchainTNT20TokenBankAddr, err := ledger.GetTokenBankContractAddress(score.CrossChainTokenTypeTNT20)
	if subchainTNT20TokenBankAddr == nil || err != nil {
		logger.Fatalf("failed to obtain SubchainTNT20TokenBank contract address: %v\n", err)
	}
	oc.subchainTNT20TokenBankAddr = *subchainTNT20TokenBankAddr
	oc.subchainTNT20TokenBank, err = scta.NewTNT20TokenBank(*subchainTNT20TokenBankAddr, oc.subchainEthRpcClient)
	if err != nil {
		logger.Fatalf("failed to set the SubchainTNT20TokenBankAddr contract: %v\n", err)
	}
}

func (oc *Orchestrator) mainloop(ctx context.Context) {
	oc.eventProcessingTicker = time.NewTicker(time.Duration(oc.updateInterval) * time.Millisecond)
	for {
		select {
		case <-ctx.Done():
			return
		case <-oc.eventProcessingTicker.C:
			// Handle token lock events
			oc.processNextTokenLockEvent(oc.mainchainID, oc.subchainID) // send token from the mainchain to the subchain
			oc.processNextTokenLockEvent(oc.subchainID, oc.mainchainID) // send token from the subchain to the mainchain

			// Handle voucher burn events
			oc.processNextVoucherBurnEvent(oc.mainchainID, oc.subchainID) // burn voucher to send token from the mainchain back to the subchain
			oc.processNextVoucherBurnEvent(oc.subchainID, oc.mainchainID) // burn voucher to send token from the subchain back to the mainchain
		}
	}
}

func (oc *Orchestrator) processNextTokenLockEvent(sourceChainID *big.Int, targetChainID *big.Int) {
	oc.processNextTFuelTokenLockEvent(sourceChainID, targetChainID)
	oc.processNextTNT20TokenLockEvent(sourceChainID, targetChainID)
}

func (oc *Orchestrator) processNextTFuelTokenLockEvent(sourceChainID *big.Int, targetChainID *big.Int) {
	targetChainTokenBank := oc.getTFuelTokenBank(targetChainID)
	maxProcessedTokenLockNonce, err := targetChainTokenBank.Getmaxprocessedtokenlocknonce(nil, sourceChainID)
	if err != nil {
		logger.Warnf("Failed to query the max processed TFuel token lock nonce for chain: %v", targetChainID.String())
		return // ignore, will retry shortly
	}

	oc.processNextEvent(sourceChainID, targetChainID, score.IMCEventTypeCrossChainTokenLockTFuel, maxProcessedTokenLockNonce)
}

func (oc *Orchestrator) processNextTNT20TokenLockEvent(sourceChainID *big.Int, targetChainID *big.Int) {
	targetChainTokenBank := oc.getTNT20TokenBank(targetChainID)
	maxProcessedTokenLockNonce, err := targetChainTokenBank.Getmaxprocessedtokenlocknonce(nil, sourceChainID)
	if err != nil {
		logger.Warnf("Failed to query the max processed TNT20 token lock nonce for chain: %v", targetChainID.String())
		return // ignore, will retry shortly
	}

	oc.processNextEvent(sourceChainID, targetChainID, score.IMCEventTypeCrossChainTokenLockTNT20, maxProcessedTokenLockNonce)
}

func (oc *Orchestrator) processNextVoucherBurnEvent(sourceChainID *big.Int, targetChainID *big.Int) {
	oc.processNextTFuelVoucherBurnEvent(sourceChainID, targetChainID)
	oc.processNextTNT20VoucherBurnEvent(sourceChainID, targetChainID)
}

func (oc *Orchestrator) processNextTFuelVoucherBurnEvent(sourceChainID *big.Int, targetChainID *big.Int) {
	targetChainTokenBank := oc.getTFuelTokenBank(targetChainID)
	maxProcessedVoucherBurnNonce, err := targetChainTokenBank.Getmaxprocessedvoucherburnnonce(nil, sourceChainID)
	if err != nil {
		logger.Warnf("Failed to query the max processed TFuel voucher burn nonce for chain: %v", targetChainID.String())
		return // ignore, will retry shortly
	}

	oc.processNextEvent(sourceChainID, targetChainID, score.IMCEventTypeCrossChainVoucherBurnTFuel, maxProcessedVoucherBurnNonce)
}

func (oc *Orchestrator) processNextTNT20VoucherBurnEvent(sourceChainID *big.Int, targetChainID *big.Int) {
	targetChainTokenBank := oc.getTNT20TokenBank(targetChainID)
	maxProcessedVoucherBurnNonce, err := targetChainTokenBank.Getmaxprocessedvoucherburnnonce(nil, sourceChainID)
	if err != nil {
		logger.Warnf("Failed to query the max processed TNT20 voucher burn nonce for chain: %v", targetChainID.String())
		return // ignore, will retry shortly
	}

	oc.processNextEvent(sourceChainID, targetChainID, score.IMCEventTypeCrossChainVoucherBurnTNT20, maxProcessedVoucherBurnNonce)
}

func (oc *Orchestrator) processNextEvent(sourceChainID *big.Int, targetChainID *big.Int, sourceChainEventType score.InterChainMessageEventType, maxProcessedNonce *big.Int) {
	oc.cleanUpInterChainEventCache(sourceChainID, sourceChainEventType, maxProcessedNonce)

	nextNonce := big.NewInt(0).Add(maxProcessedNonce, big.NewInt(1))
	sourceEvent, err := oc.interChainEventCache.Get(sourceChainID, sourceChainEventType, nextNonce)
	if err == ts.ErrKeyNotFound {
		return // the next event (e.g. Token Lock, or Voucher Burn) has not occurred yet
	}

	targetEventType := oc.getTargetChainCorrespondingEventType(sourceChainEventType)
	retryThreshold := oc.getRetryThreshold(targetChainID)
	if oc.timeElapsedSinceEventProcessed(sourceEvent) > retryThreshold { // retry if the tx has been submitted for a long time
		err := oc.callTargetContract(targetChainID, targetEventType, sourceEvent)
		if err == nil {
			oc.updateEventProcessedTimestamp(sourceEvent)
		} else {
			logger.Warnf("Failed to call target contract: %v", err)
		}
	}
}

func (oc *Orchestrator) cleanUpInterChainEventCache(sourceChainID *big.Int, eventType score.InterChainMessageEventType, maxProcessedNonce *big.Int) {
	exists, err := oc.interChainEventCache.Exists(sourceChainID, eventType, maxProcessedNonce)
	if err != nil {
		return
	}
	if exists {
		oc.interChainEventCache.Delete(sourceChainID, eventType, maxProcessedNonce)
	}
}

func (oc *Orchestrator) timeElapsedSinceEventProcessed(event *score.InterChainMessageEvent) time.Duration {
	eventID := event.ID()
	if processedTimestamp, ok := oc.eventProcessedTimestamp[eventID]; ok {
		return time.Since(processedTimestamp)
	} else { // never processed, return a large value
		return time.Since(time.Time{}) // since the Unix epoch start time (0:00:00 Jan 1st, 1970 UTC)
	}
}

func (oc *Orchestrator) updateEventProcessedTimestamp(event *score.InterChainMessageEvent) {
	eventID := event.ID()
	oc.eventProcessedTimestamp[eventID] = time.Now()
}

// For Token Lock events on the source chain, call the Mint Voucher method of the corresponding TokenBank contract on the target chain
// For Voucher Burn events on the source chain, call the Unlock Token method  of the corresponding TokenBank contract on the target chain
func (oc *Orchestrator) callTargetContract(targetChainID *big.Int, targetEventType score.InterChainMessageEventType, sourceEvent *score.InterChainMessageEvent) error {
	var err error

	targetChainEthRpcClient := oc.getEthRpcClient(targetChainID)
	txOpts, err := oc.buildTxOpts(targetChainID, targetChainEthRpcClient)
	if err != nil {
		return err
	}

	switch targetEventType {
	// Voucher Mint events
	case score.IMCEventTypeCrossChainVoucherMintTFuel:
		err = oc.mintTFuelVouchers(txOpts, targetChainID, sourceEvent)
	case score.IMCEventTypeCrossChainVoucherMintTNT20:
		err = oc.mintTNT20Vouchers(txOpts, targetChainID, sourceEvent)

	// Token Unlock events
	case score.IMCEventTypeCrossChainTokenUnlockTFuel:
		err = oc.unlockTFuelTokens(txOpts, targetChainID, sourceEvent)
	case score.IMCEventTypeCrossChainTokenUnlockTNT20:
		err = oc.unlockTNT20Tokens(txOpts, targetChainID, sourceEvent)

	default:
		return nil
	}

	if err != nil {
		logger.Warnf("Failed to call the target contract: ", err)
		return err
	}

	return nil
}

func (oc *Orchestrator) mintTFuelVouchers(txOpts *bind.TransactOpts, targetChainID *big.Int, sourceEvent *score.InterChainMessageEvent) error {
	se, err := score.ParseToCrossChainTFuelTokenLockedEvent(sourceEvent)
	if err != nil {
		return err
	}

	dynasty := oc.getDynasty()
	tfuelTokenBank := oc.getTFuelTokenBank(targetChainID)
	_, err = tfuelTokenBank.Mintvouchers(txOpts, se.Denom, se.TargetChainVoucherReceiver, se.LockedAmount, dynasty, se.TokenLockNonce)
	if err != nil {
		return err
	}

	return nil
}

func (oc *Orchestrator) mintTNT20Vouchers(txOpts *bind.TransactOpts, targetChainID *big.Int, sourceEvent *score.InterChainMessageEvent) error {
	// TODO: implementation
	return nil
}

func (oc *Orchestrator) unlockTFuelTokens(txOpts *bind.TransactOpts, targetChainID *big.Int, sourceEvent *score.InterChainMessageEvent) error {
	// TODO: implementation
	return nil
}

func (oc *Orchestrator) unlockTNT20Tokens(txOpts *bind.TransactOpts, targetChainID *big.Int, sourceEvent *score.InterChainMessageEvent) error {
	// TODO: implementation
	return nil
}

func (oc *Orchestrator) buildTxOpts(chainID *big.Int, ecClient *ec.Client) (*bind.TransactOpts, error) {
	gasPrice, err := ecClient.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	nonce, err := ecClient.PendingNonceAt(context.Background(), oc.privateKey.PublicKey().Address())
	if err != nil {
		return nil, err
	}
	txOpts, err := bind.NewKeyedTransactorWithChainID(oc.privateKey, chainID)
	if err != nil {
		return nil, err
	}
	txOpts.Nonce = big.NewInt(int64(nonce))
	txOpts.Value = big.NewInt(0)       // in wei
	txOpts.GasLimit = uint64(10000000) // in units
	txOpts.GasPrice = gasPrice
	logger.Debugf("building tx opts with address %v", oc.privateKey.PublicKey().Address())
	return txOpts, nil
}

func (oc *Orchestrator) getDynasty() *big.Int {
	return oc.ledger.GetDynasty()
}

func (oc *Orchestrator) getRetryThreshold(chainID *big.Int) time.Duration {
	var blockIntervalInSeconds int
	if chainID.Cmp(oc.mainchainID) == 0 {
		blockIntervalInSeconds = viper.GetInt(scom.CfgSubchainMainchainBlockIntervalInSeconds)
	} else {
		blockIntervalInSeconds = viper.GetInt(scom.CfgConsensusMinBlockInterval)
	}
	numBlocks := 4 // typically a tx should be finalized within 2 block intervals, here we conservatively use 4
	retryThreshold := time.Duration(numBlocks*blockIntervalInSeconds) * time.Second
	return retryThreshold
}

func (oc *Orchestrator) getEthRpcClient(chainID *big.Int) *ec.Client {
	if chainID.Cmp(oc.mainchainID) == 0 {
		return oc.mainchainEthRpcClient
	} else {
		return oc.subchainEthRpcClient
	}
}

func (oc *Orchestrator) getTFuelTokenBank(chainID *big.Int) *scta.TFuelTokenBank {
	if chainID.Cmp(oc.mainchainID) == 0 {
		return oc.mainchainTFuelTokenBank
	} else {
		return oc.subchainTFuelTokenBank
	}
}

func (oc *Orchestrator) getTNT20TokenBank(chainID *big.Int) *scta.TNT20TokenBank {
	if chainID.Cmp(oc.mainchainID) == 0 {
		return oc.mainchainTNT20TokenBank
	} else {
		return oc.subchainTNT20TokenBank
	}
}

func (oc *Orchestrator) getTargetChainCorrespondingEventType(eventType score.InterChainMessageEventType) score.InterChainMessageEventType {
	switch eventType {
	// Token Lock: the corresponding event type on the target chain is Voucher Mint
	case score.IMCEventTypeCrossChainTokenLockTFuel:
		return score.IMCEventTypeCrossChainVoucherMintTFuel
	case score.IMCEventTypeCrossChainTokenLockTNT20:
		return score.IMCEventTypeCrossChainVoucherMintTNT20
	case score.IMCEventTypeCrossChainTokenLockTNT721:
		return score.IMCEventTypeCrossChainVoucherMintTNT721

	// Voucher Burn: the corresponding event type on the target chain is Token Unlock
	case score.IMCEventTypeCrossChainVoucherBurnTFuel:
		return score.IMCEventTypeCrossChainTokenUnlockTFuel
	case score.IMCEventTypeCrossChainVoucherBurnTNT20:
		return score.IMCEventTypeCrossChainTokenUnlockTNT20
	case score.IMCEventTypeCrossChainVoucherBurnTNT721:
		return score.IMCEventTypeCrossChainTokenUnlockTNT721

	default:
		logger.Fatalf("Cannot get the counter event for type: %v", eventType)
	}

	return score.IMCEventTypeUnknown // syntactic sugar
}