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

// TestAMetaData contains all meta data concerning the TestA contract.
var TestAMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"x\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getX\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600f57600080fd5b5061270f600055608e806100246000396000f3fe6080604052348015600f57600080fd5b506004361060325760003560e01c80630c55699c1460375780635197c7aa146051575b600080fd5b603f60005481565b60405190815260200160405180910390f35b600054603f56fea26469706673582212209af9dfbd6753159278a709fdb368ff61579a19fda29ce77e4b92dbb44fdbef8864736f6c63430008070033",
}

// TestAABI is the input ABI used to generate the binding from.
// Deprecated: Use TestAMetaData.ABI instead.
var TestAABI = TestAMetaData.ABI

// TestABin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestAMetaData.Bin instead.
var TestABin = TestAMetaData.Bin

// DeployTestA deploys a new Ethereum contract, binding an instance of TestA to it.
func DeployTestA(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TestA, error) {
	parsed, err := TestAMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestABin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestA{TestACaller: TestACaller{contract: contract}, TestATransactor: TestATransactor{contract: contract}, TestAFilterer: TestAFilterer{contract: contract}}, nil
}

// TestA is an auto generated Go binding around an Ethereum contract.
type TestA struct {
	TestACaller     // Read-only binding to the contract
	TestATransactor // Write-only binding to the contract
	TestAFilterer   // Log filterer for contract events
}

// TestACaller is an auto generated read-only Go binding around an Ethereum contract.
type TestACaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestATransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestATransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestAFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestAFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestASession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestASession struct {
	Contract     *TestA            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestACallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestACallerSession struct {
	Contract *TestACaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TestATransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestATransactorSession struct {
	Contract     *TestATransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestARaw is an auto generated low-level Go binding around an Ethereum contract.
type TestARaw struct {
	Contract *TestA // Generic contract binding to access the raw methods on
}

// TestACallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestACallerRaw struct {
	Contract *TestACaller // Generic read-only contract binding to access the raw methods on
}

// TestATransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestATransactorRaw struct {
	Contract *TestATransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestA creates a new instance of TestA, bound to a specific deployed contract.
func NewTestA(address common.Address, backend bind.ContractBackend) (*TestA, error) {
	contract, err := bindTestA(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestA{TestACaller: TestACaller{contract: contract}, TestATransactor: TestATransactor{contract: contract}, TestAFilterer: TestAFilterer{contract: contract}}, nil
}

// NewTestACaller creates a new read-only instance of TestA, bound to a specific deployed contract.
func NewTestACaller(address common.Address, caller bind.ContractCaller) (*TestACaller, error) {
	contract, err := bindTestA(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestACaller{contract: contract}, nil
}

// NewTestATransactor creates a new write-only instance of TestA, bound to a specific deployed contract.
func NewTestATransactor(address common.Address, transactor bind.ContractTransactor) (*TestATransactor, error) {
	contract, err := bindTestA(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestATransactor{contract: contract}, nil
}

// NewTestAFilterer creates a new log filterer instance of TestA, bound to a specific deployed contract.
func NewTestAFilterer(address common.Address, filterer bind.ContractFilterer) (*TestAFilterer, error) {
	contract, err := bindTestA(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestAFilterer{contract: contract}, nil
}

// bindTestA binds a generic wrapper to an already deployed contract.
func bindTestA(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestAABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestA *TestARaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestA.Contract.TestACaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestA *TestARaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestA.Contract.TestATransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestA *TestARaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestA.Contract.TestATransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestA *TestACallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestA.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestA *TestATransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestA.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestA *TestATransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestA.Contract.contract.Transact(opts, method, params...)
}

// X is a free data retrieval call binding the contract method 0x0c55699c.
//
// Solidity: function x() view returns(int256)
func (_TestA *TestACaller) X(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestA.contract.Call(opts, &out, "x")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// X is a free data retrieval call binding the contract method 0x0c55699c.
//
// Solidity: function x() view returns(int256)
func (_TestA *TestASession) X() (*big.Int, error) {
	return _TestA.Contract.X(&_TestA.CallOpts)
}

// X is a free data retrieval call binding the contract method 0x0c55699c.
//
// Solidity: function x() view returns(int256)
func (_TestA *TestACallerSession) X() (*big.Int, error) {
	return _TestA.Contract.X(&_TestA.CallOpts)
}

// GetX is a paid mutator transaction binding the contract method 0x5197c7aa.
//
// Solidity: function getX() returns(int256)
func (_TestA *TestATransactor) GetX(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestA.contract.Transact(opts, "getX")
}

// GetX is a paid mutator transaction binding the contract method 0x5197c7aa.
//
// Solidity: function getX() returns(int256)
func (_TestA *TestASession) GetX() (*types.Transaction, error) {
	return _TestA.Contract.GetX(&_TestA.TransactOpts)
}

// GetX is a paid mutator transaction binding the contract method 0x5197c7aa.
//
// Solidity: function getX() returns(int256)
func (_TestA *TestATransactorSession) GetX() (*types.Transaction, error) {
	return _TestA.Contract.GetX(&_TestA.TransactOpts)
}
