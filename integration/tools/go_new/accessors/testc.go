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

// TestCMetaData contains all meta data concerning the TestC contract.
var TestCMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"x\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getX\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600f57600080fd5b5061270f600055608e806100246000396000f3fe6080604052348015600f57600080fd5b506004361060325760003560e01c80630c55699c1460375780635197c7aa146051575b600080fd5b603f60005481565b60405190815260200160405180910390f35b600054603f56fea264697066735822122073e64781d41ea24e9d353f465dd77028a38e88996f27c2cc5862e64e8b89316164736f6c63430008070033",
}

// TestCABI is the input ABI used to generate the binding from.
// Deprecated: Use TestCMetaData.ABI instead.
var TestCABI = TestCMetaData.ABI

// TestCBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestCMetaData.Bin instead.
var TestCBin = TestCMetaData.Bin

// DeployTestC deploys a new Ethereum contract, binding an instance of TestC to it.
func DeployTestC(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TestC, error) {
	parsed, err := TestCMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestCBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestC{TestCCaller: TestCCaller{contract: contract}, TestCTransactor: TestCTransactor{contract: contract}, TestCFilterer: TestCFilterer{contract: contract}}, nil
}

// TestC is an auto generated Go binding around an Ethereum contract.
type TestC struct {
	TestCCaller     // Read-only binding to the contract
	TestCTransactor // Write-only binding to the contract
	TestCFilterer   // Log filterer for contract events
}

// TestCCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestCCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestCTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestCTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestCFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestCFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestCSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestCSession struct {
	Contract     *TestC            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestCCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestCCallerSession struct {
	Contract *TestCCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TestCTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestCTransactorSession struct {
	Contract     *TestCTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestCRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestCRaw struct {
	Contract *TestC // Generic contract binding to access the raw methods on
}

// TestCCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestCCallerRaw struct {
	Contract *TestCCaller // Generic read-only contract binding to access the raw methods on
}

// TestCTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestCTransactorRaw struct {
	Contract *TestCTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestC creates a new instance of TestC, bound to a specific deployed contract.
func NewTestC(address common.Address, backend bind.ContractBackend) (*TestC, error) {
	contract, err := bindTestC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestC{TestCCaller: TestCCaller{contract: contract}, TestCTransactor: TestCTransactor{contract: contract}, TestCFilterer: TestCFilterer{contract: contract}}, nil
}

// NewTestCCaller creates a new read-only instance of TestC, bound to a specific deployed contract.
func NewTestCCaller(address common.Address, caller bind.ContractCaller) (*TestCCaller, error) {
	contract, err := bindTestC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestCCaller{contract: contract}, nil
}

// NewTestCTransactor creates a new write-only instance of TestC, bound to a specific deployed contract.
func NewTestCTransactor(address common.Address, transactor bind.ContractTransactor) (*TestCTransactor, error) {
	contract, err := bindTestC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestCTransactor{contract: contract}, nil
}

// NewTestCFilterer creates a new log filterer instance of TestC, bound to a specific deployed contract.
func NewTestCFilterer(address common.Address, filterer bind.ContractFilterer) (*TestCFilterer, error) {
	contract, err := bindTestC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestCFilterer{contract: contract}, nil
}

// bindTestC binds a generic wrapper to an already deployed contract.
func bindTestC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestCABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestC *TestCRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestC.Contract.TestCCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestC *TestCRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestC.Contract.TestCTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestC *TestCRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestC.Contract.TestCTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestC *TestCCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestC *TestCTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestC *TestCTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestC.Contract.contract.Transact(opts, method, params...)
}

// X is a free data retrieval call binding the contract method 0x0c55699c.
//
// Solidity: function x() view returns(int256)
func (_TestC *TestCCaller) X(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestC.contract.Call(opts, &out, "x")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// X is a free data retrieval call binding the contract method 0x0c55699c.
//
// Solidity: function x() view returns(int256)
func (_TestC *TestCSession) X() (*big.Int, error) {
	return _TestC.Contract.X(&_TestC.CallOpts)
}

// X is a free data retrieval call binding the contract method 0x0c55699c.
//
// Solidity: function x() view returns(int256)
func (_TestC *TestCCallerSession) X() (*big.Int, error) {
	return _TestC.Contract.X(&_TestC.CallOpts)
}

// GetX is a paid mutator transaction binding the contract method 0x5197c7aa.
//
// Solidity: function getX() returns(int256)
func (_TestC *TestCTransactor) GetX(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestC.contract.Transact(opts, "getX")
}

// GetX is a paid mutator transaction binding the contract method 0x5197c7aa.
//
// Solidity: function getX() returns(int256)
func (_TestC *TestCSession) GetX() (*types.Transaction, error) {
	return _TestC.Contract.GetX(&_TestC.TransactOpts)
}

// GetX is a paid mutator transaction binding the contract method 0x5197c7aa.
//
// Solidity: function getX() returns(int256)
func (_TestC *TestCTransactorSession) GetX() (*types.Transaction, error) {
	return _TestC.Contract.GetX(&_TestC.TransactOpts)
}
