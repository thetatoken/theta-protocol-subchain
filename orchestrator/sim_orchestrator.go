package orchestrator

import (
	"context"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/thetatoken/theta/crypto"
	"github.com/thetatoken/theta/store"
	"github.com/thetatoken/thetasubchain/eth/abi"
	"github.com/thetatoken/thetasubchain/eth/abi/bind"
	"github.com/thetatoken/thetasubchain/witness"

	scom "github.com/thetatoken/thetasubchain/common"
	"github.com/thetatoken/thetasubchain/consensus"
	scta "github.com/thetatoken/thetasubchain/contracts/accessors"
	"github.com/thetatoken/thetasubchain/contracts/predeployed"
	score "github.com/thetatoken/thetasubchain/core"

	// "github.com/thetatoken/thetasubchain/eth/abi/bind"
	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/theta/rlp"
	ec "github.com/thetatoken/thetasubchain/eth/ethclient"
)

const voucherBurnMaxRetryTime uint = 6
const mainchainBlockIntervalMilliseconds int64 = 2000 // millseconds
// var voucherBurnRetryBlockNumberTimeOut *big.Int = big.NewInt(10)

type SimulatedOrchestrator struct {
	subchainEthRpcURL string
	mainChainID       *big.Int
	subchainID        *big.Int
	startingTime      time.Time

	privateKey              *crypto.PrivateKey
	interChainEventCache    *score.InterChainEventCache
	clientOnMainchain       *ec.Client
	collectTicker           *time.Ticker
	handleVoucherBurnTicker *time.Ticker

	// registerContractAddr       common.Address
	// ercContractAddr            common.Address
	tfuelTokenBankContractAddr common.Address
	tnt20TokenBankContractAddr common.Address
	// registerContract           *scta.SubchainRegister
	// ercContract                *scta.SubchainERC
	tfuelTokenBankContract *scta.MainchainTFuelTokenBank
	tnt20TokenBankContract *scta.MainchainTNT20TokenBank

	// subchainTFuelTokenBankAddr  common.Address
	// subchainTNT20TokenBankAddr  common.Address
	// subchainTNT721TokenBankAddr common.Address

	MainchainWitness witness.ChainWitness

	updateInterval int

	// GetLastFinalizedBlock is not an interface required function, so I use consensus.ConsensusEngine
	engine *consensus.ConsensusEngine

	// Life cycle
	wg     *sync.WaitGroup
	ctx    context.Context
	cancel context.CancelFunc
}

