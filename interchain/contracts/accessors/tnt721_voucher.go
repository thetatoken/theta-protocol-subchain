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

// TNT721VoucherContractMetaData contains all meta data concerning the TNT721VoucherContract contract.
var TNT721VoucherContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"demon_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"UpdateOwner\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"updateOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voucherReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenUri\",\"type\":\"string\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voucherOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"denom\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContractOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001a0638038062001a0683398101604081905262000034916200020c565b83828281600090805190602001906200004f929190620000af565b50805162000065906001906020840190620000af565b5050600680546001600160a01b0319166001600160a01b0393909316929092179091555082516200009e906007906020860190620000af565b505060006008555062000311915050565b828054620000bd90620002be565b90600052602060002090601f016020900481019282620000e157600085556200012c565b82601f10620000fc57805160ff19168380011785556200012c565b828001600101855582156200012c579182015b828111156200012c5782518255916020019190600101906200010f565b506200013a9291506200013e565b5090565b5b808211156200013a57600081556001016200013f565b600082601f8301126200016757600080fd5b81516001600160401b0380821115620001845762000184620002fb565b604051601f8301601f19908116603f01168101908282118183101715620001af57620001af620002fb565b81604052838152602092508683858801011115620001cc57600080fd5b600091505b83821015620001f05785820183015181830184015290820190620001d1565b83821115620002025760008385830101525b9695505050505050565b600080600080608085870312156200022357600080fd5b84516001600160a01b03811681146200023b57600080fd5b60208601519094506001600160401b03808211156200025957600080fd5b620002678883890162000155565b945060408701519150808211156200027e57600080fd5b6200028c8883890162000155565b93506060870151915080821115620002a357600080fd5b50620002b28782880162000155565b91505092959194509250565b600181811c90821680620002d357607f821691505b60208210811415620002f557634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052604160045260246000fd5b6116e580620003216000396000f3fe608060405234801561001057600080fd5b506004361061012c5760003560e01c8063880cdc31116100ad578063b88d4fde11610071578063b88d4fde14610271578063c370b04214610284578063c87b56dd1461028c578063d3fc98641461029f578063e985e9c5146102b257600080fd5b8063880cdc311461021d5780638da5cb5b1461023057806395d89b41146102435780639dc29fac1461024b578063a22cb4651461025e57600080fd5b806323b872dd116100f457806323b872dd146101c057806342842e0e146101d3578063442890d5146101e65780636352211e146101f757806370a082311461020a57600080fd5b806301ffc9a71461013157806306fdde0314610159578063081812fc1461016e578063095ea7b31461019957806318160ddd146101ae575b600080fd5b61014461013f36600461143c565b6102ee565b60405190151581526020015b60405180910390f35b610161610340565b6040516101509190611519565b61018161017c366004611476565b6103d2565b6040516001600160a01b039091168152602001610150565b6101ac6101a73660046113a7565b6103f9565b005b6008545b604051908152602001610150565b6101ac6101ce3660046112b3565b610514565b6101ac6101e13660046112b3565b610545565b6006546001600160a01b0316610181565b610181610205366004611476565b610560565b6101b261021836600461125e565b6105c0565b6101ac61022b36600461125e565b610646565b600654610181906001600160a01b031681565b6101616106d9565b6101ac6102593660046113a7565b6106e8565b6101ac61026c36600461136b565b61086c565b6101ac61027f3660046112ef565b61087b565b6101616108b3565b61016161029a366004611476565b6108c2565b6101ac6102ad3660046113d1565b610964565b6101446102c0366004611280565b6001600160a01b03918216600090815260056020908152604080832093909416825291909152205460ff1690565b60006001600160e01b031982166380ac58cd60e01b148061031f57506001600160e01b03198216635b5e139f60e01b145b8061033a57506301ffc9a760e01b6001600160e01b03198316145b92915050565b60606000805461034f90611632565b80601f016020809104026020016040519081016040528092919081815260200182805461037b90611632565b80156103c85780601f1061039d576101008083540402835291602001916103c8565b820191906000526020600020905b8154815290600101906020018083116103ab57829003601f168201915b5050505050905090565b60006103dd826109d5565b506000908152600460205260409020546001600160a01b031690565b600061040482610560565b9050806001600160a01b0316836001600160a01b031614156104775760405162461bcd60e51b815260206004820152602160248201527f4552433732313a20617070726f76616c20746f2063757272656e74206f776e656044820152603960f91b60648201526084015b60405180910390fd5b336001600160a01b0382161480610493575061049381336102c0565b6105055760405162461bcd60e51b815260206004820152603e60248201527f4552433732313a20617070726f76652063616c6c6572206973206e6f7420746f60448201527f6b656e206f776e6572206e6f7220617070726f76656420666f7220616c6c0000606482015260840161046e565b61050f8383610a37565b505050565b61051e3382610aa5565b61053a5760405162461bcd60e51b815260040161046e906115b5565b61050f838383610b24565b61050f8383836040518060200160405280600081525061087b565b6000818152600260205260408120546001600160a01b03168061033a5760405162461bcd60e51b8152602060048201526018602482015277115490cdcc8c4e881a5b9d985b1a59081d1bdad95b88125160421b604482015260640161046e565b60006001600160a01b03821661062a5760405162461bcd60e51b815260206004820152602960248201527f4552433732313a2061646472657373207a65726f206973206e6f7420612076616044820152683634b21037bbb732b960b91b606482015260840161046e565b506001600160a01b031660009081526003602052604090205490565b6006546001600160a01b031633146106705760405162461bcd60e51b815260040161046e9061157e565b600654604080516001600160a01b03928316815291831660208301527fe2c7d1c4da37855e682bde14f17826d185497973b73fba7554daa6da467058d9910160405180910390a1600680546001600160a01b0319166001600160a01b0392909216919091179055565b60606001805461034f90611632565b6006546001600160a01b031633146107125760405162461bcd60e51b815260040161046e9061157e565b6000600854116107575760405162461bcd60e51b815260206004820152601060248201526f3737903a37b5b2b7103a3790313ab93760811b604482015260640161046e565b600061076282610560565b9050826001600160a01b0316816001600160a01b0316146107bb5760405162461bcd60e51b815260206004820152601360248201527237b7363c9037bbb732b91031b0b710313ab93760691b604482015260640161046e565b6006546001600160a01b03166107d18184610aa5565b61082e5760405162461bcd60e51b815260206004820152602860248201527f566f7563686572206f776e657220646964206e6f7420617070726f766520746f60448201526735b2b710313ab93760c11b606482015260840161046e565b61083783610cc0565b600083815260096020526040812061084e916110f9565b600160086000828254610861919061161b565b909155505050505050565b610877338383610d5b565b5050565b6108853383610aa5565b6108a15760405162461bcd60e51b815260040161046e906115b5565b6108ad84848484610e2a565b50505050565b60606007805461034f90611632565b60008181526009602052604090208054606091906108df90611632565b80601f016020809104026020016040519081016040528092919081815260200182805461090b90611632565b80156109585780601f1061092d57610100808354040283529160200191610958565b820191906000526020600020905b81548152906001019060200180831161093b57829003601f168201915b50505050509050919050565b6006546001600160a01b0316331461098e5760405162461bcd60e51b815260040161046e9061157e565b6109988383610e5d565b600082815260096020908152604090912082516109b792840190611133565b506001600860008282546109cb9190611603565b9091555050505050565b6000818152600260205260409020546001600160a01b0316610a345760405162461bcd60e51b8152602060048201526018602482015277115490cdcc8c4e881a5b9d985b1a59081d1bdad95b88125160421b604482015260640161046e565b50565b600081815260046020526040902080546001600160a01b0319166001600160a01b0384169081179091558190610a6c82610560565b6001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b600080610ab183610560565b9050806001600160a01b0316846001600160a01b03161480610af857506001600160a01b0380821660009081526005602090815260408083209388168352929052205460ff165b80610b1c5750836001600160a01b0316610b11846103d2565b6001600160a01b0316145b949350505050565b826001600160a01b0316610b3782610560565b6001600160a01b031614610b9b5760405162461bcd60e51b815260206004820152602560248201527f4552433732313a207472616e736665722066726f6d20696e636f72726563742060448201526437bbb732b960d91b606482015260840161046e565b6001600160a01b038216610bfd5760405162461bcd60e51b8152602060048201526024808201527f4552433732313a207472616e7366657220746f20746865207a65726f206164646044820152637265737360e01b606482015260840161046e565b610c08600082610a37565b6001600160a01b0383166000908152600360205260408120805460019290610c3190849061161b565b90915550506001600160a01b0382166000908152600360205260408120805460019290610c5f908490611603565b909155505060008181526002602052604080822080546001600160a01b0319166001600160a01b0386811691821790925591518493918716917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b6000610ccb82610560565b9050610cd8600083610a37565b6001600160a01b0381166000908152600360205260408120805460019290610d0190849061161b565b909155505060008281526002602052604080822080546001600160a01b0319169055518391906001600160a01b038416907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908390a45050565b816001600160a01b0316836001600160a01b03161415610dbd5760405162461bcd60e51b815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c657200000000000000604482015260640161046e565b6001600160a01b03838116600081815260056020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b610e35848484610b24565b610e4184848484610e77565b6108ad5760405162461bcd60e51b815260040161046e9061152c565b610877828260405180602001604052806000815250610f84565b60006001600160a01b0384163b15610f7957604051630a85bd0160e11b81526001600160a01b0385169063150b7a0290610ebb9033908990889088906004016114dc565b602060405180830381600087803b158015610ed557600080fd5b505af1925050508015610f05575060408051601f3d908101601f19168201909252610f0291810190611459565b60015b610f5f573d808015610f33576040519150601f19603f3d011682016040523d82523d6000602084013e610f38565b606091505b508051610f575760405162461bcd60e51b815260040161046e9061152c565b805181602001fd5b6001600160e01b031916630a85bd0160e11b149050610b1c565b506001949350505050565b610f8e8383610fb7565b610f9b6000848484610e77565b61050f5760405162461bcd60e51b815260040161046e9061152c565b6001600160a01b03821661100d5760405162461bcd60e51b815260206004820181905260248201527f4552433732313a206d696e7420746f20746865207a65726f2061646472657373604482015260640161046e565b6000818152600260205260409020546001600160a01b0316156110725760405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e74656400000000604482015260640161046e565b6001600160a01b038216600090815260036020526040812080546001929061109b908490611603565b909155505060008181526002602052604080822080546001600160a01b0319166001600160a01b03861690811790915590518392907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b50805461110590611632565b6000825580601f10611115575050565b601f016020900490600052602060002090810190610a3491906111b7565b82805461113f90611632565b90600052602060002090601f01602090048101928261116157600085556111a7565b82601f1061117a57805160ff19168380011785556111a7565b828001600101855582156111a7579182015b828111156111a757825182559160200191906001019061118c565b506111b39291506111b7565b5090565b5b808211156111b357600081556001016111b8565b600067ffffffffffffffff808411156111e7576111e7611683565b604051601f8501601f19908116603f0116810190828211818310171561120f5761120f611683565b8160405280935085815286868601111561122857600080fd5b858560208301376000602087830101525050509392505050565b80356001600160a01b038116811461125957600080fd5b919050565b60006020828403121561127057600080fd5b61127982611242565b9392505050565b6000806040838503121561129357600080fd5b61129c83611242565b91506112aa60208401611242565b90509250929050565b6000806000606084860312156112c857600080fd5b6112d184611242565b92506112df60208501611242565b9150604084013590509250925092565b6000806000806080858703121561130557600080fd5b61130e85611242565b935061131c60208601611242565b925060408501359150606085013567ffffffffffffffff81111561133f57600080fd5b8501601f8101871361135057600080fd5b61135f878235602084016111cc565b91505092959194509250565b6000806040838503121561137e57600080fd5b61138783611242565b91506020830135801515811461139c57600080fd5b809150509250929050565b600080604083850312156113ba57600080fd5b6113c383611242565b946020939093013593505050565b6000806000606084860312156113e657600080fd5b6113ef84611242565b925060208401359150604084013567ffffffffffffffff81111561141257600080fd5b8401601f8101861361142357600080fd5b611432868235602084016111cc565b9150509250925092565b60006020828403121561144e57600080fd5b813561127981611699565b60006020828403121561146b57600080fd5b815161127981611699565b60006020828403121561148857600080fd5b5035919050565b6000815180845260005b818110156114b557602081850181015186830182015201611499565b818111156114c7576000602083870101525b50601f01601f19169290920160200192915050565b6001600160a01b038581168252841660208201526040810183905260806060820181905260009061150f9083018461148f565b9695505050505050565b602081526000611279602083018461148f565b60208082526032908201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560408201527131b2b4bb32b91034b6b83632b6b2b73a32b960711b606082015260800190565b6020808252601c908201527f6f6e6c79206f776e65722063616e206d616b65207468652063616c6c00000000604082015260600190565b6020808252602e908201527f4552433732313a2063616c6c6572206973206e6f7420746f6b656e206f776e6560408201526d1c881b9bdc88185c1c1c9bdd995960921b606082015260800190565b600082198211156116165761161661166d565b500190565b60008282101561162d5761162d61166d565b500390565b600181811c9082168061164657607f821691505b6020821081141561166757634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160e01b031981168114610a3457600080fdfea2646970667358221220b15d589687cd043e38e292316cde83b6be826c0bf720dbdfff0cad1ccc2b331964736f6c63430008070033",
}

