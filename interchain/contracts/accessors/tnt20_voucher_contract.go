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

// TNT20VoucherContractMetaData contains all meta data concerning the TNT20VoucherContract contract.
var TNT20VoucherContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"denom_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals_\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"UpdateOwner\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"updateOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voucherReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mintedAmount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voucherOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"burnedAmount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"denom\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620011493803806200114983398101604081905262000034916200021b565b84838381600390805190602001906200004f929190620000be565b50805162000065906004906020840190620000be565b5050600580546001600160a01b0319166001600160a01b0393909316929092179091555083516200009e906006906020870190620000be565b506007805460ff191660ff92909216919091179055506200033c92505050565b828054620000cc90620002e9565b90600052602060002090601f016020900481019282620000f057600085556200013b565b82601f106200010b57805160ff19168380011785556200013b565b828001600101855582156200013b579182015b828111156200013b5782518255916020019190600101906200011e565b50620001499291506200014d565b5090565b5b808211156200014957600081556001016200014e565b600082601f8301126200017657600080fd5b81516001600160401b038082111562000193576200019362000326565b604051601f8301601f19908116603f01168101908282118183101715620001be57620001be62000326565b81604052838152602092508683858801011115620001db57600080fd5b600091505b83821015620001ff5785820183015181830184015290820190620001e0565b83821115620002115760008385830101525b9695505050505050565b600080600080600060a086880312156200023457600080fd5b85516001600160a01b03811681146200024c57600080fd5b60208701519095506001600160401b03808211156200026a57600080fd5b6200027889838a0162000164565b955060408801519150808211156200028f57600080fd5b6200029d89838a0162000164565b94506060880151915080821115620002b457600080fd5b50620002c38882890162000164565b925050608086015160ff81168114620002db57600080fd5b809150509295509295909350565b600181811c90821680620002fe57607f821691505b602082108114156200032057634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052604160045260246000fd5b610dfd806200034c6000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c8063880cdc3111610097578063a457c2d711610066578063a457c2d71461022a578063a9059cbb1461023d578063c370b04214610250578063dd62ed3e1461025857600080fd5b8063880cdc31146101d15780638da5cb5b146101e457806395d89b411461020f5780639dc29fac1461021757600080fd5b8063313ce567116100d3578063313ce5671461016b578063395093511461018057806340c10f191461019357806370a08231146101a857600080fd5b806306fdde0314610105578063095ea7b31461012357806318160ddd1461014657806323b872dd14610158575b600080fd5b61010d61026b565b60405161011a9190610cbb565b60405180910390f35b610136610131366004610c91565b6102fd565b604051901515815260200161011a565b6002545b60405190815260200161011a565b610136610166366004610c55565b610315565b60075460405160ff909116815260200161011a565b61013661018e366004610c91565b610339565b6101a66101a1366004610c91565b61035b565b005b61014a6101b6366004610c00565b6001600160a01b031660009081526020819052604090205490565b6101a66101df366004610c00565b61039c565b6005546101f7906001600160a01b031681565b6040516001600160a01b03909116815260200161011a565b61010d61042f565b6101a6610225366004610c91565b61043e565b610136610238366004610c91565b61058f565b61013661024b366004610c91565b61060a565b61010d610618565b61014a610266366004610c22565b610627565b60606003805461027a90610d76565b80601f01602080910402602001604051908101604052809291908181526020018280546102a690610d76565b80156102f35780601f106102c8576101008083540402835291602001916102f3565b820191906000526020600020905b8154815290600101906020018083116102d657829003601f168201915b5050505050905090565b60003361030b818585610652565b5060019392505050565b600033610323858285610777565b61032e8585856107f1565b506001949350505050565b60003361030b81858561034c8383610627565b6103569190610d47565b610652565b6005546001600160a01b0316331461038e5760405162461bcd60e51b815260040161038590610d10565b60405180910390fd5b61039882826109bf565b5050565b6005546001600160a01b031633146103c65760405162461bcd60e51b815260040161038590610d10565b600554604080516001600160a01b03928316815291831660208301527fe2c7d1c4da37855e682bde14f17826d185497973b73fba7554daa6da467058d9910160405180910390a1600580546001600160a01b0319166001600160a01b0392909216919091179055565b60606004805461027a90610d76565b6005546001600160a01b031633146104685760405162461bcd60e51b815260040161038590610d10565b6001600160a01b038216600090815260208190526040902054818110156104ec5760405162461bcd60e51b815260206004820152603260248201527f566f7563686572206f776e657220646f6573206e6f74206861766520656e6f7560448201527133b4103130b630b731b2903a3790313ab93760711b6064820152608401610385565b6005546001600160a01b031660006105048583610627565b9050838110156105735760405162461bcd60e51b815260206004820152603460248201527f566f7563686572206f776e657220646964206e6f7420617070726f76656420656044820152733737bab3b41030b6b7bab73a103a3790313ab93760611b6064820152608401610385565b61057e858386610777565b6105888585610a9e565b5050505050565b6000338161059d8286610627565b9050838110156105fd5760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b6064820152608401610385565b61032e8286868403610652565b60003361030b8185856107f1565b60606006805461027a90610d76565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6001600160a01b0383166106b45760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b6064820152608401610385565b6001600160a01b0382166107155760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b6064820152608401610385565b6001600160a01b0383811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b60006107838484610627565b905060001981146107eb57818110156107de5760405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e63650000006044820152606401610385565b6107eb8484848403610652565b50505050565b6001600160a01b0383166108555760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b6064820152608401610385565b6001600160a01b0382166108b75760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b6064820152608401610385565b6001600160a01b0383166000908152602081905260409020548181101561092f5760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b6064820152608401610385565b6001600160a01b03808516600090815260208190526040808220858503905591851681529081208054849290610966908490610d47565b92505081905550826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516109b291815260200190565b60405180910390a36107eb565b6001600160a01b038216610a155760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152606401610385565b8060026000828254610a279190610d47565b90915550506001600160a01b03821660009081526020819052604081208054839290610a54908490610d47565b90915550506040518181526001600160a01b038316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b6001600160a01b038216610afe5760405162461bcd60e51b815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f206164647265736044820152607360f81b6064820152608401610385565b6001600160a01b03821660009081526020819052604090205481811015610b725760405162461bcd60e51b815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e604482015261636560f01b6064820152608401610385565b6001600160a01b0383166000908152602081905260408120838303905560028054849290610ba1908490610d5f565b90915550506040518281526000906001600160a01b038516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200161076a565b80356001600160a01b0381168114610bfb57600080fd5b919050565b600060208284031215610c1257600080fd5b610c1b82610be4565b9392505050565b60008060408385031215610c3557600080fd5b610c3e83610be4565b9150610c4c60208401610be4565b90509250929050565b600080600060608486031215610c6a57600080fd5b610c7384610be4565b9250610c8160208501610be4565b9150604084013590509250925092565b60008060408385031215610ca457600080fd5b610cad83610be4565b946020939093013593505050565b600060208083528351808285015260005b81811015610ce857858101830151858201604001528201610ccc565b81811115610cfa576000604083870101525b50601f01601f1916929092016040019392505050565b6020808252601c908201527f6f6e6c79206f776e65722063616e206d616b65207468652063616c6c00000000604082015260600190565b60008219821115610d5a57610d5a610db1565b500190565b600082821015610d7157610d71610db1565b500390565b600181811c90821680610d8a57607f821691505b60208210811415610dab57634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601160045260246000fdfea2646970667358221220d1aded6c4145047407d223851ada322b32ff26f9b6a9a77f47369211fda053e564736f6c63430008070033",
}

