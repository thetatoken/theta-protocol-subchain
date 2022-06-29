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

// SubchainTFuelTokenBankMetaData contains all meta data concerning the SubchainTFuelTokenBank contract.
var SubchainTFuelTokenBankMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internaltype\":\"uint256\",\"name\":\"mainchainid_\",\"type\":\"uint256\"}],\"statemutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internaltype\":\"uint256\",\"name\":\"mainchainid\",\"type\":\"uint256\"},{\"indexed\":false,\"internaltype\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"indexed\":false,\"internaltype\":\"address\",\"name\":\"voucherowner\",\"type\":\"address\"},{\"indexed\":false,\"internaltype\":\"address\",\"name\":\"mainchaintokenreceiver\",\"type\":\"address\"},{\"indexed\":false,\"internaltype\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internaltype\":\"uint256\",\"name\":\"voucherburnnonce\",\"type\":\"uint256\"}],\"name\":\"tfuelvoucherburned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internaltype\":\"uint256\",\"name\":\"mainchainid\",\"type\":\"uint256\"},{\"indexed\":false,\"internaltype\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"indexed\":false,\"internaltype\":\"address\",\"name\":\"voucherreceiver\",\"type\":\"address\"},{\"indexed\":false,\"internaltype\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internaltype\":\"uint256\",\"name\":\"mainchaintokenlocknonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internaltype\":\"uint256\",\"name\":\"vouchermintnonce\",\"type\":\"uint256\"}],\"name\":\"tfuelvoucherminted\",\"type\":\"event\"},{\"inputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"alldenoms\",\"outputs\":[{\"internaltype\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allvouchers\",\"outputs\":[{\"internaltype\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"denomtovoucherlookup\",\"outputs\":[{\"internaltype\":\"address\",\"name\":\"contractaddress\",\"type\":\"address\"},{\"internaltype\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"name\":\"exists\",\"outputs\":[{\"internaltype\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"address\",\"name\":\"voucheraddress\",\"type\":\"address\"}],\"name\":\"exists\",\"outputs\":[{\"internaltype\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"address\",\"name\":\"vouchercontractaddr\",\"type\":\"address\"}],\"name\":\"getdenom\",\"outputs\":[{\"internaltype\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"name\":\"getvoucher\",\"outputs\":[{\"internaltype\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mainchainid\",\"outputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenlocknonce\",\"outputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenunlocknonce\",\"outputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"voucheraddresstodenomlookup\",\"outputs\":[{\"internaltype\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internaltype\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voucherburnnonce\",\"outputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vouchermintnonce\",\"outputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"address\",\"name\":\"voucherreceiver\",\"type\":\"address\"},{\"internaltype\":\"uint256\",\"name\":\"mintamount\",\"type\":\"uint256\"},{\"internaltype\":\"uint256\",\"name\":\"mainchaintokenlocknonce\",\"type\":\"uint256\"}],\"name\":\"mintvouchers\",\"outputs\":[],\"statemutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"address\",\"name\":\"mainchaintokenreceiver\",\"type\":\"address\"}],\"name\":\"burnvouchers\",\"outputs\":[],\"statemutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161149e38038061149e83398101604081905261002f9161004e565b5060016000908155600281905560038190556004819055600555610067565b60006020828403121561006057600080fd5b5051919050565b611428806100766000396000f3fe6080604052600436106100e85760003560e01c8063588b14081161008a578063a4ddaaa611610059578063a4ddaaa6146102bc578063c27d927a146102dc578063ebda9962146102f2578063f6a3d24e1461031257600080fd5b8063588b14081461022c57806360569b5e146102595780637d0fb00d14610287578063a2cc69811461029c57600080fd5b8063261a323e116100c6578063261a323e1461019857806327ca4df1146101c857806333835be6146102005780635102e7601461021657600080fd5b8063073b9502146100ed5780631527b14d14610116578063163fe97214610182575b600080fd5b3480156100f957600080fd5b5061010360015481565b6040519081526020015b60405180910390f35b34801561012257600080fd5b50610163610131366004611028565b80516020818301810180516006825292820191909301209152546001600160a01b03811690600160a01b900460ff1682565b604080516001600160a01b03909316835290151560208301520161010d565b34801561018e57600080fd5b5061010360055481565b3480156101a457600080fd5b506101b86101b3366004611028565b61034e565b604051901515815260200161010d565b3480156101d457600080fd5b506101e86101e33660046110d9565b61038e565b6040516001600160a01b03909116815260200161010d565b34801561020c57600080fd5b5061010360035481565b34801561022257600080fd5b5061010360025481565b34801561023857600080fd5b5061024c6102473660046110d9565b6103b8565b60405161010d9190611194565b34801561026557600080fd5b50610279610274366004610fd3565b610464565b60405161010d9291906111a7565b61029a610295366004610fd3565b61050b565b005b3480156102a857600080fd5b506101e86102b7366004611028565b61060d565b3480156102c857600080fd5b5061029a6102d7366004610ff5565b61067e565b3480156102e857600080fd5b5061010360045481565b3480156102fe57600080fd5b5061024c61030d366004610fd3565b610856565b34801561031e57600080fd5b506101b861032d366004610fd3565b6001600160a01b031660009081526007602052604090206001015460ff1690565b60008061035a8361094d565b905060068160405161036c919061111e565b9081526040519081900360200190205460ff600160a01b909104169392505050565b6008818154811061039e57600080fd5b6000918252602090912001546001600160a01b0316905081565b600981815481106103c857600080fd5b9060005260206000200160009150905080546103e39061130c565b80601f016020809104026020016040519081016040528092919081815260200182805461040f9061130c565b801561045c5780601f106104315761010080835404028352916020019161045c565b820191906000526020600020905b81548152906001019060200180831161043f57829003601f168201915b505050505081565b60076020526000908152604090208054819061047f9061130c565b80601f01602080910402602001604051908101604052809291908181526020018280546104ab9061130c565b80156104f85780601f106104cd576101008083540402835291602001916104f8565b820191906000526020600020905b8154815290600101906020018083116104db57829003601f168201915b5050506001909301549192505060ff1682565b600260005414156105635760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064015b60405180910390fd5b600260005533346105738161095e565b61057b610aca565b60006105bb600154604051806040016040528060018152602001600360fc1b8152506040518060600160405280602a81526020016113c9602a9139610ae4565b90507f993b083090b8404e095fca08253c428df784174f614dc6e1a1ddb5f8b189fcd6600154828587866004546040516105fa969594939291906111cb565b60405180910390a1505060016000555050565b6000806106198361094d565b9050600060068260405161062d919061111e565b908152604080516020928190038301812081830190925290546001600160a01b0381168252600160a01b900460ff1615801592820192909252915061067457519392505050565b5060009392505050565b60408051602080825281830190925260009182919060208201818036833701905050905060008060b56001600160a01b0316836040516106be919061111e565b6000604051808303816000865af19150503d80600081146106fb576040519150601f19603f3d011682016040523d82523d6000602084013e610700565b606091505b5091509150816107525760405162461bcd60e51b815260206004820181905260248201527f6661696c656420746f20636865636b2074686520616363657373206c6576656c604482015260640161055a565b600061075d82610b22565b9050806001149450846107b25760405162461bcd60e51b815260206004820152601760248201527f696e73756666696369656e742070726976696c65676573000000000000000000604482015260640161055a565b6107bc8888610bf7565b6107c4610ddd565b6000610804600154604051806040016040528060018152602001600360fc1b8152506040518060600160405280602a81526020016113c9602a9139610ae4565b90507f637cbd1bf6175910ad1e82f2be446f03227cf561a52e3e7d58792096c50c57de600154828b8b8b60055460405161084396959493929190611212565b60405180910390a1505050505050505050565b6001600160a01b038116600090815260076020526040808220815180830190925280546060939291908290829061088c9061130c565b80601f01602080910402602001604051908101604052809291908181526020018280546108b89061130c565b80156109055780601f106108da57610100808354040283529160200191610905565b820191906000526020600020905b8154815290600101906020018083116108e857829003601f168201915b50505091835250506001919091015460ff16151560209182015281015190915015610931575192915050565b5050604080516020810190915260008152919050565b50919050565b606061095882610df0565b92915050565b6040805160208082528183019092526000916020820181803683370190505090508160005b60208110156109dd5781816020811061099e5761099e61139c565b1a60f81b8382815181106109b4576109b461139c565b60200101906001600160f81b031916908160001a905350806109d581611341565b915050610983565b50600060b76001600160a01b0316836040516109f9919061111e565b6000604051808303816000865af19150503d8060008114610a36576040519150601f19603f3d011682016040523d82523d6000602084013e610a3b565b606091505b5050905080610ac45760405162461bcd60e51b815260206004820152604960248201527f537562636861696e544675656c546f6b656e42616e6b2e5f544675656c566f7560448201527f636865724275726e65643a206661696c656420746f206275726e20544675656c60648201526820766f75636865727360b81b608482015260a40161055a565b50505050565b600160046000828254610add9190611259565b9091555050565b6060610b1a610af285610e6a565b8484604051602001610b069392919061113a565b60405160208183030381529060405261094d565b949350505050565b805160009081906020811115610b885760405162461bcd60e51b815260206004820152602560248201527f627974657320746f2075696e7432353620636f6e76657273696f6e206f766572604482015264666c6f777360d81b606482015260840161055a565b60005b81811015610bee5780610b9f8360206112c9565b610ba99190611259565b610bb49060086112aa565b858281518110610bc657610bc661139c565b01602001516001600160f81b031916901c929092179180610be681611341565b915050610b8b565b50909392505050565b60408051603480825260608201909252600091602082018180368337019050509050816000805b6014811015610c86578560601b8160148110610c3c57610c3c61139c565b1a60f81b8483610c4b81611341565b945081518110610c5d57610c5d61139c565b60200101906001600160f81b031916908160001a90535080610c7e81611341565b915050610c1e565b5060005b6020811015610cef57828160208110610ca557610ca561139c565b1a60f81b8483610cb481611341565b945081518110610cc657610cc661139c565b60200101906001600160f81b031916908160001a90535080610ce781611341565b915050610c8a565b600060b66001600160a01b031685604051610d0a919061111e565b6000604051808303816000865af19150503d8060008114610d47576040519150601f19603f3d011682016040523d82523d6000602084013e610d4c565b606091505b5050905080610dd45760405162461bcd60e51b815260206004820152604860248201527f537562636861696e544675656c546f6b656e42616e6b2e544675656c566f756360448201527f6865724d696e7465643a206661696c656420746f206d696e7420544675656c20606482015267766f75636865727360c01b608482015260a40161055a565b50505050505050565b600160056000828254610add9190611259565b60608160005b8151811015610e6357610e28828281518110610e1457610e1461139c565b01602001516001600160f81b031916610f68565b828281518110610e3a57610e3a61139c565b60200101906001600160f81b031916908160001a90535080610e5b81611341565b915050610df6565b5092915050565b606081610e8e5750506040805180820190915260018152600360fc1b602082015290565b8160005b8115610eb85780610ea281611341565b9150610eb19050600a83611296565b9150610e92565b60008167ffffffffffffffff811115610ed357610ed36113b2565b6040519080825280601f01601f191660200182016040528015610efd576020820181803683370190505b5090505b8415610b1a57610f126001836112c9565b9150610f1f600a8661135c565b610f2a906030611259565b60f81b818381518110610f3f57610f3f61139c565b60200101906001600160f81b031916908160001a905350610f61600a86611296565b9450610f01565b6000604160f81b6001600160f81b0319831610801590610f965750602d60f91b6001600160f81b0319831611155b15610fb357610faa60f883901c6020611271565b60f81b92915050565b5090565b80356001600160a01b0381168114610fce57600080fd5b919050565b600060208284031215610fe557600080fd5b610fee82610fb7565b9392505050565b60008060006060848603121561100a57600080fd5b61101384610fb7565b95602085013595506040909401359392505050565b60006020828403121561103a57600080fd5b813567ffffffffffffffff8082111561105257600080fd5b818401915084601f83011261106657600080fd5b813581811115611078576110786113b2565b604051601f8201601f19908116603f011681019083821181831017156110a0576110a06113b2565b816040528281528760208487010111156110b957600080fd5b826020860160208301376000928101602001929092525095945050505050565b6000602082840312156110eb57600080fd5b5035919050565b6000815180845261110a8160208601602086016112e0565b601f01601f19169290920160200192915050565b600082516111308184602087016112e0565b9190910192915050565b6000845161114c8184602089016112e0565b8083019050602f60f81b808252855161116c816001850160208a016112e0565b600192019182015283516111878160028401602088016112e0565b0160020195945050505050565b602081526000610fee60208301846110f2565b6040815260006111ba60408301856110f2565b905082151560208301529392505050565b86815260c0602082015260006111e460c08301886110f2565b6001600160a01b039687166040840152949095166060820152608081019290925260a0909101529392505050565b86815260c06020820152600061122b60c08301886110f2565b6001600160a01b03969096166040830152506060810193909352608083019190915260a09091015292915050565b6000821982111561126c5761126c611370565b500190565b600060ff821660ff84168060ff0382111561128e5761128e611370565b019392505050565b6000826112a5576112a5611386565b500490565b60008160001904831182151516156112c4576112c4611370565b500290565b6000828210156112db576112db611370565b500390565b60005b838110156112fb5781810151838201526020016112e3565b83811115610ac45750506000910152565b600181811c9082168061132057607f821691505b6020821081141561094757634e487b7160e01b600052602260045260246000fd5b600060001982141561135557611355611370565b5060010190565b60008261136b5761136b611386565b500690565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052601260045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fdfe307830303030303030303030303030303030303030303030303030303030303030303030303030303030a2646970667358221220b239ff1a657586a8aa85f0601498d1f3839c1525944f92c67fff91ea7046588164736f6c63430008070033",
}

