// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/thetatoken/thetasubchain/eth"
	"github.com/thetatoken/thetasubchain/eth/abi"
	"github.com/thetatoken/thetasubchain/eth/abi/bind"
	"github.com/thetatoken/thetasubchain/eth/core/types"
	"github.com/thetatoken/thetasubchain/eth/event"

	"github.com/thetatoken/theta/common"
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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voucherOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"mainchainTokenReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"BurnTFuelVouchers\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voucherReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MintTFuelVouchers\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allDenoms\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allVouchers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"denomToVoucherLookup\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voucherAddress\",\"type\":\"address\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voucherContractAddr\",\"type\":\"address\"}],\"name\":\"getDenom\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"name\":\"getVoucher\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"voucherAddressToDenomLookup\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voucherBurnNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voucherReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"}],\"name\":\"mintVouchers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voucherOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mainchainTokenReceiver\",\"type\":\"address\"}],\"name\":\"burnVouchers\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060008081905550611aec806100276000396000f3fe60806040526004361061009c5760003560e01c8063a2cc698111610064578063a2cc6981146101d4578063c27d927a14610211578063d6c7e0d41461023c578063da837d5a14610258578063ebda996214610281578063f6a3d24e146102be5761009c565b80631527b14d146100a1578063261a323e146100df57806327ca4df11461011c578063588b14081461015957806360569b5e14610196575b600080fd5b3480156100ad57600080fd5b506100c860048036038101906100c391906110c4565b6102fb565b6040516100d692919061133f565b60405180910390f35b3480156100eb57600080fd5b50610106600480360381019061010191906110c4565b610362565b6040516101139190611391565b60405180910390f35b34801561012857600080fd5b50610143600480360381019061013e919061110d565b6103a7565b60405161015091906112df565b60405180910390f35b34801561016557600080fd5b50610180600480360381019061017b919061110d565b6103e6565b60405161018d91906113ac565b60405180910390f35b3480156101a257600080fd5b506101bd60048036038101906101b89190611017565b610492565b6040516101cb9291906113ce565b60405180910390f35b3480156101e057600080fd5b506101fb60048036038101906101f691906110c4565b61054b565b60405161020891906112df565b60405180910390f35b34801561021d57600080fd5b5061022661061b565b604051610233919061149e565b60405180910390f35b61025660048036038101906102519190611044565b610621565b005b34801561026457600080fd5b5061027f600480360381019061027a9190611084565b61067b565b005b34801561028d57600080fd5b506102a860048036038101906102a39190611017565b61081d565b6040516102b591906113ac565b60405180910390f35b3480156102ca57600080fd5b506102e560048036038101906102e09190611017565b61094a565b6040516102f29190611391565b60405180910390f35b6001818051602081018201805184825260208301602085012081835280955050505050506000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060000160149054906101000a900460ff16905082565b60008061036e836109a3565b905060018160405161038091906112c8565b908152602001604051809103902060000160149054906101000a900460ff16915050919050565b600381815481106103b757600080fd5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600481815481106103f657600080fd5b906000526020600020016000915090508054610411906116fe565b80601f016020809104026020016040519081016040528092919081815260200182805461043d906116fe565b801561048a5780601f1061045f5761010080835404028352916020019161048a565b820191906000526020600020905b81548152906001019060200180831161046d57829003601f168201915b505050505081565b60026020528060005260406000206000915090508060000180546104b5906116fe565b80601f01602080910402602001604051908101604052809291908181526020018280546104e1906116fe565b801561052e5780601f106105035761010080835404028352916020019161052e565b820191906000526020600020905b81548152906001019060200180831161051157829003601f168201915b5050505050908060010160009054906101000a900460ff16905082565b600080610557836109a3565b9050600060018260405161056b91906112c8565b90815260200160405180910390206040518060400160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160149054906101000a900460ff161515151581525050905080602001511561060f57806000015192505050610616565b6000925050505b919050565b60005481565b600034905061062f816109b5565b610637610b3c565b7f2d8470f382989926e91872b25e27240426cc2efe1bfef95339ae1944a221708f83838360005460405161066e94939291906112fa565b60405180910390a1505050565b600080602067ffffffffffffffff81111561069957610698611837565b5b6040519080825280601f01601f1916602001820160405280156106cb5781602001600182028036833780820191505090505b50905060008060b573ffffffffffffffffffffffffffffffffffffffff16836040516106f791906112b1565b6000604051808303816000865af19150503d8060008114610734576040519150601f19603f3d011682016040523d82523d6000602084013e610739565b606091505b50915091508161077e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610775906113fe565b60405180910390fd5b600061078982610b57565b9050600181149450846107d1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107c89061145e565b60405180910390fd5b6107db8787610c40565b7fed2a6b76dd30ca61a3c463f15ebbe687c91c02dad5ceea323d771ae5e780e3d1878760405161080c929190611368565b60405180910390a150505050505050565b60606000600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060405180604001604052908160008201805461087b906116fe565b80601f01602080910402602001604051908101604052809291908181526020018280546108a7906116fe565b80156108f45780601f106108c9576101008083540402835291602001916108f4565b820191906000526020600020905b8154815290600101906020018083116108d757829003601f168201915b505050505081526020016001820160009054906101000a900460ff1615151515815250509050806020015115610931578060000151915050610945565b604051806020016040528060008152509150505b919050565b6000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900460ff169050919050565b60606109ae82610e5f565b9050919050565b6000602067ffffffffffffffff8111156109d2576109d1611837565b5b6040519080825280601f01601f191660200182016040528015610a045781602001600182028036833780820191505090505b50905060008260001b905060005b6020811015610a8757818160208110610a2e57610a2d611808565b5b1a60f81b838281518110610a4557610a44611808565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508080610a7f90611761565b915050610a12565b50600060b773ffffffffffffffffffffffffffffffffffffffff1683604051610ab091906112b1565b6000604051808303816000865af19150503d8060008114610aed576040519150601f19603f3d011682016040523d82523d6000602084013e610af2565b606091505b5050905080610b36576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b2d9061143e565b60405180910390fd5b50505050565b6001600080828254610b4e919061154c565b92505081905550565b6000806000835190506020811115610ba4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b9b9061141e565b60405180910390fd5b60005b81811015610c3257600881836020610bbf9190611633565b610bc9919061154c565b610bd391906115d9565b60ff60f81b868381518110610beb57610bea611808565b5b602001015160f81c60f81b167effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916901c831792508080610c2a90611761565b915050610ba7565b508160001c92505050919050565b6000603467ffffffffffffffff811115610c5d57610c5c611837565b5b6040519080825280601f01601f191660200182016040528015610c8f5781602001600182028036833780820191505090505b50905060008260001b90506000805b6014811015610d22578560601b8160148110610cbd57610cbc611808565b5b1a60f81b848380610ccd90611761565b945081518110610ce057610cdf611808565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508080610d1a90611761565b915050610c9e565b600090505b6020811015610da857828160208110610d4357610d42611808565b5b1a60f81b848380610d5390611761565b945081518110610d6657610d65611808565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508080610da090611761565b915050610d27565b600060b673ffffffffffffffffffffffffffffffffffffffff1685604051610dd091906112b1565b6000604051808303816000865af19150503d8060008114610e0d576040519150601f19603f3d011682016040523d82523d6000602084013e610e12565b606091505b5050905080610e56576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e4d9061147e565b60405180910390fd5b50505050505050565b6060600082905060005b8151811015610eed57610e98828281518110610e8857610e87611808565b5b602001015160f81c60f81b610ef7565b828281518110610eab57610eaa611808565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508080610ee590611761565b915050610e69565b5080915050919050565b6000604160f81b827effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191610158015610f555750605a60f81b827effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191611155b15610f745760208260f81c610f6a91906115a2565b60f81b9050610f78565b8190505b919050565b6000610f90610f8b846114de565b6114b9565b905082815260208101848484011115610fac57610fab61186b565b5b610fb78482856116bc565b509392505050565b600081359050610fce81611a88565b92915050565b600082601f830112610fe957610fe8611866565b5b8135610ff9848260208601610f7d565b91505092915050565b60008135905061101181611a9f565b92915050565b60006020828403121561102d5761102c611875565b5b600061103b84828501610fbf565b91505092915050565b6000806040838503121561105b5761105a611875565b5b600061106985828601610fbf565b925050602061107a85828601610fbf565b9150509250929050565b6000806040838503121561109b5761109a611875565b5b60006110a985828601610fbf565b92505060206110ba85828601611002565b9150509250929050565b6000602082840312156110da576110d9611875565b5b600082013567ffffffffffffffff8111156110f8576110f7611870565b5b61110484828501610fd4565b91505092915050565b60006020828403121561112357611122611875565b5b600061113184828501611002565b91505092915050565b61114381611667565b82525050565b61115281611679565b82525050565b60006111638261150f565b61116d8185611525565b935061117d8185602086016116cb565b80840191505092915050565b60006111948261151a565b61119e8185611530565b93506111ae8185602086016116cb565b6111b78161187a565b840191505092915050565b60006111cd8261151a565b6111d78185611541565b93506111e78185602086016116cb565b80840191505092915050565b6000611200604d83611530565b915061120b8261188b565b606082019050919050565b6000611223602583611530565b915061122e82611900565b604082019050919050565b6000611246604883611530565b91506112518261194f565b606082019050919050565b6000611269603883611530565b9150611274826119c4565b604082019050919050565b600061128c604783611530565b915061129782611a13565b606082019050919050565b6112ab816116a5565b82525050565b60006112bd8284611158565b915081905092915050565b60006112d482846111c2565b915081905092915050565b60006020820190506112f4600083018461113a565b92915050565b600060808201905061130f600083018761113a565b61131c602083018661113a565b61132960408301856112a2565b61133660608301846112a2565b95945050505050565b6000604082019050611354600083018561113a565b6113616020830184611149565b9392505050565b600060408201905061137d600083018561113a565b61138a60208301846112a2565b9392505050565b60006020820190506113a66000830184611149565b92915050565b600060208201905081810360008301526113c68184611189565b905092915050565b600060408201905081810360008301526113e88185611189565b90506113f76020830184611149565b9392505050565b60006020820190508181036000830152611417816111f3565b9050919050565b6000602082019050818103600083015261143781611216565b9050919050565b6000602082019050818103600083015261145781611239565b9050919050565b600060208201905081810360008301526114778161125c565b9050919050565b600060208201905081810360008301526114978161127f565b9050919050565b60006020820190506114b360008301846112a2565b92915050565b60006114c36114d4565b90506114cf8282611730565b919050565b6000604051905090565b600067ffffffffffffffff8211156114f9576114f8611837565b5b6115028261187a565b9050602081019050919050565b600081519050919050565b600081519050919050565b600081905092915050565b600082825260208201905092915050565b600081905092915050565b6000611557826116a5565b9150611562836116a5565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611597576115966117aa565b5b828201905092915050565b60006115ad826116af565b91506115b8836116af565b92508260ff038211156115ce576115cd6117aa565b5b828201905092915050565b60006115e4826116a5565b91506115ef836116a5565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611628576116276117aa565b5b828202905092915050565b600061163e826116a5565b9150611649836116a5565b92508282101561165c5761165b6117aa565b5b828203905092915050565b600061167282611685565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b838110156116e95780820151818401526020810190506116ce565b838111156116f8576000848401525b50505050565b6000600282049050600182168061171657607f821691505b6020821081141561172a576117296117d9565b5b50919050565b6117398261187a565b810181811067ffffffffffffffff8211171561175857611757611837565b5b80604052505050565b600061176c826116a5565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82141561179f5761179e6117aa565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f537562636861696e544675656c546f6b656e42616e6b2e70726976696c65676560008201527f644163636573734f6e6c793a206661696c656420746f20636865636b2074686560208201527f20616363657373206c6576656c00000000000000000000000000000000000000604082015250565b7f627974657320746f2075696e7432353620636f6e76657273696f6e206f76657260008201527f666c6f7773000000000000000000000000000000000000000000000000000000602082015250565b7f537562636861696e544675656c546f6b656e42616e6b2e5f6275726e5446756560008201527f6c566f7563686572733a206661696c656420746f206275726e20544675656c2060208201527f766f756368657273000000000000000000000000000000000000000000000000604082015250565b7f566f75636865724d61702e70726976696c656765644163636573734f6e6c793a60008201527f20696e73756666696369656e742070726976696c656765730000000000000000602082015250565b7f537562636861696e544675656c546f6b656e42616e6b2e6d696e74544675656c60008201527f566f7563686572733a206661696c656420746f206d696e7420544675656c207660208201527f6f75636865727300000000000000000000000000000000000000000000000000604082015250565b611a9181611667565b8114611a9c57600080fd5b50565b611aa8816116a5565b8114611ab357600080fd5b5056fea2646970667358221220bf3f2dd8085dd741620d18eafb8cb501fce6a5a3bc50596665a23d23e051821364736f6c63430008070033",
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

