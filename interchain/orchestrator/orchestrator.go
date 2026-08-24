package orchestrator

import (
	"context"
	"errors"
	"fmt"
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
	ethereum "github.com/thetatoken/thetasubchain/eth"
	"github.com/thetatoken/thetasubchain/eth/core/types"
	ec "github.com/thetatoken/thetasubchain/eth/ethclient"
)

var logger *log.Entry = log.WithFields(log.Fields{"prefix": "orchestrator"})

// defaultRelayDryRunTimeout bounds both the eth_call dry run and the broadcast that
// follows it, in case the configured value is missing or nonsensical.
const defaultRelayDryRunTimeout = 5 * time.Second

// defaultUncorroboratedEventQuarantine is how long an event must stay contradicted by
// source chain state before it is discarded. It is deliberately generous: discarding
// is irreversible and jams the nonce sequence, whereas holding a forged event costs
// only a periodic re-check, and the event is never relayed while it is held.
const defaultUncorroboratedEventQuarantine = 30 * time.Minute

// relayRpcTimeout is the bound applied to every RPC read on the relay hot path.
//
// Theta's eth_call retries internally with a one-block sleep between attempts, so an
// unresponsive node holds an unbounded read open for tens of seconds. The processing
// tick is strictly sequential across four asset classes in both directions, so a
// single unbounded read starves every relay behind it. Bounding the dry run alone is
// not sufficient: the nonce and collateral reads run *before* it, on every tick.
func relayRpcTimeout() time.Duration {
	timeout := time.Duration(viper.GetInt(scom.CfgSubchainRelayDryRunTimeoutInSeconds)) * time.Second
	if timeout <= 0 {
		timeout = defaultRelayDryRunTimeout
	}
	return timeout
}

// boundedCallOpts returns CallOpts carrying a deadline. The caller must invoke the
// returned cancel func, conventionally with defer.
func boundedCallOpts() (*bind.CallOpts, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), relayRpcTimeout())
	return &bind.CallOpts{Context: ctx}, cancel
}

var (
	ErrDynastyIsNil           = errors.New("nil dynasty")
	ErrTargetChainMismatch    = errors.New("target chain mismatch")
	ErrTxWouldRevert          = errors.New("the relay transaction would revert, not broadcasting it")
	ErrDryRunTimedOut         = errors.New("the relay dry run did not complete in time, not broadcasting it")
	ErrUnlockUncollateralized = errors.New("the target chain TokenBank cannot cover the requested unlock")
)

type Orchestrator struct {
	updateInterval     int
	privateKey         *crypto.PrivateKey
	ledger             score.Ledger
	metachainWitness   witness.ChainWitness
	eventProcessedTime map[string]time.Time
	// firstContradictedTime records when an event was first found to be contradicted
	// by source chain state, keyed the same way as eventProcessedTime. See
	// quarantineExpired().
	firstContradictedTime map[string]time.Time

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
		updateInterval:        updateInterval,
		privateKey:            privateKey,
		metachainWitness:      metachainWitness,
		eventProcessedTime:    eventProcessedTime,
		firstContradictedTime: make(map[string]time.Time),

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
	// Pass the derived context, not the parent: cancel() cancels c, so a mainloop
	// selecting on the parent would never see Stop() and Wait() would hang. In
	// production Node.Stop() happens to cancel the shared parent, which masked this.
	go oc.mainloop(c)
	oc.logRelayMatrix()
	logger.Info("Metachain orchestrator started")
}

// logRelayMatrix records the relay configuration this node actually came up with.
//
// The per-path switches are the only control surface available while the TokenBank
// contracts cannot be changed, so an operator has to be able to confirm which paths
// are live from the node's own output rather than by inferring it from a config file
// that may not be the one loaded.
func (oc *Orchestrator) logRelayMatrix() {
	if !viper.GetBool(scom.CfgSubchainRelayEnabled) {
		logger.Warnf("RELAY MATRIX: relaying is DISABLED for all assets and directions (%v=false)",
			scom.CfgSubchainRelayEnabled)
		return
	}

	for _, asset := range []string{"tfuel", "tnt20", "tnt721", "tnt1155"} {
		for _, direction := range []string{"inbound", "outbound"} {
			key := fmt.Sprintf("%v.%v.%v", scom.CfgSubchainRelayPathPrefix, asset, direction)
			enabled := true
			if viper.IsSet(key) {
				enabled = viper.GetBool(key)
			}
			state := "ENABLED"
			if !enabled {
				state = "disabled"
			}
			logger.Infof("RELAY MATRIX: %-8v %-8v %v", asset, direction, state)
		}
	}
	logger.Infof("RELAY MATRIX: collateral guard=%v, dry-run/RPC timeout=%v, uncorroborated quarantine=%v",
		viper.GetBool(scom.CfgSubchainEnforceUnlockCollateral), relayRpcTimeout(), (&Orchestrator{}).getQuarantinePeriod())
}

