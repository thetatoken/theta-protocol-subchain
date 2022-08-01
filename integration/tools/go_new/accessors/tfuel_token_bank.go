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

// TFuelTokenBankMetaData contains all meta data concerning the TFuelTokenBank contract.
var TFuelTokenBankMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"mainchainID_\",\"type\":\"uint256\"},{\"internalType\":\"contractChainRegistrar\",\"name\":\"chainRegistrar_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sourceChainTokenSender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"targetChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"targetChainVoucherReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenLockNonce\",\"type\":\"uint256\"}],\"name\":\"TFuelTokenLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"targetChainTokenReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unlockedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainVoucherBurnNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenUnlockNonce\",\"type\":\"uint256\"}],\"name\":\"TFuelTokenUnlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sourceChainVoucherOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"targetChainTokenReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"burnedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"voucherBurnNonce\",\"type\":\"uint256\"}],\"name\":\"TFuelVoucherBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"targetChainVoucherReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainTokenLockNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"voucherMintNonce\",\"type\":\"uint256\"}],\"name\":\"TFuelVoucherMinted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allDenoms\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allVouchers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"denomToVoucherLookup\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voucherAddress\",\"type\":\"address\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voucherContractAddr\",\"type\":\"address\"}],\"name\":\"getDenom\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"getMaxProcessedTokenLockNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"getMaxProcessedVoucherBurnNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"name\":\"getVoucher\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOnMainchain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mainchainID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenLockNonceMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"tokenLockVotingRecords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dynasty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumlatedShares\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenUnlockNonceMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"totalLockedAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"voucherAddressToDenomLookup\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"voucherBurnNonceMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"voucherBurnVotingRecords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dynasty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumlatedShares\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"voucherMintNonceMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"id\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"targetChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"targetChainVoucherReceiver\",\"type\":\"address\"}],\"name\":\"lockTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"targetChainVoucherReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mintedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dynasty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainTokenLockNonce\",\"type\":\"uint256\"}],\"name\":\"mintVouchers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targetChainTokenReceiver\",\"type\":\"address\"}],\"name\":\"burnVouchers\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainID\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"targetChainTokenReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"unlockAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dynasty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainVoucherBurnNonce\",\"type\":\"uint256\"}],\"name\":\"unlockTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200268638038062002686833981016040819052620000349162000064565b6001600081905591909155600280546001600160a01b0319166001600160a01b03909216919091179055620000a3565b600080604083850312156200007857600080fd5b825160208401519092506001600160a01b03811681146200009857600080fd5b809150509250929050565b6125d380620000b36000396000f3fe60806040526004361061014b5760003560e01c80637d0fb00d116100b6578063ccf187c71161006f578063ccf187c714610472578063ebda99621461049f578063f6a3d24e146104bf578063f95627ac146104fb578063feaff05214610528578063ff248a441461056757600080fd5b80637d0fb00d146103bf5780638883931e146103d2578063a2cc6981146103ff578063aa68acde1461041f578063af640d0f14610432578063ca2075691461044557600080fd5b806327ca4df11161010857806327ca4df1146102b85780634250863b146102f0578063588b14081461030a57806360569b5e14610337578063740cb7f814610365578063766f8fb01461039257600080fd5b8063073b9502146101505780631527b14d1461017957806319fd1a11146101e55780631a0483d3146102125780631eb7873714610234578063261a323e14610288575b600080fd5b34801561015c57600080fd5b5061016660015481565b6040519081526020015b60405180910390f35b34801561018557600080fd5b506101c6610194366004612011565b8051602081830181018051600b825292820191909301209152546001600160a01b03811690600160a01b900460ff1682565b604080516001600160a01b039093168352901515602083015201610170565b3480156101f157600080fd5b506101666102003660046120b0565b60106020526000908152604090205481565b34801561021e57600080fd5b5061023261022d366004612046565b61057a565b005b34801561024057600080fd5b5061027361024f366004612123565b60096020908152600092835260408084209091529082529020805460029091015482565b60408051928352602083019190915201610170565b34801561029457600080fd5b506102a86102a3366004612011565b61077c565b6040519015158152602001610170565b3480156102c457600080fd5b506102d86102d33660046120b0565b6107bc565b6040516001600160a01b039091168152602001610170565b3480156102fc57600080fd5b50600f546102a89060ff1681565b34801561031657600080fd5b5061032a6103253660046120b0565b6107e6565b6040516101709190612240565b34801561034357600080fd5b50610357610352366004611f06565b610892565b604051610170929190612319565b34801561037157600080fd5b506101666103803660046120b0565b60066020526000908152604090205481565b34801561039e57600080fd5b506101666103ad3660046120b0565b60009081526008602052604090205490565b6102326103cd366004611f06565b610939565b3480156103de57600080fd5b506101666103ed3660046120b0565b60036020526000908152604090205481565b34801561040b57600080fd5b506102d861041a366004612011565b610ab1565b61023261042d3660046120c9565b610b22565b34801561043e57600080fd5b5046610166565b34801561045157600080fd5b506101666104603660046120b0565b60056020526000908152604090205481565b34801561047e57600080fd5b5061016661048d3660046120b0565b60046020526000908152604090205481565b3480156104ab57600080fd5b5061032a6104ba366004611f06565b610d65565b3480156104cb57600080fd5b506102a86104da366004611f06565b6001600160a01b03166000908152600c602052604090206001015460ff1690565b34801561050757600080fd5b506101666105163660046120b0565b60009081526007602052604090205490565b34801561053457600080fd5b50610273610543366004612123565b600a6020908152600092835260408084209091529082529020805460029091015482565b6102326105753660046120f9565b610e5c565b600260005414156105a65760405162461bcd60e51b815260040161059d9061233d565b60405180910390fd5b60026000556001544614156106235760405162461bcd60e51b815260206004820152603e60248201527f544675656c546f6b656e42616e6b2e6d696e74566f756368657273282920636160448201527f6e206f6e6c792062652063616c6c6564206f6e206120737562636861696e0000606482015260840161059d565b600061062e866110f5565b905060008061063c83611106565b91509150806106995760405162461bcd60e51b8152602060048201526024808201527f4661696c656420746f206578747261637420636861696e49442066726f6d2064604482015263656e6f6d60e01b606482015260840161059d565b60008284898989896040516020016106b6969594939291906121e7565b60405160208183030381529060405280519060200120905060006106dd848884893361118e565b9050801561076b576106ef89896111ab565b6106fb6205814961138a565b6205814960005260066020527f6b62aca1b1fa87e136d72c3c175fe558334fbeab93a31e41647f2a6c16a7e92a546040517f80742bd15a2c8c4ad5d395bcf577073110e52f0c73bf980dfa9453c1d8c354e590610761908d908d908d908c908790612253565b60405180910390a1505b505060016000555050505050505050565b600080610788836110f5565b9050600b8160405161079a9190612171565b9081526040519081900360200190205460ff600160a01b909104169392505050565b600d81815481106107cc57600080fd5b6000918252602090912001546001600160a01b0316905081565b600e81815481106107f657600080fd5b9060005260206000200160009150905080546108119061249f565b80601f016020809104026020016040519081016040528092919081815260200182805461083d9061249f565b801561088a5780601f1061085f5761010080835404028352916020019161088a565b820191906000526020600020905b81548152906001019060200180831161086d57829003601f168201915b505050505081565b600c602052600090815260409020805481906108ad9061249f565b80601f01602080910402602001604051908101604052809291908181526020018280546108d99061249f565b80156109265780601f106108fb57610100808354040283529160200191610926565b820191906000526020600020905b81548152906001019060200180831161090957829003601f168201915b5050506001909301549192505060ff1682565b6002600054141561095c5760405162461bcd60e51b815260040161059d9061233d565b60026000556001546205814914156109dc5760405162461bcd60e51b815260206004820152603e60248201527f544675656c546f6b656e42616e6b2e6275726e566f756368657273282920636160448201527f6e206f6e6c792062652063616c6c6564206f6e206120737562636861696e0000606482015260840161059d565b33346109e7816113b1565b6109f362058149611515565b6000610a33600154604051806040016040528060018152602001600360fc1b8152506040518060600160405280602a8152602001612574602a9139611534565b6205814960005260056020527fc18e43c37e0e3e2f9c59d63183efc764598410512ded25891c2225730180690f54604051919250907f40f1d475c2aa44f5c23193fab26a64d6aa4e09ab51898b10a3036baf82398ea190610a9d9084908790899088908790612293565b60405180910390a150506001600055505050565b600080610abd836110f5565b90506000600b82604051610ad19190612171565b908152604080516020928190038301812081830190925290546001600160a01b0381168252600160a01b900460ff16158015928201929092529150610b1857519392505050565b5060009392505050565b60026000541415610b455760405162461bcd60e51b815260040161059d9061233d565b60026000556001544614610bc3576040805162461bcd60e51b81526020600482015260248101919091527f544675656c546f6b656e42616e6b2e6c6f636b546f6b656e7328292063616e2060448201527f6f6e6c792062652063616c6c6564206f6e20746865206d61696e20636861696e606482015260840161059d565b6002546040516343b71f0560e01b8152600481018490526001600160a01b03909116906343b71f0590602401602060405180830381600087803b158015610c0957600080fd5b505af1158015610c1d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c419190611fef565b610c8d5760405162461bcd60e51b815260206004820181905260248201527f746172676574436861696e4944206e6f74207965742072656769737465726564604482015260640161059d565b60008281526010602052604090205434903390610caa9083611572565b600085815260106020526040902055610cc284611585565b6000610d02600154604051806040016040528060018152602001600360fc1b8152506040518060600160405280602a8152602001612574602a9139611534565b60008681526003602052604090819020549051919250907fee1ecc2b21aa613cc77cd44823a68ef1168ce1f40c2eac1d68690baf955fdbd190610d5090849086908a908a908a9088906122d3565b60405180910390a15050600160005550505050565b6001600160a01b0381166000908152600c60205260408082208151808301909252805460609392919082908290610d9b9061249f565b80601f0160208091040260200160405190810160405280929190818152602001828054610dc79061249f565b8015610e145780601f10610de957610100808354040283529160200191610e14565b820191906000526020600020905b815481529060010190602001808311610df757829003601f168201915b50505091835250506001919091015460ff16151560209182015281015190915015610e40575192915050565b5050604080516020810190915260008152919050565b50919050565b60026000541415610e7f5760405162461bcd60e51b815260040161059d9061233d565b60026000556001544614610f065760405162461bcd60e51b815260206004820152604260248201527f544675656c546f6b656e42616e6b2e756e6c6f636b546f6b656e73282920636160448201527f6e206f6e6c792062652063616c6c6564206f6e20746865206d61696e2063686160648201526134b760f11b608482015260a40161059d565b600085815260106020526040902054831115610f785760405162461bcd60e51b815260206004820152602b60248201527f43616e6e6f7420756e6c6f636b207468652072657175657374656420616d6f7560448201526a1b9d081bd98815119d595b60aa1b606482015260840161059d565b6040805160208101879052908101849052606085811b6bffffffffffffffffffffffff191690820152607481018390526094810182905260009060b40160408051601f19818403018152919052805160208201209091506000610fde88868487336115a4565b905080156110e657600088815260106020526040902054610fff90876115b7565b6000898152601060205260408082209290925590516001600160a01b0389169188156108fc02918991818181858888f19350505050158015611045573d6000803e3d6000fd5b5061104f886115c3565b600061108f600154604051806040016040528060018152602001600360fc1b8152506040518060600160405280602a8152602001612574602a9139611534565b60008a81526004602052604090819020549051919250907f5ea3a5ca7f54881fdd7781894d69709e11027910f35647f9d4cc14e6872b6f72906110db9084908c908c908b908790612253565b60405180910390a150505b50506001600055505050505050565b6060611100826115e2565b92915050565b60008060008061111c85602f60f81b600161165c565b9150915080611132575060009485945092505050565b600080611141876000866116ec565b91509150806111595750600096879650945050505050565b600080611165846117f9565b915091508061117f57506000988998509650505050505050565b50976001975095505050505050565b60006111a18686868686600760096118cb565b9695505050505050565b60408051603480825260608201909252600091602082018180368337019050509050816000805b601481101561123a578560601b81601481106111f0576111f061252f565b1a60f81b84836111ff816124d4565b9450815181106112115761121161252f565b60200101906001600160f81b031916908160001a90535080611232816124d4565b9150506111d2565b5060005b60208110156112a3578281602081106112595761125961252f565b1a60f81b8483611268816124d4565b94508151811061127a5761127a61252f565b60200101906001600160f81b031916908160001a9053508061129b816124d4565b91505061123e565b600060b66001600160a01b0316856040516112be9190612171565b6000604051808303816000865af19150503d80600081146112fb576040519150601f19603f3d011682016040523d82523d6000602084013e611300565b606091505b50509050806113815760405162461bcd60e51b815260206004820152604160248201527f544675656c546f6b656e42616e6b2e5f6d696e74544675656c566f756368657260448201527f28293a206661696c656420746f206d696e7420544675656c20766f75636865726064820152607360f81b608482015260a40161059d565b50505050505050565b60008181526006602052604081208054600192906113a99084906123c9565b909155505050565b6040805160208082528183019092526000916020820181803683370190505090508160005b6020811015611430578181602081106113f1576113f161252f565b1a60f81b8382815181106114075761140761252f565b60200101906001600160f81b031916908160001a90535080611428816124d4565b9150506113d6565b50600060b76001600160a01b03168360405161144c9190612171565b6000604051808303816000865af19150503d8060008114611489576040519150601f19603f3d011682016040523d82523d6000602084013e61148e565b606091505b505090508061150f5760405162461bcd60e51b815260206004820152604160248201527f544675656c546f6b656e42616e6b2e5f6275726e544675656c566f756368657260448201527f28293a206661696c656420746f206275726e20544675656c20766f75636865726064820152607360f81b608482015260a40161059d565b50505050565b60008181526005602052604081208054600192906113a99084906123c9565b606061156a61154285611ccb565b84846040516020016115569392919061218d565b6040516020818303038152906040526110f5565b949350505050565b600061157e82846123c9565b9392505050565b60008181526003602052604081208054600192906113a99084906123c9565b60006111a186868686866008600a6118cb565b600061157e8284612439565b60008181526004602052604081208054600192906113a99084906123c9565b60608160005b81518110156116555761161a8282815181106116065761160661252f565b01602001516001600160f81b031916611dc9565b82828151811061162c5761162c61252f565b60200101906001600160f81b031916908160001a9053508061164d816124d4565b9150506115e8565b5092915050565b82516000908190859082805b828110156116d857876001600160f81b03191684828151811061168d5761168d61252f565b01602001516001600160f81b03191614156116c6576116ad6001836123c9565b9150868214156116c6579450600193506116e492505050565b806116d0816124d4565b915050611668565b50600080945094505050505b935093915050565b825160609060009084841015611716575050604080516020810190915260008082529091506116e4565b80841115611738575050604080516020810190915260008082529091506116e4565b85600061174586886115b7565b67ffffffffffffffff81111561175d5761175d612545565b6040519080825280601f01601f191660200182016040528015611787576020820181803683370190505b509050865b8681101561117f578281815181106117a6576117a661252f565b01602001516001600160f81b031916826117c0838b6115b7565b815181106117d0576117d061252f565b60200101906001600160f81b031916908160001a905350806117f1816124d4565b91505061178c565b80516000908190839082805b828110156118be5760308482815181106118215761182161252f565b016020015160f81c10801590611851575060398482815181106118465761184661252f565b016020015160f81c11155b1561189d5761186182600a611e18565b9150611896603085838151811061187a5761187a61252f565b016020015161188c919060f81c612450565b839060ff16611572565b91506118ac565b50600096879650945050505050565b806118b6816124d4565b915050611805565b5095600195509350505050565b600060015488148061195657506002546040516343b71f0560e01b8152600481018a90526001600160a01b03909116906343b71f0590602401602060405180830381600087803b15801561191e57600080fd5b505af1158015611932573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119569190611fef565b6119945760405162461bcd60e51b815260206004820152600f60248201526e125b9d985b1a590818da185a5b9251608a1b604482015260640161059d565b6000888152602084905260409020546119ae9060016123c9565b85146119bc57506000611cc0565b600088815260208381526040808320898452909152812060015460609081908c908114156119ea5750620581495b6002546040516343f27e4560e01b815260048101839052602481018e90526001600160a01b03909116906343f27e4590604401600060405180830381600087803b158015611a3757600080fd5b505af1158015611a4b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611a739190810190611f23565b9350915060005b8251811015611bd057896001600160a01b0316838281518110611a9f57611a9f61252f565b60200260200101516001600160a01b031614611aba57611bbe565b6001955060005b6001860154811015611b5b57856001018181548110611ae257611ae261252f565b6000918252602090912001546001600160a01b038c811691161415611b495760405162461bcd60e51b815260206004820152601c60248201527f546869732076616c696461746f7220616c726561647920766f74656400000000604482015260640161059d565b80611b53816124d4565b915050611ac1565b508c85556001808601805491820181556000908152602090200180546001600160a01b031916331790558351611bb890859083908110611b9d57611b9d61252f565b6020026020010151866002015461157290919063ffffffff16565b60028601555b80611bc8816124d4565b915050611a7a565b50505082611c125760405162461bcd60e51b815260206004820152600f60248201526e2737ba1030903b30b634b230ba37b960891b604482015260640161059d565b6000805b8251811015611c5f57611c4b838281518110611c3457611c3461252f565b60200260200101518361157290919063ffffffff16565b915080611c57816124d4565b915050611c16565b50611c6b816002611e18565b6002840154611c7b906003611e18565b10611cb75760008c815260208890526040902054611c9a9060016123c9565b60008d8152602089905260409020555060019350611cc092505050565b60009450505050505b979650505050505050565b606081611cef5750506040805180820190915260018152600360fc1b602082015290565b8160005b8115611d195780611d03816124d4565b9150611d129050600a83612406565b9150611cf3565b60008167ffffffffffffffff811115611d3457611d34612545565b6040519080825280601f01601f191660200182016040528015611d5e576020820181803683370190505b5090505b841561156a57611d73600183612439565b9150611d80600a866124ef565b611d8b9060306123c9565b60f81b818381518110611da057611da061252f565b60200101906001600160f81b031916908160001a905350611dc2600a86612406565b9450611d62565b6000604160f81b6001600160f81b0319831610801590611df75750602d60f91b6001600160f81b0319831611155b15611e1457611e0b60f883901c60206123e1565b60f81b92915050565b5090565b600061157e828461241a565b600082601f830112611e3557600080fd5b81516020611e4a611e45836123a5565b612374565b80838252828201915082860187848660051b8901011115611e6a57600080fd5b60005b85811015611e8957815184529284019290840190600101611e6d565b5090979650505050505050565b600082601f830112611ea757600080fd5b813567ffffffffffffffff811115611ec157611ec1612545565b611ed4601f8201601f1916602001612374565b818152846020838601011115611ee957600080fd5b816020850160208301376000918101602001919091529392505050565b600060208284031215611f1857600080fd5b813561157e8161255b565b60008060408385031215611f3657600080fd5b825167ffffffffffffffff80821115611f4e57600080fd5b818501915085601f830112611f6257600080fd5b81516020611f72611e45836123a5565b8083825282820191508286018a848660051b8901011115611f9257600080fd5b600096505b84871015611fbe578051611faa8161255b565b835260019690960195918301918301611f97565b5091880151919650909350505080821115611fd857600080fd5b50611fe585828601611e24565b9150509250929050565b60006020828403121561200157600080fd5b8151801515811461157e57600080fd5b60006020828403121561202357600080fd5b813567ffffffffffffffff81111561203a57600080fd5b61156a84828501611e96565b600080600080600060a0868803121561205e57600080fd5b853567ffffffffffffffff81111561207557600080fd5b61208188828901611e96565b95505060208601356120928161255b565b94979496505050506040830135926060810135926080909101359150565b6000602082840312156120c257600080fd5b5035919050565b600080604083850312156120dc57600080fd5b8235915060208301356120ee8161255b565b809150509250929050565b600080600080600060a0868803121561211157600080fd5b8535945060208601356120928161255b565b6000806040838503121561213657600080fd5b50508035926020909101359150565b6000815180845261215d816020860160208601612473565b601f01601f19169290920160200192915050565b60008251612183818460208701612473565b9190910192915050565b6000845161219f818460208901612473565b8083019050602f60f81b80825285516121bf816001850160208a01612473565b600192019182015283516121da816002840160208801612473565b0160020195945050505050565b868152600086516121ff816020850160208b01612473565b60609690961b6bffffffffffffffffffffffff1916602092909601918201959095526034810193909352605483019190915260748201526094019392505050565b60208152600061157e6020830184612145565b60a08152600061226660a0830188612145565b6001600160a01b039690961660208301525060408101939093526060830191909152608090910152919050565b60a0815260006122a660a0830188612145565b6001600160a01b039687166020840152949095166040820152606081019290925260809091015292915050565b60c0815260006122e660c0830189612145565b6001600160a01b039788166020840152604083019690965250929094166060830152608082015260a00191909152919050565b60408152600061232c6040830185612145565b905082151560208301529392505050565b6020808252601f908201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604082015260600190565b604051601f8201601f1916810167ffffffffffffffff8111828210171561239d5761239d612545565b604052919050565b600067ffffffffffffffff8211156123bf576123bf612545565b5060051b60200190565b600082198211156123dc576123dc612503565b500190565b600060ff821660ff84168060ff038211156123fe576123fe612503565b019392505050565b60008261241557612415612519565b500490565b600081600019048311821515161561243457612434612503565b500290565b60008282101561244b5761244b612503565b500390565b600060ff821660ff84168082101561246a5761246a612503565b90039392505050565b60005b8381101561248e578181015183820152602001612476565b8381111561150f5750506000910152565b600181811c908216806124b357607f821691505b60208210811415610e5657634e487b7160e01b600052602260045260246000fd5b60006000198214156124e8576124e8612503565b5060010190565b6000826124fe576124fe612519565b500690565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052601260045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b038116811461257057600080fd5b5056fe307830303030303030303030303030303030303030303030303030303030303030303030303030303030a264697066735822122091273d601cd2dab4b98329887e79f115465a4c3f260792aa08516771476c89f964736f6c63430008070033",
}

