package orchestrator

import (
	"context"
	"math/big"
	"strings"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/thetatoken/theta/crypto"
	"github.com/thetatoken/theta/store"
	"github.com/thetatoken/theta/store/database"
	"github.com/thetatoken/thetasubchain/eth/abi"
	"github.com/thetatoken/thetasubchain/eth/abi/bind"
	siu "github.com/thetatoken/thetasubchain/interchain/utils"
	"github.com/thetatoken/thetasubchain/interchain/witness"

	scom "github.com/thetatoken/thetasubchain/common"
	scta "github.com/thetatoken/thetasubchain/contracts/accessors"
	"github.com/thetatoken/thetasubchain/contracts/predeployed"
	score "github.com/thetatoken/thetasubchain/core"

	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/theta/rlp"
	ec "github.com/thetatoken/thetasubchain/eth/ethclient"
)

var logger *log.Entry = log.WithFields(log.Fields{"prefix": "orchestrator"})

const voucherBurnMaxRetryTime uint = 6
const mainchainBlockIntervalMilliseconds int64 = 2000 // millseconds
// var voucherBurnRetryBlockNumberTimeOut *big.Int = big.NewInt(10)

type Orchestrator struct {
	updateInterval   int
	privateKey       *crypto.PrivateKey
	collectTicker    *time.Ticker
	ocstState        *orchestratorState
	metachainWitness witness.ChainWitness

	// The main chain
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
		logger.Fatalf("failed to get the chainID of the main chain %v\n", err)
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

	ocstState := newOrchestratorState(db)

	oc := &Orchestrator{
		updateInterval:   updateInterval,
		privateKey:       privateKey,
		ocstState:        ocstState,
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
	if oc.collectTicker != nil {
		oc.collectTicker.Stop()
	}
	oc.cancel()
	logger.Info("Metachain orchestrator stopped")
}

func (oc *Orchestrator) Wait() {
	oc.wg.Wait()
}

func (oc *Orchestrator) SetSubchainTokenBanks(ledger score.Ledger) {
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
	oc.collectTicker = time.NewTicker(3 * time.Duration(oc.updateInterval) * time.Millisecond)
	for {
		select {
		case <-ctx.Done():
			return
		case <-oc.collectTicker.C:
			oc.collect()
			oc.handleVoucherBurnTx()
		}
	}
}

func (oc *Orchestrator) buildTxOpts() *bind.TransactOpts {
	gasPrice, err := oc.mainchainEthRpcClient.SuggestGasPrice(context.Background())
	if err != nil {
		logger.Fatal(err)
	}
	nonce, err := oc.mainchainEthRpcClient.PendingNonceAt(context.Background(), oc.privateKey.PublicKey().Address())
	if err != nil {
		logger.Fatal(err)
	}
	txOpts, err := bind.NewKeyedTransactorWithChainID(oc.privateKey, oc.mainchainID)
	if err != nil {
		logger.Fatal(err)
	}
	txOpts.Nonce = big.NewInt(int64(nonce))
	txOpts.Value = big.NewInt(0)       // in wei
	txOpts.GasLimit = uint64(10000000) // in units
	txOpts.GasPrice = gasPrice
	logger.Debugf("building tx opts with address %v", oc.privateKey.PublicKey().Address())
	return txOpts
}

func (oc *Orchestrator) setEventStatus(list []*score.InterChainMessageEvent) {
	for _, event := range list {
		event := event
		eventStatus := &score.VoucherBurnEventStatusInfo{
			Type:                     event.Type,
			Nonce:                    event.Nonce,
			Status:                   score.VoucherBurnEventStatusPending,
			LastProcessedBlockHeight: common.Big0,
			RetriedTime:              0,
		}
		oc.interChainEventCache.SetVoucherBurnStatus(eventStatus)
	}
}

func (oc *Orchestrator) collect() {
	fromBlock, err := oc.interChainEventCache.GetLastQueryedHeightForType(score.IMCEventTypeCrossChainVoucherBurn)
	if err == store.ErrKeyNotFound {
		oc.interChainEventCache.SetLastQueryedHeightForType(score.IMCEventTypeCrossChainVoucherBurn, common.Big0)
	} else if err != nil {
		logger.Warnf("failed to get the last queryed height %v\n", err)
	}
	// toBlock := oc.calculateToBlock(fromBlock)
	toBlock := big.NewInt(int64(oc.engine.GetLastFinalizedBlock().Height))
	logger.Infof("trying query voucher burn on subchain from %v to %v\n", fromBlock.String(), toBlock.String())
	if fromBlock.Cmp(toBlock) == 0 {
		return
	}
	logger.Infof("query voucher burn on subchain from %v to %v\n", fromBlock.String(), toBlock.String())
	for _, imceType := range score.VoucherBurnTypes {
		var events []*score.InterChainMessageEvent
		addr, err := oc.engine.GetLedger().GetTokenBankContractAddress(score.CrossChainTokenTypeTFuel)
		if err != nil {
			logger.Errorf("Error in getting the token bank address")
		}
		switch imceType {
		case score.IMCEventTypeCrossChainVoucherBurnTFuel:
			events = score.QueryVoucherBurnEventLog(fromBlock, toBlock, *addr, imceType, oc.subchainEthRpcURL, oc.subchainID.String(), oc.mainchainID.String())
		case score.IMCEventTypeCrossChainVoucherBurnTNT20:
			events = score.QueryVoucherBurnEventLog(fromBlock, toBlock, *addr, imceType, oc.subchainEthRpcURL, oc.subchainID.String(), oc.mainchainID.String())
		}
		if len(events) == 0 {
			continue
		}
		err = oc.interChainEventCache.InsertList(events)
		oc.setEventStatus(events)
		if err != nil { // should not happen
			logger.Panicf("failed to insert events into cache")
		}
	}
	oc.interChainEventCache.SetLastQueryedHeightForType(score.IMCEventTypeCrossChainVoucherBurn, toBlock)
}

func (oc *Orchestrator) getTokenBankContractAddress(cctt score.CrossChainTokenType, isOnMainchain bool) {

}

func (oc *Orchestrator) handleVoucherBurnTx() {
	for _, imceType := range siu.VoucherBurnTypes {
		mainchainBlockHeight, err := oc.metachainWitness.GetMainchainBlockHeight()
		if err != nil {
			// Should not happen.
			logger.Panic(err)
		}
		continueProcessVoucherBurn := false
		nextEventNonce, err := oc.interChainEventCache.GetNextVoucherBurnNonceForType(imceType)
		if err == store.ErrKeyNotFound {
			oc.interChainEventCache.SetNextVoucherBurnNonceForType(imceType, common.Big1)
		} else if err != nil {
			logger.Errorf("Failed to get the next event type for nonce %v, type %v", err, imceType)
		}
		var eventStatus *score.VoucherBurnEventStatusInfo
		for { // find the next nonce to process
			statusExists, _ := oc.interChainEventCache.VoucherBurnNonceExists(imceType, nextEventNonce)
			if !statusExists {
				break
			}
			eventStatus, err = oc.interChainEventCache.GetVoucherBurnStatus(imceType, nextEventNonce)
			if err != nil {
				// Should not happen. Since statusExists
				logger.Panic(err)
			}
			// check whether this voucher burn is finalized on mainchain
			if eventStatus.Status == score.VoucherBurnEventStatusFinalized {
				oc.interChainEventCache.SetNextVoucherBurnNonceForType(imceType, new(big.Int).Add(nextEventNonce, common.Big1))
				nextEventNonce = new(big.Int).Add(nextEventNonce, common.Big1)
			} else if eventStatus.Status == score.VoucherBurnEventStatusProcessed {
				if scom.CalculateDynasty(mainchainBlockHeight).Cmp(scom.CalculateDynasty(eventStatus.LastProcessedBlockHeight)) != 0 {
					// dynasty changed and this event is not finalized on mainchain, so process it again
					continueProcessVoucherBurn = true
				}
				break
			} else if eventStatus.Status == score.VoucherBurnEventStatusPending {
				continueProcessVoucherBurn = true
				break
			} else {
				break
			}
		}

		if eventStatus == nil || !continueProcessVoucherBurn {
			break
		}

		logger.Infof("Processing voucher burn type %v nonce %v", imceType, nextEventNonce)

		// if eventStatus.Status == score.VoucherBurnEventStatusProcessed && scom.CalculateDynasty(mainchainBlockHeight).Cmp(scom.CalculateDynasty(eventStatus.LastProcessedBlockHeight)) == 0 {
		// 	// processed and dynasty unchanged, do not need to process
		// 	break
		// }
		if eventStatus.RetriedTime >= voucherBurnMaxRetryTime {
			logger.Warning("event failed many times!!")
		}
		eventStatus.LastProcessedBlockHeight = mainchainBlockHeight
		eventStatus.Status = score.VoucherBurnEventStatusProcessed
		eventStatus.RetriedTime += 1
		oc.interChainEventCache.SetVoucherBurnStatus(eventStatus)
		event, err := oc.interChainEventCache.Get(imceType, eventStatus.Nonce)
		if err != nil {
			// Should not happen. Since statusExists
			logger.Fatal(err)
		}
		oc.CallVourcherBurnOnMainchain(event)
	}
}

func (oc *Orchestrator) CallVourcherBurnOnMainchain(event *score.InterChainMessageEvent) error {
	voucherBurnData, sigData, err := oc.PrepareDataAndSignature(event)
	opts := oc.buildTxOpts()
	if err != nil {
		logger.Error("prepare data failed ! : ", err)
		return err
	}
	switch event.Type {
	case score.IMCEventTypeCrossChainVoucherBurnTFuel:
		tx, err := oc.mainchainTFuelTokenBank.Unlocktokens(opts, voucherBurnData, sigData)
		if err != nil {
			logger.Error("call unlock error! : ", err)
			return err
		}

		logger.Infof("TFuel voucher burn call tx sent: %s", tx.Hash().Hex())
	case score.IMCEventTypeCrossChainVoucherBurnTNT20:
		// oc.mainchainTNT20TokenBank.Unlock(voucherBurnData, sigData)
	}
	return nil
}

func (oc *Orchestrator) PrepareDataAndSignature(event *score.InterChainMessageEvent) ([]byte, []byte, error) {
	var data []byte
	// should get block number from witness
	mainchainBlockHeight, err := oc.metachainWitness.GetMainchainBlockHeight()
	if err != nil {
		// Should not happen.
		logger.Panic(err)
	}
	switch event.Type {
	case score.IMCEventTypeCrossChainVoucherBurnTFuel:
		var vma score.CrossChainTFuelVoucherBurnedEvent
		contractAbi, _ := abi.JSON(strings.NewReader(string(scta.TFuelTokenBankABI)))
		contractAbi.UnpackIntoInterface(&vma, "TFuelVoucherBurned", event.Data)
		logger.Infof("Preparing data %v", vma)
		// TODO: Dynasty rather mainchainBlockHeight
		data = predeployed.PrepareTFuelCalldata(oc.subchainID, mainchainBlockHeight, event.Sender, event.Receiver, vma.Amount, event.Nonce, string(score.MainnetChainID)+"/0/0x0000000000000000000000000000000000000000")
		return data, oc.signVoucherBurnData(data), err
	case score.IMCEventTypeCrossChainVoucherBurnTNT20:
		var tnt20vbma score.CrossChainTNT20VoucherBurnedEvent
		if err := rlp.DecodeBytes(event.Data, &tnt20vbma); err != nil {
			return nil, nil, err
		}
		return data, oc.signVoucherBurnData(data), err
	}

	return data, nil, nil
}

func (oc *Orchestrator) signVoucherBurnData(data []byte) []byte {
	hash := crypto.Keccak256Hash(data)
	sig, err := oc.privateKey.Sign(hash.Bytes())
	if err != nil {
		logger.Fatal(err)
	}
	return sig.ToBytes()
}
