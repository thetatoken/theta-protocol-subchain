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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"mainchainID_\",\"type\":\"uint256\"},{\"internalType\":\"contractChainRegistrar\",\"name\":\"chainRegistrar_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sourceChainTokenSender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"targetChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"targetChainVoucherReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenLockNonce\",\"type\":\"uint256\"}],\"name\":\"TFuelTokenLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"targetChainTokenReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unlockedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainVoucherBurnNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenUnlockNonce\",\"type\":\"uint256\"}],\"name\":\"TFuelTokenUnlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sourceChainVoucherOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"targetChainTokenReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"burnedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"voucherBurnNonce\",\"type\":\"uint256\"}],\"name\":\"TFuelVoucherBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"targetChainVoucherReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainTokenLockNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"voucherMintNonce\",\"type\":\"uint256\"}],\"name\":\"TFuelVoucherMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"}],\"name\":\"has\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allDenoms\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allVouchers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"denomToVoucherLookup\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voucherAddress\",\"type\":\"address\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voucherContractAddr\",\"type\":\"address\"}],\"name\":\"getDenom\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"getMaxProcessedTokenLockNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"getMaxProcessedVoucherBurnNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"name\":\"getVoucher\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOnMainchain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mainchainID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"subchainID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenLockNonceMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"tokenLockVotingRecords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dynasty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumlatedShares\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenUnlockNonceMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"totalLockedAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"voucherAddressToDenomLookup\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"voucherBurnNonceMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"voucherBurnVotingRecords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dynasty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumlatedShares\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"voucherMintNonceMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"test\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"targetChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"targetChainVoucherReceiver\",\"type\":\"address\"}],\"name\":\"lockTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"targetChainVoucherReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mintedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dynasty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainTokenLockNonce\",\"type\":\"uint256\"}],\"name\":\"mintVouchers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targetChainTokenReceiver\",\"type\":\"address\"}],\"name\":\"burnVouchers\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainID\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"targetChainTokenReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"unlockAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dynasty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainVoucherBurnNonce\",\"type\":\"uint256\"}],\"name\":\"unlockTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200267b3803806200267b833981016040819052620000349162000064565b6001600081905591909155600380546001600160a01b0319166001600160a01b03909216919091179055620000a3565b600080604083850312156200007857600080fd5b825160208401519092506001600160a01b03811681146200009857600080fd5b809150509250929050565b6125c880620000b36000396000f3fe6080604052600436106101665760003560e01c80637d0fb00d116100d1578063ebda99621161008a578063f8e8f7aa11610064578063f8e8f7aa14610518578063f95627ac1461052e578063feaff0521461055b578063ff248a441461059a57600080fd5b8063ebda9962146104a7578063f6a3d24e146104c7578063f8a8fd6d1461050357600080fd5b80637d0fb00d146103da5780638883931e146103ed578063a2cc69811461041a578063aa68acde1461043a578063ca2075691461044d578063ccf187c71461047a57600080fd5b806327ca4df11161012357806327ca4df1146102d35780634250863b1461030b578063588b14081461032557806360569b5e14610352578063740cb7f814610380578063766f8fb0146103ad57600080fd5b8063073b95021461016b5780631527b14d1461019457806319fd1a11146102005780631a0483d31461022d5780631eb787371461024f578063261a323e146102a3575b600080fd5b34801561017757600080fd5b5061018160015481565b6040519081526020015b60405180910390f35b3480156101a057600080fd5b506101e16101af366004612006565b8051602081830181018051600c825292820191909301209152546001600160a01b03811690600160a01b900460ff1682565b604080516001600160a01b03909316835290151560208301520161018b565b34801561020c57600080fd5b5061018161021b3660046120a5565b60116020526000908152604090205481565b34801561023957600080fd5b5061024d61024836600461203b565b6105ad565b005b34801561025b57600080fd5b5061028e61026a366004612118565b600a6020908152600092835260408084209091529082529020805460029091015482565b6040805192835260208301919091520161018b565b3480156102af57600080fd5b506102c36102be366004612006565b610792565b604051901515815260200161018b565b3480156102df57600080fd5b506102f36102ee3660046120a5565b6107d2565b6040516001600160a01b03909116815260200161018b565b34801561031757600080fd5b506010546102c39060ff1681565b34801561033157600080fd5b506103456103403660046120a5565b6107fc565b60405161018b9190612235565b34801561035e57600080fd5b5061037261036d366004611efb565b6108a8565b60405161018b92919061230e565b34801561038c57600080fd5b5061018161039b3660046120a5565b60076020526000908152604090205481565b3480156103b957600080fd5b506101816103c83660046120a5565b60009081526009602052604090205490565b61024d6103e8366004611efb565b61094f565b3480156103f957600080fd5b506101816104083660046120a5565b60046020526000908152604090205481565b34801561042657600080fd5b506102f3610435366004612006565b610aa7565b61024d6104483660046120be565b610b18565b34801561045957600080fd5b506101816104683660046120a5565b60066020526000908152604090205481565b34801561048657600080fd5b506101816104953660046120a5565b60056020526000908152604090205481565b3480156104b357600080fd5b506103456104c2366004611efb565b610d46565b3480156104d357600080fd5b506102c36104e2366004611efb565b6001600160a01b03166000908152600d602052604090206001015460ff1690565b34801561050f57600080fd5b50610181610e3d565b34801561052457600080fd5b5061018160025481565b34801561053a57600080fd5b506101816105493660046120a5565b60009081526008602052604090205490565b34801561056757600080fd5b5061028e610576366004612118565b600b6020908152600092835260408084209091529082529020805460029091015482565b61024d6105a83660046120ee565b610e52565b600260005414156105d95760405162461bcd60e51b81526004016105d090612332565b60405180910390fd5b600260005560015446908114156106585760405162461bcd60e51b815260206004820152603e60248201527f544675656c546f6b656e42616e6b2e6d696e74566f756368657273282920636160448201527f6e206f6e6c792062652063616c6c6564206f6e206120737562636861696e000060648201526084016105d0565b6000610663876110eb565b9050600080610671836110fc565b91509150806106ce5760405162461bcd60e51b8152602060048201526024808201527f4661696c656420746f206578747261637420636861696e49442066726f6d2064604482015263656e6f6d60e01b60648201526084016105d0565b600082848a8a8a8a6040516020016106eb969594939291906121dc565b60405160208183030381529060405280519060200120905060006107128489848a33611184565b90508015610780576107248a8a6111a1565b61072d46611380565b46600090815260076020526040908190205490517f80742bd15a2c8c4ad5d395bcf577073110e52f0c73bf980dfa9453c1d8c354e590610776908e908e908e908d908790612248565b60405180910390a1505b50506001600055505050505050505050565b60008061079e836110eb565b9050600c816040516107b09190612166565b9081526040519081900360200190205460ff600160a01b909104169392505050565b600e81815481106107e257600080fd5b6000918252602090912001546001600160a01b0316905081565b600f818154811061080c57600080fd5b90600052602060002001600091509050805461082790612494565b80601f016020809104026020016040519081016040528092919081815260200182805461085390612494565b80156108a05780601f10610875576101008083540402835291602001916108a0565b820191906000526020600020905b81548152906001019060200180831161088357829003601f168201915b505050505081565b600d602052600090815260409020805481906108c390612494565b80601f01602080910402602001604051908101604052809291908181526020018280546108ef90612494565b801561093c5780601f106109115761010080835404028352916020019161093c565b820191906000526020600020905b81548152906001019060200180831161091f57829003601f168201915b5050506001909301549192505060ff1682565b600260005414156109725760405162461bcd60e51b81526004016105d090612332565b600260005560015446908114156109f15760405162461bcd60e51b815260206004820152603e60248201527f544675656c546f6b656e42616e6b2e6275726e566f756368657273282920636160448201527f6e206f6e6c792062652063616c6c6564206f6e206120737562636861696e000060648201526084016105d0565b33346109fc816113a7565b610a054661150b565b6000610a45600154604051806040016040528060018152602001600360fc1b8152506040518060600160405280602a8152602001612569602a913961152a565b4660009081526006602052604090819020549051919250907f40f1d475c2aa44f5c23193fab26a64d6aa4e09ab51898b10a3036baf82398ea190610a9290849087908a9088908790612288565b60405180910390a15050600160005550505050565b600080610ab3836110eb565b90506000600c82604051610ac79190612166565b908152604080516020928190038301812081830190925290546001600160a01b0381168252600160a01b900460ff16158015928201929092529150610b0e57519392505050565b5060009392505050565b60026000541415610b3b5760405162461bcd60e51b81526004016105d090612332565b60026000556001544614610bb9576040805162461bcd60e51b81526020600482015260248101919091527f544675656c546f6b656e42616e6b2e6c6f636b546f6b656e7328292063616e2060448201527f6f6e6c792062652063616c6c6564206f6e20746865206d61696e20636861696e60648201526084016105d0565b6003546040516343b71f0560e01b8152600481018490526001600160a01b03909116906343b71f0590602401602060405180830381600087803b158015610bff57600080fd5b505af1158015610c13573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c379190611fe4565b610c835760405162461bcd60e51b815260206004820181905260248201527f746172676574436861696e4944206e6f7420796574207265676973746572656460448201526064016105d0565b60008281526011602052604090205434903390610ca09083611568565b600085815260116020526040902055610cb88461157b565b6000610cf8600154604051806040016040528060018152602001600360fc1b8152506040518060600160405280602a8152602001612569602a913961152a565b60008681526004602052604090819020549051919250907fee1ecc2b21aa613cc77cd44823a68ef1168ce1f40c2eac1d68690baf955fdbd190610a9290849086908a908a908a9088906122c8565b6001600160a01b0381166000908152600d60205260408082208151808301909252805460609392919082908290610d7c90612494565b80601f0160208091040260200160405190810160405280929190818152602001828054610da890612494565b8015610df55780601f10610dca57610100808354040283529160200191610df5565b820191906000526020600020905b815481529060010190602001808311610dd857829003601f168201915b50505091835250506001919091015460ff16151560209182015281015190915015610e21575192915050565b5050604080516020810190915260008152919050565b50919050565b60004660011415610e4d57600080fd5b504690565b60026000541415610e755760405162461bcd60e51b81526004016105d090612332565b60026000556001544614610efc5760405162461bcd60e51b815260206004820152604260248201527f544675656c546f6b656e42616e6b2e756e6c6f636b546f6b656e73282920636160448201527f6e206f6e6c792062652063616c6c6564206f6e20746865206d61696e2063686160648201526134b760f11b608482015260a4016105d0565b600085815260116020526040902054831115610f6e5760405162461bcd60e51b815260206004820152602b60248201527f43616e6e6f7420756e6c6f636b207468652072657175657374656420616d6f7560448201526a1b9d081bd98815119d595b60aa1b60648201526084016105d0565b6040805160208101879052908101849052606085811b6bffffffffffffffffffffffff191690820152607481018390526094810182905260009060b40160408051601f19818403018152919052805160208201209091506000610fd4888684873361159a565b905080156110dc57600088815260116020526040902054610ff590876115ad565b6000898152601160205260408082209290925590516001600160a01b0389169188156108fc02918991818181858888f1935050505015801561103b573d6000803e3d6000fd5b50611045886115b9565b6000611085600154604051806040016040528060018152602001600360fc1b8152506040518060600160405280602a8152602001612569602a913961152a565b60008a81526005602052604090819020549051919250907f5ea3a5ca7f54881fdd7781894d69709e11027910f35647f9d4cc14e6872b6f72906110d19084908c908c908b908790612248565b60405180910390a150505b50506001600055505050505050565b60606110f6826115d8565b92915050565b60008060008061111285602f60f81b6001611652565b9150915080611128575060009485945092505050565b600080611137876000866116e2565b915091508061114f5750600096879650945050505050565b60008061115b846117ef565b915091508061117557506000988998509650505050505050565b50976001975095505050505050565b600061119786868686866008600a6118c1565b9695505050505050565b60408051603480825260608201909252600091602082018180368337019050509050816000805b6014811015611230578560601b81601481106111e6576111e6612524565b1a60f81b84836111f5816124c9565b94508151811061120757611207612524565b60200101906001600160f81b031916908160001a90535080611228816124c9565b9150506111c8565b5060005b60208110156112995782816020811061124f5761124f612524565b1a60f81b848361125e816124c9565b94508151811061127057611270612524565b60200101906001600160f81b031916908160001a90535080611291816124c9565b915050611234565b600060b66001600160a01b0316856040516112b49190612166565b6000604051808303816000865af19150503d80600081146112f1576040519150601f19603f3d011682016040523d82523d6000602084013e6112f6565b606091505b50509050806113775760405162461bcd60e51b815260206004820152604160248201527f544675656c546f6b656e42616e6b2e5f6d696e74544675656c566f756368657260448201527f28293a206661696c656420746f206d696e7420544675656c20766f75636865726064820152607360f81b608482015260a4016105d0565b50505050505050565b600081815260076020526040812080546001929061139f9084906123be565b909155505050565b6040805160208082528183019092526000916020820181803683370190505090508160005b6020811015611426578181602081106113e7576113e7612524565b1a60f81b8382815181106113fd576113fd612524565b60200101906001600160f81b031916908160001a9053508061141e816124c9565b9150506113cc565b50600060b76001600160a01b0316836040516114429190612166565b6000604051808303816000865af19150503d806000811461147f576040519150601f19603f3d011682016040523d82523d6000602084013e611484565b606091505b50509050806115055760405162461bcd60e51b815260206004820152604160248201527f544675656c546f6b656e42616e6b2e5f6275726e544675656c566f756368657260448201527f28293a206661696c656420746f206275726e20544675656c20766f75636865726064820152607360f81b608482015260a4016105d0565b50505050565b600081815260066020526040812080546001929061139f9084906123be565b606061156061153885611cc0565b848460405160200161154c93929190612182565b6040516020818303038152906040526110eb565b949350505050565b600061157482846123be565b9392505050565b600081815260046020526040812080546001929061139f9084906123be565b600061119786868686866009600b6118c1565b6000611574828461242e565b600081815260056020526040812080546001929061139f9084906123be565b60608160005b815181101561164b576116108282815181106115fc576115fc612524565b01602001516001600160f81b031916611dbe565b82828151811061162257611622612524565b60200101906001600160f81b031916908160001a90535080611643816124c9565b9150506115de565b5092915050565b82516000908190859082805b828110156116ce57876001600160f81b03191684828151811061168357611683612524565b01602001516001600160f81b03191614156116bc576116a36001836123be565b9150868214156116bc579450600193506116da92505050565b806116c6816124c9565b91505061165e565b50600080945094505050505b935093915050565b82516060906000908484101561170c575050604080516020810190915260008082529091506116da565b8084111561172e575050604080516020810190915260008082529091506116da565b85600061173b86886115ad565b67ffffffffffffffff8111156117535761175361253a565b6040519080825280601f01601f19166020018201604052801561177d576020820181803683370190505b509050865b868110156111755782818151811061179c5761179c612524565b01602001516001600160f81b031916826117b6838b6115ad565b815181106117c6576117c6612524565b60200101906001600160f81b031916908160001a905350806117e7816124c9565b915050611782565b80516000908190839082805b828110156118b457603084828151811061181757611817612524565b016020015160f81c108015906118475750603984828151811061183c5761183c612524565b016020015160f81c11155b156118935761185782600a611e0d565b915061188c603085838151811061187057611870612524565b0160200151611882919060f81c612445565b839060ff16611568565b91506118a2565b50600096879650945050505050565b806118ac816124c9565b9150506117fb565b5095600195509350505050565b600060015488148061194c57506003546040516343b71f0560e01b8152600481018a90526001600160a01b03909116906343b71f0590602401602060405180830381600087803b15801561191457600080fd5b505af1158015611928573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061194c9190611fe4565b61198a5760405162461bcd60e51b815260206004820152600f60248201526e125b9d985b1a590818da185a5b9251608a1b60448201526064016105d0565b6000888152602084905260409020546119a49060016123be565b85146119b257506000611cb5565b600088815260208381526040808320898452909152812060015460609081908c908114156119df57506002545b6003546040516343f27e4560e01b815260048101839052602481018e90526001600160a01b03909116906343f27e4590604401600060405180830381600087803b158015611a2c57600080fd5b505af1158015611a40573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611a689190810190611f18565b9350915060005b8251811015611bc557896001600160a01b0316838281518110611a9457611a94612524565b60200260200101516001600160a01b031614611aaf57611bb3565b6001955060005b6001860154811015611b5057856001018181548110611ad757611ad7612524565b6000918252602090912001546001600160a01b038c811691161415611b3e5760405162461bcd60e51b815260206004820152601c60248201527f546869732076616c696461746f7220616c726561647920766f7465640000000060448201526064016105d0565b80611b48816124c9565b915050611ab6565b508c85556001808601805491820181556000908152602090200180546001600160a01b031916331790558351611bad90859083908110611b9257611b92612524565b6020026020010151866002015461156890919063ffffffff16565b60028601555b80611bbd816124c9565b915050611a6f565b50505082611c075760405162461bcd60e51b815260206004820152600f60248201526e2737ba1030903b30b634b230ba37b960891b60448201526064016105d0565b6000805b8251811015611c5457611c40838281518110611c2957611c29612524565b60200260200101518361156890919063ffffffff16565b915080611c4c816124c9565b915050611c0b565b50611c60816002611e0d565b6002840154611c70906003611e0d565b10611cac5760008c815260208890526040902054611c8f9060016123be565b60008d8152602089905260409020555060019350611cb592505050565b60009450505050505b979650505050505050565b606081611ce45750506040805180820190915260018152600360fc1b602082015290565b8160005b8115611d0e5780611cf8816124c9565b9150611d079050600a836123fb565b9150611ce8565b60008167ffffffffffffffff811115611d2957611d2961253a565b6040519080825280601f01601f191660200182016040528015611d53576020820181803683370190505b5090505b841561156057611d6860018361242e565b9150611d75600a866124e4565b611d809060306123be565b60f81b818381518110611d9557611d95612524565b60200101906001600160f81b031916908160001a905350611db7600a866123fb565b9450611d57565b6000604160f81b6001600160f81b0319831610801590611dec5750602d60f91b6001600160f81b0319831611155b15611e0957611e0060f883901c60206123d6565b60f81b92915050565b5090565b6000611574828461240f565b600082601f830112611e2a57600080fd5b81516020611e3f611e3a8361239a565b612369565b80838252828201915082860187848660051b8901011115611e5f57600080fd5b60005b85811015611e7e57815184529284019290840190600101611e62565b5090979650505050505050565b600082601f830112611e9c57600080fd5b813567ffffffffffffffff811115611eb657611eb661253a565b611ec9601f8201601f1916602001612369565b818152846020838601011115611ede57600080fd5b816020850160208301376000918101602001919091529392505050565b600060208284031215611f0d57600080fd5b813561157481612550565b60008060408385031215611f2b57600080fd5b825167ffffffffffffffff80821115611f4357600080fd5b818501915085601f830112611f5757600080fd5b81516020611f67611e3a8361239a565b8083825282820191508286018a848660051b8901011115611f8757600080fd5b600096505b84871015611fb3578051611f9f81612550565b835260019690960195918301918301611f8c565b5091880151919650909350505080821115611fcd57600080fd5b50611fda85828601611e19565b9150509250929050565b600060208284031215611ff657600080fd5b8151801515811461157457600080fd5b60006020828403121561201857600080fd5b813567ffffffffffffffff81111561202f57600080fd5b61156084828501611e8b565b600080600080600060a0868803121561205357600080fd5b853567ffffffffffffffff81111561206a57600080fd5b61207688828901611e8b565b955050602086013561208781612550565b94979496505050506040830135926060810135926080909101359150565b6000602082840312156120b757600080fd5b5035919050565b600080604083850312156120d157600080fd5b8235915060208301356120e381612550565b809150509250929050565b600080600080600060a0868803121561210657600080fd5b85359450602086013561208781612550565b6000806040838503121561212b57600080fd5b50508035926020909101359150565b60008151808452612152816020860160208601612468565b601f01601f19169290920160200192915050565b60008251612178818460208701612468565b9190910192915050565b60008451612194818460208901612468565b8083019050602f60f81b80825285516121b4816001850160208a01612468565b600192019182015283516121cf816002840160208801612468565b0160020195945050505050565b868152600086516121f4816020850160208b01612468565b60609690961b6bffffffffffffffffffffffff1916602092909601918201959095526034810193909352605483019190915260748201526094019392505050565b602081526000611574602083018461213a565b60a08152600061225b60a083018861213a565b6001600160a01b039690961660208301525060408101939093526060830191909152608090910152919050565b60a08152600061229b60a083018861213a565b6001600160a01b039687166020840152949095166040820152606081019290925260809091015292915050565b60c0815260006122db60c083018961213a565b6001600160a01b039788166020840152604083019690965250929094166060830152608082015260a00191909152919050565b604081526000612321604083018561213a565b905082151560208301529392505050565b6020808252601f908201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604082015260600190565b604051601f8201601f1916810167ffffffffffffffff811182821017156123925761239261253a565b604052919050565b600067ffffffffffffffff8211156123b4576123b461253a565b5060051b60200190565b600082198211156123d1576123d16124f8565b500190565b600060ff821660ff84168060ff038211156123f3576123f36124f8565b019392505050565b60008261240a5761240a61250e565b500490565b6000816000190483118215151615612429576124296124f8565b500290565b600082821015612440576124406124f8565b500390565b600060ff821660ff84168082101561245f5761245f6124f8565b90039392505050565b60005b8381101561248357818101518382015260200161246b565b838111156115055750506000910152565b600181811c908216806124a857607f821691505b60208210811415610e3757634e487b7160e01b600052602260045260246000fd5b60006000198214156124dd576124dd6124f8565b5060010190565b6000826124f3576124f361250e565b500690565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052601260045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b038116811461256557600080fd5b5056fe307830303030303030303030303030303030303030303030303030303030303030303030303030303030a2646970667358221220a3a8f8618ba3e77ae3e8ffd84b697d6d475f1597d3dc05befc2ac9ad15478ffb64736f6c63430008070033",
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

