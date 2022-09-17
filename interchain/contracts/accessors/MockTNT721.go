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

// MockTNT721MetaData contains all meta data concerning the MockTNT721 contract.
var MockTNT721MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604080518082018252600b8082526a4d6f636b5f544e5437323160a81b602080840182815285518087019096529285528401528151919291620000589160009162000077565b5080516200006e90600190602084019062000077565b5050506200015a565b82805462000085906200011d565b90600052602060002090601f016020900481019282620000a95760008555620000f4565b82601f10620000c457805160ff1916838001178555620000f4565b82800160010185558215620000f4579182015b82811115620000f4578251825591602001919060010190620000d7565b506200010292915062000106565b5090565b5b8082111562000102576000815560010162000107565b600181811c908216806200013257607f821691505b602082108114156200015457634e487b7160e01b600052602260045260246000fd5b50919050565b61139a806200016a6000396000f3fe608060405234801561001057600080fd5b50600436106100ca5760003560e01c80636352211e1161007c5780636352211e1461018557806370a082311461019857806395d89b41146101b9578063a22cb465146101c1578063b88d4fde146101d4578063c87b56dd146101e7578063e985e9c5146101fa57600080fd5b806301ffc9a7146100cf57806306fdde03146100f7578063081812fc1461010c578063095ea7b31461013757806323b872dd1461014c57806340c10f191461015f57806342842e0e14610172575b600080fd5b6100e26100dd366004611079565b61020d565b60405190151581526020015b60405180910390f35b6100ff61025f565b6040516100ee9190611164565b61011f61011a3660046110b3565b6102f1565b6040516001600160a01b0390911681526020016100ee565b61014a61014536600461104f565b61037e565b005b61014a61015a366004610efb565b61048f565b61014a61016d36600461104f565b6104c0565b61014a610180366004610efb565b6104ce565b61011f6101933660046110b3565b6104e9565b6101ab6101a6366004610ead565b610560565b6040519081526020016100ee565b6100ff6105e7565b61014a6101cf366004611013565b6105f6565b61014a6101e2366004610f37565b610601565b6100ff6101f53660046110b3565b610639565b6100e2610208366004610ec8565b610736565b60006001600160e01b031982166380ac58cd60e01b148061023e57506001600160e01b03198216635b5e139f60e01b145b8061025957506301ffc9a760e01b6001600160e01b03198316145b92915050565b60606000805461026e90611289565b80601f016020809104026020016040519081016040528092919081815260200182805461029a90611289565b80156102e75780601f106102bc576101008083540402835291602001916102e7565b820191906000526020600020905b8154815290600101906020018083116102ca57829003601f168201915b5050505050905090565b60006102fc82610764565b6103625760405162461bcd60e51b815260206004820152602c60248201527f4552433732313a20617070726f76656420717565727920666f72206e6f6e657860448201526b34b9ba32b73a103a37b5b2b760a11b60648201526084015b60405180910390fd5b506000908152600460205260409020546001600160a01b031690565b6000610389826104e9565b9050806001600160a01b0316836001600160a01b031614156103f75760405162461bcd60e51b815260206004820152602160248201527f4552433732313a20617070726f76616c20746f2063757272656e74206f776e656044820152603960f91b6064820152608401610359565b336001600160a01b038216148061041357506104138133610736565b6104805760405162461bcd60e51b815260206004820152603860248201527f4552433732313a20617070726f76652063616c6c6572206973206e6f74206f776044820152771b995c881b9bdc88185c1c1c9bdd995908199bdc88185b1b60421b6064820152608401610359565b61048a8383610781565b505050565b61049933826107ef565b6104b55760405162461bcd60e51b8152600401610359906111c9565b61048a8383836108b9565b6104ca8282610a55565b5050565b61048a83838360405180602001604052806000815250610601565b6000818152600260205260408120546001600160a01b0316806102595760405162461bcd60e51b815260206004820152602960248201527f4552433732313a206f776e657220717565727920666f72206e6f6e657869737460448201526832b73a103a37b5b2b760b91b6064820152608401610359565b60006001600160a01b0382166105cb5760405162461bcd60e51b815260206004820152602a60248201527f4552433732313a2062616c616e636520717565727920666f7220746865207a65604482015269726f206164647265737360b01b6064820152608401610359565b506001600160a01b031660009081526003602052604090205490565b60606001805461026e90611289565b6104ca338383610b88565b61060b33836107ef565b6106275760405162461bcd60e51b8152600401610359906111c9565b61063384848484610c53565b50505050565b606061064482610764565b6106a85760405162461bcd60e51b815260206004820152602f60248201527f4552433732314d657461646174613a2055524920717565727920666f72206e6f60448201526e3732bc34b9ba32b73a103a37b5b2b760891b6064820152608401610359565b60006106e460408051808201909152601b81527f68747470733a2f2f746e743732312e6d657461646174612e696f2f0000000000602082015290565b90506000815111610704576040518060200160405280600081525061072f565b8061070e84610c86565b60405160200161071f9291906110f8565b6040516020818303038152906040525b9392505050565b6001600160a01b03918216600090815260056020908152604080832093909416825291909152205460ff1690565b6000908152600260205260409020546001600160a01b0316151590565b600081815260046020526040902080546001600160a01b0319166001600160a01b03841690811790915581906107b6826104e9565b6001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b60006107fa82610764565b61085b5760405162461bcd60e51b815260206004820152602c60248201527f4552433732313a206f70657261746f7220717565727920666f72206e6f6e657860448201526b34b9ba32b73a103a37b5b2b760a11b6064820152608401610359565b6000610866836104e9565b9050806001600160a01b0316846001600160a01b0316148061088d575061088d8185610736565b806108b15750836001600160a01b03166108a6846102f1565b6001600160a01b0316145b949350505050565b826001600160a01b03166108cc826104e9565b6001600160a01b0316146109305760405162461bcd60e51b815260206004820152602560248201527f4552433732313a207472616e736665722066726f6d20696e636f72726563742060448201526437bbb732b960d91b6064820152608401610359565b6001600160a01b0382166109925760405162461bcd60e51b8152602060048201526024808201527f4552433732313a207472616e7366657220746f20746865207a65726f206164646044820152637265737360e01b6064820152608401610359565b61099d600082610781565b6001600160a01b03831660009081526003602052604081208054600192906109c6908490611246565b90915550506001600160a01b03821660009081526003602052604081208054600192906109f490849061121a565b909155505060008181526002602052604080822080546001600160a01b0319166001600160a01b0386811691821790925591518493918716917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b6001600160a01b038216610aab5760405162461bcd60e51b815260206004820181905260248201527f4552433732313a206d696e7420746f20746865207a65726f20616464726573736044820152606401610359565b610ab481610764565b15610b015760405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e746564000000006044820152606401610359565b6001600160a01b0382166000908152600360205260408120805460019290610b2a90849061121a565b909155505060008181526002602052604080822080546001600160a01b0319166001600160a01b03861690811790915590518392907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b816001600160a01b0316836001600160a01b03161415610be65760405162461bcd60e51b815260206004820152601960248201527822a9219b99189d1030b8383937bb32903a379031b0b63632b960391b6044820152606401610359565b6001600160a01b03838116600081815260056020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b610c5e8484846108b9565b610c6a84848484610d84565b6106335760405162461bcd60e51b815260040161035990611177565b606081610caa5750506040805180820190915260018152600360fc1b602082015290565b8160005b8115610cd45780610cbe816112c4565b9150610ccd9050600a83611232565b9150610cae565b60008167ffffffffffffffff811115610cef57610cef611335565b6040519080825280601f01601f191660200182016040528015610d19576020820181803683370190505b5090505b84156108b157610d2e600183611246565b9150610d3b600a866112df565b610d4690603061121a565b60f81b818381518110610d5b57610d5b61131f565b60200101906001600160f81b031916908160001a905350610d7d600a86611232565b9450610d1d565b60006001600160a01b0384163b15610e8657604051630a85bd0160e11b81526001600160a01b0385169063150b7a0290610dc8903390899088908890600401611127565b602060405180830381600087803b158015610de257600080fd5b505af1925050508015610e12575060408051601f3d908101601f19168201909252610e0f91810190611096565b60015b610e6c573d808015610e40576040519150601f19603f3d011682016040523d82523d6000602084013e610e45565b606091505b508051610e645760405162461bcd60e51b815260040161035990611177565b805181602001fd5b6001600160e01b031916630a85bd0160e11b1490506108b1565b506001949350505050565b80356001600160a01b0381168114610ea857600080fd5b919050565b600060208284031215610ebf57600080fd5b61072f82610e91565b60008060408385031215610edb57600080fd5b610ee483610e91565b9150610ef260208401610e91565b90509250929050565b600080600060608486031215610f1057600080fd5b610f1984610e91565b9250610f2760208501610e91565b9150604084013590509250925092565b60008060008060808587031215610f4d57600080fd5b610f5685610e91565b9350610f6460208601610e91565b925060408501359150606085013567ffffffffffffffff80821115610f8857600080fd5b818701915087601f830112610f9c57600080fd5b813581811115610fae57610fae611335565b604051601f8201601f19908116603f01168101908382118183101715610fd657610fd6611335565b816040528281528a6020848701011115610fef57600080fd5b82602086016020830137600060208483010152809550505050505092959194509250565b6000806040838503121561102657600080fd5b61102f83610e91565b91506020830135801515811461104457600080fd5b809150509250929050565b6000806040838503121561106257600080fd5b61106b83610e91565b946020939093013593505050565b60006020828403121561108b57600080fd5b813561072f8161134b565b6000602082840312156110a857600080fd5b815161072f8161134b565b6000602082840312156110c557600080fd5b5035919050565b600081518084526110e481602086016020860161125d565b601f01601f19169290920160200192915050565b6000835161110a81846020880161125d565b83519083019061111e81836020880161125d565b01949350505050565b6001600160a01b038581168252841660208201526040810183905260806060820181905260009061115a908301846110cc565b9695505050505050565b60208152600061072f60208301846110cc565b60208082526032908201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560408201527131b2b4bb32b91034b6b83632b6b2b73a32b960711b606082015260800190565b60208082526031908201527f4552433732313a207472616e736665722063616c6c6572206973206e6f74206f6040820152701ddb995c881b9bdc88185c1c1c9bdd9959607a1b606082015260800190565b6000821982111561122d5761122d6112f3565b500190565b60008261124157611241611309565b500490565b600082821015611258576112586112f3565b500390565b60005b83811015611278578181015183820152602001611260565b838111156106335750506000910152565b600181811c9082168061129d57607f821691505b602082108114156112be57634e487b7160e01b600052602260045260246000fd5b50919050565b60006000198214156112d8576112d86112f3565b5060010190565b6000826112ee576112ee611309565b500690565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052601260045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160e01b03198116811461136157600080fd5b5056fea2646970667358221220a2c34fefdaa0037b736899b56a53027fe40c5d4afb9a4f81b004bafd6459edca64736f6c63430008070033",
}