// TNT20VoucherContractABI is the input ABI used to generate the binding from.
// Deprecated: Use TNT20VoucherContractMetaData.ABI instead.
var TNT20VoucherContractABI = TNT20VoucherContractMetaData.ABI

// TNT20VoucherContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TNT20VoucherContractMetaData.Bin instead.
var TNT20VoucherContractBin = TNT20VoucherContractMetaData.Bin

// DeployTNT20VoucherContract deploys a new Ethereum contract, binding an instance of TNT20VoucherContract to it.
func DeployTNT20VoucherContract(auth *bind.TransactOpts, backend bind.ContractBackend, owner_ common.Address, denom_ string, name_ string, symbol_ string, decimals_ uint8) (common.Address, *types.Transaction, *TNT20VoucherContract, error) {
	parsed, err := TNT20VoucherContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TNT20VoucherContractBin), backend, owner_, denom_, name_, symbol_, decimals_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TNT20VoucherContract{TNT20VoucherContractCaller: TNT20VoucherContractCaller{contract: contract}, TNT20VoucherContractTransactor: TNT20VoucherContractTransactor{contract: contract}, TNT20VoucherContractFilterer: TNT20VoucherContractFilterer{contract: contract}}, nil
}

// TNT20VoucherContract is an auto generated Go binding around an Ethereum contract.
type TNT20VoucherContract struct {
	TNT20VoucherContractCaller     // Read-only binding to the contract
	TNT20VoucherContractTransactor // Write-only binding to the contract
	TNT20VoucherContractFilterer   // Log filterer for contract events
}

