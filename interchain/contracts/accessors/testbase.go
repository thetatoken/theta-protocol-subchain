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

// TestBaseMetaData contains all meta data concerning the TestBase contract.
var TestBaseMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"x_\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"x\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506040516100db3803806100db83398101604081905261002f91610037565b600055610050565b60006020828403121561004957600080fd5b5051919050565b607d8061005e6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c80630c55699c14602d575b600080fd5b603560005481565b60405190815260200160405180910390f3fea26469706673582212205e7350457f69c0cb903d2af94b36f5652165155dcffa4fdab624ef7e3ddfaf4d64736f6c63430008070033",
}

// TestBaseABI is the input ABI used to generate the binding from.
// Deprecated: Use TestBaseMetaData.ABI instead.
var TestBaseABI = TestBaseMetaData.ABI

// TestBaseBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestBaseMetaData.Bin instead.
var TestBaseBin = TestBaseMetaData.Bin

// DeployTestBase deploys a new Ethereum contract, binding an instance of TestBase to it.
func DeployTestBase(auth *bind.TransactOpts, backend bind.ContractBackend, x_ *big.Int) (common.Address, *types.Transaction, *TestBase, error) {
	parsed, err := TestBaseMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestBaseBin), backend, x_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestBase{TestBaseCaller: TestBaseCaller{contract: contract}, TestBaseTransactor: TestBaseTransactor{contract: contract}, TestBaseFilterer: TestBaseFilterer{contract: contract}}, nil
}

// TestBase is an auto generated Go binding around an Ethereum contract.
type TestBase struct {
	TestBaseCaller     // Read-only binding to the contract
	TestBaseTransactor // Write-only binding to the contract
	TestBaseFilterer   // Log filterer for contract events
}

// TestBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestBaseSession struct {
	Contract     *TestBase         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestBaseCallerSession struct {
	Contract *TestBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// TestBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestBaseTransactorSession struct {
	Contract     *TestBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TestBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestBaseRaw struct {
	Contract *TestBase // Generic contract binding to access the raw methods on
}

// TestBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestBaseCallerRaw struct {
	Contract *TestBaseCaller // Generic read-only contract binding to access the raw methods on
}

// TestBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestBaseTransactorRaw struct {
	Contract *TestBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestBase creates a new instance of TestBase, bound to a specific deployed contract.
func NewTestBase(address common.Address, backend bind.ContractBackend) (*TestBase, error) {
	contract, err := bindTestBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestBase{TestBaseCaller: TestBaseCaller{contract: contract}, TestBaseTransactor: TestBaseTransactor{contract: contract}, TestBaseFilterer: TestBaseFilterer{contract: contract}}, nil
}

// NewTestBaseCaller creates a new read-only instance of TestBase, bound to a specific deployed contract.
func NewTestBaseCaller(address common.Address, caller bind.ContractCaller) (*TestBaseCaller, error) {
	contract, err := bindTestBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestBaseCaller{contract: contract}, nil
}

// NewTestBaseTransactor creates a new write-only instance of TestBase, bound to a specific deployed contract.
func NewTestBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*TestBaseTransactor, error) {
	contract, err := bindTestBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestBaseTransactor{contract: contract}, nil
}

// NewTestBaseFilterer creates a new log filterer instance of TestBase, bound to a specific deployed contract.
func NewTestBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*TestBaseFilterer, error) {
	contract, err := bindTestBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestBaseFilterer{contract: contract}, nil
}

// bindTestBase binds a generic wrapper to an already deployed contract.
func bindTestBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestBaseABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestBase *TestBaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestBase.Contract.TestBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestBase *TestBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestBase.Contract.TestBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestBase *TestBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestBase.Contract.TestBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestBase *TestBaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestBase *TestBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestBase *TestBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestBase.Contract.contract.Transact(opts, method, params...)
}

// X is a free data retrieval call binding the contract method 0x0c55699c.
//
// Solidity: function x() view returns(int256)
func (_TestBase *TestBaseCaller) X(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestBase.contract.Call(opts, &out, "x")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// X is a free data retrieval call binding the contract method 0x0c55699c.
//
// Solidity: function x() view returns(int256)
func (_TestBase *TestBaseSession) X() (*big.Int, error) {
	return _TestBase.Contract.X(&_TestBase.CallOpts)
}

// X is a free data retrieval call binding the contract method 0x0c55699c.
//
// Solidity: function x() view returns(int256)
func (_TestBase *TestBaseCallerSession) X() (*big.Int, error) {
	return _TestBase.Contract.X(&_TestBase.CallOpts)
}
