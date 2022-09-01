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

// MockTNT1155MetaData contains all meta data concerning the MockTNT1155 contract.
var MockTNT1155MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040518060600160405280602d815260200162001c4d602d9139620000378162000077565b5062000071732e833968e5bb786ae419c4d13189fb081cc43bab60016064604051806020016040528060008152506200009060201b60201c565b620006da565b80516200008c906002906020840190620003f4565b5050565b6001600160a01b038416620000f65760405162461bcd60e51b815260206004820152602160248201527f455243313135353a206d696e7420746f20746865207a65726f206164647265736044820152607360f81b60648201526084015b60405180910390fd5b3360006200010485620001b2565b905060006200011385620001b2565b90506000868152602081815260408083206001600160a01b038b168452909152812080548792906200014790849062000579565b909155505060408051878152602081018790526001600160a01b03808a1692600092918716917fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62910160405180910390a4620001a98360008989898962000208565b50505050505050565b60408051600180825281830190925260609160009190602080830190803683370190505090508281600081518110620001ef57620001ef62000618565b602090810291909101015292915050565b505050505050565b62000227846001600160a01b0316620003e560201b620005021760201c565b15620002005760405163f23a6e6160e01b81526001600160a01b0385169063f23a6e61906200026390899089908890889088906004016200051d565b602060405180830381600087803b1580156200027e57600080fd5b505af1925050508015620002b1575060408051601f3d908101601f19168201909252620002ae918101906200049a565b60015b6200037257620002c06200062e565b806308c379a01415620003015750620002d86200064b565b80620002e5575062000303565b8060405162461bcd60e51b8152600401620000ed919062000564565b505b60405162461bcd60e51b815260206004820152603460248201527f455243313135353a207472616e7366657220746f206e6f6e204552433131353560448201527f526563656976657220696d706c656d656e7465720000000000000000000000006064820152608401620000ed565b6001600160e01b0319811663f23a6e6160e01b14620001a95760405162461bcd60e51b815260206004820152602860248201527f455243313135353a204552433131353552656365697665722072656a656374656044820152676420746f6b656e7360c01b6064820152608401620000ed565b6001600160a01b03163b151590565b8280546200040290620005a0565b90600052602060002090601f01602090048101928262000426576000855562000471565b82601f106200044157805160ff191683800117855562000471565b8280016001018555821562000471579182015b828111156200047157825182559160200191906001019062000454565b506200047f92915062000483565b5090565b5b808211156200047f576000815560010162000484565b600060208284031215620004ad57600080fd5b81516001600160e01b031981168114620004c657600080fd5b9392505050565b6000815180845260005b81811015620004f557602081850181015186830182015201620004d7565b8181111562000508576000602083870101525b50601f01601f19169290920160200192915050565b6001600160a01b03868116825285166020820152604081018490526060810183905260a0608082018190526000906200055990830184620004cd565b979650505050505050565b602081526000620004c66020830184620004cd565b600082198211156200059b57634e487b7160e01b600052601160045260246000fd5b500190565b600181811c90821680620005b557607f821691505b60208210811415620005d757634e487b7160e01b600052602260045260246000fd5b50919050565b601f8201601f191681016001600160401b03811182821017156200061157634e487b7160e01b600052604160045260246000fd5b6040525050565b634e487b7160e01b600052603260045260246000fd5b600060033d1115620006485760046000803e5060005160e01c5b90565b600060443d10156200065a5790565b6040516003193d81016004833e81513d6001600160401b0380831160248401831017156200068a57505050505090565b8285019150815181811115620006a35750505050505090565b843d8701016020828501011115620006be5750505050505090565b620006cf60208286010187620005dd565b509095945050505050565b61156380620006ea6000396000f3fe608060405234801561001057600080fd5b50600436106100925760003560e01c80634e1273f4116100665780634e1273f414610115578063731133e914610135578063a22cb46514610148578063e985e9c51461015b578063f242432a1461019757600080fd5b8062fdd58e1461009757806301ffc9a7146100bd5780630e89341c146100e05780632eb2c2d614610100575b600080fd5b6100aa6100a5366004610f12565b6101aa565b6040519081526020015b60405180910390f35b6100d06100cb36600461106e565b610240565b60405190151581526020016100b4565b6100f36100ee3660046110af565b610292565b6040516100b49190611234565b61011361010e366004610dc7565b610326565b005b610128610123366004610f9d565b610372565b6040516100b491906111f3565b610113610143366004610f3c565b61049c565b610113610156366004610ed6565b6104ae565b6100d0610169366004610d94565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205460ff1690565b6101136101a5366004610e71565b6104bd565b60006001600160a01b03831661021a5760405162461bcd60e51b815260206004820152602a60248201527f455243313135353a2061646472657373207a65726f206973206e6f742061207660448201526930b634b21037bbb732b960b11b60648201526084015b60405180910390fd5b506000908152602081815260408083206001600160a01b03949094168352929052205490565b60006001600160e01b03198216636cdb3d1360e11b148061027157506001600160e01b031982166303a24d0760e21b145b8061028c57506301ffc9a760e01b6001600160e01b03198316145b92915050565b6060600280546102a1906113a9565b80601f01602080910402602001604051908101604052809291908181526020018280546102cd906113a9565b801561031a5780601f106102ef5761010080835404028352916020019161031a565b820191906000526020600020905b8154815290600101906020018083116102fd57829003601f168201915b50505050509050919050565b6001600160a01b03851633148061034257506103428533610169565b61035e5760405162461bcd60e51b815260040161021190611247565b61036b8585858585610511565b5050505050565b606081518351146103d75760405162461bcd60e51b815260206004820152602960248201527f455243313135353a206163636f756e747320616e6420696473206c656e677468604482015268040dad2e6dac2e8c6d60bb1b6064820152608401610211565b6000835167ffffffffffffffff8111156103f3576103f3611458565b60405190808252806020026020018201604052801561041c578160200160208202803683370190505b50905060005b84518110156104945761046785828151811061044057610440611442565b602002602001015185838151811061045a5761045a611442565b60200260200101516101aa565b82828151811061047957610479611442565b602090810291909101015261048d81611411565b9050610422565b509392505050565b6104a8848484846106ee565b50505050565b6104b9338383610802565b5050565b6001600160a01b0385163314806104d957506104d98533610169565b6104f55760405162461bcd60e51b815260040161021190611247565b61036b85858585856108e3565b6001600160a01b03163b151590565b81518351146105735760405162461bcd60e51b815260206004820152602860248201527f455243313135353a2069647320616e6420616d6f756e7473206c656e677468206044820152670dad2e6dac2e8c6d60c31b6064820152608401610211565b6001600160a01b0384166105995760405162461bcd60e51b8152600401610211906112de565b3360005b84518110156106805760008582815181106105ba576105ba611442565b6020026020010151905060008583815181106105d8576105d8611442565b602090810291909101810151600084815280835260408082206001600160a01b038e1683529093529190912054909150818110156106285760405162461bcd60e51b815260040161021190611323565b6000838152602081815260408083206001600160a01b038e8116855292528083208585039055908b16825281208054849290610665908490611391565b925050819055505050508061067990611411565b905061059d565b50846001600160a01b0316866001600160a01b0316826001600160a01b03167f4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb87876040516106d0929190611206565b60405180910390a46106e6818787878787610a0d565b505050505050565b6001600160a01b03841661074e5760405162461bcd60e51b815260206004820152602160248201527f455243313135353a206d696e7420746f20746865207a65726f206164647265736044820152607360f81b6064820152608401610211565b33600061075a85610b78565b9050600061076785610b78565b90506000868152602081815260408083206001600160a01b038b16845290915281208054879290610799908490611391565b909155505060408051878152602081018790526001600160a01b03808a1692600092918716917fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62910160405180910390a46107f983600089898989610bc3565b50505050505050565b816001600160a01b0316836001600160a01b031614156108765760405162461bcd60e51b815260206004820152602960248201527f455243313135353a2073657474696e6720617070726f76616c20737461747573604482015268103337b91039b2b63360b91b6064820152608401610211565b6001600160a01b03838116600081815260016020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b6001600160a01b0384166109095760405162461bcd60e51b8152600401610211906112de565b33600061091585610b78565b9050600061092285610b78565b90506000868152602081815260408083206001600160a01b038c168452909152902054858110156109655760405162461bcd60e51b815260040161021190611323565b6000878152602081815260408083206001600160a01b038d8116855292528083208985039055908a168252812080548892906109a2908490611391565b909155505060408051888152602081018890526001600160a01b03808b16928c821692918816917fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62910160405180910390a4610a02848a8a8a8a8a610bc3565b505050505050505050565b6001600160a01b0384163b156106e65760405163bc197c8160e01b81526001600160a01b0385169063bc197c8190610a519089908990889088908890600401611150565b602060405180830381600087803b158015610a6b57600080fd5b505af1925050508015610a9b575060408051601f3d908101601f19168201909252610a9891810190611092565b60015b610b4857610aa761146e565b806308c379a01415610ae15750610abc61148a565b80610ac75750610ae3565b8060405162461bcd60e51b81526004016102119190611234565b505b60405162461bcd60e51b815260206004820152603460248201527f455243313135353a207472616e7366657220746f206e6f6e20455243313135356044820152732932b1b2b4bb32b91034b6b83632b6b2b73a32b960611b6064820152608401610211565b6001600160e01b0319811663bc197c8160e01b146107f95760405162461bcd60e51b815260040161021190611296565b60408051600180825281830190925260609160009190602080830190803683370190505090508281600081518110610bb257610bb2611442565b602090810291909101015292915050565b6001600160a01b0384163b156106e65760405163f23a6e6160e01b81526001600160a01b0385169063f23a6e6190610c0790899089908890889088906004016111ae565b602060405180830381600087803b158015610c2157600080fd5b505af1925050508015610c51575060408051601f3d908101601f19168201909252610c4e91810190611092565b60015b610c5d57610aa761146e565b6001600160e01b0319811663f23a6e6160e01b146107f95760405162461bcd60e51b815260040161021190611296565b80356001600160a01b0381168114610ca457600080fd5b919050565b600082601f830112610cba57600080fd5b81356020610cc78261136d565b604051610cd482826113e4565b8381528281019150858301600585901b87018401881015610cf457600080fd5b60005b85811015610d1357813584529284019290840190600101610cf7565b5090979650505050505050565b600082601f830112610d3157600080fd5b813567ffffffffffffffff811115610d4b57610d4b611458565b604051610d62601f8301601f1916602001826113e4565b818152846020838601011115610d7757600080fd5b816020850160208301376000918101602001919091529392505050565b60008060408385031215610da757600080fd5b610db083610c8d565b9150610dbe60208401610c8d565b90509250929050565b600080600080600060a08688031215610ddf57600080fd5b610de886610c8d565b9450610df660208701610c8d565b9350604086013567ffffffffffffffff80821115610e1357600080fd5b610e1f89838a01610ca9565b94506060880135915080821115610e3557600080fd5b610e4189838a01610ca9565b93506080880135915080821115610e5757600080fd5b50610e6488828901610d20565b9150509295509295909350565b600080600080600060a08688031215610e8957600080fd5b610e9286610c8d565b9450610ea060208701610c8d565b93506040860135925060608601359150608086013567ffffffffffffffff811115610eca57600080fd5b610e6488828901610d20565b60008060408385031215610ee957600080fd5b610ef283610c8d565b915060208301358015158114610f0757600080fd5b809150509250929050565b60008060408385031215610f2557600080fd5b610f2e83610c8d565b946020939093013593505050565b60008060008060808587031215610f5257600080fd5b610f5b85610c8d565b93506020850135925060408501359150606085013567ffffffffffffffff811115610f8557600080fd5b610f9187828801610d20565b91505092959194509250565b60008060408385031215610fb057600080fd5b823567ffffffffffffffff80821115610fc857600080fd5b818501915085601f830112610fdc57600080fd5b81356020610fe98261136d565b604051610ff682826113e4565b8381528281019150858301600585901b870184018b101561101657600080fd5b600096505b848710156110405761102c81610c8d565b83526001969096019591830191830161101b565b509650508601359250508082111561105757600080fd5b5061106485828601610ca9565b9150509250929050565b60006020828403121561108057600080fd5b813561108b81611514565b9392505050565b6000602082840312156110a457600080fd5b815161108b81611514565b6000602082840312156110c157600080fd5b5035919050565b600081518084526020808501945080840160005b838110156110f8578151875295820195908201906001016110dc565b509495945050505050565b6000815180845260005b818110156111295760208185018101518683018201520161110d565b8181111561113b576000602083870101525b50601f01601f19169290920160200192915050565b6001600160a01b0386811682528516602082015260a06040820181905260009061117c908301866110c8565b828103606084015261118e81866110c8565b905082810360808401526111a28185611103565b98975050505050505050565b6001600160a01b03868116825285166020820152604081018490526060810183905260a0608082018190526000906111e890830184611103565b979650505050505050565b60208152600061108b60208301846110c8565b60408152600061121960408301856110c8565b828103602084015261122b81856110c8565b95945050505050565b60208152600061108b6020830184611103565b6020808252602f908201527f455243313135353a2063616c6c6572206973206e6f7420746f6b656e206f776e60408201526e195c881b9bdc88185c1c1c9bdd9959608a1b606082015260800190565b60208082526028908201527f455243313135353a204552433131353552656365697665722072656a656374656040820152676420746f6b656e7360c01b606082015260800190565b60208082526025908201527f455243313135353a207472616e7366657220746f20746865207a65726f206164604082015264647265737360d81b606082015260800190565b6020808252602a908201527f455243313135353a20696e73756666696369656e742062616c616e636520666f60408201526939103a3930b739b332b960b11b606082015260800190565b600067ffffffffffffffff82111561138757611387611458565b5060051b60200190565b600082198211156113a4576113a461142c565b500190565b600181811c908216806113bd57607f821691505b602082108114156113de57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8201601f1916810167ffffffffffffffff8111828210171561140a5761140a611458565b6040525050565b60006000198214156114255761142561142c565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b600060033d11156114875760046000803e5060005160e01c5b90565b600060443d10156114985790565b6040516003193d81016004833e81513d67ffffffffffffffff81602484011181841117156114c857505050505090565b82850191508151818111156114e05750505050505090565b843d87010160208285010111156114fa5750505050505090565b611509602082860101876113e4565b509095945050505050565b6001600160e01b03198116811461152a57600080fd5b5056fea2646970667358221220097210992fd092ae52d3f64c04e6bdd4121739ee959594af830490ec4ca3482864736f6c6343000807003368747470733a2f2f7468657461746f6b656e2e6578616d706c652f6170692f6974656d2f7b69647d2e6a736f6e",
}

