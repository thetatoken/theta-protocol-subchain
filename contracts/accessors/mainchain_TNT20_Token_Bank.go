// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// MainchainTNT20TokenBankMetaData contains all meta data concerning the MainchainTNT20TokenBank contract.
var MainchainTNT20TokenBankMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"registerContractAddress_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"targetChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"mainchainTokenSender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"subchainTokenReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"TNT20Contract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"decimal\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"name\":\"TNT20TokenLocked\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allDenoms\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allVouchers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"denomToVoucherLookup\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voucherAddress\",\"type\":\"address\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voucherContractAddr\",\"type\":\"address\"}],\"name\":\"getDenom\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"name\":\"getVoucher\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"voucherAddressToDenomLookup\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"targetChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"TNT20Contract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"subchainTokenReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sendToChain\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\",\"payable\":true}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001aa138038062001aa1833981810160405281019062000037919062000096565b80600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506200011b565b600081519050620000908162000101565b92915050565b600060208284031215620000af57620000ae620000fc565b5b6000620000bf848285016200007f565b91505092915050565b6000620000d582620000dc565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b6200010c81620000c8565b81146200011857600080fd5b50565b611976806200012b6000396000f3fe6080604052600436106100865760003560e01c806360569b5e1161005957806360569b5e14610180578063a2cc6981146101be578063ebda9962146101fb578063f6a3d24e14610238578063faf9aa6e1461027557610086565b80631527b14d1461008b578063261a323e146100c957806327ca4df114610106578063588b140814610143575b600080fd5b34801561009757600080fd5b506100b260048036038101906100ad9190611013565b610291565b6040516100c0929190611315565b60405180910390f35b3480156100d557600080fd5b506100f060048036038101906100eb9190611013565b6102f8565b6040516100fd919061133e565b60405180910390f35b34801561011257600080fd5b5061012d600480360381019061012891906110a5565b61033d565b60405161013a91906112c3565b60405180910390f35b34801561014f57600080fd5b5061016a600480360381019061016591906110a5565b61037c565b6040516101779190611359565b60405180910390f35b34801561018c57600080fd5b506101a760048036038101906101a29190610fb9565b610428565b6040516101b592919061137b565b60405180910390f35b3480156101ca57600080fd5b506101e560048036038101906101e09190611013565b6104e1565b6040516101f291906112c3565b60405180910390f35b34801561020757600080fd5b50610222600480360381019061021d9190610fb9565b6105b1565b60405161022f9190611359565b60405180910390f35b34801561024457600080fd5b5061025f600480360381019061025a9190610fb9565b6106de565b60405161026c919061133e565b60405180910390f35b61028f600480360381019061028a91906110d2565b610737565b005b6002818051602081018201805184825260208301602085012081835280955050505050506000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060000160149054906101000a900460ff16905082565b60008061030483610b92565b90506002816040516103169190611261565b908152602001604051809103902060000160149054906101000a900460ff16915050919050565b6004818154811061034d57600080fd5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6005818154811061038c57600080fd5b9060005260206000200160009150905080546103a79061169d565b80601f01602080910402602001604051908101604052809291908181526020018280546103d39061169d565b80156104205780601f106103f557610100808354040283529160200191610420565b820191906000526020600020905b81548152906001019060200180831161040357829003601f168201915b505050505081565b600360205280600052604060002060009150905080600001805461044b9061169d565b80601f01602080910402602001604051908101604052809291908181526020018280546104779061169d565b80156104c45780601f10610499576101008083540402835291602001916104c4565b820191906000526020600020905b8154815290600101906020018083116104a757829003601f168201915b5050505050908060010160009054906101000a900460ff16905082565b6000806104ed83610b92565b905060006002826040516105019190611261565b90815260200160405180910390206040518060400160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160149054906101000a900460ff16151515158152505090508060200151156105a5578060000151925050506105ac565b6000925050505b919050565b60606000600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060405180604001604052908160008201805461060f9061169d565b80601f016020809104026020016040519081016040528092919081815260200182805461063b9061169d565b80156106885780601f1061065d57610100808354040283529160200191610688565b820191906000526020600020905b81548152906001019060200180831161066b57829003601f168201915b505050505081526020016001820160009054906101000a900460ff16151515158152505090508060200151156106c55780600001519150506106d9565b604051806020016040528060008152509150505b919050565b6000600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900460ff169050919050565b6000600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637d1daa96866040518263ffffffff1660e01b815260040161079491906113cb565b60206040518083038186803b1580156107ac57600080fd5b505afa1580156107c0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107e49190610fe6565b905080610826576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161081d906113ab565b60405180910390fd5b6000829050600033905060008690508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8330866040518463ffffffff1660e01b8152600401610872939291906112de565b602060405180830381600087803b15801561088c57600080fd5b505af11580156108a0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108c49190610fe6565b506108ce88610ba4565b60008060008a815260200190815260200160002054905060006109ba6108f346610bd1565b6040518060400160405280600181526020017f2f000000000000000000000000000000000000000000000000000000000000008152506040518060400160405280600281526020017f32300000000000000000000000000000000000000000000000000000000000008152506040518060400160405280600181526020017f2f000000000000000000000000000000000000000000000000000000000000008152508d6040516020016109a69190611246565b604051602081830303815290604052610d32565b90507fe5d2c5efdbbd7c86071c36e6c82f176bac6830073d3498f8641318a0763875f88a858a888d8873ffffffffffffffffffffffffffffffffffffffff166306fdde036040518163ffffffff1660e01b815260040160006040518083038186803b158015610a2857600080fd5b505afa158015610a3c573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190610a65919061105c565b8973ffffffffffffffffffffffffffffffffffffffff166395d89b416040518163ffffffff1660e01b815260040160006040518083038186803b158015610aab57600080fd5b505afa158015610abf573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190610ae8919061105c565b8a73ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b815260040160206040518083038186803b158015610b2e57600080fd5b505afa158015610b42573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b669190611139565b8a8a604051610b7e9a999897969594939291906113e6565b60405180910390a150505050505050505050565b6060610b9d82610d67565b9050919050565b60016000808381526020019081526020016000206000828254610bc79190611514565b9250508190555050565b60606000821415610c19576040518060400160405280600181526020017f30000000000000000000000000000000000000000000000000000000000000008152509050610d2d565b600082905060005b60008214610c4b578080610c3490611700565b915050600a82610c4491906115a1565b9150610c21565b60008167ffffffffffffffff811115610c6757610c6661185a565b5b6040519080825280601f01601f191660200182016040528015610c995781602001600182028036833780820191505090505b5090505b60008514610d2657600182610cb291906115d2565b9150600a85610cc1919061176d565b6030610ccd9190611514565b60f81b818381518110610ce357610ce261182b565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a85610d1f91906115a1565b9450610c9d565b8093505050505b919050565b60608585858585604051602001610d4d959493929190611278565b604051602081830303815290604052905095945050505050565b6060600082905060005b8151811015610df557610da0828281518110610d9057610d8f61182b565b5b602001015160f81c60f81b610dff565b828281518110610db357610db261182b565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508080610ded90611700565b915050610d71565b5080915050919050565b6000604160f81b827effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191610158015610e5d5750605a60f81b827effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191611155b15610e7c5760208260f81c610e72919061156a565b60f81b9050610e80565b8190505b919050565b6000610e98610e93846114bc565b611497565b905082815260208101848484011115610eb457610eb361188e565b5b610ebf84828561165b565b509392505050565b6000610eda610ed5846114bc565b611497565b905082815260208101848484011115610ef657610ef561188e565b5b610f0184828561166a565b509392505050565b600081359050610f18816118e4565b92915050565b600081519050610f2d816118fb565b92915050565b600082601f830112610f4857610f47611889565b5b8135610f58848260208601610e85565b91505092915050565b600082601f830112610f7657610f75611889565b5b8151610f86848260208601610ec7565b91505092915050565b600081359050610f9e81611912565b92915050565b600081519050610fb381611929565b92915050565b600060208284031215610fcf57610fce611898565b5b6000610fdd84828501610f09565b91505092915050565b600060208284031215610ffc57610ffb611898565b5b600061100a84828501610f1e565b91505092915050565b60006020828403121561102957611028611898565b5b600082013567ffffffffffffffff81111561104757611046611893565b5b61105384828501610f33565b91505092915050565b60006020828403121561107257611071611898565b5b600082015167ffffffffffffffff8111156110905761108f611893565b5b61109c84828501610f61565b91505092915050565b6000602082840312156110bb576110ba611898565b5b60006110c984828501610f8f565b91505092915050565b600080600080608085870312156110ec576110eb611898565b5b60006110fa87828801610f8f565b945050602061110b87828801610f09565b935050604061111c87828801610f09565b925050606061112d87828801610f8f565b91505092959194509250565b60006020828403121561114f5761114e611898565b5b600061115d84828501610fa4565b91505092915050565b61116f81611606565b82525050565b61118661118182611606565b611749565b82525050565b61119581611618565b82525050565b60006111a6826114ed565b6111b081856114f8565b93506111c081856020860161166a565b6111c98161189d565b840191505092915050565b60006111df826114ed565b6111e98185611509565b93506111f981856020860161166a565b80840191505092915050565b60006112126013836114f8565b915061121d826118bb565b602082019050919050565b61123181611644565b82525050565b6112408161164e565b82525050565b60006112528284611175565b60148201915081905092915050565b600061126d82846111d4565b915081905092915050565b600061128482886111d4565b915061129082876111d4565b915061129c82866111d4565b91506112a882856111d4565b91506112b482846111d4565b91508190509695505050505050565b60006020820190506112d86000830184611166565b92915050565b60006060820190506112f36000830186611166565b6113006020830185611166565b61130d6040830184611228565b949350505050565b600060408201905061132a6000830185611166565b611337602083018461118c565b9392505050565b6000602082019050611353600083018461118c565b92915050565b60006020820190508181036000830152611373818461119b565b905092915050565b60006040820190508181036000830152611395818561119b565b90506113a4602083018461118c565b9392505050565b600060208201905081810360008301526113c481611205565b9050919050565b60006020820190506113e06000830184611228565b92915050565b6000610140820190506113fc600083018d611228565b611409602083018c611166565b611416604083018b611166565b611423606083018a611228565b6114306080830189611166565b81810360a0830152611442818861119b565b905081810360c0830152611456818761119b565b905061146560e0830186611237565b611473610100830185611228565b818103610120830152611486818461119b565b90509b9a5050505050505050505050565b60006114a16114b2565b90506114ad82826116cf565b919050565b6000604051905090565b600067ffffffffffffffff8211156114d7576114d661185a565b5b6114e08261189d565b9050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600061151f82611644565b915061152a83611644565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561155f5761155e61179e565b5b828201905092915050565b60006115758261164e565b91506115808361164e565b92508260ff038211156115965761159561179e565b5b828201905092915050565b60006115ac82611644565b91506115b783611644565b9250826115c7576115c66117cd565b5b828204905092915050565b60006115dd82611644565b91506115e883611644565b9250828210156115fb576115fa61179e565b5b828203905092915050565b600061161182611624565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b8381101561168857808201518184015260208101905061166d565b83811115611697576000848401525b50505050565b600060028204905060018216806116b557607f821691505b602082108114156116c9576116c86117fc565b5b50919050565b6116d88261189d565b810181811067ffffffffffffffff821117156116f7576116f661185a565b5b80604052505050565b600061170b82611644565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82141561173e5761173d61179e565b5b600182019050919050565b60006117548261175b565b9050919050565b6000611766826118ae565b9050919050565b600061177882611644565b915061178383611644565b925082611793576117926117cd565b5b828206905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b60008160601b9050919050565b7f496e76616c696420537562636861696e49442100000000000000000000000000600082015250565b6118ed81611606565b81146118f857600080fd5b50565b61190481611618565b811461190f57600080fd5b50565b61191b81611644565b811461192657600080fd5b50565b6119328161164e565b811461193d57600080fd5b5056fea2646970667358221220f667cd63af7026fb075a394f6b8cbc2f8ddbd6b4982cbbd86f3aabbdc178f8fa64736f6c63430008070033",
}