// TFuelTokenBankABI is the input ABI used to generate the binding from.
// Deprecated: Use TFuelTokenBankMetaData.ABI instead.
var TFuelTokenBankABI = TFuelTokenBankMetaData.ABI

// TFuelTokenBankBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TFuelTokenBankMetaData.Bin instead.
var TFuelTokenBankBin = TFuelTokenBankMetaData.Bin

// DeployTFuelTokenBank deploys a new Ethereum contract, binding an instance of TFuelTokenBank to it.
func DeployTFuelTokenBank(auth *bind.TransactOpts, backend bind.ContractBackend, mainchainID_ *big.Int, chainRegistrar_ common.Address) (common.Address, *types.Transaction, *TFuelTokenBank, error) {
	parsed, err := TFuelTokenBankMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TFuelTokenBankBin), backend, mainchainID_, chainRegistrar_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TFuelTokenBank{TFuelTokenBankCaller: TFuelTokenBankCaller{contract: contract}, TFuelTokenBankTransactor: TFuelTokenBankTransactor{contract: contract}, TFuelTokenBankFilterer: TFuelTokenBankFilterer{contract: contract}}, nil
}

// TFuelTokenBank is an auto generated Go binding around an Ethereum contract.
type TFuelTokenBank struct {
	TFuelTokenBankCaller     // Read-only binding to the contract
	TFuelTokenBankTransactor // Write-only binding to the contract
	TFuelTokenBankFilterer   // Log filterer for contract events
}

