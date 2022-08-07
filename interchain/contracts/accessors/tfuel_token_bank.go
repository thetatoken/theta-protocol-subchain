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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"mainchainID_\",\"type\":\"uint256\"},{\"internalType\":\"contractChainRegistrar\",\"name\":\"chainRegistrar_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sourceChainTokenSender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"targetChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"targetChainVoucherReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenLockNonce\",\"type\":\"uint256\"}],\"name\":\"TFuelTokenLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"targetChainTokenReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unlockedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainVoucherBurnNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenUnlockNonce\",\"type\":\"uint256\"}],\"name\":\"TFuelTokenUnlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sourceChainVoucherOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"targetChainTokenReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"burnedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"voucherBurnNonce\",\"type\":\"uint256\"}],\"name\":\"TFuelVoucherBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"targetChainVoucherReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainTokenLockNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"voucherMintNonce\",\"type\":\"uint256\"}],\"name\":\"TFuelVoucherMinted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allDenoms\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allVouchers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"denomToVoucherLookup\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voucherAddress\",\"type\":\"address\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voucherContractAddr\",\"type\":\"address\"}],\"name\":\"getDenom\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"getMaxProcessedTokenLockNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"getMaxProcessedVoucherBurnNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"name\":\"getVoucher\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOnMainchain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mainchainID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"subchainID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenLockNonceMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"tokenLockVotingRecords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dynasty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumlatedShares\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenUnlockNonceMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"totalLockedAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"voucherAddressToDenomLookup\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"voucherBurnNonceMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"voucherBurnVotingRecords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dynasty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumlatedShares\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"voucherMintNonceMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"targetChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"targetChainVoucherReceiver\",\"type\":\"address\"}],\"name\":\"lockTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"targetChainVoucherReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mintedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dynasty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainTokenLockNonce\",\"type\":\"uint256\"}],\"name\":\"mintVouchers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targetChainTokenReceiver\",\"type\":\"address\"}],\"name\":\"burnVouchers\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainID\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"targetChainTokenReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"unlockAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dynasty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainVoucherBurnNonce\",\"type\":\"uint256\"}],\"name\":\"unlockTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200263638038062002636833981016040819052620000349162000064565b6001600081905591909155600380546001600160a01b0319166001600160a01b03909216919091179055620000a3565b600080604083850312156200007857600080fd5b825160208401519092506001600160a01b03811681146200009857600080fd5b809150509250929050565b61258380620000b36000396000f3fe60806040526004361061014b5760003560e01c80637d0fb00d116100b6578063ebda99621161006f578063ebda99621461048c578063f6a3d24e146104ac578063f8e8f7aa146104e8578063f95627ac146104fe578063feaff0521461052b578063ff248a441461056a57600080fd5b80637d0fb00d146103bf5780638883931e146103d2578063a2cc6981146103ff578063aa68acde1461041f578063ca20756914610432578063ccf187c71461045f57600080fd5b806327ca4df11161010857806327ca4df1146102b85780634250863b146102f0578063588b14081461030a57806360569b5e14610337578063740cb7f814610365578063766f8fb01461039257600080fd5b8063073b9502146101505780631527b14d1461017957806319fd1a11146101e55780631a0483d3146102125780631eb7873714610234578063261a323e14610288575b600080fd5b34801561015c57600080fd5b5061016660015481565b6040519081526020015b60405180910390f35b34801561018557600080fd5b506101c6610194366004611fc1565b8051602081830181018051600c825292820191909301209152546001600160a01b03811690600160a01b900460ff1682565b604080516001600160a01b039093168352901515602083015201610170565b3480156101f157600080fd5b50610166610200366004612060565b60116020526000908152604090205481565b34801561021e57600080fd5b5061023261022d366004611ff6565b61057d565b005b34801561024057600080fd5b5061027361024f3660046120d3565b600a6020908152600092835260408084209091529082529020805460029091015482565b60408051928352602083019190915201610170565b34801561029457600080fd5b506102a86102a3366004611fc1565b610762565b6040519015158152602001610170565b3480156102c457600080fd5b506102d86102d3366004612060565b6107a2565b6040516001600160a01b039091168152602001610170565b3480156102fc57600080fd5b506010546102a89060ff1681565b34801561031657600080fd5b5061032a610325366004612060565b6107cc565b60405161017091906121f0565b34801561034357600080fd5b50610357610352366004611eb6565b610878565b6040516101709291906122c9565b34801561037157600080fd5b50610166610380366004612060565b60076020526000908152604090205481565b34801561039e57600080fd5b506101666103ad366004612060565b60009081526009602052604090205490565b6102326103cd366004611eb6565b61091f565b3480156103de57600080fd5b506101666103ed366004612060565b60046020526000908152604090205481565b34801561040b57600080fd5b506102d861041a366004611fc1565b610a77565b61023261042d366004612079565b610ae8565b34801561043e57600080fd5b5061016661044d366004612060565b60066020526000908152604090205481565b34801561046b57600080fd5b5061016661047a366004612060565b60056020526000908152604090205481565b34801561049857600080fd5b5061032a6104a7366004611eb6565b610d16565b3480156104b857600080fd5b506102a86104c7366004611eb6565b6001600160a01b03166000908152600d602052604090206001015460ff1690565b3480156104f457600080fd5b5061016660025481565b34801561050a57600080fd5b50610166610519366004612060565b60009081526008602052604090205490565b34801561053757600080fd5b506102736105463660046120d3565b600b6020908152600092835260408084209091529082529020805460029091015482565b6102326105783660046120a9565b610e0d565b600260005414156105a95760405162461bcd60e51b81526004016105a0906122ed565b60405180910390fd5b600260005560015446908114156106285760405162461bcd60e51b815260206004820152603e60248201527f544675656c546f6b656e42616e6b2e6d696e74566f756368657273282920636160448201527f6e206f6e6c792062652063616c6c6564206f6e206120737562636861696e000060648201526084016105a0565b6000610633876110a6565b9050600080610641836110b7565b915091508061069e5760405162461bcd60e51b8152602060048201526024808201527f4661696c656420746f206578747261637420636861696e49442066726f6d2064604482015263656e6f6d60e01b60648201526084016105a0565b600082848a8a8a8a6040516020016106bb96959493929190612197565b60405160208183030381529060405280519060200120905060006106e28489848a3361113f565b90508015610750576106f48a8a61115c565b6106fd4661133b565b46600090815260076020526040908190205490517f80742bd15a2c8c4ad5d395bcf577073110e52f0c73bf980dfa9453c1d8c354e590610746908e908e908e908d908790612203565b60405180910390a1505b50506001600055505050505050505050565b60008061076e836110a6565b9050600c816040516107809190612121565b9081526040519081900360200190205460ff600160a01b909104169392505050565b600e81815481106107b257600080fd5b6000918252602090912001546001600160a01b0316905081565b600f81815481106107dc57600080fd5b9060005260206000200160009150905080546107f79061244f565b80601f01602080910402602001604051908101604052809291908181526020018280546108239061244f565b80156108705780601f1061084557610100808354040283529160200191610870565b820191906000526020600020905b81548152906001019060200180831161085357829003601f168201915b505050505081565b600d602052600090815260409020805481906108939061244f565b80601f01602080910402602001604051908101604052809291908181526020018280546108bf9061244f565b801561090c5780601f106108e15761010080835404028352916020019161090c565b820191906000526020600020905b8154815290600101906020018083116108ef57829003601f168201915b5050506001909301549192505060ff1682565b600260005414156109425760405162461bcd60e51b81526004016105a0906122ed565b600260005560015446908114156109c15760405162461bcd60e51b815260206004820152603e60248201527f544675656c546f6b656e42616e6b2e6275726e566f756368657273282920636160448201527f6e206f6e6c792062652063616c6c6564206f6e206120737562636861696e000060648201526084016105a0565b33346109cc81611362565b6109d5466114c6565b6000610a15600154604051806040016040528060018152602001600360fc1b8152506040518060600160405280602a8152602001612524602a91396114e5565b4660009081526006602052604090819020549051919250907f40f1d475c2aa44f5c23193fab26a64d6aa4e09ab51898b10a3036baf82398ea190610a6290849087908a9088908790612243565b60405180910390a15050600160005550505050565b600080610a83836110a6565b90506000600c82604051610a979190612121565b908152604080516020928190038301812081830190925290546001600160a01b0381168252600160a01b900460ff16158015928201929092529150610ade57519392505050565b5060009392505050565b60026000541415610b0b5760405162461bcd60e51b81526004016105a0906122ed565b60026000556001544614610b89576040805162461bcd60e51b81526020600482015260248101919091527f544675656c546f6b656e42616e6b2e6c6f636b546f6b656e7328292063616e2060448201527f6f6e6c792062652063616c6c6564206f6e20746865206d61696e20636861696e60648201526084016105a0565b6003546040516343b71f0560e01b8152600481018490526001600160a01b03909116906343b71f0590602401602060405180830381600087803b158015610bcf57600080fd5b505af1158015610be3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c079190611f9f565b610c535760405162461bcd60e51b815260206004820181905260248201527f746172676574436861696e4944206e6f7420796574207265676973746572656460448201526064016105a0565b60008281526011602052604090205434903390610c709083611523565b600085815260116020526040902055610c8884611536565b6000610cc8600154604051806040016040528060018152602001600360fc1b8152506040518060600160405280602a8152602001612524602a91396114e5565b60008681526004602052604090819020549051919250907fee1ecc2b21aa613cc77cd44823a68ef1168ce1f40c2eac1d68690baf955fdbd190610a6290849086908a908a908a908890612283565b6001600160a01b0381166000908152600d60205260408082208151808301909252805460609392919082908290610d4c9061244f565b80601f0160208091040260200160405190810160405280929190818152602001828054610d789061244f565b8015610dc55780601f10610d9a57610100808354040283529160200191610dc5565b820191906000526020600020905b815481529060010190602001808311610da857829003601f168201915b50505091835250506001919091015460ff16151560209182015281015190915015610df1575192915050565b5050604080516020810190915260008152919050565b50919050565b60026000541415610e305760405162461bcd60e51b81526004016105a0906122ed565b60026000556001544614610eb75760405162461bcd60e51b815260206004820152604260248201527f544675656c546f6b656e42616e6b2e756e6c6f636b546f6b656e73282920636160448201527f6e206f6e6c792062652063616c6c6564206f6e20746865206d61696e2063686160648201526134b760f11b608482015260a4016105a0565b600085815260116020526040902054831115610f295760405162461bcd60e51b815260206004820152602b60248201527f43616e6e6f7420756e6c6f636b207468652072657175657374656420616d6f7560448201526a1b9d081bd98815119d595b60aa1b60648201526084016105a0565b6040805160208101879052908101849052606085811b6bffffffffffffffffffffffff191690820152607481018390526094810182905260009060b40160408051601f19818403018152919052805160208201209091506000610f8f8886848733611555565b9050801561109757600088815260116020526040902054610fb09087611568565b6000898152601160205260408082209290925590516001600160a01b0389169188156108fc02918991818181858888f19350505050158015610ff6573d6000803e3d6000fd5b5061100088611574565b6000611040600154604051806040016040528060018152602001600360fc1b8152506040518060600160405280602a8152602001612524602a91396114e5565b60008a81526005602052604090819020549051919250907f5ea3a5ca7f54881fdd7781894d69709e11027910f35647f9d4cc14e6872b6f729061108c9084908c908c908b908790612203565b60405180910390a150505b50506001600055505050505050565b60606110b182611593565b92915050565b6000806000806110cd85602f60f81b600161160d565b91509150806110e3575060009485945092505050565b6000806110f28760008661169d565b915091508061110a5750600096879650945050505050565b600080611116846117aa565b915091508061113057506000988998509650505050505050565b50976001975095505050505050565b600061115286868686866008600a61187c565b9695505050505050565b60408051603480825260608201909252600091602082018180368337019050509050816000805b60148110156111eb578560601b81601481106111a1576111a16124df565b1a60f81b84836111b081612484565b9450815181106111c2576111c26124df565b60200101906001600160f81b031916908160001a905350806111e381612484565b915050611183565b5060005b60208110156112545782816020811061120a5761120a6124df565b1a60f81b848361121981612484565b94508151811061122b5761122b6124df565b60200101906001600160f81b031916908160001a9053508061124c81612484565b9150506111ef565b600060b66001600160a01b03168560405161126f9190612121565b6000604051808303816000865af19150503d80600081146112ac576040519150601f19603f3d011682016040523d82523d6000602084013e6112b1565b606091505b50509050806113325760405162461bcd60e51b815260206004820152604160248201527f544675656c546f6b656e42616e6b2e5f6d696e74544675656c566f756368657260448201527f28293a206661696c656420746f206d696e7420544675656c20766f75636865726064820152607360f81b608482015260a4016105a0565b50505050505050565b600081815260076020526040812080546001929061135a908490612379565b909155505050565b6040805160208082528183019092526000916020820181803683370190505090508160005b60208110156113e1578181602081106113a2576113a26124df565b1a60f81b8382815181106113b8576113b86124df565b60200101906001600160f81b031916908160001a905350806113d981612484565b915050611387565b50600060b76001600160a01b0316836040516113fd9190612121565b6000604051808303816000865af19150503d806000811461143a576040519150601f19603f3d011682016040523d82523d6000602084013e61143f565b606091505b50509050806114c05760405162461bcd60e51b815260206004820152604160248201527f544675656c546f6b656e42616e6b2e5f6275726e544675656c566f756368657260448201527f28293a206661696c656420746f206275726e20544675656c20766f75636865726064820152607360f81b608482015260a4016105a0565b50505050565b600081815260066020526040812080546001929061135a908490612379565b606061151b6114f385611c7b565b84846040516020016115079392919061213d565b6040516020818303038152906040526110a6565b949350505050565b600061152f8284612379565b9392505050565b600081815260046020526040812080546001929061135a908490612379565b600061115286868686866009600b61187c565b600061152f82846123e9565b600081815260056020526040812080546001929061135a908490612379565b60608160005b8151811015611606576115cb8282815181106115b7576115b76124df565b01602001516001600160f81b031916611d79565b8282815181106115dd576115dd6124df565b60200101906001600160f81b031916908160001a905350806115fe81612484565b915050611599565b5092915050565b82516000908190859082805b8281101561168957876001600160f81b03191684828151811061163e5761163e6124df565b01602001516001600160f81b03191614156116775761165e600183612379565b9150868214156116775794506001935061169592505050565b8061168181612484565b915050611619565b50600080945094505050505b935093915050565b8251606090600090848410156116c757505060408051602081019091526000808252909150611695565b808411156116e957505060408051602081019091526000808252909150611695565b8560006116f68688611568565b67ffffffffffffffff81111561170e5761170e6124f5565b6040519080825280601f01601f191660200182016040528015611738576020820181803683370190505b509050865b8681101561113057828181518110611757576117576124df565b01602001516001600160f81b03191682611771838b611568565b81518110611781576117816124df565b60200101906001600160f81b031916908160001a905350806117a281612484565b91505061173d565b80516000908190839082805b8281101561186f5760308482815181106117d2576117d26124df565b016020015160f81c10801590611802575060398482815181106117f7576117f76124df565b016020015160f81c11155b1561184e5761181282600a611dc8565b9150611847603085838151811061182b5761182b6124df565b016020015161183d919060f81c612400565b839060ff16611523565b915061185d565b50600096879650945050505050565b8061186781612484565b9150506117b6565b5095600195509350505050565b600060015488148061190757506003546040516343b71f0560e01b8152600481018a90526001600160a01b03909116906343b71f0590602401602060405180830381600087803b1580156118cf57600080fd5b505af11580156118e3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119079190611f9f565b6119455760405162461bcd60e51b815260206004820152600f60248201526e125b9d985b1a590818da185a5b9251608a1b60448201526064016105a0565b60008881526020849052604090205461195f906001612379565b851461196d57506000611c70565b600088815260208381526040808320898452909152812060015460609081908c9081141561199a57506002545b6003546040516343f27e4560e01b815260048101839052602481018e90526001600160a01b03909116906343f27e4590604401600060405180830381600087803b1580156119e757600080fd5b505af11580156119fb573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611a239190810190611ed3565b9350915060005b8251811015611b8057896001600160a01b0316838281518110611a4f57611a4f6124df565b60200260200101516001600160a01b031614611a6a57611b6e565b6001955060005b6001860154811015611b0b57856001018181548110611a9257611a926124df565b6000918252602090912001546001600160a01b038c811691161415611af95760405162461bcd60e51b815260206004820152601c60248201527f546869732076616c696461746f7220616c726561647920766f7465640000000060448201526064016105a0565b80611b0381612484565b915050611a71565b508c85556001808601805491820181556000908152602090200180546001600160a01b031916331790558351611b6890859083908110611b4d57611b4d6124df565b6020026020010151866002015461152390919063ffffffff16565b60028601555b80611b7881612484565b915050611a2a565b50505082611bc25760405162461bcd60e51b815260206004820152600f60248201526e2737ba1030903b30b634b230ba37b960891b60448201526064016105a0565b6000805b8251811015611c0f57611bfb838281518110611be457611be46124df565b60200260200101518361152390919063ffffffff16565b915080611c0781612484565b915050611bc6565b50611c1b816002611dc8565b6002840154611c2b906003611dc8565b10611c675760008c815260208890526040902054611c4a906001612379565b60008d8152602089905260409020555060019350611c7092505050565b60009450505050505b979650505050505050565b606081611c9f5750506040805180820190915260018152600360fc1b602082015290565b8160005b8115611cc95780611cb381612484565b9150611cc29050600a836123b6565b9150611ca3565b60008167ffffffffffffffff811115611ce457611ce46124f5565b6040519080825280601f01601f191660200182016040528015611d0e576020820181803683370190505b5090505b841561151b57611d236001836123e9565b9150611d30600a8661249f565b611d3b906030612379565b60f81b818381518110611d5057611d506124df565b60200101906001600160f81b031916908160001a905350611d72600a866123b6565b9450611d12565b6000604160f81b6001600160f81b0319831610801590611da75750602d60f91b6001600160f81b0319831611155b15611dc457611dbb60f883901c6020612391565b60f81b92915050565b5090565b600061152f82846123ca565b600082601f830112611de557600080fd5b81516020611dfa611df583612355565b612324565b80838252828201915082860187848660051b8901011115611e1a57600080fd5b60005b85811015611e3957815184529284019290840190600101611e1d565b5090979650505050505050565b600082601f830112611e5757600080fd5b813567ffffffffffffffff811115611e7157611e716124f5565b611e84601f8201601f1916602001612324565b818152846020838601011115611e9957600080fd5b816020850160208301376000918101602001919091529392505050565b600060208284031215611ec857600080fd5b813561152f8161250b565b60008060408385031215611ee657600080fd5b825167ffffffffffffffff80821115611efe57600080fd5b818501915085601f830112611f1257600080fd5b81516020611f22611df583612355565b8083825282820191508286018a848660051b8901011115611f4257600080fd5b600096505b84871015611f6e578051611f5a8161250b565b835260019690960195918301918301611f47565b5091880151919650909350505080821115611f8857600080fd5b50611f9585828601611dd4565b9150509250929050565b600060208284031215611fb157600080fd5b8151801515811461152f57600080fd5b600060208284031215611fd357600080fd5b813567ffffffffffffffff811115611fea57600080fd5b61151b84828501611e46565b600080600080600060a0868803121561200e57600080fd5b853567ffffffffffffffff81111561202557600080fd5b61203188828901611e46565b95505060208601356120428161250b565b94979496505050506040830135926060810135926080909101359150565b60006020828403121561207257600080fd5b5035919050565b6000806040838503121561208c57600080fd5b82359150602083013561209e8161250b565b809150509250929050565b600080600080600060a086880312156120c157600080fd5b8535945060208601356120428161250b565b600080604083850312156120e657600080fd5b50508035926020909101359150565b6000815180845261210d816020860160208601612423565b601f01601f19169290920160200192915050565b60008251612133818460208701612423565b9190910192915050565b6000845161214f818460208901612423565b8083019050602f60f81b808252855161216f816001850160208a01612423565b6001920191820152835161218a816002840160208801612423565b0160020195945050505050565b868152600086516121af816020850160208b01612423565b60609690961b6bffffffffffffffffffffffff1916602092909601918201959095526034810193909352605483019190915260748201526094019392505050565b60208152600061152f60208301846120f5565b60a08152600061221660a08301886120f5565b6001600160a01b039690961660208301525060408101939093526060830191909152608090910152919050565b60a08152600061225660a08301886120f5565b6001600160a01b039687166020840152949095166040820152606081019290925260809091015292915050565b60c08152600061229660c08301896120f5565b6001600160a01b039788166020840152604083019690965250929094166060830152608082015260a00191909152919050565b6040815260006122dc60408301856120f5565b905082151560208301529392505050565b6020808252601f908201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604082015260600190565b604051601f8201601f1916810167ffffffffffffffff8111828210171561234d5761234d6124f5565b604052919050565b600067ffffffffffffffff82111561236f5761236f6124f5565b5060051b60200190565b6000821982111561238c5761238c6124b3565b500190565b600060ff821660ff84168060ff038211156123ae576123ae6124b3565b019392505050565b6000826123c5576123c56124c9565b500490565b60008160001904831182151516156123e4576123e46124b3565b500290565b6000828210156123fb576123fb6124b3565b500390565b600060ff821660ff84168082101561241a5761241a6124b3565b90039392505050565b60005b8381101561243e578181015183820152602001612426565b838111156114c05750506000910152565b600181811c9082168061246357607f821691505b60208210811415610e0757634e487b7160e01b600052602260045260246000fd5b6000600019821415612498576124986124b3565b5060010190565b6000826124ae576124ae6124c9565b500690565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052601260045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b038116811461252057600080fd5b5056fe307830303030303030303030303030303030303030303030303030303030303030303030303030303030a26469706673582212201ed55750d70bec3fd257a2e59a6546840369cd00245e76a4d2c01210de8ffd2e64736f6c63430008070033",
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
