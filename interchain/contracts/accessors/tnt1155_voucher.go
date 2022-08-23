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

// TNT1155VoucherContractMetaData contains all meta data concerning the TNT1155VoucherContract contract.
var TNT1155VoucherContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"demon_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uriForm_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"UpdateOwner\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"updateOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voucherReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintedAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voucherReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIDs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"mintedAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"mintBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voucherOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"burnedAmount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voucherOwner\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIDs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"burnedAmounts\",\"type\":\"uint256[]\"}],\"name\":\"burnBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"denom\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContractOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200239d3803806200239d8339810160408190526200003491620001f7565b8281620000418162000081565b50600380546001600160a01b0319166001600160a01b03929092169190911790558151620000779060049060208501906200009a565b50505050620002d4565b8051620000969060029060208401906200009a565b5050565b828054620000a89062000281565b90600052602060002090601f016020900481019282620000cc576000855562000117565b82601f10620000e757805160ff191683800117855562000117565b8280016001018555821562000117579182015b8281111562000117578251825591602001919060010190620000fa565b506200012592915062000129565b5090565b5b808211156200012557600081556001016200012a565b600082601f8301126200015257600080fd5b81516001600160401b03808211156200016f576200016f620002be565b604051601f8301601f19908116603f011681019082821181831017156200019a576200019a620002be565b81604052838152602092508683858801011115620001b757600080fd5b600091505b83821015620001db5785820183015181830184015290820190620001bc565b83821115620001ed5760008385830101525b9695505050505050565b6000806000606084860312156200020d57600080fd5b83516001600160a01b03811681146200022557600080fd5b60208501519093506001600160401b03808211156200024357600080fd5b620002518783880162000140565b935060408601519150808211156200026857600080fd5b50620002778682870162000140565b9150509250925092565b600181811c908216806200029657607f821691505b60208210811415620002b857634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052604160045260246000fd5b6120b980620002e46000396000f3fe608060405234801561001057600080fd5b506004361061010a5760003560e01c8063731133e9116100a2578063c370b04211610071578063c370b04214610244578063c87b56dd1461024c578063e985e9c51461025f578063f242432a14610272578063f5298aca1461028557600080fd5b8063731133e9146101f8578063880cdc311461020b5780638da5cb5b1461021e578063a22cb4651461023157600080fd5b80632eb2c2d6116100de5780632eb2c2d61461018d578063442890d5146101a05780634e1273f4146101c55780636b20c454146101e557600080fd5b8062fdd58e1461010f57806301ffc9a7146101355780630e89341c146101585780631f7fdffa14610178575b600080fd5b61012261011d366004611866565b610298565b6040519081526020015b60405180910390f35b6101486101433660046119e9565b61032e565b604051901515815260200161012c565b61016b610166366004611a23565b610380565b60405161012c9190611ba8565b61018b610186366004611791565b610414565b005b61018b61019b366004611601565b610450565b6003546001600160a01b03165b6040516001600160a01b03909116815260200161012c565b6101d86101d3366004611918565b61049c565b60405161012c9190611b67565b61018b6101f3366004611710565b6105c6565b61018b6102063660046118c3565b61070f565b61018b6102193660046115ac565b610745565b6003546101ad906001600160a01b031681565b61018b61023f36600461182a565b6107d8565b61016b6107e7565b61016b61025a366004611a23565b610879565b61014861026d3660046115ce565b610884565b61018b6102803660046116ab565b6108b2565b61018b610293366004611890565b6108f7565b60006001600160a01b0383166103085760405162461bcd60e51b815260206004820152602a60248201527f455243313135353a2061646472657373207a65726f206973206e6f742061207660448201526930b634b21037bbb732b960b11b60648201526084015b60405180910390fd5b506000908152602081815260408083206001600160a01b03949094168352929052205490565b60006001600160e01b03198216636cdb3d1360e11b148061035f57506001600160e01b031982166303a24d0760e21b145b8061037a57506301ffc9a760e01b6001600160e01b03198316145b92915050565b60606002805461038f90611eff565b80601f01602080910402602001604051908101604052809291908181526020018280546103bb90611eff565b80156104085780601f106103dd57610100808354040283529160200191610408565b820191906000526020600020905b8154815290600101906020018083116103eb57829003601f168201915b50505050509050919050565b6003546001600160a01b0316331461043e5760405162461bcd60e51b81526004016102ff90611db1565b61044a84848484610987565b50505050565b6001600160a01b03851633148061046c575061046c8533610884565b6104885760405162461bcd60e51b81526004016102ff90611bbb565b6104958585858585610ad2565b5050505050565b606081518351146105015760405162461bcd60e51b815260206004820152602960248201527f455243313135353a206163636f756e747320616e6420696473206c656e677468604482015268040dad2e6dac2e8c6d60bb1b60648201526084016102ff565b6000835167ffffffffffffffff81111561051d5761051d611fae565b604051908082528060200260200182016040528015610546578160200160208202803683370190505b50905060005b84518110156105be5761059185828151811061056a5761056a611f98565b602002602001015185838151811061058457610584611f98565b6020026020010151610298565b8282815181106105a3576105a3611f98565b60209081029190910101526105b781611f67565b905061054c565b509392505050565b6003546001600160a01b031633146105f05760405162461bcd60e51b81526004016102ff90611db1565b6003546001600160a01b03166106068682610884565b6106225760405162461bcd60e51b81526004016102ff90611cdb565b60005b848110156106985783838281811061063f5761063f611f98565b905060200201356106688888888581811061065c5761065c611f98565b90506020020135610298565b10156106865760405162461bcd60e51b81526004016102ff90611e71565b8061069081611f67565b915050610625565b506107078686868080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525050604080516020808a02828101820190935289825290935089925088918291850190849080828437600092019190915250610c6692505050565b505050505050565b6003546001600160a01b031633146107395760405162461bcd60e51b81526004016102ff90611db1565b61044a84848484610df0565b6003546001600160a01b0316331461076f5760405162461bcd60e51b81526004016102ff90611db1565b600354604080516001600160a01b03928316815291831660208301527fe2c7d1c4da37855e682bde14f17826d185497973b73fba7554daa6da467058d9910160405180910390a1600380546001600160a01b0319166001600160a01b0392909216919091179055565b6107e3338383610eca565b5050565b6060600480546107f690611eff565b80601f016020809104026020016040519081016040528092919081815260200182805461082290611eff565b801561086f5780601f106108445761010080835404028352916020019161086f565b820191906000526020600020905b81548152906001019060200180831161085257829003601f168201915b5050505050905090565b606061037a82610380565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205460ff1690565b6001600160a01b0385163314806108ce57506108ce8533610884565b6108ea5760405162461bcd60e51b81526004016102ff90611bbb565b6104958585858585610fab565b6003546001600160a01b031633146109215760405162461bcd60e51b81526004016102ff90611db1565b8061092c8484610298565b101561094a5760405162461bcd60e51b81526004016102ff90611e71565b6003546001600160a01b03166109608482610884565b61097c5760405162461bcd60e51b81526004016102ff90611cdb565b61044a8484846110d5565b6001600160a01b0384166109ad5760405162461bcd60e51b81526004016102ff90611e30565b81518351146109ce5760405162461bcd60e51b81526004016102ff90611de8565b3360005b8451811015610a6a578381815181106109ed576109ed611f98565b6020026020010151600080878481518110610a0a57610a0a611f98565b602002602001015181526020019081526020016000206000886001600160a01b03166001600160a01b031681526020019081526020016000206000828254610a529190611ee7565b90915550819050610a6281611f67565b9150506109d2565b50846001600160a01b031660006001600160a01b0316826001600160a01b03167f4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb8787604051610abb929190611b7a565b60405180910390a4610495816000878787876111d9565b8151835114610af35760405162461bcd60e51b81526004016102ff90611de8565b6001600160a01b038416610b195760405162461bcd60e51b81526004016102ff90611c96565b3360005b8451811015610c00576000858281518110610b3a57610b3a611f98565b602002602001015190506000858381518110610b5857610b58611f98565b602090810291909101810151600084815280835260408082206001600160a01b038e168352909352919091205490915081811015610ba85760405162461bcd60e51b81526004016102ff90611d67565b6000838152602081815260408083206001600160a01b038e8116855292528083208585039055908b16825281208054849290610be5908490611ee7565b9250508190555050505080610bf990611f67565b9050610b1d565b50846001600160a01b0316866001600160a01b0316826001600160a01b03167f4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb8787604051610c50929190611b7a565b60405180910390a46107078187878787876111d9565b6001600160a01b038316610c8c5760405162461bcd60e51b81526004016102ff90611d24565b8051825114610cad5760405162461bcd60e51b81526004016102ff90611de8565b604080516020810190915260009081905233905b8351811015610d83576000848281518110610cde57610cde611f98565b602002602001015190506000848381518110610cfc57610cfc611f98565b602090810291909101810151600084815280835260408082206001600160a01b038c168352909352919091205490915081811015610d4c5760405162461bcd60e51b81526004016102ff90611c52565b6000928352602083815260408085206001600160a01b038b1686529091529092209103905580610d7b81611f67565b915050610cc1565b5060006001600160a01b0316846001600160a01b0316826001600160a01b03167f4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb8686604051610dd4929190611b7a565b60405180910390a460408051602081019091526000905261044a565b6001600160a01b038416610e165760405162461bcd60e51b81526004016102ff90611e30565b336000610e2285611344565b90506000610e2f85611344565b90506000868152602081815260408083206001600160a01b038b16845290915281208054879290610e61908490611ee7565b909155505060408051878152602081018790526001600160a01b03808a1692600092918716917fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62910160405180910390a4610ec18360008989898961138f565b50505050505050565b816001600160a01b0316836001600160a01b03161415610f3e5760405162461bcd60e51b815260206004820152602960248201527f455243313135353a2073657474696e6720617070726f76616c20737461747573604482015268103337b91039b2b63360b91b60648201526084016102ff565b6001600160a01b03838116600081815260016020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b6001600160a01b038416610fd15760405162461bcd60e51b81526004016102ff90611c96565b336000610fdd85611344565b90506000610fea85611344565b90506000868152602081815260408083206001600160a01b038c1684529091529020548581101561102d5760405162461bcd60e51b81526004016102ff90611d67565b6000878152602081815260408083206001600160a01b038d8116855292528083208985039055908a1682528120805488929061106a908490611ee7565b909155505060408051888152602081018890526001600160a01b03808b16928c821692918816917fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62910160405180910390a46110ca848a8a8a8a8a61138f565b505050505050505050565b6001600160a01b0383166110fb5760405162461bcd60e51b81526004016102ff90611d24565b33600061110784611344565b9050600061111484611344565b60408051602080820183526000918290528882528181528282206001600160a01b038b16835290522054909150848110156111615760405162461bcd60e51b81526004016102ff90611c52565b6000868152602081815260408083206001600160a01b038b81168086529184528285208a8703905582518b81529384018a90529092908816917fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62910160405180910390a4604080516020810190915260009052610ec1565b6001600160a01b0384163b156107075760405163bc197c8160e01b81526001600160a01b0385169063bc197c819061121d9089908990889088908890600401611ac4565b602060405180830381600087803b15801561123757600080fd5b505af1925050508015611267575060408051601f3d908101601f1916820190925261126491810190611a06565b60015b61131457611273611fc4565b806308c379a014156112ad5750611288611fe0565b8061129357506112af565b8060405162461bcd60e51b81526004016102ff9190611ba8565b505b60405162461bcd60e51b815260206004820152603460248201527f455243313135353a207472616e7366657220746f206e6f6e20455243313135356044820152732932b1b2b4bb32b91034b6b83632b6b2b73a32b960611b60648201526084016102ff565b6001600160e01b0319811663bc197c8160e01b14610ec15760405162461bcd60e51b81526004016102ff90611c0a565b6040805160018082528183019092526060916000919060208083019080368337019050509050828160008151811061137e5761137e611f98565b602090810291909101015292915050565b6001600160a01b0384163b156107075760405163f23a6e6160e01b81526001600160a01b0385169063f23a6e61906113d39089908990889088908890600401611b22565b602060405180830381600087803b1580156113ed57600080fd5b505af192505050801561141d575060408051601f3d908101601f1916820190925261141a91810190611a06565b60015b61142957611273611fc4565b6001600160e01b0319811663f23a6e6160e01b14610ec15760405162461bcd60e51b81526004016102ff90611c0a565b80356001600160a01b038116811461147057600080fd5b919050565b60008083601f84011261148757600080fd5b50813567ffffffffffffffff81111561149f57600080fd5b6020830191508360208260051b85010111156114ba57600080fd5b9250929050565b600082601f8301126114d257600080fd5b813560206114df82611ec3565b6040516114ec8282611f3a565b8381528281019150858301600585901b8701840188101561150c57600080fd5b60005b8581101561152b5781358452928401929084019060010161150f565b5090979650505050505050565b600082601f83011261154957600080fd5b813567ffffffffffffffff81111561156357611563611fae565b60405161157a601f8301601f191660200182611f3a565b81815284602083860101111561158f57600080fd5b816020850160208301376000918101602001919091529392505050565b6000602082840312156115be57600080fd5b6115c782611459565b9392505050565b600080604083850312156115e157600080fd5b6115ea83611459565b91506115f860208401611459565b90509250929050565b600080600080600060a0868803121561161957600080fd5b61162286611459565b945061163060208701611459565b9350604086013567ffffffffffffffff8082111561164d57600080fd5b61165989838a016114c1565b9450606088013591508082111561166f57600080fd5b61167b89838a016114c1565b9350608088013591508082111561169157600080fd5b5061169e88828901611538565b9150509295509295909350565b600080600080600060a086880312156116c357600080fd5b6116cc86611459565b94506116da60208701611459565b93506040860135925060608601359150608086013567ffffffffffffffff81111561170457600080fd5b61169e88828901611538565b60008060008060006060868803121561172857600080fd5b61173186611459565b9450602086013567ffffffffffffffff8082111561174e57600080fd5b61175a89838a01611475565b9096509450604088013591508082111561177357600080fd5b5061178088828901611475565b969995985093965092949392505050565b600080600080608085870312156117a757600080fd5b6117b085611459565b9350602085013567ffffffffffffffff808211156117cd57600080fd5b6117d9888389016114c1565b945060408701359150808211156117ef57600080fd5b6117fb888389016114c1565b9350606087013591508082111561181157600080fd5b5061181e87828801611538565b91505092959194509250565b6000806040838503121561183d57600080fd5b61184683611459565b91506020830135801515811461185b57600080fd5b809150509250929050565b6000806040838503121561187957600080fd5b61188283611459565b946020939093013593505050565b6000806000606084860312156118a557600080fd5b6118ae84611459565b95602085013595506040909401359392505050565b600080600080608085870312156118d957600080fd5b6118e285611459565b93506020850135925060408501359150606085013567ffffffffffffffff81111561190c57600080fd5b61181e87828801611538565b6000806040838503121561192b57600080fd5b823567ffffffffffffffff8082111561194357600080fd5b818501915085601f83011261195757600080fd5b8135602061196482611ec3565b6040516119718282611f3a565b8381528281019150858301600585901b870184018b101561199157600080fd5b600096505b848710156119bb576119a781611459565b835260019690960195918301918301611996565b50965050860135925050808211156119d257600080fd5b506119df858286016114c1565b9150509250929050565b6000602082840312156119fb57600080fd5b81356115c78161206a565b600060208284031215611a1857600080fd5b81516115c78161206a565b600060208284031215611a3557600080fd5b5035919050565b600081518084526020808501945080840160005b83811015611a6c57815187529582019590820190600101611a50565b509495945050505050565b6000815180845260005b81811015611a9d57602081850181015186830182015201611a81565b81811115611aaf576000602083870101525b50601f01601f19169290920160200192915050565b6001600160a01b0386811682528516602082015260a060408201819052600090611af090830186611a3c565b8281036060840152611b028186611a3c565b90508281036080840152611b168185611a77565b98975050505050505050565b6001600160a01b03868116825285166020820152604081018490526060810183905260a060808201819052600090611b5c90830184611a77565b979650505050505050565b6020815260006115c76020830184611a3c565b604081526000611b8d6040830185611a3c565b8281036020840152611b9f8185611a3c565b95945050505050565b6020815260006115c76020830184611a77565b6020808252602f908201527f455243313135353a2063616c6c6572206973206e6f7420746f6b656e206f776e60408201526e195c881b9bdc88185c1c1c9bdd9959608a1b606082015260800190565b60208082526028908201527f455243313135353a204552433131353552656365697665722072656a656374656040820152676420746f6b656e7360c01b606082015260800190565b60208082526024908201527f455243313135353a206275726e20616d6f756e7420657863656564732062616c604082015263616e636560e01b606082015260800190565b60208082526025908201527f455243313135353a207472616e7366657220746f20746865207a65726f206164604082015264647265737360d81b606082015260800190565b60208082526029908201527f566f7563686572206f776e657220646964206e6f7420617070726f766520796f6040820152683a903a3790313ab93760b91b606082015260800190565b60208082526023908201527f455243313135353a206275726e2066726f6d20746865207a65726f206164647260408201526265737360e81b606082015260800190565b6020808252602a908201527f455243313135353a20696e73756666696369656e742062616c616e636520666f60408201526939103a3930b739b332b960b11b606082015260800190565b6020808252601c908201527f6f6e6c79206f776e65722063616e206d616b65207468652063616c6c00000000604082015260600190565b60208082526028908201527f455243313135353a2069647320616e6420616d6f756e7473206c656e677468206040820152670dad2e6dac2e8c6d60c31b606082015260800190565b60208082526021908201527f455243313135353a206d696e7420746f20746865207a65726f206164647265736040820152607360f81b606082015260800190565b60208082526032908201527f566f7563686572206f776e657220646f6573206e6f74206861766520656e6f7560408201527133b4103130b630b731b2903a3790313ab93760711b606082015260800190565b600067ffffffffffffffff821115611edd57611edd611fae565b5060051b60200190565b60008219821115611efa57611efa611f82565b500190565b600181811c90821680611f1357607f821691505b60208210811415611f3457634e487b7160e01b600052602260045260246000fd5b50919050565b601f8201601f1916810167ffffffffffffffff81118282101715611f6057611f60611fae565b6040525050565b6000600019821415611f7b57611f7b611f82565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b600060033d1115611fdd5760046000803e5060005160e01c5b90565b600060443d1015611fee5790565b6040516003193d81016004833e81513d67ffffffffffffffff816024840111818411171561201e57505050505090565b82850191508151818111156120365750505050505090565b843d87010160208285010111156120505750505050505090565b61205f60208286010187611f3a565b509095945050505050565b6001600160e01b03198116811461208057600080fd5b5056fea264697066735822122053d506f069333f9d17a4cccb391dddf422e668da444374e8a06c3414927c16b164736f6c63430008070033",
}