// TFuelTokenBankCaller is an auto generated read-only Go binding around an Ethereum contract.
type TFuelTokenBankCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TFuelTokenBankTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TFuelTokenBankTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TFuelTokenBankFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TFuelTokenBankFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TFuelTokenBankSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TFuelTokenBankSession struct {
	Contract     *TFuelTokenBank   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TFuelTokenBankCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TFuelTokenBankCallerSession struct {
	Contract *TFuelTokenBankCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// TFuelTokenBankTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TFuelTokenBankTransactorSession struct {
	Contract     *TFuelTokenBankTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// TFuelTokenBankRaw is an auto generated low-level Go binding around an Ethereum contract.
type TFuelTokenBankRaw struct {
	Contract *TFuelTokenBank // Generic contract binding to access the raw methods on
}

// TFuelTokenBankCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TFuelTokenBankCallerRaw struct {
	Contract *TFuelTokenBankCaller // Generic read-only contract binding to access the raw methods on
}

// TFuelTokenBankTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TFuelTokenBankTransactorRaw struct {
	Contract *TFuelTokenBankTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTFuelTokenBank creates a new instance of TFuelTokenBank, bound to a specific deployed contract.
func NewTFuelTokenBank(address common.Address, backend bind.ContractBackend) (*TFuelTokenBank, error) {
	contract, err := bindTFuelTokenBank(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TFuelTokenBank{TFuelTokenBankCaller: TFuelTokenBankCaller{contract: contract}, TFuelTokenBankTransactor: TFuelTokenBankTransactor{contract: contract}, TFuelTokenBankFilterer: TFuelTokenBankFilterer{contract: contract}}, nil
}

// NewTFuelTokenBankCaller creates a new read-only instance of TFuelTokenBank, bound to a specific deployed contract.
func NewTFuelTokenBankCaller(address common.Address, caller bind.ContractCaller) (*TFuelTokenBankCaller, error) {
	contract, err := bindTFuelTokenBank(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TFuelTokenBankCaller{contract: contract}, nil
}

// NewTFuelTokenBankTransactor creates a new write-only instance of TFuelTokenBank, bound to a specific deployed contract.
func NewTFuelTokenBankTransactor(address common.Address, transactor bind.ContractTransactor) (*TFuelTokenBankTransactor, error) {
	contract, err := bindTFuelTokenBank(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TFuelTokenBankTransactor{contract: contract}, nil
}

// NewTFuelTokenBankFilterer creates a new log filterer instance of TFuelTokenBank, bound to a specific deployed contract.
func NewTFuelTokenBankFilterer(address common.Address, filterer bind.ContractFilterer) (*TFuelTokenBankFilterer, error) {
	contract, err := bindTFuelTokenBank(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TFuelTokenBankFilterer{contract: contract}, nil
}

// bindTFuelTokenBank binds a generic wrapper to an already deployed contract.
func bindTFuelTokenBank(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TFuelTokenBankABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TFuelTokenBank *TFuelTokenBankRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TFuelTokenBank.Contract.TFuelTokenBankCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TFuelTokenBank *TFuelTokenBankRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TFuelTokenBank.Contract.TFuelTokenBankTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TFuelTokenBank *TFuelTokenBankRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TFuelTokenBank.Contract.TFuelTokenBankTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TFuelTokenBank *TFuelTokenBankCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TFuelTokenBank.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TFuelTokenBank *TFuelTokenBankTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TFuelTokenBank.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TFuelTokenBank *TFuelTokenBankTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TFuelTokenBank.Contract.contract.Transact(opts, method, params...)
}

// AllDenoms is a free data retrieval call binding the contract method 0x588b1408.
//
// Solidity: function allDenoms(uint256 ) view returns(string)
func (_TFuelTokenBank *TFuelTokenBankCaller) AllDenoms(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _TFuelTokenBank.contract.Call(opts, &out, "allDenoms", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AllDenoms is a free data retrieval call binding the contract method 0x588b1408.
//
// Solidity: function allDenoms(uint256 ) view returns(string)
func (_TFuelTokenBank *TFuelTokenBankSession) AllDenoms(arg0 *big.Int) (string, error) {
	return _TFuelTokenBank.Contract.AllDenoms(&_TFuelTokenBank.CallOpts, arg0)
}

// AllDenoms is a free data retrieval call binding the contract method 0x588b1408.
//
// Solidity: function allDenoms(uint256 ) view returns(string)
func (_TFuelTokenBank *TFuelTokenBankCallerSession) AllDenoms(arg0 *big.Int) (string, error) {
	return _TFuelTokenBank.Contract.AllDenoms(&_TFuelTokenBank.CallOpts, arg0)
}

// AllVouchers is a free data retrieval call binding the contract method 0x27ca4df1.
//
// Solidity: function allVouchers(uint256 ) view returns(address)
func (_TFuelTokenBank *TFuelTokenBankCaller) AllVouchers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TFuelTokenBank.contract.Call(opts, &out, "allVouchers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllVouchers is a free data retrieval call binding the contract method 0x27ca4df1.
//
// Solidity: function allVouchers(uint256 ) view returns(address)
func (_TFuelTokenBank *TFuelTokenBankSession) AllVouchers(arg0 *big.Int) (common.Address, error) {
	return _TFuelTokenBank.Contract.AllVouchers(&_TFuelTokenBank.CallOpts, arg0)
}

// AllVouchers is a free data retrieval call binding the contract method 0x27ca4df1.
//
// Solidity: function allVouchers(uint256 ) view returns(address)
func (_TFuelTokenBank *TFuelTokenBankCallerSession) AllVouchers(arg0 *big.Int) (common.Address, error) {
	return _TFuelTokenBank.Contract.AllVouchers(&_TFuelTokenBank.CallOpts, arg0)
}

// DenomToVoucherLookup is a free data retrieval call binding the contract method 0x1527b14d.
//
// Solidity: function denomToVoucherLookup(string ) view returns(address contractAddress, bool exists)
func (_TFuelTokenBank *TFuelTokenBankCaller) DenomToVoucherLookup(opts *bind.CallOpts, arg0 string) (struct {
	ContractAddress common.Address
	Exists          bool
}, error) {
	var out []interface{}
	err := _TFuelTokenBank.contract.Call(opts, &out, "denomToVoucherLookup", arg0)

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
func (_TFuelTokenBank *TFuelTokenBankSession) DenomToVoucherLookup(arg0 string) (struct {
	ContractAddress common.Address
	Exists          bool
}, error) {
	return _TFuelTokenBank.Contract.DenomToVoucherLookup(&_TFuelTokenBank.CallOpts, arg0)
}

// DenomToVoucherLookup is a free data retrieval call binding the contract method 0x1527b14d.
//
// Solidity: function denomToVoucherLookup(string ) view returns(address contractAddress, bool exists)
func (_TFuelTokenBank *TFuelTokenBankCallerSession) DenomToVoucherLookup(arg0 string) (struct {
	ContractAddress common.Address
	Exists          bool
}, error) {
	return _TFuelTokenBank.Contract.DenomToVoucherLookup(&_TFuelTokenBank.CallOpts, arg0)
}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(string denom) view returns(bool)
func (_TFuelTokenBank *TFuelTokenBankCaller) Exists(opts *bind.CallOpts, denom string) (bool, error) {
	var out []interface{}
	err := _TFuelTokenBank.contract.Call(opts, &out, "exists", denom)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(string denom) view returns(bool)
func (_TFuelTokenBank *TFuelTokenBankSession) Exists(denom string) (bool, error) {
	return _TFuelTokenBank.Contract.Exists(&_TFuelTokenBank.CallOpts, denom)
}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(string denom) view returns(bool)
func (_TFuelTokenBank *TFuelTokenBankCallerSession) Exists(denom string) (bool, error) {
	return _TFuelTokenBank.Contract.Exists(&_TFuelTokenBank.CallOpts, denom)
}

// Exists0 is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address voucherAddress) view returns(bool)
func (_TFuelTokenBank *TFuelTokenBankCaller) Exists0(opts *bind.CallOpts, voucherAddress common.Address) (bool, error) {
	var out []interface{}
	err := _TFuelTokenBank.contract.Call(opts, &out, "exists0", voucherAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists0 is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address voucherAddress) view returns(bool)
func (_TFuelTokenBank *TFuelTokenBankSession) Exists0(voucherAddress common.Address) (bool, error) {
	return _TFuelTokenBank.Contract.Exists0(&_TFuelTokenBank.CallOpts, voucherAddress)
}

// Exists0 is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address voucherAddress) view returns(bool)
func (_TFuelTokenBank *TFuelTokenBankCallerSession) Exists0(voucherAddress common.Address) (bool, error) {
	return _TFuelTokenBank.Contract.Exists0(&_TFuelTokenBank.CallOpts, voucherAddress)
}

// GetDenom is a free data retrieval call binding the contract method 0xebda9962.
//
// Solidity: function getDenom(address voucherContractAddr) view returns(string)
func (_TFuelTokenBank *TFuelTokenBankCaller) GetDenom(opts *bind.CallOpts, voucherContractAddr common.Address) (string, error) {
	var out []interface{}
	err := _TFuelTokenBank.contract.Call(opts, &out, "getDenom", voucherContractAddr)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetDenom is a free data retrieval call binding the contract method 0xebda9962.
//
// Solidity: function getDenom(address voucherContractAddr) view returns(string)
func (_TFuelTokenBank *TFuelTokenBankSession) GetDenom(voucherContractAddr common.Address) (string, error) {
	return _TFuelTokenBank.Contract.GetDenom(&_TFuelTokenBank.CallOpts, voucherContractAddr)
}

// GetDenom is a free data retrieval call binding the contract method 0xebda9962.
//
// Solidity: function getDenom(address voucherContractAddr) view returns(string)
func (_TFuelTokenBank *TFuelTokenBankCallerSession) GetDenom(voucherContractAddr common.Address) (string, error) {
	return _TFuelTokenBank.Contract.GetDenom(&_TFuelTokenBank.CallOpts, voucherContractAddr)
}

// GetMaxProcessedTokenLockNonce is a free data retrieval call binding the contract method 0xf95627ac.
//
// Solidity: function getMaxProcessedTokenLockNonce(uint256 chainID) view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankCaller) GetMaxProcessedTokenLockNonce(opts *bind.CallOpts, chainID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TFuelTokenBank.contract.Call(opts, &out, "getMaxProcessedTokenLockNonce", chainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxProcessedTokenLockNonce is a free data retrieval call binding the contract method 0xf95627ac.
//
// Solidity: function getMaxProcessedTokenLockNonce(uint256 chainID) view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankSession) GetMaxProcessedTokenLockNonce(chainID *big.Int) (*big.Int, error) {
	return _TFuelTokenBank.Contract.GetMaxProcessedTokenLockNonce(&_TFuelTokenBank.CallOpts, chainID)
}

// GetMaxProcessedTokenLockNonce is a free data retrieval call binding the contract method 0xf95627ac.
//
// Solidity: function getMaxProcessedTokenLockNonce(uint256 chainID) view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankCallerSession) GetMaxProcessedTokenLockNonce(chainID *big.Int) (*big.Int, error) {
	return _TFuelTokenBank.Contract.GetMaxProcessedTokenLockNonce(&_TFuelTokenBank.CallOpts, chainID)
}

// GetMaxProcessedVoucherBurnNonce is a free data retrieval call binding the contract method 0x766f8fb0.
//
// Solidity: function getMaxProcessedVoucherBurnNonce(uint256 chainID) view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankCaller) GetMaxProcessedVoucherBurnNonce(opts *bind.CallOpts, chainID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TFuelTokenBank.contract.Call(opts, &out, "getMaxProcessedVoucherBurnNonce", chainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxProcessedVoucherBurnNonce is a free data retrieval call binding the contract method 0x766f8fb0.
//
// Solidity: function getMaxProcessedVoucherBurnNonce(uint256 chainID) view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankSession) GetMaxProcessedVoucherBurnNonce(chainID *big.Int) (*big.Int, error) {
	return _TFuelTokenBank.Contract.GetMaxProcessedVoucherBurnNonce(&_TFuelTokenBank.CallOpts, chainID)
}

// GetMaxProcessedVoucherBurnNonce is a free data retrieval call binding the contract method 0x766f8fb0.
//
// Solidity: function getMaxProcessedVoucherBurnNonce(uint256 chainID) view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankCallerSession) GetMaxProcessedVoucherBurnNonce(chainID *big.Int) (*big.Int, error) {
	return _TFuelTokenBank.Contract.GetMaxProcessedVoucherBurnNonce(&_TFuelTokenBank.CallOpts, chainID)
}

// GetVoucher is a free data retrieval call binding the contract method 0xa2cc6981.
//
// Solidity: function getVoucher(string denom) view returns(address)
func (_TFuelTokenBank *TFuelTokenBankCaller) GetVoucher(opts *bind.CallOpts, denom string) (common.Address, error) {
	var out []interface{}
	err := _TFuelTokenBank.contract.Call(opts, &out, "getVoucher", denom)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVoucher is a free data retrieval call binding the contract method 0xa2cc6981.
//
// Solidity: function getVoucher(string denom) view returns(address)
func (_TFuelTokenBank *TFuelTokenBankSession) GetVoucher(denom string) (common.Address, error) {
	return _TFuelTokenBank.Contract.GetVoucher(&_TFuelTokenBank.CallOpts, denom)
}

// GetVoucher is a free data retrieval call binding the contract method 0xa2cc6981.
//
// Solidity: function getVoucher(string denom) view returns(address)
func (_TFuelTokenBank *TFuelTokenBankCallerSession) GetVoucher(denom string) (common.Address, error) {
	return _TFuelTokenBank.Contract.GetVoucher(&_TFuelTokenBank.CallOpts, denom)
}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankCaller) Id(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TFuelTokenBank.contract.Call(opts, &out, "id")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankSession) Id() (*big.Int, error) {
	return _TFuelTokenBank.Contract.Id(&_TFuelTokenBank.CallOpts)
}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankCallerSession) Id() (*big.Int, error) {
	return _TFuelTokenBank.Contract.Id(&_TFuelTokenBank.CallOpts)
}

// IsOnMainchain is a free data retrieval call binding the contract method 0x4250863b.
//
// Solidity: function isOnMainchain() view returns(bool)
func (_TFuelTokenBank *TFuelTokenBankCaller) IsOnMainchain(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TFuelTokenBank.contract.Call(opts, &out, "isOnMainchain")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOnMainchain is a free data retrieval call binding the contract method 0x4250863b.
//
// Solidity: function isOnMainchain() view returns(bool)
func (_TFuelTokenBank *TFuelTokenBankSession) IsOnMainchain() (bool, error) {
	return _TFuelTokenBank.Contract.IsOnMainchain(&_TFuelTokenBank.CallOpts)
}

// IsOnMainchain is a free data retrieval call binding the contract method 0x4250863b.
//
// Solidity: function isOnMainchain() view returns(bool)
func (_TFuelTokenBank *TFuelTokenBankCallerSession) IsOnMainchain() (bool, error) {
	return _TFuelTokenBank.Contract.IsOnMainchain(&_TFuelTokenBank.CallOpts)
}

// MainchainID is a free data retrieval call binding the contract method 0x073b9502.
//
// Solidity: function mainchainID() view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankCaller) MainchainID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TFuelTokenBank.contract.Call(opts, &out, "mainchainID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MainchainID is a free data retrieval call binding the contract method 0x073b9502.
//
// Solidity: function mainchainID() view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankSession) MainchainID() (*big.Int, error) {
	return _TFuelTokenBank.Contract.MainchainID(&_TFuelTokenBank.CallOpts)
}

// MainchainID is a free data retrieval call binding the contract method 0x073b9502.
//
// Solidity: function mainchainID() view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankCallerSession) MainchainID() (*big.Int, error) {
	return _TFuelTokenBank.Contract.MainchainID(&_TFuelTokenBank.CallOpts)
}

// TokenLockNonceMap is a free data retrieval call binding the contract method 0x8883931e.
//
// Solidity: function tokenLockNonceMap(uint256 ) view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankCaller) TokenLockNonceMap(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TFuelTokenBank.contract.Call(opts, &out, "tokenLockNonceMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenLockNonceMap is a free data retrieval call binding the contract method 0x8883931e.
//
// Solidity: function tokenLockNonceMap(uint256 ) view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankSession) TokenLockNonceMap(arg0 *big.Int) (*big.Int, error) {
	return _TFuelTokenBank.Contract.TokenLockNonceMap(&_TFuelTokenBank.CallOpts, arg0)
}

// TokenLockNonceMap is a free data retrieval call binding the contract method 0x8883931e.
//
// Solidity: function tokenLockNonceMap(uint256 ) view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankCallerSession) TokenLockNonceMap(arg0 *big.Int) (*big.Int, error) {
	return _TFuelTokenBank.Contract.TokenLockNonceMap(&_TFuelTokenBank.CallOpts, arg0)
}

// TokenLockVotingRecords is a free data retrieval call binding the contract method 0x1eb78737.
//
// Solidity: function tokenLockVotingRecords(uint256 , bytes32 ) view returns(uint256 dynasty, uint256 accumlatedShares)
func (_TFuelTokenBank *TFuelTokenBankCaller) TokenLockVotingRecords(opts *bind.CallOpts, arg0 *big.Int, arg1 [32]byte) (struct {
	Dynasty          *big.Int
	AccumlatedShares *big.Int
}, error) {
	var out []interface{}
	err := _TFuelTokenBank.contract.Call(opts, &out, "tokenLockVotingRecords", arg0, arg1)

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
func (_TFuelTokenBank *TFuelTokenBankSession) TokenLockVotingRecords(arg0 *big.Int, arg1 [32]byte) (struct {
	Dynasty          *big.Int
	AccumlatedShares *big.Int
}, error) {
	return _TFuelTokenBank.Contract.TokenLockVotingRecords(&_TFuelTokenBank.CallOpts, arg0, arg1)
}

// TokenLockVotingRecords is a free data retrieval call binding the contract method 0x1eb78737.
//
// Solidity: function tokenLockVotingRecords(uint256 , bytes32 ) view returns(uint256 dynasty, uint256 accumlatedShares)
func (_TFuelTokenBank *TFuelTokenBankCallerSession) TokenLockVotingRecords(arg0 *big.Int, arg1 [32]byte) (struct {
	Dynasty          *big.Int
	AccumlatedShares *big.Int
}, error) {
	return _TFuelTokenBank.Contract.TokenLockVotingRecords(&_TFuelTokenBank.CallOpts, arg0, arg1)
}

// TokenUnlockNonceMap is a free data retrieval call binding the contract method 0xccf187c7.
//
// Solidity: function tokenUnlockNonceMap(uint256 ) view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankCaller) TokenUnlockNonceMap(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TFuelTokenBank.contract.Call(opts, &out, "tokenUnlockNonceMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenUnlockNonceMap is a free data retrieval call binding the contract method 0xccf187c7.
//
// Solidity: function tokenUnlockNonceMap(uint256 ) view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankSession) TokenUnlockNonceMap(arg0 *big.Int) (*big.Int, error) {
	return _TFuelTokenBank.Contract.TokenUnlockNonceMap(&_TFuelTokenBank.CallOpts, arg0)
}

// TokenUnlockNonceMap is a free data retrieval call binding the contract method 0xccf187c7.
//
// Solidity: function tokenUnlockNonceMap(uint256 ) view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankCallerSession) TokenUnlockNonceMap(arg0 *big.Int) (*big.Int, error) {
	return _TFuelTokenBank.Contract.TokenUnlockNonceMap(&_TFuelTokenBank.CallOpts, arg0)
}

// TotalLockedAmounts is a free data retrieval call binding the contract method 0x19fd1a11.
//
// Solidity: function totalLockedAmounts(uint256 ) view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankCaller) TotalLockedAmounts(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TFuelTokenBank.contract.Call(opts, &out, "totalLockedAmounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalLockedAmounts is a free data retrieval call binding the contract method 0x19fd1a11.
//
// Solidity: function totalLockedAmounts(uint256 ) view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankSession) TotalLockedAmounts(arg0 *big.Int) (*big.Int, error) {
	return _TFuelTokenBank.Contract.TotalLockedAmounts(&_TFuelTokenBank.CallOpts, arg0)
}

// TotalLockedAmounts is a free data retrieval call binding the contract method 0x19fd1a11.
//
// Solidity: function totalLockedAmounts(uint256 ) view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankCallerSession) TotalLockedAmounts(arg0 *big.Int) (*big.Int, error) {
	return _TFuelTokenBank.Contract.TotalLockedAmounts(&_TFuelTokenBank.CallOpts, arg0)
}

// VoucherAddressToDenomLookup is a free data retrieval call binding the contract method 0x60569b5e.
//
// Solidity: function voucherAddressToDenomLookup(address ) view returns(string denom, bool exists)
func (_TFuelTokenBank *TFuelTokenBankCaller) VoucherAddressToDenomLookup(opts *bind.CallOpts, arg0 common.Address) (struct {
	Denom  string
	Exists bool
}, error) {
	var out []interface{}
	err := _TFuelTokenBank.contract.Call(opts, &out, "voucherAddressToDenomLookup", arg0)

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
func (_TFuelTokenBank *TFuelTokenBankSession) VoucherAddressToDenomLookup(arg0 common.Address) (struct {
	Denom  string
	Exists bool
}, error) {
	return _TFuelTokenBank.Contract.VoucherAddressToDenomLookup(&_TFuelTokenBank.CallOpts, arg0)
}

// VoucherAddressToDenomLookup is a free data retrieval call binding the contract method 0x60569b5e.
//
// Solidity: function voucherAddressToDenomLookup(address ) view returns(string denom, bool exists)
func (_TFuelTokenBank *TFuelTokenBankCallerSession) VoucherAddressToDenomLookup(arg0 common.Address) (struct {
	Denom  string
	Exists bool
}, error) {
	return _TFuelTokenBank.Contract.VoucherAddressToDenomLookup(&_TFuelTokenBank.CallOpts, arg0)
}

// VoucherBurnNonceMap is a free data retrieval call binding the contract method 0xca207569.
//
// Solidity: function voucherBurnNonceMap(uint256 ) view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankCaller) VoucherBurnNonceMap(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TFuelTokenBank.contract.Call(opts, &out, "voucherBurnNonceMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VoucherBurnNonceMap is a free data retrieval call binding the contract method 0xca207569.
//
// Solidity: function voucherBurnNonceMap(uint256 ) view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankSession) VoucherBurnNonceMap(arg0 *big.Int) (*big.Int, error) {
	return _TFuelTokenBank.Contract.VoucherBurnNonceMap(&_TFuelTokenBank.CallOpts, arg0)
}

// VoucherBurnNonceMap is a free data retrieval call binding the contract method 0xca207569.
//
// Solidity: function voucherBurnNonceMap(uint256 ) view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankCallerSession) VoucherBurnNonceMap(arg0 *big.Int) (*big.Int, error) {
	return _TFuelTokenBank.Contract.VoucherBurnNonceMap(&_TFuelTokenBank.CallOpts, arg0)
}