// TNT20VoucherContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type TNT20VoucherContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TNT20VoucherContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TNT20VoucherContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TNT20VoucherContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TNT20VoucherContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TNT20VoucherContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TNT20VoucherContractSession struct {
	Contract     *TNT20VoucherContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TNT20VoucherContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TNT20VoucherContractCallerSession struct {
	Contract *TNT20VoucherContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// TNT20VoucherContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TNT20VoucherContractTransactorSession struct {
	Contract     *TNT20VoucherContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// TNT20VoucherContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type TNT20VoucherContractRaw struct {
	Contract *TNT20VoucherContract // Generic contract binding to access the raw methods on
}

// TNT20VoucherContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TNT20VoucherContractCallerRaw struct {
	Contract *TNT20VoucherContractCaller // Generic read-only contract binding to access the raw methods on
}

// TNT20VoucherContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TNT20VoucherContractTransactorRaw struct {
	Contract *TNT20VoucherContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTNT20VoucherContract creates a new instance of TNT20VoucherContract, bound to a specific deployed contract.
func NewTNT20VoucherContract(address common.Address, backend bind.ContractBackend) (*TNT20VoucherContract, error) {
	contract, err := bindTNT20VoucherContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TNT20VoucherContract{TNT20VoucherContractCaller: TNT20VoucherContractCaller{contract: contract}, TNT20VoucherContractTransactor: TNT20VoucherContractTransactor{contract: contract}, TNT20VoucherContractFilterer: TNT20VoucherContractFilterer{contract: contract}}, nil
}

// NewTNT20VoucherContractCaller creates a new read-only instance of TNT20VoucherContract, bound to a specific deployed contract.
func NewTNT20VoucherContractCaller(address common.Address, caller bind.ContractCaller) (*TNT20VoucherContractCaller, error) {
	contract, err := bindTNT20VoucherContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TNT20VoucherContractCaller{contract: contract}, nil
}

// NewTNT20VoucherContractTransactor creates a new write-only instance of TNT20VoucherContract, bound to a specific deployed contract.
func NewTNT20VoucherContractTransactor(address common.Address, transactor bind.ContractTransactor) (*TNT20VoucherContractTransactor, error) {
	contract, err := bindTNT20VoucherContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TNT20VoucherContractTransactor{contract: contract}, nil
}

