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

// TestBMetaData contains all meta data concerning the TestB contract.
var TestBMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"x\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getX\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600f57600080fd5b50608e8061001e6000396000f3fe6080604052348015600f57600080fd5b506004361060325760003560e01c80630c55699c1460375780635197c7aa146051575b600080fd5b603f61270f81565b60405190815260200160405180910390f35b61270f603f56fea26469706673582212201d824df8af0267e5734b5208d0398782bdfb989cf00e435b480f1f9207531c1864736f6c63430008070033",
}

// TestBABI is the input ABI used to generate the binding from.
// Deprecated: Use TestBMetaData.ABI instead.
var TestBABI = TestBMetaData.ABI

// TestBBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestBMetaData.Bin instead.
var TestBBin = TestBMetaData.Bin

// DeployTestB deploys a new Ethereum contract, binding an instance of TestB to it.
func DeployTestB(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TestB, error) {
	parsed, err := TestBMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestBBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestB{TestBCaller: TestBCaller{contract: contract}, TestBTransactor: TestBTransactor{contract: contract}, TestBFilterer: TestBFilterer{contract: contract}}, nil
}

// TestB is an auto generated Go binding around an Ethereum contract.
type TestB struct {
	TestBCaller     // Read-only binding to the contract
	TestBTransactor // Write-only binding to the contract
	TestBFilterer   // Log filterer for contract events
}

// TestBCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestBCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestBTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestBTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestBFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestBFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestBSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestBSession struct {
	Contract     *TestB            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestBCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestBCallerSession struct {
	Contract *TestBCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TestBTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestBTransactorSession struct {
	Contract     *TestBTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestBRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestBRaw struct {
	Contract *TestB // Generic contract binding to access the raw methods on
}

// TestBCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestBCallerRaw struct {
	Contract *TestBCaller // Generic read-only contract binding to access the raw methods on
}

// TestBTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestBTransactorRaw struct {
	Contract *TestBTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestB creates a new instance of TestB, bound to a specific deployed contract.
func NewTestB(address common.Address, backend bind.ContractBackend) (*TestB, error) {
	contract, err := bindTestB(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestB{TestBCaller: TestBCaller{contract: contract}, TestBTransactor: TestBTransactor{contract: contract}, TestBFilterer: TestBFilterer{contract: contract}}, nil
}

// NewTestBCaller creates a new read-only instance of TestB, bound to a specific deployed contract.
func NewTestBCaller(address common.Address, caller bind.ContractCaller) (*TestBCaller, error) {
	contract, err := bindTestB(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestBCaller{contract: contract}, nil
}

// NewTestBTransactor creates a new write-only instance of TestB, bound to a specific deployed contract.
func NewTestBTransactor(address common.Address, transactor bind.ContractTransactor) (*TestBTransactor, error) {
	contract, err := bindTestB(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestBTransactor{contract: contract}, nil
}

// NewTestBFilterer creates a new log filterer instance of TestB, bound to a specific deployed contract.
func NewTestBFilterer(address common.Address, filterer bind.ContractFilterer) (*TestBFilterer, error) {
	contract, err := bindTestB(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestBFilterer{contract: contract}, nil
}

// bindTestB binds a generic wrapper to an already deployed contract.
func bindTestB(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestBABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestB *TestBRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestB.Contract.TestBCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestB *TestBRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestB.Contract.TestBTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestB *TestBRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestB.Contract.TestBTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestB *TestBCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestB.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestB *TestBTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestB.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestB *TestBTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestB.Contract.contract.Transact(opts, method, params...)
}

// X is a free data retrieval call binding the contract method 0x0c55699c.
//
// Solidity: function x() view returns(int256)
func (_TestB *TestBCaller) X(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestB.contract.Call(opts, &out, "x")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// X is a free data retrieval call binding the contract method 0x0c55699c.
//
// Solidity: function x() view returns(int256)
func (_TestB *TestBSession) X() (*big.Int, error) {
	return _TestB.Contract.X(&_TestB.CallOpts)
}

// X is a free data retrieval call binding the contract method 0x0c55699c.
//
// Solidity: function x() view returns(int256)
func (_TestB *TestBCallerSession) X() (*big.Int, error) {
	return _TestB.Contract.X(&_TestB.CallOpts)
}

// GetX is a paid mutator transaction binding the contract method 0x5197c7aa.
//
// Solidity: function getX() returns(int256)
func (_TestB *TestBTransactor) GetX(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestB.contract.Transact(opts, "getX")
}

// GetX is a paid mutator transaction binding the contract method 0x5197c7aa.
//
// Solidity: function getX() returns(int256)
func (_TestB *TestBSession) GetX() (*types.Transaction, error) {
	return _TestB.Contract.GetX(&_TestB.TransactOpts)
}

// GetX is a paid mutator transaction binding the contract method 0x5197c7aa.
//
// Solidity: function getX() returns(int256)
func (_TestB *TestBTransactorSession) GetX() (*types.Transaction, error) {
	return _TestB.Contract.GetX(&_TestB.TransactOpts)
}