// AllDenoms is a free data retrieval call binding the contract method 0x588b1408.
//
// Solidity: function allDenoms(uint256 ) view returns(string)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankCaller) AllDenoms(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _SubchainTFuelTokenBank.contract.Call(opts, &out, "allDenoms", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AllDenoms is a free data retrieval call binding the contract method 0x588b1408.
//
// Solidity: function allDenoms(uint256 ) view returns(string)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankSession) AllDenoms(arg0 *big.Int) (string, error) {
	return _SubchainTFuelTokenBank.Contract.AllDenoms(&_SubchainTFuelTokenBank.CallOpts, arg0)
}

// AllDenoms is a free data retrieval call binding the contract method 0x588b1408.
//
// Solidity: function allDenoms(uint256 ) view returns(string)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankCallerSession) AllDenoms(arg0 *big.Int) (string, error) {
	return _SubchainTFuelTokenBank.Contract.AllDenoms(&_SubchainTFuelTokenBank.CallOpts, arg0)
}

// AllVouchers is a free data retrieval call binding the contract method 0x27ca4df1.
//
// Solidity: function allVouchers(uint256 ) view returns(address)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankCaller) AllVouchers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SubchainTFuelTokenBank.contract.Call(opts, &out, "allVouchers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllVouchers is a free data retrieval call binding the contract method 0x27ca4df1.
//
// Solidity: function allVouchers(uint256 ) view returns(address)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankSession) AllVouchers(arg0 *big.Int) (common.Address, error) {
	return _SubchainTFuelTokenBank.Contract.AllVouchers(&_SubchainTFuelTokenBank.CallOpts, arg0)
}