// TNT721VoucherContractABI is the input ABI used to generate the binding from.
// Deprecated: Use TNT721VoucherContractMetaData.ABI instead.
var TNT721VoucherContractABI = TNT721VoucherContractMetaData.ABI

// TNT721VoucherContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TNT721VoucherContractMetaData.Bin instead.
var TNT721VoucherContractBin = TNT721VoucherContractMetaData.Bin

// DeployTNT721VoucherContract deploys a new Ethereum contract, binding an instance of TNT721VoucherContract to it.
func DeployTNT721VoucherContract(auth *bind.TransactOpts, backend bind.ContractBackend, owner_ common.Address, demon_ string, name_ string, symbol_ string) (common.Address, *types.Transaction, *TNT721VoucherContract, error) {
	parsed, err := TNT721VoucherContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TNT721VoucherContractBin), backend, owner_, demon_, name_, symbol_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TNT721VoucherContract{TNT721VoucherContractCaller: TNT721VoucherContractCaller{contract: contract}, TNT721VoucherContractTransactor: TNT721VoucherContractTransactor{contract: contract}, TNT721VoucherContractFilterer: TNT721VoucherContractFilterer{contract: contract}}, nil
}

// TNT721VoucherContract is an auto generated Go binding around an Ethereum contract.
type TNT721VoucherContract struct {
	TNT721VoucherContractCaller     // Read-only binding to the contract
	TNT721VoucherContractTransactor // Write-only binding to the contract
	TNT721VoucherContractFilterer   // Log filterer for contract events
}