// MainchainTNT20TokenBankABI is the input ABI used to generate the binding from.
// Deprecated: Use MainchainTNT20TokenBankMetaData.ABI instead.
var MainchainTNT20TokenBankABI = MainchainTNT20TokenBankMetaData.ABI

// MainchainTNT20TokenBankBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MainchainTNT20TokenBankMetaData.Bin instead.
var MainchainTNT20TokenBankBin = MainchainTNT20TokenBankMetaData.Bin

// DeployMainchainTNT20TokenBank deploys a new Ethereum contract, binding an instance of MainchainTNT20TokenBank to it.
func DeployMainchainTNT20TokenBank(auth *bind.TransactOpts, backend bind.ContractBackend, registerContractAddress_ common.Address) (common.Address, *types.Transaction, *MainchainTNT20TokenBank, error) {
	parsed, err := MainchainTNT20TokenBankMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MainchainTNT20TokenBankBin), backend, registerContractAddress_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MainchainTNT20TokenBank{MainchainTNT20TokenBankCaller: MainchainTNT20TokenBankCaller{contract: contract}, MainchainTNT20TokenBankTransactor: MainchainTNT20TokenBankTransactor{contract: contract}, MainchainTNT20TokenBankFilterer: MainchainTNT20TokenBankFilterer{contract: contract}}, nil
}