// AllVouchers is a free data retrieval call binding the contract method 0x27ca4df1.
//
// Solidity: function allVouchers(uint256 ) view returns(address)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankCallerSession) AllVouchers(arg0 *big.Int) (common.Address, error) {
	return _SubchainTFuelTokenBank.Contract.AllVouchers(&_SubchainTFuelTokenBank.CallOpts, arg0)
}

// DenomToVoucherLookup is a free data retrieval call binding the contract method 0x1527b14d.
//
// Solidity: function denomToVoucherLookup(string ) view returns(address contractAddress, bool exists)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankCaller) DenomToVoucherLookup(opts *bind.CallOpts, arg0 string) (struct {
	ContractAddress common.Address
	Exists          bool
}, error) {
	var out []interface{}
	err := _SubchainTFuelTokenBank.contract.Call(opts, &out, "denomToVoucherLookup", arg0)

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
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankSession) DenomToVoucherLookup(arg0 string) (struct {
	ContractAddress common.Address
	Exists          bool
}, error) {
	return _SubchainTFuelTokenBank.Contract.DenomToVoucherLookup(&_SubchainTFuelTokenBank.CallOpts, arg0)
}

// DenomToVoucherLookup is a free data retrieval call binding the contract method 0x1527b14d.
//
// Solidity: function denomToVoucherLookup(string ) view returns(address contractAddress, bool exists)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankCallerSession) DenomToVoucherLookup(arg0 string) (struct {
	ContractAddress common.Address
	Exists          bool
}, error) {
	return _SubchainTFuelTokenBank.Contract.DenomToVoucherLookup(&_SubchainTFuelTokenBank.CallOpts, arg0)
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
// Solidity: function exists(address voucherAddress) view returns(bool)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankCaller) Exists0(opts *bind.CallOpts, voucherAddress common.Address) (bool, error) {
	var out []interface{}
	err := _SubchainTFuelTokenBank.contract.Call(opts, &out, "exists0", voucherAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists0 is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address voucherAddress) view returns(bool)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankSession) Exists0(voucherAddress common.Address) (bool, error) {
	return _SubchainTFuelTokenBank.Contract.Exists0(&_SubchainTFuelTokenBank.CallOpts, voucherAddress)
}