// SubchainTFuelTokenBankABI is the input ABI used to generate the binding from.
// Deprecated: Use SubchainTFuelTokenBankMetaData.ABI instead.
var SubchainTFuelTokenBankABI = SubchainTFuelTokenBankMetaData.ABI

// SubchainTFuelTokenBankBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SubchainTFuelTokenBankMetaData.Bin instead.
var SubchainTFuelTokenBankBin = SubchainTFuelTokenBankMetaData.Bin

// DeploySubchainTFuelTokenBank deploys a new Ethereum contract, binding an instance of SubchainTFuelTokenBank to it.
func DeploySubchainTFuelTokenBank(auth *bind.TransactOpts, backend bind.ContractBackend, mainchainid_ *big.Int) (common.Address, *types.Transaction, *SubchainTFuelTokenBank, error) {
	parsed, err := SubchainTFuelTokenBankMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SubchainTFuelTokenBankBin), backend, mainchainid_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SubchainTFuelTokenBank{SubchainTFuelTokenBankCaller: SubchainTFuelTokenBankCaller{contract: contract}, SubchainTFuelTokenBankTransactor: SubchainTFuelTokenBankTransactor{contract: contract}, SubchainTFuelTokenBankFilterer: SubchainTFuelTokenBankFilterer{contract: contract}}, nil
}