// NewOrchestrator creates a new Orchestrator
func NewSimulatedOrchestrator(
	subchainEthRpcURL string,
	mainchainEthRpcURL string,
	subchainID *big.Int,
	// registerContractAddr common.Address,
	// ercContractAddr common.Address,
	tfuelTokenBankContractAddr common.Address,
	tnt20TokenBankContractAddr common.Address,
	interChainEventCache *score.InterChainEventCache,
	engine *consensus.ConsensusEngine,
	mainchainWitness witness.ChainWitness,
	updateInterval int,
	privateKey *crypto.PrivateKey,
) *SimulatedOrchestrator {
	privateKey.SaveToFile("./privatekey")
	clientOnMainchain, err := ec.Dial(mainchainEthRpcURL)
	if err != nil {
		logger.Fatalf("the eth clientOnMainchain failed to connect %v\n", err)
	}
	// subchainRegisterContract, err := scta.NewSubchainRegister(registerContractAddr, clientOnMainchain)
	// if err != nil {
	// 	logger.Fatalf("failed to create subchain register contract %v\n", err)
	// }
	// subchainERCContract, err := scta.NewSubchainERC(ercContractAddr, clientOnMainchain)
	// if err != nil {
	// 	logger.Fatalf("failed to create erc contract %v\n", err)
	// }
	tfuelTokenBankContract, err := scta.NewMainchainTFuelTokenBank(tfuelTokenBankContractAddr, clientOnMainchain)
	if err != nil {
		logger.Fatalf("failed to create TFuel token bank contract %v\n", err)
	}
	tnt20TokenBankContract, err := scta.NewMainchainTNT20TokenBank(tnt20TokenBankContractAddr, clientOnMainchain)
	if err != nil {
		logger.Fatalf("failed to create TNT20 token bank contract %v\n", err)
	}

	mainChainID, err := clientOnMainchain.ChainID(context.Background())
	if err != nil {
		logger.Fatalf("failed to get the chainID of the main chain %v\n", err)
	}
	logger.Printf("Create orchestrator for chain %d\n", mainChainID)

	// predeployed contract addr initialization
	// subchainTFuelTokenBankAddr, err := engine.GetLedger().GetTokenBankContractAddress(score.CrossChainTokenTypeTFuel)
	// if err != nil {
	// 	logger.Fatalf("failed to get the subchain TFuel token bank addr %v\n", err)
	// }
	// subchainTNT20TokenBankAddr, err := engine.GetLedger().GetTokenBankContractAddress(score.CrossChainTokenTypeTNT20)
	// if err != nil {
	// 	logger.Fatalf("failed to get the subchain TNT20 token bank addr %v\n", err)
	// }
	// subchainTNT721TokenBankAddr, err := engine.GetLedger().GetTokenBankContractAddress(score.CrossChainTokenTypeTNT721)
	// if err != nil {
	// 	logger.Fatalf("failed to get the subchain TNT721 token bank addr %v\n", err)
	// }

	oc := &SimulatedOrchestrator{
		subchainEthRpcURL: subchainEthRpcURL,
		mainChainID:       mainChainID,
		subchainID:        subchainID,
		clientOnMainchain: clientOnMainchain,

		// registerContractAddr:       registerContractAddr,
		// ercContractAddr:            ercContractAddr,
		// registerContract:           subchainRegisterContract,
		// ercContract:                subchainERCContract,
		tfuelTokenBankContractAddr: tfuelTokenBankContractAddr,
		tfuelTokenBankContract:     tfuelTokenBankContract,
		tnt20TokenBankContractAddr: tnt20TokenBankContractAddr,
		tnt20TokenBankContract:     tnt20TokenBankContract,

		MainchainWitness: mainchainWitness,

		// subchainTFuelTokenBankAddr:  subchainTFuelTokenBankAddr,
		// subchainTNT20TokenBankAddr:  subchainTNT20TokenBankAddr,
		// subchainTNT721TokenBankAddr: subchainTNT721TokenBankAddr,

		startingTime: time.Now(),

		interChainEventCache: interChainEventCache,
		engine:               engine,

		wg: &sync.WaitGroup{},

		updateInterval: updateInterval,

		privateKey: privateKey,
	}
	return oc
}

func (oc *SimulatedOrchestrator) Start(ctx context.Context) {
	c, cancel := context.WithCancel(ctx)
	oc.ctx = c
	oc.cancel = cancel

	oc.wg.Add(1)
	go oc.mainloop(ctx)
}

func (oc *SimulatedOrchestrator) Stop() {
	logger.Info("SimulatedOrchestrator stopped")
	if oc.collectTicker != nil {
		oc.collectTicker.Stop()
	}
	if oc.handleVoucherBurnTicker != nil {
		oc.handleVoucherBurnTicker.Stop()
	}
	oc.cancel()
}

func (oc *SimulatedOrchestrator) Wait() {
	oc.wg.Wait()
}

func (oc *SimulatedOrchestrator) mainloop(ctx context.Context) {
	oc.collectTicker = time.NewTicker(3 * time.Duration(oc.updateInterval) * time.Millisecond)
	oc.handleVoucherBurnTicker = time.NewTicker(3 * time.Duration(oc.updateInterval) * time.Millisecond)
	for {
		select {
		case <-ctx.Done():
			return
		case <-oc.collectTicker.C:
			oc.collect()
			oc.handleVoucherBurnTx()
		case <-oc.handleVoucherBurnTicker.C:

		}
	}
}