// Exists0 is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address voucherAddress) view returns(bool)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankCallerSession) Exists0(voucherAddress common.Address) (bool, error) {
	return _SubchainTFuelTokenBank.Contract.Exists0(&_SubchainTFuelTokenBank.CallOpts, voucherAddress)
}

// GetDenom is a free data retrieval call binding the contract method 0xebda9962.
//
// Solidity: function getDenom(address voucherContractAddr) view returns(string)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankCaller) GetDenom(opts *bind.CallOpts, voucherContractAddr common.Address) (string, error) {
	var out []interface{}
	err := _SubchainTFuelTokenBank.contract.Call(opts, &out, "getDenom", voucherContractAddr)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetDenom is a free data retrieval call binding the contract method 0xebda9962.
//
// Solidity: function getDenom(address voucherContractAddr) view returns(string)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankSession) GetDenom(voucherContractAddr common.Address) (string, error) {
	return _SubchainTFuelTokenBank.Contract.GetDenom(&_SubchainTFuelTokenBank.CallOpts, voucherContractAddr)
}

// GetDenom is a free data retrieval call binding the contract method 0xebda9962.
//
// Solidity: function getDenom(address voucherContractAddr) view returns(string)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankCallerSession) GetDenom(voucherContractAddr common.Address) (string, error) {
	return _SubchainTFuelTokenBank.Contract.GetDenom(&_SubchainTFuelTokenBank.CallOpts, voucherContractAddr)
}

