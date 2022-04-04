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

const mainchainBlockIntervalMilliseconds int64 = 1000 // seconds

// SimulatedMainchainWitness is a simulated mainchain witness for end-to-end testing
type SimulatedMainchainWitness struct {
	mu sync.RWMutex

	subchainID        *big.Int
	witnessedDynasty  *big.Int
	validatorSetCache map[*big.Int]*score.ValidatorSet
	updateTimer       *time.Timer
	startingTime      time.Time
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
		witnessedDynasty:  big.NewInt(0), // will be updated in the first update() call
		validatorSetCache: make(map[*big.Int]*score.ValidatorSet),
		updateTimer:       time.NewTimer(time.Duration(100) * time.Millisecond),
		startingTime:      time.Now(),
	}

	return mw
}

func (mw *SimulatedMainchainWitness) Start(ctx context.Context) {
	go mw.mainloop(ctx)
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
	for {
		select {
		case <-ctx.Done():
			return
		case <-mw.updateTimer.C:
			mw.update()
		}
	}
}

func (mw *SimulatedMainchainWitness) update() {
	mainchainBlockNumber, err := mw.GetMainchainBlockNumber()
	if err != nil {
		logger.Warnf("failed to get the mainchain block number %v\n", err)
		return
	}

	dynasty := scom.CalculateDynasty(mainchainBlockNumber)
	if dynasty.Cmp(mw.witnessedDynasty) > 0 { // needs to update the cache
		mw.updateValidatorSetCache(dynasty)
		mw.witnessedDynasty = dynasty
		logger.Infof("updated the witnessed dynasty to %v\n", dynasty)
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
	if dynasty.Cmp(big.NewInt(0)) == 0 { // start with a single validator
		stake := big.NewInt(100000000000)
		v2 := score.NewValidator(validatorAddrList[1], stake)
		validatorSet.AddValidator(v2)
	} else if dynasty.Cmp(big.NewInt(1)) == 0 { // add more validators
		stake := big.NewInt(300000000000)
		v1 := score.NewValidator(validatorAddrList[0], stake)
		v2 := score.NewValidator(validatorAddrList[1], stake)
		v3 := score.NewValidator(validatorAddrList[2], stake)
		validatorSet.AddValidator(v1)
		validatorSet.AddValidator(v2)
		validatorSet.AddValidator(v3)
	} else if dynasty.Cmp(big.NewInt(2)) == 0 { // remove validators
		stake1 := big.NewInt(300000000000)
		stake2 := big.NewInt(6000000000)
		v1 := score.NewValidator(validatorAddrList[0], stake1)
		v2 := score.NewValidator(validatorAddrList[1], stake2)
		validatorSet.AddValidator(v1)
		validatorSet.AddValidator(v2)
	} else { // swap out the entire validator set
		stake3 := big.NewInt(7000000000)
		stake4 := big.NewInt(22000000000)
		v3 := score.NewValidator(validatorAddrList[2], stake3)
		v4 := score.NewValidator(validatorAddrList[3], stake4)
		validatorSet.AddValidator(v3)
		validatorSet.AddValidator(v4)
	}

	mw.validatorSetCache[dynasty] = validatorSet

	logger.Infof("Witnessed validator set for dynasty %v\n", dynasty)
	for _, v := range validatorSet.Validators() {
		logger.Infof("Validator: %v", v)
	}

	return validatorSet, nil
}
