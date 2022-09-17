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

// MockTNT1155MetaData contains all meta data concerning the MockTNT1155 contract.
var MockTNT1155MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040805160208101909152600081526200002c8162000033565b506200012f565b8051620000489060029060208401906200004c565b5050565b8280546200005a90620000f2565b90600052602060002090601f0160209004810192826200007e5760008555620000c9565b82601f106200009957805160ff1916838001178555620000c9565b82800160010185558215620000c9579182015b82811115620000c9578251825591602001919060010190620000ac565b50620000d7929150620000db565b5090565b5b80821115620000d75760008155600101620000dc565b600181811c908216806200010757607f821691505b602082108114156200012957634e487b7160e01b600052602260045260246000fd5b50919050565b611689806200013f6000396000f3fe608060405234801561001057600080fd5b50600436106100825760003560e01c8062fdd58e1461008757806301ffc9a7146100ad5780630e89341c146100d05780632eb2c2d6146100f05780634e1273f414610105578063731133e914610125578063a22cb46514610138578063e985e9c51461014b578063f242432a14610187575b600080fd5b61009a610095366004611022565b61019a565b6040519081526020015b60405180910390f35b6100c06100bb36600461117c565b610231565b60405190151581526020016100a4565b6100e36100de3660046111bd565b610283565b6040516100a49190611366565b6101036100fe366004610ed9565b6102b7565b005b6101186101133660046110ac565b61034e565b6040516100a49190611325565b61010361013336600461104c565b610477565b610103610146366004610fe6565b610489565b6100c0610159366004610ea6565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205460ff1690565b610103610195366004610f82565b610498565b60006001600160a01b03831661020b5760405162461bcd60e51b815260206004820152602b60248201527f455243313135353a2062616c616e636520717565727920666f7220746865207a60448201526a65726f206164647265737360a81b60648201526084015b60405180910390fd5b506000908152602081815260408083206001600160a01b03949094168352929052205490565b60006001600160e01b03198216636cdb3d1360e11b148061026257506001600160e01b031982166303a24d0760e21b145b8061027d57506301ffc9a760e01b6001600160e01b03198316145b92915050565b606060006102908361051f565b6040516020016102a0919061123d565b60408051601f198184030181529190529392505050565b6001600160a01b0385163314806102d357506102d38533610159565b61033a5760405162461bcd60e51b815260206004820152603260248201527f455243313135353a207472616e736665722063616c6c6572206973206e6f74206044820152711bdddb995c881b9bdc88185c1c1c9bdd995960721b6064820152608401610202565b6103478585858585610624565b5050505050565b606081518351146103b35760405162461bcd60e51b815260206004820152602960248201527f455243313135353a206163636f756e747320616e6420696473206c656e677468604482015268040dad2e6dac2e8c6d60bb1b6064820152608401610202565b600083516001600160401b038111156103ce576103ce61157f565b6040519080825280602002602001820160405280156103f7578160200160208202803683370190505b50905060005b845181101561046f5761044285828151811061041b5761041b611569565b602002602001015185838151811061043557610435611569565b602002602001015161019a565b82828151811061045457610454611569565b60209081029190910101526104688161150e565b90506103fd565b509392505050565b61048384848484610801565b50505050565b610494338383610915565b5050565b6001600160a01b0385163314806104b457506104b48533610159565b6105125760405162461bcd60e51b815260206004820152602960248201527f455243313135353a2063616c6c6572206973206e6f74206f776e6572206e6f7260448201526808185c1c1c9bdd995960ba1b6064820152608401610202565b61034785858585856109f6565b6060816105435750506040805180820190915260018152600360fc1b602082015290565b8160005b811561056d57806105578161150e565b91506105669050600a8361148b565b9150610547565b6000816001600160401b038111156105875761058761157f565b6040519080825280601f01601f1916602001820160405280156105b1576020820181803683370190505b5090505b841561061c576105c660018361149f565b91506105d3600a86611529565b6105de906030611473565b60f81b8183815181106105f3576105f3611569565b60200101906001600160f81b031916908160001a905350610615600a8661148b565b94506105b5565b949350505050565b81518351146106865760405162461bcd60e51b815260206004820152602860248201527f455243313135353a2069647320616e6420616d6f756e7473206c656e677468206044820152670dad2e6dac2e8c6d60c31b6064820152608401610202565b6001600160a01b0384166106ac5760405162461bcd60e51b8152600401610202906113c1565b3360005b84518110156107935760008582815181106106cd576106cd611569565b6020026020010151905060008583815181106106eb576106eb611569565b602090810291909101810151600084815280835260408082206001600160a01b038e16835290935291909120549091508181101561073b5760405162461bcd60e51b815260040161020290611406565b6000838152602081815260408083206001600160a01b038e8116855292528083208585039055908b16825281208054849290610778908490611473565b925050819055505050508061078c9061150e565b90506106b0565b50846001600160a01b0316866001600160a01b0316826001600160a01b03167f4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb87876040516107e3929190611338565b60405180910390a46107f9818787878787610b20565b505050505050565b6001600160a01b0384166108615760405162461bcd60e51b815260206004820152602160248201527f455243313135353a206d696e7420746f20746865207a65726f206164647265736044820152607360f81b6064820152608401610202565b33600061086d85610c8b565b9050600061087a85610c8b565b90506000868152602081815260408083206001600160a01b038b168452909152812080548792906108ac908490611473565b909155505060408051878152602081018790526001600160a01b03808a1692600092918716917fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62910160405180910390a461090c83600089898989610cd6565b50505050505050565b816001600160a01b0316836001600160a01b031614156109895760405162461bcd60e51b815260206004820152602960248201527f455243313135353a2073657474696e6720617070726f76616c20737461747573604482015268103337b91039b2b63360b91b6064820152608401610202565b6001600160a01b03838116600081815260016020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b6001600160a01b038416610a1c5760405162461bcd60e51b8152600401610202906113c1565b336000610a2885610c8b565b90506000610a3585610c8b565b90506000868152602081815260408083206001600160a01b038c16845290915290205485811015610a785760405162461bcd60e51b815260040161020290611406565b6000878152602081815260408083206001600160a01b038d8116855292528083208985039055908a16825281208054889290610ab5908490611473565b909155505060408051888152602081018890526001600160a01b03808b16928c821692918816917fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62910160405180910390a4610b15848a8a8a8a8a610cd6565b505050505050505050565b6001600160a01b0384163b156107f95760405163bc197c8160e01b81526001600160a01b0385169063bc197c8190610b649089908990889088908890600401611282565b602060405180830381600087803b158015610b7e57600080fd5b505af1925050508015610bae575060408051601f3d908101601f19168201909252610bab918101906111a0565b60015b610c5b57610bba611595565b806308c379a01415610bf45750610bcf6115b1565b80610bda5750610bf6565b8060405162461bcd60e51b81526004016102029190611366565b505b60405162461bcd60e51b815260206004820152603460248201527f455243313135353a207472616e7366657220746f206e6f6e20455243313135356044820152732932b1b2b4bb32b91034b6b83632b6b2b73a32b960611b6064820152608401610202565b6001600160e01b0319811663bc197c8160e01b1461090c5760405162461bcd60e51b815260040161020290611379565b60408051600180825281830190925260609160009190602080830190803683370190505090508281600081518110610cc557610cc5611569565b602090810291909101015292915050565b6001600160a01b0384163b156107f95760405163f23a6e6160e01b81526001600160a01b0385169063f23a6e6190610d1a90899089908890889088906004016112e0565b602060405180830381600087803b158015610d3457600080fd5b505af1925050508015610d64575060408051601f3d908101601f19168201909252610d61918101906111a0565b60015b610d7057610bba611595565b6001600160e01b0319811663f23a6e6160e01b1461090c5760405162461bcd60e51b815260040161020290611379565b80356001600160a01b0381168114610db757600080fd5b919050565b600082601f830112610dcd57600080fd5b81356020610dda82611450565b604051610de782826114e2565b8381528281019150858301600585901b87018401881015610e0757600080fd5b60005b85811015610e2657813584529284019290840190600101610e0a565b5090979650505050505050565b600082601f830112610e4457600080fd5b81356001600160401b03811115610e5d57610e5d61157f565b604051610e74601f8301601f1916602001826114e2565b818152846020838601011115610e8957600080fd5b816020850160208301376000918101602001919091529392505050565b60008060408385031215610eb957600080fd5b610ec283610da0565b9150610ed060208401610da0565b90509250929050565b600080600080600060a08688031215610ef157600080fd5b610efa86610da0565b9450610f0860208701610da0565b935060408601356001600160401b0380821115610f2457600080fd5b610f3089838a01610dbc565b94506060880135915080821115610f4657600080fd5b610f5289838a01610dbc565b93506080880135915080821115610f6857600080fd5b50610f7588828901610e33565b9150509295509295909350565b600080600080600060a08688031215610f9a57600080fd5b610fa386610da0565b9450610fb160208701610da0565b9350604086013592506060860135915060808601356001600160401b03811115610fda57600080fd5b610f7588828901610e33565b60008060408385031215610ff957600080fd5b61100283610da0565b91506020830135801515811461101757600080fd5b809150509250929050565b6000806040838503121561103557600080fd5b61103e83610da0565b946020939093013593505050565b6000806000806080858703121561106257600080fd5b61106b85610da0565b9350602085013592506040850135915060608501356001600160401b0381111561109457600080fd5b6110a087828801610e33565b91505092959194509250565b600080604083850312156110bf57600080fd5b82356001600160401b03808211156110d657600080fd5b818501915085601f8301126110ea57600080fd5b813560206110f782611450565b60405161110482826114e2565b8381528281019150858301600585901b870184018b101561112457600080fd5b600096505b8487101561114e5761113a81610da0565b835260019690960195918301918301611129565b509650508601359250508082111561116557600080fd5b5061117285828601610dbc565b9150509250929050565b60006020828403121561118e57600080fd5b81356111998161163a565b9392505050565b6000602082840312156111b257600080fd5b81516111998161163a565b6000602082840312156111cf57600080fd5b5035919050565b600081518084526020808501945080840160005b83811015611206578151875295820195908201906001016111ea565b509495945050505050565b600081518084526112298160208601602086016114b6565b601f01601f19169290920160200192915050565b7f68747470733a2f2f746e74313135352e6d657461646174612e696f2f0000000081526000825161127581601c8501602087016114b6565b91909101601c0192915050565b6001600160a01b0386811682528516602082015260a0604082018190526000906112ae908301866111d6565b82810360608401526112c081866111d6565b905082810360808401526112d48185611211565b98975050505050505050565b6001600160a01b03868116825285166020820152604081018490526060810183905260a06080820181905260009061131a90830184611211565b979650505050505050565b60208152600061119960208301846111d6565b60408152600061134b60408301856111d6565b828103602084015261135d81856111d6565b95945050505050565b6020815260006111996020830184611211565b60208082526028908201527f455243313135353a204552433131353552656365697665722072656a656374656040820152676420746f6b656e7360c01b606082015260800190565b60208082526025908201527f455243313135353a207472616e7366657220746f20746865207a65726f206164604082015264647265737360d81b606082015260800190565b6020808252602a908201527f455243313135353a20696e73756666696369656e742062616c616e636520666f60408201526939103a3930b739b332b960b11b606082015260800190565b60006001600160401b038211156114695761146961157f565b5060051b60200190565b600082198211156114865761148661153d565b500190565b60008261149a5761149a611553565b500490565b6000828210156114b1576114b161153d565b500390565b60005b838110156114d15781810151838201526020016114b9565b838111156104835750506000910152565b601f8201601f191681016001600160401b03811182821017156115075761150761157f565b6040525050565b60006000198214156115225761152261153d565b5060010190565b60008261153857611538611553565b500690565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052601260045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b600060033d11156115ae5760046000803e5060005160e01c5b90565b600060443d10156115bf5790565b6040516003193d81016004833e81513d6001600160401b0381602484011181841117156115ee57505050505090565b82850191508151818111156116065750505050505090565b843d87010160208285010111156116205750505050505090565b61162f602082860101876114e2565b509095945050505050565b6001600160e01b03198116811461165057600080fd5b5056fea2646970667358221220d08cc5a4a913cec51ef87ab889e3cb65ea9ddf85eba1b1c6e1aa9c599e05616564736f6c63430008070033",
}

