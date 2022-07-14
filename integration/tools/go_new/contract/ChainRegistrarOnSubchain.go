// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ChainRegistrarOnSubchainMetaData contains all meta data concerning the ChainRegistrarOnSubchain contract.
var ChainRegistrarOnSubchainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"isARegisteredSubchain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subchainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dynasty\",\"type\":\"uint256\"}],\"name\":\"getValidatorSet\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"shareAmounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610676806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806343b71f051461003b57806343f27e4514610064575b600080fd5b61004f610049366004610497565b50600190565b60405190151581526020015b60405180910390f35b6100776100723660046104b0565b610085565b60405161005b92919061050d565b604080518181526060818101835291829160009160208201818036833701905050905084846000805b6020811015610113578381602081106100c9576100c9610614565b1a60f81b85836100d8816105eb565b9450815181106100ea576100ea610614565b60200101906001600160f81b031916908160001a9053508061010b816105eb565b9150506100ae565b5060005b602081101561017c5782816020811061013257610132610614565b1a60f81b8583610141816105eb565b94508151811061015357610153610614565b60200101906001600160f81b031916908160001a90535080610174816105eb565b915050610117565b5060008060b56001600160a01b03168660405161019991906104d2565b600060405180830381855afa9150503d80600081146101d4576040519150601f19603f3d011682016040523d82523d6000602084013e6101d9565b606091505b5091509150816102555760405162461bcd60e51b815260206004820152603a60248201527f73746174696363616c6c20746f207468652076616c69646174726f536574207060448201527f7265636f6d70696c656420636f6e7472616374206661696c6564000000000000606482015260840160405180910390fd5b60008180602001905181019061026b91906103b2565b80519091508067ffffffffffffffff8111156102895761028961062a565b6040519080825280602002602001820160405280156102b2578160200160208202803683370190505b5099508067ffffffffffffffff8111156102ce576102ce61062a565b6040519080825280602002602001820160405280156102f7578160200160208202803683370190505b50985060005b818110156103a25782818151811061031757610317610614565b6020026020010151600001518b828151811061033557610335610614565b60200260200101906001600160a01b031690816001600160a01b03168152505082818151811061036757610367610614565b6020026020010151602001518a828151811061038557610385610614565b60209081029190910101528061039a816105eb565b9150506102fd565b5050505050505050509250929050565b600060208083850312156103c557600080fd5b825167ffffffffffffffff808211156103dd57600080fd5b818501915085601f8301126103f157600080fd5b8151818111156104035761040361062a565b610411848260051b016105ba565b8181528481019250838501600683901b8501860189101561043157600080fd5b600094505b8285101561048b57604080828b03121561044f57600080fd5b610457610591565b82516001600160a01b038116811461046e57600080fd5b815282880151888201528552600195909501949386019301610436565b50979650505050505050565b6000602082840312156104a957600080fd5b5035919050565b600080604083850312156104c357600080fd5b50508035926020909101359150565b6000825160005b818110156104f357602081860181015185830152016104d9565b81811115610502576000828501525b509190910192915050565b604080825283519082018190526000906020906060840190828701845b8281101561054f5781516001600160a01b03168452928401929084019060010161052a565b5050508381038285015284518082528583019183019060005b8181101561058457835183529284019291840191600101610568565b5090979650505050505050565b6040805190810167ffffffffffffffff811182821017156105b4576105b461062a565b60405290565b604051601f8201601f1916810167ffffffffffffffff811182821017156105e3576105e361062a565b604052919050565b600060001982141561060d57634e487b7160e01b600052601160045260246000fd5b5060010190565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fdfea2646970667358221220ffd99dde72c299218261ad7b4082257cdac3461e5957c028bd61e5b8fbdc232264736f6c63430008070033",
}

// ChainRegistrarOnSubchainABI is the input ABI used to generate the binding from.
// Deprecated: Use ChainRegistrarOnSubchainMetaData.ABI instead.
var ChainRegistrarOnSubchainABI = ChainRegistrarOnSubchainMetaData.ABI

// ChainRegistrarOnSubchainBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ChainRegistrarOnSubchainMetaData.Bin instead.
var ChainRegistrarOnSubchainBin = ChainRegistrarOnSubchainMetaData.Bin

