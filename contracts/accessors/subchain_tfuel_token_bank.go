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
	ABI: "[{\"inputs\":[],\"statemutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internaltype\":\"address\",\"name\":\"voucherowner\",\"type\":\"address\"},{\"indexed\":false,\"internaltype\":\"address\",\"name\":\"mainchaintokenreceiver\",\"type\":\"address\"},{\"indexed\":false,\"internaltype\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internaltype\":\"uint256\",\"name\":\"voucherburnnonce\",\"type\":\"uint256\"}],\"name\":\"burntfuelvouchers\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internaltype\":\"address\",\"name\":\"voucherreceiver\",\"type\":\"address\"},{\"indexed\":false,\"internaltype\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internaltype\":\"uint256\",\"name\":\"mainchaintokenlocknonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internaltype\":\"uint256\",\"name\":\"vouchermintnonce\",\"type\":\"uint256\"}],\"name\":\"minttfuelvouchers\",\"type\":\"event\"},{\"inputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"alldenoms\",\"outputs\":[{\"internaltype\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allvouchers\",\"outputs\":[{\"internaltype\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"denomtovoucherlookup\",\"outputs\":[{\"internaltype\":\"address\",\"name\":\"contractaddress\",\"type\":\"address\"},{\"internaltype\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"name\":\"exists\",\"outputs\":[{\"internaltype\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"address\",\"name\":\"voucheraddress\",\"type\":\"address\"}],\"name\":\"exists\",\"outputs\":[{\"internaltype\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"address\",\"name\":\"vouchercontractaddr\",\"type\":\"address\"}],\"name\":\"getdenom\",\"outputs\":[{\"internaltype\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"name\":\"getvoucher\",\"outputs\":[{\"internaltype\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenlocknonce\",\"outputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenunlocknonce\",\"outputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"voucheraddresstodenomlookup\",\"outputs\":[{\"internaltype\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internaltype\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voucherburnnonce\",\"outputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vouchermintnonce\",\"outputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"address\",\"name\":\"voucherreceiver\",\"type\":\"address\"},{\"internaltype\":\"uint256\",\"name\":\"mintamount\",\"type\":\"uint256\"},{\"internaltype\":\"uint256\",\"name\":\"mainchaintokenlocknonce\",\"type\":\"uint256\"}],\"name\":\"mintvouchers\",\"outputs\":[],\"statemutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"address\",\"name\":\"mainchaintokenreceiver\",\"type\":\"address\"}],\"name\":\"burnvouchers\",\"outputs\":[],\"statemutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50600160008181559081905560028190556003819055600455611110806100386000396000f3fe6080604052600436106100dd5760003560e01c806360569b5e1161007f578063a4ddaaa611610059578063a4ddaaa61461029b578063c27d927a146102bb578063ebda9962146102d1578063f6a3d24e146102f157600080fd5b806360569b5e146102385780637d0fb00d14610266578063a2cc69811461027b57600080fd5b806327ca4df1116100bb57806327ca4df1146101a757806333835be6146101df5780635102e760146101f5578063588b14081461020b57600080fd5b80631527b14d146100e2578063163fe97214610153578063261a323e14610177575b600080fd5b3480156100ee57600080fd5b5061012f6100fd366004610e60565b80516020818301810180516005825292820191909301209152546001600160a01b03811690600160a01b900460ff1682565b604080516001600160a01b0390931683529015156020830152015b60405180910390f35b34801561015f57600080fd5b5061016960045481565b60405190815260200161014a565b34801561018357600080fd5b50610197610192366004610e60565b61032d565b604051901515815260200161014a565b3480156101b357600080fd5b506101c76101c2366004610f11565b61036d565b6040516001600160a01b03909116815260200161014a565b3480156101eb57600080fd5b5061016960025481565b34801561020157600080fd5b5061016960015481565b34801561021757600080fd5b5061022b610226366004610f11565b610397565b60405161014a9190610f72565b34801561024457600080fd5b50610258610253366004610e0b565b610443565b60405161014a929190610f85565b610279610274366004610e0b565b6104ea565b005b34801561028757600080fd5b506101c7610296366004610e60565b6105b8565b3480156102a757600080fd5b506102796102b6366004610e2d565b610629565b3480156102c757600080fd5b5061016960035481565b3480156102dd57600080fd5b5061022b6102ec366004610e0b565b6107cc565b3480156102fd57600080fd5b5061019761030c366004610e0b565b6001600160a01b031660009081526006602052604090206001015460ff1690565b600080610339836108c3565b905060058160405161034b9190610f56565b9081526040519081900360200190205460ff600160a01b909104169392505050565b6007818154811061037d57600080fd5b6000918252602090912001546001600160a01b0316905081565b600881815481106103a757600080fd5b9060005260206000200160009150905080546103c290611048565b80601f01602080910402602001604051908101604052809291908181526020018280546103ee90611048565b801561043b5780601f106104105761010080835404028352916020019161043b565b820191906000526020600020905b81548152906001019060200180831161041e57829003601f168201915b505050505081565b60066020526000908152604090208054819061045e90611048565b80601f016020809104026020016040519081016040528092919081815260200182805461048a90611048565b80156104d75780601f106104ac576101008083540402835291602001916104d7565b820191906000526020600020905b8154815290600101906020018083116104ba57829003601f168201915b5050506001909301549192505060ff1682565b600260005414156105425760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064015b60405180910390fd5b60026000553334610552816108d4565b61055a610a3f565b600354604080516001600160a01b038581168252861660208201528082018490526060810192909252517f2d8470f382989926e91872b25e27240426cc2efe1bfef95339ae1944a221708f9181900360800190a15050600160005550565b6000806105c4836108c3565b905060006005826040516105d89190610f56565b908152604080516020928190038301812081830190925290546001600160a01b0381168252600160a01b900460ff1615801592820192909252915061061f57519392505050565b5060009392505050565b60408051602080825281830190925260009182919060208201818036833701905050905060008060b56001600160a01b0316836040516106699190610f56565b6000604051808303816000865af19150503d80600081146106a6576040519150601f19603f3d011682016040523d82523d6000602084013e6106ab565b606091505b5091509150816106fd5760405162461bcd60e51b815260206004820181905260248201527f6661696c656420746f20636865636b2074686520616363657373206c6576656c6044820152606401610539565b600061070882610a59565b90508060011494508461075d5760405162461bcd60e51b815260206004820152601760248201527f696e73756666696369656e742070726976696c656765730000000000000000006044820152606401610539565b6107678888610b2e565b61076f610d13565b600454604080516001600160a01b038b168152602081018a90528082018990526060810192909252517f96d5e853bcbdb7781e569193053bec1adfa85361c6cfb27abdfd6dd100b0907d9181900360800190a15050505050505050565b6001600160a01b038116600090815260066020526040808220815180830190925280546060939291908290829061080290611048565b80601f016020809104026020016040519081016040528092919081815260200182805461082e90611048565b801561087b5780601f106108505761010080835404028352916020019161087b565b820191906000526020600020905b81548152906001019060200180831161085e57829003601f168201915b50505091835250506001919091015460ff161515602091820152810151909150156108a7575192915050565b5050604080516020810190915260008152919050565b50919050565b60606108ce82610d26565b92915050565b6040805160208082528183019092526000916020820181803683370190505090508160005b602081101561095357818160208110610914576109146110ae565b1a60f81b83828151811061092a5761092a6110ae565b60200101906001600160f81b031916908160001a9053508061094b8161107d565b9150506108f9565b50600060b76001600160a01b03168360405161096f9190610f56565b6000604051808303816000865af19150503d80600081146109ac576040519150601f19603f3d011682016040523d82523d6000602084013e6109b1565b606091505b5050905080610a395760405162461bcd60e51b815260206004820152604860248201527f537562636861696e544675656c546f6b656e42616e6b2e5f6275726e5446756560448201527f6c566f7563686572733a206661696c656420746f206275726e20544675656c20606482015267766f75636865727360c01b608482015260a401610539565b50505050565b600160036000828254610a529190610fa9565b9091555050565b805160009081906020811115610abf5760405162461bcd60e51b815260206004820152602560248201527f627974657320746f2075696e7432353620636f6e76657273696f6e206f766572604482015264666c6f777360d81b6064820152608401610539565b60005b81811015610b255780610ad6836020611005565b610ae09190610fa9565b610aeb906008610fe6565b858281518110610afd57610afd6110ae565b01602001516001600160f81b031916901c929092179180610b1d8161107d565b915050610ac2565b50909392505050565b60408051603480825260608201909252600091602082018180368337019050509050816000805b6014811015610bbd578560601b8160148110610b7357610b736110ae565b1a60f81b8483610b828161107d565b945081518110610b9457610b946110ae565b60200101906001600160f81b031916908160001a90535080610bb58161107d565b915050610b55565b5060005b6020811015610c2657828160208110610bdc57610bdc6110ae565b1a60f81b8483610beb8161107d565b945081518110610bfd57610bfd6110ae565b60200101906001600160f81b031916908160001a90535080610c1e8161107d565b915050610bc1565b600060b66001600160a01b031685604051610c419190610f56565b6000604051808303816000865af19150503d8060008114610c7e576040519150601f19603f3d011682016040523d82523d6000602084013e610c83565b606091505b5050905080610d0a5760405162461bcd60e51b815260206004820152604760248201527f537562636861696e544675656c546f6b656e42616e6b2e6d696e74544675656c60448201527f566f7563686572733a206661696c656420746f206d696e7420544675656c20766064820152666f75636865727360c81b608482015260a401610539565b50505050505050565b600160046000828254610a529190610fa9565b60608160005b8151811015610d9957610d5e828281518110610d4a57610d4a6110ae565b01602001516001600160f81b031916610da0565b828281518110610d7057610d706110ae565b60200101906001600160f81b031916908160001a90535080610d918161107d565b915050610d2c565b5092915050565b6000604160f81b6001600160f81b0319831610801590610dce5750602d60f91b6001600160f81b0319831611155b15610deb57610de260f883901c6020610fc1565b60f81b92915050565b5090565b80356001600160a01b0381168114610e0657600080fd5b919050565b600060208284031215610e1d57600080fd5b610e2682610def565b9392505050565b600080600060608486031215610e4257600080fd5b610e4b84610def565b95602085013595506040909401359392505050565b600060208284031215610e7257600080fd5b813567ffffffffffffffff80821115610e8a57600080fd5b818401915084601f830112610e9e57600080fd5b813581811115610eb057610eb06110c4565b604051601f8201601f19908116603f01168101908382118183101715610ed857610ed86110c4565b81604052828152876020848701011115610ef157600080fd5b826020860160208301376000928101602001929092525095945050505050565b600060208284031215610f2357600080fd5b5035919050565b60008151808452610f4281602086016020860161101c565b601f01601f19169290920160200192915050565b60008251610f6881846020870161101c565b9190910192915050565b602081526000610e266020830184610f2a565b604081526000610f986040830185610f2a565b905082151560208301529392505050565b60008219821115610fbc57610fbc611098565b500190565b600060ff821660ff84168060ff03821115610fde57610fde611098565b019392505050565b600081600019048311821515161561100057611000611098565b500290565b60008282101561101757611017611098565b500390565b60005b8381101561103757818101518382015260200161101f565b83811115610a395750506000910152565b600181811c9082168061105c57607f821691505b602082108114156108bd57634e487b7160e01b600052602260045260246000fd5b600060001982141561109157611091611098565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fdfea26469706673582212208edf96ffb79590d7177e046cf60ed3bf0866ed9eaf1fc202e91f133019d4b5a764736f6c63430008070033",
}