func (oc *Orchestrator) Stop() {
	// See MetachainWitness.Stop(): the ticker belongs to mainloop.
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
	defer oc.wg.Done() // Start() does wg.Add(1); without this Wait() blocks forever
	ticker := time.NewTicker(time.Duration(oc.updateInterval) * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if !viper.GetBool(scom.CfgSubchainRelayEnabled) {
				// Relaying is suspended by configuration. The witness keeps
				// collecting events, so the pipeline resumes where it left off
				// once relaying is re-enabled.
				continue
			}

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
	if !oc.relayPathEnabled(sourceChainID, score.IMCEventTypeCrossChainTokenLockTFuel) {
		return // suspended by configuration: do not spend an RPC round trip on it
	}

	targetChainTokenBank := oc.getTFuelTokenBank(targetChainID)
	opts, cancel := boundedCallOpts()
	defer cancel()
	maxProcessedTokenLockNonce, err := targetChainTokenBank.GetMaxProcessedTokenLockNonce(opts, sourceChainID)
	if err != nil {
		logger.Warnf("Failed to query the max processed TFuel token lock nonce for chain: %v, err: %v", targetChainID.String(), err)
		return // ignore
	}

	oc.processNextEvent(sourceChainID, targetChainID, score.IMCEventTypeCrossChainTokenLockTFuel, maxProcessedTokenLockNonce)
}

func (oc *Orchestrator) processNextTNT20TokenLockEvent(sourceChainID *big.Int, targetChainID *big.Int) {
	if !oc.relayPathEnabled(sourceChainID, score.IMCEventTypeCrossChainTokenLockTNT20) {
		return // suspended by configuration: do not spend an RPC round trip on it
	}

	targetChainTokenBank := oc.getTNT20TokenBank(targetChainID)
	opts, cancel := boundedCallOpts()
	defer cancel()
	maxProcessedTokenLockNonce, err := targetChainTokenBank.GetMaxProcessedTokenLockNonce(opts, sourceChainID)
	if err != nil {
		logger.Warnf("Failed to query the max processed TNT20 token lock nonce for chain: %v, err: %v", targetChainID.String(), err)
		return // ignore
	}
	oc.processNextEvent(sourceChainID, targetChainID, score.IMCEventTypeCrossChainTokenLockTNT20, maxProcessedTokenLockNonce)
}

func (oc *Orchestrator) processNextTNT721TokenLockEvent(sourceChainID *big.Int, targetChainID *big.Int) {
	if !oc.relayPathEnabled(sourceChainID, score.IMCEventTypeCrossChainTokenLockTNT721) {
		return // suspended by configuration: do not spend an RPC round trip on it
	}

	targetChainTokenBank := oc.getTNT721TokenBank(targetChainID)
	opts, cancel := boundedCallOpts()
	defer cancel()
	maxProcessedTokenLockNonce, err := targetChainTokenBank.GetMaxProcessedTokenLockNonce(opts, sourceChainID)
	if err != nil {
		logger.Warnf("Failed to query the max processed TNT721 token lock nonce for chain: %v, err: %v", targetChainID.String(), err)
		return // ignore
	}
	oc.processNextEvent(sourceChainID, targetChainID, score.IMCEventTypeCrossChainTokenLockTNT721, maxProcessedTokenLockNonce)
}

func (oc *Orchestrator) processNextTNT1155TokenLockEvent(sourceChainID *big.Int, targetChainID *big.Int) {
	if !oc.relayPathEnabled(sourceChainID, score.IMCEventTypeCrossChainTokenLockTNT1155) {
		return // suspended by configuration: do not spend an RPC round trip on it
	}

	targetChainTokenBank := oc.getTNT1155TokenBank(targetChainID)
	opts, cancel := boundedCallOpts()
	defer cancel()
	maxProcessedTokenLockNonce, err := targetChainTokenBank.GetMaxProcessedTokenLockNonce(opts, sourceChainID)
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
	if !oc.relayPathEnabled(sourceChainID, score.IMCEventTypeCrossChainVoucherBurnTFuel) {
		return // suspended by configuration: do not spend an RPC round trip on it
	}

	targetChainTokenBank := oc.getTFuelTokenBank(targetChainID)
	opts, cancel := boundedCallOpts()
	defer cancel()
	maxProcessedVoucherBurnNonce, err := targetChainTokenBank.GetMaxProcessedVoucherBurnNonce(opts, sourceChainID)
	if err != nil {
		logger.Warnf("Failed to query the max processed TFuel voucher burn nonce for chain: %v", targetChainID.String())
		return // ignore
	}

	oc.processNextEvent(sourceChainID, targetChainID, score.IMCEventTypeCrossChainVoucherBurnTFuel, maxProcessedVoucherBurnNonce)
}

func (oc *Orchestrator) processNextTNT20VoucherBurnEvent(sourceChainID *big.Int, targetChainID *big.Int) {
	if !oc.relayPathEnabled(sourceChainID, score.IMCEventTypeCrossChainVoucherBurnTNT20) {
		return // suspended by configuration: do not spend an RPC round trip on it
	}

	targetChainTokenBank := oc.getTNT20TokenBank(targetChainID)
	opts, cancel := boundedCallOpts()
	defer cancel()
	maxProcessedVoucherBurnNonce, err := targetChainTokenBank.GetMaxProcessedVoucherBurnNonce(opts, sourceChainID)
	if err != nil {
		logger.Warnf("Failed to query the max processed TNT20 voucher burn nonce for chain: %v", targetChainID.String())
		return // ignore
	}

	oc.processNextEvent(sourceChainID, targetChainID, score.IMCEventTypeCrossChainVoucherBurnTNT20, maxProcessedVoucherBurnNonce)
}

func (oc *Orchestrator) processNextTNT721VoucherBurnEvent(sourceChainID *big.Int, targetChainID *big.Int) {
	if !oc.relayPathEnabled(sourceChainID, score.IMCEventTypeCrossChainVoucherBurnTNT721) {
		return // suspended by configuration: do not spend an RPC round trip on it
	}

	targetChainTokenBank := oc.getTNT721TokenBank(targetChainID)
	opts, cancel := boundedCallOpts()
	defer cancel()
	maxProcessedVoucherBurnNonce, err := targetChainTokenBank.GetMaxProcessedVoucherBurnNonce(opts, sourceChainID)
	if err != nil {
		logger.Warnf("Failed to query the max processed TNT721 voucher burn nonce for chain: %v", targetChainID.String())
		return // ignore
	}

	oc.processNextEvent(sourceChainID, targetChainID, score.IMCEventTypeCrossChainVoucherBurnTNT721, maxProcessedVoucherBurnNonce)
}

func (oc *Orchestrator) processNextTNT1155VoucherBurnEvent(sourceChainID *big.Int, targetChainID *big.Int) {
	if !oc.relayPathEnabled(sourceChainID, score.IMCEventTypeCrossChainVoucherBurnTNT1155) {
		return // suspended by configuration: do not spend an RPC round trip on it
	}

	targetChainTokenBank := oc.getTNT1155TokenBank(targetChainID)
	opts, cancel := boundedCallOpts()
	defer cancel()
	maxProcessedVoucherBurnNonce, err := targetChainTokenBank.GetMaxProcessedVoucherBurnNonce(opts, sourceChainID)
	if err != nil {
		logger.Warnf("Failed to query the max processed TNT1155 voucher burn nonce for chain: %v", targetChainID.String())
		return // ignore
	}

	oc.processNextEvent(sourceChainID, targetChainID, score.IMCEventTypeCrossChainVoucherBurnTNT1155, maxProcessedVoucherBurnNonce)
}

func (oc *Orchestrator) processNextEvent(sourceChainID *big.Int, targetChainID *big.Int, sourceChainEventType score.InterChainMessageEventType, maxProcessedNonce *big.Int) {
	if !oc.relayPathEnabled(sourceChainID, sourceChainEventType) {
		return // this asset class is suspended in this direction by configuration
	}

	oc.cleanUpInterChainEventCache(sourceChainID, sourceChainEventType, maxProcessedNonce)

	nextNonce := big.NewInt(0).Add(maxProcessedNonce, big.NewInt(1))
	sourceEvent, err := oc.interChainEventCache.Get(sourceChainID, sourceChainEventType, nextNonce)

	logger.Debugf("Query next event, sourceChainID: %v, targetChainID: %v, sourceChainEventType: %v, nextNonce: %v, err: %v",
		sourceChainID, targetChainID, sourceChainEventType, nextNonce, err)

	if err == ts.ErrKeyNotFound {
		return // the next event (e.g. Token Lock, or Voucher Burn) has not occurred yet
	}

	logger.Debugf("Process next event, sourceChainID: %v, targetChainID: %v, sourceChainEventType: %v, nextNonce: %v",
		sourceChainID, targetChainID, sourceChainEventType, nextNonce)

	targetEventType := oc.getTargetChainCorrespondingEventType(sourceChainEventType)
	retryThreshold := oc.getRetryThreshold(targetChainID)
	if oc.timeElapsedSinceEventProcessed(sourceEvent) > retryThreshold { // retry if the tx has been submitted for a long time
		// Never vote on an event that the source chain's own contract state does
		// not confirm. A log is not sufficient evidence on its own that the
		// corresponding lock or burn was committed.
		if err := oc.verifyEventAgainstSourceChain(sourceChainID, sourceEvent); err != nil {
			if siu.IsEventNotCorroborated(err) {
				// Quarantine rather than delete. A contradiction can be produced by a
				// lagging or load-balanced RPC backend as easily as by a forged event,
				// and deleting on the first such read discards a genuine transfer
				// permanently: the scan window has moved on and nothing rescans it.
				// The event is not relayed either way -- this path never votes -- so
				// holding it costs nothing but a retry.
				if oc.quarantineExpired(sourceEvent) {
					logger.Errorf("Discarding an inter-chain event that source chain state has contradicted "+
						"for longer than the quarantine period: %v", err)
					oc.interChainEventCache.Delete(sourceChainID, sourceChainEventType, nextNonce)
					oc.clearContradicted(sourceEvent)
				} else {
					logger.Errorf("Refusing to relay an inter-chain event that the source chain state does not "+
						"confirm; holding it for re-check: %v", err)
				}
			} else {
				oc.clearContradicted(sourceEvent)
				logger.Warnf("Skipping the event for now, could not verify it against the source chain: %v", err)
			}
			oc.updateEventProcessedTime(sourceEvent)
			return
		}
		oc.clearContradicted(sourceEvent)

		err := oc.callTargetContract(targetChainID, targetEventType, sourceEvent)

		// Record the attempt whether or not it succeeded. A relay that cannot
		// currently succeed -- e.g. one whose dry run reverts because the vault
		// is short of collateral -- must back off to the retry threshold rather
		// than be re-attempted on every tick.
		oc.updateEventProcessedTime(sourceEvent)

		if err != nil {
			logger.Warnf("Failed to call target contract: %v", err)
		}
	}
}

// verifyEventAgainstSourceChain re-derives the event from the emitting TokenBank's
// storage on the source chain. Returns an ErrEventNotCorroborated error if the
// contract state contradicts the event, or an ErrEventVerificationUnavailable
// error if the check could not be carried out, in which case the caller must not
// relay the event either -- unlike the witness, the orchestrator fails closed,
// since skipping a round merely delays a genuine transfer.
// quarantineExpired reports whether an event has been contradicted continuously for
// longer than the quarantine period, and records the first contradiction if this is
// the first one seen. A transient disagreement -- a backend a few blocks behind, a
// node restarting -- resolves well inside the window; a forged event never does.
func (oc *Orchestrator) quarantineExpired(event *score.InterChainMessageEvent) bool {
	if oc.firstContradictedTime == nil {
		// Writing to a nil map panics, and this path is only reached once a
		// contradiction has already occurred -- i.e. exactly when the node is under
		// attack or degraded. Tolerate a construction path that missed the map
		// rather than taking the validator down at that moment.
		oc.firstContradictedTime = make(map[string]time.Time)
	}

	key := event.ID()
	firstSeen, ok := oc.firstContradictedTime[key]
	if !ok {
		oc.firstContradictedTime[key] = time.Now()
		return false
	}
	return time.Since(firstSeen) > oc.getQuarantinePeriod()
}

// clearContradicted forgets a past contradiction, so that the quarantine measures a
// continuous disagreement rather than an accumulation of unrelated blips.
func (oc *Orchestrator) clearContradicted(event *score.InterChainMessageEvent) {
	delete(oc.firstContradictedTime, event.ID())
}

func (oc *Orchestrator) getQuarantinePeriod() time.Duration {
	seconds := viper.GetInt(scom.CfgSubchainUncorroboratedEventQuarantineInSeconds)
	if seconds <= 0 {
		return defaultUncorroboratedEventQuarantine
	}
	return time.Duration(seconds) * time.Second
}

func (oc *Orchestrator) verifyEventAgainstSourceChain(sourceChainID *big.Int, event *score.InterChainMessageEvent) error {
	tokenBank := oc.getTokenBankForEventType(sourceChainID, event.Type)
	return siu.VerifyEventAgainstTokenBankState(tokenBank, event, oc.getChainHeadHeight(sourceChainID))
}

// getTokenBankForEventType returns the TokenBank on the given chain that emits
// the given event type, or nil if there is no such accessor.
func (oc *Orchestrator) getTokenBankForEventType(chainID *big.Int, eventType score.InterChainMessageEventType) siu.TokenBankEventHeightReader {
	switch eventType {
	case score.IMCEventTypeCrossChainTokenLockTFuel, score.IMCEventTypeCrossChainVoucherBurnTFuel,
		score.IMCEventTypeCrossChainVoucherMintTFuel, score.IMCEventTypeCrossChainTokenUnlockTFuel:
		return oc.getTFuelTokenBank(chainID)
	case score.IMCEventTypeCrossChainTokenLockTNT20, score.IMCEventTypeCrossChainVoucherBurnTNT20,
		score.IMCEventTypeCrossChainVoucherMintTNT20, score.IMCEventTypeCrossChainTokenUnlockTNT20:
		return oc.getTNT20TokenBank(chainID)
	case score.IMCEventTypeCrossChainTokenLockTNT721, score.IMCEventTypeCrossChainVoucherBurnTNT721,
		score.IMCEventTypeCrossChainVoucherMintTNT721, score.IMCEventTypeCrossChainTokenUnlockTNT721:
		return oc.getTNT721TokenBank(chainID)
	case score.IMCEventTypeCrossChainTokenLockTNT1155, score.IMCEventTypeCrossChainVoucherBurnTNT1155,
		score.IMCEventTypeCrossChainVoucherMintTNT1155, score.IMCEventTypeCrossChainTokenUnlockTNT1155:
		return oc.getTNT1155TokenBank(chainID)
	}
	return nil
}

// relayPathEnabled reports whether this node should relay the given asset class in
// the given direction.
//
// CfgSubchainRelayEnabled is all-or-nothing, which is too blunt when one asset's
// pipeline has to stay closed while the others keep running. The processing tick is
// strictly sequential -- TFuel, then TNT20, then TNT721, then TNT1155, in both
// directions -- so an asset that can never make progress otherwise consumes the loop
// ahead of the healthy ones on every retry. It also lets an operator close a
// direction that must not accept new traffic, which is the only lever available when
// the TokenBank contracts cannot be changed.
//
// A path that is not configured is enabled, so existing configs keep working.
func (oc *Orchestrator) relayPathEnabled(sourceChainID *big.Int, eventType score.InterChainMessageEventType) bool {
	asset := relayAssetName(eventType)
	if asset == "" {
		return true // unknown event type, leave it to the rest of the pipeline
	}

	key := fmt.Sprintf("%v.%v.%v", scom.CfgSubchainRelayPathPrefix, asset, oc.relayDirectionName(sourceChainID))
	if !viper.IsSet(key) {
		return true
	}
	if enabled := viper.GetBool(key); !enabled {
		logger.Debugf("Relaying is suspended for %v, skipping", key)
		return false
	}
	return true
}

// relayDirectionName names the direction of travel from the perspective of the
// subchain: "inbound" is main chain -> subchain, "outbound" is subchain -> main chain.
func (oc *Orchestrator) relayDirectionName(sourceChainID *big.Int) string {
	if sourceChainID.Cmp(oc.mainchainID) == 0 {
		return "inbound"
	}
	return "outbound"
}

func relayAssetName(eventType score.InterChainMessageEventType) string {
	switch eventType {
	case score.IMCEventTypeCrossChainTokenLockTFuel, score.IMCEventTypeCrossChainVoucherBurnTFuel,
		score.IMCEventTypeCrossChainVoucherMintTFuel, score.IMCEventTypeCrossChainTokenUnlockTFuel:
		return "tfuel"
	case score.IMCEventTypeCrossChainTokenLockTNT20, score.IMCEventTypeCrossChainVoucherBurnTNT20,
		score.IMCEventTypeCrossChainVoucherMintTNT20, score.IMCEventTypeCrossChainTokenUnlockTNT20:
		return "tnt20"
	case score.IMCEventTypeCrossChainTokenLockTNT721, score.IMCEventTypeCrossChainVoucherBurnTNT721,
		score.IMCEventTypeCrossChainVoucherMintTNT721, score.IMCEventTypeCrossChainTokenUnlockTNT721:
		return "tnt721"
	case score.IMCEventTypeCrossChainTokenLockTNT1155, score.IMCEventTypeCrossChainVoucherBurnTNT1155,
		score.IMCEventTypeCrossChainVoucherMintTNT1155, score.IMCEventTypeCrossChainTokenUnlockTNT1155:
		return "tnt1155"
	}
	return ""
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

		// Diagnostic only: the results are used solely for the two Debugf calls
		// below. Skip the round trips entirely unless debug logging is on, and bound
		// them when it is -- these run on every relay attempt, so an unresponsive
		// node would otherwise stall the tick for a log line nobody reads.
		if logger.Logger.IsLevelEnabled(log.DebugLevel) {
			opts, cancel := boundedCallOpts()
			defer cancel()

			vsQueriedFromMC, mcErr := oc.mainchainTFuelTokenBank.GetAdjustedValidatorSet(opts, oc.subchainID, dynasty)
			vsQueriedFromSC, scErr := oc.subchainTNT20TokenBank.GetAdjustedValidatorSet(opts, oc.subchainID, dynasty)
			logger.Debugf("Subchain %v adjusted ValSet queried from the Mainchain for dynasty %v: %v (err: %v)", oc.subchainID, dynasty, vsQueriedFromMC, mcErr)
			logger.Debugf("Subchain %v adjusted ValSet queried from the Subchain  for dynasty %v: %v (err: %v)", oc.subchainID, dynasty, vsQueriedFromSC, scErr)
		}
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
		logger.Debugf("Sending tokens between two subchains is not supported yet: %v", err)
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
		logger.Debugf("mintTFuelVouchers, skip minting since target chain ID is neither the current chain nor the mainchain: %v vs %v", targetChainID, se.TargetChainID)
		return common.Hash{}, ErrTargetChainMismatch
	}
	dynasty := oc.getDynasty()
	if dynasty == nil {
		return common.Hash{}, ErrDynastyIsNil
	}
	if !oc.checkChainIDCompatability(se.Denom) {
		return common.Hash{}, fmt.Errorf("incompatiable chainID (subchainID: %v, denom: %v)", oc.subchainID, se.Denom)
	}
	tfuelTokenBank := oc.getTFuelTokenBank(targetChainID)
	tx, err := tfuelTokenBank.MintVouchers(txOpts, se.Denom, se.TargetChainVoucherReceiver, se.LockedAmount, dynasty, se.TokenLockNonce)
	if err != nil {
		return common.Hash{}, err
	}
	txHash, err := oc.simulateAndSend(targetChainID, tx)
	if err != nil {
		return common.Hash{}, err
	}
	logger.Debugf("mintTFuelVouchers, dynasty: %v, targetChainID: %v, denom: %v, tokenLockNonce: %v, tx: %v", dynasty, targetChainID, se.Denom, se.TokenLockNonce, txHash.Hex())
	return txHash, nil
}