// MainchainTNT20TokenBank is an auto generated Go binding around an Ethereum contract.
type MainchainTNT20TokenBank struct {
	MainchainTNT20TokenBankCaller     // Read-only binding to the contract
	MainchainTNT20TokenBankTransactor // Write-only binding to the contract
	MainchainTNT20TokenBankFilterer   // Log filterer for contract events
}

// MainchainTNT20TokenBankCaller is an auto generated read-only Go binding around an Ethereum contract.
type MainchainTNT20TokenBankCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainchainTNT20TokenBankTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MainchainTNT20TokenBankTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainchainTNT20TokenBankFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MainchainTNT20TokenBankFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainchainTNT20TokenBankSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MainchainTNT20TokenBankSession struct {
	Contract     *MainchainTNT20TokenBank // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// MainchainTNT20TokenBankCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MainchainTNT20TokenBankCallerSession struct {
	Contract *MainchainTNT20TokenBankCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// MainchainTNT20TokenBankTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MainchainTNT20TokenBankTransactorSession struct {
	Contract     *MainchainTNT20TokenBankTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// MainchainTNT20TokenBankRaw is an auto generated low-level Go binding around an Ethereum contract.
type MainchainTNT20TokenBankRaw struct {
	Contract *MainchainTNT20TokenBank // Generic contract binding to access the raw methods on
}

// MainchainTNT20TokenBankCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MainchainTNT20TokenBankCallerRaw struct {
	Contract *MainchainTNT20TokenBankCaller // Generic read-only contract binding to access the raw methods on
}