// SubchainTFuelTokenBank is an auto generated Go binding around an Ethereum contract.
type SubchainTFuelTokenBank struct {
	SubchainTFuelTokenBankCaller     // Read-only binding to the contract
	SubchainTFuelTokenBankTransactor // Write-only binding to the contract
	SubchainTFuelTokenBankFilterer   // Log filterer for contract events
}

// SubchainTFuelTokenBankCaller is an auto generated read-only Go binding around an Ethereum contract.
type SubchainTFuelTokenBankCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubchainTFuelTokenBankTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SubchainTFuelTokenBankTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubchainTFuelTokenBankFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SubchainTFuelTokenBankFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubchainTFuelTokenBankSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SubchainTFuelTokenBankSession struct {
	Contract     *SubchainTFuelTokenBank // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// SubchainTFuelTokenBankCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SubchainTFuelTokenBankCallerSession struct {
	Contract *SubchainTFuelTokenBankCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// SubchainTFuelTokenBankTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SubchainTFuelTokenBankTransactorSession struct {
	Contract     *SubchainTFuelTokenBankTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// SubchainTFuelTokenBankRaw is an auto generated low-level Go binding around an Ethereum contract.
type SubchainTFuelTokenBankRaw struct {
	Contract *SubchainTFuelTokenBank // Generic contract binding to access the raw methods on
}

// SubchainTFuelTokenBankCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SubchainTFuelTokenBankCallerRaw struct {
	Contract *SubchainTFuelTokenBankCaller // Generic read-only contract binding to access the raw methods on
}

// SubchainTFuelTokenBankTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SubchainTFuelTokenBankTransactorRaw struct {
	Contract *SubchainTFuelTokenBankTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSubchainTFuelTokenBank creates a new instance of SubchainTFuelTokenBank, bound to a specific deployed contract.
func NewSubchainTFuelTokenBank(address common.Address, backend bind.ContractBackend) (*SubchainTFuelTokenBank, error) {
	contract, err := bindSubchainTFuelTokenBank(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SubchainTFuelTokenBank{SubchainTFuelTokenBankCaller: SubchainTFuelTokenBankCaller{contract: contract}, SubchainTFuelTokenBankTransactor: SubchainTFuelTokenBankTransactor{contract: contract}, SubchainTFuelTokenBankFilterer: SubchainTFuelTokenBankFilterer{contract: contract}}, nil
}

// NewSubchainTFuelTokenBankCaller creates a new read-only instance of SubchainTFuelTokenBank, bound to a specific deployed contract.
func NewSubchainTFuelTokenBankCaller(address common.Address, caller bind.ContractCaller) (*SubchainTFuelTokenBankCaller, error) {
	contract, err := bindSubchainTFuelTokenBank(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SubchainTFuelTokenBankCaller{contract: contract}, nil
}

// NewSubchainTFuelTokenBankTransactor creates a new write-only instance of SubchainTFuelTokenBank, bound to a specific deployed contract.
func NewSubchainTFuelTokenBankTransactor(address common.Address, transactor bind.ContractTransactor) (*SubchainTFuelTokenBankTransactor, error) {
	contract, err := bindSubchainTFuelTokenBank(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SubchainTFuelTokenBankTransactor{contract: contract}, nil
}

// NewSubchainTFuelTokenBankFilterer creates a new log filterer instance of SubchainTFuelTokenBank, bound to a specific deployed contract.
func NewSubchainTFuelTokenBankFilterer(address common.Address, filterer bind.ContractFilterer) (*SubchainTFuelTokenBankFilterer, error) {
	contract, err := bindSubchainTFuelTokenBank(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SubchainTFuelTokenBankFilterer{contract: contract}, nil
}

// bindSubchainTFuelTokenBank binds a generic wrapper to an already deployed contract.
func bindSubchainTFuelTokenBank(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SubchainTFuelTokenBankABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SubchainTFuelTokenBank.Contract.SubchainTFuelTokenBankCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SubchainTFuelTokenBank.Contract.SubchainTFuelTokenBankTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SubchainTFuelTokenBank.Contract.SubchainTFuelTokenBankTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SubchainTFuelTokenBank.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SubchainTFuelTokenBank.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SubchainTFuelTokenBank.Contract.contract.Transact(opts, method, params...)
}

// Alldenoms is a free data retrieval call binding the contract method 0xbc15cedc.
//
// Solidity: function alldenoms(uint256 ) view returns(string)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankCaller) Alldenoms(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _SubchainTFuelTokenBank.contract.Call(opts, &out, "alldenoms", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Alldenoms is a free data retrieval call binding the contract method 0xbc15cedc.
//
// Solidity: function alldenoms(uint256 ) view returns(string)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankSession) Alldenoms(arg0 *big.Int) (string, error) {
	return _SubchainTFuelTokenBank.Contract.Alldenoms(&_SubchainTFuelTokenBank.CallOpts, arg0)
}

// Alldenoms is a free data retrieval call binding the contract method 0xbc15cedc.
//
// Solidity: function alldenoms(uint256 ) view returns(string)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankCallerSession) Alldenoms(arg0 *big.Int) (string, error) {
	return _SubchainTFuelTokenBank.Contract.Alldenoms(&_SubchainTFuelTokenBank.CallOpts, arg0)
}

// Allvouchers is a free data retrieval call binding the contract method 0x11418b8e.
//
// Solidity: function allvouchers(uint256 ) view returns(address)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankCaller) Allvouchers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SubchainTFuelTokenBank.contract.Call(opts, &out, "allvouchers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Allvouchers is a free data retrieval call binding the contract method 0x11418b8e.
//
// Solidity: function allvouchers(uint256 ) view returns(address)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankSession) Allvouchers(arg0 *big.Int) (common.Address, error) {
	return _SubchainTFuelTokenBank.Contract.Allvouchers(&_SubchainTFuelTokenBank.CallOpts, arg0)
}

// Allvouchers is a free data retrieval call binding the contract method 0x11418b8e.
//
// Solidity: function allvouchers(uint256 ) view returns(address)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankCallerSession) Allvouchers(arg0 *big.Int) (common.Address, error) {
	return _SubchainTFuelTokenBank.Contract.Allvouchers(&_SubchainTFuelTokenBank.CallOpts, arg0)
}