// TNT721VoucherContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type TNT721VoucherContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TNT721VoucherContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TNT721VoucherContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TNT721VoucherContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TNT721VoucherContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TNT721VoucherContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TNT721VoucherContractSession struct {
	Contract     *TNT721VoucherContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TNT721VoucherContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TNT721VoucherContractCallerSession struct {
	Contract *TNT721VoucherContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// TNT721VoucherContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TNT721VoucherContractTransactorSession struct {
	Contract     *TNT721VoucherContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// TNT721VoucherContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type TNT721VoucherContractRaw struct {
	Contract *TNT721VoucherContract // Generic contract binding to access the raw methods on
}

// TNT721VoucherContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TNT721VoucherContractCallerRaw struct {
	Contract *TNT721VoucherContractCaller // Generic read-only contract binding to access the raw methods on
}

// TNT721VoucherContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TNT721VoucherContractTransactorRaw struct {
	Contract *TNT721VoucherContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTNT721VoucherContract creates a new instance of TNT721VoucherContract, bound to a specific deployed contract.
func NewTNT721VoucherContract(address common.Address, backend bind.ContractBackend) (*TNT721VoucherContract, error) {
	contract, err := bindTNT721VoucherContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TNT721VoucherContract{TNT721VoucherContractCaller: TNT721VoucherContractCaller{contract: contract}, TNT721VoucherContractTransactor: TNT721VoucherContractTransactor{contract: contract}, TNT721VoucherContractFilterer: TNT721VoucherContractFilterer{contract: contract}}, nil
}

