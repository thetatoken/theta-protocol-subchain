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

// SubchainGovernanceTokenMetaData contains all meta data concerning the SubchainGovernanceToken contract.
var SubchainGovernanceTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxSupply_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"minter_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stakerRewardPerBlock_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"initDistrWallet_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initMintAmount_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"admin_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"UpdateAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newMinter\",\"type\":\"address\"}],\"name\":\"UpdateMinter\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newStakerRewardPerBlock\",\"type\":\"uint256\"}],\"name\":\"UpdateStakerRewardPerBlock\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"minter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mintStakerReward\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakerRewardPerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakerRewardPerBlock_\",\"type\":\"uint256\"}],\"name\":\"updateStakerRewardPerBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter_\",\"type\":\"address\"}],\"name\":\"updateMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin_\",\"type\":\"address\"}],\"name\":\"updateAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620013c0380380620013c08339810160408190526200003491620004b3565b8851899089906200004d90600390602085019062000339565b5080516200006390600490602084019062000339565b5050506c010000000000000000000000008610620000c85760405162461bcd60e51b815260206004820152601360248201527f6d6178537570706c7920746f6f206c617267650000000000000000000000000060448201526064015b60405180910390fd5b85821115620001325760405162461bcd60e51b815260206004820152602f60248201527f696e697469616c20737570706c792073686f756c64206e6f742065786365656460448201526e20746865206d617820737570706c7960881b6064820152608401620000bf565b6005805460ff891660ff199091161790556007869055600880546001600160a01b038088166001600160a01b0319928316179092556006869055600980549284169290911691909117905562000189838362000254565b6008546040516001600160a01b0390911681527f94a0a7c0a7a455351029a521046d4438bc63e58a1ee8d984d624eb3161583b2c9060200160405180910390a16009546040516001600160a01b0390911681527fbfc8d7754fec5096becc28e0816011e8d0adcfe752ffa1c88924a7f88b00a41d9060200160405180910390a17fd65d781db82f0bd55f732ea1b96f4c6565750df815c18a70962cd8405274e2436006546040516200023d91815260200190565b60405180910390a150505050505050505062000607565b6001600160a01b038216620002ac5760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152606401620000bf565b8060026000828254620002c091906200058d565b90915550506001600160a01b03821660009081526020819052604081208054839290620002ef9084906200058d565b90915550506040518181526001600160a01b038316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b8280546200034790620005b4565b90600052602060002090601f0160209004810192826200036b5760008555620003b6565b82601f106200038657805160ff1916838001178555620003b6565b82800160010185558215620003b6579182015b82811115620003b657825182559160200191906001019062000399565b50620003c4929150620003c8565b5090565b5b80821115620003c45760008155600101620003c9565b80516001600160a01b0381168114620003f757600080fd5b919050565b600082601f8301126200040e57600080fd5b81516001600160401b03808211156200042b576200042b620005f1565b604051601f8301601f19908116603f01168101908282118183101715620004565762000456620005f1565b816040528381526020925086838588010111156200047357600080fd5b600091505b8382101562000497578582018301518183018401529082019062000478565b83821115620004a95760008385830101525b9695505050505050565b60008060008060008060008060006101208a8c031215620004d357600080fd5b89516001600160401b0380821115620004eb57600080fd5b620004f98d838e01620003fc565b9a5060208c01519150808211156200051057600080fd5b506200051f8c828d01620003fc565b98505060408a015160ff811681146200053757600080fd5b60608b015190975095506200054f60808b01620003df565b945060a08a015193506200056660c08b01620003df565b925060e08a015191506200057e6101008b01620003df565b90509295985092959850929598565b60008219821115620005af57634e487b7160e01b600052601160045260246000fd5b500190565b600181811c90821680620005c957607f821691505b60208210811415620005eb57634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052604160045260246000fd5b610da980620006176000396000f3fe608060405234801561001057600080fd5b50600436106101215760003560e01c80634eb03f6e116100ad578063ab8e786f11610071578063ab8e786f14610266578063d5abeb0114610279578063dd62ed3e14610282578063e2f273bd14610295578063f851a440146102a857600080fd5b80634eb03f6e146101fc57806370a082311461020f57806395d89b4114610238578063a457c2d714610240578063a9059cbb1461025357600080fd5b806323b872dd116100f457806323b872dd146101a457806324e770d0146101b757806329fb7ef6146101bf578063313ce567146101d457806339509351146101e957600080fd5b806306fdde03146101265780630754617214610144578063095ea7b31461016f57806318160ddd14610192575b600080fd5b61012e6102bb565b60405161013b9190610c67565b60405180910390f35b600854610157906001600160a01b031681565b6040516001600160a01b03909116815260200161013b565b61018261017d366004610c0b565b61034d565b604051901515815260200161013b565b6002545b60405190815260200161013b565b6101826101b2366004610bcf565b610367565b600654610196565b6101d26101cd366004610c35565b61038b565b005b60055460405160ff909116815260200161013b565b6101826101f7366004610c0b565b6103fa565b6101d261020a366004610b81565b61041c565b61019661021d366004610b81565b6001600160a01b031660009081526020819052604090205490565b61012e610494565b61018261024e366004610c0b565b6104a3565b610182610261366004610c0b565b61051e565b610182610274366004610c0b565b61052c565b61019660075481565b610196610290366004610b9c565b610658565b6101d26102a3366004610b81565b610683565b600954610157906001600160a01b031681565b6060600380546102ca90610d22565b80601f01602080910402602001604051908101604052809291908181526020018280546102f690610d22565b80156103435780601f1061031857610100808354040283529160200191610343565b820191906000526020600020905b81548152906001019060200180831161032657829003601f168201915b5050505050905090565b60003361035b8185856106fb565b60019150505b92915050565b60003361037585828561081f565b610380858585610899565b506001949350505050565b6009546001600160a01b031633146103be5760405162461bcd60e51b81526004016103b590610cbc565b60405180910390fd5b60068190556040518181527fd65d781db82f0bd55f732ea1b96f4c6565750df815c18a70962cd8405274e243906020015b60405180910390a150565b60003361035b81858561040d8383610658565b6104179190610cf3565b6106fb565b6009546001600160a01b031633146104465760405162461bcd60e51b81526004016103b590610cbc565b600880546001600160a01b0319166001600160a01b0383169081179091556040519081527f94a0a7c0a7a455351029a521046d4438bc63e58a1ee8d984d624eb3161583b2c906020016103ef565b6060600480546102ca90610d22565b600033816104b18286610658565b9050838110156105115760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b60648201526084016103b5565b61038082868684036106fb565b60003361035b818585610899565b6008546000906001600160a01b031633146105895760405162461bcd60e51b815260206004820152601e60248201527f4f6e6c79206d696e7465722063616e206d616b6520746869732063616c6c000060448201526064016103b5565b6008546001600160a01b031633146105a357506000610361565b6000306001600160a01b03166318160ddd6040518163ffffffff1660e01b815260040160206040518083038186803b1580156105de57600080fd5b505afa1580156105f2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106169190610c4e565b9050600754811061062b576000915050610361565b6007546106388285610a67565b111561064e5760075461064b9082610a7a565b92505b61035b8484610a86565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6009546001600160a01b031633146106ad5760405162461bcd60e51b81526004016103b590610cbc565b600980546001600160a01b0319166001600160a01b0383169081179091556040519081527fbfc8d7754fec5096becc28e0816011e8d0adcfe752ffa1c88924a7f88b00a41d906020016103ef565b6001600160a01b03831661075d5760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b60648201526084016103b5565b6001600160a01b0382166107be5760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b60648201526084016103b5565b6001600160a01b0383811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b600061082b8484610658565b9050600019811461089357818110156108865760405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e636500000060448201526064016103b5565b61089384848484036106fb565b50505050565b6001600160a01b0383166108fd5760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b60648201526084016103b5565b6001600160a01b03821661095f5760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b60648201526084016103b5565b6001600160a01b038316600090815260208190526040902054818110156109d75760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b60648201526084016103b5565b6001600160a01b03808516600090815260208190526040808220858503905591851681529081208054849290610a0e908490610cf3565b92505081905550826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610a5a91815260200190565b60405180910390a3610893565b6000610a738284610cf3565b9392505050565b6000610a738284610d0b565b6001600160a01b038216610adc5760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f20616464726573730060448201526064016103b5565b8060026000828254610aee9190610cf3565b90915550506001600160a01b03821660009081526020819052604081208054839290610b1b908490610cf3565b90915550506040518181526001600160a01b038316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b80356001600160a01b0381168114610b7c57600080fd5b919050565b600060208284031215610b9357600080fd5b610a7382610b65565b60008060408385031215610baf57600080fd5b610bb883610b65565b9150610bc660208401610b65565b90509250929050565b600080600060608486031215610be457600080fd5b610bed84610b65565b9250610bfb60208501610b65565b9150604084013590509250925092565b60008060408385031215610c1e57600080fd5b610c2783610b65565b946020939093013593505050565b600060208284031215610c4757600080fd5b5035919050565b600060208284031215610c6057600080fd5b5051919050565b600060208083528351808285015260005b81811015610c9457858101830151858201604001528201610c78565b81811115610ca6576000604083870101525b50601f01601f1916929092016040019392505050565b6020808252601d908201527f4f6e6c792061646d696e2063616e206d616b6520746869732063616c6c000000604082015260600190565b60008219821115610d0657610d06610d5d565b500190565b600082821015610d1d57610d1d610d5d565b500390565b600181811c90821680610d3657607f821691505b60208210811415610d5757634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601160045260246000fdfea2646970667358221220230433e5cb82157e9488ab13df024c650e78fa9eb771e599c8d37547351839bc64736f6c63430008070033",
}

// SubchainGovernanceTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use SubchainGovernanceTokenMetaData.ABI instead.
var SubchainGovernanceTokenABI = SubchainGovernanceTokenMetaData.ABI

// SubchainGovernanceTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SubchainGovernanceTokenMetaData.Bin instead.
var SubchainGovernanceTokenBin = SubchainGovernanceTokenMetaData.Bin

// DeploySubchainGovernanceToken deploys a new Ethereum contract, binding an instance of SubchainGovernanceToken to it.
func DeploySubchainGovernanceToken(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ string, decimals_ uint8, maxSupply_ *big.Int, minter_ common.Address, stakerRewardPerBlock_ *big.Int, initDistrWallet_ common.Address, initMintAmount_ *big.Int, admin_ common.Address) (common.Address, *types.Transaction, *SubchainGovernanceToken, error) {
	parsed, err := SubchainGovernanceTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SubchainGovernanceTokenBin), backend, name_, symbol_, decimals_, maxSupply_, minter_, stakerRewardPerBlock_, initDistrWallet_, initMintAmount_, admin_)
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

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_SubchainGovernanceToken *SubchainGovernanceTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SubchainGovernanceToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_SubchainGovernanceToken *SubchainGovernanceTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _SubchainGovernanceToken.Contract.BalanceOf(&_SubchainGovernanceToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_SubchainGovernanceToken *SubchainGovernanceTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _SubchainGovernanceToken.Contract.BalanceOf(&_SubchainGovernanceToken.CallOpts, account)
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

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_SubchainGovernanceToken *SubchainGovernanceTokenCaller) MaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SubchainGovernanceToken.contract.Call(opts, &out, "maxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_SubchainGovernanceToken *SubchainGovernanceTokenSession) MaxSupply() (*big.Int, error) {
	return _SubchainGovernanceToken.Contract.MaxSupply(&_SubchainGovernanceToken.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_SubchainGovernanceToken *SubchainGovernanceTokenCallerSession) MaxSupply() (*big.Int, error) {
	return _SubchainGovernanceToken.Contract.MaxSupply(&_SubchainGovernanceToken.CallOpts)
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

// StakerRewardPerBlock is a free data retrieval call binding the contract method 0x24e770d0.
//
// Solidity: function stakerRewardPerBlock() view returns(uint256)
func (_SubchainGovernanceToken *SubchainGovernanceTokenCaller) StakerRewardPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SubchainGovernanceToken.contract.Call(opts, &out, "stakerRewardPerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerRewardPerBlock is a free data retrieval call binding the contract method 0x24e770d0.
//
// Solidity: function stakerRewardPerBlock() view returns(uint256)
func (_SubchainGovernanceToken *SubchainGovernanceTokenSession) StakerRewardPerBlock() (*big.Int, error) {
	return _SubchainGovernanceToken.Contract.StakerRewardPerBlock(&_SubchainGovernanceToken.CallOpts)
}

// StakerRewardPerBlock is a free data retrieval call binding the contract method 0x24e770d0.
//
// Solidity: function stakerRewardPerBlock() view returns(uint256)
func (_SubchainGovernanceToken *SubchainGovernanceTokenCallerSession) StakerRewardPerBlock() (*big.Int, error) {
	return _SubchainGovernanceToken.Contract.StakerRewardPerBlock(&_SubchainGovernanceToken.CallOpts)
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

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SubchainGovernanceToken *SubchainGovernanceTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SubchainGovernanceToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SubchainGovernanceToken *SubchainGovernanceTokenSession) TotalSupply() (*big.Int, error) {
	return _SubchainGovernanceToken.Contract.TotalSupply(&_SubchainGovernanceToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SubchainGovernanceToken *SubchainGovernanceTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _SubchainGovernanceToken.Contract.TotalSupply(&_SubchainGovernanceToken.CallOpts)
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

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_SubchainGovernanceToken *SubchainGovernanceTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _SubchainGovernanceToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_SubchainGovernanceToken *SubchainGovernanceTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _SubchainGovernanceToken.Contract.DecreaseAllowance(&_SubchainGovernanceToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_SubchainGovernanceToken *SubchainGovernanceTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _SubchainGovernanceToken.Contract.DecreaseAllowance(&_SubchainGovernanceToken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_SubchainGovernanceToken *SubchainGovernanceTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _SubchainGovernanceToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_SubchainGovernanceToken *SubchainGovernanceTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _SubchainGovernanceToken.Contract.IncreaseAllowance(&_SubchainGovernanceToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_SubchainGovernanceToken *SubchainGovernanceTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _SubchainGovernanceToken.Contract.IncreaseAllowance(&_SubchainGovernanceToken.TransactOpts, spender, addedValue)
}

// MintStakerReward is a paid mutator transaction binding the contract method 0xab8e786f.
//
// Solidity: function mintStakerReward(address account, uint256 amount) returns(bool)
func (_SubchainGovernanceToken *SubchainGovernanceTokenTransactor) MintStakerReward(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SubchainGovernanceToken.contract.Transact(opts, "mintStakerReward", account, amount)
}

// MintStakerReward is a paid mutator transaction binding the contract method 0xab8e786f.
//
// Solidity: function mintStakerReward(address account, uint256 amount) returns(bool)
func (_SubchainGovernanceToken *SubchainGovernanceTokenSession) MintStakerReward(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SubchainGovernanceToken.Contract.MintStakerReward(&_SubchainGovernanceToken.TransactOpts, account, amount)
}

// MintStakerReward is a paid mutator transaction binding the contract method 0xab8e786f.
//
// Solidity: function mintStakerReward(address account, uint256 amount) returns(bool)
func (_SubchainGovernanceToken *SubchainGovernanceTokenTransactorSession) MintStakerReward(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SubchainGovernanceToken.Contract.MintStakerReward(&_SubchainGovernanceToken.TransactOpts, account, amount)
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

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_SubchainGovernanceToken *SubchainGovernanceTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SubchainGovernanceToken.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_SubchainGovernanceToken *SubchainGovernanceTokenSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SubchainGovernanceToken.Contract.TransferFrom(&_SubchainGovernanceToken.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_SubchainGovernanceToken *SubchainGovernanceTokenTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SubchainGovernanceToken.Contract.TransferFrom(&_SubchainGovernanceToken.TransactOpts, from, to, amount)
}

// UpdateAdmin is a paid mutator transaction binding the contract method 0xe2f273bd.
//
// Solidity: function updateAdmin(address admin_) returns()
func (_SubchainGovernanceToken *SubchainGovernanceTokenTransactor) UpdateAdmin(opts *bind.TransactOpts, admin_ common.Address) (*types.Transaction, error) {
	return _SubchainGovernanceToken.contract.Transact(opts, "updateAdmin", admin_)
}

// UpdateAdmin is a paid mutator transaction binding the contract method 0xe2f273bd.
//
// Solidity: function updateAdmin(address admin_) returns()
func (_SubchainGovernanceToken *SubchainGovernanceTokenSession) UpdateAdmin(admin_ common.Address) (*types.Transaction, error) {
	return _SubchainGovernanceToken.Contract.UpdateAdmin(&_SubchainGovernanceToken.TransactOpts, admin_)
}

// UpdateAdmin is a paid mutator transaction binding the contract method 0xe2f273bd.
//
// Solidity: function updateAdmin(address admin_) returns()
func (_SubchainGovernanceToken *SubchainGovernanceTokenTransactorSession) UpdateAdmin(admin_ common.Address) (*types.Transaction, error) {
	return _SubchainGovernanceToken.Contract.UpdateAdmin(&_SubchainGovernanceToken.TransactOpts, admin_)
}

// UpdateMinter is a paid mutator transaction binding the contract method 0x4eb03f6e.
//
// Solidity: function updateMinter(address minter_) returns()
func (_SubchainGovernanceToken *SubchainGovernanceTokenTransactor) UpdateMinter(opts *bind.TransactOpts, minter_ common.Address) (*types.Transaction, error) {
	return _SubchainGovernanceToken.contract.Transact(opts, "updateMinter", minter_)
}

// UpdateMinter is a paid mutator transaction binding the contract method 0x4eb03f6e.
//
// Solidity: function updateMinter(address minter_) returns()
func (_SubchainGovernanceToken *SubchainGovernanceTokenSession) UpdateMinter(minter_ common.Address) (*types.Transaction, error) {
	return _SubchainGovernanceToken.Contract.UpdateMinter(&_SubchainGovernanceToken.TransactOpts, minter_)
}

// UpdateMinter is a paid mutator transaction binding the contract method 0x4eb03f6e.
//
// Solidity: function updateMinter(address minter_) returns()
func (_SubchainGovernanceToken *SubchainGovernanceTokenTransactorSession) UpdateMinter(minter_ common.Address) (*types.Transaction, error) {
	return _SubchainGovernanceToken.Contract.UpdateMinter(&_SubchainGovernanceToken.TransactOpts, minter_)
}

// UpdateStakerRewardPerBlock is a paid mutator transaction binding the contract method 0x29fb7ef6.
//
// Solidity: function updateStakerRewardPerBlock(uint256 stakerRewardPerBlock_) returns()
func (_SubchainGovernanceToken *SubchainGovernanceTokenTransactor) UpdateStakerRewardPerBlock(opts *bind.TransactOpts, stakerRewardPerBlock_ *big.Int) (*types.Transaction, error) {
	return _SubchainGovernanceToken.contract.Transact(opts, "updateStakerRewardPerBlock", stakerRewardPerBlock_)
}

// UpdateStakerRewardPerBlock is a paid mutator transaction binding the contract method 0x29fb7ef6.
//
// Solidity: function updateStakerRewardPerBlock(uint256 stakerRewardPerBlock_) returns()
func (_SubchainGovernanceToken *SubchainGovernanceTokenSession) UpdateStakerRewardPerBlock(stakerRewardPerBlock_ *big.Int) (*types.Transaction, error) {
	return _SubchainGovernanceToken.Contract.UpdateStakerRewardPerBlock(&_SubchainGovernanceToken.TransactOpts, stakerRewardPerBlock_)
}

// UpdateStakerRewardPerBlock is a paid mutator transaction binding the contract method 0x29fb7ef6.
//
// Solidity: function updateStakerRewardPerBlock(uint256 stakerRewardPerBlock_) returns()
func (_SubchainGovernanceToken *SubchainGovernanceTokenTransactorSession) UpdateStakerRewardPerBlock(stakerRewardPerBlock_ *big.Int) (*types.Transaction, error) {
	return _SubchainGovernanceToken.Contract.UpdateStakerRewardPerBlock(&_SubchainGovernanceToken.TransactOpts, stakerRewardPerBlock_)
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

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SubchainGovernanceToken *SubchainGovernanceTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*SubchainGovernanceTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SubchainGovernanceToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &SubchainGovernanceTokenApprovalIterator{contract: _SubchainGovernanceToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SubchainGovernanceToken *SubchainGovernanceTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SubchainGovernanceTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SubchainGovernanceToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
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
				if err := _SubchainGovernanceToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_SubchainGovernanceToken *SubchainGovernanceTokenFilterer) ParseApproval(log types.Log) (*SubchainGovernanceTokenApproval, error) {
	event := new(SubchainGovernanceTokenApproval)
	if err := _SubchainGovernanceToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SubchainGovernanceToken *SubchainGovernanceTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SubchainGovernanceTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SubchainGovernanceToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SubchainGovernanceTokenTransferIterator{contract: _SubchainGovernanceToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SubchainGovernanceToken *SubchainGovernanceTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SubchainGovernanceTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SubchainGovernanceToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
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
				if err := _SubchainGovernanceToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_SubchainGovernanceToken *SubchainGovernanceTokenFilterer) ParseTransfer(log types.Log) (*SubchainGovernanceTokenTransfer, error) {
	event := new(SubchainGovernanceTokenTransfer)
	if err := _SubchainGovernanceToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubchainGovernanceTokenUpdateAdminIterator is returned from FilterUpdateAdmin and is used to iterate over the raw logs and unpacked data for UpdateAdmin events raised by the SubchainGovernanceToken contract.
type SubchainGovernanceTokenUpdateAdminIterator struct {
	Event *SubchainGovernanceTokenUpdateAdmin // Event containing the contract specifics and raw log

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
func (it *SubchainGovernanceTokenUpdateAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubchainGovernanceTokenUpdateAdmin)
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
		it.Event = new(SubchainGovernanceTokenUpdateAdmin)
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
func (it *SubchainGovernanceTokenUpdateAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubchainGovernanceTokenUpdateAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubchainGovernanceTokenUpdateAdmin represents a UpdateAdmin event raised by the SubchainGovernanceToken contract.
type SubchainGovernanceTokenUpdateAdmin struct {
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdateAdmin is a free log retrieval operation binding the contract event 0xbfc8d7754fec5096becc28e0816011e8d0adcfe752ffa1c88924a7f88b00a41d.
//
// Solidity: event UpdateAdmin(address newAdmin)
func (_SubchainGovernanceToken *SubchainGovernanceTokenFilterer) FilterUpdateAdmin(opts *bind.FilterOpts) (*SubchainGovernanceTokenUpdateAdminIterator, error) {

	logs, sub, err := _SubchainGovernanceToken.contract.FilterLogs(opts, "UpdateAdmin")
	if err != nil {
		return nil, err
	}
	return &SubchainGovernanceTokenUpdateAdminIterator{contract: _SubchainGovernanceToken.contract, event: "UpdateAdmin", logs: logs, sub: sub}, nil
}

// WatchUpdateAdmin is a free log subscription operation binding the contract event 0xbfc8d7754fec5096becc28e0816011e8d0adcfe752ffa1c88924a7f88b00a41d.
//
// Solidity: event UpdateAdmin(address newAdmin)
func (_SubchainGovernanceToken *SubchainGovernanceTokenFilterer) WatchUpdateAdmin(opts *bind.WatchOpts, sink chan<- *SubchainGovernanceTokenUpdateAdmin) (event.Subscription, error) {

	logs, sub, err := _SubchainGovernanceToken.contract.WatchLogs(opts, "UpdateAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubchainGovernanceTokenUpdateAdmin)
				if err := _SubchainGovernanceToken.contract.UnpackLog(event, "UpdateAdmin", log); err != nil {
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

// ParseUpdateAdmin is a log parse operation binding the contract event 0xbfc8d7754fec5096becc28e0816011e8d0adcfe752ffa1c88924a7f88b00a41d.
//
// Solidity: event UpdateAdmin(address newAdmin)
func (_SubchainGovernanceToken *SubchainGovernanceTokenFilterer) ParseUpdateAdmin(log types.Log) (*SubchainGovernanceTokenUpdateAdmin, error) {
	event := new(SubchainGovernanceTokenUpdateAdmin)
	if err := _SubchainGovernanceToken.contract.UnpackLog(event, "UpdateAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubchainGovernanceTokenUpdateMinterIterator is returned from FilterUpdateMinter and is used to iterate over the raw logs and unpacked data for UpdateMinter events raised by the SubchainGovernanceToken contract.
type SubchainGovernanceTokenUpdateMinterIterator struct {
	Event *SubchainGovernanceTokenUpdateMinter // Event containing the contract specifics and raw log

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
func (it *SubchainGovernanceTokenUpdateMinterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubchainGovernanceTokenUpdateMinter)
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
		it.Event = new(SubchainGovernanceTokenUpdateMinter)
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
func (it *SubchainGovernanceTokenUpdateMinterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubchainGovernanceTokenUpdateMinterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubchainGovernanceTokenUpdateMinter represents a UpdateMinter event raised by the SubchainGovernanceToken contract.
type SubchainGovernanceTokenUpdateMinter struct {
	NewMinter common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdateMinter is a free log retrieval operation binding the contract event 0x94a0a7c0a7a455351029a521046d4438bc63e58a1ee8d984d624eb3161583b2c.
//
// Solidity: event UpdateMinter(address newMinter)
func (_SubchainGovernanceToken *SubchainGovernanceTokenFilterer) FilterUpdateMinter(opts *bind.FilterOpts) (*SubchainGovernanceTokenUpdateMinterIterator, error) {

	logs, sub, err := _SubchainGovernanceToken.contract.FilterLogs(opts, "UpdateMinter")
	if err != nil {
		return nil, err
	}
	return &SubchainGovernanceTokenUpdateMinterIterator{contract: _SubchainGovernanceToken.contract, event: "UpdateMinter", logs: logs, sub: sub}, nil
}

// WatchUpdateMinter is a free log subscription operation binding the contract event 0x94a0a7c0a7a455351029a521046d4438bc63e58a1ee8d984d624eb3161583b2c.
//
// Solidity: event UpdateMinter(address newMinter)
func (_SubchainGovernanceToken *SubchainGovernanceTokenFilterer) WatchUpdateMinter(opts *bind.WatchOpts, sink chan<- *SubchainGovernanceTokenUpdateMinter) (event.Subscription, error) {

	logs, sub, err := _SubchainGovernanceToken.contract.WatchLogs(opts, "UpdateMinter")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubchainGovernanceTokenUpdateMinter)
				if err := _SubchainGovernanceToken.contract.UnpackLog(event, "UpdateMinter", log); err != nil {
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

// ParseUpdateMinter is a log parse operation binding the contract event 0x94a0a7c0a7a455351029a521046d4438bc63e58a1ee8d984d624eb3161583b2c.
//
// Solidity: event UpdateMinter(address newMinter)
func (_SubchainGovernanceToken *SubchainGovernanceTokenFilterer) ParseUpdateMinter(log types.Log) (*SubchainGovernanceTokenUpdateMinter, error) {
	event := new(SubchainGovernanceTokenUpdateMinter)
	if err := _SubchainGovernanceToken.contract.UnpackLog(event, "UpdateMinter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubchainGovernanceTokenUpdateStakerRewardPerBlockIterator is returned from FilterUpdateStakerRewardPerBlock and is used to iterate over the raw logs and unpacked data for UpdateStakerRewardPerBlock events raised by the SubchainGovernanceToken contract.
type SubchainGovernanceTokenUpdateStakerRewardPerBlockIterator struct {
	Event *SubchainGovernanceTokenUpdateStakerRewardPerBlock // Event containing the contract specifics and raw log

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
func (it *SubchainGovernanceTokenUpdateStakerRewardPerBlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubchainGovernanceTokenUpdateStakerRewardPerBlock)
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
		it.Event = new(SubchainGovernanceTokenUpdateStakerRewardPerBlock)
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
func (it *SubchainGovernanceTokenUpdateStakerRewardPerBlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubchainGovernanceTokenUpdateStakerRewardPerBlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubchainGovernanceTokenUpdateStakerRewardPerBlock represents a UpdateStakerRewardPerBlock event raised by the SubchainGovernanceToken contract.
type SubchainGovernanceTokenUpdateStakerRewardPerBlock struct {
	NewStakerRewardPerBlock *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterUpdateStakerRewardPerBlock is a free log retrieval operation binding the contract event 0xd65d781db82f0bd55f732ea1b96f4c6565750df815c18a70962cd8405274e243.
//
// Solidity: event UpdateStakerRewardPerBlock(uint256 newStakerRewardPerBlock)
func (_SubchainGovernanceToken *SubchainGovernanceTokenFilterer) FilterUpdateStakerRewardPerBlock(opts *bind.FilterOpts) (*SubchainGovernanceTokenUpdateStakerRewardPerBlockIterator, error) {

	logs, sub, err := _SubchainGovernanceToken.contract.FilterLogs(opts, "UpdateStakerRewardPerBlock")
	if err != nil {
		return nil, err
	}
	return &SubchainGovernanceTokenUpdateStakerRewardPerBlockIterator{contract: _SubchainGovernanceToken.contract, event: "UpdateStakerRewardPerBlock", logs: logs, sub: sub}, nil
}

// WatchUpdateStakerRewardPerBlock is a free log subscription operation binding the contract event 0xd65d781db82f0bd55f732ea1b96f4c6565750df815c18a70962cd8405274e243.
//
// Solidity: event UpdateStakerRewardPerBlock(uint256 newStakerRewardPerBlock)
func (_SubchainGovernanceToken *SubchainGovernanceTokenFilterer) WatchUpdateStakerRewardPerBlock(opts *bind.WatchOpts, sink chan<- *SubchainGovernanceTokenUpdateStakerRewardPerBlock) (event.Subscription, error) {

	logs, sub, err := _SubchainGovernanceToken.contract.WatchLogs(opts, "UpdateStakerRewardPerBlock")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubchainGovernanceTokenUpdateStakerRewardPerBlock)
				if err := _SubchainGovernanceToken.contract.UnpackLog(event, "UpdateStakerRewardPerBlock", log); err != nil {
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

// ParseUpdateStakerRewardPerBlock is a log parse operation binding the contract event 0xd65d781db82f0bd55f732ea1b96f4c6565750df815c18a70962cd8405274e243.
//
// Solidity: event UpdateStakerRewardPerBlock(uint256 newStakerRewardPerBlock)
func (_SubchainGovernanceToken *SubchainGovernanceTokenFilterer) ParseUpdateStakerRewardPerBlock(log types.Log) (*SubchainGovernanceTokenUpdateStakerRewardPerBlock, error) {
	event := new(SubchainGovernanceTokenUpdateStakerRewardPerBlock)
	if err := _SubchainGovernanceToken.contract.UnpackLog(event, "UpdateStakerRewardPerBlock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