// SubchainID is a free data retrieval call binding the contract method 0xf8e8f7aa.
//
// Solidity: function subchainID() view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankCaller) SubchainID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TFuelTokenBank.contract.Call(opts, &out, "subchainID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SubchainID is a free data retrieval call binding the contract method 0xf8e8f7aa.
//
// Solidity: function subchainID() view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankSession) SubchainID() (*big.Int, error) {
	return _TFuelTokenBank.Contract.SubchainID(&_TFuelTokenBank.CallOpts)
}

// SubchainID is a free data retrieval call binding the contract method 0xf8e8f7aa.
//
// Solidity: function subchainID() view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankCallerSession) SubchainID() (*big.Int, error) {
	return _TFuelTokenBank.Contract.SubchainID(&_TFuelTokenBank.CallOpts)
}

// Test is a free data retrieval call binding the contract method 0xf8a8fd6d.
//
// Solidity: function test() view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankCaller) Test(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TFuelTokenBank.contract.Call(opts, &out, "test")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Test is a free data retrieval call binding the contract method 0xf8a8fd6d.
//
// Solidity: function test() view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankSession) Test() (*big.Int, error) {
	return _TFuelTokenBank.Contract.Test(&_TFuelTokenBank.CallOpts)
}

// Test is a free data retrieval call binding the contract method 0xf8a8fd6d.
//
// Solidity: function test() view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankCallerSession) Test() (*big.Int, error) {
	return _TFuelTokenBank.Contract.Test(&_TFuelTokenBank.CallOpts)
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