func (oc *Orchestrator) mintTNT20Vouchers(txOpts *bind.TransactOpts, targetChainID *big.Int, sourceEvent *score.InterChainMessageEvent) (common.Hash, error) {
	se, err := score.ParseToCrossChainTNT20TokenLockedEvent(sourceEvent)
	if err != nil {
		return common.Hash{}, err
	}
	if targetChainID.Cmp(se.TargetChainID) != 0 {
		logger.Debugf("mintTNT20Vouchers, skip minting since target chain ID is neither the current chain nor the mainchain: %v vs %v", targetChainID, se.TargetChainID)
		return common.Hash{}, ErrTargetChainMismatch
	}
	dynasty := oc.getDynasty()
	if dynasty == nil {
		return common.Hash{}, ErrDynastyIsNil
	}
	if !oc.checkChainIDCompatability(se.Denom) {
		return common.Hash{}, fmt.Errorf("incompatiable chainID (subchainID: %v, denom: %v)", oc.subchainID, se.Denom)
	}
	TNT20TokenBank := oc.getTNT20TokenBank(targetChainID)
	tx, err := TNT20TokenBank.MintVouchers(txOpts, se.Denom, se.Name, se.Symbol, se.Decimals, se.TargetChainVoucherReceiver, se.LockedAmount, dynasty, se.TokenLockNonce)
	if err != nil {
		return common.Hash{}, err
	}
	txHash, err := oc.simulateAndSend(targetChainID, tx)
	if err != nil {
		return common.Hash{}, err
	}
	logger.Debugf("mintTNT20Vouchers, dynasty: %v, targetChainID: %v, denom: %v, tokenLockNonce: %v, tx: %v", dynasty, targetChainID, se.Denom, se.TokenLockNonce, txHash.Hex())
	return txHash, nil
}