// VoucherBurnVotingRecords is a free data retrieval call binding the contract method 0xfeaff052.
//
// Solidity: function voucherBurnVotingRecords(uint256 , bytes32 ) view returns(uint256 dynasty, uint256 accumlatedShares)
func (_TFuelTokenBank *TFuelTokenBankCaller) VoucherBurnVotingRecords(opts *bind.CallOpts, arg0 *big.Int, arg1 [32]byte) (struct {
	Dynasty          *big.Int
	AccumlatedShares *big.Int
}, error) {
	var out []interface{}
	err := _TFuelTokenBank.contract.Call(opts, &out, "voucherBurnVotingRecords", arg0, arg1)

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
func (_TFuelTokenBank *TFuelTokenBankSession) VoucherBurnVotingRecords(arg0 *big.Int, arg1 [32]byte) (struct {
	Dynasty          *big.Int
	AccumlatedShares *big.Int
}, error) {
	return _TFuelTokenBank.Contract.VoucherBurnVotingRecords(&_TFuelTokenBank.CallOpts, arg0, arg1)
}

// VoucherBurnVotingRecords is a free data retrieval call binding the contract method 0xfeaff052.
//
// Solidity: function voucherBurnVotingRecords(uint256 , bytes32 ) view returns(uint256 dynasty, uint256 accumlatedShares)
func (_TFuelTokenBank *TFuelTokenBankCallerSession) VoucherBurnVotingRecords(arg0 *big.Int, arg1 [32]byte) (struct {
	Dynasty          *big.Int
	AccumlatedShares *big.Int
}, error) {
	return _TFuelTokenBank.Contract.VoucherBurnVotingRecords(&_TFuelTokenBank.CallOpts, arg0, arg1)
}

// VoucherMintNonceMap is a free data retrieval call binding the contract method 0x740cb7f8.
//
// Solidity: function voucherMintNonceMap(uint256 ) view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankCaller) VoucherMintNonceMap(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TFuelTokenBank.contract.Call(opts, &out, "voucherMintNonceMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VoucherMintNonceMap is a free data retrieval call binding the contract method 0x740cb7f8.
//
// Solidity: function voucherMintNonceMap(uint256 ) view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankSession) VoucherMintNonceMap(arg0 *big.Int) (*big.Int, error) {
	return _TFuelTokenBank.Contract.VoucherMintNonceMap(&_TFuelTokenBank.CallOpts, arg0)
}