// MockTNT1155ABI is the input ABI used to generate the binding from.
// Deprecated: Use MockTNT1155MetaData.ABI instead.
var MockTNT1155ABI = MockTNT1155MetaData.ABI

// MockTNT1155Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MockTNT1155MetaData.Bin instead.
var MockTNT1155Bin = MockTNT1155MetaData.Bin

// DeployMockTNT1155 deploys a new Ethereum contract, binding an instance of MockTNT1155 to it.
func DeployMockTNT1155(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MockTNT1155, error) {
	parsed, err := MockTNT1155MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockTNT1155Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockTNT1155{MockTNT1155Caller: MockTNT1155Caller{contract: contract}, MockTNT1155Transactor: MockTNT1155Transactor{contract: contract}, MockTNT1155Filterer: MockTNT1155Filterer{contract: contract}}, nil
}

// MockTNT1155 is an auto generated Go binding around an Ethereum contract.
type MockTNT1155 struct {
	MockTNT1155Caller     // Read-only binding to the contract
	MockTNT1155Transactor // Write-only binding to the contract
	MockTNT1155Filterer   // Log filterer for contract events
}

// MockTNT1155Caller is an auto generated read-only Go binding around an Ethereum contract.
type MockTNT1155Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockTNT1155Transactor is an auto generated write-only Go binding around an Ethereum contract.
type MockTNT1155Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockTNT1155Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockTNT1155Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockTNT1155Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockTNT1155Session struct {
	Contract     *MockTNT1155      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MockTNT1155CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockTNT1155CallerSession struct {
	Contract *MockTNT1155Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MockTNT1155TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockTNT1155TransactorSession struct {
	Contract     *MockTNT1155Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MockTNT1155Raw is an auto generated low-level Go binding around an Ethereum contract.
type MockTNT1155Raw struct {
	Contract *MockTNT1155 // Generic contract binding to access the raw methods on
}

// MockTNT1155CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockTNT1155CallerRaw struct {
	Contract *MockTNT1155Caller // Generic read-only contract binding to access the raw methods on
}

// MockTNT1155TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockTNT1155TransactorRaw struct {
	Contract *MockTNT1155Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMockTNT1155 creates a new instance of MockTNT1155, bound to a specific deployed contract.
func NewMockTNT1155(address common.Address, backend bind.ContractBackend) (*MockTNT1155, error) {
	contract, err := bindMockTNT1155(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockTNT1155{MockTNT1155Caller: MockTNT1155Caller{contract: contract}, MockTNT1155Transactor: MockTNT1155Transactor{contract: contract}, MockTNT1155Filterer: MockTNT1155Filterer{contract: contract}}, nil
}

// NewMockTNT1155Caller creates a new read-only instance of MockTNT1155, bound to a specific deployed contract.
func NewMockTNT1155Caller(address common.Address, caller bind.ContractCaller) (*MockTNT1155Caller, error) {
	contract, err := bindMockTNT1155(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockTNT1155Caller{contract: contract}, nil
}

// NewMockTNT1155Transactor creates a new write-only instance of MockTNT1155, bound to a specific deployed contract.
func NewMockTNT1155Transactor(address common.Address, transactor bind.ContractTransactor) (*MockTNT1155Transactor, error) {
	contract, err := bindMockTNT1155(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockTNT1155Transactor{contract: contract}, nil
}

// NewMockTNT1155Filterer creates a new log filterer instance of MockTNT1155, bound to a specific deployed contract.
func NewMockTNT1155Filterer(address common.Address, filterer bind.ContractFilterer) (*MockTNT1155Filterer, error) {
	contract, err := bindMockTNT1155(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockTNT1155Filterer{contract: contract}, nil
}

// bindMockTNT1155 binds a generic wrapper to an already deployed contract.
func bindMockTNT1155(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MockTNT1155ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockTNT1155 *MockTNT1155Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockTNT1155.Contract.MockTNT1155Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockTNT1155 *MockTNT1155Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockTNT1155.Contract.MockTNT1155Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockTNT1155 *MockTNT1155Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockTNT1155.Contract.MockTNT1155Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockTNT1155 *MockTNT1155CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockTNT1155.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockTNT1155 *MockTNT1155TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockTNT1155.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockTNT1155 *MockTNT1155TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockTNT1155.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_MockTNT1155 *MockTNT1155Caller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MockTNT1155.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_MockTNT1155 *MockTNT1155Session) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _MockTNT1155.Contract.BalanceOf(&_MockTNT1155.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_MockTNT1155 *MockTNT1155CallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _MockTNT1155.Contract.BalanceOf(&_MockTNT1155.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_MockTNT1155 *MockTNT1155Caller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _MockTNT1155.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_MockTNT1155 *MockTNT1155Session) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _MockTNT1155.Contract.BalanceOfBatch(&_MockTNT1155.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_MockTNT1155 *MockTNT1155CallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _MockTNT1155.Contract.BalanceOfBatch(&_MockTNT1155.CallOpts, accounts, ids)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_MockTNT1155 *MockTNT1155Caller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _MockTNT1155.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_MockTNT1155 *MockTNT1155Session) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _MockTNT1155.Contract.IsApprovedForAll(&_MockTNT1155.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_MockTNT1155 *MockTNT1155CallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _MockTNT1155.Contract.IsApprovedForAll(&_MockTNT1155.CallOpts, account, operator)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MockTNT1155 *MockTNT1155Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _MockTNT1155.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MockTNT1155 *MockTNT1155Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MockTNT1155.Contract.SupportsInterface(&_MockTNT1155.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MockTNT1155 *MockTNT1155CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MockTNT1155.Contract.SupportsInterface(&_MockTNT1155.CallOpts, interfaceId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenID) pure returns(string)
func (_MockTNT1155 *MockTNT1155Caller) Uri(opts *bind.CallOpts, tokenID *big.Int) (string, error) {
	var out []interface{}
	err := _MockTNT1155.contract.Call(opts, &out, "uri", tokenID)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenID) pure returns(string)
func (_MockTNT1155 *MockTNT1155Session) Uri(tokenID *big.Int) (string, error) {
	return _MockTNT1155.Contract.Uri(&_MockTNT1155.CallOpts, tokenID)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenID) pure returns(string)
func (_MockTNT1155 *MockTNT1155CallerSession) Uri(tokenID *big.Int) (string, error) {
	return _MockTNT1155.Contract.Uri(&_MockTNT1155.CallOpts, tokenID)
}

// Mint is a paid mutator transaction binding the contract method 0x731133e9.
//
// Solidity: function mint(address to, uint256 id, uint256 amount, bytes data) returns()
func (_MockTNT1155 *MockTNT1155Transactor) Mint(opts *bind.TransactOpts, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _MockTNT1155.contract.Transact(opts, "mint", to, id, amount, data)
}

// Mint is a paid mutator transaction binding the contract method 0x731133e9.
//
// Solidity: function mint(address to, uint256 id, uint256 amount, bytes data) returns()
func (_MockTNT1155 *MockTNT1155Session) Mint(to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _MockTNT1155.Contract.Mint(&_MockTNT1155.TransactOpts, to, id, amount, data)
}

// Mint is a paid mutator transaction binding the contract method 0x731133e9.
//
// Solidity: function mint(address to, uint256 id, uint256 amount, bytes data) returns()
func (_MockTNT1155 *MockTNT1155TransactorSession) Mint(to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _MockTNT1155.Contract.Mint(&_MockTNT1155.TransactOpts, to, id, amount, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_MockTNT1155 *MockTNT1155Transactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _MockTNT1155.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_MockTNT1155 *MockTNT1155Session) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _MockTNT1155.Contract.SafeBatchTransferFrom(&_MockTNT1155.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_MockTNT1155 *MockTNT1155TransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _MockTNT1155.Contract.SafeBatchTransferFrom(&_MockTNT1155.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_MockTNT1155 *MockTNT1155Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _MockTNT1155.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_MockTNT1155 *MockTNT1155Session) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _MockTNT1155.Contract.SafeTransferFrom(&_MockTNT1155.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_MockTNT1155 *MockTNT1155TransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _MockTNT1155.Contract.SafeTransferFrom(&_MockTNT1155.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MockTNT1155 *MockTNT1155Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _MockTNT1155.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MockTNT1155 *MockTNT1155Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _MockTNT1155.Contract.SetApprovalForAll(&_MockTNT1155.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MockTNT1155 *MockTNT1155TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _MockTNT1155.Contract.SetApprovalForAll(&_MockTNT1155.TransactOpts, operator, approved)
}

// MockTNT1155ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the MockTNT1155 contract.
type MockTNT1155ApprovalForAllIterator struct {
	Event *MockTNT1155ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *MockTNT1155ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockTNT1155ApprovalForAll)
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
		it.Event = new(MockTNT1155ApprovalForAll)
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
func (it *MockTNT1155ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockTNT1155ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockTNT1155ApprovalForAll represents a ApprovalForAll event raised by the MockTNT1155 contract.
type MockTNT1155ApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_MockTNT1155 *MockTNT1155Filterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*MockTNT1155ApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _MockTNT1155.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &MockTNT1155ApprovalForAllIterator{contract: _MockTNT1155.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_MockTNT1155 *MockTNT1155Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *MockTNT1155ApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _MockTNT1155.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockTNT1155ApprovalForAll)
				if err := _MockTNT1155.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_MockTNT1155 *MockTNT1155Filterer) ParseApprovalForAll(log types.Log) (*MockTNT1155ApprovalForAll, error) {
	event := new(MockTNT1155ApprovalForAll)
	if err := _MockTNT1155.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockTNT1155TransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the MockTNT1155 contract.
type MockTNT1155TransferBatchIterator struct {
	Event *MockTNT1155TransferBatch // Event containing the contract specifics and raw log

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
func (it *MockTNT1155TransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockTNT1155TransferBatch)
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
		it.Event = new(MockTNT1155TransferBatch)
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
func (it *MockTNT1155TransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockTNT1155TransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockTNT1155TransferBatch represents a TransferBatch event raised by the MockTNT1155 contract.
type MockTNT1155TransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_MockTNT1155 *MockTNT1155Filterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*MockTNT1155TransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MockTNT1155.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MockTNT1155TransferBatchIterator{contract: _MockTNT1155.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_MockTNT1155 *MockTNT1155Filterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *MockTNT1155TransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MockTNT1155.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockTNT1155TransferBatch)
				if err := _MockTNT1155.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_MockTNT1155 *MockTNT1155Filterer) ParseTransferBatch(log types.Log) (*MockTNT1155TransferBatch, error) {
	event := new(MockTNT1155TransferBatch)
	if err := _MockTNT1155.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockTNT1155TransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the MockTNT1155 contract.
type MockTNT1155TransferSingleIterator struct {
	Event *MockTNT1155TransferSingle // Event containing the contract specifics and raw log

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
func (it *MockTNT1155TransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockTNT1155TransferSingle)
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
		it.Event = new(MockTNT1155TransferSingle)
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
func (it *MockTNT1155TransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockTNT1155TransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockTNT1155TransferSingle represents a TransferSingle event raised by the MockTNT1155 contract.
type MockTNT1155TransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_MockTNT1155 *MockTNT1155Filterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*MockTNT1155TransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MockTNT1155.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MockTNT1155TransferSingleIterator{contract: _MockTNT1155.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_MockTNT1155 *MockTNT1155Filterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *MockTNT1155TransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MockTNT1155.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockTNT1155TransferSingle)
				if err := _MockTNT1155.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_MockTNT1155 *MockTNT1155Filterer) ParseTransferSingle(log types.Log) (*MockTNT1155TransferSingle, error) {
	event := new(MockTNT1155TransferSingle)
	if err := _MockTNT1155.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockTNT1155URIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the MockTNT1155 contract.
type MockTNT1155URIIterator struct {
	Event *MockTNT1155URI // Event containing the contract specifics and raw log

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
func (it *MockTNT1155URIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockTNT1155URI)
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
		it.Event = new(MockTNT1155URI)
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
func (it *MockTNT1155URIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockTNT1155URIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockTNT1155URI represents a URI event raised by the MockTNT1155 contract.
type MockTNT1155URI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_MockTNT1155 *MockTNT1155Filterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*MockTNT1155URIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MockTNT1155.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &MockTNT1155URIIterator{contract: _MockTNT1155.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_MockTNT1155 *MockTNT1155Filterer) WatchURI(opts *bind.WatchOpts, sink chan<- *MockTNT1155URI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MockTNT1155.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockTNT1155URI)
				if err := _MockTNT1155.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_MockTNT1155 *MockTNT1155Filterer) ParseURI(log types.Log) (*MockTNT1155URI, error) {
	event := new(MockTNT1155URI)
	if err := _MockTNT1155.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
