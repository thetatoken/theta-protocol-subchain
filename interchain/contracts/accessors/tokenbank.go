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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"mainchainID_\",\"type\":\"uint256\"},{\"internalType\":\"contractChainRegistrar\",\"name\":\"chainRegistrar_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FailedToSendTFuel\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"mainchainID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenLockNonceMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"tokenLockVotingRecords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dynasty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumlatedShares\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenUnlockNonceMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"voucherBurnNonceMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"voucherBurnVotingRecords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dynasty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumlatedShares\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"voucherMintNonceMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"getMaxProcessedTokenLockNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"eventNonce\",\"type\":\"uint256\"}],\"name\":\"getTokenLockEventHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"getMaxProcessedVoucherBurnNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"eventNonce\",\"type\":\"uint256\"}],\"name\":\"getVoucherBurnEventHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"dynasty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainTokenLockNonce\",\"type\":\"uint256\"}],\"name\":\"forceIncrementMaxProcessedTokenLockNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"dynasty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainVoucherBurnNonce\",\"type\":\"uint256\"}],\"name\":\"forceIncrementMaxProcessedVoucherBurnNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subchainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dynasty\",\"type\":\"uint256\"}],\"name\":\"getAdjustedValidatorSet\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"shareAmounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161148a38038061148a83398101604081905261002f9161005e565b6001600081905591909155600280546001600160a01b0319166001600160a01b0390921691909117905561009b565b6000806040838503121561007157600080fd5b825160208401519092506001600160a01b038116811461009057600080fd5b809150509250929050565b6113e0806100aa6000396000f3fe608060405234801561001057600080fd5b50600436106100ca5760003560e01c8063aa861c151161007c578063aa861c15146101d2578063b4baab85146101f3578063ca20756914610206578063ccf187c714610226578063dd17eb6d14610246578063f95627ac14610271578063feaff0521461029157600080fd5b8063060cb552146100cf578063073b9502146100e45780631eb78737146101005780636ac739b914610147578063740cb7f814610172578063766f8fb0146101925780638883931e146101b2575b600080fd5b6100e26100dd366004610ff7565b6102c3565b005b6100ed60015481565b6040519081526020015b60405180910390f35b61013261010e3660046110fb565b600b6020908152600092835260408084209091529082529020805460029091015482565b604080519283526020830191909152016100f7565b6100ed6101553660046110fb565b600091825260086020908152604080842092845291905290205490565b6100ed61018036600461109d565b60066020526000908152604090205481565b6100ed6101a036600461109d565b6000908152600a602052604090205490565b6100ed6101c036600461109d565b60036020526000908152604090205481565b6101e56101e03660046110fb565b61037b565b6040516100f7929190611170565b6100e2610201366004610ff7565b610618565b6100ed61021436600461109d565b60056020526000908152604090205481565b6100ed61023436600461109d565b60046020526000908152604090205481565b6100ed6102543660046110fb565b600091825260076020908152604080842092845291905290205490565b6100ed61027f36600461109d565b60009081526009602052604090205490565b61013261029f3660046110fb565b600c6020908152600092835260408084209091529082529020805460029091015482565b600260005414156102ef5760405162461bcd60e51b81526004016102e69061122b565b60405180910390fd5b600260009081556102ff846106b8565b905060008061030d836106c9565b915091508061032e5760405162461bcd60e51b81526004016102e6906111e7565b600082848787604051602001610347949392919061111d565b60405160208183030381529060405280519060200120905061036c8387838833610751565b50506001600055505050505050565b6060806000600154461490506000600260009054906101000a90046001600160a01b03166001600160a01b031663a7464b126040518163ffffffff1660e01b815260040160206040518083038186803b1580156103d757600080fd5b505afa1580156103eb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061040f91906110b6565b600254604051632f2c13b560e01b81526004810189905291925060009182916001600160a01b031690632f2c13b590602401604080518083038186803b15801561045857600080fd5b505afa15801561046c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061049091906110cf565b9150915060008180156104ac57506104a8888561076e565b8310155b80156104cb57506104c8846104c28a6001610781565b9061076e565b83105b90508480156104d75750805b156105805760006104e9896001610781565b6002546040516343f27e4560e01b8152600481018d9052602481018390529192506001600160a01b0316906343f27e459060440160006040518083038186803b15801561053557600080fd5b505afa158015610549573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526105719190810190610f04565b97509750505050505050610611565b6002546040516343f27e4560e01b8152600481018b9052602481018a90526001600160a01b03909116906343f27e459060440160006040518083038186803b1580156105cb57600080fd5b505afa1580156105df573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526106079190810190610f04565b9650965050505050505b9250929050565b6002600054141561063b5760405162461bcd60e51b81526004016102e69061122b565b6002600090815561064b846106b8565b9050600080610659836106c9565b915091508061067a5760405162461bcd60e51b81526004016102e6906111e7565b600082848787604051602001610693949392919061111d565b60405160208183030381529060405280519060200120905061036c838783883361078d565b60606106c3826107a0565b92915050565b6000806000806106df85602f60f81b600161081a565b91509150806106f5575060009485945092505050565b600080610704876000866108aa565b915091508061071c5750600096879650945050505050565b600080610728846109b7565b915091508061074257506000988998509650505050505050565b50976001975095505050505050565b60006107648686868686600a600c610a89565b9695505050505050565b600061077a82846112f4565b9392505050565b600061077a82846112b7565b600061076486868686866009600b610a89565b60608160005b8151811015610813576107d88282815181106107c4576107c461137e565b01602001516001600160f81b031916610aa8565b8282815181106107ea576107ea61137e565b60200101906001600160f81b031916908160001a9053508061080b8161134d565b9150506107a6565b5092915050565b82516000908190859082805b8281101561089657876001600160f81b03191684828151811061084b5761084b61137e565b01602001516001600160f81b03191614156108845761086b6001836112b7565b915086821415610884579450600193506108a292505050565b8061088e8161134d565b915050610826565b50600080945094505050505b935093915050565b8251606090600090848410156108d4575050604080516020810190915260008082529091506108a2565b808411156108f6575050604080516020810190915260008082529091506108a2565b8560006109038688610af7565b67ffffffffffffffff81111561091b5761091b611394565b6040519080825280601f01601f191660200182016040528015610945576020820181803683370190505b509050865b86811015610742578281815181106109645761096461137e565b01602001516001600160f81b0319168261097e838b610af7565b8151811061098e5761098e61137e565b60200101906001600160f81b031916908160001a905350806109af8161134d565b91505061094a565b80516000908190839082805b82811015610a7c5760308482815181106109df576109df61137e565b016020015160f81c10801590610a0f57506039848281518110610a0457610a0461137e565b016020015160f81c11155b15610a5b57610a1f82600a61076e565b9150610a546030858381518110610a3857610a3861137e565b0160200151610a4a919060f81c61132a565b839060ff16610781565b9150610a6a565b50600096879650945050505050565b80610a748161134d565b9150506109c3565b5095600195509350505050565b6000610a9a88888888888888610b03565b90505b979650505050505050565b6000604160f81b6001600160f81b0319831610801590610ad65750602d60f91b6001600160f81b0319831611155b15610af357610aea60f883901c60206112cf565b60f81b92915050565b5090565b600061077a8284611313565b6000600154881480610b8c57506002546040516343b71f0560e01b8152600481018a90526001600160a01b03909116906343b71f059060240160206040518083038186803b158015610b5457600080fd5b505afa158015610b68573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b8c9190610fdc565b610bca5760405162461bcd60e51b815260206004820152600f60248201526e125b9d985b1a590818da185a5b9251608a1b60448201526064016102e6565b600088815260208490526040902054610be49060016112b7565b8514610bf257506000610a9d565b600088815260208381526040808320898452909152812060015460609081908c90811415610c1d5750465b610c27818d61037b565b9350915060005b8251811015610d8457896001600160a01b0316838281518110610c5357610c5361137e565b60200260200101516001600160a01b031614610c6e57610d72565b6001955060005b6001860154811015610d0f57856001018181548110610c9657610c9661137e565b6000918252602090912001546001600160a01b038c811691161415610cfd5760405162461bcd60e51b815260206004820152601c60248201527f546869732076616c696461746f7220616c726561647920766f7465640000000060448201526064016102e6565b80610d078161134d565b915050610c75565b508c85556001808601805491820181556000908152602090200180546001600160a01b031916331790558351610d6c90859083908110610d5157610d5161137e565b6020026020010151866002015461078190919063ffffffff16565b60028601555b80610d7c8161134d565b915050610c2e565b50505082610dc65760405162461bcd60e51b815260206004820152600f60248201526e2737ba1030903b30b634b230ba37b960891b60448201526064016102e6565b6000805b8251811015610e1357610dff838281518110610de857610de861137e565b60200260200101518361078190919063ffffffff16565b915080610e0b8161134d565b915050610dca565b50610e1f81600261076e565b6002840154610e2f90600361076e565b10610e6b5760008c815260208890526040902054610e4e9060016112b7565b60008d8152602089905260409020555060019350610a9d92505050565b5060009b9a5050505050505050505050565b600082601f830112610e8e57600080fd5b81516020610ea3610e9e83611293565b611262565b80838252828201915082860187848660051b8901011115610ec357600080fd5b60005b85811015610ee257815184529284019290840190600101610ec6565b5090979650505050505050565b80518015158114610eff57600080fd5b919050565b60008060408385031215610f1757600080fd5b825167ffffffffffffffff80821115610f2f57600080fd5b818501915085601f830112610f4357600080fd5b81516020610f53610e9e83611293565b8083825282820191508286018a848660051b8901011115610f7357600080fd5b600096505b84871015610fab5780516001600160a01b0381168114610f9757600080fd5b835260019690960195918301918301610f78565b5091880151919650909350505080821115610fc557600080fd5b50610fd285828601610e7d565b9150509250929050565b600060208284031215610fee57600080fd5b61077a82610eef565b60008060006060848603121561100c57600080fd5b833567ffffffffffffffff8082111561102457600080fd5b818601915086601f83011261103857600080fd5b813560208282111561104c5761104c611394565b61105e601f8301601f19168201611262565b9250818352888183860101111561107457600080fd5b818185018285013760009183018101919091529097908601359650604090950135949350505050565b6000602082840312156110af57600080fd5b5035919050565b6000602082840312156110c857600080fd5b5051919050565b600080604083850312156110e257600080fd5b825191506110f260208401610eef565b90509250929050565b6000806040838503121561110e57600080fd5b50508035926020909101359150565b8481526000845160005b8181101561114357602081880181015185830182015201611127565b81811115611155576000602083860101525b50909101602081019390935250604082015260600192915050565b604080825283519082018190526000906020906060840190828701845b828110156111b25781516001600160a01b03168452928401929084019060010161118d565b5050508381038285015284518082528583019183019060005b81811015610ee2578351835292840192918401916001016111cb565b60208082526024908201527f4661696c656420746f206578747261637420636861696e49442066726f6d2064604082015263656e6f6d60e01b606082015260800190565b6020808252601f908201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604082015260600190565b604051601f8201601f1916810167ffffffffffffffff8111828210171561128b5761128b611394565b604052919050565b600067ffffffffffffffff8211156112ad576112ad611394565b5060051b60200190565b600082198211156112ca576112ca611368565b500190565b600060ff821660ff84168060ff038211156112ec576112ec611368565b019392505050565b600081600019048311821515161561130e5761130e611368565b500290565b60008282101561132557611325611368565b500390565b600060ff821660ff84168082101561134457611344611368565b90039392505050565b600060001982141561136157611361611368565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fdfea2646970667358221220bcf0acd0fc15b33fa80828f483fb477336af38e3db08fa57e2d517f36ee9540964736f6c63430008070033",
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

