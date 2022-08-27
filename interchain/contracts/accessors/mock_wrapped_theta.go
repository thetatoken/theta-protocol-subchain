// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package accessors

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/thetatoken/thetasubchain/eth"
	"github.com/thetatoken/thetasubchain/eth/abi"
	"github.com/thetatoken/thetasubchain/eth/abi/bind"
	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/thetasubchain/eth/core/types"
	"github.com/thetatoken/thetasubchain/eth/event"
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

// MockWrappedThetaMetaData contains all meta data concerning the MockWrappedTheta contract.
var MockWrappedThetaMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604080518082018252600d81526c5772617070656420546865746160981b60208083019182528351808501909452600684526557544845544160d01b9084015281519192916100629160039161007e565b50805161007690600490602084019061007e565b505050610152565b82805461008a90610117565b90600052602060002090601f0160209004810192826100ac57600085556100f2565b82601f106100c557805160ff19168380011785556100f2565b828001600101855582156100f2579182015b828111156100f25782518255916020019190600101906100d7565b506100fe929150610102565b5090565b5b808211156100fe5760008155600101610103565b600181811c9082168061012b57607f821691505b6020821081141561014c57634e487b7160e01b600052602260045260246000fd5b50919050565b610994806101616000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c806340c10f191161007157806340c10f191461014157806370a082311461015657806395d89b411461017f578063a457c2d714610187578063a9059cbb1461019a578063dd62ed3e146101ad57600080fd5b806306fdde03146100b9578063095ea7b3146100d757806318160ddd146100fa57806323b872dd1461010c578063313ce5671461011f578063395093511461012e575b600080fd5b6100c16101c0565b6040516100ce91906108a8565b60405180910390f35b6100ea6100e536600461087e565b610252565b60405190151581526020016100ce565b6002545b6040519081526020016100ce565b6100ea61011a366004610842565b61026a565b604051601281526020016100ce565b6100ea61013c36600461087e565b61028e565b61015461014f36600461087e565b6102b0565b005b6100fe6101643660046107ed565b6001600160a01b031660009081526020819052604090205490565b6100c16102be565b6100ea61019536600461087e565b6102cd565b6100ea6101a836600461087e565b61034d565b6100fe6101bb36600461080f565b61035b565b6060600380546101cf90610923565b80601f01602080910402602001604051908101604052809291908181526020018280546101fb90610923565b80156102485780601f1061021d57610100808354040283529160200191610248565b820191906000526020600020905b81548152906001019060200180831161022b57829003601f168201915b5050505050905090565b600033610260818585610386565b5060019392505050565b6000336102788582856104aa565b610283858585610524565b506001949350505050565b6000336102608185856102a1838361035b565b6102ab91906108fd565b610386565b6102ba82826106f2565b5050565b6060600480546101cf90610923565b600033816102db828661035b565b9050838110156103405760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b60648201526084015b60405180910390fd5b6102838286868403610386565b600033610260818585610524565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6001600160a01b0383166103e85760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b6064820152608401610337565b6001600160a01b0382166104495760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b6064820152608401610337565b6001600160a01b0383811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b60006104b6848461035b565b9050600019811461051e57818110156105115760405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e63650000006044820152606401610337565b61051e8484848403610386565b50505050565b6001600160a01b0383166105885760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b6064820152608401610337565b6001600160a01b0382166105ea5760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b6064820152608401610337565b6001600160a01b038316600090815260208190526040902054818110156106625760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b6064820152608401610337565b6001600160a01b038085166000908152602081905260408082208585039055918516815290812080548492906106999084906108fd565b92505081905550826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516106e591815260200190565b60405180910390a361051e565b6001600160a01b0382166107485760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152606401610337565b806002600082825461075a91906108fd565b90915550506001600160a01b038216600090815260208190526040812080548392906107879084906108fd565b90915550506040518181526001600160a01b038316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b80356001600160a01b03811681146107e857600080fd5b919050565b6000602082840312156107ff57600080fd5b610808826107d1565b9392505050565b6000806040838503121561082257600080fd5b61082b836107d1565b9150610839602084016107d1565b90509250929050565b60008060006060848603121561085757600080fd5b610860846107d1565b925061086e602085016107d1565b9150604084013590509250925092565b6000806040838503121561089157600080fd5b61089a836107d1565b946020939093013593505050565b600060208083528351808285015260005b818110156108d5578581018301518582016040015282016108b9565b818111156108e7576000604083870101525b50601f01601f1916929092016040019392505050565b6000821982111561091e57634e487b7160e01b600052601160045260246000fd5b500190565b600181811c9082168061093757607f821691505b6020821081141561095857634e487b7160e01b600052602260045260246000fd5b5091905056fea26469706673582212209d7b12dbe79f231789ff403041944625f6f4f3c62105c095278b977caf83dd1364736f6c63430008070033",
}