// VoucherMintNonceMap is a free data retrieval call binding the contract method 0x740cb7f8.
//
// Solidity: function voucherMintNonceMap(uint256 ) view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankCallerSession) VoucherMintNonceMap(arg0 *big.Int) (*big.Int, error) {
	return _TFuelTokenBank.Contract.VoucherMintNonceMap(&_TFuelTokenBank.CallOpts, arg0)
}

// BurnVouchers is a paid mutator transaction binding the contract method 0x7d0fb00d.
//
// Solidity: function burnVouchers(address targetChainTokenReceiver) payable returns()
func (_TFuelTokenBank *TFuelTokenBankTransactor) BurnVouchers(opts *bind.TransactOpts, targetChainTokenReceiver common.Address) (*types.Transaction, error) {
	return _TFuelTokenBank.contract.Transact(opts, "burnVouchers", targetChainTokenReceiver)
}

// BurnVouchers is a paid mutator transaction binding the contract method 0x7d0fb00d.
//
// Solidity: function burnVouchers(address targetChainTokenReceiver) payable returns()
func (_TFuelTokenBank *TFuelTokenBankSession) BurnVouchers(targetChainTokenReceiver common.Address) (*types.Transaction, error) {
	return _TFuelTokenBank.Contract.BurnVouchers(&_TFuelTokenBank.TransactOpts, targetChainTokenReceiver)
}

