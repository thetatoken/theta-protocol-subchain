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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"mainchainID_\",\"type\":\"uint256\"},{\"internalType\":\"contractChainRegistrar\",\"name\":\"chainRegistrar_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sourceChainTokenSender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"targetChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"targetChainVoucherReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenLockNonce\",\"type\":\"uint256\"}],\"name\":\"TFuelTokenLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"targetChainTokenReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unlockedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainVoucherBurnNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenUnlockNonce\",\"type\":\"uint256\"}],\"name\":\"TFuelTokenUnlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sourceChainVoucherOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"targetChainTokenReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"burnedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"voucherBurnNonce\",\"type\":\"uint256\"}],\"name\":\"TFuelVoucherBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"targetChainVoucherReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainTokenLockNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"voucherMintNonce\",\"type\":\"uint256\"}],\"name\":\"TFuelVoucherMinted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allDenoms\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allVouchers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"denomToVoucherLookup\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voucherAddress\",\"type\":\"address\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voucherContractAddr\",\"type\":\"address\"}],\"name\":\"getDenom\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"getMaxProcessedTokenLockNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"getMaxProcessedVoucherBurnNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"name\":\"getVoucher\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"isOnMainchain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"mainchainID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenLockNonceMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"tokenLockVotingRecords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dynasty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumlatedShares\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenUnlockNonceMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"totalLockedAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"voucherAddressToDenomLookup\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"voucherBurnNonceMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"voucherBurnVotingRecords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dynasty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumlatedShares\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"voucherMintNonceMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"targetChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"targetChainVoucherReceiver\",\"type\":\"address\"}],\"name\":\"lockTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\",\"payable\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"targetChainVoucherReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mintedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dynasty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainTokenLockNonce\",\"type\":\"uint256\"}],\"name\":\"mintVouchers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targetChainTokenReceiver\",\"type\":\"address\"}],\"name\":\"burnVouchers\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\",\"payable\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainID\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"targetChainTokenReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"unlockAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dynasty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainVoucherBurnNonce\",\"type\":\"uint256\"}],\"name\":\"unlockTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\",\"payable\":true}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200262538038062002625833981016040819052620000349162000064565b6001600081905591909155600280546001600160a01b0319166001600160a01b03909216919091179055620000a3565b600080604083850312156200007857600080fd5b825160208401519092506001600160a01b03811681146200009857600080fd5b809150509250929050565b61257280620000b36000396000f3fe6080604052600436106101405760003560e01c8063766f8fb0116100b6578063ccf187c71161006f578063ccf187c714610454578063ebda996214610481578063f6a3d24e146104a1578063f95627ac146104dd578063feaff0521461050a578063ff248a441461054957600080fd5b8063766f8fb0146103875780637d0fb00d146103b45780638883931e146103c7578063a2cc6981146103f4578063aa68acde14610414578063ca2075691461042757600080fd5b8063261a323e11610108578063261a323e1461027d57806327ca4df1146102ad5780634250863b146102e5578063588b1408146102ff57806360569b5e1461032c578063740cb7f81461035a57600080fd5b8063073b9502146101455780631527b14d1461016e57806319fd1a11146101da5780631a0483d3146102075780631eb7873714610229575b600080fd5b34801561015157600080fd5b5061015b60015481565b6040519081526020015b60405180910390f35b34801561017a57600080fd5b506101bb610189366004611fb0565b8051602081830181018051600b825292820191909301209152546001600160a01b03811690600160a01b900460ff1682565b604080516001600160a01b039093168352901515602083015201610165565b3480156101e657600080fd5b5061015b6101f536600461204f565b60106020526000908152604090205481565b34801561021357600080fd5b50610227610222366004611fe5565b61055c565b005b34801561023557600080fd5b506102686102443660046120c2565b60096020908152600092835260408084209091529082529020805460029091015482565b60408051928352602083019190915201610165565b34801561028957600080fd5b5061029d610298366004611fb0565b61073e565b6040519015158152602001610165565b3480156102b957600080fd5b506102cd6102c836600461204f565b61077e565b6040516001600160a01b039091168152602001610165565b3480156102f157600080fd5b50600f5461029d9060ff1681565b34801561030b57600080fd5b5061031f61031a36600461204f565b6107a8565b60405161016591906121df565b34801561033857600080fd5b5061034c610347366004611ea5565b610854565b6040516101659291906122b8565b34801561036657600080fd5b5061015b61037536600461204f565b60066020526000908152604090205481565b34801561039357600080fd5b5061015b6103a236600461204f565b60009081526008602052604090205490565b6102276103c2366004611ea5565b6108fb565b3480156103d357600080fd5b5061015b6103e236600461204f565b60036020526000908152604090205481565b34801561040057600080fd5b506102cd61040f366004611fb0565b610a50565b610227610422366004612068565b610ac1565b34801561043357600080fd5b5061015b61044236600461204f565b60056020526000908152604090205481565b34801561046057600080fd5b5061015b61046f36600461204f565b60046020526000908152604090205481565b34801561048d57600080fd5b5061031f61049c366004611ea5565b610d04565b3480156104ad57600080fd5b5061029d6104bc366004611ea5565b6001600160a01b03166000908152600c602052604090206001015460ff1690565b3480156104e957600080fd5b5061015b6104f836600461204f565b60009081526007602052604090205490565b34801561051657600080fd5b506102686105253660046120c2565b600a6020908152600092835260408084209091529082529020805460029091015482565b610227610557366004612098565b610dfb565b600260005414156105885760405162461bcd60e51b815260040161057f906122dc565b60405180910390fd5b60026000556001544614156106055760405162461bcd60e51b815260206004820152603e60248201527f544675656c546f6b656e42616e6b2e6d696e74566f756368657273282920636160448201527f6e206f6e6c792062652063616c6c6564206f6e206120737562636861696e0000606482015260840161057f565b600061061086611094565b905060008061061e836110a5565b915091508061067b5760405162461bcd60e51b8152602060048201526024808201527f4661696c656420746f206578747261637420636861696e49442066726f6d2064604482015263656e6f6d60e01b606482015260840161057f565b600082848989898960405160200161069896959493929190612186565b60405160208183030381529060405280519060200120905060006106bf848884893361112d565b9050801561072d576106d1898961114a565b6106da46611329565b46600090815260066020526040908190205490517f80742bd15a2c8c4ad5d395bcf577073110e52f0c73bf980dfa9453c1d8c354e590610723908d908d908d908c9087906121f2565b60405180910390a1505b505060016000555050505050505050565b60008061074a83611094565b9050600b8160405161075c9190612110565b9081526040519081900360200190205460ff600160a01b909104169392505050565b600d818154811061078e57600080fd5b6000918252602090912001546001600160a01b0316905081565b600e81815481106107b857600080fd5b9060005260206000200160009150905080546107d39061243e565b80601f01602080910402602001604051908101604052809291908181526020018280546107ff9061243e565b801561084c5780601f106108215761010080835404028352916020019161084c565b820191906000526020600020905b81548152906001019060200180831161082f57829003601f168201915b505050505081565b600c6020526000908152604090208054819061086f9061243e565b80601f016020809104026020016040519081016040528092919081815260200182805461089b9061243e565b80156108e85780601f106108bd576101008083540402835291602001916108e8565b820191906000526020600020905b8154815290600101906020018083116108cb57829003601f168201915b5050506001909301549192505060ff1682565b6002600054141561091e5760405162461bcd60e51b815260040161057f906122dc565b600260005560015446141561099b5760405162461bcd60e51b815260206004820152603e60248201527f544675656c546f6b656e42616e6b2e6275726e566f756368657273282920636160448201527f6e206f6e6c792062652063616c6c6564206f6e206120737562636861696e0000606482015260840161057f565b33346109a681611350565b6109af466114b4565b60006109ef600154604051806040016040528060018152602001600360fc1b8152506040518060600160405280602a8152602001612513602a91396114d3565b4660009081526005602052604090819020549051919250907f40f1d475c2aa44f5c23193fab26a64d6aa4e09ab51898b10a3036baf82398ea190610a3c9084908790899088908790612232565b60405180910390a150506001600055505050565b600080610a5c83611094565b90506000600b82604051610a709190612110565b908152604080516020928190038301812081830190925290546001600160a01b0381168252600160a01b900460ff16158015928201929092529150610ab757519392505050565b5060009392505050565b60026000541415610ae45760405162461bcd60e51b815260040161057f906122dc565b60026000556001544614610b62576040805162461bcd60e51b81526020600482015260248101919091527f544675656c546f6b656e42616e6b2e6c6f636b546f6b656e7328292063616e2060448201527f6f6e6c792062652063616c6c6564206f6e20746865206d61696e20636861696e606482015260840161057f565b6002546040516343b71f0560e01b8152600481018490526001600160a01b03909116906343b71f0590602401602060405180830381600087803b158015610ba857600080fd5b505af1158015610bbc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610be09190611f8e565b610c2c5760405162461bcd60e51b815260206004820181905260248201527f746172676574436861696e4944206e6f74207965742072656769737465726564604482015260640161057f565b60008281526010602052604090205434903390610c499083611511565b600085815260106020526040902055610c6184611524565b6000610ca1600154604051806040016040528060018152602001600360fc1b8152506040518060600160405280602a8152602001612513602a91396114d3565b60008681526003602052604090819020549051919250907fee1ecc2b21aa613cc77cd44823a68ef1168ce1f40c2eac1d68690baf955fdbd190610cef90849086908a908a908a908890612272565b60405180910390a15050600160005550505050565b6001600160a01b0381166000908152600c60205260408082208151808301909252805460609392919082908290610d3a9061243e565b80601f0160208091040260200160405190810160405280929190818152602001828054610d669061243e565b8015610db35780601f10610d8857610100808354040283529160200191610db3565b820191906000526020600020905b815481529060010190602001808311610d9657829003601f168201915b50505091835250506001919091015460ff16151560209182015281015190915015610ddf575192915050565b5050604080516020810190915260008152919050565b50919050565b60026000541415610e1e5760405162461bcd60e51b815260040161057f906122dc565b60026000556001544614610ea55760405162461bcd60e51b815260206004820152604260248201527f544675656c546f6b656e42616e6b2e756e6c6f636b546f6b656e73282920636160448201527f6e206f6e6c792062652063616c6c6564206f6e20746865206d61696e2063686160648201526134b760f11b608482015260a40161057f565b600085815260106020526040902054831115610f175760405162461bcd60e51b815260206004820152602b60248201527f43616e6e6f7420756e6c6f636b207468652072657175657374656420616d6f7560448201526a1b9d081bd98815119d595b60aa1b606482015260840161057f565b6040805160208101879052908101849052606085811b6bffffffffffffffffffffffff191690820152607481018390526094810182905260009060b40160408051601f19818403018152919052805160208201209091506000610f7d8886848733611543565b9050801561108557600088815260106020526040902054610f9e9087611556565b6000898152601060205260408082209290925590516001600160a01b0389169188156108fc02918991818181858888f19350505050158015610fe4573d6000803e3d6000fd5b50610fee88611562565b600061102e600154604051806040016040528060018152602001600360fc1b8152506040518060600160405280602a8152602001612513602a91396114d3565b60008a81526004602052604090819020549051919250907f5ea3a5ca7f54881fdd7781894d69709e11027910f35647f9d4cc14e6872b6f729061107a9084908c908c908b9087906121f2565b60405180910390a150505b50506001600055505050505050565b606061109f82611581565b92915050565b6000806000806110bb85602f60f81b60016115fb565b91509150806110d1575060009485945092505050565b6000806110e08760008661168b565b91509150806110f85750600096879650945050505050565b60008061110484611798565b915091508061111e57506000988998509650505050505050565b50976001975095505050505050565b600061114086868686866007600961186a565b9695505050505050565b60408051603480825260608201909252600091602082018180368337019050509050816000805b60148110156111d9578560601b816014811061118f5761118f6124ce565b1a60f81b848361119e81612473565b9450815181106111b0576111b06124ce565b60200101906001600160f81b031916908160001a905350806111d181612473565b915050611171565b5060005b6020811015611242578281602081106111f8576111f86124ce565b1a60f81b848361120781612473565b945081518110611219576112196124ce565b60200101906001600160f81b031916908160001a9053508061123a81612473565b9150506111dd565b600060b66001600160a01b03168560405161125d9190612110565b6000604051808303816000865af19150503d806000811461129a576040519150601f19603f3d011682016040523d82523d6000602084013e61129f565b606091505b50509050806113205760405162461bcd60e51b815260206004820152604160248201527f544675656c546f6b656e42616e6b2e5f6d696e74544675656c566f756368657260448201527f28293a206661696c656420746f206d696e7420544675656c20766f75636865726064820152607360f81b608482015260a40161057f565b50505050505050565b6000818152600660205260408120805460019290611348908490612368565b909155505050565b6040805160208082528183019092526000916020820181803683370190505090508160005b60208110156113cf57818160208110611390576113906124ce565b1a60f81b8382815181106113a6576113a66124ce565b60200101906001600160f81b031916908160001a905350806113c781612473565b915050611375565b50600060b76001600160a01b0316836040516113eb9190612110565b6000604051808303816000865af19150503d8060008114611428576040519150601f19603f3d011682016040523d82523d6000602084013e61142d565b606091505b50509050806114ae5760405162461bcd60e51b815260206004820152604160248201527f544675656c546f6b656e42616e6b2e5f6275726e544675656c566f756368657260448201527f28293a206661696c656420746f206275726e20544675656c20766f75636865726064820152607360f81b608482015260a40161057f565b50505050565b6000818152600560205260408120805460019290611348908490612368565b60606115096114e185611c6a565b84846040516020016114f59392919061212c565b604051602081830303815290604052611094565b949350505050565b600061151d8284612368565b9392505050565b6000818152600360205260408120805460019290611348908490612368565b600061114086868686866008600a61186a565b600061151d82846123d8565b6000818152600460205260408120805460019290611348908490612368565b60608160005b81518110156115f4576115b98282815181106115a5576115a56124ce565b01602001516001600160f81b031916611d68565b8282815181106115cb576115cb6124ce565b60200101906001600160f81b031916908160001a905350806115ec81612473565b915050611587565b5092915050565b82516000908190859082805b8281101561167757876001600160f81b03191684828151811061162c5761162c6124ce565b01602001516001600160f81b03191614156116655761164c600183612368565b9150868214156116655794506001935061168392505050565b8061166f81612473565b915050611607565b50600080945094505050505b935093915050565b8251606090600090848410156116b557505060408051602081019091526000808252909150611683565b808411156116d757505060408051602081019091526000808252909150611683565b8560006116e48688611556565b67ffffffffffffffff8111156116fc576116fc6124e4565b6040519080825280601f01601f191660200182016040528015611726576020820181803683370190505b509050865b8681101561111e57828181518110611745576117456124ce565b01602001516001600160f81b0319168261175f838b611556565b8151811061176f5761176f6124ce565b60200101906001600160f81b031916908160001a9053508061179081612473565b91505061172b565b80516000908190839082805b8281101561185d5760308482815181106117c0576117c06124ce565b016020015160f81c108015906117f0575060398482815181106117e5576117e56124ce565b016020015160f81c11155b1561183c5761180082600a611db7565b91506118356030858381518110611819576118196124ce565b016020015161182b919060f81c6123ef565b839060ff16611511565b915061184b565b50600096879650945050505050565b8061185581612473565b9150506117a4565b5095600195509350505050565b60006001548814806118f557506002546040516343b71f0560e01b8152600481018a90526001600160a01b03909116906343b71f0590602401602060405180830381600087803b1580156118bd57600080fd5b505af11580156118d1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118f59190611f8e565b6119335760405162461bcd60e51b815260206004820152600f60248201526e125b9d985b1a590818da185a5b9251608a1b604482015260640161057f565b60008881526020849052604090205461194d906001612368565b851461195b57506000611c5f565b600088815260208381526040808320898452909152812060015460609081908c908114156119895750620581495b6002546040516343f27e4560e01b815260048101839052602481018e90526001600160a01b03909116906343f27e4590604401600060405180830381600087803b1580156119d657600080fd5b505af11580156119ea573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611a129190810190611ec2565b9350915060005b8251811015611b6f57896001600160a01b0316838281518110611a3e57611a3e6124ce565b60200260200101516001600160a01b031614611a5957611b5d565b6001955060005b6001860154811015611afa57856001018181548110611a8157611a816124ce565b6000918252602090912001546001600160a01b038c811691161415611ae85760405162461bcd60e51b815260206004820152601c60248201527f546869732076616c696461746f7220616c726561647920766f74656400000000604482015260640161057f565b80611af281612473565b915050611a60565b508c85556001808601805491820181556000908152602090200180546001600160a01b031916331790558351611b5790859083908110611b3c57611b3c6124ce565b6020026020010151866002015461151190919063ffffffff16565b60028601555b80611b6781612473565b915050611a19565b50505082611bb15760405162461bcd60e51b815260206004820152600f60248201526e2737ba1030903b30b634b230ba37b960891b604482015260640161057f565b6000805b8251811015611bfe57611bea838281518110611bd357611bd36124ce565b60200260200101518361151190919063ffffffff16565b915080611bf681612473565b915050611bb5565b50611c0a816002611db7565b6002840154611c1a906003611db7565b10611c565760008c815260208890526040902054611c39906001612368565b60008d8152602089905260409020555060019350611c5f92505050565b60009450505050505b979650505050505050565b606081611c8e5750506040805180820190915260018152600360fc1b602082015290565b8160005b8115611cb85780611ca281612473565b9150611cb19050600a836123a5565b9150611c92565b60008167ffffffffffffffff811115611cd357611cd36124e4565b6040519080825280601f01601f191660200182016040528015611cfd576020820181803683370190505b5090505b841561150957611d126001836123d8565b9150611d1f600a8661248e565b611d2a906030612368565b60f81b818381518110611d3f57611d3f6124ce565b60200101906001600160f81b031916908160001a905350611d61600a866123a5565b9450611d01565b6000604160f81b6001600160f81b0319831610801590611d965750602d60f91b6001600160f81b0319831611155b15611db357611daa60f883901c6020612380565b60f81b92915050565b5090565b600061151d82846123b9565b600082601f830112611dd457600080fd5b81516020611de9611de483612344565b612313565b80838252828201915082860187848660051b8901011115611e0957600080fd5b60005b85811015611e2857815184529284019290840190600101611e0c565b5090979650505050505050565b600082601f830112611e4657600080fd5b813567ffffffffffffffff811115611e6057611e606124e4565b611e73601f8201601f1916602001612313565b818152846020838601011115611e8857600080fd5b816020850160208301376000918101602001919091529392505050565b600060208284031215611eb757600080fd5b813561151d816124fa565b60008060408385031215611ed557600080fd5b825167ffffffffffffffff80821115611eed57600080fd5b818501915085601f830112611f0157600080fd5b81516020611f11611de483612344565b8083825282820191508286018a848660051b8901011115611f3157600080fd5b600096505b84871015611f5d578051611f49816124fa565b835260019690960195918301918301611f36565b5091880151919650909350505080821115611f7757600080fd5b50611f8485828601611dc3565b9150509250929050565b600060208284031215611fa057600080fd5b8151801515811461151d57600080fd5b600060208284031215611fc257600080fd5b813567ffffffffffffffff811115611fd957600080fd5b61150984828501611e35565b600080600080600060a08688031215611ffd57600080fd5b853567ffffffffffffffff81111561201457600080fd5b61202088828901611e35565b9550506020860135612031816124fa565b94979496505050506040830135926060810135926080909101359150565b60006020828403121561206157600080fd5b5035919050565b6000806040838503121561207b57600080fd5b82359150602083013561208d816124fa565b809150509250929050565b600080600080600060a086880312156120b057600080fd5b853594506020860135612031816124fa565b600080604083850312156120d557600080fd5b50508035926020909101359150565b600081518084526120fc816020860160208601612412565b601f01601f19169290920160200192915050565b60008251612122818460208701612412565b9190910192915050565b6000845161213e818460208901612412565b8083019050602f60f81b808252855161215e816001850160208a01612412565b60019201918201528351612179816002840160208801612412565b0160020195945050505050565b8681526000865161219e816020850160208b01612412565b60609690961b6bffffffffffffffffffffffff1916602092909601918201959095526034810193909352605483019190915260748201526094019392505050565b60208152600061151d60208301846120e4565b60a08152600061220560a08301886120e4565b6001600160a01b039690961660208301525060408101939093526060830191909152608090910152919050565b60a08152600061224560a08301886120e4565b6001600160a01b039687166020840152949095166040820152606081019290925260809091015292915050565b60c08152600061228560c08301896120e4565b6001600160a01b039788166020840152604083019690965250929094166060830152608082015260a00191909152919050565b6040815260006122cb60408301856120e4565b905082151560208301529392505050565b6020808252601f908201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604082015260600190565b604051601f8201601f1916810167ffffffffffffffff8111828210171561233c5761233c6124e4565b604052919050565b600067ffffffffffffffff82111561235e5761235e6124e4565b5060051b60200190565b6000821982111561237b5761237b6124a2565b500190565b600060ff821660ff84168060ff0382111561239d5761239d6124a2565b019392505050565b6000826123b4576123b46124b8565b500490565b60008160001904831182151516156123d3576123d36124a2565b500290565b6000828210156123ea576123ea6124a2565b500390565b600060ff821660ff841680821015612409576124096124a2565b90039392505050565b60005b8381101561242d578181015183820152602001612415565b838111156114ae5750506000910152565b600181811c9082168061245257607f821691505b60208210811415610df557634e487b7160e01b600052602260045260246000fd5b6000600019821415612487576124876124a2565b5060010190565b60008261249d5761249d6124b8565b500690565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052601260045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b038116811461250f57600080fd5b5056fe307830303030303030303030303030303030303030303030303030303030303030303030303030303030a2646970667358221220e1d800603c0e3a973ed00a8838488168a9b5bedfdd2e6d70cc16eb15b9189b1564736f6c63430008070033",
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