// MockWrappedThetaABI is the input ABI used to generate the binding from.
// Deprecated: Use MockWrappedThetaMetaData.ABI instead.
var MockWrappedThetaABI = MockWrappedThetaMetaData.ABI

// MockWrappedThetaBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MockWrappedThetaMetaData.Bin instead.
var MockWrappedThetaBin = MockWrappedThetaMetaData.Bin

// DeployMockWrappedTheta deploys a new Ethereum contract, binding an instance of MockWrappedTheta to it.
func DeployMockWrappedTheta(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MockWrappedTheta, error) {
	parsed, err := MockWrappedThetaMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockWrappedThetaBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockWrappedTheta{MockWrappedThetaCaller: MockWrappedThetaCaller{contract: contract}, MockWrappedThetaTransactor: MockWrappedThetaTransactor{contract: contract}, MockWrappedThetaFilterer: MockWrappedThetaFilterer{contract: contract}}, nil
}

// MockWrappedTheta is an auto generated Go binding around an Ethereum contract.
type MockWrappedTheta struct {
	MockWrappedThetaCaller     // Read-only binding to the contract
	MockWrappedThetaTransactor // Write-only binding to the contract
	MockWrappedThetaFilterer   // Log filterer for contract events
}

// MockWrappedThetaCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockWrappedThetaCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockWrappedThetaTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockWrappedThetaTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockWrappedThetaFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockWrappedThetaFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockWrappedThetaSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockWrappedThetaSession struct {
	Contract     *MockWrappedTheta // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MockWrappedThetaCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockWrappedThetaCallerSession struct {
	Contract *MockWrappedThetaCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// MockWrappedThetaTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockWrappedThetaTransactorSession struct {
	Contract     *MockWrappedThetaTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// MockWrappedThetaRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockWrappedThetaRaw struct {
	Contract *MockWrappedTheta // Generic contract binding to access the raw methods on
}

// MockWrappedThetaCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockWrappedThetaCallerRaw struct {
	Contract *MockWrappedThetaCaller // Generic read-only contract binding to access the raw methods on
}

// MockWrappedThetaTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockWrappedThetaTransactorRaw struct {
	Contract *MockWrappedThetaTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMockWrappedTheta creates a new instance of MockWrappedTheta, bound to a specific deployed contract.