// Denomtovoucherlookup is a free data retrieval call binding the contract method 0x4c68012b.
//
// Solidity: function denomtovoucherlookup(string ) view returns(address contractaddress, bool exists)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankCaller) Denomtovoucherlookup(opts *bind.CallOpts, arg0 string) (struct {
	Contractaddress common.Address
	Exists          bool
}, error) {
	var out []interface{}
	err := _SubchainTFuelTokenBank.contract.Call(opts, &out, "denomtovoucherlookup", arg0)

	outstruct := new(struct {
		Contractaddress common.Address
		Exists          bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Contractaddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Exists = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// Denomtovoucherlookup is a free data retrieval call binding the contract method 0x4c68012b.
//
// Solidity: function denomtovoucherlookup(string ) view returns(address contractaddress, bool exists)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankSession) Denomtovoucherlookup(arg0 string) (struct {
	Contractaddress common.Address
	Exists          bool
}, error) {
	return _SubchainTFuelTokenBank.Contract.Denomtovoucherlookup(&_SubchainTFuelTokenBank.CallOpts, arg0)
}

// Denomtovoucherlookup is a free data retrieval call binding the contract method 0x4c68012b.
//
// Solidity: function denomtovoucherlookup(string ) view returns(address contractaddress, bool exists)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankCallerSession) Denomtovoucherlookup(arg0 string) (struct {
	Contractaddress common.Address
	Exists          bool
}, error) {
	return _SubchainTFuelTokenBank.Contract.Denomtovoucherlookup(&_SubchainTFuelTokenBank.CallOpts, arg0)
}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(string denom) view returns(bool)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankCaller) Exists(opts *bind.CallOpts, denom string) (bool, error) {
	var out []interface{}
	err := _SubchainTFuelTokenBank.contract.Call(opts, &out, "exists", denom)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(string denom) view returns(bool)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankSession) Exists(denom string) (bool, error) {
	return _SubchainTFuelTokenBank.Contract.Exists(&_SubchainTFuelTokenBank.CallOpts, denom)
}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(string denom) view returns(bool)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankCallerSession) Exists(denom string) (bool, error) {
	return _SubchainTFuelTokenBank.Contract.Exists(&_SubchainTFuelTokenBank.CallOpts, denom)
}

// Exists0 is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address voucheraddress) view returns(bool)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankCaller) Exists0(opts *bind.CallOpts, voucheraddress common.Address) (bool, error) {
	var out []interface{}
	err := _SubchainTFuelTokenBank.contract.Call(opts, &out, "exists0", voucheraddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists0 is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address voucheraddress) view returns(bool)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankSession) Exists0(voucheraddress common.Address) (bool, error) {
	return _SubchainTFuelTokenBank.Contract.Exists0(&_SubchainTFuelTokenBank.CallOpts, voucheraddress)
}

// Exists0 is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address voucheraddress) view returns(bool)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankCallerSession) Exists0(voucheraddress common.Address) (bool, error) {
	return _SubchainTFuelTokenBank.Contract.Exists0(&_SubchainTFuelTokenBank.CallOpts, voucheraddress)
}

// Getdenom is a free data retrieval call binding the contract method 0xaf46078f.
//
// Solidity: function getdenom(address vouchercontractaddr) view returns(string)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankCaller) Getdenom(opts *bind.CallOpts, vouchercontractaddr common.Address) (string, error) {
	var out []interface{}
	err := _SubchainTFuelTokenBank.contract.Call(opts, &out, "getdenom", vouchercontractaddr)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Getdenom is a free data retrieval call binding the contract method 0xaf46078f.
//
// Solidity: function getdenom(address vouchercontractaddr) view returns(string)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankSession) Getdenom(vouchercontractaddr common.Address) (string, error) {
	return _SubchainTFuelTokenBank.Contract.Getdenom(&_SubchainTFuelTokenBank.CallOpts, vouchercontractaddr)
}

// Getdenom is a free data retrieval call binding the contract method 0xaf46078f.
//
// Solidity: function getdenom(address vouchercontractaddr) view returns(string)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankCallerSession) Getdenom(vouchercontractaddr common.Address) (string, error) {
	return _SubchainTFuelTokenBank.Contract.Getdenom(&_SubchainTFuelTokenBank.CallOpts, vouchercontractaddr)
}

// Getvoucher is a free data retrieval call binding the contract method 0xd2990e7d.
//
// Solidity: function getvoucher(string denom) view returns(address)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankCaller) Getvoucher(opts *bind.CallOpts, denom string) (common.Address, error) {
	var out []interface{}
	err := _SubchainTFuelTokenBank.contract.Call(opts, &out, "getvoucher", denom)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Getvoucher is a free data retrieval call binding the contract method 0xd2990e7d.
//
// Solidity: function getvoucher(string denom) view returns(address)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankSession) Getvoucher(denom string) (common.Address, error) {
	return _SubchainTFuelTokenBank.Contract.Getvoucher(&_SubchainTFuelTokenBank.CallOpts, denom)
}

// Getvoucher is a free data retrieval call binding the contract method 0xd2990e7d.
//
// Solidity: function getvoucher(string denom) view returns(address)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankCallerSession) Getvoucher(denom string) (common.Address, error) {
	return _SubchainTFuelTokenBank.Contract.Getvoucher(&_SubchainTFuelTokenBank.CallOpts, denom)
}

// Mainchainid is a free data retrieval call binding the contract method 0xba3599b3.
//
// Solidity: function mainchainid() view returns(uint256)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankCaller) Mainchainid(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SubchainTFuelTokenBank.contract.Call(opts, &out, "mainchainid")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Mainchainid is a free data retrieval call binding the contract method 0xba3599b3.
//
// Solidity: function mainchainid() view returns(uint256)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankSession) Mainchainid() (*big.Int, error) {
	return _SubchainTFuelTokenBank.Contract.Mainchainid(&_SubchainTFuelTokenBank.CallOpts)
}

