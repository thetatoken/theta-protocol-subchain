package witness

import (
	// "context"
	"log"
	"math/big"
	"sync"

	// "github.com/thetatoken/theta/crypto"
	scom "github.com/thetatoken/thetasubchain/common"
	// sct "github.com/thetatoken/thetasubchain/contracts"

	// "github.com/thetatoken/thetasubchain/eth/abi/bind"
	"github.com/thetatoken/theta/common"
	ec "github.com/thetatoken/thetasubchain/eth/ethclient"
)

type MainchainWitness struct {
	mu                   sync.RWMutex
	MainchainBlockNumber *big.Int
	SubchainID           *big.Int
	gasPriceLimit        *big.Int
	Dynasty              *big.Int

	chainID *big.Int
	// privateKey           *crypto.PrivateKey
	registerContractAddr common.Address
	ercContractAddr      common.Address

	// client *ethclient.Client
	// registerContract *sct.SubchainRegister
	// ercContract      *sct.SubchainERC
}

// NewMainchainWitness creates a new MainchainWitness
func NewMainchainWitness(
	// privateKey *crypto.PrivateKey,
	gasPriceLimit *big.Int,
	subchainID *big.Int,
	registerContractAddr common.Address,
	ercContractAddr common.Address,
) *MainchainWitness {
	client, err := ec.Dial(scom.CfgMainchainAdaptorURL)
	if err != nil {
		log.Fatalf("the eth client failed to connect %v\n", err)
	}
	log.Fatalf("%v", client)
	// subchainRegisterContract, err := sct.NewSubchainRegister(registerContractAddr, client)
	// if err != nil {
	// 	log.Fatalf("failed to create subchain register contract %v\n", err)
	// }
	// subchainERCContract, err := sct.NewSubchainERC(ercContractAddr, client)
	// if err != nil {
	// 	log.Fatalf("failed to create erc contract %v\n", err)
	// }
	// chainID, err := client.ChainID(context.Background())
	// if err != nil {
	// 	log.Fatalf("failed to get chainID %v\n", err)
	// }
	// log.Printf("Create transfer validator for chain %d\n", chainID)
	// tipBlockHeader, err := client.HeaderByNumber(context.Background(), nil)
	// if err != nil {
	// 	log.Fatalf("failed to get the highest block number %v\n", err)
	// }
	// dynasty := new(big.Int).Div(tipBlockHeader.Number, big.NewInt(100))
	// mm := &MainchainWitness{
	// 	MainchainBlockNumber: tipBlockHeader.Number,
	// 	SubchainID:           subchainID,
	// 	Dynasty:              dynasty,
	// 	gasPriceLimit:        gasPriceLimit,

	// 	client:  client,
	// 	chainID: chainID,
	// 	// privateKey: privateKey,

	// 	registerContractAddr: registerContractAddr,
	// 	ercContractAddr:      ercContractAddr,
	// 	// registerContract:     subchainRegisterContract,
	// 	// ercContract:          subchainERCContract,
	// }

	mm := &MainchainWitness{}
	return mm
}

// func (_MainchainWitness *MainchainWitness) GetValidatorSetWithBlockHeight(opts *bind.CallOpts, subchainID *big.Int) ([]ethcommon.Address, error) {
// 	// TODO: later test block height
// 	return _MainchainWitness.registerContract.GetLegalValidators(opts, subchainID)
// }
