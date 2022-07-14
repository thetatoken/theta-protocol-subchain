// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package accessors

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

// SubchainGovernanceTokenMetaData contains all meta data concerning the SubchainGovernanceToken contract.
var SubchainGovernanceTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internaltype\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internaltype\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internaltype\":\"uint8\",\"name\":\"decimals_\",\"type\":\"uint8\"},{\"internaltype\":\"uint256\",\"name\":\"maxsupply_\",\"type\":\"uint256\"},{\"internaltype\":\"address\",\"name\":\"minter_\",\"type\":\"address\"},{\"internaltype\":\"uint256\",\"name\":\"stakerrewardperblock_\",\"type\":\"uint256\"},{\"internaltype\":\"address\",\"name\":\"initdistrwallet_\",\"type\":\"address\"},{\"internaltype\":\"uint256\",\"name\":\"initmintamount_\",\"type\":\"uint256\"},{\"internaltype\":\"address\",\"name\":\"admin_\",\"type\":\"address\"}],\"statemutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internaltype\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internaltype\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internaltype\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internaltype\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internaltype\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internaltype\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internaltype\":\"address\",\"name\":\"newadmin\",\"type\":\"address\"}],\"name\":\"updateadmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internaltype\":\"address\",\"name\":\"newminter\",\"type\":\"address\"}],\"name\":\"updateminter\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internaltype\":\"uint256\",\"name\":\"newstakerrewardperblock\",\"type\":\"uint256\"}],\"name\":\"updatestakerrewardperblock\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internaltype\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"statemutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internaltype\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internaltype\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"statemutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internaltype\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internaltype\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internaltype\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"statemutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceof\",\"outputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"statemutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internaltype\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internaltype\":\"uint256\",\"name\":\"subtractedvalue\",\"type\":\"uint256\"}],\"name\":\"decreaseallowance\",\"outputs\":[{\"internaltype\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"statemutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internaltype\":\"uint256\",\"name\":\"addedvalue\",\"type\":\"uint256\"}],\"name\":\"increaseallowance\",\"outputs\":[{\"internaltype\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"statemutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxsupply\",\"outputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"statemutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"minter\",\"outputs\":[{\"internaltype\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"statemutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internaltype\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"statemutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internaltype\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"statemutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"totalsupply\",\"outputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"statemutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internaltype\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internaltype\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internaltype\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"statemutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internaltype\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internaltype\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferfrom\",\"outputs\":[{\"internaltype\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"statemutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internaltype\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"statemutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internaltype\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internaltype\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mintstakerreward\",\"outputs\":[{\"internaltype\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"statemutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakerrewardperblock\",\"outputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"statemutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internaltype\":\"uint256\",\"name\":\"stakerrewardperblock_\",\"type\":\"uint256\"}],\"name\":\"updatestakerrewardperblock\",\"outputs\":[],\"statemutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"address\",\"name\":\"minter_\",\"type\":\"address\"}],\"name\":\"updateminter\",\"outputs\":[],\"statemutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"address\",\"name\":\"admin_\",\"type\":\"address\"}],\"name\":\"updateadmin\",\"outputs\":[],\"statemutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620013c0380380620013c08339810160408190526200003491620004b3565b8851899089906200004d90600390602085019062000339565b5080516200006390600490602084019062000339565b5050506c010000000000000000000000008610620000c85760405162461bcd60e51b815260206004820152601360248201527f6d6178537570706c7920746f6f206c617267650000000000000000000000000060448201526064015b60405180910390fd5b85821115620001325760405162461bcd60e51b815260206004820152602f60248201527f696e697469616c20737570706c792073686f756c64206e6f742065786365656460448201526e20746865206d617820737570706c7960881b6064820152608401620000bf565b6005805460ff891660ff199091161790556007869055600880546001600160a01b038088166001600160a01b0319928316179092556006869055600980549284169290911691909117905562000189838362000254565b6008546040516001600160a01b0390911681527f94a0a7c0a7a455351029a521046d4438bc63e58a1ee8d984d624eb3161583b2c9060200160405180910390a16009546040516001600160a01b0390911681527fbfc8d7754fec5096becc28e0816011e8d0adcfe752ffa1c88924a7f88b00a41d9060200160405180910390a17fd65d781db82f0bd55f732ea1b96f4c6565750df815c18a70962cd8405274e2436006546040516200023d91815260200190565b60405180910390a150505050505050505062000607565b6001600160a01b038216620002ac5760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152606401620000bf565b8060026000828254620002c091906200058d565b90915550506001600160a01b03821660009081526020819052604081208054839290620002ef9084906200058d565b90915550506040518181526001600160a01b038316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b8280546200034790620005b4565b90600052602060002090601f0160209004810192826200036b5760008555620003b6565b82601f106200038657805160ff1916838001178555620003b6565b82800160010185558215620003b6579182015b82811115620003b657825182559160200191906001019062000399565b50620003c4929150620003c8565b5090565b5b80821115620003c45760008155600101620003c9565b80516001600160a01b0381168114620003f757600080fd5b919050565b600082601f8301126200040e57600080fd5b81516001600160401b03808211156200042b576200042b620005f1565b604051601f8301601f19908116603f01168101908282118183101715620004565762000456620005f1565b816040528381526020925086838588010111156200047357600080fd5b600091505b8382101562000497578582018301518183018401529082019062000478565b83821115620004a95760008385830101525b9695505050505050565b60008060008060008060008060006101208a8c031215620004d357600080fd5b89516001600160401b0380821115620004eb57600080fd5b620004f98d838e01620003fc565b9a5060208c01519150808211156200051057600080fd5b506200051f8c828d01620003fc565b98505060408a015160ff811681146200053757600080fd5b60608b015190975095506200054f60808b01620003df565b945060a08a015193506200056660c08b01620003df565b925060e08a015191506200057e6101008b01620003df565b90509295985092959850929598565b60008219821115620005af57634e487b7160e01b600052601160045260246000fd5b500190565b600181811c90821680620005c957607f821691505b60208210811415620005eb57634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052604160045260246000fd5b610da980620006176000396000f3fe608060405234801561001057600080fd5b50600436106101215760003560e01c80634eb03f6e116100ad578063ab8e786f11610071578063ab8e786f14610266578063d5abeb0114610279578063dd62ed3e14610282578063e2f273bd14610295578063f851a440146102a857600080fd5b80634eb03f6e146101fc57806370a082311461020f57806395d89b4114610238578063a457c2d714610240578063a9059cbb1461025357600080fd5b806323b872dd116100f457806323b872dd146101a457806324e770d0146101b757806329fb7ef6146101bf578063313ce567146101d457806339509351146101e957600080fd5b806306fdde03146101265780630754617214610144578063095ea7b31461016f57806318160ddd14610192575b600080fd5b61012e6102bb565b60405161013b9190610c67565b60405180910390f35b600854610157906001600160a01b031681565b6040516001600160a01b03909116815260200161013b565b61018261017d366004610c0b565b61034d565b604051901515815260200161013b565b6002545b60405190815260200161013b565b6101826101b2366004610bcf565b610367565b600654610196565b6101d26101cd366004610c35565b61038b565b005b60055460405160ff909116815260200161013b565b6101826101f7366004610c0b565b6103fa565b6101d261020a366004610b81565b61041c565b61019661021d366004610b81565b6001600160a01b031660009081526020819052604090205490565b61012e610494565b61018261024e366004610c0b565b6104a3565b610182610261366004610c0b565b61051e565b610182610274366004610c0b565b61052c565b61019660075481565b610196610290366004610b9c565b610658565b6101d26102a3366004610b81565b610683565b600954610157906001600160a01b031681565b6060600380546102ca90610d22565b80601f01602080910402602001604051908101604052809291908181526020018280546102f690610d22565b80156103435780601f1061031857610100808354040283529160200191610343565b820191906000526020600020905b81548152906001019060200180831161032657829003601f168201915b5050505050905090565b60003361035b8185856106fb565b60019150505b92915050565b60003361037585828561081f565b610380858585610899565b506001949350505050565b6009546001600160a01b031633146103be5760405162461bcd60e51b81526004016103b590610cbc565b60405180910390fd5b60068190556040518181527fd65d781db82f0bd55f732ea1b96f4c6565750df815c18a70962cd8405274e243906020015b60405180910390a150565b60003361035b81858561040d8383610658565b6104179190610cf3565b6106fb565b6009546001600160a01b031633146104465760405162461bcd60e51b81526004016103b590610cbc565b600880546001600160a01b0319166001600160a01b0383169081179091556040519081527f94a0a7c0a7a455351029a521046d4438bc63e58a1ee8d984d624eb3161583b2c906020016103ef565b6060600480546102ca90610d22565b600033816104b18286610658565b9050838110156105115760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b60648201526084016103b5565b61038082868684036106fb565b60003361035b818585610899565b6008546000906001600160a01b031633146105895760405162461bcd60e51b815260206004820152601e60248201527f4f6e6c79206d696e7465722063616e206d616b6520746869732063616c6c000060448201526064016103b5565b6008546001600160a01b031633146105a357506000610361565b6000306001600160a01b03166318160ddd6040518163ffffffff1660e01b815260040160206040518083038186803b1580156105de57600080fd5b505afa1580156105f2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106169190610c4e565b9050600754811061062b576000915050610361565b6007546106388285610a67565b111561064e5760075461064b9082610a7a565b92505b61035b8484610a86565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6009546001600160a01b031633146106ad5760405162461bcd60e51b81526004016103b590610cbc565b600980546001600160a01b0319166001600160a01b0383169081179091556040519081527fbfc8d7754fec5096becc28e0816011e8d0adcfe752ffa1c88924a7f88b00a41d906020016103ef565b6001600160a01b03831661075d5760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b60648201526084016103b5565b6001600160a01b0382166107be5760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b60648201526084016103b5565b6001600160a01b0383811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b600061082b8484610658565b9050600019811461089357818110156108865760405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e636500000060448201526064016103b5565b61089384848484036106fb565b50505050565b6001600160a01b0383166108fd5760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b60648201526084016103b5565b6001600160a01b03821661095f5760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b60648201526084016103b5565b6001600160a01b038316600090815260208190526040902054818110156109d75760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b60648201526084016103b5565b6001600160a01b03808516600090815260208190526040808220858503905591851681529081208054849290610a0e908490610cf3565b92505081905550826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610a5a91815260200190565b60405180910390a3610893565b6000610a738284610cf3565b9392505050565b6000610a738284610d0b565b6001600160a01b038216610adc5760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f20616464726573730060448201526064016103b5565b8060026000828254610aee9190610cf3565b90915550506001600160a01b03821660009081526020819052604081208054839290610b1b908490610cf3565b90915550506040518181526001600160a01b038316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b80356001600160a01b0381168114610b7c57600080fd5b919050565b600060208284031215610b9357600080fd5b610a7382610b65565b60008060408385031215610baf57600080fd5b610bb883610b65565b9150610bc660208401610b65565b90509250929050565b600080600060608486031215610be457600080fd5b610bed84610b65565b9250610bfb60208501610b65565b9150604084013590509250925092565b60008060408385031215610c1e57600080fd5b610c2783610b65565b946020939093013593505050565b600060208284031215610c4757600080fd5b5035919050565b600060208284031215610c6057600080fd5b5051919050565b600060208083528351808285015260005b81811015610c9457858101830151858201604001528201610c78565b81811115610ca6576000604083870101525b50601f01601f1916929092016040019392505050565b6020808252601d908201527f4f6e6c792061646d696e2063616e206d616b6520746869732063616c6c000000604082015260600190565b60008219821115610d0657610d06610d5d565b500190565b600082821015610d1d57610d1d610d5d565b500390565b600181811c90821680610d3657607f821691505b60208210811415610d5757634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601160045260246000fdfea26469706673582212200d5799869cc21955452ba6f29be63f355d4150762fec30cb2fae8c16110a324364736f6c63430008070033",
}

// SubchainGovernanceTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use SubchainGovernanceTokenMetaData.ABI instead.
var SubchainGovernanceTokenABI = SubchainGovernanceTokenMetaData.ABI

// SubchainGovernanceTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SubchainGovernanceTokenMetaData.Bin instead.
var SubchainGovernanceTokenBin = SubchainGovernanceTokenMetaData.Bin

// DeploySubchainGovernanceToken deploys a new Ethereum contract, binding an instance of SubchainGovernanceToken to it.
func DeploySubchainGovernanceToken(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ string, decimals_ uint8, maxsupply_ *big.Int, minter_ common.Address, stakerrewardperblock_ *big.Int, initdistrwallet_ common.Address, initmintamount_ *big.Int, admin_ common.Address) (common.Address, *types.Transaction, *SubchainGovernanceToken, error) {
	parsed, err := SubchainGovernanceTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SubchainGovernanceTokenBin), backend, name_, symbol_, decimals_, maxsupply_, minter_, stakerrewardperblock_, initdistrwallet_, initmintamount_, admin_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SubchainGovernanceToken{SubchainGovernanceTokenCaller: SubchainGovernanceTokenCaller{contract: contract}, SubchainGovernanceTokenTransactor: SubchainGovernanceTokenTransactor{contract: contract}, SubchainGovernanceTokenFilterer: SubchainGovernanceTokenFilterer{contract: contract}}, nil
}

