package witness

import (
	"context"
	"math/big"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	// "github.com/thetatoken/theta/crypto"
	// scom "github.com/thetatoken/thetasubchain/common"
	sct "github.com/thetatoken/thetasubchain/contracts"

	// "github.com/thetatoken/thetasubchain/eth/abi/bind"
	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/thetasubchain/eth/abi/bind"
	ec "github.com/thetatoken/thetasubchain/eth/ethclient"
)

type MainchainWitness struct {
	mu                   sync.RWMutex
	MainchainBlockNumber *big.Int
	SubchainID           *big.Int
	gasPriceLimit        *big.Int
	StableDynasty        *big.Int
	Dynasty              *big.Int
	IsDynastyChanged     bool

	chainID *big.Int
	// privateKey           *crypto.PrivateKey
	RegisterContractAddr common.Address
	ErcContractAddr      common.Address

	updateTimer *time.Timer

	ValidatorAndStakeCache map[*big.Int]ValidatorAndStakeCheckpoint
	Client                 *ec.Client

	RegisterContract *sct.SubchainRegister
	ErcContract      *sct.SubchainERC
}

type ValidatorAndStakeCheckpoint struct {
	ValidatorSet    []common.Address
	ValidatorStakes []*big.Int
}

// NewMainchainWitness creates a new MainchainWitness
func NewMainchainWitness(
	// privateKey *crypto.PrivateKey,
	ethClientAddress string,
	subchainID *big.Int,
	registerContractAddr common.Address,
	ercContractAddr common.Address,
) *MainchainWitness {
	client, err := ec.Dial(ethClientAddress)
	if err != nil {
		log.Fatalf("the eth client failed to connect %v\n", err)
	}
	// log.Fatalf("client connected !!! %v\n subchain ID is %v, register addr is %v, erc addr is %v", client, subchainID, registerContractAddr, ercContractAddr)
	subchainRegisterContract, err := sct.NewSubchainRegister(registerContractAddr, client)
	if err != nil {
		log.Fatalf("failed to create subchain register contract %v\n", err)
	}
	subchainERCContract, err := sct.NewSubchainERC(ercContractAddr, client)
	if err != nil {
		log.Fatalf("failed to create erc contract %v\n", err)
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatalf("failed to get chainID %v\n", err)
	}
	log.Printf("Create transfer validator for chain %d\n", chainID)
	tipBlockHeader, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("failed to get the highest block number %v\n", err)
	}
	dynasty := new(big.Int).Div(tipBlockHeader.Number, big.NewInt(100))
	mw := &MainchainWitness{
		MainchainBlockNumber: tipBlockHeader.Number,
		SubchainID:           subchainID,
		Dynasty:              dynasty,

		Client:  client,
		chainID: chainID,

		RegisterContractAddr: registerContractAddr,
		ErcContractAddr:      ercContractAddr,
		RegisterContract:     subchainRegisterContract,
		ErcContract:          subchainERCContract,
	}
	mw.updateTimer = time.NewTimer(time.Duration(1000) * time.Millisecond)
	mw.ValidatorAndStakeCache = make(map[*big.Int]ValidatorAndStakeCheckpoint)
	return mw
}

// func (_MainchainWitness *MainchainWitness) GetValidatorSetWithDynasty(opts *bind.CallOpts, subchainID *big.Int) ([]ethcommon.Address, error) {
// 	// TODO: later test block height
// 	return _MainchainWitness.registerContract.GetLegalValidators(opts, subchainID)
// }

func (mw *MainchainWitness) Start(ctx context.Context) {
	go mw.mainloop(ctx)
}

func (mw *MainchainWitness) GetMainchainBlockNumber() *big.Int {
	tipBlockNumber, err := mw.Client.BlockNumber(context.Background())
	if err != nil {
		log.Fatalf("failed to get the highest block number %v\n", err)
	}
	return big.NewInt(int64(tipBlockNumber))
}

func (mw *MainchainWitness) GetMainchainBlockNumberUint() uint64 {
	tipBlockNumber, err := mw.Client.BlockNumber(context.Background())
	if err != nil {
		log.Fatalf("failed to get the highest block number %v\n", err)
	}
	return tipBlockNumber
}

func (mw *MainchainWitness) GetValidatorSetWithDynasty(opts *bind.CallOpts, subchainID *big.Int, dynasty *big.Int) ([]common.Address, error) {
	// TODO: later test block height
	// dynasty := new(big.Int).Div(big.NewInt(tipBlockNumber), big.NewInt(100))
	vsc, ok := mw.ValidatorAndStakeCache[dynasty]
	if ok {
		return vsc.ValidatorSet, nil
	} else {
		resValidatorSet, err := mw.RegisterContract.GetLegalValidators(opts, mw.SubchainID)
		if err != nil {
			log.Fatalf("failed to get the validator set from mainchain %v\n", err)
		}
		new_vsc := ValidatorAndStakeCheckpoint{
			ValidatorSet:    resValidatorSet,
			ValidatorStakes: nil,
		}
		mw.ValidatorAndStakeCache[dynasty] = new_vsc
		return new_vsc.ValidatorSet, nil
	}
	// return mw.RegisterContract.GetLegalValidators(opts, mw.SubchainID)
}

func (mw *MainchainWitness) mainloop(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-mw.updateTimer.C:
			mw.update()
		}
	}
}

func (mw *MainchainWitness) update() {
	// 更新区块高度、Dynasty
	// 如果 Dynasty变更，更新Validator set、 stake的情况
	mw.MainchainBlockNumber = mw.GetMainchainBlockNumber()
	newDynasty := new(big.Int).Div(mw.MainchainBlockNumber, big.NewInt(100))
	if mw.StableDynasty == nil {
		mw.StableDynasty = newDynasty
	}
	// validatorSet, _ := mw.GetValidatorSetWithDynasty(nil, mw.SubchainID)
	// log.Infof("validators %v", validatorSet)
	// log.Fatalf("hihihi")
	if newDynasty != mw.Dynasty {
		log.Infof("Dynasty changed! ", newDynasty)
		validatorSet, _ := mw.GetValidatorSetWithDynasty(nil, mw.SubchainID, mw.Dynasty)
		log.Infof("validators %v", validatorSet)
		// mw.IsDynastyChanged = true
		// validatorSet, validatorStake := mw.RegisterContract.GetLegalValidatorsAndStakes(mw.SubchainID)
		// newDynastyCheckpoint := &ValidatorAndStakeCheckpoint{
		// 	ValidatorSet:    validatorSet,
		// 	ValidatorStakes: validatorStake,
		// }
		// mw.ValidatorAndStakeCache[newDynasty] = newDynastyCheckpoint
	}
}