// MainchainTNT20TokenBankTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MainchainTNT20TokenBankTransactorRaw struct {
	Contract *MainchainTNT20TokenBankTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMainchainTNT20TokenBank creates a new instance of MainchainTNT20TokenBank, bound to a specific deployed contract.
func NewMainchainTNT20TokenBank(address common.Address, backend bind.ContractBackend) (*MainchainTNT20TokenBank, error) {
	contract, err := bindMainchainTNT20TokenBank(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MainchainTNT20TokenBank{MainchainTNT20TokenBankCaller: MainchainTNT20TokenBankCaller{contract: contract}, MainchainTNT20TokenBankTransactor: MainchainTNT20TokenBankTransactor{contract: contract}, MainchainTNT20TokenBankFilterer: MainchainTNT20TokenBankFilterer{contract: contract}}, nil
}

// NewMainchainTNT20TokenBankCaller creates a new read-only instance of MainchainTNT20TokenBank, bound to a specific deployed contract.
func NewMainchainTNT20TokenBankCaller(address common.Address, caller bind.ContractCaller) (*MainchainTNT20TokenBankCaller, error) {
	contract, err := bindMainchainTNT20TokenBank(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MainchainTNT20TokenBankCaller{contract: contract}, nil
}

// NewMainchainTNT20TokenBankTransactor creates a new write-only instance of MainchainTNT20TokenBank, bound to a specific deployed contract.
func NewMainchainTNT20TokenBankTransactor(address common.Address, transactor bind.ContractTransactor) (*MainchainTNT20TokenBankTransactor, error) {
	contract, err := bindMainchainTNT20TokenBank(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MainchainTNT20TokenBankTransactor{contract: contract}, nil
}

// NewMainchainTNT20TokenBankFilterer creates a new log filterer instance of MainchainTNT20TokenBank, bound to a specific deployed contract.
func NewMainchainTNT20TokenBankFilterer(address common.Address, filterer bind.ContractFilterer) (*MainchainTNT20TokenBankFilterer, error) {
	contract, err := bindMainchainTNT20TokenBank(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MainchainTNT20TokenBankFilterer{contract: contract}, nil
}

// bindMainchainTNT20TokenBank binds a generic wrapper to an already deployed contract.
func bindMainchainTNT20TokenBank(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MainchainTNT20TokenBankABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MainchainTNT20TokenBank *MainchainTNT20TokenBankRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MainchainTNT20TokenBank.Contract.MainchainTNT20TokenBankCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MainchainTNT20TokenBank *MainchainTNT20TokenBankRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MainchainTNT20TokenBank.Contract.MainchainTNT20TokenBankTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MainchainTNT20TokenBank *MainchainTNT20TokenBankRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MainchainTNT20TokenBank.Contract.MainchainTNT20TokenBankTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MainchainTNT20TokenBank *MainchainTNT20TokenBankCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MainchainTNT20TokenBank.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MainchainTNT20TokenBank *MainchainTNT20TokenBankTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MainchainTNT20TokenBank.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MainchainTNT20TokenBank *MainchainTNT20TokenBankTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MainchainTNT20TokenBank.Contract.contract.Transact(opts, method, params...)
}

// AllDenoms is a free data retrieval call binding the contract method 0x588b1408.
//
// Solidity: function allDenoms(uint256 ) view returns(string)
func (_MainchainTNT20TokenBank *MainchainTNT20TokenBankCaller) AllDenoms(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _MainchainTNT20TokenBank.contract.Call(opts, &out, "allDenoms", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AllDenoms is a free data retrieval call binding the contract method 0x588b1408.
//
// Solidity: function allDenoms(uint256 ) view returns(string)
func (_MainchainTNT20TokenBank *MainchainTNT20TokenBankSession) AllDenoms(arg0 *big.Int) (string, error) {
	return _MainchainTNT20TokenBank.Contract.AllDenoms(&_MainchainTNT20TokenBank.CallOpts, arg0)
}

// AllDenoms is a free data retrieval call binding the contract method 0x588b1408.
//
// Solidity: function allDenoms(uint256 ) view returns(string)
func (_MainchainTNT20TokenBank *MainchainTNT20TokenBankCallerSession) AllDenoms(arg0 *big.Int) (string, error) {
	return _MainchainTNT20TokenBank.Contract.AllDenoms(&_MainchainTNT20TokenBank.CallOpts, arg0)
}

// AllVouchers is a free data retrieval call binding the contract method 0x27ca4df1.
//
// Solidity: function allVouchers(uint256 ) view returns(address)
func (_MainchainTNT20TokenBank *MainchainTNT20TokenBankCaller) AllVouchers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MainchainTNT20TokenBank.contract.Call(opts, &out, "allVouchers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllVouchers is a free data retrieval call binding the contract method 0x27ca4df1.
//
// Solidity: function allVouchers(uint256 ) view returns(address)
func (_MainchainTNT20TokenBank *MainchainTNT20TokenBankSession) AllVouchers(arg0 *big.Int) (common.Address, error) {
	return _MainchainTNT20TokenBank.Contract.AllVouchers(&_MainchainTNT20TokenBank.CallOpts, arg0)
}

// AllVouchers is a free data retrieval call binding the contract method 0x27ca4df1.
//
// Solidity: function allVouchers(uint256 ) view returns(address)
func (_MainchainTNT20TokenBank *MainchainTNT20TokenBankCallerSession) AllVouchers(arg0 *big.Int) (common.Address, error) {
	return _MainchainTNT20TokenBank.Contract.AllVouchers(&_MainchainTNT20TokenBank.CallOpts, arg0)
}

// DenomToVoucherLookup is a free data retrieval call binding the contract method 0x1527b14d.
//
// Solidity: function denomToVoucherLookup(string ) view returns(address contractAddress, bool exists)
func (_MainchainTNT20TokenBank *MainchainTNT20TokenBankCaller) DenomToVoucherLookup(opts *bind.CallOpts, arg0 string) (struct {
	ContractAddress common.Address
	Exists          bool
}, error) {
	var out []interface{}
	err := _MainchainTNT20TokenBank.contract.Call(opts, &out, "denomToVoucherLookup", arg0)

	outstruct := new(struct {
		ContractAddress common.Address
		Exists          bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ContractAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Exists = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// DenomToVoucherLookup is a free data retrieval call binding the contract method 0x1527b14d.
//
// Solidity: function denomToVoucherLookup(string ) view returns(address contractAddress, bool exists)
func (_MainchainTNT20TokenBank *MainchainTNT20TokenBankSession) DenomToVoucherLookup(arg0 string) (struct {
	ContractAddress common.Address
	Exists          bool
}, error) {
	return _MainchainTNT20TokenBank.Contract.DenomToVoucherLookup(&_MainchainTNT20TokenBank.CallOpts, arg0)
}

// DenomToVoucherLookup is a free data retrieval call binding the contract method 0x1527b14d.
//
// Solidity: function denomToVoucherLookup(string ) view returns(address contractAddress, bool exists)
func (_MainchainTNT20TokenBank *MainchainTNT20TokenBankCallerSession) DenomToVoucherLookup(arg0 string) (struct {
	ContractAddress common.Address
	Exists          bool
}, error) {
	return _MainchainTNT20TokenBank.Contract.DenomToVoucherLookup(&_MainchainTNT20TokenBank.CallOpts, arg0)
}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(string denom) view returns(bool)
func (_MainchainTNT20TokenBank *MainchainTNT20TokenBankCaller) Exists(opts *bind.CallOpts, denom string) (bool, error) {
	var out []interface{}
	err := _MainchainTNT20TokenBank.contract.Call(opts, &out, "exists", denom)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(string denom) view returns(bool)
func (_MainchainTNT20TokenBank *MainchainTNT20TokenBankSession) Exists(denom string) (bool, error) {
	return _MainchainTNT20TokenBank.Contract.Exists(&_MainchainTNT20TokenBank.CallOpts, denom)
}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(string denom) view returns(bool)
func (_MainchainTNT20TokenBank *MainchainTNT20TokenBankCallerSession) Exists(denom string) (bool, error) {
	return _MainchainTNT20TokenBank.Contract.Exists(&_MainchainTNT20TokenBank.CallOpts, denom)
}

// Exists0 is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address voucherAddress) view returns(bool)
func (_MainchainTNT20TokenBank *MainchainTNT20TokenBankCaller) Exists0(opts *bind.CallOpts, voucherAddress common.Address) (bool, error) {
	var out []interface{}
	err := _MainchainTNT20TokenBank.contract.Call(opts, &out, "exists0", voucherAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists0 is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address voucherAddress) view returns(bool)
func (_MainchainTNT20TokenBank *MainchainTNT20TokenBankSession) Exists0(voucherAddress common.Address) (bool, error) {
	return _MainchainTNT20TokenBank.Contract.Exists0(&_MainchainTNT20TokenBank.CallOpts, voucherAddress)
}

// Exists0 is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address voucherAddress) view returns(bool)
func (_MainchainTNT20TokenBank *MainchainTNT20TokenBankCallerSession) Exists0(voucherAddress common.Address) (bool, error) {
	return _MainchainTNT20TokenBank.Contract.Exists0(&_MainchainTNT20TokenBank.CallOpts, voucherAddress)
}

// GetDenom is a free data retrieval call binding the contract method 0xebda9962.
//
// Solidity: function getDenom(address voucherContractAddr) view returns(string)
func (_MainchainTNT20TokenBank *MainchainTNT20TokenBankCaller) GetDenom(opts *bind.CallOpts, voucherContractAddr common.Address) (string, error) {
	var out []interface{}
	err := _MainchainTNT20TokenBank.contract.Call(opts, &out, "getDenom", voucherContractAddr)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetDenom is a free data retrieval call binding the contract method 0xebda9962.
//
// Solidity: function getDenom(address voucherContractAddr) view returns(string)
func (_MainchainTNT20TokenBank *MainchainTNT20TokenBankSession) GetDenom(voucherContractAddr common.Address) (string, error) {
	return _MainchainTNT20TokenBank.Contract.GetDenom(&_MainchainTNT20TokenBank.CallOpts, voucherContractAddr)
}

// GetDenom is a free data retrieval call binding the contract method 0xebda9962.
//
// Solidity: function getDenom(address voucherContractAddr) view returns(string)
func (_MainchainTNT20TokenBank *MainchainTNT20TokenBankCallerSession) GetDenom(voucherContractAddr common.Address) (string, error) {
	return _MainchainTNT20TokenBank.Contract.GetDenom(&_MainchainTNT20TokenBank.CallOpts, voucherContractAddr)
}

// GetVoucher is a free data retrieval call binding the contract method 0xa2cc6981.
//
// Solidity: function getVoucher(string denom) view returns(address)
func (_MainchainTNT20TokenBank *MainchainTNT20TokenBankCaller) GetVoucher(opts *bind.CallOpts, denom string) (common.Address, error) {
	var out []interface{}
	err := _MainchainTNT20TokenBank.contract.Call(opts, &out, "getVoucher", denom)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVoucher is a free data retrieval call binding the contract method 0xa2cc6981.
//
// Solidity: function getVoucher(string denom) view returns(address)
func (_MainchainTNT20TokenBank *MainchainTNT20TokenBankSession) GetVoucher(denom string) (common.Address, error) {
	return _MainchainTNT20TokenBank.Contract.GetVoucher(&_MainchainTNT20TokenBank.CallOpts, denom)
}

// GetVoucher is a free data retrieval call binding the contract method 0xa2cc6981.
//
// Solidity: function getVoucher(string denom) view returns(address)
func (_MainchainTNT20TokenBank *MainchainTNT20TokenBankCallerSession) GetVoucher(denom string) (common.Address, error) {
	return _MainchainTNT20TokenBank.Contract.GetVoucher(&_MainchainTNT20TokenBank.CallOpts, denom)
}

// VoucherAddressToDenomLookup is a free data retrieval call binding the contract method 0x60569b5e.
//
// Solidity: function voucherAddressToDenomLookup(address ) view returns(string denom, bool exists)
func (_MainchainTNT20TokenBank *MainchainTNT20TokenBankCaller) VoucherAddressToDenomLookup(opts *bind.CallOpts, arg0 common.Address) (struct {
	Denom  string
	Exists bool
}, error) {
	var out []interface{}
	err := _MainchainTNT20TokenBank.contract.Call(opts, &out, "voucherAddressToDenomLookup", arg0)

	outstruct := new(struct {
		Denom  string
		Exists bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Denom = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Exists = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// VoucherAddressToDenomLookup is a free data retrieval call binding the contract method 0x60569b5e.
//
// Solidity: function voucherAddressToDenomLookup(address ) view returns(string denom, bool exists)
func (_MainchainTNT20TokenBank *MainchainTNT20TokenBankSession) VoucherAddressToDenomLookup(arg0 common.Address) (struct {
	Denom  string
	Exists bool
}, error) {
	return _MainchainTNT20TokenBank.Contract.VoucherAddressToDenomLookup(&_MainchainTNT20TokenBank.CallOpts, arg0)
}

// VoucherAddressToDenomLookup is a free data retrieval call binding the contract method 0x60569b5e.
//
// Solidity: function voucherAddressToDenomLookup(address ) view returns(string denom, bool exists)
func (_MainchainTNT20TokenBank *MainchainTNT20TokenBankCallerSession) VoucherAddressToDenomLookup(arg0 common.Address) (struct {
	Denom  string
	Exists bool
}, error) {
	return _MainchainTNT20TokenBank.Contract.VoucherAddressToDenomLookup(&_MainchainTNT20TokenBank.CallOpts, arg0)
}

// SendToChain is a paid mutator transaction binding the contract method 0xfaf9aa6e.
//
// Solidity: function sendToChain(uint256 targetChainID, address TNT20Contract, address subchainTokenReceiver, uint256 amount) payable returns()
func (_MainchainTNT20TokenBank *MainchainTNT20TokenBankTransactor) SendToChain(opts *bind.TransactOpts, targetChainID *big.Int, TNT20Contract common.Address, subchainTokenReceiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MainchainTNT20TokenBank.contract.Transact(opts, "sendToChain", targetChainID, TNT20Contract, subchainTokenReceiver, amount)
}

// SendToChain is a paid mutator transaction binding the contract method 0xfaf9aa6e.
//
// Solidity: function sendToChain(uint256 targetChainID, address TNT20Contract, address subchainTokenReceiver, uint256 amount) payable returns()
func (_MainchainTNT20TokenBank *MainchainTNT20TokenBankSession) SendToChain(targetChainID *big.Int, TNT20Contract common.Address, subchainTokenReceiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MainchainTNT20TokenBank.Contract.SendToChain(&_MainchainTNT20TokenBank.TransactOpts, targetChainID, TNT20Contract, subchainTokenReceiver, amount)
}

// SendToChain is a paid mutator transaction binding the contract method 0xfaf9aa6e.
//
// Solidity: function sendToChain(uint256 targetChainID, address TNT20Contract, address subchainTokenReceiver, uint256 amount) payable returns()
func (_MainchainTNT20TokenBank *MainchainTNT20TokenBankTransactorSession) SendToChain(targetChainID *big.Int, TNT20Contract common.Address, subchainTokenReceiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MainchainTNT20TokenBank.Contract.SendToChain(&_MainchainTNT20TokenBank.TransactOpts, targetChainID, TNT20Contract, subchainTokenReceiver, amount)
}

// MainchainTNT20TokenBankTNT20TokenLockedIterator is returned from FilterTNT20TokenLocked and is used to iterate over the raw logs and unpacked data for TNT20TokenLocked events raised by the MainchainTNT20TokenBank contract.
type MainchainTNT20TokenBankTNT20TokenLockedIterator struct {
	Event *MainchainTNT20TokenBankTNT20TokenLocked // Event containing the contract specifics and raw log

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
func (it *MainchainTNT20TokenBankTNT20TokenLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainchainTNT20TokenBankTNT20TokenLocked)
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
		it.Event = new(MainchainTNT20TokenBankTNT20TokenLocked)
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
func (it *MainchainTNT20TokenBankTNT20TokenLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainchainTNT20TokenBankTNT20TokenLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainchainTNT20TokenBankTNT20TokenLocked represents a TNT20TokenLocked event raised by the MainchainTNT20TokenBank contract.
type MainchainTNT20TokenBankTNT20TokenLocked struct {
	TargetChainID         *big.Int
	MainchainTokenSender  common.Address
	SubchainTokenReceiver common.Address
	LockedAmount          *big.Int
	TNT20Contract         common.Address
	Name                  string
	Symbol                string
	Decimal               uint8
	Nonce                 *big.Int
	Denom                 string
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterTNT20TokenLocked is a free log retrieval operation binding the contract event 0xe5d2c5efdbbd7c86071c36e6c82f176bac6830073d3498f8641318a0763875f8.
//
// Solidity: event TNT20TokenLocked(uint256 targetChainID, address mainchainTokenSender, address subchainTokenReceiver, uint256 lockedAmount, address TNT20Contract, string name, string symbol, uint8 decimal, uint256 nonce, string denom)
func (_MainchainTNT20TokenBank *MainchainTNT20TokenBankFilterer) FilterTNT20TokenLocked(opts *bind.FilterOpts) (*MainchainTNT20TokenBankTNT20TokenLockedIterator, error) {

	logs, sub, err := _MainchainTNT20TokenBank.contract.FilterLogs(opts, "TNT20TokenLocked")
	if err != nil {
		return nil, err
	}
	return &MainchainTNT20TokenBankTNT20TokenLockedIterator{contract: _MainchainTNT20TokenBank.contract, event: "TNT20TokenLocked", logs: logs, sub: sub}, nil
}

// WatchTNT20TokenLocked is a free log subscription operation binding the contract event 0xe5d2c5efdbbd7c86071c36e6c82f176bac6830073d3498f8641318a0763875f8.
//
// Solidity: event TNT20TokenLocked(uint256 targetChainID, address mainchainTokenSender, address subchainTokenReceiver, uint256 lockedAmount, address TNT20Contract, string name, string symbol, uint8 decimal, uint256 nonce, string denom)
func (_MainchainTNT20TokenBank *MainchainTNT20TokenBankFilterer) WatchTNT20TokenLocked(opts *bind.WatchOpts, sink chan<- *MainchainTNT20TokenBankTNT20TokenLocked) (event.Subscription, error) {

	logs, sub, err := _MainchainTNT20TokenBank.contract.WatchLogs(opts, "TNT20TokenLocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainchainTNT20TokenBankTNT20TokenLocked)
				if err := _MainchainTNT20TokenBank.contract.UnpackLog(event, "TNT20TokenLocked", log); err != nil {
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

// ParseTNT20TokenLocked is a log parse operation binding the contract event 0xe5d2c5efdbbd7c86071c36e6c82f176bac6830073d3498f8641318a0763875f8.
//
// Solidity: event TNT20TokenLocked(uint256 targetChainID, address mainchainTokenSender, address subchainTokenReceiver, uint256 lockedAmount, address TNT20Contract, string name, string symbol, uint8 decimal, uint256 nonce, string denom)
func (_MainchainTNT20TokenBank *MainchainTNT20TokenBankFilterer) ParseTNT20TokenLocked(log types.Log) (*MainchainTNT20TokenBankTNT20TokenLocked, error) {
	event := new(MainchainTNT20TokenBankTNT20TokenLocked)
	if err := _MainchainTNT20TokenBank.contract.UnpackLog(event, "TNT20TokenLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}