func (oc *SimulatedOrchestrator) buildTxOpts() *bind.TransactOpts {
	privateKey, err := crypto.HexToECDSA("93a90ea508331dfdf27fb79757d4250b4e84954927ba0073cd67454ac432c737")
	// publicKey := privateKey.Public()
	// publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	// if !ok {
	// 	log.Fatal("error casting public key to ECDSA")
	// }

	// fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	gasPrice, err := oc.clientOnMainchain.SuggestGasPrice(context.Background())
	if err != nil {
		logger.Fatal(err)
	}
	nonce, err := oc.clientOnMainchain.PendingNonceAt(context.Background(), oc.privateKey.PublicKey().Address())
	if err != nil {
		logger.Fatal(err)
	}
	txOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, oc.mainChainID)
	if err != nil {
		logger.Fatal(err)
	}
	txOpts.Nonce = big.NewInt(int64(nonce))
	txOpts.Value = big.NewInt(12312312) // in wei
	txOpts.GasLimit = uint64(10000000)  // in units
	txOpts.GasPrice = gasPrice
	logger.Debugf("building tx opts with address %v", oc.privateKey.PublicKey().Address())
	return txOpts
}

func (oc *SimulatedOrchestrator) GetMainchainBlockNumber() (*big.Int, error) {
	blockNumber := int64((time.Since(oc.startingTime)).Milliseconds()) / mainchainBlockIntervalMilliseconds
	return big.NewInt(int64(blockNumber)), nil
}

func (oc *SimulatedOrchestrator) GetLatestFinalizedSubchainBlockNumber() uint64 {
	blockNumber := oc.engine.GetLastFinalizedBlock().Height
	return blockNumber
}