// TNT1155VoucherContractABI is the input ABI used to generate the binding from.
// Deprecated: Use TNT1155VoucherContractMetaData.ABI instead.
var TNT1155VoucherContractABI = TNT1155VoucherContractMetaData.ABI

// TNT1155VoucherContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TNT1155VoucherContractMetaData.Bin instead.
var TNT1155VoucherContractBin = TNT1155VoucherContractMetaData.Bin

// DeployTNT1155VoucherContract deploys a new Ethereum contract, binding an instance of TNT1155VoucherContract to it.
func DeployTNT1155VoucherContract(auth *bind.TransactOpts, backend bind.ContractBackend, owner_ common.Address, demon_ string, uriForm_ string) (common.Address, *types.Transaction, *TNT1155VoucherContract, error) {
	parsed, err := TNT1155VoucherContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TNT1155VoucherContractBin), backend, owner_, demon_, uriForm_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TNT1155VoucherContract{TNT1155VoucherContractCaller: TNT1155VoucherContractCaller{contract: contract}, TNT1155VoucherContractTransactor: TNT1155VoucherContractTransactor{contract: contract}, TNT1155VoucherContractFilterer: TNT1155VoucherContractFilterer{contract: contract}}, nil
}

// TNT1155VoucherContract is an auto generated Go binding around an Ethereum contract.
type TNT1155VoucherContract struct {
	TNT1155VoucherContractCaller     // Read-only binding to the contract
	TNT1155VoucherContractTransactor // Write-only binding to the contract
	TNT1155VoucherContractFilterer   // Log filterer for contract events
}