// ForceIncrementMaxProcessedTokenLockNonce is a paid mutator transaction binding the contract method 0xb4baab85.
//
// Solidity: function forceIncrementMaxProcessedTokenLockNonce(string denom, uint256 dynasty, uint256 sourceChainTokenLockNonce) returns()
func (_TokenBank *TokenBankTransactor) ForceIncrementMaxProcessedTokenLockNonce(opts *bind.TransactOpts, denom string, dynasty *big.Int, sourceChainTokenLockNonce *big.Int) (*types.Transaction, error) {
	return _TokenBank.contract.Transact(opts, "forceIncrementMaxProcessedTokenLockNonce", denom, dynasty, sourceChainTokenLockNonce)
}

// ForceIncrementMaxProcessedTokenLockNonce is a paid mutator transaction binding the contract method 0xb4baab85.
//
// Solidity: function forceIncrementMaxProcessedTokenLockNonce(string denom, uint256 dynasty, uint256 sourceChainTokenLockNonce) returns()
func (_TokenBank *TokenBankSession) ForceIncrementMaxProcessedTokenLockNonce(denom string, dynasty *big.Int, sourceChainTokenLockNonce *big.Int) (*types.Transaction, error) {
	return _TokenBank.Contract.ForceIncrementMaxProcessedTokenLockNonce(&_TokenBank.TransactOpts, denom, dynasty, sourceChainTokenLockNonce)
}