// SubchainGovernanceToken is an auto generated Go binding around an Ethereum contract.
type SubchainGovernanceToken struct {
	SubchainGovernanceTokenCaller     // Read-only binding to the contract
	SubchainGovernanceTokenTransactor // Write-only binding to the contract
	SubchainGovernanceTokenFilterer   // Log filterer for contract events
}

// SubchainGovernanceTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type SubchainGovernanceTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubchainGovernanceTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SubchainGovernanceTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubchainGovernanceTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SubchainGovernanceTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubchainGovernanceTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SubchainGovernanceTokenSession struct {
	Contract     *SubchainGovernanceToken // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SubchainGovernanceTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SubchainGovernanceTokenCallerSession struct {
	Contract *SubchainGovernanceTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// SubchainGovernanceTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SubchainGovernanceTokenTransactorSession struct {
	Contract     *SubchainGovernanceTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// SubchainGovernanceTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type SubchainGovernanceTokenRaw struct {
	Contract *SubchainGovernanceToken // Generic contract binding to access the raw methods on
}

// SubchainGovernanceTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SubchainGovernanceTokenCallerRaw struct {
	Contract *SubchainGovernanceTokenCaller // Generic read-only contract binding to access the raw methods on
}

// SubchainGovernanceTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SubchainGovernanceTokenTransactorRaw struct {
	Contract *SubchainGovernanceTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSubchainGovernanceToken creates a new instance of SubchainGovernanceToken, bound to a specific deployed contract.
func NewSubchainGovernanceToken(address common.Address, backend bind.ContractBackend) (*SubchainGovernanceToken, error) {
	contract, err := bindSubchainGovernanceToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SubchainGovernanceToken{SubchainGovernanceTokenCaller: SubchainGovernanceTokenCaller{contract: contract}, SubchainGovernanceTokenTransactor: SubchainGovernanceTokenTransactor{contract: contract}, SubchainGovernanceTokenFilterer: SubchainGovernanceTokenFilterer{contract: contract}}, nil
}

// NewSubchainGovernanceTokenCaller creates a new read-only instance of SubchainGovernanceToken, bound to a specific deployed contract.
func NewSubchainGovernanceTokenCaller(address common.Address, caller bind.ContractCaller) (*SubchainGovernanceTokenCaller, error) {
	contract, err := bindSubchainGovernanceToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SubchainGovernanceTokenCaller{contract: contract}, nil
}

// NewSubchainGovernanceTokenTransactor creates a new write-only instance of SubchainGovernanceToken, bound to a specific deployed contract.
func NewSubchainGovernanceTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*SubchainGovernanceTokenTransactor, error) {
	contract, err := bindSubchainGovernanceToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SubchainGovernanceTokenTransactor{contract: contract}, nil
}

// NewSubchainGovernanceTokenFilterer creates a new log filterer instance of SubchainGovernanceToken, bound to a specific deployed contract.
func NewSubchainGovernanceTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*SubchainGovernanceTokenFilterer, error) {
	contract, err := bindSubchainGovernanceToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SubchainGovernanceTokenFilterer{contract: contract}, nil
}

