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
	Bin: "0x60806040523480156200001157600080fd5b5060405162002384380380620023848339810160408190526200003491620001f7565b8281620000418162000081565b50600380546001600160a01b0319166001600160a01b03929092169190911790558151620000779060049060208501906200009a565b50505050620002d4565b8051620000969060029060208401906200009a565b5050565b828054620000a89062000281565b90600052602060002090601f016020900481019282620000cc576000855562000117565b82601f10620000e757805160ff191683800117855562000117565b8280016001018555821562000117579182015b8281111562000117578251825591602001919060010190620000fa565b506200012592915062000129565b5090565b5b808211156200012557600081556001016200012a565b600082601f8301126200015257600080fd5b81516001600160401b03808211156200016f576200016f620002be565b604051601f8301601f19908116603f011681019082821181831017156200019a576200019a620002be565b81604052838152602092508683858801011115620001b757600080fd5b600091505b83821015620001db5785820183015181830184015290820190620001bc565b83821115620001ed5760008385830101525b9695505050505050565b6000806000606084860312156200020d57600080fd5b83516001600160a01b03811681146200022557600080fd5b60208501519093506001600160401b03808211156200024357600080fd5b620002518783880162000140565b935060408601519150808211156200026857600080fd5b50620002778682870162000140565b9150509250925092565b600181811c908216806200029657607f821691505b60208210811415620002b857634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052604160045260246000fd5b6120a080620002e46000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c8063731133e911610092578063731133e9146101d8578063880cdc31146101eb5780638da5cb5b146101fe578063a22cb46514610211578063c370b04214610224578063c87b56dd1461022c578063e985e9c51461023f578063f242432a14610252578063f5298aca1461026557600080fd5b8062fdd58e146100ef57806301ffc9a7146101155780630e89341c146101385780631f7fdffa146101585780632eb2c2d61461016d578063442890d5146101805780634e1273f4146101a55780636b20c454146101c5575b600080fd5b6101026100fd366004611861565b610278565b6040519081526020015b60405180910390f35b6101286101233660046119e2565b61030f565b604051901515815260200161010c565b61014b610146366004611a1c565b610361565b60405161010c9190611ba1565b61016b61016636600461178d565b6103f5565b005b61016b61017b366004611600565b610431565b6003546001600160a01b03165b6040516001600160a01b03909116815260200161010c565b6101b86101b3366004611912565b6104c8565b60405161010c9190611b60565b61016b6101d336600461170d565b6105f1565b61016b6101e63660046118be565b61073a565b61016b6101f93660046115ab565b610770565b60035461018d906001600160a01b031681565b61016b61021f366004611825565b610803565b61014b610812565b61014b61023a366004611a1c565b6108a4565b61012861024d3660046115cd565b6108af565b61016b6102603660046116a9565b6108dd565b61016b61027336600461188b565b610964565b60006001600160a01b0383166102e95760405162461bcd60e51b815260206004820152602b60248201527f455243313135353a2062616c616e636520717565727920666f7220746865207a60448201526a65726f206164647265737360a81b60648201526084015b60405180910390fd5b506000908152602081815260408083206001600160a01b03949094168352929052205490565b60006001600160e01b03198216636cdb3d1360e11b148061034057506001600160e01b031982166303a24d0760e21b145b8061035b57506301ffc9a760e01b6001600160e01b03198316145b92915050565b60606002805461037090611ea8565b80601f016020809104026020016040519081016040528092919081815260200182805461039c90611ea8565b80156103e95780601f106103be576101008083540402835291602001916103e9565b820191906000526020600020905b8154815290600101906020018083116103cc57829003601f168201915b50505050509050919050565b6003546001600160a01b0316331461041f5760405162461bcd60e51b81526004016102e090611d5b565b61042b848484846109f4565b50505050565b6001600160a01b03851633148061044d575061044d85336108af565b6104b45760405162461bcd60e51b815260206004820152603260248201527f455243313135353a207472616e736665722063616c6c6572206973206e6f74206044820152711bdddb995c881b9bdc88185c1c1c9bdd995960721b60648201526084016102e0565b6104c18585858585610b2d565b5050505050565b6060815183511461052d5760405162461bcd60e51b815260206004820152602960248201527f455243313135353a206163636f756e747320616e6420696473206c656e677468604482015268040dad2e6dac2e8c6d60bb1b60648201526084016102e0565b600083516001600160401b0381111561054857610548611f56565b604051908082528060200260200182016040528015610571578160200160208202803683370190505b50905060005b84518110156105e9576105bc85828151811061059557610595611f40565b60200260200101518583815181106105af576105af611f40565b6020026020010151610278565b8282815181106105ce576105ce611f40565b60209081029190910101526105e281611f0f565b9050610577565b509392505050565b6003546001600160a01b0316331461061b5760405162461bcd60e51b81526004016102e090611d5b565b6003546001600160a01b031661063186826108af565b61064d5760405162461bcd60e51b81526004016102e090611c85565b60005b848110156106c35783838281811061066a5761066a611f40565b905060200201356106938888888581811061068757610687611f40565b90506020020135610278565b10156106b15760405162461bcd60e51b81526004016102e090611e1b565b806106bb81611f0f565b915050610650565b506107328686868080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525050604080516020808a02828101820190935289825290935089925088918291850190849080828437600092019190915250610caf92505050565b505050505050565b6003546001600160a01b031633146107645760405162461bcd60e51b81526004016102e090611d5b565b61042b84848484610e27565b6003546001600160a01b0316331461079a5760405162461bcd60e51b81526004016102e090611d5b565b600354604080516001600160a01b03928316815291831660208301527fe2c7d1c4da37855e682bde14f17826d185497973b73fba7554daa6da467058d9910160405180910390a1600380546001600160a01b0319166001600160a01b0392909216919091179055565b61080e338383610eef565b5050565b60606004805461082190611ea8565b80601f016020809104026020016040519081016040528092919081815260200182805461084d90611ea8565b801561089a5780601f1061086f5761010080835404028352916020019161089a565b820191906000526020600020905b81548152906001019060200180831161087d57829003601f168201915b5050505050905090565b606061035b82610361565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205460ff1690565b6001600160a01b0385163314806108f957506108f985336108af565b6109575760405162461bcd60e51b815260206004820152602960248201527f455243313135353a2063616c6c6572206973206e6f74206f776e6572206e6f7260448201526808185c1c1c9bdd995960ba1b60648201526084016102e0565b6104c18585858585610fd0565b6003546001600160a01b0316331461098e5760405162461bcd60e51b81526004016102e090611d5b565b806109998484610278565b10156109b75760405162461bcd60e51b81526004016102e090611e1b565b6003546001600160a01b03166109cd84826108af565b6109e95760405162461bcd60e51b81526004016102e090611c85565b61042b8484846110e8565b6001600160a01b038416610a1a5760405162461bcd60e51b81526004016102e090611dda565b8151835114610a3b5760405162461bcd60e51b81526004016102e090611d92565b3360005b8451811015610ad757838181518110610a5a57610a5a611f40565b6020026020010151600080878481518110610a7757610a77611f40565b602002602001015181526020019081526020016000206000886001600160a01b03166001600160a01b031681526020019081526020016000206000828254610abf9190611e90565b90915550819050610acf81611f0f565b915050610a3f565b50846001600160a01b031660006001600160a01b0316826001600160a01b031660008051602061202b8339815191528787604051610b16929190611b73565b60405180910390a46104c1816000878787876111da565b8151835114610b4e5760405162461bcd60e51b81526004016102e090611d92565b6001600160a01b038416610b745760405162461bcd60e51b81526004016102e090611c40565b3360005b8451811015610c5b576000858281518110610b9557610b95611f40565b602002602001015190506000858381518110610bb357610bb3611f40565b602090810291909101810151600084815280835260408082206001600160a01b038e168352909352919091205490915081811015610c035760405162461bcd60e51b81526004016102e090611d11565b6000838152602081815260408083206001600160a01b038e8116855292528083208585039055908b16825281208054849290610c40908490611e90565b9250508190555050505080610c5490611f0f565b9050610b78565b50846001600160a01b0316866001600160a01b0316826001600160a01b031660008051602061202b8339815191528787604051610c99929190611b73565b60405180910390a46107328187878787876111da565b6001600160a01b038316610cd55760405162461bcd60e51b81526004016102e090611cce565b8051825114610cf65760405162461bcd60e51b81526004016102e090611d92565b604080516020810190915260009081905233905b8351811015610dcc576000848281518110610d2757610d27611f40565b602002602001015190506000848381518110610d4557610d45611f40565b602090810291909101810151600084815280835260408082206001600160a01b038c168352909352919091205490915081811015610d955760405162461bcd60e51b81526004016102e090611bfc565b6000928352602083815260408085206001600160a01b038b1686529091529092209103905580610dc481611f0f565b915050610d0a565b5060006001600160a01b0316846001600160a01b0316826001600160a01b031660008051602061202b8339815191528686604051610e0b929190611b73565b60405180910390a460408051602081019091526000905261042b565b6001600160a01b038416610e4d5760405162461bcd60e51b81526004016102e090611dda565b336000610e5985611345565b90506000610e6685611345565b90506000868152602081815260408083206001600160a01b038b16845290915281208054879290610e98908490611e90565b909155505060408051878152602081018790526001600160a01b03808a16926000929187169160008051602061204b833981519152910160405180910390a4610ee683600089898989611390565b50505050505050565b816001600160a01b0316836001600160a01b03161415610f635760405162461bcd60e51b815260206004820152602960248201527f455243313135353a2073657474696e6720617070726f76616c20737461747573604482015268103337b91039b2b63360b91b60648201526084016102e0565b6001600160a01b03838116600081815260016020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b6001600160a01b038416610ff65760405162461bcd60e51b81526004016102e090611c40565b33600061100285611345565b9050600061100f85611345565b90506000868152602081815260408083206001600160a01b038c168452909152902054858110156110525760405162461bcd60e51b81526004016102e090611d11565b6000878152602081815260408083206001600160a01b038d8116855292528083208985039055908a1682528120805488929061108f908490611e90565b909155505060408051888152602081018890526001600160a01b03808b16928c8216929188169160008051602061204b833981519152910160405180910390a46110dd848a8a8a8a8a611390565b505050505050505050565b6001600160a01b03831661110e5760405162461bcd60e51b81526004016102e090611cce565b33600061111a84611345565b9050600061112784611345565b60408051602080820183526000918290528882528181528282206001600160a01b038b16835290522054909150848110156111745760405162461bcd60e51b81526004016102e090611bfc565b6000868152602081815260408083206001600160a01b038b81168086529184528285208a8703905582518b81529384018a905290929088169160008051602061204b833981519152910160405180910390a4604080516020810190915260009052610ee6565b6001600160a01b0384163b156107325760405163bc197c8160e01b81526001600160a01b0385169063bc197c819061121e9089908990889088908890600401611abd565b602060405180830381600087803b15801561123857600080fd5b505af1925050508015611268575060408051601f3d908101601f19168201909252611265918101906119ff565b60015b61131557611274611f6c565b806308c379a014156112ae5750611289611f88565b8061129457506112b0565b8060405162461bcd60e51b81526004016102e09190611ba1565b505b60405162461bcd60e51b815260206004820152603460248201527f455243313135353a207472616e7366657220746f206e6f6e20455243313135356044820152732932b1b2b4bb32b91034b6b83632b6b2b73a32b960611b60648201526084016102e0565b6001600160e01b0319811663bc197c8160e01b14610ee65760405162461bcd60e51b81526004016102e090611bb4565b6040805160018082528183019092526060916000919060208083019080368337019050509050828160008151811061137f5761137f611f40565b602090810291909101015292915050565b6001600160a01b0384163b156107325760405163f23a6e6160e01b81526001600160a01b0385169063f23a6e61906113d49089908990889088908890600401611b1b565b602060405180830381600087803b1580156113ee57600080fd5b505af192505050801561141e575060408051601f3d908101601f1916820190925261141b918101906119ff565b60015b61142a57611274611f6c565b6001600160e01b0319811663f23a6e6160e01b14610ee65760405162461bcd60e51b81526004016102e090611bb4565b80356001600160a01b038116811461147157600080fd5b919050565b60008083601f84011261148857600080fd5b5081356001600160401b0381111561149f57600080fd5b6020830191508360208260051b85010111156114ba57600080fd5b9250929050565b600082601f8301126114d257600080fd5b813560206114df82611e6d565b6040516114ec8282611ee3565b8381528281019150858301600585901b8701840188101561150c57600080fd5b60005b8581101561152b5781358452928401929084019060010161150f565b5090979650505050505050565b600082601f83011261154957600080fd5b81356001600160401b0381111561156257611562611f56565b604051611579601f8301601f191660200182611ee3565b81815284602083860101111561158e57600080fd5b816020850160208301376000918101602001919091529392505050565b6000602082840312156115bd57600080fd5b6115c68261145a565b9392505050565b600080604083850312156115e057600080fd5b6115e98361145a565b91506115f76020840161145a565b90509250929050565b600080600080600060a0868803121561161857600080fd5b6116218661145a565b945061162f6020870161145a565b935060408601356001600160401b038082111561164b57600080fd5b61165789838a016114c1565b9450606088013591508082111561166d57600080fd5b61167989838a016114c1565b9350608088013591508082111561168f57600080fd5b5061169c88828901611538565b9150509295509295909350565b600080600080600060a086880312156116c157600080fd5b6116ca8661145a565b94506116d86020870161145a565b9350604086013592506060860135915060808601356001600160401b0381111561170157600080fd5b61169c88828901611538565b60008060008060006060868803121561172557600080fd5b61172e8661145a565b945060208601356001600160401b038082111561174a57600080fd5b61175689838a01611476565b9096509450604088013591508082111561176f57600080fd5b5061177c88828901611476565b969995985093965092949392505050565b600080600080608085870312156117a357600080fd5b6117ac8561145a565b935060208501356001600160401b03808211156117c857600080fd5b6117d4888389016114c1565b945060408701359150808211156117ea57600080fd5b6117f6888389016114c1565b9350606087013591508082111561180c57600080fd5b5061181987828801611538565b91505092959194509250565b6000806040838503121561183857600080fd5b6118418361145a565b91506020830135801515811461185657600080fd5b809150509250929050565b6000806040838503121561187457600080fd5b61187d8361145a565b946020939093013593505050565b6000806000606084860312156118a057600080fd5b6118a98461145a565b95602085013595506040909401359392505050565b600080600080608085870312156118d457600080fd5b6118dd8561145a565b9350602085013592506040850135915060608501356001600160401b0381111561190657600080fd5b61181987828801611538565b6000806040838503121561192557600080fd5b82356001600160401b038082111561193c57600080fd5b818501915085601f83011261195057600080fd5b8135602061195d82611e6d565b60405161196a8282611ee3565b8381528281019150858301600585901b870184018b101561198a57600080fd5b600096505b848710156119b4576119a08161145a565b83526001969096019591830191830161198f565b50965050860135925050808211156119cb57600080fd5b506119d8858286016114c1565b9150509250929050565b6000602082840312156119f457600080fd5b81356115c681612011565b600060208284031215611a1157600080fd5b81516115c681612011565b600060208284031215611a2e57600080fd5b5035919050565b600081518084526020808501945080840160005b83811015611a6557815187529582019590820190600101611a49565b509495945050505050565b6000815180845260005b81811015611a9657602081850181015186830182015201611a7a565b81811115611aa8576000602083870101525b50601f01601f19169290920160200192915050565b6001600160a01b0386811682528516602082015260a060408201819052600090611ae990830186611a35565b8281036060840152611afb8186611a35565b90508281036080840152611b0f8185611a70565b98975050505050505050565b6001600160a01b03868116825285166020820152604081018490526060810183905260a060808201819052600090611b5590830184611a70565b979650505050505050565b6020815260006115c66020830184611a35565b604081526000611b866040830185611a35565b8281036020840152611b988185611a35565b95945050505050565b6020815260006115c66020830184611a70565b60208082526028908201527f455243313135353a204552433131353552656365697665722072656a656374656040820152676420746f6b656e7360c01b606082015260800190565b60208082526024908201527f455243313135353a206275726e20616d6f756e7420657863656564732062616c604082015263616e636560e01b606082015260800190565b60208082526025908201527f455243313135353a207472616e7366657220746f20746865207a65726f206164604082015264647265737360d81b606082015260800190565b60208082526029908201527f566f7563686572206f776e657220646964206e6f7420617070726f766520796f6040820152683a903a3790313ab93760b91b606082015260800190565b60208082526023908201527f455243313135353a206275726e2066726f6d20746865207a65726f206164647260408201526265737360e81b606082015260800190565b6020808252602a908201527f455243313135353a20696e73756666696369656e742062616c616e636520666f60408201526939103a3930b739b332b960b11b606082015260800190565b6020808252601c908201527f6f6e6c79206f776e65722063616e206d616b65207468652063616c6c00000000604082015260600190565b60208082526028908201527f455243313135353a2069647320616e6420616d6f756e7473206c656e677468206040820152670dad2e6dac2e8c6d60c31b606082015260800190565b60208082526021908201527f455243313135353a206d696e7420746f20746865207a65726f206164647265736040820152607360f81b606082015260800190565b60208082526032908201527f566f7563686572206f776e657220646f6573206e6f74206861766520656e6f7560408201527133b4103130b630b731b2903a3790313ab93760711b606082015260800190565b60006001600160401b03821115611e8657611e86611f56565b5060051b60200190565b60008219821115611ea357611ea3611f2a565b500190565b600181811c90821680611ebc57607f821691505b60208210811415611edd57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8201601f191681016001600160401b0381118282101715611f0857611f08611f56565b6040525050565b6000600019821415611f2357611f23611f2a565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b600060033d1115611f855760046000803e5060005160e01c5b90565b600060443d1015611f965790565b6040516003193d81016004833e81513d6001600160401b038160248401118184111715611fc557505050505090565b8285019150815181811115611fdd5750505050505090565b843d8701016020828501011115611ff75750505050505090565b61200660208286010187611ee3565b509095945050505050565b6001600160e01b03198116811461202757600080fd5b5056fe4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fbc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62a26469706673582212203dae927920410fa1cb299b6b6fd9d75884b3bebf4a74c1845bf64a000437222a64736f6c63430008070033",
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