// GetVoucher is a free data retrieval call binding the contract method 0xa2cc6981.
//
// Solidity: function getVoucher(string denom) view returns(address)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankCaller) GetVoucher(opts *bind.CallOpts, denom string) (common.Address, error) {
	var out []interface{}
	err := _SubchainTFuelTokenBank.contract.Call(opts, &out, "getVoucher", denom)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVoucher is a free data retrieval call binding the contract method 0xa2cc6981.
//
// Solidity: function getVoucher(string denom) view returns(address)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankSession) GetVoucher(denom string) (common.Address, error) {
	return _SubchainTFuelTokenBank.Contract.GetVoucher(&_SubchainTFuelTokenBank.CallOpts, denom)
}

// GetVoucher is a free data retrieval call binding the contract method 0xa2cc6981.
//
// Solidity: function getVoucher(string denom) view returns(address)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankCallerSession) GetVoucher(denom string) (common.Address, error) {
	return _SubchainTFuelTokenBank.Contract.GetVoucher(&_SubchainTFuelTokenBank.CallOpts, denom)
}

// VoucherAddressToDenomLookup is a free data retrieval call binding the contract method 0x60569b5e.
//
// Solidity: function voucherAddressToDenomLookup(address ) view returns(string denom, bool exists)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankCaller) VoucherAddressToDenomLookup(opts *bind.CallOpts, arg0 common.Address) (struct {
	Denom  string
	Exists bool
}, error) {
	var out []interface{}
	err := _SubchainTFuelTokenBank.contract.Call(opts, &out, "voucherAddressToDenomLookup", arg0)

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
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankSession) VoucherAddressToDenomLookup(arg0 common.Address) (struct {
	Denom  string
	Exists bool
}, error) {
	return _SubchainTFuelTokenBank.Contract.VoucherAddressToDenomLookup(&_SubchainTFuelTokenBank.CallOpts, arg0)
}