// ForceIncrementMaxProcessedTokenLockNonce is a paid mutator transaction binding the contract method 0xb4baab85.
//
// Solidity: function forceIncrementMaxProcessedTokenLockNonce(string denom, uint256 dynasty, uint256 sourceChainTokenLockNonce) returns()
func (_TokenBank *TokenBankTransactorSession) ForceIncrementMaxProcessedTokenLockNonce(denom string, dynasty *big.Int, sourceChainTokenLockNonce *big.Int) (*types.Transaction, error) {
	return _TokenBank.Contract.ForceIncrementMaxProcessedTokenLockNonce(&_TokenBank.TransactOpts, denom, dynasty, sourceChainTokenLockNonce)
}

// ForceIncrementMaxProcessedVoucherBurnNonce is a paid mutator transaction binding the contract method 0x060cb552.
//
// Solidity: function forceIncrementMaxProcessedVoucherBurnNonce(string denom, uint256 dynasty, uint256 sourceChainVoucherBurnNonce) returns()
func (_TokenBank *TokenBankTransactor) ForceIncrementMaxProcessedVoucherBurnNonce(opts *bind.TransactOpts, denom string, dynasty *big.Int, sourceChainVoucherBurnNonce *big.Int) (*types.Transaction, error) {
	return _TokenBank.contract.Transact(opts, "forceIncrementMaxProcessedVoucherBurnNonce", denom, dynasty, sourceChainVoucherBurnNonce)
}