// SubchainTFuelTokenBankABI is the input ABI used to generate the binding from.
// Deprecated: Use SubchainTFuelTokenBankMetaData.ABI instead.
var SubchainTFuelTokenBankABI = SubchainTFuelTokenBankMetaData.ABI

// SubchainTFuelTokenBankBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SubchainTFuelTokenBankMetaData.Bin instead.
var SubchainTFuelTokenBankBin = SubchainTFuelTokenBankMetaData.Bin

// DeploySubchainTFuelTokenBank deploys a new Ethereum contract, binding an instance of SubchainTFuelTokenBank to it.
func DeploySubchainTFuelTokenBank(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SubchainTFuelTokenBank, error) {
	parsed, err := SubchainTFuelTokenBankMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SubchainTFuelTokenBankBin), backend)
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

// SubchainTFuelTokenBankBurntfuelvouchersIterator is returned from FilterBurntfuelvouchers and is used to iterate over the raw logs and unpacked data for Burntfuelvouchers events raised by the SubchainTFuelTokenBank contract.
type SubchainTFuelTokenBankBurntfuelvouchersIterator struct {
	Event *SubchainTFuelTokenBankBurntfuelvouchers // Event containing the contract specifics and raw log

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
func (it *SubchainTFuelTokenBankBurntfuelvouchersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubchainTFuelTokenBankBurntfuelvouchers)
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
		it.Event = new(SubchainTFuelTokenBankBurntfuelvouchers)
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
func (it *SubchainTFuelTokenBankBurntfuelvouchersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubchainTFuelTokenBankBurntfuelvouchersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubchainTFuelTokenBankBurntfuelvouchers represents a Burntfuelvouchers event raised by the SubchainTFuelTokenBank contract.
type SubchainTFuelTokenBankBurntfuelvouchers struct {
	Voucherowner           common.Address
	Mainchaintokenreceiver common.Address
	Amount                 *big.Int
	Voucherburnnonce       *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterBurntfuelvouchers is a free log retrieval operation binding the contract event 0xe557265d897cf38e6c2a28a3a53b0a186397ceeb4308813ab3636bb3b2c1217c.
//
// Solidity: event burntfuelvouchers(address voucherowner, address mainchaintokenreceiver, uint256 amount, uint256 voucherburnnonce)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankFilterer) FilterBurntfuelvouchers(opts *bind.FilterOpts) (*SubchainTFuelTokenBankBurntfuelvouchersIterator, error) {

	logs, sub, err := _SubchainTFuelTokenBank.contract.FilterLogs(opts, "burntfuelvouchers")
	if err != nil {
		return nil, err
	}
	return &SubchainTFuelTokenBankBurntfuelvouchersIterator{contract: _SubchainTFuelTokenBank.contract, event: "burntfuelvouchers", logs: logs, sub: sub}, nil
}

// WatchBurntfuelvouchers is a free log subscription operation binding the contract event 0xe557265d897cf38e6c2a28a3a53b0a186397ceeb4308813ab3636bb3b2c1217c.
//
// Solidity: event burntfuelvouchers(address voucherowner, address mainchaintokenreceiver, uint256 amount, uint256 voucherburnnonce)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankFilterer) WatchBurntfuelvouchers(opts *bind.WatchOpts, sink chan<- *SubchainTFuelTokenBankBurntfuelvouchers) (event.Subscription, error) {

	logs, sub, err := _SubchainTFuelTokenBank.contract.WatchLogs(opts, "burntfuelvouchers")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubchainTFuelTokenBankBurntfuelvouchers)
				if err := _SubchainTFuelTokenBank.contract.UnpackLog(event, "burntfuelvouchers", log); err != nil {
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

// ParseBurntfuelvouchers is a log parse operation binding the contract event 0xe557265d897cf38e6c2a28a3a53b0a186397ceeb4308813ab3636bb3b2c1217c.
//
// Solidity: event burntfuelvouchers(address voucherowner, address mainchaintokenreceiver, uint256 amount, uint256 voucherburnnonce)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankFilterer) ParseBurntfuelvouchers(log types.Log) (*SubchainTFuelTokenBankBurntfuelvouchers, error) {
	event := new(SubchainTFuelTokenBankBurntfuelvouchers)
	if err := _SubchainTFuelTokenBank.contract.UnpackLog(event, "burntfuelvouchers", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubchainTFuelTokenBankMinttfuelvouchersIterator is returned from FilterMinttfuelvouchers and is used to iterate over the raw logs and unpacked data for Minttfuelvouchers events raised by the SubchainTFuelTokenBank contract.
type SubchainTFuelTokenBankMinttfuelvouchersIterator struct {
	Event *SubchainTFuelTokenBankMinttfuelvouchers // Event containing the contract specifics and raw log

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
func (it *SubchainTFuelTokenBankMinttfuelvouchersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubchainTFuelTokenBankMinttfuelvouchers)
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
		it.Event = new(SubchainTFuelTokenBankMinttfuelvouchers)
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
func (it *SubchainTFuelTokenBankMinttfuelvouchersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubchainTFuelTokenBankMinttfuelvouchersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubchainTFuelTokenBankMinttfuelvouchers represents a Minttfuelvouchers event raised by the SubchainTFuelTokenBank contract.
type SubchainTFuelTokenBankMinttfuelvouchers struct {
	Voucherreceiver         common.Address
	Amount                  *big.Int
	Mainchaintokenlocknonce *big.Int
	Vouchermintnonce        *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterMinttfuelvouchers is a free log retrieval operation binding the contract event 0x92c9b88a988ad885ea7d46bf30711e85e9e12bec8ec93ad7878a9bf03c73a36b.
//
// Solidity: event minttfuelvouchers(address voucherreceiver, uint256 amount, uint256 mainchaintokenlocknonce, uint256 vouchermintnonce)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankFilterer) FilterMinttfuelvouchers(opts *bind.FilterOpts) (*SubchainTFuelTokenBankMinttfuelvouchersIterator, error) {

	logs, sub, err := _SubchainTFuelTokenBank.contract.FilterLogs(opts, "minttfuelvouchers")
	if err != nil {
		return nil, err
	}
	return &SubchainTFuelTokenBankMinttfuelvouchersIterator{contract: _SubchainTFuelTokenBank.contract, event: "minttfuelvouchers", logs: logs, sub: sub}, nil
}

// WatchMinttfuelvouchers is a free log subscription operation binding the contract event 0x92c9b88a988ad885ea7d46bf30711e85e9e12bec8ec93ad7878a9bf03c73a36b.
//
// Solidity: event minttfuelvouchers(address voucherreceiver, uint256 amount, uint256 mainchaintokenlocknonce, uint256 vouchermintnonce)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankFilterer) WatchMinttfuelvouchers(opts *bind.WatchOpts, sink chan<- *SubchainTFuelTokenBankMinttfuelvouchers) (event.Subscription, error) {

	logs, sub, err := _SubchainTFuelTokenBank.contract.WatchLogs(opts, "minttfuelvouchers")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubchainTFuelTokenBankMinttfuelvouchers)
				if err := _SubchainTFuelTokenBank.contract.UnpackLog(event, "minttfuelvouchers", log); err != nil {
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

// ParseMinttfuelvouchers is a log parse operation binding the contract event 0x92c9b88a988ad885ea7d46bf30711e85e9e12bec8ec93ad7878a9bf03c73a36b.
//
// Solidity: event minttfuelvouchers(address voucherreceiver, uint256 amount, uint256 mainchaintokenlocknonce, uint256 vouchermintnonce)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankFilterer) ParseMinttfuelvouchers(log types.Log) (*SubchainTFuelTokenBankMinttfuelvouchers, error) {
	event := new(SubchainTFuelTokenBankMinttfuelvouchers)
	if err := _SubchainTFuelTokenBank.contract.UnpackLog(event, "minttfuelvouchers", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