// MockTNT721ABI is the input ABI used to generate the binding from.
// Deprecated: Use MockTNT721MetaData.ABI instead.
var MockTNT721ABI = MockTNT721MetaData.ABI

// MockTNT721Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MockTNT721MetaData.Bin instead.
var MockTNT721Bin = MockTNT721MetaData.Bin

// DeployMockTNT721 deploys a new Ethereum contract, binding an instance of MockTNT721 to it.
func DeployMockTNT721(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MockTNT721, error) {
	parsed, err := MockTNT721MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockTNT721Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockTNT721{MockTNT721Caller: MockTNT721Caller{contract: contract}, MockTNT721Transactor: MockTNT721Transactor{contract: contract}, MockTNT721Filterer: MockTNT721Filterer{contract: contract}}, nil
}

// MockTNT721 is an auto generated Go binding around an Ethereum contract.
type MockTNT721 struct {
	MockTNT721Caller     // Read-only binding to the contract
	MockTNT721Transactor // Write-only binding to the contract
	MockTNT721Filterer   // Log filterer for contract events
}

// MockTNT721Caller is an auto generated read-only Go binding around an Ethereum contract.
type MockTNT721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockTNT721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type MockTNT721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockTNT721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockTNT721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockTNT721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockTNT721Session struct {
	Contract     *MockTNT721       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MockTNT721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockTNT721CallerSession struct {
	Contract *MockTNT721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MockTNT721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockTNT721TransactorSession struct {
	Contract     *MockTNT721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MockTNT721Raw is an auto generated low-level Go binding around an Ethereum contract.
type MockTNT721Raw struct {
	Contract *MockTNT721 // Generic contract binding to access the raw methods on
}

// MockTNT721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockTNT721CallerRaw struct {
	Contract *MockTNT721Caller // Generic read-only contract binding to access the raw methods on
}

// MockTNT721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockTNT721TransactorRaw struct {
	Contract *MockTNT721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMockTNT721 creates a new instance of MockTNT721, bound to a specific deployed contract.
func NewMockTNT721(address common.Address, backend bind.ContractBackend) (*MockTNT721, error) {
	contract, err := bindMockTNT721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockTNT721{MockTNT721Caller: MockTNT721Caller{contract: contract}, MockTNT721Transactor: MockTNT721Transactor{contract: contract}, MockTNT721Filterer: MockTNT721Filterer{contract: contract}}, nil
}

// NewMockTNT721Caller creates a new read-only instance of MockTNT721, bound to a specific deployed contract.
func NewMockTNT721Caller(address common.Address, caller bind.ContractCaller) (*MockTNT721Caller, error) {
	contract, err := bindMockTNT721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockTNT721Caller{contract: contract}, nil
}

// NewMockTNT721Transactor creates a new write-only instance of MockTNT721, bound to a specific deployed contract.
func NewMockTNT721Transactor(address common.Address, transactor bind.ContractTransactor) (*MockTNT721Transactor, error) {
	contract, err := bindMockTNT721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockTNT721Transactor{contract: contract}, nil
}

// NewMockTNT721Filterer creates a new log filterer instance of MockTNT721, bound to a specific deployed contract.
func NewMockTNT721Filterer(address common.Address, filterer bind.ContractFilterer) (*MockTNT721Filterer, error) {
	contract, err := bindMockTNT721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockTNT721Filterer{contract: contract}, nil
}

// bindMockTNT721 binds a generic wrapper to an already deployed contract.
func bindMockTNT721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MockTNT721ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockTNT721 *MockTNT721Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockTNT721.Contract.MockTNT721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockTNT721 *MockTNT721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockTNT721.Contract.MockTNT721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockTNT721 *MockTNT721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockTNT721.Contract.MockTNT721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockTNT721 *MockTNT721CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockTNT721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockTNT721 *MockTNT721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockTNT721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockTNT721 *MockTNT721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockTNT721.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_MockTNT721 *MockTNT721Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MockTNT721.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_MockTNT721 *MockTNT721Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _MockTNT721.Contract.BalanceOf(&_MockTNT721.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_MockTNT721 *MockTNT721CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _MockTNT721.Contract.BalanceOf(&_MockTNT721.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_MockTNT721 *MockTNT721Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MockTNT721.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_MockTNT721 *MockTNT721Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _MockTNT721.Contract.GetApproved(&_MockTNT721.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_MockTNT721 *MockTNT721CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _MockTNT721.Contract.GetApproved(&_MockTNT721.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_MockTNT721 *MockTNT721Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _MockTNT721.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_MockTNT721 *MockTNT721Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _MockTNT721.Contract.IsApprovedForAll(&_MockTNT721.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_MockTNT721 *MockTNT721CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _MockTNT721.Contract.IsApprovedForAll(&_MockTNT721.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MockTNT721 *MockTNT721Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MockTNT721.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MockTNT721 *MockTNT721Session) Name() (string, error) {
	return _MockTNT721.Contract.Name(&_MockTNT721.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MockTNT721 *MockTNT721CallerSession) Name() (string, error) {
	return _MockTNT721.Contract.Name(&_MockTNT721.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_MockTNT721 *MockTNT721Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MockTNT721.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_MockTNT721 *MockTNT721Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _MockTNT721.Contract.OwnerOf(&_MockTNT721.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_MockTNT721 *MockTNT721CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _MockTNT721.Contract.OwnerOf(&_MockTNT721.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MockTNT721 *MockTNT721Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _MockTNT721.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MockTNT721 *MockTNT721Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MockTNT721.Contract.SupportsInterface(&_MockTNT721.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MockTNT721 *MockTNT721CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MockTNT721.Contract.SupportsInterface(&_MockTNT721.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MockTNT721 *MockTNT721Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MockTNT721.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MockTNT721 *MockTNT721Session) Symbol() (string, error) {
	return _MockTNT721.Contract.Symbol(&_MockTNT721.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MockTNT721 *MockTNT721CallerSession) Symbol() (string, error) {
	return _MockTNT721.Contract.Symbol(&_MockTNT721.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_MockTNT721 *MockTNT721Caller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _MockTNT721.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_MockTNT721 *MockTNT721Session) TokenURI(tokenId *big.Int) (string, error) {
	return _MockTNT721.Contract.TokenURI(&_MockTNT721.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_MockTNT721 *MockTNT721CallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _MockTNT721.Contract.TokenURI(&_MockTNT721.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_MockTNT721 *MockTNT721Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MockTNT721.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_MockTNT721 *MockTNT721Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MockTNT721.Contract.Approve(&_MockTNT721.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_MockTNT721 *MockTNT721TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MockTNT721.Contract.Approve(&_MockTNT721.TransactOpts, to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 tokenID) returns()
func (_MockTNT721 *MockTNT721Transactor) Mint(opts *bind.TransactOpts, account common.Address, tokenID *big.Int) (*types.Transaction, error) {
	return _MockTNT721.contract.Transact(opts, "mint", account, tokenID)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 tokenID) returns()
func (_MockTNT721 *MockTNT721Session) Mint(account common.Address, tokenID *big.Int) (*types.Transaction, error) {
	return _MockTNT721.Contract.Mint(&_MockTNT721.TransactOpts, account, tokenID)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 tokenID) returns()
func (_MockTNT721 *MockTNT721TransactorSession) Mint(account common.Address, tokenID *big.Int) (*types.Transaction, error) {
	return _MockTNT721.Contract.Mint(&_MockTNT721.TransactOpts, account, tokenID)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_MockTNT721 *MockTNT721Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MockTNT721.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_MockTNT721 *MockTNT721Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MockTNT721.Contract.SafeTransferFrom(&_MockTNT721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_MockTNT721 *MockTNT721TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MockTNT721.Contract.SafeTransferFrom(&_MockTNT721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_MockTNT721 *MockTNT721Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _MockTNT721.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_MockTNT721 *MockTNT721Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _MockTNT721.Contract.SafeTransferFrom0(&_MockTNT721.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_MockTNT721 *MockTNT721TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _MockTNT721.Contract.SafeTransferFrom0(&_MockTNT721.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MockTNT721 *MockTNT721Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _MockTNT721.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MockTNT721 *MockTNT721Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _MockTNT721.Contract.SetApprovalForAll(&_MockTNT721.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MockTNT721 *MockTNT721TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _MockTNT721.Contract.SetApprovalForAll(&_MockTNT721.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_MockTNT721 *MockTNT721Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MockTNT721.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_MockTNT721 *MockTNT721Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MockTNT721.Contract.TransferFrom(&_MockTNT721.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_MockTNT721 *MockTNT721TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MockTNT721.Contract.TransferFrom(&_MockTNT721.TransactOpts, from, to, tokenId)
}

// MockTNT721ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MockTNT721 contract.
type MockTNT721ApprovalIterator struct {
	Event *MockTNT721Approval // Event containing the contract specifics and raw log

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
func (it *MockTNT721ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockTNT721Approval)
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
		it.Event = new(MockTNT721Approval)
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
func (it *MockTNT721ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockTNT721ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockTNT721Approval represents a Approval event raised by the MockTNT721 contract.
type MockTNT721Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_MockTNT721 *MockTNT721Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*MockTNT721ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _MockTNT721.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MockTNT721ApprovalIterator{contract: _MockTNT721.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_MockTNT721 *MockTNT721Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MockTNT721Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _MockTNT721.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockTNT721Approval)
				if err := _MockTNT721.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_MockTNT721 *MockTNT721Filterer) ParseApproval(log types.Log) (*MockTNT721Approval, error) {
	event := new(MockTNT721Approval)
	if err := _MockTNT721.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockTNT721ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the MockTNT721 contract.
type MockTNT721ApprovalForAllIterator struct {
	Event *MockTNT721ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *MockTNT721ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockTNT721ApprovalForAll)
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
		it.Event = new(MockTNT721ApprovalForAll)
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
func (it *MockTNT721ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockTNT721ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockTNT721ApprovalForAll represents a ApprovalForAll event raised by the MockTNT721 contract.
type MockTNT721ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_MockTNT721 *MockTNT721Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*MockTNT721ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _MockTNT721.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &MockTNT721ApprovalForAllIterator{contract: _MockTNT721.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_MockTNT721 *MockTNT721Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *MockTNT721ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _MockTNT721.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockTNT721ApprovalForAll)
				if err := _MockTNT721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_MockTNT721 *MockTNT721Filterer) ParseApprovalForAll(log types.Log) (*MockTNT721ApprovalForAll, error) {
	event := new(MockTNT721ApprovalForAll)
	if err := _MockTNT721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockTNT721TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MockTNT721 contract.
type MockTNT721TransferIterator struct {
	Event *MockTNT721Transfer // Event containing the contract specifics and raw log

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
func (it *MockTNT721TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockTNT721Transfer)
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
		it.Event = new(MockTNT721Transfer)
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
func (it *MockTNT721TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockTNT721TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockTNT721Transfer represents a Transfer event raised by the MockTNT721 contract.
type MockTNT721Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_MockTNT721 *MockTNT721Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*MockTNT721TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _MockTNT721.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MockTNT721TransferIterator{contract: _MockTNT721.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_MockTNT721 *MockTNT721Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MockTNT721Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _MockTNT721.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockTNT721Transfer)
				if err := _MockTNT721.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_MockTNT721 *MockTNT721Filterer) ParseTransfer(log types.Log) (*MockTNT721Transfer, error) {
	event := new(MockTNT721Transfer)
	if err := _MockTNT721.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