func NewMockWrappedTheta(address common.Address, backend bind.ContractBackend) (*MockWrappedTheta, error) {
	contract, err := bindMockWrappedTheta(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockWrappedTheta{MockWrappedThetaCaller: MockWrappedThetaCaller{contract: contract}, MockWrappedThetaTransactor: MockWrappedThetaTransactor{contract: contract}, MockWrappedThetaFilterer: MockWrappedThetaFilterer{contract: contract}}, nil
}

// NewMockWrappedThetaCaller creates a new read-only instance of MockWrappedTheta, bound to a specific deployed contract.
func NewMockWrappedThetaCaller(address common.Address, caller bind.ContractCaller) (*MockWrappedThetaCaller, error) {
	contract, err := bindMockWrappedTheta(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockWrappedThetaCaller{contract: contract}, nil
}

// NewMockWrappedThetaTransactor creates a new write-only instance of MockWrappedTheta, bound to a specific deployed contract.
func NewMockWrappedThetaTransactor(address common.Address, transactor bind.ContractTransactor) (*MockWrappedThetaTransactor, error) {
	contract, err := bindMockWrappedTheta(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockWrappedThetaTransactor{contract: contract}, nil
}

// NewMockWrappedThetaFilterer creates a new log filterer instance of MockWrappedTheta, bound to a specific deployed contract.
func NewMockWrappedThetaFilterer(address common.Address, filterer bind.ContractFilterer) (*MockWrappedThetaFilterer, error) {
	contract, err := bindMockWrappedTheta(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockWrappedThetaFilterer{contract: contract}, nil
}

// bindMockWrappedTheta binds a generic wrapper to an already deployed contract.
func bindMockWrappedTheta(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MockWrappedThetaABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockWrappedTheta *MockWrappedThetaRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockWrappedTheta.Contract.MockWrappedThetaCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockWrappedTheta *MockWrappedThetaRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockWrappedTheta.Contract.MockWrappedThetaTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockWrappedTheta *MockWrappedThetaRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockWrappedTheta.Contract.MockWrappedThetaTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockWrappedTheta *MockWrappedThetaCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockWrappedTheta.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockWrappedTheta *MockWrappedThetaTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockWrappedTheta.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockWrappedTheta *MockWrappedThetaTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockWrappedTheta.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MockWrappedTheta *MockWrappedThetaCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MockWrappedTheta.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MockWrappedTheta *MockWrappedThetaSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MockWrappedTheta.Contract.Allowance(&_MockWrappedTheta.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MockWrappedTheta *MockWrappedThetaCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MockWrappedTheta.Contract.Allowance(&_MockWrappedTheta.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MockWrappedTheta *MockWrappedThetaCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MockWrappedTheta.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MockWrappedTheta *MockWrappedThetaSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _MockWrappedTheta.Contract.BalanceOf(&_MockWrappedTheta.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MockWrappedTheta *MockWrappedThetaCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _MockWrappedTheta.Contract.BalanceOf(&_MockWrappedTheta.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MockWrappedTheta *MockWrappedThetaCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MockWrappedTheta.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MockWrappedTheta *MockWrappedThetaSession) Decimals() (uint8, error) {
	return _MockWrappedTheta.Contract.Decimals(&_MockWrappedTheta.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MockWrappedTheta *MockWrappedThetaCallerSession) Decimals() (uint8, error) {
	return _MockWrappedTheta.Contract.Decimals(&_MockWrappedTheta.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MockWrappedTheta *MockWrappedThetaCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MockWrappedTheta.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MockWrappedTheta *MockWrappedThetaSession) Name() (string, error) {
	return _MockWrappedTheta.Contract.Name(&_MockWrappedTheta.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MockWrappedTheta *MockWrappedThetaCallerSession) Name() (string, error) {
	return _MockWrappedTheta.Contract.Name(&_MockWrappedTheta.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MockWrappedTheta *MockWrappedThetaCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MockWrappedTheta.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MockWrappedTheta *MockWrappedThetaSession) Symbol() (string, error) {
	return _MockWrappedTheta.Contract.Symbol(&_MockWrappedTheta.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MockWrappedTheta *MockWrappedThetaCallerSession) Symbol() (string, error) {
	return _MockWrappedTheta.Contract.Symbol(&_MockWrappedTheta.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MockWrappedTheta *MockWrappedThetaCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockWrappedTheta.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MockWrappedTheta *MockWrappedThetaSession) TotalSupply() (*big.Int, error) {
	return _MockWrappedTheta.Contract.TotalSupply(&_MockWrappedTheta.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MockWrappedTheta *MockWrappedThetaCallerSession) TotalSupply() (*big.Int, error) {
	return _MockWrappedTheta.Contract.TotalSupply(&_MockWrappedTheta.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_MockWrappedTheta *MockWrappedThetaTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockWrappedTheta.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_MockWrappedTheta *MockWrappedThetaSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockWrappedTheta.Contract.Approve(&_MockWrappedTheta.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_MockWrappedTheta *MockWrappedThetaTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockWrappedTheta.Contract.Approve(&_MockWrappedTheta.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_MockWrappedTheta *MockWrappedThetaTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _MockWrappedTheta.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_MockWrappedTheta *MockWrappedThetaSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _MockWrappedTheta.Contract.DecreaseAllowance(&_MockWrappedTheta.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_MockWrappedTheta *MockWrappedThetaTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _MockWrappedTheta.Contract.DecreaseAllowance(&_MockWrappedTheta.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_MockWrappedTheta *MockWrappedThetaTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _MockWrappedTheta.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_MockWrappedTheta *MockWrappedThetaSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _MockWrappedTheta.Contract.IncreaseAllowance(&_MockWrappedTheta.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_MockWrappedTheta *MockWrappedThetaTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _MockWrappedTheta.Contract.IncreaseAllowance(&_MockWrappedTheta.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_MockWrappedTheta *MockWrappedThetaTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockWrappedTheta.contract.Transact(opts, "mint", account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_MockWrappedTheta *MockWrappedThetaSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockWrappedTheta.Contract.Mint(&_MockWrappedTheta.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_MockWrappedTheta *MockWrappedThetaTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockWrappedTheta.Contract.Mint(&_MockWrappedTheta.TransactOpts, account, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_MockWrappedTheta *MockWrappedThetaTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockWrappedTheta.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_MockWrappedTheta *MockWrappedThetaSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockWrappedTheta.Contract.Transfer(&_MockWrappedTheta.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_MockWrappedTheta *MockWrappedThetaTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockWrappedTheta.Contract.Transfer(&_MockWrappedTheta.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_MockWrappedTheta *MockWrappedThetaTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockWrappedTheta.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_MockWrappedTheta *MockWrappedThetaSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockWrappedTheta.Contract.TransferFrom(&_MockWrappedTheta.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_MockWrappedTheta *MockWrappedThetaTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockWrappedTheta.Contract.TransferFrom(&_MockWrappedTheta.TransactOpts, from, to, amount)
}

// MockWrappedThetaApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MockWrappedTheta contract.
type MockWrappedThetaApprovalIterator struct {
	Event *MockWrappedThetaApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MockWrappedThetaApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockWrappedThetaApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MockWrappedThetaApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MockWrappedThetaApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockWrappedThetaApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockWrappedThetaApproval represents a Approval event raised by the MockWrappedTheta contract.
type MockWrappedThetaApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MockWrappedTheta *MockWrappedThetaFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MockWrappedThetaApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MockWrappedTheta.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MockWrappedThetaApprovalIterator{contract: _MockWrappedTheta.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MockWrappedTheta *MockWrappedThetaFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MockWrappedThetaApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MockWrappedTheta.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockWrappedThetaApproval)
				if err := _MockWrappedTheta.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MockWrappedTheta *MockWrappedThetaFilterer) ParseApproval(log types.Log) (*MockWrappedThetaApproval, error) {
	event := new(MockWrappedThetaApproval)
	if err := _MockWrappedTheta.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockWrappedThetaTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MockWrappedTheta contract.
type MockWrappedThetaTransferIterator struct {
	Event *MockWrappedThetaTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MockWrappedThetaTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockWrappedThetaTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MockWrappedThetaTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MockWrappedThetaTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockWrappedThetaTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockWrappedThetaTransfer represents a Transfer event raised by the MockWrappedTheta contract.
type MockWrappedThetaTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MockWrappedTheta *MockWrappedThetaFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MockWrappedThetaTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MockWrappedTheta.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MockWrappedThetaTransferIterator{contract: _MockWrappedTheta.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MockWrappedTheta *MockWrappedThetaFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MockWrappedThetaTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MockWrappedTheta.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockWrappedThetaTransfer)
				if err := _MockWrappedTheta.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MockWrappedTheta *MockWrappedThetaFilterer) ParseTransfer(log types.Log) (*MockWrappedThetaTransfer, error) {
	event := new(MockWrappedThetaTransfer)
	if err := _MockWrappedTheta.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