// TNT1155VoucherContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type TNT1155VoucherContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TNT1155VoucherContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TNT1155VoucherContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TNT1155VoucherContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TNT1155VoucherContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TNT1155VoucherContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TNT1155VoucherContractSession struct {
	Contract     *TNT1155VoucherContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TNT1155VoucherContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TNT1155VoucherContractCallerSession struct {
	Contract *TNT1155VoucherContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// TNT1155VoucherContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TNT1155VoucherContractTransactorSession struct {
	Contract     *TNT1155VoucherContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// TNT1155VoucherContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type TNT1155VoucherContractRaw struct {
	Contract *TNT1155VoucherContract // Generic contract binding to access the raw methods on
}

// TNT1155VoucherContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TNT1155VoucherContractCallerRaw struct {
	Contract *TNT1155VoucherContractCaller // Generic read-only contract binding to access the raw methods on
}

// TNT1155VoucherContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TNT1155VoucherContractTransactorRaw struct {
	Contract *TNT1155VoucherContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTNT1155VoucherContract creates a new instance of TNT1155VoucherContract, bound to a specific deployed contract.
func NewTNT1155VoucherContract(address common.Address, backend bind.ContractBackend) (*TNT1155VoucherContract, error) {
	contract, err := bindTNT1155VoucherContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TNT1155VoucherContract{TNT1155VoucherContractCaller: TNT1155VoucherContractCaller{contract: contract}, TNT1155VoucherContractTransactor: TNT1155VoucherContractTransactor{contract: contract}, TNT1155VoucherContractFilterer: TNT1155VoucherContractFilterer{contract: contract}}, nil
}