// bindSubchainGovernanceToken binds a generic wrapper to an already deployed contract.
func bindSubchainGovernanceToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SubchainGovernanceTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SubchainGovernanceToken *SubchainGovernanceTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SubchainGovernanceToken.Contract.SubchainGovernanceTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SubchainGovernanceToken *SubchainGovernanceTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SubchainGovernanceToken.Contract.SubchainGovernanceTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SubchainGovernanceToken *SubchainGovernanceTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SubchainGovernanceToken.Contract.SubchainGovernanceTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SubchainGovernanceToken *SubchainGovernanceTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SubchainGovernanceToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SubchainGovernanceToken *SubchainGovernanceTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SubchainGovernanceToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SubchainGovernanceToken *SubchainGovernanceTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SubchainGovernanceToken.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_SubchainGovernanceToken *SubchainGovernanceTokenCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SubchainGovernanceToken.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_SubchainGovernanceToken *SubchainGovernanceTokenSession) Admin() (common.Address, error) {
	return _SubchainGovernanceToken.Contract.Admin(&_SubchainGovernanceToken.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_SubchainGovernanceToken *SubchainGovernanceTokenCallerSession) Admin() (common.Address, error) {
	return _SubchainGovernanceToken.Contract.Admin(&_SubchainGovernanceToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_SubchainGovernanceToken *SubchainGovernanceTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SubchainGovernanceToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_SubchainGovernanceToken *SubchainGovernanceTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _SubchainGovernanceToken.Contract.Allowance(&_SubchainGovernanceToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_SubchainGovernanceToken *SubchainGovernanceTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _SubchainGovernanceToken.Contract.Allowance(&_SubchainGovernanceToken.CallOpts, owner, spender)
}

// Balanceof is a free data retrieval call binding the contract method 0x3d64125b.
//
// Solidity: function balanceof(address account) view returns(uint256)
func (_SubchainGovernanceToken *SubchainGovernanceTokenCaller) Balanceof(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SubchainGovernanceToken.contract.Call(opts, &out, "balanceof", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balanceof is a free data retrieval call binding the contract method 0x3d64125b.
//
// Solidity: function balanceof(address account) view returns(uint256)
func (_SubchainGovernanceToken *SubchainGovernanceTokenSession) Balanceof(account common.Address) (*big.Int, error) {
	return _SubchainGovernanceToken.Contract.Balanceof(&_SubchainGovernanceToken.CallOpts, account)
}

// Balanceof is a free data retrieval call binding the contract method 0x3d64125b.
//
// Solidity: function balanceof(address account) view returns(uint256)
func (_SubchainGovernanceToken *SubchainGovernanceTokenCallerSession) Balanceof(account common.Address) (*big.Int, error) {
	return _SubchainGovernanceToken.Contract.Balanceof(&_SubchainGovernanceToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SubchainGovernanceToken *SubchainGovernanceTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SubchainGovernanceToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SubchainGovernanceToken *SubchainGovernanceTokenSession) Decimals() (uint8, error) {
	return _SubchainGovernanceToken.Contract.Decimals(&_SubchainGovernanceToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SubchainGovernanceToken *SubchainGovernanceTokenCallerSession) Decimals() (uint8, error) {
	return _SubchainGovernanceToken.Contract.Decimals(&_SubchainGovernanceToken.CallOpts)
}

// Maxsupply is a free data retrieval call binding the contract method 0x4b6406d1.
//
// Solidity: function maxsupply() view returns(uint256)
func (_SubchainGovernanceToken *SubchainGovernanceTokenCaller) Maxsupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SubchainGovernanceToken.contract.Call(opts, &out, "maxsupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Maxsupply is a free data retrieval call binding the contract method 0x4b6406d1.
//
// Solidity: function maxsupply() view returns(uint256)
func (_SubchainGovernanceToken *SubchainGovernanceTokenSession) Maxsupply() (*big.Int, error) {
	return _SubchainGovernanceToken.Contract.Maxsupply(&_SubchainGovernanceToken.CallOpts)
}

// Maxsupply is a free data retrieval call binding the contract method 0x4b6406d1.
//
// Solidity: function maxsupply() view returns(uint256)
func (_SubchainGovernanceToken *SubchainGovernanceTokenCallerSession) Maxsupply() (*big.Int, error) {
	return _SubchainGovernanceToken.Contract.Maxsupply(&_SubchainGovernanceToken.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_SubchainGovernanceToken *SubchainGovernanceTokenCaller) Minter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SubchainGovernanceToken.contract.Call(opts, &out, "minter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_SubchainGovernanceToken *SubchainGovernanceTokenSession) Minter() (common.Address, error) {
	return _SubchainGovernanceToken.Contract.Minter(&_SubchainGovernanceToken.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_SubchainGovernanceToken *SubchainGovernanceTokenCallerSession) Minter() (common.Address, error) {
	return _SubchainGovernanceToken.Contract.Minter(&_SubchainGovernanceToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SubchainGovernanceToken *SubchainGovernanceTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SubchainGovernanceToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SubchainGovernanceToken *SubchainGovernanceTokenSession) Name() (string, error) {
	return _SubchainGovernanceToken.Contract.Name(&_SubchainGovernanceToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SubchainGovernanceToken *SubchainGovernanceTokenCallerSession) Name() (string, error) {
	return _SubchainGovernanceToken.Contract.Name(&_SubchainGovernanceToken.CallOpts)
}

// Stakerrewardperblock is a free data retrieval call binding the contract method 0x4aade615.
//
// Solidity: function stakerrewardperblock() view returns(uint256)
func (_SubchainGovernanceToken *SubchainGovernanceTokenCaller) Stakerrewardperblock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SubchainGovernanceToken.contract.Call(opts, &out, "stakerrewardperblock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Stakerrewardperblock is a free data retrieval call binding the contract method 0x4aade615.
//
// Solidity: function stakerrewardperblock() view returns(uint256)
func (_SubchainGovernanceToken *SubchainGovernanceTokenSession) Stakerrewardperblock() (*big.Int, error) {
	return _SubchainGovernanceToken.Contract.Stakerrewardperblock(&_SubchainGovernanceToken.CallOpts)
}

// Stakerrewardperblock is a free data retrieval call binding the contract method 0x4aade615.
//
// Solidity: function stakerrewardperblock() view returns(uint256)
func (_SubchainGovernanceToken *SubchainGovernanceTokenCallerSession) Stakerrewardperblock() (*big.Int, error) {
	return _SubchainGovernanceToken.Contract.Stakerrewardperblock(&_SubchainGovernanceToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SubchainGovernanceToken *SubchainGovernanceTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SubchainGovernanceToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SubchainGovernanceToken *SubchainGovernanceTokenSession) Symbol() (string, error) {
	return _SubchainGovernanceToken.Contract.Symbol(&_SubchainGovernanceToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SubchainGovernanceToken *SubchainGovernanceTokenCallerSession) Symbol() (string, error) {
	return _SubchainGovernanceToken.Contract.Symbol(&_SubchainGovernanceToken.CallOpts)
}

// Totalsupply is a free data retrieval call binding the contract method 0x72dd529b.
//
// Solidity: function totalsupply() view returns(uint256)
func (_SubchainGovernanceToken *SubchainGovernanceTokenCaller) Totalsupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SubchainGovernanceToken.contract.Call(opts, &out, "totalsupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Totalsupply is a free data retrieval call binding the contract method 0x72dd529b.
//
// Solidity: function totalsupply() view returns(uint256)
func (_SubchainGovernanceToken *SubchainGovernanceTokenSession) Totalsupply() (*big.Int, error) {
	return _SubchainGovernanceToken.Contract.Totalsupply(&_SubchainGovernanceToken.CallOpts)
}

// Totalsupply is a free data retrieval call binding the contract method 0x72dd529b.
//
// Solidity: function totalsupply() view returns(uint256)
func (_SubchainGovernanceToken *SubchainGovernanceTokenCallerSession) Totalsupply() (*big.Int, error) {
	return _SubchainGovernanceToken.Contract.Totalsupply(&_SubchainGovernanceToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_SubchainGovernanceToken *SubchainGovernanceTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SubchainGovernanceToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_SubchainGovernanceToken *SubchainGovernanceTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SubchainGovernanceToken.Contract.Approve(&_SubchainGovernanceToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_SubchainGovernanceToken *SubchainGovernanceTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SubchainGovernanceToken.Contract.Approve(&_SubchainGovernanceToken.TransactOpts, spender, amount)
}

// Decreaseallowance is a paid mutator transaction binding the contract method 0x7c179d92.
//
// Solidity: function decreaseallowance(address spender, uint256 subtractedvalue) returns(bool)
func (_SubchainGovernanceToken *SubchainGovernanceTokenTransactor) Decreaseallowance(opts *bind.TransactOpts, spender common.Address, subtractedvalue *big.Int) (*types.Transaction, error) {
	return _SubchainGovernanceToken.contract.Transact(opts, "decreaseallowance", spender, subtractedvalue)
}

// Decreaseallowance is a paid mutator transaction binding the contract method 0x7c179d92.
//
// Solidity: function decreaseallowance(address spender, uint256 subtractedvalue) returns(bool)
func (_SubchainGovernanceToken *SubchainGovernanceTokenSession) Decreaseallowance(spender common.Address, subtractedvalue *big.Int) (*types.Transaction, error) {
	return _SubchainGovernanceToken.Contract.Decreaseallowance(&_SubchainGovernanceToken.TransactOpts, spender, subtractedvalue)
}

// Decreaseallowance is a paid mutator transaction binding the contract method 0x7c179d92.
//
// Solidity: function decreaseallowance(address spender, uint256 subtractedvalue) returns(bool)
func (_SubchainGovernanceToken *SubchainGovernanceTokenTransactorSession) Decreaseallowance(spender common.Address, subtractedvalue *big.Int) (*types.Transaction, error) {
	return _SubchainGovernanceToken.Contract.Decreaseallowance(&_SubchainGovernanceToken.TransactOpts, spender, subtractedvalue)
}

// Increaseallowance is a paid mutator transaction binding the contract method 0x4d3eaaa8.
//
// Solidity: function increaseallowance(address spender, uint256 addedvalue) returns(bool)
func (_SubchainGovernanceToken *SubchainGovernanceTokenTransactor) Increaseallowance(opts *bind.TransactOpts, spender common.Address, addedvalue *big.Int) (*types.Transaction, error) {
	return _SubchainGovernanceToken.contract.Transact(opts, "increaseallowance", spender, addedvalue)
}

// Increaseallowance is a paid mutator transaction binding the contract method 0x4d3eaaa8.
//
// Solidity: function increaseallowance(address spender, uint256 addedvalue) returns(bool)
func (_SubchainGovernanceToken *SubchainGovernanceTokenSession) Increaseallowance(spender common.Address, addedvalue *big.Int) (*types.Transaction, error) {
	return _SubchainGovernanceToken.Contract.Increaseallowance(&_SubchainGovernanceToken.TransactOpts, spender, addedvalue)
}

// Increaseallowance is a paid mutator transaction binding the contract method 0x4d3eaaa8.
//
// Solidity: function increaseallowance(address spender, uint256 addedvalue) returns(bool)
func (_SubchainGovernanceToken *SubchainGovernanceTokenTransactorSession) Increaseallowance(spender common.Address, addedvalue *big.Int) (*types.Transaction, error) {
	return _SubchainGovernanceToken.Contract.Increaseallowance(&_SubchainGovernanceToken.TransactOpts, spender, addedvalue)
}

// Mintstakerreward is a paid mutator transaction binding the contract method 0xb55f0767.
//
// Solidity: function mintstakerreward(address account, uint256 amount) returns(bool)
func (_SubchainGovernanceToken *SubchainGovernanceTokenTransactor) Mintstakerreward(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SubchainGovernanceToken.contract.Transact(opts, "mintstakerreward", account, amount)
}

// Mintstakerreward is a paid mutator transaction binding the contract method 0xb55f0767.
//
// Solidity: function mintstakerreward(address account, uint256 amount) returns(bool)
func (_SubchainGovernanceToken *SubchainGovernanceTokenSession) Mintstakerreward(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SubchainGovernanceToken.Contract.Mintstakerreward(&_SubchainGovernanceToken.TransactOpts, account, amount)
}

// Mintstakerreward is a paid mutator transaction binding the contract method 0xb55f0767.
//
// Solidity: function mintstakerreward(address account, uint256 amount) returns(bool)
func (_SubchainGovernanceToken *SubchainGovernanceTokenTransactorSession) Mintstakerreward(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SubchainGovernanceToken.Contract.Mintstakerreward(&_SubchainGovernanceToken.TransactOpts, account, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_SubchainGovernanceToken *SubchainGovernanceTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SubchainGovernanceToken.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_SubchainGovernanceToken *SubchainGovernanceTokenSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SubchainGovernanceToken.Contract.Transfer(&_SubchainGovernanceToken.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_SubchainGovernanceToken *SubchainGovernanceTokenTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SubchainGovernanceToken.Contract.Transfer(&_SubchainGovernanceToken.TransactOpts, to, amount)
}

// Transferfrom is a paid mutator transaction binding the contract method 0x63b0545f.
//
// Solidity: function transferfrom(address from, address to, uint256 amount) returns(bool)
func (_SubchainGovernanceToken *SubchainGovernanceTokenTransactor) Transferfrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SubchainGovernanceToken.contract.Transact(opts, "transferfrom", from, to, amount)
}

// Transferfrom is a paid mutator transaction binding the contract method 0x63b0545f.
//
// Solidity: function transferfrom(address from, address to, uint256 amount) returns(bool)
func (_SubchainGovernanceToken *SubchainGovernanceTokenSession) Transferfrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SubchainGovernanceToken.Contract.Transferfrom(&_SubchainGovernanceToken.TransactOpts, from, to, amount)
}

// Transferfrom is a paid mutator transaction binding the contract method 0x63b0545f.
//
// Solidity: function transferfrom(address from, address to, uint256 amount) returns(bool)
func (_SubchainGovernanceToken *SubchainGovernanceTokenTransactorSession) Transferfrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SubchainGovernanceToken.Contract.Transferfrom(&_SubchainGovernanceToken.TransactOpts, from, to, amount)
}

// Updateadmin is a paid mutator transaction binding the contract method 0x47fd8035.
//
// Solidity: function updateadmin(address admin_) returns()
func (_SubchainGovernanceToken *SubchainGovernanceTokenTransactor) Updateadmin(opts *bind.TransactOpts, admin_ common.Address) (*types.Transaction, error) {
	return _SubchainGovernanceToken.contract.Transact(opts, "updateadmin", admin_)
}

// Updateadmin is a paid mutator transaction binding the contract method 0x47fd8035.
//
// Solidity: function updateadmin(address admin_) returns()
func (_SubchainGovernanceToken *SubchainGovernanceTokenSession) Updateadmin(admin_ common.Address) (*types.Transaction, error) {
	return _SubchainGovernanceToken.Contract.Updateadmin(&_SubchainGovernanceToken.TransactOpts, admin_)
}

// Updateadmin is a paid mutator transaction binding the contract method 0x47fd8035.
//
// Solidity: function updateadmin(address admin_) returns()
func (_SubchainGovernanceToken *SubchainGovernanceTokenTransactorSession) Updateadmin(admin_ common.Address) (*types.Transaction, error) {
	return _SubchainGovernanceToken.Contract.Updateadmin(&_SubchainGovernanceToken.TransactOpts, admin_)
}

// Updateminter is a paid mutator transaction binding the contract method 0x454c76be.
//
// Solidity: function updateminter(address minter_) returns()
func (_SubchainGovernanceToken *SubchainGovernanceTokenTransactor) Updateminter(opts *bind.TransactOpts, minter_ common.Address) (*types.Transaction, error) {
	return _SubchainGovernanceToken.contract.Transact(opts, "updateminter", minter_)
}

// Updateminter is a paid mutator transaction binding the contract method 0x454c76be.
//
// Solidity: function updateminter(address minter_) returns()
func (_SubchainGovernanceToken *SubchainGovernanceTokenSession) Updateminter(minter_ common.Address) (*types.Transaction, error) {
	return _SubchainGovernanceToken.Contract.Updateminter(&_SubchainGovernanceToken.TransactOpts, minter_)
}

// Updateminter is a paid mutator transaction binding the contract method 0x454c76be.
//
// Solidity: function updateminter(address minter_) returns()
func (_SubchainGovernanceToken *SubchainGovernanceTokenTransactorSession) Updateminter(minter_ common.Address) (*types.Transaction, error) {
	return _SubchainGovernanceToken.Contract.Updateminter(&_SubchainGovernanceToken.TransactOpts, minter_)
}

// Updatestakerrewardperblock is a paid mutator transaction binding the contract method 0x788b35af.
//
// Solidity: function updatestakerrewardperblock(uint256 stakerrewardperblock_) returns()
func (_SubchainGovernanceToken *SubchainGovernanceTokenTransactor) Updatestakerrewardperblock(opts *bind.TransactOpts, stakerrewardperblock_ *big.Int) (*types.Transaction, error) {
	return _SubchainGovernanceToken.contract.Transact(opts, "updatestakerrewardperblock", stakerrewardperblock_)
}

// Updatestakerrewardperblock is a paid mutator transaction binding the contract method 0x788b35af.
//
// Solidity: function updatestakerrewardperblock(uint256 stakerrewardperblock_) returns()
func (_SubchainGovernanceToken *SubchainGovernanceTokenSession) Updatestakerrewardperblock(stakerrewardperblock_ *big.Int) (*types.Transaction, error) {
	return _SubchainGovernanceToken.Contract.Updatestakerrewardperblock(&_SubchainGovernanceToken.TransactOpts, stakerrewardperblock_)
}

// Updatestakerrewardperblock is a paid mutator transaction binding the contract method 0x788b35af.
//
// Solidity: function updatestakerrewardperblock(uint256 stakerrewardperblock_) returns()
func (_SubchainGovernanceToken *SubchainGovernanceTokenTransactorSession) Updatestakerrewardperblock(stakerrewardperblock_ *big.Int) (*types.Transaction, error) {
	return _SubchainGovernanceToken.Contract.Updatestakerrewardperblock(&_SubchainGovernanceToken.TransactOpts, stakerrewardperblock_)
}

// SubchainGovernanceTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the SubchainGovernanceToken contract.
type SubchainGovernanceTokenApprovalIterator struct {
	Event *SubchainGovernanceTokenApproval // Event containing the contract specifics and raw log

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
func (it *SubchainGovernanceTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubchainGovernanceTokenApproval)
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
		it.Event = new(SubchainGovernanceTokenApproval)
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
func (it *SubchainGovernanceTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubchainGovernanceTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubchainGovernanceTokenApproval represents a Approval event raised by the SubchainGovernanceToken contract.
type SubchainGovernanceTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x5c52a5f2b86fd16be577188b5a83ef1165faddc00b137b10285f16162e17792a.
//
// Solidity: event approval(address indexed owner, address indexed spender, uint256 value)
func (_SubchainGovernanceToken *SubchainGovernanceTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*SubchainGovernanceTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SubchainGovernanceToken.contract.FilterLogs(opts, "approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &SubchainGovernanceTokenApprovalIterator{contract: _SubchainGovernanceToken.contract, event: "approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x5c52a5f2b86fd16be577188b5a83ef1165faddc00b137b10285f16162e17792a.
//
// Solidity: event approval(address indexed owner, address indexed spender, uint256 value)
func (_SubchainGovernanceToken *SubchainGovernanceTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SubchainGovernanceTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SubchainGovernanceToken.contract.WatchLogs(opts, "approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubchainGovernanceTokenApproval)
				if err := _SubchainGovernanceToken.contract.UnpackLog(event, "approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x5c52a5f2b86fd16be577188b5a83ef1165faddc00b137b10285f16162e17792a.
//
// Solidity: event approval(address indexed owner, address indexed spender, uint256 value)
func (_SubchainGovernanceToken *SubchainGovernanceTokenFilterer) ParseApproval(log types.Log) (*SubchainGovernanceTokenApproval, error) {
	event := new(SubchainGovernanceTokenApproval)
	if err := _SubchainGovernanceToken.contract.UnpackLog(event, "approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubchainGovernanceTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the SubchainGovernanceToken contract.
type SubchainGovernanceTokenTransferIterator struct {
	Event *SubchainGovernanceTokenTransfer // Event containing the contract specifics and raw log

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
func (it *SubchainGovernanceTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubchainGovernanceTokenTransfer)
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
		it.Event = new(SubchainGovernanceTokenTransfer)
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
func (it *SubchainGovernanceTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubchainGovernanceTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubchainGovernanceTokenTransfer represents a Transfer event raised by the SubchainGovernanceToken contract.
type SubchainGovernanceTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xbeabacc8ffedac16e9a60acdb2ca743d80c2ebb44977a93fa8e483c74d2b35a8.
//
// Solidity: event transfer(address indexed from, address indexed to, uint256 value)
func (_SubchainGovernanceToken *SubchainGovernanceTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SubchainGovernanceTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SubchainGovernanceToken.contract.FilterLogs(opts, "transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SubchainGovernanceTokenTransferIterator{contract: _SubchainGovernanceToken.contract, event: "transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xbeabacc8ffedac16e9a60acdb2ca743d80c2ebb44977a93fa8e483c74d2b35a8.
//
// Solidity: event transfer(address indexed from, address indexed to, uint256 value)
func (_SubchainGovernanceToken *SubchainGovernanceTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SubchainGovernanceTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SubchainGovernanceToken.contract.WatchLogs(opts, "transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubchainGovernanceTokenTransfer)
				if err := _SubchainGovernanceToken.contract.UnpackLog(event, "transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xbeabacc8ffedac16e9a60acdb2ca743d80c2ebb44977a93fa8e483c74d2b35a8.
//
// Solidity: event transfer(address indexed from, address indexed to, uint256 value)
func (_SubchainGovernanceToken *SubchainGovernanceTokenFilterer) ParseTransfer(log types.Log) (*SubchainGovernanceTokenTransfer, error) {
	event := new(SubchainGovernanceTokenTransfer)
	if err := _SubchainGovernanceToken.contract.UnpackLog(event, "transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubchainGovernanceTokenUpdateadminIterator is returned from FilterUpdateadmin and is used to iterate over the raw logs and unpacked data for Updateadmin events raised by the SubchainGovernanceToken contract.
type SubchainGovernanceTokenUpdateadminIterator struct {
	Event *SubchainGovernanceTokenUpdateadmin // Event containing the contract specifics and raw log

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
func (it *SubchainGovernanceTokenUpdateadminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubchainGovernanceTokenUpdateadmin)
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
		it.Event = new(SubchainGovernanceTokenUpdateadmin)
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
func (it *SubchainGovernanceTokenUpdateadminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubchainGovernanceTokenUpdateadminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubchainGovernanceTokenUpdateadmin represents a Updateadmin event raised by the SubchainGovernanceToken contract.
type SubchainGovernanceTokenUpdateadmin struct {
	Newadmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdateadmin is a free log retrieval operation binding the contract event 0x47fd8035c198d5e177f5dc8bc0abb26ddac92ae03c8830691f0b2e67fa126250.
//
// Solidity: event updateadmin(address newadmin)
func (_SubchainGovernanceToken *SubchainGovernanceTokenFilterer) FilterUpdateadmin(opts *bind.FilterOpts) (*SubchainGovernanceTokenUpdateadminIterator, error) {

	logs, sub, err := _SubchainGovernanceToken.contract.FilterLogs(opts, "updateadmin")
	if err != nil {
		return nil, err
	}
	return &SubchainGovernanceTokenUpdateadminIterator{contract: _SubchainGovernanceToken.contract, event: "updateadmin", logs: logs, sub: sub}, nil
}

// WatchUpdateadmin is a free log subscription operation binding the contract event 0x47fd8035c198d5e177f5dc8bc0abb26ddac92ae03c8830691f0b2e67fa126250.
//
// Solidity: event updateadmin(address newadmin)
func (_SubchainGovernanceToken *SubchainGovernanceTokenFilterer) WatchUpdateadmin(opts *bind.WatchOpts, sink chan<- *SubchainGovernanceTokenUpdateadmin) (event.Subscription, error) {

	logs, sub, err := _SubchainGovernanceToken.contract.WatchLogs(opts, "updateadmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubchainGovernanceTokenUpdateadmin)
				if err := _SubchainGovernanceToken.contract.UnpackLog(event, "updateadmin", log); err != nil {
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

// ParseUpdateadmin is a log parse operation binding the contract event 0x47fd8035c198d5e177f5dc8bc0abb26ddac92ae03c8830691f0b2e67fa126250.
//
// Solidity: event updateadmin(address newadmin)
func (_SubchainGovernanceToken *SubchainGovernanceTokenFilterer) ParseUpdateadmin(log types.Log) (*SubchainGovernanceTokenUpdateadmin, error) {
	event := new(SubchainGovernanceTokenUpdateadmin)
	if err := _SubchainGovernanceToken.contract.UnpackLog(event, "updateadmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubchainGovernanceTokenUpdateminterIterator is returned from FilterUpdateminter and is used to iterate over the raw logs and unpacked data for Updateminter events raised by the SubchainGovernanceToken contract.
type SubchainGovernanceTokenUpdateminterIterator struct {
	Event *SubchainGovernanceTokenUpdateminter // Event containing the contract specifics and raw log

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
func (it *SubchainGovernanceTokenUpdateminterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubchainGovernanceTokenUpdateminter)
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
		it.Event = new(SubchainGovernanceTokenUpdateminter)
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
func (it *SubchainGovernanceTokenUpdateminterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubchainGovernanceTokenUpdateminterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubchainGovernanceTokenUpdateminter represents a Updateminter event raised by the SubchainGovernanceToken contract.
type SubchainGovernanceTokenUpdateminter struct {
	Newminter common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdateminter is a free log retrieval operation binding the contract event 0x454c76beeeae1bd18ef6ca90f8bb38bf662066f775ccd172eab022da15960c72.
//
// Solidity: event updateminter(address newminter)
func (_SubchainGovernanceToken *SubchainGovernanceTokenFilterer) FilterUpdateminter(opts *bind.FilterOpts) (*SubchainGovernanceTokenUpdateminterIterator, error) {

	logs, sub, err := _SubchainGovernanceToken.contract.FilterLogs(opts, "updateminter")
	if err != nil {
		return nil, err
	}
	return &SubchainGovernanceTokenUpdateminterIterator{contract: _SubchainGovernanceToken.contract, event: "updateminter", logs: logs, sub: sub}, nil
}

// WatchUpdateminter is a free log subscription operation binding the contract event 0x454c76beeeae1bd18ef6ca90f8bb38bf662066f775ccd172eab022da15960c72.
//
// Solidity: event updateminter(address newminter)
func (_SubchainGovernanceToken *SubchainGovernanceTokenFilterer) WatchUpdateminter(opts *bind.WatchOpts, sink chan<- *SubchainGovernanceTokenUpdateminter) (event.Subscription, error) {

	logs, sub, err := _SubchainGovernanceToken.contract.WatchLogs(opts, "updateminter")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubchainGovernanceTokenUpdateminter)
				if err := _SubchainGovernanceToken.contract.UnpackLog(event, "updateminter", log); err != nil {
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

// ParseUpdateminter is a log parse operation binding the contract event 0x454c76beeeae1bd18ef6ca90f8bb38bf662066f775ccd172eab022da15960c72.
//
// Solidity: event updateminter(address newminter)
func (_SubchainGovernanceToken *SubchainGovernanceTokenFilterer) ParseUpdateminter(log types.Log) (*SubchainGovernanceTokenUpdateminter, error) {
	event := new(SubchainGovernanceTokenUpdateminter)
	if err := _SubchainGovernanceToken.contract.UnpackLog(event, "updateminter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubchainGovernanceTokenUpdatestakerrewardperblockIterator is returned from FilterUpdatestakerrewardperblock and is used to iterate over the raw logs and unpacked data for Updatestakerrewardperblock events raised by the SubchainGovernanceToken contract.
type SubchainGovernanceTokenUpdatestakerrewardperblockIterator struct {
	Event *SubchainGovernanceTokenUpdatestakerrewardperblock // Event containing the contract specifics and raw log

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
func (it *SubchainGovernanceTokenUpdatestakerrewardperblockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubchainGovernanceTokenUpdatestakerrewardperblock)
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
		it.Event = new(SubchainGovernanceTokenUpdatestakerrewardperblock)
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
func (it *SubchainGovernanceTokenUpdatestakerrewardperblockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubchainGovernanceTokenUpdatestakerrewardperblockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubchainGovernanceTokenUpdatestakerrewardperblock represents a Updatestakerrewardperblock event raised by the SubchainGovernanceToken contract.
type SubchainGovernanceTokenUpdatestakerrewardperblock struct {
	Newstakerrewardperblock *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterUpdatestakerrewardperblock is a free log retrieval operation binding the contract event 0x788b35af76927a343eb5b1ac1390916980ee421d99d204801c5b3bf4469b3e28.
//
// Solidity: event updatestakerrewardperblock(uint256 newstakerrewardperblock)
func (_SubchainGovernanceToken *SubchainGovernanceTokenFilterer) FilterUpdatestakerrewardperblock(opts *bind.FilterOpts) (*SubchainGovernanceTokenUpdatestakerrewardperblockIterator, error) {

	logs, sub, err := _SubchainGovernanceToken.contract.FilterLogs(opts, "updatestakerrewardperblock")
	if err != nil {
		return nil, err
	}
	return &SubchainGovernanceTokenUpdatestakerrewardperblockIterator{contract: _SubchainGovernanceToken.contract, event: "updatestakerrewardperblock", logs: logs, sub: sub}, nil
}

// WatchUpdatestakerrewardperblock is a free log subscription operation binding the contract event 0x788b35af76927a343eb5b1ac1390916980ee421d99d204801c5b3bf4469b3e28.
//
// Solidity: event updatestakerrewardperblock(uint256 newstakerrewardperblock)
func (_SubchainGovernanceToken *SubchainGovernanceTokenFilterer) WatchUpdatestakerrewardperblock(opts *bind.WatchOpts, sink chan<- *SubchainGovernanceTokenUpdatestakerrewardperblock) (event.Subscription, error) {

	logs, sub, err := _SubchainGovernanceToken.contract.WatchLogs(opts, "updatestakerrewardperblock")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubchainGovernanceTokenUpdatestakerrewardperblock)
				if err := _SubchainGovernanceToken.contract.UnpackLog(event, "updatestakerrewardperblock", log); err != nil {
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

// ParseUpdatestakerrewardperblock is a log parse operation binding the contract event 0x788b35af76927a343eb5b1ac1390916980ee421d99d204801c5b3bf4469b3e28.
//
// Solidity: event updatestakerrewardperblock(uint256 newstakerrewardperblock)
func (_SubchainGovernanceToken *SubchainGovernanceTokenFilterer) ParseUpdatestakerrewardperblock(log types.Log) (*SubchainGovernanceTokenUpdatestakerrewardperblock, error) {
	event := new(SubchainGovernanceTokenUpdatestakerrewardperblock)
	if err := _SubchainGovernanceToken.contract.UnpackLog(event, "updatestakerrewardperblock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