// ForceIncrementMaxProcessedVoucherBurnNonce is a paid mutator transaction binding the contract method 0x060cb552.
//
// Solidity: function forceIncrementMaxProcessedVoucherBurnNonce(string denom, uint256 dynasty, uint256 sourceChainVoucherBurnNonce) returns()
func (_TokenBank *TokenBankSession) ForceIncrementMaxProcessedVoucherBurnNonce(denom string, dynasty *big.Int, sourceChainVoucherBurnNonce *big.Int) (*types.Transaction, error) {
	return _TokenBank.Contract.ForceIncrementMaxProcessedVoucherBurnNonce(&_TokenBank.TransactOpts, denom, dynasty, sourceChainVoucherBurnNonce)
}

// ForceIncrementMaxProcessedVoucherBurnNonce is a paid mutator transaction binding the contract method 0x060cb552.
//
// Solidity: function forceIncrementMaxProcessedVoucherBurnNonce(string denom, uint256 dynasty, uint256 sourceChainVoucherBurnNonce) returns()
func (_TokenBank *TokenBankTransactorSession) ForceIncrementMaxProcessedVoucherBurnNonce(denom string, dynasty *big.Int, sourceChainVoucherBurnNonce *big.Int) (*types.Transaction, error) {
	return _TokenBank.Contract.ForceIncrementMaxProcessedVoucherBurnNonce(&_TokenBank.TransactOpts, denom, dynasty, sourceChainVoucherBurnNonce)
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