// BurnVouchers is a paid mutator transaction binding the contract method 0x7d0fb00d.
//
// Solidity: function burnVouchers(address targetChainTokenReceiver) payable returns()
func (_TFuelTokenBank *TFuelTokenBankTransactorSession) BurnVouchers(targetChainTokenReceiver common.Address) (*types.Transaction, error) {
	return _TFuelTokenBank.Contract.BurnVouchers(&_TFuelTokenBank.TransactOpts, targetChainTokenReceiver)
}

// LockTokens is a paid mutator transaction binding the contract method 0xaa68acde.
//
// Solidity: function lockTokens(uint256 targetChainID, address targetChainVoucherReceiver) payable returns()
func (_TFuelTokenBank *TFuelTokenBankTransactor) LockTokens(opts *bind.TransactOpts, targetChainID *big.Int, targetChainVoucherReceiver common.Address) (*types.Transaction, error) {
	return _TFuelTokenBank.contract.Transact(opts, "lockTokens", targetChainID, targetChainVoucherReceiver)
}

// LockTokens is a paid mutator transaction binding the contract method 0xaa68acde.
//
// Solidity: function lockTokens(uint256 targetChainID, address targetChainVoucherReceiver) payable returns()
func (_TFuelTokenBank *TFuelTokenBankSession) LockTokens(targetChainID *big.Int, targetChainVoucherReceiver common.Address) (*types.Transaction, error) {
	return _TFuelTokenBank.Contract.LockTokens(&_TFuelTokenBank.TransactOpts, targetChainID, targetChainVoucherReceiver)
}

// LockTokens is a paid mutator transaction binding the contract method 0xaa68acde.
//
// Solidity: function lockTokens(uint256 targetChainID, address targetChainVoucherReceiver) payable returns()
func (_TFuelTokenBank *TFuelTokenBankTransactorSession) LockTokens(targetChainID *big.Int, targetChainVoucherReceiver common.Address) (*types.Transaction, error) {
	return _TFuelTokenBank.Contract.LockTokens(&_TFuelTokenBank.TransactOpts, targetChainID, targetChainVoucherReceiver)
}

// MintVouchers is a paid mutator transaction binding the contract method 0x1a0483d3.
//
// Solidity: function mintVouchers(string denom, address targetChainVoucherReceiver, uint256 mintedAmount, uint256 dynasty, uint256 sourceChainTokenLockNonce) returns()
func (_TFuelTokenBank *TFuelTokenBankTransactor) MintVouchers(opts *bind.TransactOpts, denom string, targetChainVoucherReceiver common.Address, mintedAmount *big.Int, dynasty *big.Int, sourceChainTokenLockNonce *big.Int) (*types.Transaction, error) {
	return _TFuelTokenBank.contract.Transact(opts, "mintVouchers", denom, targetChainVoucherReceiver, mintedAmount, dynasty, sourceChainTokenLockNonce)
}

// MintVouchers is a paid mutator transaction binding the contract method 0x1a0483d3.
//
// Solidity: function mintVouchers(string denom, address targetChainVoucherReceiver, uint256 mintedAmount, uint256 dynasty, uint256 sourceChainTokenLockNonce) returns()
func (_TFuelTokenBank *TFuelTokenBankSession) MintVouchers(denom string, targetChainVoucherReceiver common.Address, mintedAmount *big.Int, dynasty *big.Int, sourceChainTokenLockNonce *big.Int) (*types.Transaction, error) {
	return _TFuelTokenBank.Contract.MintVouchers(&_TFuelTokenBank.TransactOpts, denom, targetChainVoucherReceiver, mintedAmount, dynasty, sourceChainTokenLockNonce)
}