// DeployChainRegistrarOnSubchain deploys a new Ethereum contract, binding an instance of ChainRegistrarOnSubchain to it.
func DeployChainRegistrarOnSubchain(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ChainRegistrarOnSubchain, error) {
	parsed, err := ChainRegistrarOnSubchainMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ChainRegistrarOnSubchainBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ChainRegistrarOnSubchain{ChainRegistrarOnSubchainCaller: ChainRegistrarOnSubchainCaller{contract: contract}, ChainRegistrarOnSubchainTransactor: ChainRegistrarOnSubchainTransactor{contract: contract}, ChainRegistrarOnSubchainFilterer: ChainRegistrarOnSubchainFilterer{contract: contract}}, nil
}

// ChainRegistrarOnSubchain is an auto generated Go binding around an Ethereum contract.
type ChainRegistrarOnSubchain struct {
	ChainRegistrarOnSubchainCaller     // Read-only binding to the contract
	ChainRegistrarOnSubchainTransactor // Write-only binding to the contract
	ChainRegistrarOnSubchainFilterer   // Log filterer for contract events
}

// ChainRegistrarOnSubchainCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChainRegistrarOnSubchainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainRegistrarOnSubchainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChainRegistrarOnSubchainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainRegistrarOnSubchainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChainRegistrarOnSubchainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainRegistrarOnSubchainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChainRegistrarOnSubchainSession struct {
	Contract     *ChainRegistrarOnSubchain // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ChainRegistrarOnSubchainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChainRegistrarOnSubchainCallerSession struct {
	Contract *ChainRegistrarOnSubchainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// ChainRegistrarOnSubchainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChainRegistrarOnSubchainTransactorSession struct {
	Contract     *ChainRegistrarOnSubchainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// ChainRegistrarOnSubchainRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChainRegistrarOnSubchainRaw struct {
	Contract *ChainRegistrarOnSubchain // Generic contract binding to access the raw methods on
}

// ChainRegistrarOnSubchainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChainRegistrarOnSubchainCallerRaw struct {
	Contract *ChainRegistrarOnSubchainCaller // Generic read-only contract binding to access the raw methods on
}

// ChainRegistrarOnSubchainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChainRegistrarOnSubchainTransactorRaw struct {
	Contract *ChainRegistrarOnSubchainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChainRegistrarOnSubchain creates a new instance of ChainRegistrarOnSubchain, bound to a specific deployed contract.
func NewChainRegistrarOnSubchain(address common.Address, backend bind.ContractBackend) (*ChainRegistrarOnSubchain, error) {
	contract, err := bindChainRegistrarOnSubchain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChainRegistrarOnSubchain{ChainRegistrarOnSubchainCaller: ChainRegistrarOnSubchainCaller{contract: contract}, ChainRegistrarOnSubchainTransactor: ChainRegistrarOnSubchainTransactor{contract: contract}, ChainRegistrarOnSubchainFilterer: ChainRegistrarOnSubchainFilterer{contract: contract}}, nil
}

// NewChainRegistrarOnSubchainCaller creates a new read-only instance of ChainRegistrarOnSubchain, bound to a specific deployed contract.
func NewChainRegistrarOnSubchainCaller(address common.Address, caller bind.ContractCaller) (*ChainRegistrarOnSubchainCaller, error) {
	contract, err := bindChainRegistrarOnSubchain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChainRegistrarOnSubchainCaller{contract: contract}, nil
}

// NewChainRegistrarOnSubchainTransactor creates a new write-only instance of ChainRegistrarOnSubchain, bound to a specific deployed contract.
func NewChainRegistrarOnSubchainTransactor(address common.Address, transactor bind.ContractTransactor) (*ChainRegistrarOnSubchainTransactor, error) {
	contract, err := bindChainRegistrarOnSubchain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChainRegistrarOnSubchainTransactor{contract: contract}, nil
}

// NewChainRegistrarOnSubchainFilterer creates a new log filterer instance of ChainRegistrarOnSubchain, bound to a specific deployed contract.
func NewChainRegistrarOnSubchainFilterer(address common.Address, filterer bind.ContractFilterer) (*ChainRegistrarOnSubchainFilterer, error) {
	contract, err := bindChainRegistrarOnSubchain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChainRegistrarOnSubchainFilterer{contract: contract}, nil
}

