package witness

import (
	"context"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/rlp"
	scom "github.com/thetatoken/thetasubchain/common"
	score "github.com/thetatoken/thetasubchain/core"

	"github.com/thetatoken/theta/common"
)

const mainchainBlockIntervalMilliseconds int64 = 2000 // millseconds

// SimulatedMainchainWitness is a simulated mainchain witness for end-to-end testing
type SimulatedMainchainWitness struct {
	subchainID        *big.Int
	witnessedDynasty  *big.Int
	validatorSetCache map[string]*score.ValidatorSet
	updateTicker      *time.Ticker
	startingTime      time.Time

	// Life cycle
	wg                   *sync.WaitGroup
	ctx                  context.Context
	cancel               context.CancelFunc
	crossChainEventCache *score.InterChainEventCache

	lastSimEventNonce *big.Int
}

// NewSimulatedMainchainWitness creates a new SimulatedMainchainWitness
func NewSimulatedMainchainWitness(
	ethClientAddress string,
	subchainID *big.Int,
	registerContractAddr common.Address,
	ercContractAddr common.Address,
	crossChainEventCache *score.InterChainEventCache,
) *SimulatedMainchainWitness {
	mw := &SimulatedMainchainWitness{
		subchainID:           subchainID,
		witnessedDynasty:     nil, // will be updated in the first update() call
		validatorSetCache:    make(map[string]*score.ValidatorSet),
		startingTime:         time.Now(),
		crossChainEventCache: crossChainEventCache,
		wg:                   &sync.WaitGroup{},
		lastSimEventNonce:    big.NewInt(2),
	}

	amount := big.NewInt(88888)
	nonce := big.NewInt(1)
	mainchainBlockNumber := big.NewInt(0)
	event := mw.generateInterChainEventForTFuelTransfer(amount, nonce, mainchainBlockNumber)

	err := mw.crossChainEventCache.Insert(event)
	if err != nil {
		logger.Panicf("Insert Fail!! %v", err)
	}
	return mw
}

func (mw *SimulatedMainchainWitness) Start(ctx context.Context) {
	c, cancel := context.WithCancel(ctx)
	mw.ctx = c
	mw.cancel = cancel

	mw.wg.Add(1)
	go mw.mainloop(ctx)
}

func (mw *SimulatedMainchainWitness) Stop() {
	if mw.updateTicker != nil {
		mw.updateTicker.Stop()
	}
	mw.cancel()
}

func (mw *SimulatedMainchainWitness) Wait() {
	mw.wg.Wait()
}

func (mw *SimulatedMainchainWitness) GetMainchainBlockNumber() (*big.Int, error) {
	blockNumber := int64((time.Since(mw.startingTime)).Milliseconds()) / mainchainBlockIntervalMilliseconds
	return big.NewInt(int64(blockNumber)), nil
}

func (mw *SimulatedMainchainWitness) GetMainchainBlockNumberUint() (uint64, error) {
	blockNumber, _ := mw.GetMainchainBlockNumber()
	return blockNumber.Uint64(), nil
}

func (mw *SimulatedMainchainWitness) GetValidatorSetByDynasty(dynasty *big.Int) (*score.ValidatorSet, error) {
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

func (mw *SimulatedMainchainWitness) mainloop(ctx context.Context) {
	mw.updateTicker = time.NewTicker(time.Duration(1000) * time.Millisecond)
	for {
		select {
		case <-ctx.Done():
			return
		case <-mw.updateTicker.C:
			mw.update()
		}
	}
}

func (mw *SimulatedMainchainWitness) update() {
	mainchainBlockNumber, err := mw.GetMainchainBlockNumber()
	logger.Debugf("witnessed main chain block height: %v", mainchainBlockNumber)

	if err != nil {
		logger.Warnf("failed to get the mainchain block number %v", err)
		return
	}

	dynasty := scom.CalculateDynasty(mainchainBlockNumber)
	if mw.witnessedDynasty == nil || dynasty.Cmp(mw.witnessedDynasty) > 0 { // needs to update the cache
		mw.updateValidatorSetCache(dynasty)
		mw.witnessedDynasty = dynasty
		logger.Infof("updated the witnessed dynasty to %v", dynasty)
	}

	for i := 0; i < 3; i++ {
		amount := big.NewInt(int64(9999999 + i*1234567))
		event := mw.generateInterChainEventForTFuelTransfer(amount, mw.lastSimEventNonce, mainchainBlockNumber)
		mw.crossChainEventCache.Insert(event)
		// logger.Infof("Inserted Event %v", event)
		mw.lastSimEventNonce = new(big.Int).Add(mw.lastSimEventNonce, big.NewInt(1))
	}
}

func (mw *SimulatedMainchainWitness) updateValidatorSetCache(dynasty *big.Int) (*score.ValidatorSet, error) {
	// Simulate validator set updates
	validatorAddrList := []string{
		"0x2E833968E5bB786Ae419c4d13189fB081Cc43bab",
	}
	validatorSet := score.NewValidatorSet(dynasty)
	stake := big.NewInt(100000000)
	v := score.NewValidator(validatorAddrList[0], stake)
	validatorSet.AddValidator(v)
	mw.validatorSetCache[dynasty.String()] = validatorSet

	logger.Infof("Witnessed validator set for dynasty %v", dynasty)
	for _, v := range validatorSet.Validators() {
		logger.Infof("Validator: %v", v)
	}

	return validatorSet, nil
}

func (mw *SimulatedMainchainWitness) GetInterChainEventCache() *score.InterChainEventCache {
	return mw.crossChainEventCache
}

func (mw *SimulatedMainchainWitness) generateInterChainEventForTFuelTransfer(amount *big.Int, nonce *big.Int, mainchainBlockNumber *big.Int) *score.InterChainMessageEvent {
	tfuelDenom := score.TFuelDenom("mainnet")
	data, err := rlp.EncodeToBytes(score.TfuelTransferMetaData{
		Denom:  tfuelDenom,
		Amount: amount,
	})
	if err != nil {
		logger.Panicf("failed to get encode inter-chain message event data for TFuel transfer: %v", err)
	}

	event := &score.InterChainMessageEvent{
		Type:        score.IMCEventTypeCrossChainTFuelTransfer,
		Sender:      common.HexToAddress("0x2E833968E5bB786Ae419c4d13189fB081Cc43bab"),
		Receiver:    common.HexToAddress("0x2E833968E5bB786Ae419c4d13189fB081Cc43bab"),
		Data:        data,
		Nonce:       nonce,
		BlockNumber: mainchainBlockNumber,
	}

	return event
}