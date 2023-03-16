package orchestrator

import (
	"context"
	"errors"
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

var (
	ErrDynastyIsNil        = errors.New("nil dynasty")
	ErrTargetChainMismatch = errors.New("target chain mismatch")
)

type Orchestrator struct {
	updateInterval        int
	privateKey            *crypto.PrivateKey
	ledger                score.Ledger
	eventProcessingTicker *time.Ticker
	metachainWitness      witness.ChainWitness
	eventProcessedTime    map[string]time.Time

	// The mainchain
	mainchainID                   *big.Int
	mainchainEthRpcURL            string
	mainchainEthRpcClient         *ec.Client
	mainchainTFuelTokenBankAddr   common.Address
	mainchainTFuelTokenBank       *scta.TFuelTokenBank
	mainchainTNT20TokenBankAddr   common.Address
	mainchainTNT20TokenBank       *scta.TNT20TokenBank
	mainchainTNT721TokenBankAddr  common.Address
	mainchainTNT721TokenBank      *scta.TNT721TokenBank
	mainchainTNT1155TokenBankAddr common.Address
	mainchainTNT1155TokenBank     *scta.TNT1155TokenBank
	// The subchain
	subchainID                   *big.Int
	subchainEthRpcURL            string
	subchainEthRpcClient         *ec.Client
	subchainTFuelTokenBankAddr   common.Address
	subchainTFuelTokenBank       *scta.TFuelTokenBank
	subchainTNT20TokenBankAddr   common.Address
	subchainTNT20TokenBank       *scta.TNT20TokenBank
	subchainTNT721TokenBankAddr  common.Address
	subchainTNT721TokenBank      *scta.TNT721TokenBank
	subchainTNT1155TokenBankAddr common.Address
	subchainTNT1155TokenBank     *scta.TNT1155TokenBank
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
		logger.Fatalf("failed to get the chainID of the mainchain, is the mainchain RPC API service running? error: %v\n", err)
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
		logger.Fatalf("failed to create MainchainTNT721TokenBank contract: %v\n", err)
	}
	mainchainTNT1155TokenBankAddr := common.HexToAddress(viper.GetString(scom.CfgMainchainTNT1155TokenBankContractAddress))
	mainchainTNT1155TokenBank, err := scta.NewTNT1155TokenBank(mainchainTNT1155TokenBankAddr, mainchainEthRpcClient)
	if err != nil {
		logger.Fatalf("failed to create MainchainTNT1155TokenBank contract %v\n", err)
	}
	subchainID := big.NewInt(viper.GetInt64(scom.CfgSubchainID))
	subchainEthRpcURL := viper.GetString(scom.CfgSubchainEthRpcURL)
	subchainEthRpcClient, err := ec.Dial(subchainEthRpcURL)
	if err != nil {
		logger.Fatalf("the ETH client failed to connect to the subchain ETH RPC: %v\n", err)
	}
	eventProcessedTime := make(map[string]time.Time)
	oc := &Orchestrator{
		updateInterval:     updateInterval,
		privateKey:         privateKey,
		metachainWitness:   metachainWitness,
		eventProcessedTime: eventProcessedTime,

		mainchainID:                   mainchainID,
		mainchainEthRpcURL:            mainchainEthRpcURL,
		mainchainEthRpcClient:         mainchainEthRpcClient,
		mainchainTFuelTokenBankAddr:   mainchainTFuelTokenBankAddr,
		mainchainTFuelTokenBank:       mainchainTFuelTokenBank,
		mainchainTNT20TokenBankAddr:   mainchainTNT20TokenBankAddr,
		mainchainTNT20TokenBank:       mainchainTNT20TokenBank,
		mainchainTNT721TokenBankAddr:  mainchainTNT721TokenBankAddr,
		mainchainTNT721TokenBank:      mainchainTNT721TokenBank,
		mainchainTNT1155TokenBankAddr: mainchainTNT1155TokenBankAddr,
		mainchainTNT1155TokenBank:     mainchainTNT1155TokenBank,

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

	var err error
	subchainTFuelTokenBankAddr := ledger.GetTokenBankContractAddress(score.CrossChainTokenTypeTFuel)
	if subchainTFuelTokenBankAddr == nil {
		logger.Fatalf("failed to obtain SubchainTFuelTokenBank contract address\n")
	}
	oc.subchainTFuelTokenBankAddr = *subchainTFuelTokenBankAddr
	oc.subchainTFuelTokenBank, err = scta.NewTFuelTokenBank(*subchainTFuelTokenBankAddr, oc.subchainEthRpcClient)
	if err != nil {
		logger.Fatalf("failed to set the SubchainTFuelTokenBank contract: %v\n", err)
	}

	subchainTNT20TokenBankAddr := ledger.GetTokenBankContractAddress(score.CrossChainTokenTypeTNT20)
	if subchainTNT20TokenBankAddr == nil {
		logger.Fatalf("failed to obtain SubchainTNT20TokenBank contract address\n")
	}
	oc.subchainTNT20TokenBankAddr = *subchainTNT20TokenBankAddr
	oc.subchainTNT20TokenBank, err = scta.NewTNT20TokenBank(*subchainTNT20TokenBankAddr, oc.subchainEthRpcClient)
	if err != nil {
		logger.Fatalf("failed to set the SubchainTNT20TokenBankAddr contract: %v\n", err)
	}

	subchainTNT721TokenBankAddr := ledger.GetTokenBankContractAddress(score.CrossChainTokenTypeTNT721)
	if subchainTNT721TokenBankAddr == nil {
		logger.Fatalf("failed to obtain SubchainTNT721TokenBank contract address\n")
	}
	oc.subchainTNT721TokenBankAddr = *subchainTNT721TokenBankAddr
	oc.subchainTNT721TokenBank, err = scta.NewTNT721TokenBank(*subchainTNT721TokenBankAddr, oc.subchainEthRpcClient)
	if err != nil {
		logger.Fatalf("failed to set the SubchainTNT721TokenBankAddr contract: %v\n", err)
	}

	subchainTNT1155TokenBankAddr := ledger.GetTokenBankContractAddress(score.CrossChainTokenTypeTNT1155)
	if subchainTNT1155TokenBankAddr == nil || err != nil {
		logger.Fatalf("failed to obtain SubchainTNT1155TokenBank contract address: %v\n", err)
	}
	oc.subchainTNT1155TokenBankAddr = *subchainTNT1155TokenBankAddr
	oc.subchainTNT1155TokenBank, err = scta.NewTNT1155TokenBank(*subchainTNT1155TokenBankAddr, oc.subchainEthRpcClient)
	if err != nil {
		logger.Fatalf("failed to set the SubchainTNT1155TokenBankAddr contract: %v\n", err)
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
	oc.processNextTNT721TokenLockEvent(sourceChainID, targetChainID)
	oc.processNextTNT1155TokenLockEvent(sourceChainID, targetChainID)
}

func (oc *Orchestrator) processNextTFuelTokenLockEvent(sourceChainID *big.Int, targetChainID *big.Int) {
	targetChainTokenBank := oc.getTFuelTokenBank(targetChainID)
	maxProcessedTokenLockNonce, err := targetChainTokenBank.GetMaxProcessedTokenLockNonce(nil, sourceChainID)
	if err != nil {
		logger.Warnf("Failed to query the max processed TFuel token lock nonce for chain: %v, err: %v", targetChainID.String(), err)
		return // ignore
	}

	oc.processNextEvent(sourceChainID, targetChainID, score.IMCEventTypeCrossChainTokenLockTFuel, maxProcessedTokenLockNonce)
}

func (oc *Orchestrator) processNextTNT20TokenLockEvent(sourceChainID *big.Int, targetChainID *big.Int) {
	targetChainTokenBank := oc.getTNT20TokenBank(targetChainID)
	maxProcessedTokenLockNonce, err := targetChainTokenBank.GetMaxProcessedTokenLockNonce(nil, sourceChainID)
	if err != nil {
		logger.Warnf("Failed to query the max processed TNT20 token lock nonce for chain: %v, err: %v", targetChainID.String(), err)
		return // ignore
	}
	oc.processNextEvent(sourceChainID, targetChainID, score.IMCEventTypeCrossChainTokenLockTNT20, maxProcessedTokenLockNonce)
}

func (oc *Orchestrator) processNextTNT721TokenLockEvent(sourceChainID *big.Int, targetChainID *big.Int) {
	targetChainTokenBank := oc.getTNT721TokenBank(targetChainID)
	maxProcessedTokenLockNonce, err := targetChainTokenBank.GetMaxProcessedTokenLockNonce(nil, sourceChainID)
	if err != nil {
		logger.Warnf("Failed to query the max processed TNT721 token lock nonce for chain: %v, err: %v", targetChainID.String(), err)
		return // ignore
	}
	oc.processNextEvent(sourceChainID, targetChainID, score.IMCEventTypeCrossChainTokenLockTNT721, maxProcessedTokenLockNonce)
}

func (oc *Orchestrator) processNextTNT1155TokenLockEvent(sourceChainID *big.Int, targetChainID *big.Int) {
	targetChainTokenBank := oc.getTNT1155TokenBank(targetChainID)
	maxProcessedTokenLockNonce, err := targetChainTokenBank.GetMaxProcessedTokenLockNonce(nil, sourceChainID)
	if err != nil {
		logger.Warnf("Failed to query the max processed TNT1155 token lock nonce for chain: %v, err: %v", targetChainID.String(), err)
		return // ignore
	}
	oc.processNextEvent(sourceChainID, targetChainID, score.IMCEventTypeCrossChainTokenLockTNT1155, maxProcessedTokenLockNonce)
}

func (oc *Orchestrator) processNextVoucherBurnEvent(sourceChainID *big.Int, targetChainID *big.Int) {
	oc.processNextTFuelVoucherBurnEvent(sourceChainID, targetChainID)
	oc.processNextTNT20VoucherBurnEvent(sourceChainID, targetChainID)
	oc.processNextTNT721VoucherBurnEvent(sourceChainID, targetChainID)
	oc.processNextTNT1155VoucherBurnEvent(sourceChainID, targetChainID)
}

func (oc *Orchestrator) processNextTFuelVoucherBurnEvent(sourceChainID *big.Int, targetChainID *big.Int) {
	targetChainTokenBank := oc.getTFuelTokenBank(targetChainID)
	maxProcessedVoucherBurnNonce, err := targetChainTokenBank.GetMaxProcessedVoucherBurnNonce(nil, sourceChainID)
	if err != nil {
		logger.Warnf("Failed to query the max processed TFuel voucher burn nonce for chain: %v", targetChainID.String())
		return // ignore
	}

	oc.processNextEvent(sourceChainID, targetChainID, score.IMCEventTypeCrossChainVoucherBurnTFuel, maxProcessedVoucherBurnNonce)
}

func (oc *Orchestrator) processNextTNT20VoucherBurnEvent(sourceChainID *big.Int, targetChainID *big.Int) {
	targetChainTokenBank := oc.getTNT20TokenBank(targetChainID)
	maxProcessedVoucherBurnNonce, err := targetChainTokenBank.GetMaxProcessedVoucherBurnNonce(nil, sourceChainID)
	if err != nil {
		logger.Warnf("Failed to query the max processed TNT20 voucher burn nonce for chain: %v", targetChainID.String())
		return // ignore
	}

	oc.processNextEvent(sourceChainID, targetChainID, score.IMCEventTypeCrossChainVoucherBurnTNT20, maxProcessedVoucherBurnNonce)
}

func (oc *Orchestrator) processNextTNT721VoucherBurnEvent(sourceChainID *big.Int, targetChainID *big.Int) {
	targetChainTokenBank := oc.getTNT721TokenBank(targetChainID)
	maxProcessedVoucherBurnNonce, err := targetChainTokenBank.GetMaxProcessedVoucherBurnNonce(nil, sourceChainID)
	if err != nil {
		logger.Warnf("Failed to query the max processed TNT721 voucher burn nonce for chain: %v", targetChainID.String())
		return // ignore
	}

	oc.processNextEvent(sourceChainID, targetChainID, score.IMCEventTypeCrossChainVoucherBurnTNT721, maxProcessedVoucherBurnNonce)
}

func (oc *Orchestrator) processNextTNT1155VoucherBurnEvent(sourceChainID *big.Int, targetChainID *big.Int) {
	targetChainTokenBank := oc.getTNT1155TokenBank(targetChainID)
	maxProcessedVoucherBurnNonce, err := targetChainTokenBank.GetMaxProcessedVoucherBurnNonce(nil, sourceChainID)
	if err != nil {
		logger.Warnf("Failed to query the max processed TNT1155 voucher burn nonce for chain: %v", targetChainID.String())
		return // ignore
	}

	oc.processNextEvent(sourceChainID, targetChainID, score.IMCEventTypeCrossChainVoucherBurnTNT1155, maxProcessedVoucherBurnNonce)
}

func (oc *Orchestrator) processNextEvent(sourceChainID *big.Int, targetChainID *big.Int, sourceChainEventType score.InterChainMessageEventType, maxProcessedNonce *big.Int) {
	oc.cleanUpInterChainEventCache(sourceChainID, sourceChainEventType, maxProcessedNonce)

	nextNonce := big.NewInt(0).Add(maxProcessedNonce, big.NewInt(1))
	sourceEvent, err := oc.interChainEventCache.Get(sourceChainID, sourceChainEventType, nextNonce)
	if err == ts.ErrKeyNotFound {
		return // the next event (e.g. Token Lock, or Voucher Burn) has not occurred yet
	}

	logger.Debugf("Process next event, sourceChainID: %v, targetChainID: %v, sourceChainEventType: %v, nextNonce: %v",
		sourceChainID, targetChainID, sourceChainEventType, nextNonce)

	targetEventType := oc.getTargetChainCorrespondingEventType(sourceChainEventType)
	retryThreshold := oc.getRetryThreshold(targetChainID)
	if oc.timeElapsedSinceEventProcessed(sourceEvent) > retryThreshold { // retry if the tx has been submitted for a long time
		err := oc.callTargetContract(targetChainID, targetEventType, sourceEvent)
		if err == nil {
			oc.updateEventProcessedTime(sourceEvent)
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
	if processedTime, ok := oc.eventProcessedTime[eventID]; ok {
		return time.Since(processedTime)
	} else { // never processed, return a large value
		return time.Since(time.Time{}) // since the Unix epoch start time (0:00:00 Jan 1st, 1970 UTC)
	}
}

func (oc *Orchestrator) updateEventProcessedTime(event *score.InterChainMessageEvent) {
	eventID := event.ID()
	oc.eventProcessedTime[eventID] = time.Now()
}

// For Token Lock events on the source chain, call the Mint Voucher method of the corresponding TokenBank contract on the target chain
// For Voucher Burn events on the source chain, call the Unlock Token method  of the corresponding TokenBank contract on the target chain
func (oc *Orchestrator) callTargetContract(targetChainID *big.Int, targetEventType score.InterChainMessageEventType, sourceEvent *score.InterChainMessageEvent) error {
	var err error
	var txHash common.Hash

	dynasty := oc.getDynasty()
	if dynasty != nil {
		logger.Infof("calling contracts on target chain %v for event type %v, current dynasty: %v", targetChainID, targetEventType, dynasty)

		vsQueriedFromMC, _ := oc.mainchainTFuelTokenBank.GetAdjustedValidatorSet(nil, oc.subchainID, dynasty)
		vsQueriedFromSC, _ := oc.subchainTNT20TokenBank.GetAdjustedValidatorSet(nil, oc.subchainID, dynasty)
		logger.Debugf("Subchain %v adjusted ValSet queried from the Mainchain for dynasty %v: %v", oc.subchainID, dynasty, vsQueriedFromMC)
		logger.Debugf("Subchain %v adjusted ValSet queried from the Subchain  for dynasty %v: %v", oc.subchainID, dynasty, vsQueriedFromSC)
	}

	targetChainEthRpcClient := oc.getEthRpcClient(targetChainID)
	txOpts, err := oc.buildTxOpts(targetChainID, targetChainEthRpcClient)
	if err != nil {
		return err
	}
	switch targetEventType {
	// Voucher Mint events
	case score.IMCEventTypeCrossChainVoucherMintTFuel:
		txHash, err = oc.mintTFuelVouchers(txOpts, targetChainID, sourceEvent)
	case score.IMCEventTypeCrossChainVoucherMintTNT20:
		txHash, err = oc.mintTNT20Vouchers(txOpts, targetChainID, sourceEvent)
	case score.IMCEventTypeCrossChainVoucherMintTNT721:
		txHash, err = oc.mintTN721Vouchers(txOpts, targetChainID, sourceEvent)
	case score.IMCEventTypeCrossChainVoucherMintTNT1155:
		txHash, err = oc.mintTN1155Vouchers(txOpts, targetChainID, sourceEvent)

	// Token Unlock events
	case score.IMCEventTypeCrossChainTokenUnlockTFuel:
		txHash, err = oc.unlockTFuelTokens(txOpts, targetChainID, sourceEvent)
	case score.IMCEventTypeCrossChainTokenUnlockTNT20:
		txHash, err = oc.unlockTNT20Tokens(txOpts, targetChainID, sourceEvent)
	case score.IMCEventTypeCrossChainTokenUnlockTNT721:
		txHash, err = oc.unlockTNT721Tokens(txOpts, targetChainID, sourceEvent)
	case score.IMCEventTypeCrossChainTokenUnlockTNT1155:
		txHash, err = oc.unlockTNT1155Tokens(txOpts, targetChainID, sourceEvent)
	default:
		return nil
	}

	if err == ErrTargetChainMismatch {
		// this may happen when a user sends tokens from a subchain to another subchain, which is not supported yet
		// the current orchestrator only supports transfers between the mainchain and a subchain
		logger.Warnf("Sending tokens between two subchains is not supported yet: %v", err)
		return nil // ignore
	}

	if err != nil {
		logger.Warnf("Failed to call the target contract: %v", err)
		return err
	}

	logger.Infof("contract call tx hash: %v, chain: %v", txHash.Hex(), targetChainID)

	return nil
}

func (oc *Orchestrator) mintTFuelVouchers(txOpts *bind.TransactOpts, targetChainID *big.Int, sourceEvent *score.InterChainMessageEvent) (common.Hash, error) {
	se, err := score.ParseToCrossChainTFuelTokenLockedEvent(sourceEvent)
	if err != nil {
		return common.Hash{}, err
	}
	if targetChainID.Cmp(se.TargetChainID) != 0 {
		logger.Warnf("mintTFuelVouchers, target chain ID mismatch: %v vs %v", targetChainID, se.TargetChainID)
		return common.Hash{}, ErrTargetChainMismatch
	}
	dynasty := oc.getDynasty()
	if dynasty == nil {
		return common.Hash{}, ErrDynastyIsNil
	}
	tfuelTokenBank := oc.getTFuelTokenBank(targetChainID)
	tx, err := tfuelTokenBank.MintVouchers(txOpts, se.Denom, se.TargetChainVoucherReceiver, se.LockedAmount, dynasty, se.TokenLockNonce)
	if err != nil {
		return common.Hash{}, err
	}
	txHash := tx.Hash()
	logger.Debugf("mintTFuelVouchers, dynasty: %v, targetChainID: %v, denom: %v, tokenLockNonce: %v, tx: %v", dynasty, targetChainID, se.Denom, se.TokenLockNonce, txHash.Hex())
	return txHash, nil
}

func (oc *Orchestrator) mintTNT20Vouchers(txOpts *bind.TransactOpts, targetChainID *big.Int, sourceEvent *score.InterChainMessageEvent) (common.Hash, error) {
	se, err := score.ParseToCrossChainTNT20TokenLockedEvent(sourceEvent)
	if err != nil {
		return common.Hash{}, err
	}
	if targetChainID.Cmp(se.TargetChainID) != 0 {
		logger.Warnf("mintTNT20Vouchers, target chain ID mismatch: %v vs %v", targetChainID, se.TargetChainID)
		return common.Hash{}, ErrTargetChainMismatch
	}
	dynasty := oc.getDynasty()
	if dynasty == nil {
		return common.Hash{}, ErrDynastyIsNil
	}
	TNT20TokenBank := oc.getTNT20TokenBank(targetChainID)
	tx, err := TNT20TokenBank.MintVouchers(txOpts, se.Denom, se.Name, se.Symbol, se.Decimals, se.TargetChainVoucherReceiver, se.LockedAmount, dynasty, se.TokenLockNonce)
	if err != nil {
		return common.Hash{}, err
	}
	txHash := tx.Hash()
	logger.Debugf("mintTNT20Vouchers, dynasty: %v, targetChainID: %v, denom: %v, tokenLockNonce: %v, tx: %v", dynasty, targetChainID, se.Denom, se.TokenLockNonce, txHash.Hex())
	return txHash, nil
}

func (oc *Orchestrator) mintTN721Vouchers(txOpts *bind.TransactOpts, targetChainID *big.Int, sourceEvent *score.InterChainMessageEvent) (common.Hash, error) {
	se, err := score.ParseToCrossChainTNT721TokenLockedEvent(sourceEvent)
	if err != nil {
		return common.Hash{}, err
	}
	if targetChainID.Cmp(se.TargetChainID) != 0 {
		logger.Warnf("mintTN721Vouchers, target chain ID mismatch: %v vs %v", targetChainID, se.TargetChainID)
		return common.Hash{}, ErrTargetChainMismatch
	}
	dynasty := oc.getDynasty()
	if dynasty == nil {
		return common.Hash{}, ErrDynastyIsNil
	}
	TNT721TokenBank := oc.getTNT721TokenBank(targetChainID)
	tx, err := TNT721TokenBank.MintVouchers(txOpts, se.Denom, se.Name, se.Symbol, se.TargetChainVoucherReceiver, se.TokenID, se.TokenURI, dynasty, se.TokenLockNonce)
	if err != nil {
		return common.Hash{}, err
	}
	txHash := tx.Hash()
	logger.Debugf("mintTN721Vouchers, dynasty: %v, targetChainID: %v, denom: %v, tokenLockNonce: %v, tx: %v", dynasty, targetChainID, se.Denom, se.TokenLockNonce, txHash.Hex())
	return txHash, nil
}

func (oc *Orchestrator) mintTN1155Vouchers(txOpts *bind.TransactOpts, targetChainID *big.Int, sourceEvent *score.InterChainMessageEvent) (common.Hash, error) {
	se, err := score.ParseToCrossChainTNT1155TokenLockedEvent(sourceEvent)
	if err != nil {
		return common.Hash{}, err
	}
	if targetChainID.Cmp(se.TargetChainID) != 0 {
		logger.Warnf("mintTN1155Vouchers, target chain ID mismatch: %v vs %v", targetChainID, se.TargetChainID)
		return common.Hash{}, ErrTargetChainMismatch
	}
	dynasty := oc.getDynasty()
	if dynasty == nil {
		return common.Hash{}, ErrDynastyIsNil
	}
	TNT1155TokenBank := oc.getTNT1155TokenBank(targetChainID)
	tx, err := TNT1155TokenBank.MintVouchers(txOpts, se.Denom, se.TargetChainVoucherReceiver, se.TokenID, se.LockedAmount, se.TokenURI, dynasty, se.TokenLockNonce)
	if err != nil {
		return common.Hash{}, err
	}
	txHash := tx.Hash()
	logger.Debugf("se.TargetChainID: %v", se.TargetChainID)
	logger.Debugf("mintTN1155Vouchers, dynasty: %v, targetChainID: %v, denom: %v, tokenLockNonce: %v, tx: %v", dynasty, targetChainID, se.Denom, se.TokenLockNonce, txHash.Hex())
	return txHash, nil
}

func (oc *Orchestrator) unlockTFuelTokens(txOpts *bind.TransactOpts, targetChainID *big.Int, sourceEvent *score.InterChainMessageEvent) (common.Hash, error) {
	se, err := score.ParseToCrossChainTFuelVoucherBurnedEvent(sourceEvent)
	if err != nil {
		return common.Hash{}, err
	}
	dynasty := oc.getDynasty()
	if dynasty == nil {
		return common.Hash{}, ErrDynastyIsNil
	}
	tfuelTokenBank := oc.getTFuelTokenBank(targetChainID)
	tx, err := tfuelTokenBank.UnlockTokens(txOpts, sourceEvent.SourceChainID, se.TargetChainTokenReceiver, se.BurnedAmount, dynasty, se.VoucherBurnNonce)
	if err != nil {
		return common.Hash{}, err
	}
	txHash := tx.Hash()
	logger.Debugf("unlockTFuelTokens, dynasty: %v, targetChainID: %v, denom: %v, tokenLockNonce: %v, tx: %v", dynasty, targetChainID, se.Denom, se.VoucherBurnNonce, txHash.Hex())
	return txHash, nil
}

func (oc *Orchestrator) unlockTNT20Tokens(txOpts *bind.TransactOpts, targetChainID *big.Int, sourceEvent *score.InterChainMessageEvent) (common.Hash, error) {
	se, err := score.ParseToCrossChainTNT20VoucherBurnedEvent(sourceEvent)
	if err != nil {
		return common.Hash{}, err
	}
	dynasty := oc.getDynasty()
	if dynasty == nil {
		return common.Hash{}, ErrDynastyIsNil
	}
	TNT20TokenBank := oc.getTNT20TokenBank(targetChainID)
	tx, err := TNT20TokenBank.UnlockTokens(txOpts, sourceEvent.SourceChainID, se.Denom, se.TargetChainTokenReceiver, se.BurnedAmount, dynasty, se.VoucherBurnNonce)
	if err != nil {
		return common.Hash{}, err
	}
	txHash := tx.Hash()
	logger.Debugf("unlockTNT20Tokens, dynasty: %v, targetChainID: %v, denom: %v, tokenLockNonce: %v, tx: %v", dynasty, targetChainID, se.Denom, se.VoucherBurnNonce, txHash.Hex())
	return txHash, nil
}

func (oc *Orchestrator) unlockTNT721Tokens(txOpts *bind.TransactOpts, targetChainID *big.Int, sourceEvent *score.InterChainMessageEvent) (common.Hash, error) {
	se, err := score.ParseToCrossChainTNT721VoucherBurnedEvent(sourceEvent)
	if err != nil {
		return common.Hash{}, err
	}
	dynasty := oc.getDynasty()
	if dynasty == nil {
		return common.Hash{}, ErrDynastyIsNil
	}
	TNT721TokenBank := oc.getTNT721TokenBank(targetChainID)
	tx, err := TNT721TokenBank.UnlockTokens(txOpts, sourceEvent.SourceChainID, se.Denom, se.TargetChainTokenReceiver, se.TokenID, dynasty, se.VoucherBurnNonce)
	if err != nil {
		return common.Hash{}, err
	}
	txHash := tx.Hash()
	logger.Debugf("unlockTNT721Tokens, dynasty: %v, targetChainID: %v, denom: %v, tokenLockNonce: %v, tx: %v", dynasty, targetChainID, se.Denom, se.VoucherBurnNonce, txHash.Hex())
	return txHash, nil
}

func (oc *Orchestrator) unlockTNT1155Tokens(txOpts *bind.TransactOpts, targetChainID *big.Int, sourceEvent *score.InterChainMessageEvent) (common.Hash, error) {
	se, err := score.ParseToCrossChainTNT1155VoucherBurnedEvent(sourceEvent)
	if err != nil {
		return common.Hash{}, err
	}
	dynasty := oc.getDynasty()
	if dynasty == nil {
		return common.Hash{}, ErrDynastyIsNil
	}
	TNT1155TokenBank := oc.getTNT1155TokenBank(targetChainID)
	tx, err := TNT1155TokenBank.UnlockTokens(txOpts, sourceEvent.SourceChainID, se.Denom, se.TargetChainTokenReceiver, se.TokenID, se.BurnedAmount, dynasty, se.VoucherBurnNonce)
	if err != nil {
		return common.Hash{}, err
	}
	txHash := tx.Hash()
	logger.Debugf("unlockTNT1155Tokens, dynasty: %v, targetChainID: %v, denom: %v, tokenLockNonce: %v, tx: %v", dynasty, targetChainID, se.Denom, se.VoucherBurnNonce, txHash.Hex())
	return txHash, nil
}

func (oc *Orchestrator) buildTxOpts(chainID *big.Int, ecClient *ec.Client) (*bind.TransactOpts, error) {
	var gasPrice *big.Int
	var err error
	if chainID.Cmp(oc.mainchainID) == 0 {
		gasPrice, err = ecClient.SuggestGasPrice(context.Background())
		if err != nil {
			return nil, err
		}
	} else {
		// eth_gasPrice returns a hardcoded nubmer for the mainchain, which could be much higher than min gasPrice required by the subchain
		// TODO: parameterize the subchain ETH RPC service to suggest the proper gasPrice for different chains
		// gasPrice = big.NewInt(int64(scom.MinimumGasPrice) * 2)
		gasPrice = common.Big0
	}

	nonce, err := ecClient.PendingNonceAt(context.Background(), oc.privateKey.PublicKey().Address())
	if err != nil {
		return nil, err
	}
	txOpts, err := bind.NewKeyedTransactorWithChainID(oc.privateKey, chainID) //chainID)
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

func (oc *Orchestrator) getTNT721TokenBank(chainID *big.Int) *scta.TNT721TokenBank {
	if chainID.Cmp(oc.mainchainID) == 0 {
		return oc.mainchainTNT721TokenBank
	} else {
		return oc.subchainTNT721TokenBank
	}
}

func (oc *Orchestrator) getTNT1155TokenBank(chainID *big.Int) *scta.TNT1155TokenBank {
	if chainID.Cmp(oc.mainchainID) == 0 {
		return oc.mainchainTNT1155TokenBank
	} else {
		return oc.subchainTNT1155TokenBank
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
	case score.IMCEventTypeCrossChainTokenLockTNT1155:
		return score.IMCEventTypeCrossChainVoucherMintTNT1155

	// Voucher Burn: the corresponding event type on the target chain is Token Unlock
	case score.IMCEventTypeCrossChainVoucherBurnTFuel:
		return score.IMCEventTypeCrossChainTokenUnlockTFuel
	case score.IMCEventTypeCrossChainVoucherBurnTNT20:
		return score.IMCEventTypeCrossChainTokenUnlockTNT20
	case score.IMCEventTypeCrossChainVoucherBurnTNT721:
		return score.IMCEventTypeCrossChainTokenUnlockTNT721
	case score.IMCEventTypeCrossChainVoucherBurnTNT1155:
		return score.IMCEventTypeCrossChainTokenUnlockTNT1155

	default:
		logger.Fatalf("Cannot get the counter event for type: %v", eventType)
	}

	return score.IMCEventTypeUnknown // syntactic sugar
}