// Mainchainid is a free data retrieval call binding the contract method 0xba3599b3.
//
// Solidity: function mainchainid() view returns(uint256)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankCallerSession) Mainchainid() (*big.Int, error) {
	return _SubchainTFuelTokenBank.Contract.Mainchainid(&_SubchainTFuelTokenBank.CallOpts)
}

// Tokenlocknonce is a free data retrieval call binding the contract method 0x03ffa7b2.
//
// Solidity: function tokenlocknonce() view returns(uint256)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankCaller) Tokenlocknonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SubchainTFuelTokenBank.contract.Call(opts, &out, "tokenlocknonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Tokenlocknonce is a free data retrieval call binding the contract method 0x03ffa7b2.
//
// Solidity: function tokenlocknonce() view returns(uint256)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankSession) Tokenlocknonce() (*big.Int, error) {
	return _SubchainTFuelTokenBank.Contract.Tokenlocknonce(&_SubchainTFuelTokenBank.CallOpts)
}

// Tokenlocknonce is a free data retrieval call binding the contract method 0x03ffa7b2.
//
// Solidity: function tokenlocknonce() view returns(uint256)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankCallerSession) Tokenlocknonce() (*big.Int, error) {
	return _SubchainTFuelTokenBank.Contract.Tokenlocknonce(&_SubchainTFuelTokenBank.CallOpts)
}

// Tokenunlocknonce is a free data retrieval call binding the contract method 0xf89b5efd.
//
// Solidity: function tokenunlocknonce() view returns(uint256)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankCaller) Tokenunlocknonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SubchainTFuelTokenBank.contract.Call(opts, &out, "tokenunlocknonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Tokenunlocknonce is a free data retrieval call binding the contract method 0xf89b5efd.
//
// Solidity: function tokenunlocknonce() view returns(uint256)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankSession) Tokenunlocknonce() (*big.Int, error) {
	return _SubchainTFuelTokenBank.Contract.Tokenunlocknonce(&_SubchainTFuelTokenBank.CallOpts)
}

// Tokenunlocknonce is a free data retrieval call binding the contract method 0xf89b5efd.
//
// Solidity: function tokenunlocknonce() view returns(uint256)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankCallerSession) Tokenunlocknonce() (*big.Int, error) {
	return _SubchainTFuelTokenBank.Contract.Tokenunlocknonce(&_SubchainTFuelTokenBank.CallOpts)
}

// Voucheraddresstodenomlookup is a free data retrieval call binding the contract method 0xcd22e450.
//
// Solidity: function voucheraddresstodenomlookup(address ) view returns(string denom, bool exists)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankCaller) Voucheraddresstodenomlookup(opts *bind.CallOpts, arg0 common.Address) (struct {
	Denom  string
	Exists bool
}, error) {
	var out []interface{}
	err := _SubchainTFuelTokenBank.contract.Call(opts, &out, "voucheraddresstodenomlookup", arg0)

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

// Voucheraddresstodenomlookup is a free data retrieval call binding the contract method 0xcd22e450.
//
// Solidity: function voucheraddresstodenomlookup(address ) view returns(string denom, bool exists)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankSession) Voucheraddresstodenomlookup(arg0 common.Address) (struct {
	Denom  string
	Exists bool
}, error) {
	return _SubchainTFuelTokenBank.Contract.Voucheraddresstodenomlookup(&_SubchainTFuelTokenBank.CallOpts, arg0)
}

// Voucheraddresstodenomlookup is a free data retrieval call binding the contract method 0xcd22e450.
//
// Solidity: function voucheraddresstodenomlookup(address ) view returns(string denom, bool exists)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankCallerSession) Voucheraddresstodenomlookup(arg0 common.Address) (struct {
	Denom  string
	Exists bool
}, error) {
	return _SubchainTFuelTokenBank.Contract.Voucheraddresstodenomlookup(&_SubchainTFuelTokenBank.CallOpts, arg0)
}

// Voucherburnnonce is a free data retrieval call binding the contract method 0x11bc48fb.
//
// Solidity: function voucherburnnonce() view returns(uint256)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankCaller) Voucherburnnonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SubchainTFuelTokenBank.contract.Call(opts, &out, "voucherburnnonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Voucherburnnonce is a free data retrieval call binding the contract method 0x11bc48fb.
//
// Solidity: function voucherburnnonce() view returns(uint256)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankSession) Voucherburnnonce() (*big.Int, error) {
	return _SubchainTFuelTokenBank.Contract.Voucherburnnonce(&_SubchainTFuelTokenBank.CallOpts)
}

// Voucherburnnonce is a free data retrieval call binding the contract method 0x11bc48fb.
//
// Solidity: function voucherburnnonce() view returns(uint256)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankCallerSession) Voucherburnnonce() (*big.Int, error) {
	return _SubchainTFuelTokenBank.Contract.Voucherburnnonce(&_SubchainTFuelTokenBank.CallOpts)
}

// Vouchermintnonce is a free data retrieval call binding the contract method 0xb516d30f.
//
// Solidity: function vouchermintnonce() view returns(uint256)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankCaller) Vouchermintnonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SubchainTFuelTokenBank.contract.Call(opts, &out, "vouchermintnonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Vouchermintnonce is a free data retrieval call binding the contract method 0xb516d30f.
//
// Solidity: function vouchermintnonce() view returns(uint256)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankSession) Vouchermintnonce() (*big.Int, error) {
	return _SubchainTFuelTokenBank.Contract.Vouchermintnonce(&_SubchainTFuelTokenBank.CallOpts)
}

