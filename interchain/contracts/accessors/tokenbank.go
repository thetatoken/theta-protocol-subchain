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

// TokenBankMetaData contains all meta data concerning the TokenBank contract.
var TokenBankMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"mainchainID_\",\"type\":\"uint256\"},{\"internalType\":\"contractChainRegistrar\",\"name\":\"chainRegistrar_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FailedToSendTFuel\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"mainchainID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenLockNonceMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"tokenLockVotingRecords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dynasty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumlatedShares\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenUnlockNonceMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"voucherBurnNonceMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"voucherBurnVotingRecords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dynasty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumlatedShares\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"voucherMintNonceMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"getMaxProcessedTokenLockNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"eventNonce\",\"type\":\"uint256\"}],\"name\":\"getTokenLockEventHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"getMaxProcessedVoucherBurnNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"eventNonce\",\"type\":\"uint256\"}],\"name\":\"getVoucherBurnEventHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subchainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dynasty\",\"type\":\"uint256\"}],\"name\":\"getAdjustedValidatorSet\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"shareAmounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161092338038061092383398101604081905261002f9161005e565b6001600081905591909155600280546001600160a01b0319166001600160a01b0390921691909117905561009b565b6000806040838503121561007157600080fd5b825160208401519092506001600160a01b038116811461009057600080fd5b809150509250929050565b610879806100aa6000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c8063aa861c1511610071578063aa861c15146101a7578063ca207569146101c8578063ccf187c7146101e8578063dd17eb6d14610208578063f95627ac14610233578063feaff0521461025357600080fd5b8063073b9502146100b95780631eb78737146100d55780636ac739b91461011c578063740cb7f814610147578063766f8fb0146101675780638883931e14610187575b600080fd5b6100c260015481565b6040519081526020015b60405180910390f35b6101076100e33660046106f2565b600b6020908152600092835260408084209091529082529020805460029091015482565b604080519283526020830191909152016100cc565b6100c261012a3660046106f2565b600091825260086020908152604080842092845291905290205490565b6100c261015536600461068b565b60066020526000908152604090205481565b6100c261017536600461068b565b6000908152600a602052604090205490565b6100c261019536600461068b565b60036020526000908152604090205481565b6101ba6101b53660046106f2565b610285565b6040516100cc929190610714565b6100c26101d636600461068b565b60056020526000908152604090205481565b6100c26101f636600461068b565b60046020526000908152604090205481565b6100c26102163660046106f2565b600091825260076020908152604080842092845291905290205490565b6100c261024136600461068b565b60009081526009602052604090205490565b6101076102613660046106f2565b600c6020908152600092835260408084209091529082529020805460029091015482565b6060806000600154461490506000600260009054906101000a90046001600160a01b03166001600160a01b031663a7464b126040518163ffffffff1660e01b815260040160206040518083038186803b1580156102e157600080fd5b505afa1580156102f5573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061031991906106a4565b600254604051632f2c13b560e01b81526004810189905291925060009182916001600160a01b031690632f2c13b590602401604080518083038186803b15801561036257600080fd5b505afa158015610376573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061039a91906106bd565b9150915060008180156103b657506103b28885610522565b8310155b80156103d557506103d2846103cc8a6001610535565b90610522565b83105b90508480156103e15750805b1561048a5760006103f3896001610535565b6002546040516343f27e4560e01b8152600481018d9052602481018390529192506001600160a01b0316906343f27e459060440160006040518083038186803b15801561043f57600080fd5b505afa158015610453573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261047b91908101906105b3565b9750975050505050505061051b565b6002546040516343f27e4560e01b8152600481018b9052602481018a90526001600160a01b03909116906343f27e459060440160006040518083038186803b1580156104d557600080fd5b505afa1580156104e9573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261051191908101906105b3565b9650965050505050505b9250929050565b600061052e82846107f8565b9392505050565b600061052e82846107e0565b600082601f83011261055257600080fd5b81516020610567610562836107bc565b61078b565b80838252828201915082860187848660051b890101111561058757600080fd5b60005b858110156105a65781518452928401929084019060010161058a565b5090979650505050505050565b600080604083850312156105c657600080fd5b825167ffffffffffffffff808211156105de57600080fd5b818501915085601f8301126105f257600080fd5b81516020610602610562836107bc565b8083825282820191508286018a848660051b890101111561062257600080fd5b600096505b8487101561065a5780516001600160a01b038116811461064657600080fd5b835260019690960195918301918301610627565b509188015191965090935050508082111561067457600080fd5b5061068185828601610541565b9150509250929050565b60006020828403121561069d57600080fd5b5035919050565b6000602082840312156106b657600080fd5b5051919050565b600080604083850312156106d057600080fd5b82519150602083015180151581146106e757600080fd5b809150509250929050565b6000806040838503121561070557600080fd5b50508035926020909101359150565b604080825283519082018190526000906020906060840190828701845b828110156107565781516001600160a01b031684529284019290840190600101610731565b5050508381038285015284518082528583019183019060005b818110156105a65783518352928401929184019160010161076f565b604051601f8201601f1916810167ffffffffffffffff811182821017156107b4576107b461082d565b604052919050565b600067ffffffffffffffff8211156107d6576107d661082d565b5060051b60200190565b600082198211156107f3576107f3610817565b500190565b600081600019048311821515161561081257610812610817565b500290565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fdfea2646970667358221220a1c547d2927abcaa30b3634b7a1f23db99435e471eeef155fa439e909336d1ba64736f6c63430008070033",
}

// TokenBankABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenBankMetaData.ABI instead.
var TokenBankABI = TokenBankMetaData.ABI

// TokenBankBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TokenBankMetaData.Bin instead.
var TokenBankBin = TokenBankMetaData.Bin

// DeployTokenBank deploys a new Ethereum contract, binding an instance of TokenBank to it.
func DeployTokenBank(auth *bind.TransactOpts, backend bind.ContractBackend, mainchainID_ *big.Int, chainRegistrar_ common.Address) (common.Address, *types.Transaction, *TokenBank, error) {
	parsed, err := TokenBankMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TokenBankBin), backend, mainchainID_, chainRegistrar_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenBank{TokenBankCaller: TokenBankCaller{contract: contract}, TokenBankTransactor: TokenBankTransactor{contract: contract}, TokenBankFilterer: TokenBankFilterer{contract: contract}}, nil
}

// TokenBank is an auto generated Go binding around an Ethereum contract.
type TokenBank struct {
	TokenBankCaller     // Read-only binding to the contract
	TokenBankTransactor // Write-only binding to the contract
	TokenBankFilterer   // Log filterer for contract events
}

// TokenBankCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenBankCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenBankTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenBankTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenBankFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenBankFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenBankSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenBankSession struct {
	Contract     *TokenBank        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenBankCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenBankCallerSession struct {
	Contract *TokenBankCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// TokenBankTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenBankTransactorSession struct {
	Contract     *TokenBankTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TokenBankRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenBankRaw struct {
	Contract *TokenBank // Generic contract binding to access the raw methods on
}

// TokenBankCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenBankCallerRaw struct {
	Contract *TokenBankCaller // Generic read-only contract binding to access the raw methods on
}

// TokenBankTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenBankTransactorRaw struct {
	Contract *TokenBankTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenBank creates a new instance of TokenBank, bound to a specific deployed contract.