// MockTNT1155ABI is the input ABI used to generate the binding from.
// Deprecated: Use MockTNT1155MetaData.ABI instead.
var MockTNT1155ABI = MockTNT1155MetaData.ABI

// MockTNT1155Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MockTNT1155MetaData.Bin instead.
var MockTNT1155Bin = MockTNT1155MetaData.Bin

// DeployMockTNT1155 deploys a new Ethereum contract, binding an instance of MockTNT1155 to it.
func DeployMockTNT1155(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MockTNT1155, error) {
	parsed, err := MockTNT1155MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockTNT1155Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockTNT1155{MockTNT1155Caller: MockTNT1155Caller{contract: contract}, MockTNT1155Transactor: MockTNT1155Transactor{contract: contract}, MockTNT1155Filterer: MockTNT1155Filterer{contract: contract}}, nil
}

// MockTNT1155 is an auto generated Go binding around an Ethereum contract.
type MockTNT1155 struct {
	MockTNT1155Caller     // Read-only binding to the contract
	MockTNT1155Transactor // Write-only binding to the contract
	MockTNT1155Filterer   // Log filterer for contract events
}

// MockTNT1155Caller is an auto generated read-only Go binding around an Ethereum contract.
type MockTNT1155Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockTNT1155Transactor is an auto generated write-only Go binding around an Ethereum contract.
type MockTNT1155Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockTNT1155Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockTNT1155Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockTNT1155Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockTNT1155Session struct {
	Contract     *MockTNT1155      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MockTNT1155CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockTNT1155CallerSession struct {
	Contract *MockTNT1155Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MockTNT1155TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockTNT1155TransactorSession struct {
	Contract     *MockTNT1155Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MockTNT1155Raw is an auto generated low-level Go binding around an Ethereum contract.
type MockTNT1155Raw struct {
	Contract *MockTNT1155 // Generic contract binding to access the raw methods on
}

// MockTNT1155CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockTNT1155CallerRaw struct {
	Contract *MockTNT1155Caller // Generic read-only contract binding to access the raw methods on
}

// MockTNT1155TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockTNT1155TransactorRaw struct {
	Contract *MockTNT1155Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMockTNT1155 creates a new instance of MockTNT1155, bound to a specific deployed contract.
func NewMockTNT1155(address common.Address, backend bind.ContractBackend) (*MockTNT1155, error) {
	contract, err := bindMockTNT1155(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockTNT1155{MockTNT1155Caller: MockTNT1155Caller{contract: contract}, MockTNT1155Transactor: MockTNT1155Transactor{contract: contract}, MockTNT1155Filterer: MockTNT1155Filterer{contract: contract}}, nil
}

// NewMockTNT1155Caller creates a new read-only instance of MockTNT1155, bound to a specific deployed contract.
func NewMockTNT1155Caller(address common.Address, caller bind.ContractCaller) (*MockTNT1155Caller, error) {
	contract, err := bindMockTNT1155(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockTNT1155Caller{contract: contract}, nil
}

// NewMockTNT1155Transactor creates a new write-only instance of MockTNT1155, bound to a specific deployed contract.
func NewMockTNT1155Transactor(address common.Address, transactor bind.ContractTransactor) (*MockTNT1155Transactor, error) {
	contract, err := bindMockTNT1155(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockTNT1155Transactor{contract: contract}, nil
}

// NewMockTNT1155Filterer creates a new log filterer instance of MockTNT1155, bound to a specific deployed contract.
func NewMockTNT1155Filterer(address common.Address, filterer bind.ContractFilterer) (*MockTNT1155Filterer, error) {
	contract, err := bindMockTNT1155(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockTNT1155Filterer{contract: contract}, nil
}

// bindMockTNT1155 binds a generic wrapper to an already deployed contract.
func bindMockTNT1155(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MockTNT1155ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockTNT1155 *MockTNT1155Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockTNT1155.Contract.MockTNT1155Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockTNT1155 *MockTNT1155Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockTNT1155.Contract.MockTNT1155Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockTNT1155 *MockTNT1155Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockTNT1155.Contract.MockTNT1155Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockTNT1155 *MockTNT1155CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockTNT1155.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockTNT1155 *MockTNT1155TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockTNT1155.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockTNT1155 *MockTNT1155TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockTNT1155.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_MockTNT1155 *MockTNT1155Caller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MockTNT1155.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_MockTNT1155 *MockTNT1155Session) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _MockTNT1155.Contract.BalanceOf(&_MockTNT1155.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_MockTNT1155 *MockTNT1155CallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _MockTNT1155.Contract.BalanceOf(&_MockTNT1155.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_MockTNT1155 *MockTNT1155Caller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _MockTNT1155.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_MockTNT1155 *MockTNT1155Session) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _MockTNT1155.Contract.BalanceOfBatch(&_MockTNT1155.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_MockTNT1155 *MockTNT1155CallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _MockTNT1155.Contract.BalanceOfBatch(&_MockTNT1155.CallOpts, accounts, ids)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_MockTNT1155 *MockTNT1155Caller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _MockTNT1155.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_MockTNT1155 *MockTNT1155Session) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _MockTNT1155.Contract.IsApprovedForAll(&_MockTNT1155.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_MockTNT1155 *MockTNT1155CallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _MockTNT1155.Contract.IsApprovedForAll(&_MockTNT1155.CallOpts, account, operator)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MockTNT1155 *MockTNT1155Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _MockTNT1155.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MockTNT1155 *MockTNT1155Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MockTNT1155.Contract.SupportsInterface(&_MockTNT1155.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MockTNT1155 *MockTNT1155CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MockTNT1155.Contract.SupportsInterface(&_MockTNT1155.CallOpts, interfaceId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_MockTNT1155 *MockTNT1155Caller) Uri(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _MockTNT1155.contract.Call(opts, &out, "uri", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_MockTNT1155 *MockTNT1155Session) Uri(arg0 *big.Int) (string, error) {
	return _MockTNT1155.Contract.Uri(&_MockTNT1155.CallOpts, arg0)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_MockTNT1155 *MockTNT1155CallerSession) Uri(arg0 *big.Int) (string, error) {
	return _MockTNT1155.Contract.Uri(&_MockTNT1155.CallOpts, arg0)
}

// Mint is a paid mutator transaction binding the contract method 0x731133e9.
//
// Solidity: function mint(address to, uint256 id, uint256 amount, bytes data) returns()
func (_MockTNT1155 *MockTNT1155Transactor) Mint(opts *bind.TransactOpts, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _MockTNT1155.contract.Transact(opts, "mint", to, id, amount, data)
}

// Mint is a paid mutator transaction binding the contract method 0x731133e9.
//
// Solidity: function mint(address to, uint256 id, uint256 amount, bytes data) returns()
func (_MockTNT1155 *MockTNT1155Session) Mint(to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _MockTNT1155.Contract.Mint(&_MockTNT1155.TransactOpts, to, id, amount, data)
}

// Mint is a paid mutator transaction binding the contract method 0x731133e9.
//
// Solidity: function mint(address to, uint256 id, uint256 amount, bytes data) returns()
func (_MockTNT1155 *MockTNT1155TransactorSession) Mint(to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _MockTNT1155.Contract.Mint(&_MockTNT1155.TransactOpts, to, id, amount, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_MockTNT1155 *MockTNT1155Transactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _MockTNT1155.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_MockTNT1155 *MockTNT1155Session) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _MockTNT1155.Contract.SafeBatchTransferFrom(&_MockTNT1155.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_MockTNT1155 *MockTNT1155TransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _MockTNT1155.Contract.SafeBatchTransferFrom(&_MockTNT1155.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_MockTNT1155 *MockTNT1155Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _MockTNT1155.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_MockTNT1155 *MockTNT1155Session) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _MockTNT1155.Contract.SafeTransferFrom(&_MockTNT1155.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_MockTNT1155 *MockTNT1155TransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _MockTNT1155.Contract.SafeTransferFrom(&_MockTNT1155.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MockTNT1155 *MockTNT1155Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _MockTNT1155.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MockTNT1155 *MockTNT1155Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _MockTNT1155.Contract.SetApprovalForAll(&_MockTNT1155.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MockTNT1155 *MockTNT1155TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _MockTNT1155.Contract.SetApprovalForAll(&_MockTNT1155.TransactOpts, operator, approved)
}

// MockTNT1155ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the MockTNT1155 contract.
type MockTNT1155ApprovalForAllIterator struct {
	Event *MockTNT1155ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *MockTNT1155ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockTNT1155ApprovalForAll)
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
		it.Event = new(MockTNT1155ApprovalForAll)
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
func (it *MockTNT1155ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockTNT1155ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockTNT1155ApprovalForAll represents a ApprovalForAll event raised by the MockTNT1155 contract.
type MockTNT1155ApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_MockTNT1155 *MockTNT1155Filterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*MockTNT1155ApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _MockTNT1155.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &MockTNT1155ApprovalForAllIterator{contract: _MockTNT1155.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_MockTNT1155 *MockTNT1155Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *MockTNT1155ApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _MockTNT1155.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockTNT1155ApprovalForAll)
				if err := _MockTNT1155.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_MockTNT1155 *MockTNT1155Filterer) ParseApprovalForAll(log types.Log) (*MockTNT1155ApprovalForAll, error) {
	event := new(MockTNT1155ApprovalForAll)
	if err := _MockTNT1155.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockTNT1155TransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the MockTNT1155 contract.
type MockTNT1155TransferBatchIterator struct {
	Event *MockTNT1155TransferBatch // Event containing the contract specifics and raw log

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
func (it *MockTNT1155TransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockTNT1155TransferBatch)
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
		it.Event = new(MockTNT1155TransferBatch)
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
func (it *MockTNT1155TransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockTNT1155TransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockTNT1155TransferBatch represents a TransferBatch event raised by the MockTNT1155 contract.
type MockTNT1155TransferBatch struct {
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
func (_MockTNT1155 *MockTNT1155Filterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*MockTNT1155TransferBatchIterator, error) {

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

	logs, sub, err := _MockTNT1155.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MockTNT1155TransferBatchIterator{contract: _MockTNT1155.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_MockTNT1155 *MockTNT1155Filterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *MockTNT1155TransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _MockTNT1155.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockTNT1155TransferBatch)
				if err := _MockTNT1155.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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
func (_MockTNT1155 *MockTNT1155Filterer) ParseTransferBatch(log types.Log) (*MockTNT1155TransferBatch, error) {
	event := new(MockTNT1155TransferBatch)
	if err := _MockTNT1155.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockTNT1155TransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the MockTNT1155 contract.
type MockTNT1155TransferSingleIterator struct {
	Event *MockTNT1155TransferSingle // Event containing the contract specifics and raw log

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
func (it *MockTNT1155TransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockTNT1155TransferSingle)
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
		it.Event = new(MockTNT1155TransferSingle)
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
func (it *MockTNT1155TransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockTNT1155TransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockTNT1155TransferSingle represents a TransferSingle event raised by the MockTNT1155 contract.
type MockTNT1155TransferSingle struct {
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
func (_MockTNT1155 *MockTNT1155Filterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*MockTNT1155TransferSingleIterator, error) {

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

	logs, sub, err := _MockTNT1155.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MockTNT1155TransferSingleIterator{contract: _MockTNT1155.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_MockTNT1155 *MockTNT1155Filterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *MockTNT1155TransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _MockTNT1155.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockTNT1155TransferSingle)
				if err := _MockTNT1155.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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
func (_MockTNT1155 *MockTNT1155Filterer) ParseTransferSingle(log types.Log) (*MockTNT1155TransferSingle, error) {
	event := new(MockTNT1155TransferSingle)
	if err := _MockTNT1155.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockTNT1155URIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the MockTNT1155 contract.
type MockTNT1155URIIterator struct {
	Event *MockTNT1155URI // Event containing the contract specifics and raw log

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
func (it *MockTNT1155URIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockTNT1155URI)
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
		it.Event = new(MockTNT1155URI)
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
func (it *MockTNT1155URIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockTNT1155URIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockTNT1155URI represents a URI event raised by the MockTNT1155 contract.
type MockTNT1155URI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_MockTNT1155 *MockTNT1155Filterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*MockTNT1155URIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MockTNT1155.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &MockTNT1155URIIterator{contract: _MockTNT1155.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_MockTNT1155 *MockTNT1155Filterer) WatchURI(opts *bind.WatchOpts, sink chan<- *MockTNT1155URI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MockTNT1155.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockTNT1155URI)
				if err := _MockTNT1155.contract.UnpackLog(event, "URI", log); err != nil {
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
func (_MockTNT1155 *MockTNT1155Filterer) ParseURI(log types.Log) (*MockTNT1155URI, error) {
	event := new(MockTNT1155URI)
	if err := _MockTNT1155.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