// Vouchermintnonce is a free data retrieval call binding the contract method 0xb516d30f.
//
// Solidity: function vouchermintnonce() view returns(uint256)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankCallerSession) Vouchermintnonce() (*big.Int, error) {
	return _SubchainTFuelTokenBank.Contract.Vouchermintnonce(&_SubchainTFuelTokenBank.CallOpts)
}

// Burnvouchers is a paid mutator transaction binding the contract method 0xc18b2085.
//
// Solidity: function burnvouchers(address mainchaintokenreceiver) payable returns()
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankTransactor) Burnvouchers(opts *bind.TransactOpts, mainchaintokenreceiver common.Address) (*types.Transaction, error) {
	return _SubchainTFuelTokenBank.contract.Transact(opts, "burnvouchers", mainchaintokenreceiver)
}

// Burnvouchers is a paid mutator transaction binding the contract method 0xc18b2085.
//
// Solidity: function burnvouchers(address mainchaintokenreceiver) payable returns()
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankSession) Burnvouchers(mainchaintokenreceiver common.Address) (*types.Transaction, error) {
	return _SubchainTFuelTokenBank.Contract.Burnvouchers(&_SubchainTFuelTokenBank.TransactOpts, mainchaintokenreceiver)
}

// Burnvouchers is a paid mutator transaction binding the contract method 0xc18b2085.
//
// Solidity: function burnvouchers(address mainchaintokenreceiver) payable returns()
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankTransactorSession) Burnvouchers(mainchaintokenreceiver common.Address) (*types.Transaction, error) {
	return _SubchainTFuelTokenBank.Contract.Burnvouchers(&_SubchainTFuelTokenBank.TransactOpts, mainchaintokenreceiver)
}

// Mintvouchers is a paid mutator transaction binding the contract method 0x0e44406d.
//
// Solidity: function mintvouchers(address voucherreceiver, uint256 mintamount, uint256 mainchaintokenlocknonce) returns()
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankTransactor) Mintvouchers(opts *bind.TransactOpts, voucherreceiver common.Address, mintamount *big.Int, mainchaintokenlocknonce *big.Int) (*types.Transaction, error) {
	return _SubchainTFuelTokenBank.contract.Transact(opts, "mintvouchers", voucherreceiver, mintamount, mainchaintokenlocknonce)
}

// Mintvouchers is a paid mutator transaction binding the contract method 0x0e44406d.
//
// Solidity: function mintvouchers(address voucherreceiver, uint256 mintamount, uint256 mainchaintokenlocknonce) returns()
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankSession) Mintvouchers(voucherreceiver common.Address, mintamount *big.Int, mainchaintokenlocknonce *big.Int) (*types.Transaction, error) {
	return _SubchainTFuelTokenBank.Contract.Mintvouchers(&_SubchainTFuelTokenBank.TransactOpts, voucherreceiver, mintamount, mainchaintokenlocknonce)
}

// Mintvouchers is a paid mutator transaction binding the contract method 0x0e44406d.
//
// Solidity: function mintvouchers(address voucherreceiver, uint256 mintamount, uint256 mainchaintokenlocknonce) returns()
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankTransactorSession) Mintvouchers(voucherreceiver common.Address, mintamount *big.Int, mainchaintokenlocknonce *big.Int) (*types.Transaction, error) {
	return _SubchainTFuelTokenBank.Contract.Mintvouchers(&_SubchainTFuelTokenBank.TransactOpts, voucherreceiver, mintamount, mainchaintokenlocknonce)
}

