package witness

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	// "github.com/thetatoken/theta/crypto"
	scom "github.com/thetatoken/thetasubchain/common"
	sct "github.com/thetatoken/thetasubchain/contracts"
	score "github.com/thetatoken/thetasubchain/core"

	// "github.com/thetatoken/thetasubchain/eth/abi/bind"
	"github.com/thetatoken/theta/common"
	ec "github.com/thetatoken/thetasubchain/eth/ethclient"
)

// SubchainRegisterSendToSubchainEvent
var logger *log.Entry = log.WithFields(log.Fields{"prefix": "witness"})

type MainchainWitness struct {
	mainChainID          *big.Int
	subchainID           *big.Int
	witnessedDynasty     *big.Int
	validatorSetCache    map[string]*score.ValidatorSet
	crossChainEventCache *score.CrossChainEventCache
	client               *ec.Client
	updateTicker         *time.Ticker

	registerContractAddr common.Address
	ercContractAddr      common.Address
	registerContract     *sct.SubchainRegister
	ercContract          *sct.SubchainERC

	updateInterval int

	// Life cycle
	wg     *sync.WaitGroup
	ctx    context.Context
	cancel context.CancelFunc

	lastQueryedMainChainHeight *big.Int
}

// NewMainchainWitness creates a new MainchainWitness
func NewMainchainWitness(
	ethClientAddress string,
	subchainID *big.Int,
	registerContractAddr common.Address,
	ercContractAddr common.Address,
	updateInterval int,
	crossChainEventCache *score.CrossChainEventCache,
) *MainchainWitness {
	client, err := ec.Dial(ethClientAddress)
	if err != nil {
		logger.Fatalf("the eth client failed to connect %v\n", err)
	}
	subchainRegisterContract, err := sct.NewSubchainRegister(registerContractAddr, client)
	if err != nil {
		logger.Fatalf("failed to create subchain register contract %v\n", err)
	}
	subchainERCContract, err := sct.NewSubchainERC(ercContractAddr, client)
	if err != nil {
		logger.Fatalf("failed to create erc contract %v\n", err)
	}
	mainChainID, err := client.ChainID(context.Background())
	if err != nil {
		logger.Fatalf("failed to get the chainID of the main chain %v\n", err)
	}
	logger.Printf("Create transfer validator for chain %d\n", mainChainID)

	mw := &MainchainWitness{
		mainChainID:       mainChainID,
		subchainID:        subchainID,
		witnessedDynasty:  nil, // will be updated in the first update() call
		validatorSetCache: make(map[string]*score.ValidatorSet),
		client:            client,

		registerContractAddr: registerContractAddr,
		ercContractAddr:      ercContractAddr,
		registerContract:     subchainRegisterContract,
		ercContract:          subchainERCContract,

		updateInterval: updateInterval,

		crossChainEventCache: crossChainEventCache,

		lastQueryedMainChainHeight: big.NewInt(0),

		wg: &sync.WaitGroup{},
	}
	return mw
}

func (mw *MainchainWitness) Start(ctx context.Context) {
	c, cancel := context.WithCancel(ctx)
	mw.ctx = c
	mw.cancel = cancel

	mw.wg.Add(1)
	go mw.mainloop(ctx)
}

func (mw *MainchainWitness) Stop() {
	if mw.updateTicker != nil {
		mw.updateTicker.Stop()
	}
	mw.cancel()
}

func (mw *MainchainWitness) Wait() {
	mw.wg.Wait()
}

// TODO: make sure the block number returned by the client.BlockNumber() call is the lastest *finalized* block number
func (mw *MainchainWitness) GetMainchainBlockNumber() (*big.Int, error) {
	blockNumber, err := mw.client.BlockNumber(context.Background())
	if err != nil {
		return nil, err
	}
	return big.NewInt(int64(blockNumber)), nil
}

// TODO: make sure the block number returned by the client.BlockNumber() call is the lastest *finalized* block number
func (mw *MainchainWitness) GetMainchainBlockNumberUint() (uint64, error) {
	blockNumber, err := mw.client.BlockNumber(context.Background())
	if err != nil {
		return 0, err
	}
	return blockNumber, nil
}

func (mw *MainchainWitness) GetValidatorSetByDynasty(dynasty *big.Int) (*score.ValidatorSet, error) {
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

func (mw *MainchainWitness) mainloop(ctx context.Context) {
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

func (mw *MainchainWitness) update() {
	mainchainBlockNumber, err := mw.GetMainchainBlockNumber()
	if err != nil {
		logger.Warnf("failed to get the mainchain block number %v\n", err)
		return
	}

	dynasty := scom.CalculateDynasty(mainchainBlockNumber)
	if mw.witnessedDynasty == nil || dynasty.Cmp(mw.witnessedDynasty) > 0 { // needs to update the cache
		mw.updateValidatorSetCache(dynasty)
		mw.witnessedDynasty = dynasty
	}
}

func (mw *MainchainWitness) collectCrossChainTransferEvent() []score.CrossChainTransferEvent {
	fromBlock := mw.lastQueryedMainChainHeight
	toBlock, err := mw.GetMainchainBlockNumber()
	if err != nil {
		logger.Warnf("failed to get the mainchain block number %v\n", err)
		return make([]score.CrossChainTransferEvent, 1)
	}
	return scom.RpcEventLogQuery(fromBlock, toBlock, mw.registerContractAddr, mw.crossChainEventCache)
}

func (mw *MainchainWitness) updateValidatorSetCache(dynasty *big.Int) (*score.ValidatorSet, error) {
	validatorAddrs, validatorStakes, err := mw.registerContract.GetValidatorAndStakeSetWithBlockHeight(nil, mw.subchainID, dynasty.Mul(dynasty, big.NewInt(100)))
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

	return validatorSet, nil
}

func (mw *MainchainWitness) GetCrossChainEventCache() *score.CrossChainEventCache {
	return mw.crossChainEventCache
}