func (oc *Orchestrator) mintTN721Vouchers(txOpts *bind.TransactOpts, targetChainID *big.Int, sourceEvent *score.InterChainMessageEvent) (common.Hash, error) {
	se, err := score.ParseToCrossChainTNT721TokenLockedEvent(sourceEvent)
	if err != nil {
		return common.Hash{}, err
	}
	if targetChainID.Cmp(se.TargetChainID) != 0 {
		logger.Debugf("mintTN721Vouchers, skip minting since target chain ID is neither the current chain nor the mainchain: %v vs %v", targetChainID, se.TargetChainID)
		return common.Hash{}, ErrTargetChainMismatch
	}
	dynasty := oc.getDynasty()
	if dynasty == nil {
		return common.Hash{}, ErrDynastyIsNil
	}
	if !oc.checkChainIDCompatability(se.Denom) {
		return common.Hash{}, fmt.Errorf("incompatiable chainID (subchainID: %v, denom: %v)", oc.subchainID, se.Denom)
	}
	TNT721TokenBank := oc.getTNT721TokenBank(targetChainID)
	tx, err := TNT721TokenBank.MintVouchers(txOpts, se.Denom, se.Name, se.Symbol, se.TargetChainVoucherReceiver, se.TokenID, se.TokenURI, dynasty, se.TokenLockNonce)
	if err != nil {
		return common.Hash{}, err
	}
	txHash, err := oc.simulateAndSend(targetChainID, tx)
	if err != nil {
		return common.Hash{}, err
	}
	logger.Debugf("mintTN721Vouchers, dynasty: %v, targetChainID: %v, denom: %v, tokenLockNonce: %v, tx: %v", dynasty, targetChainID, se.Denom, se.TokenLockNonce, txHash.Hex())
	return txHash, nil
}