// MintVouchers is a paid mutator transaction binding the contract method 0x1a0483d3.
//
// Solidity: function mintVouchers(string denom, address targetChainVoucherReceiver, uint256 mintedAmount, uint256 dynasty, uint256 sourceChainTokenLockNonce) returns()
func (_TFuelTokenBank *TFuelTokenBankTransactorSession) MintVouchers(denom string, targetChainVoucherReceiver common.Address, mintedAmount *big.Int, dynasty *big.Int, sourceChainTokenLockNonce *big.Int) (*types.Transaction, error) {
	return _TFuelTokenBank.Contract.MintVouchers(&_TFuelTokenBank.TransactOpts, denom, targetChainVoucherReceiver, mintedAmount, dynasty, sourceChainTokenLockNonce)
}

// UnlockTokens is a paid mutator transaction binding the contract method 0xff248a44.
//
// Solidity: function unlockTokens(uint256 sourceChainID, address targetChainTokenReceiver, uint256 unlockAmount, uint256 dynasty, uint256 sourceChainVoucherBurnNonce) payable returns()
func (_TFuelTokenBank *TFuelTokenBankTransactor) UnlockTokens(opts *bind.TransactOpts, sourceChainID *big.Int, targetChainTokenReceiver common.Address, unlockAmount *big.Int, dynasty *big.Int, sourceChainVoucherBurnNonce *big.Int) (*types.Transaction, error) {
	return _TFuelTokenBank.contract.Transact(opts, "unlockTokens", sourceChainID, targetChainTokenReceiver, unlockAmount, dynasty, sourceChainVoucherBurnNonce)
}

// UnlockTokens is a paid mutator transaction binding the contract method 0xff248a44.
//
// Solidity: function unlockTokens(uint256 sourceChainID, address targetChainTokenReceiver, uint256 unlockAmount, uint256 dynasty, uint256 sourceChainVoucherBurnNonce) payable returns()
func (_TFuelTokenBank *TFuelTokenBankSession) UnlockTokens(sourceChainID *big.Int, targetChainTokenReceiver common.Address, unlockAmount *big.Int, dynasty *big.Int, sourceChainVoucherBurnNonce *big.Int) (*types.Transaction, error) {
	return _TFuelTokenBank.Contract.UnlockTokens(&_TFuelTokenBank.TransactOpts, sourceChainID, targetChainTokenReceiver, unlockAmount, dynasty, sourceChainVoucherBurnNonce)
}

// UnlockTokens is a paid mutator transaction binding the contract method 0xff248a44.
//
// Solidity: function unlockTokens(uint256 sourceChainID, address targetChainTokenReceiver, uint256 unlockAmount, uint256 dynasty, uint256 sourceChainVoucherBurnNonce) payable returns()
func (_TFuelTokenBank *TFuelTokenBankTransactorSession) UnlockTokens(sourceChainID *big.Int, targetChainTokenReceiver common.Address, unlockAmount *big.Int, dynasty *big.Int, sourceChainVoucherBurnNonce *big.Int) (*types.Transaction, error) {
	return _TFuelTokenBank.Contract.UnlockTokens(&_TFuelTokenBank.TransactOpts, sourceChainID, targetChainTokenReceiver, unlockAmount, dynasty, sourceChainVoucherBurnNonce)
}

