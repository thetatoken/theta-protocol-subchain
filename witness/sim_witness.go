package witness

import (
	"context"
	"math/big"
	"sync"
	"time"

	scom "github.com/thetatoken/thetasubchain/common"
	score "github.com/thetatoken/thetasubchain/core"

	"github.com/thetatoken/theta/common"
)

const mainchainBlockIntervalMilliseconds int64 = 2000 // millseconds

// SimulatedMainchainWitness is a simulated mainchain witness for end-to-end testing
type SimulatedMainchainWitness struct {
	subchainID        *big.Int
	witnessedDynasty  *big.Int
	validatorSetCache map[*big.Int]*score.ValidatorSet
	updateTicker      *time.Ticker
	startingTime      time.Time

	// Life cycle
	wg     *sync.WaitGroup
	ctx    context.Context
	cancel context.CancelFunc
}

// NewSimulatedMainchainWitness creates a new SimulatedMainchainWitness
func NewSimulatedMainchainWitness(
	ethClientAddress string,
	subchainID *big.Int,
	registerContractAddr common.Address,
	ercContractAddr common.Address,
) *SimulatedMainchainWitness {
	mw := &SimulatedMainchainWitness{
		subchainID:        subchainID,
		witnessedDynasty:  nil, // will be updated in the first update() call
		validatorSetCache: make(map[*big.Int]*score.ValidatorSet),
		startingTime:      time.Now(),
		wg:                &sync.WaitGroup{},
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
	blockNumber := int64((time.Now().Sub(mw.startingTime)).Milliseconds()) / mainchainBlockIntervalMilliseconds
	return big.NewInt(int64(blockNumber)), nil
}

func (mw *SimulatedMainchainWitness) GetMainchainBlockNumberUint() (uint64, error) {
	blockNumber, _ := mw.GetMainchainBlockNumber()
	return blockNumber.Uint64(), nil
}

func (mw *SimulatedMainchainWitness) GetValidatorSetByDynasty(dynasty *big.Int) (*score.ValidatorSet, error) {
	validatorSet, ok := mw.validatorSetCache[dynasty]
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
}

func (mw *SimulatedMainchainWitness) updateValidatorSetCache(dynasty *big.Int) (*score.ValidatorSet, error) {
	// Simulate validator set updates
	validatorAddrList := []string{
		"0x9F1233798E905E173560071255140b4A8aBd3Ec6",
		"0x2E833968E5bB786Ae419c4d13189fB081Cc43bab",
		"0xC15E24083152dD76Ae6FC2aEb5269FF23d70330B",
		"0x7631958d57Cf6a5605635a5F06Aa2ae2e000820e",
	}

	validatorSet := score.NewValidatorSet(dynasty)
	if dynasty.Cmp(big.NewInt(0)) == 0 {
		// Dynasty 0: Start with a single validator

		stake := big.NewInt(100000000)
		v2 := score.NewValidator(validatorAddrList[1], stake)
		validatorSet.AddValidator(v2)
	} else if dynasty.Cmp(big.NewInt(1)) == 0 {
		// Dynasty 1: Add more validators

		stake1 := big.NewInt(6000000000)
		stake2 := big.NewInt(300000000000)
		stake3 := big.NewInt(6000000000)
		v1 := score.NewValidator(validatorAddrList[0], stake1)
		v2 := score.NewValidator(validatorAddrList[1], stake2)
		v3 := score.NewValidator(validatorAddrList[2], stake3)
		validatorSet.AddValidator(v1)
		validatorSet.AddValidator(v2)
		validatorSet.AddValidator(v3)
	} else if dynasty.Cmp(big.NewInt(2)) == 0 {
		// Dynasty 2: Remove validators

		stake1 := big.NewInt(6000000000)
		stake2 := big.NewInt(300000000000)
		v1 := score.NewValidator(validatorAddrList[0], stake1)
		v2 := score.NewValidator(validatorAddrList[1], stake2)
		validatorSet.AddValidator(v1)
		validatorSet.AddValidator(v2)
	} else {
		// Dynasty 3: Remove some validators, and add additional validators

		stake2 := big.NewInt(300000000000)
		stake3 := big.NewInt(7000000000)
		stake4 := big.NewInt(22000000000)
		v2 := score.NewValidator(validatorAddrList[1], stake2)
		v3 := score.NewValidator(validatorAddrList[2], stake3)
		v4 := score.NewValidator(validatorAddrList[3], stake4)
		validatorSet.AddValidator(v2)
		validatorSet.AddValidator(v3)
		validatorSet.AddValidator(v4)
	}

	mw.validatorSetCache[dynasty] = validatorSet

	logger.Infof("Witnessed validator set for dynasty %v", dynasty)
	for _, v := range validatorSet.Validators() {
		logger.Infof("Validator: %v", v)
	}

	return validatorSet, nil
}