func (oc *Orchestrator) mintTN1155Vouchers(txOpts *bind.TransactOpts, targetChainID *big.Int, sourceEvent *score.InterChainMessageEvent) (common.Hash, error) {
	se, err := score.ParseToCrossChainTNT1155TokenLockedEvent(sourceEvent)
	if err != nil {
		return common.Hash{}, err
	}
	if targetChainID.Cmp(se.TargetChainID) != 0 {
		logger.Debugf("mintTN1155Vouchers, skip minting since target chain ID is neither the current chain nor the mainchain: %v vs %v", targetChainID, se.TargetChainID)
		return common.Hash{}, ErrTargetChainMismatch
	}
	dynasty := oc.getDynasty()
	if dynasty == nil {
		return common.Hash{}, ErrDynastyIsNil
	}
	if !oc.checkChainIDCompatability(se.Denom) {
		return common.Hash{}, fmt.Errorf("incompatiable chainID (subchainID: %v, denom: %v)", oc.subchainID, se.Denom)
	}
	TNT1155TokenBank := oc.getTNT1155TokenBank(targetChainID)
	tx, err := TNT1155TokenBank.MintVouchers(txOpts, se.Denom, se.TargetChainVoucherReceiver, se.TokenID, se.LockedAmount, se.TokenURI, dynasty, se.TokenLockNonce)
	if err != nil {
		return common.Hash{}, err
	}
	txHash, err := oc.simulateAndSend(targetChainID, tx)
	if err != nil {
		return common.Hash{}, err
	}
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
	if !oc.checkChainIDCompatability(se.Denom) {
		return common.Hash{}, fmt.Errorf("incompatiable chainID (subchainID: %v, denom: %v)", oc.subchainID, se.Denom)
	}
	tfuelTokenBank := oc.getTFuelTokenBank(targetChainID)
	tx, err := tfuelTokenBank.UnlockTokens(txOpts, sourceEvent.SourceChainID, se.TargetChainTokenReceiver, se.BurnedAmount, dynasty, se.VoucherBurnNonce)
	if err != nil {
		return common.Hash{}, err
	}
	txHash, err := oc.simulateAndSend(targetChainID, tx)
	if err != nil {
		return common.Hash{}, err
	}
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
	if !oc.checkChainIDCompatability(se.Denom) {
		return common.Hash{}, fmt.Errorf("incompatiable chainID (subchainID: %v, denom: %v)", oc.subchainID, se.Denom)
	}
	if err := oc.verifyTNT20UnlockCollateral(targetChainID, se.Denom, se.BurnedAmount); err != nil {
		return common.Hash{}, err
	}
	TNT20TokenBank := oc.getTNT20TokenBank(targetChainID)
	tx, err := TNT20TokenBank.UnlockTokens(txOpts, sourceEvent.SourceChainID, se.Denom, se.TargetChainTokenReceiver, se.BurnedAmount, dynasty, se.VoucherBurnNonce)
	if err != nil {
		return common.Hash{}, err
	}
	txHash, err := oc.simulateAndSend(targetChainID, tx)
	if err != nil {
		return common.Hash{}, err
	}
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
	if !oc.checkChainIDCompatability(se.Denom) {
		return common.Hash{}, fmt.Errorf("incompatiable chainID (subchainID: %v, denom: %v)", oc.subchainID, se.Denom)
	}
	if err := oc.verifyTNT721UnlockCollateral(targetChainID, se.Denom, se.TokenID); err != nil {
		return common.Hash{}, err
	}
	TNT721TokenBank := oc.getTNT721TokenBank(targetChainID)
	tx, err := TNT721TokenBank.UnlockTokens(txOpts, sourceEvent.SourceChainID, se.Denom, se.TargetChainTokenReceiver, se.TokenID, dynasty, se.VoucherBurnNonce)
	if err != nil {
		return common.Hash{}, err
	}
	txHash, err := oc.simulateAndSend(targetChainID, tx)
	if err != nil {
		return common.Hash{}, err
	}
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
	if !oc.checkChainIDCompatability(se.Denom) {
		return common.Hash{}, fmt.Errorf("incompatiable chainID (subchainID: %v, denom: %v)", oc.subchainID, se.Denom)
	}
	if err := oc.verifyTNT1155UnlockCollateral(targetChainID, se.Denom, se.TokenID, se.BurnedAmount); err != nil {
		return common.Hash{}, err
	}
	TNT1155TokenBank := oc.getTNT1155TokenBank(targetChainID)
	tx, err := TNT1155TokenBank.UnlockTokens(txOpts, sourceEvent.SourceChainID, se.Denom, se.TargetChainTokenReceiver, se.TokenID, se.BurnedAmount, dynasty, se.VoucherBurnNonce)
	if err != nil {
		return common.Hash{}, err
	}
	txHash, err := oc.simulateAndSend(targetChainID, tx)
	if err != nil {
		return common.Hash{}, err
	}
	logger.Debugf("unlockTNT1155Tokens, dynasty: %v, targetChainID: %v, denom: %v, tokenLockNonce: %v, tx: %v", dynasty, targetChainID, se.Denom, se.VoucherBurnNonce, txHash.Hex())
	return txHash, nil
}

