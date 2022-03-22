package witness

import (
	// "context"
	"log"
	"math/big"
	"sync"

	// "github.com/thetatoken/theta/crypto"
	scom "github.com/thetatoken/thetasubchain/common"
	// sct "github.com/thetatoken/thetasubchain/contracts"

	// "github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ec "github.com/ethereum/go-ethereum/ethclient"
)

type MainchainMonitor struct {
	mu                   sync.RWMutex
	MainchainBlockNumber *big.Int
	SubchainID           *big.Int
	gasPriceLimit        *big.Int
	Dynasty              *big.Int

	chainID *big.Int
	// privateKey           *crypto.PrivateKey
	registerContractAddr ethcommon.Address
	ercContractAddr      ethcommon.Address

	// client *ethclient.Client
	// registerContract *sct.SubchainRegister
	// ercContract      *sct.SubchainERC
}

// NewMainchainMonitor creates a new MainchainMonitor
func NewMainchainMonitor(
	// privateKey *crypto.PrivateKey,
	gasPriceLimit *big.Int,
	subchainID *big.Int,
	registerContractAddr ethcommon.Address,
	ercContractAddr ethcommon.Address,
) *MainchainMonitor {
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
	// mm := &MainchainMonitor{
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

	mm := &MainchainMonitor{}
	return mm
}

// func (_MainchainMonitor *MainchainMonitor) GetValidatorSetWithBlockHeight(opts *bind.CallOpts, subchainID *big.Int) ([]ethcommon.Address, error) {
// 	// TODO: later test block height
// 	return _MainchainMonitor.registerContract.GetLegalValidators(opts, subchainID)
// }