// NewTNT1155VoucherContractCaller creates a new read-only instance of TNT1155VoucherContract, bound to a specific deployed contract.
func NewTNT1155VoucherContractCaller(address common.Address, caller bind.ContractCaller) (*TNT1155VoucherContractCaller, error) {
	contract, err := bindTNT1155VoucherContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TNT1155VoucherContractCaller{contract: contract}, nil
}

// NewTNT1155VoucherContractTransactor creates a new write-only instance of TNT1155VoucherContract, bound to a specific deployed contract.
func NewTNT1155VoucherContractTransactor(address common.Address, transactor bind.ContractTransactor) (*TNT1155VoucherContractTransactor, error) {
	contract, err := bindTNT1155VoucherContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TNT1155VoucherContractTransactor{contract: contract}, nil
}

// NewTNT1155VoucherContractFilterer creates a new log filterer instance of TNT1155VoucherContract, bound to a specific deployed contract.
func NewTNT1155VoucherContractFilterer(address common.Address, filterer bind.ContractFilterer) (*TNT1155VoucherContractFilterer, error) {
	contract, err := bindTNT1155VoucherContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TNT1155VoucherContractFilterer{contract: contract}, nil
}

// bindTNT1155VoucherContract binds a generic wrapper to an already deployed contract.
func bindTNT1155VoucherContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TNT1155VoucherContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TNT1155VoucherContract *TNT1155VoucherContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TNT1155VoucherContract.Contract.TNT1155VoucherContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TNT1155VoucherContract *TNT1155VoucherContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TNT1155VoucherContract.Contract.TNT1155VoucherContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TNT1155VoucherContract *TNT1155VoucherContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TNT1155VoucherContract.Contract.TNT1155VoucherContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TNT1155VoucherContract *TNT1155VoucherContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TNT1155VoucherContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TNT1155VoucherContract *TNT1155VoucherContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TNT1155VoucherContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TNT1155VoucherContract *TNT1155VoucherContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TNT1155VoucherContract.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_TNT1155VoucherContract *TNT1155VoucherContractCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TNT1155VoucherContract.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_TNT1155VoucherContract *TNT1155VoucherContractSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _TNT1155VoucherContract.Contract.BalanceOf(&_TNT1155VoucherContract.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_TNT1155VoucherContract *TNT1155VoucherContractCallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _TNT1155VoucherContract.Contract.BalanceOf(&_TNT1155VoucherContract.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_TNT1155VoucherContract *TNT1155VoucherContractCaller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _TNT1155VoucherContract.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_TNT1155VoucherContract *TNT1155VoucherContractSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _TNT1155VoucherContract.Contract.BalanceOfBatch(&_TNT1155VoucherContract.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_TNT1155VoucherContract *TNT1155VoucherContractCallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _TNT1155VoucherContract.Contract.BalanceOfBatch(&_TNT1155VoucherContract.CallOpts, accounts, ids)
}

// Denom is a free data retrieval call binding the contract method 0xc370b042.
//
// Solidity: function denom() view returns(string)
func (_TNT1155VoucherContract *TNT1155VoucherContractCaller) Denom(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TNT1155VoucherContract.contract.Call(opts, &out, "denom")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Denom is a free data retrieval call binding the contract method 0xc370b042.
//
// Solidity: function denom() view returns(string)
func (_TNT1155VoucherContract *TNT1155VoucherContractSession) Denom() (string, error) {
	return _TNT1155VoucherContract.Contract.Denom(&_TNT1155VoucherContract.CallOpts)
}

// Denom is a free data retrieval call binding the contract method 0xc370b042.
//
// Solidity: function denom() view returns(string)
func (_TNT1155VoucherContract *TNT1155VoucherContractCallerSession) Denom() (string, error) {
	return _TNT1155VoucherContract.Contract.Denom(&_TNT1155VoucherContract.CallOpts)
}

// GetContractOwner is a free data retrieval call binding the contract method 0x442890d5.
//
// Solidity: function getContractOwner() view returns(address)
func (_TNT1155VoucherContract *TNT1155VoucherContractCaller) GetContractOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TNT1155VoucherContract.contract.Call(opts, &out, "getContractOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetContractOwner is a free data retrieval call binding the contract method 0x442890d5.
//
// Solidity: function getContractOwner() view returns(address)
func (_TNT1155VoucherContract *TNT1155VoucherContractSession) GetContractOwner() (common.Address, error) {
	return _TNT1155VoucherContract.Contract.GetContractOwner(&_TNT1155VoucherContract.CallOpts)
}

// GetContractOwner is a free data retrieval call binding the contract method 0x442890d5.
//
// Solidity: function getContractOwner() view returns(address)
func (_TNT1155VoucherContract *TNT1155VoucherContractCallerSession) GetContractOwner() (common.Address, error) {
	return _TNT1155VoucherContract.Contract.GetContractOwner(&_TNT1155VoucherContract.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_TNT1155VoucherContract *TNT1155VoucherContractCaller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _TNT1155VoucherContract.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_TNT1155VoucherContract *TNT1155VoucherContractSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _TNT1155VoucherContract.Contract.IsApprovedForAll(&_TNT1155VoucherContract.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_TNT1155VoucherContract *TNT1155VoucherContractCallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _TNT1155VoucherContract.Contract.IsApprovedForAll(&_TNT1155VoucherContract.CallOpts, account, operator)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TNT1155VoucherContract *TNT1155VoucherContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TNT1155VoucherContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TNT1155VoucherContract *TNT1155VoucherContractSession) Owner() (common.Address, error) {
	return _TNT1155VoucherContract.Contract.Owner(&_TNT1155VoucherContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TNT1155VoucherContract *TNT1155VoucherContractCallerSession) Owner() (common.Address, error) {
	return _TNT1155VoucherContract.Contract.Owner(&_TNT1155VoucherContract.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TNT1155VoucherContract *TNT1155VoucherContractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _TNT1155VoucherContract.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TNT1155VoucherContract *TNT1155VoucherContractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TNT1155VoucherContract.Contract.SupportsInterface(&_TNT1155VoucherContract.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TNT1155VoucherContract *TNT1155VoucherContractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TNT1155VoucherContract.Contract.SupportsInterface(&_TNT1155VoucherContract.CallOpts, interfaceId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenID) view returns(string)
func (_TNT1155VoucherContract *TNT1155VoucherContractCaller) TokenURI(opts *bind.CallOpts, tokenID *big.Int) (string, error) {
	var out []interface{}
	err := _TNT1155VoucherContract.contract.Call(opts, &out, "tokenURI", tokenID)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenID) view returns(string)
func (_TNT1155VoucherContract *TNT1155VoucherContractSession) TokenURI(tokenID *big.Int) (string, error) {
	return _TNT1155VoucherContract.Contract.TokenURI(&_TNT1155VoucherContract.CallOpts, tokenID)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenID) view returns(string)
func (_TNT1155VoucherContract *TNT1155VoucherContractCallerSession) TokenURI(tokenID *big.Int) (string, error) {
	return _TNT1155VoucherContract.Contract.TokenURI(&_TNT1155VoucherContract.CallOpts, tokenID)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_TNT1155VoucherContract *TNT1155VoucherContractCaller) Uri(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _TNT1155VoucherContract.contract.Call(opts, &out, "uri", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_TNT1155VoucherContract *TNT1155VoucherContractSession) Uri(arg0 *big.Int) (string, error) {
	return _TNT1155VoucherContract.Contract.Uri(&_TNT1155VoucherContract.CallOpts, arg0)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_TNT1155VoucherContract *TNT1155VoucherContractCallerSession) Uri(arg0 *big.Int) (string, error) {
	return _TNT1155VoucherContract.Contract.Uri(&_TNT1155VoucherContract.CallOpts, arg0)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address voucherOwner, uint256 tokenID, uint256 burnedAmount) returns()
func (_TNT1155VoucherContract *TNT1155VoucherContractTransactor) Burn(opts *bind.TransactOpts, voucherOwner common.Address, tokenID *big.Int, burnedAmount *big.Int) (*types.Transaction, error) {
	return _TNT1155VoucherContract.contract.Transact(opts, "burn", voucherOwner, tokenID, burnedAmount)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address voucherOwner, uint256 tokenID, uint256 burnedAmount) returns()
func (_TNT1155VoucherContract *TNT1155VoucherContractSession) Burn(voucherOwner common.Address, tokenID *big.Int, burnedAmount *big.Int) (*types.Transaction, error) {
	return _TNT1155VoucherContract.Contract.Burn(&_TNT1155VoucherContract.TransactOpts, voucherOwner, tokenID, burnedAmount)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address voucherOwner, uint256 tokenID, uint256 burnedAmount) returns()
func (_TNT1155VoucherContract *TNT1155VoucherContractTransactorSession) Burn(voucherOwner common.Address, tokenID *big.Int, burnedAmount *big.Int) (*types.Transaction, error) {
	return _TNT1155VoucherContract.Contract.Burn(&_TNT1155VoucherContract.TransactOpts, voucherOwner, tokenID, burnedAmount)
}

// BurnBatch is a paid mutator transaction binding the contract method 0x6b20c454.
//
// Solidity: function burnBatch(address voucherOwner, uint256[] tokenIDs, uint256[] burnedAmounts) returns()
func (_TNT1155VoucherContract *TNT1155VoucherContractTransactor) BurnBatch(opts *bind.TransactOpts, voucherOwner common.Address, tokenIDs []*big.Int, burnedAmounts []*big.Int) (*types.Transaction, error) {
	return _TNT1155VoucherContract.contract.Transact(opts, "burnBatch", voucherOwner, tokenIDs, burnedAmounts)
}

// BurnBatch is a paid mutator transaction binding the contract method 0x6b20c454.
//
// Solidity: function burnBatch(address voucherOwner, uint256[] tokenIDs, uint256[] burnedAmounts) returns()
func (_TNT1155VoucherContract *TNT1155VoucherContractSession) BurnBatch(voucherOwner common.Address, tokenIDs []*big.Int, burnedAmounts []*big.Int) (*types.Transaction, error) {
	return _TNT1155VoucherContract.Contract.BurnBatch(&_TNT1155VoucherContract.TransactOpts, voucherOwner, tokenIDs, burnedAmounts)
}

// BurnBatch is a paid mutator transaction binding the contract method 0x6b20c454.
//
// Solidity: function burnBatch(address voucherOwner, uint256[] tokenIDs, uint256[] burnedAmounts) returns()
func (_TNT1155VoucherContract *TNT1155VoucherContractTransactorSession) BurnBatch(voucherOwner common.Address, tokenIDs []*big.Int, burnedAmounts []*big.Int) (*types.Transaction, error) {
	return _TNT1155VoucherContract.Contract.BurnBatch(&_TNT1155VoucherContract.TransactOpts, voucherOwner, tokenIDs, burnedAmounts)
}

// Mint is a paid mutator transaction binding the contract method 0x731133e9.
//
// Solidity: function mint(address voucherReceiver, uint256 tokenID, uint256 mintedAmount, bytes data) returns()
func (_TNT1155VoucherContract *TNT1155VoucherContractTransactor) Mint(opts *bind.TransactOpts, voucherReceiver common.Address, tokenID *big.Int, mintedAmount *big.Int, data []byte) (*types.Transaction, error) {
	return _TNT1155VoucherContract.contract.Transact(opts, "mint", voucherReceiver, tokenID, mintedAmount, data)
}

// Mint is a paid mutator transaction binding the contract method 0x731133e9.
//
// Solidity: function mint(address voucherReceiver, uint256 tokenID, uint256 mintedAmount, bytes data) returns()
func (_TNT1155VoucherContract *TNT1155VoucherContractSession) Mint(voucherReceiver common.Address, tokenID *big.Int, mintedAmount *big.Int, data []byte) (*types.Transaction, error) {
	return _TNT1155VoucherContract.Contract.Mint(&_TNT1155VoucherContract.TransactOpts, voucherReceiver, tokenID, mintedAmount, data)
}

// Mint is a paid mutator transaction binding the contract method 0x731133e9.
//
// Solidity: function mint(address voucherReceiver, uint256 tokenID, uint256 mintedAmount, bytes data) returns()
func (_TNT1155VoucherContract *TNT1155VoucherContractTransactorSession) Mint(voucherReceiver common.Address, tokenID *big.Int, mintedAmount *big.Int, data []byte) (*types.Transaction, error) {
	return _TNT1155VoucherContract.Contract.Mint(&_TNT1155VoucherContract.TransactOpts, voucherReceiver, tokenID, mintedAmount, data)
}

// MintBatch is a paid mutator transaction binding the contract method 0x1f7fdffa.
//
// Solidity: function mintBatch(address voucherReceiver, uint256[] tokenIDs, uint256[] mintedAmounts, bytes data) returns()
func (_TNT1155VoucherContract *TNT1155VoucherContractTransactor) MintBatch(opts *bind.TransactOpts, voucherReceiver common.Address, tokenIDs []*big.Int, mintedAmounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _TNT1155VoucherContract.contract.Transact(opts, "mintBatch", voucherReceiver, tokenIDs, mintedAmounts, data)
}

// MintBatch is a paid mutator transaction binding the contract method 0x1f7fdffa.
//
// Solidity: function mintBatch(address voucherReceiver, uint256[] tokenIDs, uint256[] mintedAmounts, bytes data) returns()
func (_TNT1155VoucherContract *TNT1155VoucherContractSession) MintBatch(voucherReceiver common.Address, tokenIDs []*big.Int, mintedAmounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _TNT1155VoucherContract.Contract.MintBatch(&_TNT1155VoucherContract.TransactOpts, voucherReceiver, tokenIDs, mintedAmounts, data)
}

// MintBatch is a paid mutator transaction binding the contract method 0x1f7fdffa.
//
// Solidity: function mintBatch(address voucherReceiver, uint256[] tokenIDs, uint256[] mintedAmounts, bytes data) returns()
func (_TNT1155VoucherContract *TNT1155VoucherContractTransactorSession) MintBatch(voucherReceiver common.Address, tokenIDs []*big.Int, mintedAmounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _TNT1155VoucherContract.Contract.MintBatch(&_TNT1155VoucherContract.TransactOpts, voucherReceiver, tokenIDs, mintedAmounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_TNT1155VoucherContract *TNT1155VoucherContractTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _TNT1155VoucherContract.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_TNT1155VoucherContract *TNT1155VoucherContractSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _TNT1155VoucherContract.Contract.SafeBatchTransferFrom(&_TNT1155VoucherContract.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_TNT1155VoucherContract *TNT1155VoucherContractTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _TNT1155VoucherContract.Contract.SafeBatchTransferFrom(&_TNT1155VoucherContract.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_TNT1155VoucherContract *TNT1155VoucherContractTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _TNT1155VoucherContract.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_TNT1155VoucherContract *TNT1155VoucherContractSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _TNT1155VoucherContract.Contract.SafeTransferFrom(&_TNT1155VoucherContract.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_TNT1155VoucherContract *TNT1155VoucherContractTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _TNT1155VoucherContract.Contract.SafeTransferFrom(&_TNT1155VoucherContract.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_TNT1155VoucherContract *TNT1155VoucherContractTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _TNT1155VoucherContract.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_TNT1155VoucherContract *TNT1155VoucherContractSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _TNT1155VoucherContract.Contract.SetApprovalForAll(&_TNT1155VoucherContract.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_TNT1155VoucherContract *TNT1155VoucherContractTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _TNT1155VoucherContract.Contract.SetApprovalForAll(&_TNT1155VoucherContract.TransactOpts, operator, approved)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0x880cdc31.
//
// Solidity: function updateOwner(address newOwner) returns()
func (_TNT1155VoucherContract *TNT1155VoucherContractTransactor) UpdateOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TNT1155VoucherContract.contract.Transact(opts, "updateOwner", newOwner)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0x880cdc31.
//
// Solidity: function updateOwner(address newOwner) returns()
func (_TNT1155VoucherContract *TNT1155VoucherContractSession) UpdateOwner(newOwner common.Address) (*types.Transaction, error) {
	return _TNT1155VoucherContract.Contract.UpdateOwner(&_TNT1155VoucherContract.TransactOpts, newOwner)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0x880cdc31.
//
// Solidity: function updateOwner(address newOwner) returns()
func (_TNT1155VoucherContract *TNT1155VoucherContractTransactorSession) UpdateOwner(newOwner common.Address) (*types.Transaction, error) {
	return _TNT1155VoucherContract.Contract.UpdateOwner(&_TNT1155VoucherContract.TransactOpts, newOwner)
}

// TNT1155VoucherContractApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the TNT1155VoucherContract contract.
type TNT1155VoucherContractApprovalForAllIterator struct {
	Event *TNT1155VoucherContractApprovalForAll // Event containing the contract specifics and raw log

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
func (it *TNT1155VoucherContractApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TNT1155VoucherContractApprovalForAll)
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
		it.Event = new(TNT1155VoucherContractApprovalForAll)
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
func (it *TNT1155VoucherContractApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TNT1155VoucherContractApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TNT1155VoucherContractApprovalForAll represents a ApprovalForAll event raised by the TNT1155VoucherContract contract.
type TNT1155VoucherContractApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_TNT1155VoucherContract *TNT1155VoucherContractFilterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*TNT1155VoucherContractApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TNT1155VoucherContract.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &TNT1155VoucherContractApprovalForAllIterator{contract: _TNT1155VoucherContract.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_TNT1155VoucherContract *TNT1155VoucherContractFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *TNT1155VoucherContractApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TNT1155VoucherContract.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TNT1155VoucherContractApprovalForAll)
				if err := _TNT1155VoucherContract.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_TNT1155VoucherContract *TNT1155VoucherContractFilterer) ParseApprovalForAll(log types.Log) (*TNT1155VoucherContractApprovalForAll, error) {
	event := new(TNT1155VoucherContractApprovalForAll)
	if err := _TNT1155VoucherContract.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TNT1155VoucherContractTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the TNT1155VoucherContract contract.
type TNT1155VoucherContractTransferBatchIterator struct {
	Event *TNT1155VoucherContractTransferBatch // Event containing the contract specifics and raw log

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
func (it *TNT1155VoucherContractTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TNT1155VoucherContractTransferBatch)
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
		it.Event = new(TNT1155VoucherContractTransferBatch)
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
func (it *TNT1155VoucherContractTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TNT1155VoucherContractTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TNT1155VoucherContractTransferBatch represents a TransferBatch event raised by the TNT1155VoucherContract contract.
type TNT1155VoucherContractTransferBatch struct {
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
func (_TNT1155VoucherContract *TNT1155VoucherContractFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*TNT1155VoucherContractTransferBatchIterator, error) {

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

	logs, sub, err := _TNT1155VoucherContract.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TNT1155VoucherContractTransferBatchIterator{contract: _TNT1155VoucherContract.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_TNT1155VoucherContract *TNT1155VoucherContractFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *TNT1155VoucherContractTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _TNT1155VoucherContract.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TNT1155VoucherContractTransferBatch)
				if err := _TNT1155VoucherContract.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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
func (_TNT1155VoucherContract *TNT1155VoucherContractFilterer) ParseTransferBatch(log types.Log) (*TNT1155VoucherContractTransferBatch, error) {
	event := new(TNT1155VoucherContractTransferBatch)
	if err := _TNT1155VoucherContract.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TNT1155VoucherContractTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the TNT1155VoucherContract contract.
type TNT1155VoucherContractTransferSingleIterator struct {
	Event *TNT1155VoucherContractTransferSingle // Event containing the contract specifics and raw log

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
func (it *TNT1155VoucherContractTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TNT1155VoucherContractTransferSingle)
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
		it.Event = new(TNT1155VoucherContractTransferSingle)
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
func (it *TNT1155VoucherContractTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TNT1155VoucherContractTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TNT1155VoucherContractTransferSingle represents a TransferSingle event raised by the TNT1155VoucherContract contract.
type TNT1155VoucherContractTransferSingle struct {
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
func (_TNT1155VoucherContract *TNT1155VoucherContractFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*TNT1155VoucherContractTransferSingleIterator, error) {

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

	logs, sub, err := _TNT1155VoucherContract.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TNT1155VoucherContractTransferSingleIterator{contract: _TNT1155VoucherContract.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_TNT1155VoucherContract *TNT1155VoucherContractFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *TNT1155VoucherContractTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _TNT1155VoucherContract.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TNT1155VoucherContractTransferSingle)
				if err := _TNT1155VoucherContract.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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
func (_TNT1155VoucherContract *TNT1155VoucherContractFilterer) ParseTransferSingle(log types.Log) (*TNT1155VoucherContractTransferSingle, error) {
	event := new(TNT1155VoucherContractTransferSingle)
	if err := _TNT1155VoucherContract.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TNT1155VoucherContractURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the TNT1155VoucherContract contract.
type TNT1155VoucherContractURIIterator struct {
	Event *TNT1155VoucherContractURI // Event containing the contract specifics and raw log

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
func (it *TNT1155VoucherContractURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TNT1155VoucherContractURI)
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
		it.Event = new(TNT1155VoucherContractURI)
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
func (it *TNT1155VoucherContractURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TNT1155VoucherContractURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TNT1155VoucherContractURI represents a URI event raised by the TNT1155VoucherContract contract.
type TNT1155VoucherContractURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_TNT1155VoucherContract *TNT1155VoucherContractFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*TNT1155VoucherContractURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _TNT1155VoucherContract.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &TNT1155VoucherContractURIIterator{contract: _TNT1155VoucherContract.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_TNT1155VoucherContract *TNT1155VoucherContractFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *TNT1155VoucherContractURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _TNT1155VoucherContract.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TNT1155VoucherContractURI)
				if err := _TNT1155VoucherContract.contract.UnpackLog(event, "URI", log); err != nil {
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
func (_TNT1155VoucherContract *TNT1155VoucherContractFilterer) ParseURI(log types.Log) (*TNT1155VoucherContractURI, error) {
	event := new(TNT1155VoucherContractURI)
	if err := _TNT1155VoucherContract.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TNT1155VoucherContractUpdateOwnerIterator is returned from FilterUpdateOwner and is used to iterate over the raw logs and unpacked data for UpdateOwner events raised by the TNT1155VoucherContract contract.
type TNT1155VoucherContractUpdateOwnerIterator struct {
	Event *TNT1155VoucherContractUpdateOwner // Event containing the contract specifics and raw log

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
func (it *TNT1155VoucherContractUpdateOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TNT1155VoucherContractUpdateOwner)
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
		it.Event = new(TNT1155VoucherContractUpdateOwner)
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
func (it *TNT1155VoucherContractUpdateOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TNT1155VoucherContractUpdateOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TNT1155VoucherContractUpdateOwner represents a UpdateOwner event raised by the TNT1155VoucherContract contract.
type TNT1155VoucherContractUpdateOwner struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdateOwner is a free log retrieval operation binding the contract event 0xe2c7d1c4da37855e682bde14f17826d185497973b73fba7554daa6da467058d9.
//
// Solidity: event UpdateOwner(address oldOwner, address newOwner)
func (_TNT1155VoucherContract *TNT1155VoucherContractFilterer) FilterUpdateOwner(opts *bind.FilterOpts) (*TNT1155VoucherContractUpdateOwnerIterator, error) {

	logs, sub, err := _TNT1155VoucherContract.contract.FilterLogs(opts, "UpdateOwner")
	if err != nil {
		return nil, err
	}
	return &TNT1155VoucherContractUpdateOwnerIterator{contract: _TNT1155VoucherContract.contract, event: "UpdateOwner", logs: logs, sub: sub}, nil
}

// WatchUpdateOwner is a free log subscription operation binding the contract event 0xe2c7d1c4da37855e682bde14f17826d185497973b73fba7554daa6da467058d9.
//
// Solidity: event UpdateOwner(address oldOwner, address newOwner)
func (_TNT1155VoucherContract *TNT1155VoucherContractFilterer) WatchUpdateOwner(opts *bind.WatchOpts, sink chan<- *TNT1155VoucherContractUpdateOwner) (event.Subscription, error) {

	logs, sub, err := _TNT1155VoucherContract.contract.WatchLogs(opts, "UpdateOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TNT1155VoucherContractUpdateOwner)
				if err := _TNT1155VoucherContract.contract.UnpackLog(event, "UpdateOwner", log); err != nil {
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
func (_TNT1155VoucherContract *TNT1155VoucherContractFilterer) ParseUpdateOwner(log types.Log) (*TNT1155VoucherContractUpdateOwner, error) {
	event := new(TNT1155VoucherContractUpdateOwner)
	if err := _TNT1155VoucherContract.contract.UnpackLog(event, "UpdateOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