// VoucherAddressToDenomLookup is a free data retrieval call binding the contract method 0x60569b5e.
//
// Solidity: function voucherAddressToDenomLookup(address ) view returns(string denom, bool exists)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankCallerSession) VoucherAddressToDenomLookup(arg0 common.Address) (struct {
	Denom  string
	Exists bool
}, error) {
	return _SubchainTFuelTokenBank.Contract.VoucherAddressToDenomLookup(&_SubchainTFuelTokenBank.CallOpts, arg0)
}

// VoucherBurnNonce is a free data retrieval call binding the contract method 0xc27d927a.
//
// Solidity: function voucherBurnNonce() view returns(uint256)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankCaller) VoucherBurnNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SubchainTFuelTokenBank.contract.Call(opts, &out, "voucherBurnNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VoucherBurnNonce is a free data retrieval call binding the contract method 0xc27d927a.
//
// Solidity: function voucherBurnNonce() view returns(uint256)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankSession) VoucherBurnNonce() (*big.Int, error) {
	return _SubchainTFuelTokenBank.Contract.VoucherBurnNonce(&_SubchainTFuelTokenBank.CallOpts)
}

// VoucherBurnNonce is a free data retrieval call binding the contract method 0xc27d927a.
//
// Solidity: function voucherBurnNonce() view returns(uint256)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankCallerSession) VoucherBurnNonce() (*big.Int, error) {
	return _SubchainTFuelTokenBank.Contract.VoucherBurnNonce(&_SubchainTFuelTokenBank.CallOpts)
}

// BurnVouchers is a paid mutator transaction binding the contract method 0xd6c7e0d4.
//
// Solidity: function burnVouchers(address voucherOwner, address mainchainTokenReceiver) payable returns()
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankTransactor) BurnVouchers(opts *bind.TransactOpts, voucherOwner common.Address, mainchainTokenReceiver common.Address) (*types.Transaction, error) {
	return _SubchainTFuelTokenBank.contract.Transact(opts, "burnVouchers", voucherOwner, mainchainTokenReceiver)
}

// BurnVouchers is a paid mutator transaction binding the contract method 0xd6c7e0d4.
//
// Solidity: function burnVouchers(address voucherOwner, address mainchainTokenReceiver) payable returns()
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankSession) BurnVouchers(voucherOwner common.Address, mainchainTokenReceiver common.Address) (*types.Transaction, error) {
	return _SubchainTFuelTokenBank.Contract.BurnVouchers(&_SubchainTFuelTokenBank.TransactOpts, voucherOwner, mainchainTokenReceiver)
}

// BurnVouchers is a paid mutator transaction binding the contract method 0xd6c7e0d4.
//
// Solidity: function burnVouchers(address voucherOwner, address mainchainTokenReceiver) payable returns()
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankTransactorSession) BurnVouchers(voucherOwner common.Address, mainchainTokenReceiver common.Address) (*types.Transaction, error) {
	return _SubchainTFuelTokenBank.Contract.BurnVouchers(&_SubchainTFuelTokenBank.TransactOpts, voucherOwner, mainchainTokenReceiver)
}

// MintVouchers is a paid mutator transaction binding the contract method 0xda837d5a.
//
// Solidity: function mintVouchers(address voucherReceiver, uint256 mintAmount) returns()
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankTransactor) MintVouchers(opts *bind.TransactOpts, voucherReceiver common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _SubchainTFuelTokenBank.contract.Transact(opts, "mintVouchers", voucherReceiver, mintAmount)
}