func NewTokenBank(address common.Address, backend bind.ContractBackend) (*TokenBank, error) {
	contract, err := bindTokenBank(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenBank{TokenBankCaller: TokenBankCaller{contract: contract}, TokenBankTransactor: TokenBankTransactor{contract: contract}, TokenBankFilterer: TokenBankFilterer{contract: contract}}, nil
}

// NewTokenBankCaller creates a new read-only instance of TokenBank, bound to a specific deployed contract.
func NewTokenBankCaller(address common.Address, caller bind.ContractCaller) (*TokenBankCaller, error) {
	contract, err := bindTokenBank(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenBankCaller{contract: contract}, nil
}

// NewTokenBankTransactor creates a new write-only instance of TokenBank, bound to a specific deployed contract.
func NewTokenBankTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenBankTransactor, error) {
	contract, err := bindTokenBank(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenBankTransactor{contract: contract}, nil
}

// NewTokenBankFilterer creates a new log filterer instance of TokenBank, bound to a specific deployed contract.
func NewTokenBankFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenBankFilterer, error) {
	contract, err := bindTokenBank(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenBankFilterer{contract: contract}, nil
}

// bindTokenBank binds a generic wrapper to an already deployed contract.
func bindTokenBank(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenBankABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenBank *TokenBankRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenBank.Contract.TokenBankCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenBank *TokenBankRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenBank.Contract.TokenBankTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenBank *TokenBankRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenBank.Contract.TokenBankTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenBank *TokenBankCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenBank.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenBank *TokenBankTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenBank.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenBank *TokenBankTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenBank.Contract.contract.Transact(opts, method, params...)
}

// GetAdjustedValidatorSet is a free data retrieval call binding the contract method 0xaa861c15.
//
// Solidity: function getAdjustedValidatorSet(uint256 subchainID, uint256 dynasty) view returns(address[] validators, uint256[] shareAmounts)
func (_TokenBank *TokenBankCaller) GetAdjustedValidatorSet(opts *bind.CallOpts, subchainID *big.Int, dynasty *big.Int) (struct {
	Validators   []common.Address
	ShareAmounts []*big.Int
}, error) {
	var out []interface{}
	err := _TokenBank.contract.Call(opts, &out, "getAdjustedValidatorSet", subchainID, dynasty)

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

// GetAdjustedValidatorSet is a free data retrieval call binding the contract method 0xaa861c15.
//
// Solidity: function getAdjustedValidatorSet(uint256 subchainID, uint256 dynasty) view returns(address[] validators, uint256[] shareAmounts)
func (_TokenBank *TokenBankSession) GetAdjustedValidatorSet(subchainID *big.Int, dynasty *big.Int) (struct {
	Validators   []common.Address
	ShareAmounts []*big.Int
}, error) {
	return _TokenBank.Contract.GetAdjustedValidatorSet(&_TokenBank.CallOpts, subchainID, dynasty)
}

// GetAdjustedValidatorSet is a free data retrieval call binding the contract method 0xaa861c15.
//
// Solidity: function getAdjustedValidatorSet(uint256 subchainID, uint256 dynasty) view returns(address[] validators, uint256[] shareAmounts)
func (_TokenBank *TokenBankCallerSession) GetAdjustedValidatorSet(subchainID *big.Int, dynasty *big.Int) (struct {
	Validators   []common.Address
	ShareAmounts []*big.Int
}, error) {
	return _TokenBank.Contract.GetAdjustedValidatorSet(&_TokenBank.CallOpts, subchainID, dynasty)
}

// GetMaxProcessedTokenLockNonce is a free data retrieval call binding the contract method 0xf95627ac.
//
// Solidity: function getMaxProcessedTokenLockNonce(uint256 chainID) view returns(uint256)
func (_TokenBank *TokenBankCaller) GetMaxProcessedTokenLockNonce(opts *bind.CallOpts, chainID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TokenBank.contract.Call(opts, &out, "getMaxProcessedTokenLockNonce", chainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxProcessedTokenLockNonce is a free data retrieval call binding the contract method 0xf95627ac.
//
// Solidity: function getMaxProcessedTokenLockNonce(uint256 chainID) view returns(uint256)
func (_TokenBank *TokenBankSession) GetMaxProcessedTokenLockNonce(chainID *big.Int) (*big.Int, error) {
	return _TokenBank.Contract.GetMaxProcessedTokenLockNonce(&_TokenBank.CallOpts, chainID)
}

// GetMaxProcessedTokenLockNonce is a free data retrieval call binding the contract method 0xf95627ac.
//
// Solidity: function getMaxProcessedTokenLockNonce(uint256 chainID) view returns(uint256)
func (_TokenBank *TokenBankCallerSession) GetMaxProcessedTokenLockNonce(chainID *big.Int) (*big.Int, error) {
	return _TokenBank.Contract.GetMaxProcessedTokenLockNonce(&_TokenBank.CallOpts, chainID)
}

// GetMaxProcessedVoucherBurnNonce is a free data retrieval call binding the contract method 0x766f8fb0.
//
// Solidity: function getMaxProcessedVoucherBurnNonce(uint256 chainID) view returns(uint256)
func (_TokenBank *TokenBankCaller) GetMaxProcessedVoucherBurnNonce(opts *bind.CallOpts, chainID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TokenBank.contract.Call(opts, &out, "getMaxProcessedVoucherBurnNonce", chainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxProcessedVoucherBurnNonce is a free data retrieval call binding the contract method 0x766f8fb0.
//
// Solidity: function getMaxProcessedVoucherBurnNonce(uint256 chainID) view returns(uint256)
func (_TokenBank *TokenBankSession) GetMaxProcessedVoucherBurnNonce(chainID *big.Int) (*big.Int, error) {
	return _TokenBank.Contract.GetMaxProcessedVoucherBurnNonce(&_TokenBank.CallOpts, chainID)
}

// GetMaxProcessedVoucherBurnNonce is a free data retrieval call binding the contract method 0x766f8fb0.
//
// Solidity: function getMaxProcessedVoucherBurnNonce(uint256 chainID) view returns(uint256)
func (_TokenBank *TokenBankCallerSession) GetMaxProcessedVoucherBurnNonce(chainID *big.Int) (*big.Int, error) {
	return _TokenBank.Contract.GetMaxProcessedVoucherBurnNonce(&_TokenBank.CallOpts, chainID)
}

// GetTokenLockEventHeight is a free data retrieval call binding the contract method 0xdd17eb6d.
//
// Solidity: function getTokenLockEventHeight(uint256 chainID, uint256 eventNonce) view returns(uint256)
func (_TokenBank *TokenBankCaller) GetTokenLockEventHeight(opts *bind.CallOpts, chainID *big.Int, eventNonce *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TokenBank.contract.Call(opts, &out, "getTokenLockEventHeight", chainID, eventNonce)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenLockEventHeight is a free data retrieval call binding the contract method 0xdd17eb6d.
//
// Solidity: function getTokenLockEventHeight(uint256 chainID, uint256 eventNonce) view returns(uint256)
func (_TokenBank *TokenBankSession) GetTokenLockEventHeight(chainID *big.Int, eventNonce *big.Int) (*big.Int, error) {
	return _TokenBank.Contract.GetTokenLockEventHeight(&_TokenBank.CallOpts, chainID, eventNonce)
}

// GetTokenLockEventHeight is a free data retrieval call binding the contract method 0xdd17eb6d.
//
// Solidity: function getTokenLockEventHeight(uint256 chainID, uint256 eventNonce) view returns(uint256)
func (_TokenBank *TokenBankCallerSession) GetTokenLockEventHeight(chainID *big.Int, eventNonce *big.Int) (*big.Int, error) {
	return _TokenBank.Contract.GetTokenLockEventHeight(&_TokenBank.CallOpts, chainID, eventNonce)
}

// GetVoucherBurnEventHeight is a free data retrieval call binding the contract method 0x6ac739b9.
//
// Solidity: function getVoucherBurnEventHeight(uint256 chainID, uint256 eventNonce) view returns(uint256)
func (_TokenBank *TokenBankCaller) GetVoucherBurnEventHeight(opts *bind.CallOpts, chainID *big.Int, eventNonce *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TokenBank.contract.Call(opts, &out, "getVoucherBurnEventHeight", chainID, eventNonce)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVoucherBurnEventHeight is a free data retrieval call binding the contract method 0x6ac739b9.
//
// Solidity: function getVoucherBurnEventHeight(uint256 chainID, uint256 eventNonce) view returns(uint256)
func (_TokenBank *TokenBankSession) GetVoucherBurnEventHeight(chainID *big.Int, eventNonce *big.Int) (*big.Int, error) {
	return _TokenBank.Contract.GetVoucherBurnEventHeight(&_TokenBank.CallOpts, chainID, eventNonce)
}

// GetVoucherBurnEventHeight is a free data retrieval call binding the contract method 0x6ac739b9.
//
// Solidity: function getVoucherBurnEventHeight(uint256 chainID, uint256 eventNonce) view returns(uint256)
func (_TokenBank *TokenBankCallerSession) GetVoucherBurnEventHeight(chainID *big.Int, eventNonce *big.Int) (*big.Int, error) {
	return _TokenBank.Contract.GetVoucherBurnEventHeight(&_TokenBank.CallOpts, chainID, eventNonce)
}

// MainchainID is a free data retrieval call binding the contract method 0x073b9502.
//
// Solidity: function mainchainID() view returns(uint256)
func (_TokenBank *TokenBankCaller) MainchainID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenBank.contract.Call(opts, &out, "mainchainID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MainchainID is a free data retrieval call binding the contract method 0x073b9502.
//
// Solidity: function mainchainID() view returns(uint256)
func (_TokenBank *TokenBankSession) MainchainID() (*big.Int, error) {
	return _TokenBank.Contract.MainchainID(&_TokenBank.CallOpts)
}

// MainchainID is a free data retrieval call binding the contract method 0x073b9502.
//
// Solidity: function mainchainID() view returns(uint256)
func (_TokenBank *TokenBankCallerSession) MainchainID() (*big.Int, error) {
	return _TokenBank.Contract.MainchainID(&_TokenBank.CallOpts)
}

// TokenLockNonceMap is a free data retrieval call binding the contract method 0x8883931e.
//
// Solidity: function tokenLockNonceMap(uint256 ) view returns(uint256)
func (_TokenBank *TokenBankCaller) TokenLockNonceMap(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TokenBank.contract.Call(opts, &out, "tokenLockNonceMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenLockNonceMap is a free data retrieval call binding the contract method 0x8883931e.
//
// Solidity: function tokenLockNonceMap(uint256 ) view returns(uint256)
func (_TokenBank *TokenBankSession) TokenLockNonceMap(arg0 *big.Int) (*big.Int, error) {
	return _TokenBank.Contract.TokenLockNonceMap(&_TokenBank.CallOpts, arg0)
}

// TokenLockNonceMap is a free data retrieval call binding the contract method 0x8883931e.
//
// Solidity: function tokenLockNonceMap(uint256 ) view returns(uint256)
func (_TokenBank *TokenBankCallerSession) TokenLockNonceMap(arg0 *big.Int) (*big.Int, error) {
	return _TokenBank.Contract.TokenLockNonceMap(&_TokenBank.CallOpts, arg0)
}

// TokenLockVotingRecords is a free data retrieval call binding the contract method 0x1eb78737.
//
// Solidity: function tokenLockVotingRecords(uint256 , bytes32 ) view returns(uint256 dynasty, uint256 accumlatedShares)
func (_TokenBank *TokenBankCaller) TokenLockVotingRecords(opts *bind.CallOpts, arg0 *big.Int, arg1 [32]byte) (struct {
	Dynasty          *big.Int
	AccumlatedShares *big.Int
}, error) {
	var out []interface{}
	err := _TokenBank.contract.Call(opts, &out, "tokenLockVotingRecords", arg0, arg1)

	outstruct := new(struct {
		Dynasty          *big.Int
		AccumlatedShares *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Dynasty = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AccumlatedShares = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TokenLockVotingRecords is a free data retrieval call binding the contract method 0x1eb78737.
//
// Solidity: function tokenLockVotingRecords(uint256 , bytes32 ) view returns(uint256 dynasty, uint256 accumlatedShares)
func (_TokenBank *TokenBankSession) TokenLockVotingRecords(arg0 *big.Int, arg1 [32]byte) (struct {
	Dynasty          *big.Int
	AccumlatedShares *big.Int
}, error) {
	return _TokenBank.Contract.TokenLockVotingRecords(&_TokenBank.CallOpts, arg0, arg1)
}

// TokenLockVotingRecords is a free data retrieval call binding the contract method 0x1eb78737.
//
// Solidity: function tokenLockVotingRecords(uint256 , bytes32 ) view returns(uint256 dynasty, uint256 accumlatedShares)
func (_TokenBank *TokenBankCallerSession) TokenLockVotingRecords(arg0 *big.Int, arg1 [32]byte) (struct {
	Dynasty          *big.Int
	AccumlatedShares *big.Int
}, error) {
	return _TokenBank.Contract.TokenLockVotingRecords(&_TokenBank.CallOpts, arg0, arg1)
}

// TokenUnlockNonceMap is a free data retrieval call binding the contract method 0xccf187c7.
//
// Solidity: function tokenUnlockNonceMap(uint256 ) view returns(uint256)
func (_TokenBank *TokenBankCaller) TokenUnlockNonceMap(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TokenBank.contract.Call(opts, &out, "tokenUnlockNonceMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenUnlockNonceMap is a free data retrieval call binding the contract method 0xccf187c7.
//
// Solidity: function tokenUnlockNonceMap(uint256 ) view returns(uint256)
func (_TokenBank *TokenBankSession) TokenUnlockNonceMap(arg0 *big.Int) (*big.Int, error) {
	return _TokenBank.Contract.TokenUnlockNonceMap(&_TokenBank.CallOpts, arg0)
}

// TokenUnlockNonceMap is a free data retrieval call binding the contract method 0xccf187c7.
//
// Solidity: function tokenUnlockNonceMap(uint256 ) view returns(uint256)
func (_TokenBank *TokenBankCallerSession) TokenUnlockNonceMap(arg0 *big.Int) (*big.Int, error) {
	return _TokenBank.Contract.TokenUnlockNonceMap(&_TokenBank.CallOpts, arg0)
}

// VoucherBurnNonceMap is a free data retrieval call binding the contract method 0xca207569.
//
// Solidity: function voucherBurnNonceMap(uint256 ) view returns(uint256)
func (_TokenBank *TokenBankCaller) VoucherBurnNonceMap(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TokenBank.contract.Call(opts, &out, "voucherBurnNonceMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VoucherBurnNonceMap is a free data retrieval call binding the contract method 0xca207569.
//
// Solidity: function voucherBurnNonceMap(uint256 ) view returns(uint256)
func (_TokenBank *TokenBankSession) VoucherBurnNonceMap(arg0 *big.Int) (*big.Int, error) {
	return _TokenBank.Contract.VoucherBurnNonceMap(&_TokenBank.CallOpts, arg0)
}

// VoucherBurnNonceMap is a free data retrieval call binding the contract method 0xca207569.
//
// Solidity: function voucherBurnNonceMap(uint256 ) view returns(uint256)
func (_TokenBank *TokenBankCallerSession) VoucherBurnNonceMap(arg0 *big.Int) (*big.Int, error) {
	return _TokenBank.Contract.VoucherBurnNonceMap(&_TokenBank.CallOpts, arg0)
}

// VoucherBurnVotingRecords is a free data retrieval call binding the contract method 0xfeaff052.
//
// Solidity: function voucherBurnVotingRecords(uint256 , bytes32 ) view returns(uint256 dynasty, uint256 accumlatedShares)
func (_TokenBank *TokenBankCaller) VoucherBurnVotingRecords(opts *bind.CallOpts, arg0 *big.Int, arg1 [32]byte) (struct {
	Dynasty          *big.Int
	AccumlatedShares *big.Int
}, error) {
	var out []interface{}
	err := _TokenBank.contract.Call(opts, &out, "voucherBurnVotingRecords", arg0, arg1)

	outstruct := new(struct {
		Dynasty          *big.Int
		AccumlatedShares *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Dynasty = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AccumlatedShares = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// VoucherBurnVotingRecords is a free data retrieval call binding the contract method 0xfeaff052.
//
// Solidity: function voucherBurnVotingRecords(uint256 , bytes32 ) view returns(uint256 dynasty, uint256 accumlatedShares)
func (_TokenBank *TokenBankSession) VoucherBurnVotingRecords(arg0 *big.Int, arg1 [32]byte) (struct {
	Dynasty          *big.Int
	AccumlatedShares *big.Int
}, error) {
	return _TokenBank.Contract.VoucherBurnVotingRecords(&_TokenBank.CallOpts, arg0, arg1)
}

// VoucherBurnVotingRecords is a free data retrieval call binding the contract method 0xfeaff052.
//
// Solidity: function voucherBurnVotingRecords(uint256 , bytes32 ) view returns(uint256 dynasty, uint256 accumlatedShares)
func (_TokenBank *TokenBankCallerSession) VoucherBurnVotingRecords(arg0 *big.Int, arg1 [32]byte) (struct {
	Dynasty          *big.Int
	AccumlatedShares *big.Int
}, error) {
	return _TokenBank.Contract.VoucherBurnVotingRecords(&_TokenBank.CallOpts, arg0, arg1)
}

// VoucherMintNonceMap is a free data retrieval call binding the contract method 0x740cb7f8.
//
// Solidity: function voucherMintNonceMap(uint256 ) view returns(uint256)
func (_TokenBank *TokenBankCaller) VoucherMintNonceMap(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TokenBank.contract.Call(opts, &out, "voucherMintNonceMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VoucherMintNonceMap is a free data retrieval call binding the contract method 0x740cb7f8.
//
// Solidity: function voucherMintNonceMap(uint256 ) view returns(uint256)
func (_TokenBank *TokenBankSession) VoucherMintNonceMap(arg0 *big.Int) (*big.Int, error) {
	return _TokenBank.Contract.VoucherMintNonceMap(&_TokenBank.CallOpts, arg0)
}

// VoucherMintNonceMap is a free data retrieval call binding the contract method 0x740cb7f8.
//
// Solidity: function voucherMintNonceMap(uint256 ) view returns(uint256)
func (_TokenBank *TokenBankCallerSession) VoucherMintNonceMap(arg0 *big.Int) (*big.Int, error) {
	return _TokenBank.Contract.VoucherMintNonceMap(&_TokenBank.CallOpts, arg0)
}

// TokenBankFailedToSendTFuelIterator is returned from FilterFailedToSendTFuel and is used to iterate over the raw logs and unpacked data for FailedToSendTFuel events raised by the TokenBank contract.
type TokenBankFailedToSendTFuelIterator struct {
	Event *TokenBankFailedToSendTFuel // Event containing the contract specifics and raw log

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
func (it *TokenBankFailedToSendTFuelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenBankFailedToSendTFuel)
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
		it.Event = new(TokenBankFailedToSendTFuel)
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
func (it *TokenBankFailedToSendTFuelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenBankFailedToSendTFuelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenBankFailedToSendTFuel represents a FailedToSendTFuel event raised by the TokenBank contract.
type TokenBankFailedToSendTFuel struct {
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFailedToSendTFuel is a free log retrieval operation binding the contract event 0x562a1007af95860758404d928a251ad8b0062ac50058db9f82dab3fe379f4885.
//
// Solidity: event FailedToSendTFuel(address indexed receiver, uint256 amount)
func (_TokenBank *TokenBankFilterer) FilterFailedToSendTFuel(opts *bind.FilterOpts, receiver []common.Address) (*TokenBankFailedToSendTFuelIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _TokenBank.contract.FilterLogs(opts, "FailedToSendTFuel", receiverRule)
	if err != nil {
		return nil, err
	}
	return &TokenBankFailedToSendTFuelIterator{contract: _TokenBank.contract, event: "FailedToSendTFuel", logs: logs, sub: sub}, nil
}

// WatchFailedToSendTFuel is a free log subscription operation binding the contract event 0x562a1007af95860758404d928a251ad8b0062ac50058db9f82dab3fe379f4885.
//
// Solidity: event FailedToSendTFuel(address indexed receiver, uint256 amount)
func (_TokenBank *TokenBankFilterer) WatchFailedToSendTFuel(opts *bind.WatchOpts, sink chan<- *TokenBankFailedToSendTFuel, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _TokenBank.contract.WatchLogs(opts, "FailedToSendTFuel", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenBankFailedToSendTFuel)
				if err := _TokenBank.contract.UnpackLog(event, "FailedToSendTFuel", log); err != nil {
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

// ParseFailedToSendTFuel is a log parse operation binding the contract event 0x562a1007af95860758404d928a251ad8b0062ac50058db9f82dab3fe379f4885.
//
// Solidity: event FailedToSendTFuel(address indexed receiver, uint256 amount)
func (_TokenBank *TokenBankFilterer) ParseFailedToSendTFuel(log types.Log) (*TokenBankFailedToSendTFuel, error) {
	event := new(TokenBankFailedToSendTFuel)
	if err := _TokenBank.contract.UnpackLog(event, "FailedToSendTFuel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
