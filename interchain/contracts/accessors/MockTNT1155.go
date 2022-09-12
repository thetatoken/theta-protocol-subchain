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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040518060600160405280602d815260200162001c81602d9139620000378162000077565b5062000071732e833968e5bb786ae419c4d13189fb081cc43bab60016064604051806020016040528060008152506200009060201b60201c565b620006da565b80516200008c906002906020840190620003f4565b5050565b6001600160a01b038416620000f65760405162461bcd60e51b815260206004820152602160248201527f455243313135353a206d696e7420746f20746865207a65726f206164647265736044820152607360f81b60648201526084015b60405180910390fd5b3360006200010485620001b2565b905060006200011385620001b2565b90506000868152602081815260408083206001600160a01b038b168452909152812080548792906200014790849062000579565b909155505060408051878152602081018790526001600160a01b03808a1692600092918716917fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62910160405180910390a4620001a98360008989898962000208565b50505050505050565b60408051600180825281830190925260609160009190602080830190803683370190505090508281600081518110620001ef57620001ef62000618565b602090810291909101015292915050565b505050505050565b62000227846001600160a01b0316620003e560201b6200057f1760201c565b15620002005760405163f23a6e6160e01b81526001600160a01b0385169063f23a6e61906200026390899089908890889088906004016200051d565b602060405180830381600087803b1580156200027e57600080fd5b505af1925050508015620002b1575060408051601f3d908101601f19168201909252620002ae918101906200049a565b60015b6200037257620002c06200062e565b806308c379a01415620003015750620002d86200064b565b80620002e5575062000303565b8060405162461bcd60e51b8152600401620000ed919062000564565b505b60405162461bcd60e51b815260206004820152603460248201527f455243313135353a207472616e7366657220746f206e6f6e204552433131353560448201527f526563656976657220696d706c656d656e7465720000000000000000000000006064820152608401620000ed565b6001600160e01b0319811663f23a6e6160e01b14620001a95760405162461bcd60e51b815260206004820152602860248201527f455243313135353a204552433131353552656365697665722072656a656374656044820152676420746f6b656e7360c01b6064820152608401620000ed565b6001600160a01b03163b151590565b8280546200040290620005a0565b90600052602060002090601f01602090048101928262000426576000855562000471565b82601f106200044157805160ff191683800117855562000471565b8280016001018555821562000471579182015b828111156200047157825182559160200191906001019062000454565b506200047f92915062000483565b5090565b5b808211156200047f576000815560010162000484565b600060208284031215620004ad57600080fd5b81516001600160e01b031981168114620004c657600080fd5b9392505050565b6000815180845260005b81811015620004f557602081850181015186830182015201620004d7565b8181111562000508576000602083870101525b50601f01601f19169290920160200192915050565b6001600160a01b03868116825285166020820152604081018490526060810183905260a0608082018190526000906200055990830184620004cd565b979650505050505050565b602081526000620004c66020830184620004cd565b600082198211156200059b57634e487b7160e01b600052601160045260246000fd5b500190565b600181811c90821680620005b557607f821691505b60208210811415620005d757634e487b7160e01b600052602260045260246000fd5b50919050565b601f8201601f191681016001600160401b03811182821017156200061157634e487b7160e01b600052604160045260246000fd5b6040525050565b634e487b7160e01b600052603260045260246000fd5b600060033d1115620006485760046000803e5060005160e01c5b90565b600060443d10156200065a5790565b6040516003193d81016004833e81513d6001600160401b0380831160248401831017156200068a57505050505090565b8285019150815181811115620006a35750505050505090565b843d8701016020828501011115620006be5750505050505090565b620006cf60208286010187620005dd565b509095945050505050565b61159780620006ea6000396000f3fe608060405234801561001057600080fd5b50600436106100825760003560e01c8062fdd58e1461008757806301ffc9a7146100ad5780630e89341c146100d05780632eb2c2d6146100f05780634e1273f414610105578063731133e914610125578063a22cb46514610138578063e985e9c51461014b578063f242432a14610187575b600080fd5b61009a610095366004610f9a565b61019a565b6040519081526020015b60405180910390f35b6100c06100bb3660046110f4565b610231565b60405190151581526020016100a4565b6100e36100de366004611135565b610283565b6040516100a491906112ba565b6101036100fe366004610e51565b610317565b005b610118610113366004611024565b6103ae565b6040516100a49190611279565b610103610133366004610fc4565b6104d7565b610103610146366004610f5e565b6104e9565b6100c0610159366004610e1e565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205460ff1690565b610103610195366004610efa565b6104f8565b60006001600160a01b03831661020b5760405162461bcd60e51b815260206004820152602b60248201527f455243313135353a2062616c616e636520717565727920666f7220746865207a60448201526a65726f206164647265737360a81b60648201526084015b60405180910390fd5b506000908152602081815260408083206001600160a01b03949094168352929052205490565b60006001600160e01b03198216636cdb3d1360e11b148061026257506001600160e01b031982166303a24d0760e21b145b8061027d57506301ffc9a760e01b6001600160e01b03198316145b92915050565b606060028054610292906113df565b80601f01602080910402602001604051908101604052809291908181526020018280546102be906113df565b801561030b5780601f106102e05761010080835404028352916020019161030b565b820191906000526020600020905b8154815290600101906020018083116102ee57829003601f168201915b50505050509050919050565b6001600160a01b03851633148061033357506103338533610159565b61039a5760405162461bcd60e51b815260206004820152603260248201527f455243313135353a207472616e736665722063616c6c6572206973206e6f74206044820152711bdddb995c881b9bdc88185c1c1c9bdd995960721b6064820152608401610202565b6103a7858585858561058e565b5050505050565b606081518351146104135760405162461bcd60e51b815260206004820152602960248201527f455243313135353a206163636f756e747320616e6420696473206c656e677468604482015268040dad2e6dac2e8c6d60bb1b6064820152608401610202565b600083516001600160401b0381111561042e5761042e61148d565b604051908082528060200260200182016040528015610457578160200160208202803683370190505b50905060005b84518110156104cf576104a285828151811061047b5761047b611477565b602002602001015185838151811061049557610495611477565b602002602001015161019a565b8282815181106104b4576104b4611477565b60209081029190910101526104c881611446565b905061045d565b509392505050565b6104e38484848461076b565b50505050565b6104f433838361087f565b5050565b6001600160a01b03851633148061051457506105148533610159565b6105725760405162461bcd60e51b815260206004820152602960248201527f455243313135353a2063616c6c6572206973206e6f74206f776e6572206e6f7260448201526808185c1c1c9bdd995960ba1b6064820152608401610202565b6103a78585858585610960565b6001600160a01b03163b151590565b81518351146105f05760405162461bcd60e51b815260206004820152602860248201527f455243313135353a2069647320616e6420616d6f756e7473206c656e677468206044820152670dad2e6dac2e8c6d60c31b6064820152608401610202565b6001600160a01b0384166106165760405162461bcd60e51b815260040161020290611315565b3360005b84518110156106fd57600085828151811061063757610637611477565b60200260200101519050600085838151811061065557610655611477565b602090810291909101810151600084815280835260408082206001600160a01b038e1683529093529190912054909150818110156106a55760405162461bcd60e51b81526004016102029061135a565b6000838152602081815260408083206001600160a01b038e8116855292528083208585039055908b168252812080548492906106e29084906113c7565b92505081905550505050806106f690611446565b905061061a565b50846001600160a01b0316866001600160a01b0316826001600160a01b03167f4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb878760405161074d92919061128c565b60405180910390a4610763818787878787610a8a565b505050505050565b6001600160a01b0384166107cb5760405162461bcd60e51b815260206004820152602160248201527f455243313135353a206d696e7420746f20746865207a65726f206164647265736044820152607360f81b6064820152608401610202565b3360006107d785610bfc565b905060006107e485610bfc565b90506000868152602081815260408083206001600160a01b038b168452909152812080548792906108169084906113c7565b909155505060408051878152602081018790526001600160a01b03808a1692600092918716917fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62910160405180910390a461087683600089898989610c47565b50505050505050565b816001600160a01b0316836001600160a01b031614156108f35760405162461bcd60e51b815260206004820152602960248201527f455243313135353a2073657474696e6720617070726f76616c20737461747573604482015268103337b91039b2b63360b91b6064820152608401610202565b6001600160a01b03838116600081815260016020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b6001600160a01b0384166109865760405162461bcd60e51b815260040161020290611315565b33600061099285610bfc565b9050600061099f85610bfc565b90506000868152602081815260408083206001600160a01b038c168452909152902054858110156109e25760405162461bcd60e51b81526004016102029061135a565b6000878152602081815260408083206001600160a01b038d8116855292528083208985039055908a16825281208054889290610a1f9084906113c7565b909155505060408051888152602081018890526001600160a01b03808b16928c821692918816917fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62910160405180910390a4610a7f848a8a8a8a8a610c47565b505050505050505050565b610a9c846001600160a01b031661057f565b156107635760405163bc197c8160e01b81526001600160a01b0385169063bc197c8190610ad590899089908890889088906004016111d6565b602060405180830381600087803b158015610aef57600080fd5b505af1925050508015610b1f575060408051601f3d908101601f19168201909252610b1c91810190611118565b60015b610bcc57610b2b6114a3565b806308c379a01415610b655750610b406114bf565b80610b4b5750610b67565b8060405162461bcd60e51b815260040161020291906112ba565b505b60405162461bcd60e51b815260206004820152603460248201527f455243313135353a207472616e7366657220746f206e6f6e20455243313135356044820152732932b1b2b4bb32b91034b6b83632b6b2b73a32b960611b6064820152608401610202565b6001600160e01b0319811663bc197c8160e01b146108765760405162461bcd60e51b8152600401610202906112cd565b60408051600180825281830190925260609160009190602080830190803683370190505090508281600081518110610c3657610c36611477565b602090810291909101015292915050565b610c59846001600160a01b031661057f565b156107635760405163f23a6e6160e01b81526001600160a01b0385169063f23a6e6190610c929089908990889088908890600401611234565b602060405180830381600087803b158015610cac57600080fd5b505af1925050508015610cdc575060408051601f3d908101601f19168201909252610cd991810190611118565b60015b610ce857610b2b6114a3565b6001600160e01b0319811663f23a6e6160e01b146108765760405162461bcd60e51b8152600401610202906112cd565b80356001600160a01b0381168114610d2f57600080fd5b919050565b600082601f830112610d4557600080fd5b81356020610d52826113a4565b604051610d5f828261141a565b8381528281019150858301600585901b87018401881015610d7f57600080fd5b60005b85811015610d9e57813584529284019290840190600101610d82565b5090979650505050505050565b600082601f830112610dbc57600080fd5b81356001600160401b03811115610dd557610dd561148d565b604051610dec601f8301601f19166020018261141a565b818152846020838601011115610e0157600080fd5b816020850160208301376000918101602001919091529392505050565b60008060408385031215610e3157600080fd5b610e3a83610d18565b9150610e4860208401610d18565b90509250929050565b600080600080600060a08688031215610e6957600080fd5b610e7286610d18565b9450610e8060208701610d18565b935060408601356001600160401b0380821115610e9c57600080fd5b610ea889838a01610d34565b94506060880135915080821115610ebe57600080fd5b610eca89838a01610d34565b93506080880135915080821115610ee057600080fd5b50610eed88828901610dab565b9150509295509295909350565b600080600080600060a08688031215610f1257600080fd5b610f1b86610d18565b9450610f2960208701610d18565b9350604086013592506060860135915060808601356001600160401b03811115610f5257600080fd5b610eed88828901610dab565b60008060408385031215610f7157600080fd5b610f7a83610d18565b915060208301358015158114610f8f57600080fd5b809150509250929050565b60008060408385031215610fad57600080fd5b610fb683610d18565b946020939093013593505050565b60008060008060808587031215610fda57600080fd5b610fe385610d18565b9350602085013592506040850135915060608501356001600160401b0381111561100c57600080fd5b61101887828801610dab565b91505092959194509250565b6000806040838503121561103757600080fd5b82356001600160401b038082111561104e57600080fd5b818501915085601f83011261106257600080fd5b8135602061106f826113a4565b60405161107c828261141a565b8381528281019150858301600585901b870184018b101561109c57600080fd5b600096505b848710156110c6576110b281610d18565b8352600196909601959183019183016110a1565b50965050860135925050808211156110dd57600080fd5b506110ea85828601610d34565b9150509250929050565b60006020828403121561110657600080fd5b813561111181611548565b9392505050565b60006020828403121561112a57600080fd5b815161111181611548565b60006020828403121561114757600080fd5b5035919050565b600081518084526020808501945080840160005b8381101561117e57815187529582019590820190600101611162565b509495945050505050565b6000815180845260005b818110156111af57602081850181015186830182015201611193565b818111156111c1576000602083870101525b50601f01601f19169290920160200192915050565b6001600160a01b0386811682528516602082015260a0604082018190526000906112029083018661114e565b8281036060840152611214818661114e565b905082810360808401526112288185611189565b98975050505050505050565b6001600160a01b03868116825285166020820152604081018490526060810183905260a06080820181905260009061126e90830184611189565b979650505050505050565b602081526000611111602083018461114e565b60408152600061129f604083018561114e565b82810360208401526112b1818561114e565b95945050505050565b6020815260006111116020830184611189565b60208082526028908201527f455243313135353a204552433131353552656365697665722072656a656374656040820152676420746f6b656e7360c01b606082015260800190565b60208082526025908201527f455243313135353a207472616e7366657220746f20746865207a65726f206164604082015264647265737360d81b606082015260800190565b6020808252602a908201527f455243313135353a20696e73756666696369656e742062616c616e636520666f60408201526939103a3930b739b332b960b11b606082015260800190565b60006001600160401b038211156113bd576113bd61148d565b5060051b60200190565b600082198211156113da576113da611461565b500190565b600181811c908216806113f357607f821691505b6020821081141561141457634e487b7160e01b600052602260045260246000fd5b50919050565b601f8201601f191681016001600160401b038111828210171561143f5761143f61148d565b6040525050565b600060001982141561145a5761145a611461565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b600060033d11156114bc5760046000803e5060005160e01c5b90565b600060443d10156114cd5790565b6040516003193d81016004833e81513d6001600160401b0381602484011181841117156114fc57505050505090565b82850191508151818111156115145750505050505090565b843d870101602082850101111561152e5750505050505090565b61153d6020828601018761141a565b509095945050505050565b6001600160e01b03198116811461155e57600080fd5b5056fea2646970667358221220b9c5361461526612dbc02eb64330af10d1f31973d3faaccf8fa16801f9229dff64736f6c6343000807003368747470733a2f2f7468657461746f6b656e2e6578616d706c652f6170692f6974656d2f7b69647d2e6a736f6e",
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