// NewTNT721VoucherContractCaller creates a new read-only instance of TNT721VoucherContract, bound to a specific deployed contract.
func NewTNT721VoucherContractCaller(address common.Address, caller bind.ContractCaller) (*TNT721VoucherContractCaller, error) {
	contract, err := bindTNT721VoucherContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TNT721VoucherContractCaller{contract: contract}, nil
}

// NewTNT721VoucherContractTransactor creates a new write-only instance of TNT721VoucherContract, bound to a specific deployed contract.
func NewTNT721VoucherContractTransactor(address common.Address, transactor bind.ContractTransactor) (*TNT721VoucherContractTransactor, error) {
	contract, err := bindTNT721VoucherContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TNT721VoucherContractTransactor{contract: contract}, nil
}

// NewTNT721VoucherContractFilterer creates a new log filterer instance of TNT721VoucherContract, bound to a specific deployed contract.
func NewTNT721VoucherContractFilterer(address common.Address, filterer bind.ContractFilterer) (*TNT721VoucherContractFilterer, error) {
	contract, err := bindTNT721VoucherContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TNT721VoucherContractFilterer{contract: contract}, nil
}

// bindTNT721VoucherContract binds a generic wrapper to an already deployed contract.
func bindTNT721VoucherContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TNT721VoucherContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TNT721VoucherContract *TNT721VoucherContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TNT721VoucherContract.Contract.TNT721VoucherContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TNT721VoucherContract *TNT721VoucherContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TNT721VoucherContract.Contract.TNT721VoucherContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TNT721VoucherContract *TNT721VoucherContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TNT721VoucherContract.Contract.TNT721VoucherContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TNT721VoucherContract *TNT721VoucherContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TNT721VoucherContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TNT721VoucherContract *TNT721VoucherContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TNT721VoucherContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TNT721VoucherContract *TNT721VoucherContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TNT721VoucherContract.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_TNT721VoucherContract *TNT721VoucherContractCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TNT721VoucherContract.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_TNT721VoucherContract *TNT721VoucherContractSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _TNT721VoucherContract.Contract.BalanceOf(&_TNT721VoucherContract.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_TNT721VoucherContract *TNT721VoucherContractCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _TNT721VoucherContract.Contract.BalanceOf(&_TNT721VoucherContract.CallOpts, owner)
}