// TFuelTokenBankTFuelTokenLockedIterator is returned from FilterTFuelTokenLocked and is used to iterate over the raw logs and unpacked data for TFuelTokenLocked events raised by the TFuelTokenBank contract.
type TFuelTokenBankTFuelTokenLockedIterator struct {
	Event *TFuelTokenBankTFuelTokenLocked // Event containing the contract specifics and raw log

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
func (it *TFuelTokenBankTFuelTokenLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TFuelTokenBankTFuelTokenLocked)
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
		it.Event = new(TFuelTokenBankTFuelTokenLocked)
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
func (it *TFuelTokenBankTFuelTokenLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TFuelTokenBankTFuelTokenLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TFuelTokenBankTFuelTokenLocked represents a TFuelTokenLocked event raised by the TFuelTokenBank contract.
type TFuelTokenBankTFuelTokenLocked struct {
	Denom                      string
	SourceChainTokenSender     common.Address
	TargetChainID              *big.Int
	TargetChainVoucherReceiver common.Address
	LockedAmount               *big.Int
	TokenLockNonce             *big.Int
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterTFuelTokenLocked is a free log retrieval operation binding the contract event 0xee1ecc2b21aa613cc77cd44823a68ef1168ce1f40c2eac1d68690baf955fdbd1.
//
// Solidity: event TFuelTokenLocked(string denom, address sourceChainTokenSender, uint256 targetChainID, address targetChainVoucherReceiver, uint256 lockedAmount, uint256 tokenLockNonce)
func (_TFuelTokenBank *TFuelTokenBankFilterer) FilterTFuelTokenLocked(opts *bind.FilterOpts) (*TFuelTokenBankTFuelTokenLockedIterator, error) {

	logs, sub, err := _TFuelTokenBank.contract.FilterLogs(opts, "TFuelTokenLocked")
	if err != nil {
		return nil, err
	}
	return &TFuelTokenBankTFuelTokenLockedIterator{contract: _TFuelTokenBank.contract, event: "TFuelTokenLocked", logs: logs, sub: sub}, nil
}

// WatchTFuelTokenLocked is a free log subscription operation binding the contract event 0xee1ecc2b21aa613cc77cd44823a68ef1168ce1f40c2eac1d68690baf955fdbd1.
//
// Solidity: event TFuelTokenLocked(string denom, address sourceChainTokenSender, uint256 targetChainID, address targetChainVoucherReceiver, uint256 lockedAmount, uint256 tokenLockNonce)
func (_TFuelTokenBank *TFuelTokenBankFilterer) WatchTFuelTokenLocked(opts *bind.WatchOpts, sink chan<- *TFuelTokenBankTFuelTokenLocked) (event.Subscription, error) {

	logs, sub, err := _TFuelTokenBank.contract.WatchLogs(opts, "TFuelTokenLocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TFuelTokenBankTFuelTokenLocked)
				if err := _TFuelTokenBank.contract.UnpackLog(event, "TFuelTokenLocked", log); err != nil {
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

// ParseTFuelTokenLocked is a log parse operation binding the contract event 0xee1ecc2b21aa613cc77cd44823a68ef1168ce1f40c2eac1d68690baf955fdbd1.
//
// Solidity: event TFuelTokenLocked(string denom, address sourceChainTokenSender, uint256 targetChainID, address targetChainVoucherReceiver, uint256 lockedAmount, uint256 tokenLockNonce)
func (_TFuelTokenBank *TFuelTokenBankFilterer) ParseTFuelTokenLocked(log types.Log) (*TFuelTokenBankTFuelTokenLocked, error) {
	event := new(TFuelTokenBankTFuelTokenLocked)
	if err := _TFuelTokenBank.contract.UnpackLog(event, "TFuelTokenLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TFuelTokenBankTFuelTokenUnlockedIterator is returned from FilterTFuelTokenUnlocked and is used to iterate over the raw logs and unpacked data for TFuelTokenUnlocked events raised by the TFuelTokenBank contract.
type TFuelTokenBankTFuelTokenUnlockedIterator struct {
	Event *TFuelTokenBankTFuelTokenUnlocked // Event containing the contract specifics and raw log

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
func (it *TFuelTokenBankTFuelTokenUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TFuelTokenBankTFuelTokenUnlocked)
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
		it.Event = new(TFuelTokenBankTFuelTokenUnlocked)
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
func (it *TFuelTokenBankTFuelTokenUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TFuelTokenBankTFuelTokenUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TFuelTokenBankTFuelTokenUnlocked represents a TFuelTokenUnlocked event raised by the TFuelTokenBank contract.
type TFuelTokenBankTFuelTokenUnlocked struct {
	Denom                       string
	TargetChainTokenReceiver    common.Address
	UnlockedAmount              *big.Int
	SourceChainVoucherBurnNonce *big.Int
	TokenUnlockNonce            *big.Int
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterTFuelTokenUnlocked is a free log retrieval operation binding the contract event 0x5ea3a5ca7f54881fdd7781894d69709e11027910f35647f9d4cc14e6872b6f72.
//
// Solidity: event TFuelTokenUnlocked(string denom, address targetChainTokenReceiver, uint256 unlockedAmount, uint256 sourceChainVoucherBurnNonce, uint256 tokenUnlockNonce)
func (_TFuelTokenBank *TFuelTokenBankFilterer) FilterTFuelTokenUnlocked(opts *bind.FilterOpts) (*TFuelTokenBankTFuelTokenUnlockedIterator, error) {

	logs, sub, err := _TFuelTokenBank.contract.FilterLogs(opts, "TFuelTokenUnlocked")
	if err != nil {
		return nil, err
	}
	return &TFuelTokenBankTFuelTokenUnlockedIterator{contract: _TFuelTokenBank.contract, event: "TFuelTokenUnlocked", logs: logs, sub: sub}, nil
}

// WatchTFuelTokenUnlocked is a free log subscription operation binding the contract event 0x5ea3a5ca7f54881fdd7781894d69709e11027910f35647f9d4cc14e6872b6f72.
//
// Solidity: event TFuelTokenUnlocked(string denom, address targetChainTokenReceiver, uint256 unlockedAmount, uint256 sourceChainVoucherBurnNonce, uint256 tokenUnlockNonce)
func (_TFuelTokenBank *TFuelTokenBankFilterer) WatchTFuelTokenUnlocked(opts *bind.WatchOpts, sink chan<- *TFuelTokenBankTFuelTokenUnlocked) (event.Subscription, error) {

	logs, sub, err := _TFuelTokenBank.contract.WatchLogs(opts, "TFuelTokenUnlocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TFuelTokenBankTFuelTokenUnlocked)
				if err := _TFuelTokenBank.contract.UnpackLog(event, "TFuelTokenUnlocked", log); err != nil {
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

// ParseTFuelTokenUnlocked is a log parse operation binding the contract event 0x5ea3a5ca7f54881fdd7781894d69709e11027910f35647f9d4cc14e6872b6f72.
//
// Solidity: event TFuelTokenUnlocked(string denom, address targetChainTokenReceiver, uint256 unlockedAmount, uint256 sourceChainVoucherBurnNonce, uint256 tokenUnlockNonce)
func (_TFuelTokenBank *TFuelTokenBankFilterer) ParseTFuelTokenUnlocked(log types.Log) (*TFuelTokenBankTFuelTokenUnlocked, error) {
	event := new(TFuelTokenBankTFuelTokenUnlocked)
	if err := _TFuelTokenBank.contract.UnpackLog(event, "TFuelTokenUnlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TFuelTokenBankTFuelVoucherBurnedIterator is returned from FilterTFuelVoucherBurned and is used to iterate over the raw logs and unpacked data for TFuelVoucherBurned events raised by the TFuelTokenBank contract.
type TFuelTokenBankTFuelVoucherBurnedIterator struct {
	Event *TFuelTokenBankTFuelVoucherBurned // Event containing the contract specifics and raw log

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
func (it *TFuelTokenBankTFuelVoucherBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TFuelTokenBankTFuelVoucherBurned)
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
		it.Event = new(TFuelTokenBankTFuelVoucherBurned)
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
func (it *TFuelTokenBankTFuelVoucherBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TFuelTokenBankTFuelVoucherBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TFuelTokenBankTFuelVoucherBurned represents a TFuelVoucherBurned event raised by the TFuelTokenBank contract.
type TFuelTokenBankTFuelVoucherBurned struct {
	Denom                    string
	SourceChainVoucherOwner  common.Address
	TargetChainTokenReceiver common.Address
	BurnedAmount             *big.Int
	VoucherBurnNonce         *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterTFuelVoucherBurned is a free log retrieval operation binding the contract event 0x40f1d475c2aa44f5c23193fab26a64d6aa4e09ab51898b10a3036baf82398ea1.
//
// Solidity: event TFuelVoucherBurned(string denom, address sourceChainVoucherOwner, address targetChainTokenReceiver, uint256 burnedAmount, uint256 voucherBurnNonce)
func (_TFuelTokenBank *TFuelTokenBankFilterer) FilterTFuelVoucherBurned(opts *bind.FilterOpts) (*TFuelTokenBankTFuelVoucherBurnedIterator, error) {

	logs, sub, err := _TFuelTokenBank.contract.FilterLogs(opts, "TFuelVoucherBurned")
	if err != nil {
		return nil, err
	}
	return &TFuelTokenBankTFuelVoucherBurnedIterator{contract: _TFuelTokenBank.contract, event: "TFuelVoucherBurned", logs: logs, sub: sub}, nil
}

// WatchTFuelVoucherBurned is a free log subscription operation binding the contract event 0x40f1d475c2aa44f5c23193fab26a64d6aa4e09ab51898b10a3036baf82398ea1.
//
// Solidity: event TFuelVoucherBurned(string denom, address sourceChainVoucherOwner, address targetChainTokenReceiver, uint256 burnedAmount, uint256 voucherBurnNonce)
func (_TFuelTokenBank *TFuelTokenBankFilterer) WatchTFuelVoucherBurned(opts *bind.WatchOpts, sink chan<- *TFuelTokenBankTFuelVoucherBurned) (event.Subscription, error) {

	logs, sub, err := _TFuelTokenBank.contract.WatchLogs(opts, "TFuelVoucherBurned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TFuelTokenBankTFuelVoucherBurned)
				if err := _TFuelTokenBank.contract.UnpackLog(event, "TFuelVoucherBurned", log); err != nil {
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

// ParseTFuelVoucherBurned is a log parse operation binding the contract event 0x40f1d475c2aa44f5c23193fab26a64d6aa4e09ab51898b10a3036baf82398ea1.
//
// Solidity: event TFuelVoucherBurned(string denom, address sourceChainVoucherOwner, address targetChainTokenReceiver, uint256 burnedAmount, uint256 voucherBurnNonce)
func (_TFuelTokenBank *TFuelTokenBankFilterer) ParseTFuelVoucherBurned(log types.Log) (*TFuelTokenBankTFuelVoucherBurned, error) {
	event := new(TFuelTokenBankTFuelVoucherBurned)
	if err := _TFuelTokenBank.contract.UnpackLog(event, "TFuelVoucherBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TFuelTokenBankTFuelVoucherMintedIterator is returned from FilterTFuelVoucherMinted and is used to iterate over the raw logs and unpacked data for TFuelVoucherMinted events raised by the TFuelTokenBank contract.
type TFuelTokenBankTFuelVoucherMintedIterator struct {
	Event *TFuelTokenBankTFuelVoucherMinted // Event containing the contract specifics and raw log

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
func (it *TFuelTokenBankTFuelVoucherMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TFuelTokenBankTFuelVoucherMinted)
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
		it.Event = new(TFuelTokenBankTFuelVoucherMinted)
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
func (it *TFuelTokenBankTFuelVoucherMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TFuelTokenBankTFuelVoucherMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TFuelTokenBankTFuelVoucherMinted represents a TFuelVoucherMinted event raised by the TFuelTokenBank contract.
type TFuelTokenBankTFuelVoucherMinted struct {
	Denom                      string
	TargetChainVoucherReceiver common.Address
	MintedAmount               *big.Int
	SourceChainTokenLockNonce  *big.Int
	VoucherMintNonce           *big.Int
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterTFuelVoucherMinted is a free log retrieval operation binding the contract event 0x80742bd15a2c8c4ad5d395bcf577073110e52f0c73bf980dfa9453c1d8c354e5.
//
// Solidity: event TFuelVoucherMinted(string denom, address targetChainVoucherReceiver, uint256 mintedAmount, uint256 sourceChainTokenLockNonce, uint256 voucherMintNonce)
func (_TFuelTokenBank *TFuelTokenBankFilterer) FilterTFuelVoucherMinted(opts *bind.FilterOpts) (*TFuelTokenBankTFuelVoucherMintedIterator, error) {

	logs, sub, err := _TFuelTokenBank.contract.FilterLogs(opts, "TFuelVoucherMinted")
	if err != nil {
		return nil, err
	}
	return &TFuelTokenBankTFuelVoucherMintedIterator{contract: _TFuelTokenBank.contract, event: "TFuelVoucherMinted", logs: logs, sub: sub}, nil
}

// WatchTFuelVoucherMinted is a free log subscription operation binding the contract event 0x80742bd15a2c8c4ad5d395bcf577073110e52f0c73bf980dfa9453c1d8c354e5.
//
// Solidity: event TFuelVoucherMinted(string denom, address targetChainVoucherReceiver, uint256 mintedAmount, uint256 sourceChainTokenLockNonce, uint256 voucherMintNonce)
func (_TFuelTokenBank *TFuelTokenBankFilterer) WatchTFuelVoucherMinted(opts *bind.WatchOpts, sink chan<- *TFuelTokenBankTFuelVoucherMinted) (event.Subscription, error) {

	logs, sub, err := _TFuelTokenBank.contract.WatchLogs(opts, "TFuelVoucherMinted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TFuelTokenBankTFuelVoucherMinted)
				if err := _TFuelTokenBank.contract.UnpackLog(event, "TFuelVoucherMinted", log); err != nil {
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

// ParseTFuelVoucherMinted is a log parse operation binding the contract event 0x80742bd15a2c8c4ad5d395bcf577073110e52f0c73bf980dfa9453c1d8c354e5.
//
// Solidity: event TFuelVoucherMinted(string denom, address targetChainVoucherReceiver, uint256 mintedAmount, uint256 sourceChainTokenLockNonce, uint256 voucherMintNonce)
func (_TFuelTokenBank *TFuelTokenBankFilterer) ParseTFuelVoucherMinted(log types.Log) (*TFuelTokenBankTFuelVoucherMinted, error) {
	event := new(TFuelTokenBankTFuelVoucherMinted)
	if err := _TFuelTokenBank.contract.UnpackLog(event, "TFuelVoucherMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