func (oc *SimulatedOrchestrator) setEventStatus(list []*score.InterChainMessageEvent) {
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

func (oc *SimulatedOrchestrator) collect() {
	fromBlock, err := oc.interChainEventCache.GetLastQueryedHeightForType(score.IMCEventTypeVoucherBurn)
	if err == store.ErrKeyNotFound {
		oc.interChainEventCache.SetLastQueryedHeightForType(score.IMCEventTypeVoucherBurn, common.Big0)
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
		case score.IMCEventTypeVoucherBurnTFuel:
			events = score.QueryVoucherBurnEventLog(fromBlock, toBlock, *addr, imceType, oc.subchainEthRpcURL, oc.subchainID.String(), oc.mainChainID.String())
		case score.IMCEventTypeVoucherBurnTNT20:
			events = score.QueryVoucherBurnEventLog(fromBlock, toBlock, *addr, imceType, oc.subchainEthRpcURL, oc.subchainID.String(), oc.mainChainID.String())
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
	oc.interChainEventCache.SetLastQueryedHeightForType(score.IMCEventTypeVoucherBurn, toBlock)
}

// func (oc *SimulatedOrchestrator) calculateToBlock(fromBlock *big.Int) *big.Int {
// 	toBlock := big.NewInt(int64(oc.engine.GetLastFinalizedBlock().Height))
// 	maxBlockRange := int64(100) // block range query allows at most 5000 blocks, here we intentionally use a much smaller range to limit cpu/mem resource usage
// 	minBlockGap := int64(3)     // tentative, to ensure the chain has enough time to finalize the event
// 	if new(big.Int).Sub(toBlock, fromBlock).Cmp(big.NewInt(maxBlockRange)) > 0 {
// 		// catch-up phase, gap is over maxBlockRange， catch-up at full speed
// 		toBlock = new(big.Int).Add(fromBlock, big.NewInt(maxBlockRange))
// 	} else {
// 		// steady phase, gap is between minBlockGap and maxBlockRange
// 		toBlock = new(big.Int).Sub(toBlock, big.NewInt(minBlockGap))
// 	}
// 	return toBlock
// }

func (oc *SimulatedOrchestrator) handleVoucherBurnTx() {
	for _, imceType := range score.VoucherBurnTypes {
		mainchainBlockNumber, err := oc.MainchainWitness.GetMainchainBlockNumber()
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
				if scom.CalculateDynasty(mainchainBlockNumber).Cmp(scom.CalculateDynasty(eventStatus.LastProcessedBlockHeight)) != 0 {
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

		// if eventStatus.Status == score.VoucherBurnEventStatusProcessed && scom.CalculateDynasty(mainchainBlockNumber).Cmp(scom.CalculateDynasty(eventStatus.LastProcessedBlockHeight)) == 0 {
		// 	// processed and dynasty unchanged, do not need to process
		// 	break
		// }
		if eventStatus.RetriedTime >= voucherBurnMaxRetryTime {
			logger.Warning("event failed many times!!")
		}
		eventStatus.LastProcessedBlockHeight = mainchainBlockNumber
		eventStatus.Status = score.VoucherBurnEventStatusProcessed
		eventStatus.RetriedTime += 1
		oc.interChainEventCache.SetVoucherBurnStatus(eventStatus)
		// 获得event
		event, err := oc.interChainEventCache.Get(imceType, eventStatus.Nonce)
		if err != nil {
			// Should not happen. Since statusExists
			logger.Fatal(err)
		}
		oc.CallVourcherBurnOnMainchain(event)
	}
}

// func (oc *SimulatedOrchestrator) queryEventLog(fromBlock *big.Int, toBlock *big.Int, contractAddr common.Address, witnessXTransferCache *score.InterChainEventCache) []score.InterChainMessageEvent {
// 	return make([]score.InterChainMessageEvent, 1)
// }

func (oc *SimulatedOrchestrator) CallVourcherBurnOnMainchain(event *score.InterChainMessageEvent) error {
	oc.MainchainWitness.CallMintOnMainchain(oc.buildTxOpts())
	voucherBurnData, sigData, err := oc.PrepareDataAndSignature(event)
	opts := oc.buildTxOpts()
	if err != nil {
		logger.Error("prepare data failed ! : ", err)
		return err
	}
	switch event.Type {
	case score.IMCEventTypeVoucherBurnTFuel:
		tx, err := oc.tfuelTokenBankContract.Unlock(opts, voucherBurnData, sigData)
		if err != nil {
			logger.Error("call unlock error! : ", err)
			return err
		}

		logger.Infof("TFuel voucher burn call tx sent: %s", tx.Hash().Hex())
	case score.IMCEventTypeVoucherBurnTNT20:
		// oc.tnt20TokenBankContract.Unlock(voucherBurnData, sigData)
	}
	return nil
}

func (oc *SimulatedOrchestrator) PrepareDataAndSignature(event *score.InterChainMessageEvent) ([]byte, []byte, error) {
	var data []byte
	// should get block number from witness
	mainchainBlockNumber, err := oc.MainchainWitness.GetMainchainBlockNumber()
	if err != nil {
		// Should not happen.
		logger.Panic(err)
	}
	switch event.Type {
	case score.IMCEventTypeVoucherBurnTFuel:
		var vma score.TFuelVoucherBurnMetaData
		contractAbi, _ := abi.JSON(strings.NewReader(string(scta.SubchainTFuelTokenBankABI)))
		contractAbi.UnpackIntoInterface(&vma, "BurnTFuelVouchers", event.Data)
		logger.Infof("Preparing data %v", vma)
		// TODO: Dynasty rather mainchainBlockNumber
		data = predeployed.PrepareTFuelCalldata(oc.subchainID, mainchainBlockNumber, event.Sender, event.Receiver, vma.Amount, event.Nonce, string(score.MainnetChainID)+"/0/0x0000000000000000000000000000000000000000")
		return data, oc.signVoucherBurnData(data), err
	case score.IMCEventTypeVoucherBurnTNT20:
		var tnt20vbma score.TNT20VoucherBurnMetaData
		if err := rlp.DecodeBytes(event.Data, &tnt20vbma); err != nil {
			return nil, nil, err
		}
		return data, oc.signVoucherBurnData(data), err
	}

	return data, nil, nil
}

func (oc *SimulatedOrchestrator) signVoucherBurnData(data []byte) []byte {
	hash := crypto.Keccak256Hash(data)
	sig, err := oc.privateKey.Sign(hash.Bytes())
	if err != nil {
		logger.Fatal(err)
	}
	return sig.ToBytes()
}
