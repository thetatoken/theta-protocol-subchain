package orchestrator

import (
	"context"
	"math/big"
	"sync"
	"time"

	"github.com/thetatoken/theta/crypto"
	"github.com/thetatoken/thetasubchain/eth/abi/bind"

	scom "github.com/thetatoken/thetasubchain/common"
	"github.com/thetatoken/thetasubchain/consensus"
	scta "github.com/thetatoken/thetasubchain/contracts/accessors"
	score "github.com/thetatoken/thetasubchain/core"

	// "github.com/thetatoken/thetasubchain/eth/abi/bind"
	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/theta/rlp"
	ec "github.com/thetatoken/thetasubchain/eth/ethclient"
)

const voucherBurnMaxRetryTime uint = 6
const mainchainBlockIntervalMilliseconds int64 = 2000 // millseconds
var voucherBurnRetryBlockNumberTimeOut *big.Int = big.NewInt(1000)

type SimulatedOrchestrator struct {
	mainChainID  *big.Int
	subchainID   *big.Int
	startingTime time.Time

	privateKey           *crypto.PrivateKey
	interChainEventCache *score.InterChainEventCache
	client               *ec.Client
	collectTicker        *time.Ticker

	registerContractAddr       common.Address
	ercContractAddr            common.Address
	tfuelTokenBankContractAddr common.Address
	tnt20TokenBankContractAddr common.Address
	registerContract           *scta.SubchainRegister
	ercContract                *scta.SubchainERC
	tfuelTokenBankContract     *scta.MainchainTFuelTokenBank
	tnt20TokenBankContract     *scta.MainchainTNT20TokenBank

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
	ethClientAddress string,
	subchainID *big.Int,
	registerContractAddr common.Address,
	ercContractAddr common.Address,
	tfuelTokenBankContractAddr common.Address,
	tnt20TokenBankContractAddr common.Address,
	interChainEventCache *score.InterChainEventCache,
	engine *consensus.ConsensusEngine,
) *SimulatedOrchestrator {
	client, err := ec.Dial(ethClientAddress)
	if err != nil {
		logger.Fatalf("the eth client failed to connect %v\n", err)
	}
	subchainRegisterContract, err := scta.NewSubchainRegister(registerContractAddr, client)
	if err != nil {
		logger.Fatalf("failed to create subchain register contract %v\n", err)
	}
	subchainERCContract, err := scta.NewSubchainERC(ercContractAddr, client)
	if err != nil {
		logger.Fatalf("failed to create erc contract %v\n", err)
	}
	tfuelTokenBankContract, err := scta.NewMainchainTFuelTokenBank(tfuelTokenBankContractAddr, client)
	if err != nil {
		logger.Fatalf("failed to create TFuel token bank contract %v\n", err)
	}
	tnt20TokenBankContract, err := scta.NewMainchainTNT20TokenBank(tnt20TokenBankContractAddr, client)
	if err != nil {
		logger.Fatalf("failed to create TNT20 token bank contract %v\n", err)
	}

	mainChainID, err := client.ChainID(context.Background())
	if err != nil {
		logger.Fatalf("failed to get the chainID of the main chain %v\n", err)
	}
	logger.Printf("Create transfer validator for chain %d\n", mainChainID)

	oc := &SimulatedOrchestrator{
		mainChainID: mainChainID,
		subchainID:  subchainID,
		client:      client,

		registerContractAddr:       registerContractAddr,
		ercContractAddr:            ercContractAddr,
		registerContract:           subchainRegisterContract,
		ercContract:                subchainERCContract,
		tfuelTokenBankContractAddr: tfuelTokenBankContractAddr,
		tfuelTokenBankContract:     tfuelTokenBankContract,
		tnt20TokenBankContractAddr: tnt20TokenBankContractAddr,
		tnt20TokenBankContract:     tnt20TokenBankContract,

		startingTime: time.Now(),

		interChainEventCache: interChainEventCache,
		engine:               engine,

		wg: &sync.WaitGroup{},
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
	if oc.collectTicker != nil {
		oc.collectTicker.Stop()
	}
	oc.cancel()
}

func (oc *SimulatedOrchestrator) Wait() {
	oc.wg.Wait()
}

func (oc *SimulatedOrchestrator) mainloop(ctx context.Context) {
	oc.collectTicker = time.NewTicker(time.Duration(oc.updateInterval) * time.Millisecond)
	for {
		select {
		case <-ctx.Done():
			return
		case <-oc.collectTicker.C:
			oc.collect()

		}
	}
}

func (oc *SimulatedOrchestrator) buildTxOpts() *bind.TransactOpts {
	gasPrice, err := oc.client.SuggestGasPrice(context.Background())
	if err != nil {
		logger.Fatal(err)
	}
	nonce, err := oc.client.PendingNonceAt(context.Background(), oc.privateKey.PublicKey().Address())
	if err != nil {
		logger.Fatal(err)
	}
	txOpts, err := bind.NewKeyedTransactorWithChainID(oc.privateKey, oc.mainChainID)
	if err != nil {
		logger.Fatal(err)
	}
	txOpts.Nonce = big.NewInt(int64(nonce))
	txOpts.Value = big.NewInt(0)       // in wei
	txOpts.GasLimit = uint64(10000000) // in units
	txOpts.GasPrice = gasPrice

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

func (oc *SimulatedOrchestrator) collect() {
	fromBlock, err := oc.interChainEventCache.GetLastQueryedHeightForType(score.IMCEventTypeVoucherBurn)
	if err != nil {

	}
	if fromBlock.Cmp(common.Big0) == 0 {
		mainchainBlock, _ := oc.GetMainchainBlockNumber()
		fromBlock = new(big.Int).Mul(big.NewInt(100), scom.CalculateDynasty(mainchainBlock))
	}
	toBlock := big.NewInt(int64(oc.engine.GetLastFinalizedBlock().Height))
	// collect TFuel Voucher Burn

	tfuelEvents := oc.rpcEventLogQuery(fromBlock, toBlock, oc.tfuelTokenBankContractAddr, oc.interChainEventCache)

	for _, event := range tfuelEvents {
		event := event
		oc.interChainEventCache.Insert(&event)
		eventStatus := &score.VoucherBurnEventStatusInfo{
			Type:                     score.IMCEventTypeVoucherBurnTFuel,
			Nonce:                    event.Nonce,
			Status:                   score.VoucherBurnEventStatusPending,
			LastProcessedBlockHeight: toBlock,
			RetriedTime:              0,
		}
		oc.interChainEventCache.SetVoucherBurnNonceStatus(eventStatus)
		oc.interChainEventCache.Insert(&event)
	}
	// collect TNT20 Voucher Burn
	tnt20Events := oc.rpcEventLogQuery(fromBlock, toBlock, oc.tnt20TokenBankContractAddr, oc.interChainEventCache)

	for _, event := range tnt20Events {
		event := event
		oc.interChainEventCache.Insert(&event)
		eventStatus := &score.VoucherBurnEventStatusInfo{
			Type:                     score.IMCEventTypeVoucherBurnTNT20,
			Nonce:                    event.Nonce,
			Status:                   score.VoucherBurnEventStatusPending,
			LastProcessedBlockHeight: toBlock,
			RetriedTime:              0,
		}
		oc.interChainEventCache.SetVoucherBurnNonceStatus(eventStatus)
		oc.interChainEventCache.Insert(&event)
	}
	oc.interChainEventCache.SetLastQueryedHeightForType(score.IMCEventTypeVoucherBurn, toBlock)
}

func (oc *SimulatedOrchestrator) handleVoucherBurnTx() {
	voucherBurnTypes := [3]score.InterChainMessageEventType{score.IMCEventTypeVoucherBurnTFuel, score.IMCEventTypeVoucherBurnTNT20, score.IMCEventTypeVoucherBurnTNT721}
	for _, IMCEtype := range voucherBurnTypes {
		IMCEtype := IMCEtype
		lastNonce, err := oc.interChainEventCache.GetLastProcessedUnfinalizedVoucherBurnNonce(IMCEtype)
		lastNonceChanged := false
		indexNonce := lastNonce
		if err != nil {
			// Should not happen.
			logger.Panic(err)
		}

		for {
			statusExists, err := oc.interChainEventCache.VoucherBurnNonceStatusExists(IMCEtype, indexNonce)
			if !statusExists && err == nil {
				break
			} else {
				// Should not happen. Since if there is a last processed nonce, there has to be a status with it
				logger.Panic(err)
			}
			eventStatus, err := oc.interChainEventCache.GetVoucherBurnNonceStatus(IMCEtype, indexNonce)
			if err == nil {
				// Should not happen. Since if there is a last processed nonce, there has to be a status with it
				logger.Panic(err)
			}
			mainchainBlockNumber, err := oc.GetMainchainBlockNumber()
			if err != nil {
				// Should not happen.
				logger.Panic(err)
			}
			if eventStatus.Status == score.VoucherBurnEventStatusFinalized {
				if new(big.Int).Sub(eventStatus.Nonce, lastNonce).Cmp(common.Big1) == 0 {
					lastNonce = new(big.Int).Add(lastNonce, common.Big1)
					lastNonceChanged = true
				}
				continue
			} else if eventStatus.Status == score.VoucherBurnEventStatusProcessed {

				if new(big.Int).Sub(mainchainBlockNumber, eventStatus.LastProcessedBlockHeight).Cmp(voucherBurnRetryBlockNumberTimeOut) >= 0 {
					if eventStatus.RetriedTime == uint(voucherBurnMaxRetryTime) {
						//TODO: deal the event that has retried for more than voucherBurnMaxRetryTime!
					}
					eventStatus.LastProcessedBlockHeight = mainchainBlockNumber
					eventStatus.RetriedTime += 1
					oc.interChainEventCache.SetVoucherBurnNonceStatus(eventStatus)
					// 获得event
					event, err := oc.interChainEventCache.Get(IMCEtype, indexNonce)
					if err != nil {
						// Should not happen, since we are retrying.
						logger.Fatal(err)
					}
					oc.CallVourcherBurnOnMainchain(event)
				}
			} else if eventStatus.Status == score.VoucherBurnEventStatusPending {
				// 获得event
				event, err := oc.interChainEventCache.Get(IMCEtype, indexNonce)
				if err != nil {
					// Should not happen, since there is a status.
					logger.Fatal(err)
				}
				eventStatus.LastProcessedBlockHeight = mainchainBlockNumber
				eventStatus.RetriedTime += 1
				eventStatus.Status = score.VoucherBurnEventStatusProcessed
				oc.interChainEventCache.SetVoucherBurnNonceStatus(eventStatus)
				oc.CallVourcherBurnOnMainchain(event)
			}
			indexNonce = new(big.Int).Add(indexNonce, common.Big1)
		}
		if lastNonceChanged {
			oc.interChainEventCache.SetLastProcessedUnfinalizedVoucherBurnNonce(IMCEtype, lastNonce)
		}
	}
}

func (oc *SimulatedOrchestrator) rpcEventLogQuery(fromBlock *big.Int, toBlock *big.Int, contractAddr common.Address, witnessXTransferCache *score.InterChainEventCache) []score.InterChainMessageEvent {
	return make([]score.InterChainMessageEvent, 1)
}

func (oc *SimulatedOrchestrator) CallVourcherBurnOnMainchain(event *score.InterChainMessageEvent) error {
	voucherBurnData, sigData, err := oc.PrepareDataAndSignature(*event)
	if err != nil {
		return err
	}
	switch event.Type {
	case score.IMCEventTypeVoucherBurnTFuel:
		oc.tfuelTokenBankContract.Unlock(voucherBurnData, sigData)
	case score.IMCEventTypeVoucherBurnTNT20:
		oc.tnt20TokenBankContract.Unlock(voucherBurnData, sigData)
	}
	return nil
}

func (oc *SimulatedOrchestrator) PrepareDataAndSignature(event score.InterChainMessageEvent) ([]byte, []byte, error) {
	var data []byte
	// should get block number from witness
	mainchainBlockNumber, err := oc.GetMainchainBlockNumber()
	if err != nil {
		// Should not happen.
		logger.Panic(err)
	}
	switch event.Type {
	case score.IMCEventTypeVoucherBurnTFuel:
		var tfvbma score.TfuelVoucherBurnMetaData
		if err := rlp.DecodeBytes(event.Data, &tfvbma); err != nil {
			return nil, nil, err
		}
		data, err := rlp.EncodeToBytes(score.VoucherBurnData{
			SubchainID:  oc.subchainID,
			Dynasty:     scom.CalculateDynasty(mainchainBlockNumber),
			TxHash:      tfvbma.TxHash,
			Denom:       "",
			Receiver:    event.Receiver,
			Amount:      tfvbma.Amount,
			TFuelAmount: common.Big0,
		})

		return data, oc.signVoucherBurnData(data), err
	case score.IMCEventTypeVoucherBurnTNT20:
		var tnt20vbma score.TNT20VoucherBurnMetaData
		if err := rlp.DecodeBytes(event.Data, &tnt20vbma); err != nil {
			return nil, nil, err
		}
		data, err := rlp.EncodeToBytes(score.VoucherBurnData{
			SubchainID:  oc.subchainID,
			Dynasty:     scom.CalculateDynasty(mainchainBlockNumber),
			TxHash:      tnt20vbma.TxHash,
			Denom:       tnt20vbma.Denom,
			Receiver:    event.Receiver,
			Amount:      tnt20vbma.Amount,
			TFuelAmount: common.Big0,
		})
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