// TFuelTokenBankHasIterator is returned from FilterHas and is used to iterate over the raw logs and unpacked data for Has events raised by the TFuelTokenBank contract.
type TFuelTokenBankHasIterator struct {
	Event *TFuelTokenBankHas // Event containing the contract specifics and raw log

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
func (it *TFuelTokenBankHasIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TFuelTokenBankHas)
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
		it.Event = new(TFuelTokenBankHas)
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
func (it *TFuelTokenBankHasIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TFuelTokenBankHasIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TFuelTokenBankHas represents a Has event raised by the TFuelTokenBank contract.
type TFuelTokenBankHas struct {
	X   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterHas is a free log retrieval operation binding the contract event 0xcccf7a8e37b5b5f22d02b4793041d83cdfe84b965b1a0f07e23dfc2dc81520e9.
//
// Solidity: event has(uint256 x)
func (_TFuelTokenBank *TFuelTokenBankFilterer) FilterHas(opts *bind.FilterOpts) (*TFuelTokenBankHasIterator, error) {

	logs, sub, err := _TFuelTokenBank.contract.FilterLogs(opts, "has")
	if err != nil {
		return nil, err
	}
	return &TFuelTokenBankHasIterator{contract: _TFuelTokenBank.contract, event: "has", logs: logs, sub: sub}, nil
}

// WatchHas is a free log subscription operation binding the contract event 0xcccf7a8e37b5b5f22d02b4793041d83cdfe84b965b1a0f07e23dfc2dc81520e9.
//
// Solidity: event has(uint256 x)
func (_TFuelTokenBank *TFuelTokenBankFilterer) WatchHas(opts *bind.WatchOpts, sink chan<- *TFuelTokenBankHas) (event.Subscription, error) {

	logs, sub, err := _TFuelTokenBank.contract.WatchLogs(opts, "has")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TFuelTokenBankHas)
				if err := _TFuelTokenBank.contract.UnpackLog(event, "has", log); err != nil {
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

// ParseHas is a log parse operation binding the contract event 0xcccf7a8e37b5b5f22d02b4793041d83cdfe84b965b1a0f07e23dfc2dc81520e9.
//
// Solidity: event has(uint256 x)
func (_TFuelTokenBank *TFuelTokenBankFilterer) ParseHas(log types.Log) (*TFuelTokenBankHas, error) {
	event := new(TFuelTokenBankHas)
	if err := _TFuelTokenBank.contract.UnpackLog(event, "has", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