// SubchainTFuelTokenBankTfuelvoucherburnedIterator is returned from FilterTfuelvoucherburned and is used to iterate over the raw logs and unpacked data for Tfuelvoucherburned events raised by the SubchainTFuelTokenBank contract.
type SubchainTFuelTokenBankTfuelvoucherburnedIterator struct {
	Event *SubchainTFuelTokenBankTfuelvoucherburned // Event containing the contract specifics and raw log

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
func (it *SubchainTFuelTokenBankTfuelvoucherburnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubchainTFuelTokenBankTfuelvoucherburned)
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
		it.Event = new(SubchainTFuelTokenBankTfuelvoucherburned)
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
func (it *SubchainTFuelTokenBankTfuelvoucherburnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubchainTFuelTokenBankTfuelvoucherburnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubchainTFuelTokenBankTfuelvoucherburned represents a Tfuelvoucherburned event raised by the SubchainTFuelTokenBank contract.
type SubchainTFuelTokenBankTfuelvoucherburned struct {
	Mainchainid            *big.Int
	Denom                  string
	Voucherowner           common.Address
	Mainchaintokenreceiver common.Address
	Amount                 *big.Int
	Voucherburnnonce       *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterTfuelvoucherburned is a free log retrieval operation binding the contract event 0xa7cfa6d5d6a7158eb281bcae45cf205fdb1fbf491ddf68b86f663ed94a7e9181.
//
// Solidity: event tfuelvoucherburned(uint256 mainchainid, string denom, address voucherowner, address mainchaintokenreceiver, uint256 amount, uint256 voucherburnnonce)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankFilterer) FilterTfuelvoucherburned(opts *bind.FilterOpts) (*SubchainTFuelTokenBankTfuelvoucherburnedIterator, error) {

	logs, sub, err := _SubchainTFuelTokenBank.contract.FilterLogs(opts, "tfuelvoucherburned")
	if err != nil {
		return nil, err
	}
	return &SubchainTFuelTokenBankTfuelvoucherburnedIterator{contract: _SubchainTFuelTokenBank.contract, event: "tfuelvoucherburned", logs: logs, sub: sub}, nil
}

// WatchTfuelvoucherburned is a free log subscription operation binding the contract event 0xa7cfa6d5d6a7158eb281bcae45cf205fdb1fbf491ddf68b86f663ed94a7e9181.
//
// Solidity: event tfuelvoucherburned(uint256 mainchainid, string denom, address voucherowner, address mainchaintokenreceiver, uint256 amount, uint256 voucherburnnonce)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankFilterer) WatchTfuelvoucherburned(opts *bind.WatchOpts, sink chan<- *SubchainTFuelTokenBankTfuelvoucherburned) (event.Subscription, error) {

	logs, sub, err := _SubchainTFuelTokenBank.contract.WatchLogs(opts, "tfuelvoucherburned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubchainTFuelTokenBankTfuelvoucherburned)
				if err := _SubchainTFuelTokenBank.contract.UnpackLog(event, "tfuelvoucherburned", log); err != nil {
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

// ParseTfuelvoucherburned is a log parse operation binding the contract event 0xa7cfa6d5d6a7158eb281bcae45cf205fdb1fbf491ddf68b86f663ed94a7e9181.
//
// Solidity: event tfuelvoucherburned(uint256 mainchainid, string denom, address voucherowner, address mainchaintokenreceiver, uint256 amount, uint256 voucherburnnonce)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankFilterer) ParseTfuelvoucherburned(log types.Log) (*SubchainTFuelTokenBankTfuelvoucherburned, error) {
	event := new(SubchainTFuelTokenBankTfuelvoucherburned)
	if err := _SubchainTFuelTokenBank.contract.UnpackLog(event, "tfuelvoucherburned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubchainTFuelTokenBankTfuelvouchermintedIterator is returned from FilterTfuelvoucherminted and is used to iterate over the raw logs and unpacked data for Tfuelvoucherminted events raised by the SubchainTFuelTokenBank contract.
type SubchainTFuelTokenBankTfuelvouchermintedIterator struct {
	Event *SubchainTFuelTokenBankTfuelvoucherminted // Event containing the contract specifics and raw log

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
func (it *SubchainTFuelTokenBankTfuelvouchermintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubchainTFuelTokenBankTfuelvoucherminted)
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
		it.Event = new(SubchainTFuelTokenBankTfuelvoucherminted)
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
func (it *SubchainTFuelTokenBankTfuelvouchermintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubchainTFuelTokenBankTfuelvouchermintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubchainTFuelTokenBankTfuelvoucherminted represents a Tfuelvoucherminted event raised by the SubchainTFuelTokenBank contract.
type SubchainTFuelTokenBankTfuelvoucherminted struct {
	Mainchainid             *big.Int
	Denom                   string
	Voucherreceiver         common.Address
	Amount                  *big.Int
	Mainchaintokenlocknonce *big.Int
	Vouchermintnonce        *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterTfuelvoucherminted is a free log retrieval operation binding the contract event 0x71009aa2fbc71446ff84aeaa2ea78c58edfeaa64c6828d7557178906ce2ca30d.
//
// Solidity: event tfuelvoucherminted(uint256 mainchainid, string denom, address voucherreceiver, uint256 amount, uint256 mainchaintokenlocknonce, uint256 vouchermintnonce)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankFilterer) FilterTfuelvoucherminted(opts *bind.FilterOpts) (*SubchainTFuelTokenBankTfuelvouchermintedIterator, error) {

	logs, sub, err := _SubchainTFuelTokenBank.contract.FilterLogs(opts, "tfuelvoucherminted")
	if err != nil {
		return nil, err
	}
	return &SubchainTFuelTokenBankTfuelvouchermintedIterator{contract: _SubchainTFuelTokenBank.contract, event: "tfuelvoucherminted", logs: logs, sub: sub}, nil
}

// WatchTfuelvoucherminted is a free log subscription operation binding the contract event 0x71009aa2fbc71446ff84aeaa2ea78c58edfeaa64c6828d7557178906ce2ca30d.
//
// Solidity: event tfuelvoucherminted(uint256 mainchainid, string denom, address voucherreceiver, uint256 amount, uint256 mainchaintokenlocknonce, uint256 vouchermintnonce)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankFilterer) WatchTfuelvoucherminted(opts *bind.WatchOpts, sink chan<- *SubchainTFuelTokenBankTfuelvoucherminted) (event.Subscription, error) {

	logs, sub, err := _SubchainTFuelTokenBank.contract.WatchLogs(opts, "tfuelvoucherminted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubchainTFuelTokenBankTfuelvoucherminted)
				if err := _SubchainTFuelTokenBank.contract.UnpackLog(event, "tfuelvoucherminted", log); err != nil {
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

// ParseTfuelvoucherminted is a log parse operation binding the contract event 0x71009aa2fbc71446ff84aeaa2ea78c58edfeaa64c6828d7557178906ce2ca30d.
//
// Solidity: event tfuelvoucherminted(uint256 mainchainid, string denom, address voucherreceiver, uint256 amount, uint256 mainchaintokenlocknonce, uint256 vouchermintnonce)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankFilterer) ParseTfuelvoucherminted(log types.Log) (*SubchainTFuelTokenBankTfuelvoucherminted, error) {
	event := new(SubchainTFuelTokenBankTfuelvoucherminted)
	if err := _SubchainTFuelTokenBank.contract.UnpackLog(event, "tfuelvoucherminted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