// NewTNT20VoucherContractFilterer creates a new log filterer instance of TNT20VoucherContract, bound to a specific deployed contract.
func NewTNT20VoucherContractFilterer(address common.Address, filterer bind.ContractFilterer) (*TNT20VoucherContractFilterer, error) {
	contract, err := bindTNT20VoucherContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TNT20VoucherContractFilterer{contract: contract}, nil
}

// bindTNT20VoucherContract binds a generic wrapper to an already deployed contract.
func bindTNT20VoucherContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TNT20VoucherContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TNT20VoucherContract *TNT20VoucherContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TNT20VoucherContract.Contract.TNT20VoucherContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TNT20VoucherContract *TNT20VoucherContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TNT20VoucherContract.Contract.TNT20VoucherContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TNT20VoucherContract *TNT20VoucherContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TNT20VoucherContract.Contract.TNT20VoucherContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TNT20VoucherContract *TNT20VoucherContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TNT20VoucherContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TNT20VoucherContract *TNT20VoucherContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TNT20VoucherContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TNT20VoucherContract *TNT20VoucherContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TNT20VoucherContract.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TNT20VoucherContract *TNT20VoucherContractCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TNT20VoucherContract.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TNT20VoucherContract *TNT20VoucherContractSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _TNT20VoucherContract.Contract.Allowance(&_TNT20VoucherContract.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TNT20VoucherContract *TNT20VoucherContractCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _TNT20VoucherContract.Contract.Allowance(&_TNT20VoucherContract.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TNT20VoucherContract *TNT20VoucherContractCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TNT20VoucherContract.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TNT20VoucherContract *TNT20VoucherContractSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _TNT20VoucherContract.Contract.BalanceOf(&_TNT20VoucherContract.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TNT20VoucherContract *TNT20VoucherContractCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _TNT20VoucherContract.Contract.BalanceOf(&_TNT20VoucherContract.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TNT20VoucherContract *TNT20VoucherContractCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _TNT20VoucherContract.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TNT20VoucherContract *TNT20VoucherContractSession) Decimals() (uint8, error) {
	return _TNT20VoucherContract.Contract.Decimals(&_TNT20VoucherContract.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TNT20VoucherContract *TNT20VoucherContractCallerSession) Decimals() (uint8, error) {
	return _TNT20VoucherContract.Contract.Decimals(&_TNT20VoucherContract.CallOpts)
}

// Denom is a free data retrieval call binding the contract method 0xc370b042.
//
// Solidity: function denom() view returns(string)
func (_TNT20VoucherContract *TNT20VoucherContractCaller) Denom(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TNT20VoucherContract.contract.Call(opts, &out, "denom")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Denom is a free data retrieval call binding the contract method 0xc370b042.
//
// Solidity: function denom() view returns(string)
func (_TNT20VoucherContract *TNT20VoucherContractSession) Denom() (string, error) {
	return _TNT20VoucherContract.Contract.Denom(&_TNT20VoucherContract.CallOpts)
}

// Denom is a free data retrieval call binding the contract method 0xc370b042.
//
// Solidity: function denom() view returns(string)
func (_TNT20VoucherContract *TNT20VoucherContractCallerSession) Denom() (string, error) {
	return _TNT20VoucherContract.Contract.Denom(&_TNT20VoucherContract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TNT20VoucherContract *TNT20VoucherContractCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TNT20VoucherContract.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TNT20VoucherContract *TNT20VoucherContractSession) Name() (string, error) {
	return _TNT20VoucherContract.Contract.Name(&_TNT20VoucherContract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TNT20VoucherContract *TNT20VoucherContractCallerSession) Name() (string, error) {
	return _TNT20VoucherContract.Contract.Name(&_TNT20VoucherContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TNT20VoucherContract *TNT20VoucherContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TNT20VoucherContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TNT20VoucherContract *TNT20VoucherContractSession) Owner() (common.Address, error) {
	return _TNT20VoucherContract.Contract.Owner(&_TNT20VoucherContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TNT20VoucherContract *TNT20VoucherContractCallerSession) Owner() (common.Address, error) {
	return _TNT20VoucherContract.Contract.Owner(&_TNT20VoucherContract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TNT20VoucherContract *TNT20VoucherContractCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TNT20VoucherContract.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TNT20VoucherContract *TNT20VoucherContractSession) Symbol() (string, error) {
	return _TNT20VoucherContract.Contract.Symbol(&_TNT20VoucherContract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TNT20VoucherContract *TNT20VoucherContractCallerSession) Symbol() (string, error) {
	return _TNT20VoucherContract.Contract.Symbol(&_TNT20VoucherContract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TNT20VoucherContract *TNT20VoucherContractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TNT20VoucherContract.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TNT20VoucherContract *TNT20VoucherContractSession) TotalSupply() (*big.Int, error) {
	return _TNT20VoucherContract.Contract.TotalSupply(&_TNT20VoucherContract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TNT20VoucherContract *TNT20VoucherContractCallerSession) TotalSupply() (*big.Int, error) {
	return _TNT20VoucherContract.Contract.TotalSupply(&_TNT20VoucherContract.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TNT20VoucherContract *TNT20VoucherContractTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TNT20VoucherContract.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TNT20VoucherContract *TNT20VoucherContractSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TNT20VoucherContract.Contract.Approve(&_TNT20VoucherContract.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TNT20VoucherContract *TNT20VoucherContractTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TNT20VoucherContract.Contract.Approve(&_TNT20VoucherContract.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address voucherOwner, uint256 burnedAmount) returns()
func (_TNT20VoucherContract *TNT20VoucherContractTransactor) Burn(opts *bind.TransactOpts, voucherOwner common.Address, burnedAmount *big.Int) (*types.Transaction, error) {
	return _TNT20VoucherContract.contract.Transact(opts, "burn", voucherOwner, burnedAmount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address voucherOwner, uint256 burnedAmount) returns()
func (_TNT20VoucherContract *TNT20VoucherContractSession) Burn(voucherOwner common.Address, burnedAmount *big.Int) (*types.Transaction, error) {
	return _TNT20VoucherContract.Contract.Burn(&_TNT20VoucherContract.TransactOpts, voucherOwner, burnedAmount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address voucherOwner, uint256 burnedAmount) returns()
func (_TNT20VoucherContract *TNT20VoucherContractTransactorSession) Burn(voucherOwner common.Address, burnedAmount *big.Int) (*types.Transaction, error) {
	return _TNT20VoucherContract.Contract.Burn(&_TNT20VoucherContract.TransactOpts, voucherOwner, burnedAmount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TNT20VoucherContract *TNT20VoucherContractTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TNT20VoucherContract.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TNT20VoucherContract *TNT20VoucherContractSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TNT20VoucherContract.Contract.DecreaseAllowance(&_TNT20VoucherContract.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TNT20VoucherContract *TNT20VoucherContractTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TNT20VoucherContract.Contract.DecreaseAllowance(&_TNT20VoucherContract.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TNT20VoucherContract *TNT20VoucherContractTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TNT20VoucherContract.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TNT20VoucherContract *TNT20VoucherContractSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TNT20VoucherContract.Contract.IncreaseAllowance(&_TNT20VoucherContract.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TNT20VoucherContract *TNT20VoucherContractTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TNT20VoucherContract.Contract.IncreaseAllowance(&_TNT20VoucherContract.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address voucherReceiver, uint256 mintedAmount) returns()
func (_TNT20VoucherContract *TNT20VoucherContractTransactor) Mint(opts *bind.TransactOpts, voucherReceiver common.Address, mintedAmount *big.Int) (*types.Transaction, error) {
	return _TNT20VoucherContract.contract.Transact(opts, "mint", voucherReceiver, mintedAmount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address voucherReceiver, uint256 mintedAmount) returns()
func (_TNT20VoucherContract *TNT20VoucherContractSession) Mint(voucherReceiver common.Address, mintedAmount *big.Int) (*types.Transaction, error) {
	return _TNT20VoucherContract.Contract.Mint(&_TNT20VoucherContract.TransactOpts, voucherReceiver, mintedAmount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address voucherReceiver, uint256 mintedAmount) returns()
func (_TNT20VoucherContract *TNT20VoucherContractTransactorSession) Mint(voucherReceiver common.Address, mintedAmount *big.Int) (*types.Transaction, error) {
	return _TNT20VoucherContract.Contract.Mint(&_TNT20VoucherContract.TransactOpts, voucherReceiver, mintedAmount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_TNT20VoucherContract *TNT20VoucherContractTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TNT20VoucherContract.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_TNT20VoucherContract *TNT20VoucherContractSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TNT20VoucherContract.Contract.Transfer(&_TNT20VoucherContract.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_TNT20VoucherContract *TNT20VoucherContractTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TNT20VoucherContract.Contract.Transfer(&_TNT20VoucherContract.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_TNT20VoucherContract *TNT20VoucherContractTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TNT20VoucherContract.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_TNT20VoucherContract *TNT20VoucherContractSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TNT20VoucherContract.Contract.TransferFrom(&_TNT20VoucherContract.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_TNT20VoucherContract *TNT20VoucherContractTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TNT20VoucherContract.Contract.TransferFrom(&_TNT20VoucherContract.TransactOpts, from, to, amount)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0x880cdc31.
//
// Solidity: function updateOwner(address newOwner) returns()
func (_TNT20VoucherContract *TNT20VoucherContractTransactor) UpdateOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TNT20VoucherContract.contract.Transact(opts, "updateOwner", newOwner)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0x880cdc31.
//
// Solidity: function updateOwner(address newOwner) returns()
func (_TNT20VoucherContract *TNT20VoucherContractSession) UpdateOwner(newOwner common.Address) (*types.Transaction, error) {
	return _TNT20VoucherContract.Contract.UpdateOwner(&_TNT20VoucherContract.TransactOpts, newOwner)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0x880cdc31.
//
// Solidity: function updateOwner(address newOwner) returns()
func (_TNT20VoucherContract *TNT20VoucherContractTransactorSession) UpdateOwner(newOwner common.Address) (*types.Transaction, error) {
	return _TNT20VoucherContract.Contract.UpdateOwner(&_TNT20VoucherContract.TransactOpts, newOwner)
}

// TNT20VoucherContractApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the TNT20VoucherContract contract.
type TNT20VoucherContractApprovalIterator struct {
	Event *TNT20VoucherContractApproval // Event containing the contract specifics and raw log

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
func (it *TNT20VoucherContractApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TNT20VoucherContractApproval)
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
		it.Event = new(TNT20VoucherContractApproval)
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
func (it *TNT20VoucherContractApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TNT20VoucherContractApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TNT20VoucherContractApproval represents a Approval event raised by the TNT20VoucherContract contract.
type TNT20VoucherContractApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TNT20VoucherContract *TNT20VoucherContractFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*TNT20VoucherContractApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TNT20VoucherContract.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TNT20VoucherContractApprovalIterator{contract: _TNT20VoucherContract.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TNT20VoucherContract *TNT20VoucherContractFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TNT20VoucherContractApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TNT20VoucherContract.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TNT20VoucherContractApproval)
				if err := _TNT20VoucherContract.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_TNT20VoucherContract *TNT20VoucherContractFilterer) ParseApproval(log types.Log) (*TNT20VoucherContractApproval, error) {
	event := new(TNT20VoucherContractApproval)
	if err := _TNT20VoucherContract.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TNT20VoucherContractTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the TNT20VoucherContract contract.
type TNT20VoucherContractTransferIterator struct {
	Event *TNT20VoucherContractTransfer // Event containing the contract specifics and raw log

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
func (it *TNT20VoucherContractTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TNT20VoucherContractTransfer)
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
		it.Event = new(TNT20VoucherContractTransfer)
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
func (it *TNT20VoucherContractTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TNT20VoucherContractTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TNT20VoucherContractTransfer represents a Transfer event raised by the TNT20VoucherContract contract.
type TNT20VoucherContractTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TNT20VoucherContract *TNT20VoucherContractFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TNT20VoucherContractTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TNT20VoucherContract.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TNT20VoucherContractTransferIterator{contract: _TNT20VoucherContract.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TNT20VoucherContract *TNT20VoucherContractFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TNT20VoucherContractTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TNT20VoucherContract.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TNT20VoucherContractTransfer)
				if err := _TNT20VoucherContract.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_TNT20VoucherContract *TNT20VoucherContractFilterer) ParseTransfer(log types.Log) (*TNT20VoucherContractTransfer, error) {
	event := new(TNT20VoucherContractTransfer)
	if err := _TNT20VoucherContract.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TNT20VoucherContractUpdateOwnerIterator is returned from FilterUpdateOwner and is used to iterate over the raw logs and unpacked data for UpdateOwner events raised by the TNT20VoucherContract contract.
type TNT20VoucherContractUpdateOwnerIterator struct {
	Event *TNT20VoucherContractUpdateOwner // Event containing the contract specifics and raw log

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
func (it *TNT20VoucherContractUpdateOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TNT20VoucherContractUpdateOwner)
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
		it.Event = new(TNT20VoucherContractUpdateOwner)
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
func (it *TNT20VoucherContractUpdateOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TNT20VoucherContractUpdateOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TNT20VoucherContractUpdateOwner represents a UpdateOwner event raised by the TNT20VoucherContract contract.
type TNT20VoucherContractUpdateOwner struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdateOwner is a free log retrieval operation binding the contract event 0xe2c7d1c4da37855e682bde14f17826d185497973b73fba7554daa6da467058d9.
//
// Solidity: event UpdateOwner(address oldOwner, address newOwner)
func (_TNT20VoucherContract *TNT20VoucherContractFilterer) FilterUpdateOwner(opts *bind.FilterOpts) (*TNT20VoucherContractUpdateOwnerIterator, error) {

	logs, sub, err := _TNT20VoucherContract.contract.FilterLogs(opts, "UpdateOwner")
	if err != nil {
		return nil, err
	}
	return &TNT20VoucherContractUpdateOwnerIterator{contract: _TNT20VoucherContract.contract, event: "UpdateOwner", logs: logs, sub: sub}, nil
}

// WatchUpdateOwner is a free log subscription operation binding the contract event 0xe2c7d1c4da37855e682bde14f17826d185497973b73fba7554daa6da467058d9.
//
// Solidity: event UpdateOwner(address oldOwner, address newOwner)
func (_TNT20VoucherContract *TNT20VoucherContractFilterer) WatchUpdateOwner(opts *bind.WatchOpts, sink chan<- *TNT20VoucherContractUpdateOwner) (event.Subscription, error) {

	logs, sub, err := _TNT20VoucherContract.contract.WatchLogs(opts, "UpdateOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TNT20VoucherContractUpdateOwner)
				if err := _TNT20VoucherContract.contract.UnpackLog(event, "UpdateOwner", log); err != nil {
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

// ParseUpdateOwner is a log parse operation binding the contract event 0xe2c7d1c4da37855e682bde14f17826d185497973b73fba7554daa6da467058d9.
//
// Solidity: event UpdateOwner(address oldOwner, address newOwner)
func (_TNT20VoucherContract *TNT20VoucherContractFilterer) ParseUpdateOwner(log types.Log) (*TNT20VoucherContractUpdateOwner, error) {
	event := new(TNT20VoucherContractUpdateOwner)
	if err := _TNT20VoucherContract.contract.UnpackLog(event, "UpdateOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