func (oc *Orchestrator) checkChainIDCompatability(denom string) bool {
	originalChainID, err := score.ExtractOriginatedChainIDFromDenom(denom)
	if err != nil {
		return false
	}

	compatible := (oc.mainchainID.Cmp(originalChainID) == 0) || (oc.subchainID.Cmp(originalChainID) == 0)

	logger.Debugf("denom: %v, originalChainID: %v, mainchainID: %v, suchainID: %v, compatible: %v",
		denom, originalChainID, oc.mainchainID, oc.subchainID, compatible)
	return compatible
}

func (oc *Orchestrator) buildTxOpts(chainID *big.Int, ecClient *ec.Client) (*bind.TransactOpts, error) {
	var gasPrice *big.Int
	var err error
	if chainID.Cmp(oc.mainchainID) == 0 {
		gasCtx, cancelGas := context.WithTimeout(context.Background(), relayRpcTimeout())
		defer cancelGas()
		gasPrice, err = ecClient.SuggestGasPrice(gasCtx)
		if err != nil {
			return nil, err
		}
	} else {
		// eth_gasPrice returns a hardcoded nubmer for the mainchain, which could be much higher than min gasPrice required by the subchain
		// TODO: parameterize the subchain ETH RPC service to suggest the proper gasPrice for different chains
		// gasPrice = big.NewInt(int64(scom.MinimumGasPrice) * 2)
		gasPrice = common.Big0
	}

	nonceCtx, cancelNonce := context.WithTimeout(context.Background(), relayRpcTimeout())
	defer cancelNonce()
	nonce, err := ecClient.PendingNonceAt(nonceCtx, oc.privateKey.PublicKey().Address())
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
	txOpts.NoSend = true // the tx is dry-run through eth_call by simulateAndSend() before it is broadcast
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

// simulateAndSend dry-runs a relay transaction with eth_call before broadcasting it.
//
// Validator votes are ordinary transactions, so a vote that cannot succeed --
// because the bank is short of the requested amount, because this validator
// already voted, or because the dynasty has aged out -- still costs gas and still
// lands on the target chain. The orchestrator also retries such a transaction
// indefinitely, since the event nonce it is trying to advance never moves. Dry
// running first turns that broadcast loop into a cheap local query and surfaces
// the revert reason in the logs.
func (oc *Orchestrator) simulateAndSend(chainID *big.Int, tx *types.Transaction) (common.Hash, error) {
	client := oc.getEthRpcClient(chainID)
	if client == nil {
		return common.Hash{}, fmt.Errorf("no ETH RPC client for chain %v", chainID)
	}

	callMsg := ethereum.CallMsg{
		From:     oc.privateKey.PublicKey().Address(),
		To:       tx.To(),
		Gas:      tx.Gas(),
		GasPrice: tx.GasPrice(),
		Value:    tx.Value(),
		Data:     tx.Data(),
	}
	// Bound the dry run. Theta's eth_call retries internally on error, sleeping one
	// block between attempts, so a reverting dry run takes the better part of a minute
	// to come back. The processing tick is sequential across all four asset classes in
	// both directions, so an unbounded call here starves every relay behind it. That is
	// not hypothetical: a voucher burn the target bank can never cover stays in the
	// failing state permanently.
	timeout := relayRpcTimeout()

	callCtx, cancelCall := context.WithTimeout(context.Background(), timeout)
	defer cancelCall()
	if _, err := client.CallContract(callCtx, callMsg, nil); err != nil {
		if callCtx.Err() != nil {
			// The dry run did not finish, so we do not know whether the transaction
			// would succeed. Fail closed: skipping a round only delays a genuine
			// transfer, whereas broadcasting blind can burn gas indefinitely on a
			// transaction that cannot succeed.
			return common.Hash{}, fmt.Errorf("%w on chain %v after %v", ErrDryRunTimedOut, chainID, timeout)
		}
		return common.Hash{}, fmt.Errorf("%w on chain %v: %v", ErrTxWouldRevert, chainID, err)
	}

	sendCtx, cancelSend := context.WithTimeout(context.Background(), timeout)
	defer cancelSend()
	if err := client.SendTransaction(sendCtx, tx); err != nil {
		return common.Hash{}, err
	}

	return tx.Hash(), nil
}

// Collateral guards.
//
// TFuelTokenBank enforces "unlockAmount <= totalLockedAmounts" on-chain, and that
// single line is what bounded the loss during the 2026-08-09 incident: the drain
// stopped when the vault ran dry instead of continuing. The TNT banks omit the
// equivalent check deliberately, because elastic-supply tokens (e.g. AMPL) would
// violate the conservation rule and stall the pipeline. The side effect is that a
// forged voucher burn against a TNT bank drains it with nothing to stop it.
//
// The dry run does not cover this: TNT unlockTokens() swallows a failed transfer in
// a try/catch, so the nonce still advances and eth_call still reports success. The
// bank's holdings therefore have to be read explicitly.
//
// These checks are fail-closed, and will stall the nonce sequence for a token whose
// supply genuinely shrinks out from under the bank. Set
// subchain.enforceUnlockCollateral to false on a chain that bridges such a token.

// resolveCollateralTarget returns the token contract carried by the denom together
// with an RPC client for the target chain, or ok=false if the check is switched off.
func (oc *Orchestrator) resolveCollateralTarget(targetChainID *big.Int, denom string) (tokenAddr common.Address, client *ec.Client, ok bool, err error) {
	if !viper.GetBool(scom.CfgSubchainEnforceUnlockCollateral) {
		return common.Address{}, nil, false, nil
	}

	tokenAddr, err = score.ExtractContractAddressFromDenom(denom)
	if err != nil {
		return common.Address{}, nil, false, fmt.Errorf("%w: %v", ErrUnlockUncollateralized, err)
	}

	client = oc.getEthRpcClient(targetChainID)
	if client == nil {
		return common.Address{}, nil, false, fmt.Errorf("no ETH RPC client for chain %v", targetChainID)
	}
	return tokenAddr, client, true, nil
}

// verifyTNT20UnlockCollateral refuses to vote on a TNT20 unlock that the target
// chain TokenBank cannot cover.
func (oc *Orchestrator) verifyTNT20UnlockCollateral(targetChainID *big.Int, denom string, unlockAmount *big.Int) error {
	tokenAddr, client, ok, err := oc.resolveCollateralTarget(targetChainID, denom)
	if err != nil || !ok {
		return err
	}
	if unlockAmount == nil {
		return fmt.Errorf("%w: the event carries no unlock amount", ErrUnlockUncollateralized)
	}

	token, err := scta.NewTNT20VoucherContract(tokenAddr, client)
	if err != nil {
		return fmt.Errorf("failed to bind the TNT20 token %v on chain %v: %v", tokenAddr.Hex(), targetChainID, err)
	}

	bankAddr := oc.getTNT20TokenBankAddr(targetChainID)
	opts, cancel := boundedCallOpts()
	defer cancel()
	balance, err := token.BalanceOf(opts, bankAddr)
	if err != nil {
		// Could not check. Fail closed, consistent with the rest of the vote path:
		// skipping a round merely delays a genuine transfer.
		return fmt.Errorf("failed to read the TNT20 balance of the TokenBank %v on chain %v: %v", bankAddr.Hex(), targetChainID, err)
	}

	if balance.Cmp(unlockAmount) < 0 {
		return fmt.Errorf("%w: the TNT20TokenBank %v on chain %v holds %v of token %v, but the unlock asks for %v",
			ErrUnlockUncollateralized, bankAddr.Hex(), targetChainID, balance, tokenAddr.Hex(), unlockAmount)
	}
	return nil
}

// verifyTNT721UnlockCollateral refuses to vote on an NFT unlock unless the target
// chain TokenBank actually holds that token. Unlike the fungible case there is no
// amount to compare: either the bank is the current owner or the unlock is baseless.
func (oc *Orchestrator) verifyTNT721UnlockCollateral(targetChainID *big.Int, denom string, tokenID *big.Int) error {
	tokenAddr, client, ok, err := oc.resolveCollateralTarget(targetChainID, denom)
	if err != nil || !ok {
		return err
	}
	if tokenID == nil {
		return fmt.Errorf("%w: the event carries no token ID", ErrUnlockUncollateralized)
	}

	token, err := scta.NewTNT721VoucherContract(tokenAddr, client)
	if err != nil {
		return fmt.Errorf("failed to bind the TNT721 token %v on chain %v: %v", tokenAddr.Hex(), targetChainID, err)
	}

	bankAddr := oc.getTNT721TokenBankAddr(targetChainID)
	opts, cancel := boundedCallOpts()
	defer cancel()
	owner, err := token.OwnerOf(opts, tokenID)
	if err != nil {
		return fmt.Errorf("failed to read the owner of TNT721 token %v (%v) on chain %v: %v", tokenID, tokenAddr.Hex(), targetChainID, err)
	}

	if owner != bankAddr {
		return fmt.Errorf("%w: TNT721 token %v of %v on chain %v is held by %v, not by the TokenBank %v",
			ErrUnlockUncollateralized, tokenID, tokenAddr.Hex(), targetChainID, owner.Hex(), bankAddr.Hex())
	}
	return nil
}

// verifyTNT1155UnlockCollateral refuses to vote on a TNT1155 unlock the target chain
// TokenBank cannot cover for that specific token ID.
func (oc *Orchestrator) verifyTNT1155UnlockCollateral(targetChainID *big.Int, denom string, tokenID *big.Int, unlockAmount *big.Int) error {
	tokenAddr, client, ok, err := oc.resolveCollateralTarget(targetChainID, denom)
	if err != nil || !ok {
		return err
	}
	if tokenID == nil || unlockAmount == nil {
		return fmt.Errorf("%w: the event carries no token ID or no unlock amount", ErrUnlockUncollateralized)
	}

	token, err := scta.NewTNT1155VoucherContract(tokenAddr, client)
	if err != nil {
		return fmt.Errorf("failed to bind the TNT1155 token %v on chain %v: %v", tokenAddr.Hex(), targetChainID, err)
	}

	bankAddr := oc.getTNT1155TokenBankAddr(targetChainID)
	opts, cancel := boundedCallOpts()
	defer cancel()
	balance, err := token.BalanceOf(opts, bankAddr, tokenID)
	if err != nil {
		return fmt.Errorf("failed to read the TNT1155 balance of the TokenBank %v on chain %v: %v", bankAddr.Hex(), targetChainID, err)
	}

	if balance.Cmp(unlockAmount) < 0 {
		return fmt.Errorf("%w: the TNT1155TokenBank %v on chain %v holds %v of token %v id %v, but the unlock asks for %v",
			ErrUnlockUncollateralized, bankAddr.Hex(), targetChainID, balance, tokenAddr.Hex(), tokenID, unlockAmount)
	}
	return nil
}

func (oc *Orchestrator) getTNT20TokenBankAddr(chainID *big.Int) common.Address {
	if chainID.Cmp(oc.mainchainID) == 0 {
		return oc.mainchainTNT20TokenBankAddr
	}
	return oc.subchainTNT20TokenBankAddr
}

func (oc *Orchestrator) getTNT721TokenBankAddr(chainID *big.Int) common.Address {
	if chainID.Cmp(oc.mainchainID) == 0 {
		return oc.mainchainTNT721TokenBankAddr
	}
	return oc.subchainTNT721TokenBankAddr
}

func (oc *Orchestrator) getTNT1155TokenBankAddr(chainID *big.Int) common.Address {
	if chainID.Cmp(oc.mainchainID) == 0 {
		return oc.mainchainTNT1155TokenBankAddr
	}
	return oc.subchainTNT1155TokenBankAddr
}

// getChainHeadHeight returns the given chain's head as seen by the node this
// orchestrator queries, or nil if it cannot be read. nil disables the verifier's
// staleness check rather than failing verification outright.
func (oc *Orchestrator) getChainHeadHeight(chainID *big.Int) *big.Int {
	client := oc.getEthRpcClient(chainID)
	if client == nil {
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), relayRpcTimeout())
	defer cancel()

	height, err := client.BlockNumber(ctx)
	if err != nil {
		logger.Warnf("Could not read the head height of chain %v, skipping the staleness check: %v", chainID, err)
		return nil
	}
	return new(big.Int).SetUint64(height)
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