// Denom is a free data retrieval call binding the contract method 0xc370b042.
//
// Solidity: function denom() view returns(string)
func (_TNT721VoucherContract *TNT721VoucherContractCaller) Denom(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TNT721VoucherContract.contract.Call(opts, &out, "denom")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Denom is a free data retrieval call binding the contract method 0xc370b042.
//
// Solidity: function denom() view returns(string)
func (_TNT721VoucherContract *TNT721VoucherContractSession) Denom() (string, error) {
	return _TNT721VoucherContract.Contract.Denom(&_TNT721VoucherContract.CallOpts)
}

// Denom is a free data retrieval call binding the contract method 0xc370b042.
//
// Solidity: function denom() view returns(string)
func (_TNT721VoucherContract *TNT721VoucherContractCallerSession) Denom() (string, error) {
	return _TNT721VoucherContract.Contract.Denom(&_TNT721VoucherContract.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_TNT721VoucherContract *TNT721VoucherContractCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TNT721VoucherContract.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_TNT721VoucherContract *TNT721VoucherContractSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _TNT721VoucherContract.Contract.GetApproved(&_TNT721VoucherContract.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_TNT721VoucherContract *TNT721VoucherContractCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _TNT721VoucherContract.Contract.GetApproved(&_TNT721VoucherContract.CallOpts, tokenId)
}

// GetContractOwner is a free data retrieval call binding the contract method 0x442890d5.
//
// Solidity: function getContractOwner() view returns(address)
func (_TNT721VoucherContract *TNT721VoucherContractCaller) GetContractOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TNT721VoucherContract.contract.Call(opts, &out, "getContractOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetContractOwner is a free data retrieval call binding the contract method 0x442890d5.
//
// Solidity: function getContractOwner() view returns(address)
func (_TNT721VoucherContract *TNT721VoucherContractSession) GetContractOwner() (common.Address, error) {
	return _TNT721VoucherContract.Contract.GetContractOwner(&_TNT721VoucherContract.CallOpts)
}

// GetContractOwner is a free data retrieval call binding the contract method 0x442890d5.
//
// Solidity: function getContractOwner() view returns(address)
func (_TNT721VoucherContract *TNT721VoucherContractCallerSession) GetContractOwner() (common.Address, error) {
	return _TNT721VoucherContract.Contract.GetContractOwner(&_TNT721VoucherContract.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_TNT721VoucherContract *TNT721VoucherContractCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _TNT721VoucherContract.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_TNT721VoucherContract *TNT721VoucherContractSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _TNT721VoucherContract.Contract.IsApprovedForAll(&_TNT721VoucherContract.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_TNT721VoucherContract *TNT721VoucherContractCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _TNT721VoucherContract.Contract.IsApprovedForAll(&_TNT721VoucherContract.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TNT721VoucherContract *TNT721VoucherContractCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TNT721VoucherContract.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TNT721VoucherContract *TNT721VoucherContractSession) Name() (string, error) {
	return _TNT721VoucherContract.Contract.Name(&_TNT721VoucherContract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TNT721VoucherContract *TNT721VoucherContractCallerSession) Name() (string, error) {
	return _TNT721VoucherContract.Contract.Name(&_TNT721VoucherContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TNT721VoucherContract *TNT721VoucherContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TNT721VoucherContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TNT721VoucherContract *TNT721VoucherContractSession) Owner() (common.Address, error) {
	return _TNT721VoucherContract.Contract.Owner(&_TNT721VoucherContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TNT721VoucherContract *TNT721VoucherContractCallerSession) Owner() (common.Address, error) {
	return _TNT721VoucherContract.Contract.Owner(&_TNT721VoucherContract.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_TNT721VoucherContract *TNT721VoucherContractCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TNT721VoucherContract.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_TNT721VoucherContract *TNT721VoucherContractSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _TNT721VoucherContract.Contract.OwnerOf(&_TNT721VoucherContract.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_TNT721VoucherContract *TNT721VoucherContractCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _TNT721VoucherContract.Contract.OwnerOf(&_TNT721VoucherContract.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TNT721VoucherContract *TNT721VoucherContractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _TNT721VoucherContract.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TNT721VoucherContract *TNT721VoucherContractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TNT721VoucherContract.Contract.SupportsInterface(&_TNT721VoucherContract.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TNT721VoucherContract *TNT721VoucherContractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TNT721VoucherContract.Contract.SupportsInterface(&_TNT721VoucherContract.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TNT721VoucherContract *TNT721VoucherContractCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TNT721VoucherContract.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TNT721VoucherContract *TNT721VoucherContractSession) Symbol() (string, error) {
	return _TNT721VoucherContract.Contract.Symbol(&_TNT721VoucherContract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TNT721VoucherContract *TNT721VoucherContractCallerSession) Symbol() (string, error) {
	return _TNT721VoucherContract.Contract.Symbol(&_TNT721VoucherContract.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenID) view returns(string)
func (_TNT721VoucherContract *TNT721VoucherContractCaller) TokenURI(opts *bind.CallOpts, tokenID *big.Int) (string, error) {
	var out []interface{}
	err := _TNT721VoucherContract.contract.Call(opts, &out, "tokenURI", tokenID)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenID) view returns(string)
func (_TNT721VoucherContract *TNT721VoucherContractSession) TokenURI(tokenID *big.Int) (string, error) {
	return _TNT721VoucherContract.Contract.TokenURI(&_TNT721VoucherContract.CallOpts, tokenID)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenID) view returns(string)
func (_TNT721VoucherContract *TNT721VoucherContractCallerSession) TokenURI(tokenID *big.Int) (string, error) {
	return _TNT721VoucherContract.Contract.TokenURI(&_TNT721VoucherContract.CallOpts, tokenID)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TNT721VoucherContract *TNT721VoucherContractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TNT721VoucherContract.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TNT721VoucherContract *TNT721VoucherContractSession) TotalSupply() (*big.Int, error) {
	return _TNT721VoucherContract.Contract.TotalSupply(&_TNT721VoucherContract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TNT721VoucherContract *TNT721VoucherContractCallerSession) TotalSupply() (*big.Int, error) {
	return _TNT721VoucherContract.Contract.TotalSupply(&_TNT721VoucherContract.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_TNT721VoucherContract *TNT721VoucherContractTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TNT721VoucherContract.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_TNT721VoucherContract *TNT721VoucherContractSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TNT721VoucherContract.Contract.Approve(&_TNT721VoucherContract.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_TNT721VoucherContract *TNT721VoucherContractTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TNT721VoucherContract.Contract.Approve(&_TNT721VoucherContract.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address voucherOwner, uint256 tokenID) returns()
func (_TNT721VoucherContract *TNT721VoucherContractTransactor) Burn(opts *bind.TransactOpts, voucherOwner common.Address, tokenID *big.Int) (*types.Transaction, error) {
	return _TNT721VoucherContract.contract.Transact(opts, "burn", voucherOwner, tokenID)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address voucherOwner, uint256 tokenID) returns()
func (_TNT721VoucherContract *TNT721VoucherContractSession) Burn(voucherOwner common.Address, tokenID *big.Int) (*types.Transaction, error) {
	return _TNT721VoucherContract.Contract.Burn(&_TNT721VoucherContract.TransactOpts, voucherOwner, tokenID)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address voucherOwner, uint256 tokenID) returns()
func (_TNT721VoucherContract *TNT721VoucherContractTransactorSession) Burn(voucherOwner common.Address, tokenID *big.Int) (*types.Transaction, error) {
	return _TNT721VoucherContract.Contract.Burn(&_TNT721VoucherContract.TransactOpts, voucherOwner, tokenID)
}

// Mint is a paid mutator transaction binding the contract method 0xd3fc9864.
//
// Solidity: function mint(address voucherReceiver, uint256 tokenID, string tokenUri) returns()
func (_TNT721VoucherContract *TNT721VoucherContractTransactor) Mint(opts *bind.TransactOpts, voucherReceiver common.Address, tokenID *big.Int, tokenUri string) (*types.Transaction, error) {
	return _TNT721VoucherContract.contract.Transact(opts, "mint", voucherReceiver, tokenID, tokenUri)
}

// Mint is a paid mutator transaction binding the contract method 0xd3fc9864.
//
// Solidity: function mint(address voucherReceiver, uint256 tokenID, string tokenUri) returns()
func (_TNT721VoucherContract *TNT721VoucherContractSession) Mint(voucherReceiver common.Address, tokenID *big.Int, tokenUri string) (*types.Transaction, error) {
	return _TNT721VoucherContract.Contract.Mint(&_TNT721VoucherContract.TransactOpts, voucherReceiver, tokenID, tokenUri)
}

// Mint is a paid mutator transaction binding the contract method 0xd3fc9864.
//
// Solidity: function mint(address voucherReceiver, uint256 tokenID, string tokenUri) returns()
func (_TNT721VoucherContract *TNT721VoucherContractTransactorSession) Mint(voucherReceiver common.Address, tokenID *big.Int, tokenUri string) (*types.Transaction, error) {
	return _TNT721VoucherContract.Contract.Mint(&_TNT721VoucherContract.TransactOpts, voucherReceiver, tokenID, tokenUri)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_TNT721VoucherContract *TNT721VoucherContractTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TNT721VoucherContract.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_TNT721VoucherContract *TNT721VoucherContractSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TNT721VoucherContract.Contract.SafeTransferFrom(&_TNT721VoucherContract.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_TNT721VoucherContract *TNT721VoucherContractTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TNT721VoucherContract.Contract.SafeTransferFrom(&_TNT721VoucherContract.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_TNT721VoucherContract *TNT721VoucherContractTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _TNT721VoucherContract.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_TNT721VoucherContract *TNT721VoucherContractSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _TNT721VoucherContract.Contract.SafeTransferFrom0(&_TNT721VoucherContract.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_TNT721VoucherContract *TNT721VoucherContractTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _TNT721VoucherContract.Contract.SafeTransferFrom0(&_TNT721VoucherContract.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_TNT721VoucherContract *TNT721VoucherContractTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _TNT721VoucherContract.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_TNT721VoucherContract *TNT721VoucherContractSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _TNT721VoucherContract.Contract.SetApprovalForAll(&_TNT721VoucherContract.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_TNT721VoucherContract *TNT721VoucherContractTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _TNT721VoucherContract.Contract.SetApprovalForAll(&_TNT721VoucherContract.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_TNT721VoucherContract *TNT721VoucherContractTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TNT721VoucherContract.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_TNT721VoucherContract *TNT721VoucherContractSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TNT721VoucherContract.Contract.TransferFrom(&_TNT721VoucherContract.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_TNT721VoucherContract *TNT721VoucherContractTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TNT721VoucherContract.Contract.TransferFrom(&_TNT721VoucherContract.TransactOpts, from, to, tokenId)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0x880cdc31.
//
// Solidity: function updateOwner(address newOwner) returns()
func (_TNT721VoucherContract *TNT721VoucherContractTransactor) UpdateOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TNT721VoucherContract.contract.Transact(opts, "updateOwner", newOwner)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0x880cdc31.
//
// Solidity: function updateOwner(address newOwner) returns()
func (_TNT721VoucherContract *TNT721VoucherContractSession) UpdateOwner(newOwner common.Address) (*types.Transaction, error) {
	return _TNT721VoucherContract.Contract.UpdateOwner(&_TNT721VoucherContract.TransactOpts, newOwner)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0x880cdc31.
//
// Solidity: function updateOwner(address newOwner) returns()
func (_TNT721VoucherContract *TNT721VoucherContractTransactorSession) UpdateOwner(newOwner common.Address) (*types.Transaction, error) {
	return _TNT721VoucherContract.Contract.UpdateOwner(&_TNT721VoucherContract.TransactOpts, newOwner)
}

// TNT721VoucherContractApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the TNT721VoucherContract contract.
type TNT721VoucherContractApprovalIterator struct {
	Event *TNT721VoucherContractApproval // Event containing the contract specifics and raw log

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
func (it *TNT721VoucherContractApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TNT721VoucherContractApproval)
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
		it.Event = new(TNT721VoucherContractApproval)
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
func (it *TNT721VoucherContractApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TNT721VoucherContractApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TNT721VoucherContractApproval represents a Approval event raised by the TNT721VoucherContract contract.
type TNT721VoucherContractApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_TNT721VoucherContract *TNT721VoucherContractFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*TNT721VoucherContractApprovalIterator, error) {

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

	logs, sub, err := _TNT721VoucherContract.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &TNT721VoucherContractApprovalIterator{contract: _TNT721VoucherContract.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_TNT721VoucherContract *TNT721VoucherContractFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TNT721VoucherContractApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _TNT721VoucherContract.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TNT721VoucherContractApproval)
				if err := _TNT721VoucherContract.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_TNT721VoucherContract *TNT721VoucherContractFilterer) ParseApproval(log types.Log) (*TNT721VoucherContractApproval, error) {
	event := new(TNT721VoucherContractApproval)
	if err := _TNT721VoucherContract.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TNT721VoucherContractApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the TNT721VoucherContract contract.
type TNT721VoucherContractApprovalForAllIterator struct {
	Event *TNT721VoucherContractApprovalForAll // Event containing the contract specifics and raw log

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
func (it *TNT721VoucherContractApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TNT721VoucherContractApprovalForAll)
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
		it.Event = new(TNT721VoucherContractApprovalForAll)
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
func (it *TNT721VoucherContractApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TNT721VoucherContractApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TNT721VoucherContractApprovalForAll represents a ApprovalForAll event raised by the TNT721VoucherContract contract.
type TNT721VoucherContractApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_TNT721VoucherContract *TNT721VoucherContractFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*TNT721VoucherContractApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TNT721VoucherContract.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &TNT721VoucherContractApprovalForAllIterator{contract: _TNT721VoucherContract.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_TNT721VoucherContract *TNT721VoucherContractFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *TNT721VoucherContractApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TNT721VoucherContract.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TNT721VoucherContractApprovalForAll)
				if err := _TNT721VoucherContract.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_TNT721VoucherContract *TNT721VoucherContractFilterer) ParseApprovalForAll(log types.Log) (*TNT721VoucherContractApprovalForAll, error) {
	event := new(TNT721VoucherContractApprovalForAll)
	if err := _TNT721VoucherContract.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TNT721VoucherContractTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the TNT721VoucherContract contract.
type TNT721VoucherContractTransferIterator struct {
	Event *TNT721VoucherContractTransfer // Event containing the contract specifics and raw log

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
func (it *TNT721VoucherContractTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TNT721VoucherContractTransfer)
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
		it.Event = new(TNT721VoucherContractTransfer)
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
func (it *TNT721VoucherContractTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TNT721VoucherContractTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TNT721VoucherContractTransfer represents a Transfer event raised by the TNT721VoucherContract contract.
type TNT721VoucherContractTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_TNT721VoucherContract *TNT721VoucherContractFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*TNT721VoucherContractTransferIterator, error) {

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

	logs, sub, err := _TNT721VoucherContract.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &TNT721VoucherContractTransferIterator{contract: _TNT721VoucherContract.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_TNT721VoucherContract *TNT721VoucherContractFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TNT721VoucherContractTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _TNT721VoucherContract.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TNT721VoucherContractTransfer)
				if err := _TNT721VoucherContract.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_TNT721VoucherContract *TNT721VoucherContractFilterer) ParseTransfer(log types.Log) (*TNT721VoucherContractTransfer, error) {
	event := new(TNT721VoucherContractTransfer)
	if err := _TNT721VoucherContract.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TNT721VoucherContractUpdateOwnerIterator is returned from FilterUpdateOwner and is used to iterate over the raw logs and unpacked data for UpdateOwner events raised by the TNT721VoucherContract contract.
type TNT721VoucherContractUpdateOwnerIterator struct {
	Event *TNT721VoucherContractUpdateOwner // Event containing the contract specifics and raw log

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
func (it *TNT721VoucherContractUpdateOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TNT721VoucherContractUpdateOwner)
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
		it.Event = new(TNT721VoucherContractUpdateOwner)
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
func (it *TNT721VoucherContractUpdateOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TNT721VoucherContractUpdateOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TNT721VoucherContractUpdateOwner represents a UpdateOwner event raised by the TNT721VoucherContract contract.
type TNT721VoucherContractUpdateOwner struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdateOwner is a free log retrieval operation binding the contract event 0xe2c7d1c4da37855e682bde14f17826d185497973b73fba7554daa6da467058d9.
//
// Solidity: event UpdateOwner(address oldOwner, address newOwner)
func (_TNT721VoucherContract *TNT721VoucherContractFilterer) FilterUpdateOwner(opts *bind.FilterOpts) (*TNT721VoucherContractUpdateOwnerIterator, error) {

	logs, sub, err := _TNT721VoucherContract.contract.FilterLogs(opts, "UpdateOwner")
	if err != nil {
		return nil, err
	}
	return &TNT721VoucherContractUpdateOwnerIterator{contract: _TNT721VoucherContract.contract, event: "UpdateOwner", logs: logs, sub: sub}, nil
}

// WatchUpdateOwner is a free log subscription operation binding the contract event 0xe2c7d1c4da37855e682bde14f17826d185497973b73fba7554daa6da467058d9.
//
// Solidity: event UpdateOwner(address oldOwner, address newOwner)
func (_TNT721VoucherContract *TNT721VoucherContractFilterer) WatchUpdateOwner(opts *bind.WatchOpts, sink chan<- *TNT721VoucherContractUpdateOwner) (event.Subscription, error) {

	logs, sub, err := _TNT721VoucherContract.contract.WatchLogs(opts, "UpdateOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TNT721VoucherContractUpdateOwner)
				if err := _TNT721VoucherContract.contract.UnpackLog(event, "UpdateOwner", log); err != nil {
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

// ParseUpdateOwner is a log parse operation binding the contract event 0xe2c7d1c4da37855e682bde14f17826d185497973b73fba7554daa6da467058d9.
//
// Solidity: event UpdateOwner(address oldOwner, address newOwner)
func (_TNT721VoucherContract *TNT721VoucherContractFilterer) ParseUpdateOwner(log types.Log) (*TNT721VoucherContractUpdateOwner, error) {
	event := new(TNT721VoucherContractUpdateOwner)
	if err := _TNT721VoucherContract.contract.UnpackLog(event, "UpdateOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