// MintVouchers is a paid mutator transaction binding the contract method 0xda837d5a.
//
// Solidity: function mintVouchers(address voucherReceiver, uint256 mintAmount) returns()
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankSession) MintVouchers(voucherReceiver common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _SubchainTFuelTokenBank.Contract.MintVouchers(&_SubchainTFuelTokenBank.TransactOpts, voucherReceiver, mintAmount)
}

// MintVouchers is a paid mutator transaction binding the contract method 0xda837d5a.
//
// Solidity: function mintVouchers(address voucherReceiver, uint256 mintAmount) returns()
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankTransactorSession) MintVouchers(voucherReceiver common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _SubchainTFuelTokenBank.Contract.MintVouchers(&_SubchainTFuelTokenBank.TransactOpts, voucherReceiver, mintAmount)
}

// SubchainTFuelTokenBankBurnTFuelVouchersIterator is returned from FilterBurnTFuelVouchers and is used to iterate over the raw logs and unpacked data for BurnTFuelVouchers events raised by the SubchainTFuelTokenBank contract.
type SubchainTFuelTokenBankBurnTFuelVouchersIterator struct {
	Event *SubchainTFuelTokenBankBurnTFuelVouchers // Event containing the contract specifics and raw log

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
func (it *SubchainTFuelTokenBankBurnTFuelVouchersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubchainTFuelTokenBankBurnTFuelVouchers)
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
		it.Event = new(SubchainTFuelTokenBankBurnTFuelVouchers)
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
func (it *SubchainTFuelTokenBankBurnTFuelVouchersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubchainTFuelTokenBankBurnTFuelVouchersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubchainTFuelTokenBankBurnTFuelVouchers represents a BurnTFuelVouchers event raised by the SubchainTFuelTokenBank contract.
type SubchainTFuelTokenBankBurnTFuelVouchers struct {
	VoucherOwner           common.Address
	MainchainTokenReceiver common.Address
	Amount                 *big.Int
	Nonce                  *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterBurnTFuelVouchers is a free log retrieval operation binding the contract event 0x2d8470f382989926e91872b25e27240426cc2efe1bfef95339ae1944a221708f.
//
// Solidity: event BurnTFuelVouchers(address voucherOwner, address mainchainTokenReceiver, uint256 amount, uint256 nonce)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankFilterer) FilterBurnTFuelVouchers(opts *bind.FilterOpts) (*SubchainTFuelTokenBankBurnTFuelVouchersIterator, error) {

	logs, sub, err := _SubchainTFuelTokenBank.contract.FilterLogs(opts, "BurnTFuelVouchers")
	if err != nil {
		return nil, err
	}
	return &SubchainTFuelTokenBankBurnTFuelVouchersIterator{contract: _SubchainTFuelTokenBank.contract, event: "BurnTFuelVouchers", logs: logs, sub: sub}, nil
}

// WatchBurnTFuelVouchers is a free log subscription operation binding the contract event 0x2d8470f382989926e91872b25e27240426cc2efe1bfef95339ae1944a221708f.
//
// Solidity: event BurnTFuelVouchers(address voucherOwner, address mainchainTokenReceiver, uint256 amount, uint256 nonce)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankFilterer) WatchBurnTFuelVouchers(opts *bind.WatchOpts, sink chan<- *SubchainTFuelTokenBankBurnTFuelVouchers) (event.Subscription, error) {

	logs, sub, err := _SubchainTFuelTokenBank.contract.WatchLogs(opts, "BurnTFuelVouchers")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubchainTFuelTokenBankBurnTFuelVouchers)
				if err := _SubchainTFuelTokenBank.contract.UnpackLog(event, "BurnTFuelVouchers", log); err != nil {
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

// ParseBurnTFuelVouchers is a log parse operation binding the contract event 0x2d8470f382989926e91872b25e27240426cc2efe1bfef95339ae1944a221708f.
//
// Solidity: event BurnTFuelVouchers(address voucherOwner, address mainchainTokenReceiver, uint256 amount, uint256 nonce)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankFilterer) ParseBurnTFuelVouchers(log types.Log) (*SubchainTFuelTokenBankBurnTFuelVouchers, error) {
	event := new(SubchainTFuelTokenBankBurnTFuelVouchers)
	if err := _SubchainTFuelTokenBank.contract.UnpackLog(event, "BurnTFuelVouchers", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubchainTFuelTokenBankMintTFuelVouchersIterator is returned from FilterMintTFuelVouchers and is used to iterate over the raw logs and unpacked data for MintTFuelVouchers events raised by the SubchainTFuelTokenBank contract.
type SubchainTFuelTokenBankMintTFuelVouchersIterator struct {
	Event *SubchainTFuelTokenBankMintTFuelVouchers // Event containing the contract specifics and raw log

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
func (it *SubchainTFuelTokenBankMintTFuelVouchersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubchainTFuelTokenBankMintTFuelVouchers)
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
		it.Event = new(SubchainTFuelTokenBankMintTFuelVouchers)
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
func (it *SubchainTFuelTokenBankMintTFuelVouchersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubchainTFuelTokenBankMintTFuelVouchersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubchainTFuelTokenBankMintTFuelVouchers represents a MintTFuelVouchers event raised by the SubchainTFuelTokenBank contract.
type SubchainTFuelTokenBankMintTFuelVouchers struct {
	VoucherReceiver common.Address
	Amount          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMintTFuelVouchers is a free log retrieval operation binding the contract event 0xed2a6b76dd30ca61a3c463f15ebbe687c91c02dad5ceea323d771ae5e780e3d1.
//
// Solidity: event MintTFuelVouchers(address voucherReceiver, uint256 amount)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankFilterer) FilterMintTFuelVouchers(opts *bind.FilterOpts) (*SubchainTFuelTokenBankMintTFuelVouchersIterator, error) {

	logs, sub, err := _SubchainTFuelTokenBank.contract.FilterLogs(opts, "MintTFuelVouchers")
	if err != nil {
		return nil, err
	}
	return &SubchainTFuelTokenBankMintTFuelVouchersIterator{contract: _SubchainTFuelTokenBank.contract, event: "MintTFuelVouchers", logs: logs, sub: sub}, nil
}

// WatchMintTFuelVouchers is a free log subscription operation binding the contract event 0xed2a6b76dd30ca61a3c463f15ebbe687c91c02dad5ceea323d771ae5e780e3d1.
//
// Solidity: event MintTFuelVouchers(address voucherReceiver, uint256 amount)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankFilterer) WatchMintTFuelVouchers(opts *bind.WatchOpts, sink chan<- *SubchainTFuelTokenBankMintTFuelVouchers) (event.Subscription, error) {

	logs, sub, err := _SubchainTFuelTokenBank.contract.WatchLogs(opts, "MintTFuelVouchers")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubchainTFuelTokenBankMintTFuelVouchers)
				if err := _SubchainTFuelTokenBank.contract.UnpackLog(event, "MintTFuelVouchers", log); err != nil {
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

// ParseMintTFuelVouchers is a log parse operation binding the contract event 0xed2a6b76dd30ca61a3c463f15ebbe687c91c02dad5ceea323d771ae5e780e3d1.
//
// Solidity: event MintTFuelVouchers(address voucherReceiver, uint256 amount)
func (_SubchainTFuelTokenBank *SubchainTFuelTokenBankFilterer) ParseMintTFuelVouchers(log types.Log) (*SubchainTFuelTokenBankMintTFuelVouchers, error) {
	event := new(SubchainTFuelTokenBankMintTFuelVouchers)
	if err := _SubchainTFuelTokenBank.contract.UnpackLog(event, "MintTFuelVouchers", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}