// bindChainRegistrarOnSubchain binds a generic wrapper to an already deployed contract.
func bindChainRegistrarOnSubchain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChainRegistrarOnSubchainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChainRegistrarOnSubchain *ChainRegistrarOnSubchainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChainRegistrarOnSubchain.Contract.ChainRegistrarOnSubchainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChainRegistrarOnSubchain *ChainRegistrarOnSubchainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainRegistrarOnSubchain.Contract.ChainRegistrarOnSubchainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChainRegistrarOnSubchain *ChainRegistrarOnSubchainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChainRegistrarOnSubchain.Contract.ChainRegistrarOnSubchainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChainRegistrarOnSubchain *ChainRegistrarOnSubchainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChainRegistrarOnSubchain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChainRegistrarOnSubchain *ChainRegistrarOnSubchainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainRegistrarOnSubchain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChainRegistrarOnSubchain *ChainRegistrarOnSubchainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChainRegistrarOnSubchain.Contract.contract.Transact(opts, method, params...)
}

// GetValidatorSet is a free data retrieval call binding the contract method 0x43f27e45.
//
// Solidity: function getValidatorSet(uint256 subchainID, uint256 dynasty) view returns(address[] validators, uint256[] shareAmounts)
func (_ChainRegistrarOnSubchain *ChainRegistrarOnSubchainCaller) GetValidatorSet(opts *bind.CallOpts, subchainID *big.Int, dynasty *big.Int) (struct {
	Validators   []common.Address
	ShareAmounts []*big.Int
}, error) {
	var out []interface{}
	err := _ChainRegistrarOnSubchain.contract.Call(opts, &out, "getValidatorSet", subchainID, dynasty)

	outstruct := new(struct {
		Validators   []common.Address
		ShareAmounts []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Validators = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.ShareAmounts = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetValidatorSet is a free data retrieval call binding the contract method 0x43f27e45.
//
// Solidity: function getValidatorSet(uint256 subchainID, uint256 dynasty) view returns(address[] validators, uint256[] shareAmounts)
func (_ChainRegistrarOnSubchain *ChainRegistrarOnSubchainSession) GetValidatorSet(subchainID *big.Int, dynasty *big.Int) (struct {
	Validators   []common.Address
	ShareAmounts []*big.Int
}, error) {
	return _ChainRegistrarOnSubchain.Contract.GetValidatorSet(&_ChainRegistrarOnSubchain.CallOpts, subchainID, dynasty)
}

// GetValidatorSet is a free data retrieval call binding the contract method 0x43f27e45.
//
// Solidity: function getValidatorSet(uint256 subchainID, uint256 dynasty) view returns(address[] validators, uint256[] shareAmounts)
func (_ChainRegistrarOnSubchain *ChainRegistrarOnSubchainCallerSession) GetValidatorSet(subchainID *big.Int, dynasty *big.Int) (struct {
	Validators   []common.Address
	ShareAmounts []*big.Int
}, error) {
	return _ChainRegistrarOnSubchain.Contract.GetValidatorSet(&_ChainRegistrarOnSubchain.CallOpts, subchainID, dynasty)
}

// IsARegisteredSubchain is a free data retrieval call binding the contract method 0x43b71f05.
//
// Solidity: function isARegisteredSubchain(uint256 ) pure returns(bool)
func (_ChainRegistrarOnSubchain *ChainRegistrarOnSubchainCaller) IsARegisteredSubchain(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _ChainRegistrarOnSubchain.contract.Call(opts, &out, "isARegisteredSubchain", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsARegisteredSubchain is a free data retrieval call binding the contract method 0x43b71f05.
//
// Solidity: function isARegisteredSubchain(uint256 ) pure returns(bool)
func (_ChainRegistrarOnSubchain *ChainRegistrarOnSubchainSession) IsARegisteredSubchain(arg0 *big.Int) (bool, error) {
	return _ChainRegistrarOnSubchain.Contract.IsARegisteredSubchain(&_ChainRegistrarOnSubchain.CallOpts, arg0)
}

// IsARegisteredSubchain is a free data retrieval call binding the contract method 0x43b71f05.
//
// Solidity: function isARegisteredSubchain(uint256 ) pure returns(bool)
func (_ChainRegistrarOnSubchain *ChainRegistrarOnSubchainCallerSession) IsARegisteredSubchain(arg0 *big.Int) (bool, error) {
	return _ChainRegistrarOnSubchain.Contract.IsARegisteredSubchain(&_ChainRegistrarOnSubchain.CallOpts, arg0)
}
