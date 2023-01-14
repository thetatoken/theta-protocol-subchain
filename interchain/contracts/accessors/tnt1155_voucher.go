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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"demon_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"UpdateOwner\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"updateOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voucherReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintedAmount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenUri\",\"type\":\"string\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voucherOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"burnedAmount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"denom\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContractOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001e2638038062001e2683398101604081905262000034916200014e565b604080516020810190915260008152829062000050816200008f565b50600380546001600160a01b0319166001600160a01b0392909216919091179055805162000086906004906020840190620000a8565b505050620002a1565b8051620000a4906002906020840190620000a8565b5050565b828054620000b6906200024e565b90600052602060002090601f016020900481019282620000da576000855562000125565b82601f10620000f557805160ff191683800117855562000125565b8280016001018555821562000125579182015b828111156200012557825182559160200191906001019062000108565b506200013392915062000137565b5090565b5b8082111562000133576000815560010162000138565b600080604083850312156200016257600080fd5b82516001600160a01b03811681146200017a57600080fd5b602084810151919350906001600160401b03808211156200019a57600080fd5b818601915086601f830112620001af57600080fd5b815181811115620001c457620001c46200028b565b604051601f8201601f19908116603f01168101908382118183101715620001ef57620001ef6200028b565b8160405282815289868487010111156200020857600080fd5b600093505b828410156200022c57848401860151818501870152928501926200020d565b828411156200023e5760008684830101525b8096505050505050509250929050565b600181811c908216806200026357607f821691505b602082108114156200028557634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052604160045260246000fd5b611b7580620002b16000396000f3fe608060405234801561001057600080fd5b50600436106100c95760003560e01c80638da5cb5b1161007c5780638da5cb5b146101a4578063a22cb465146101b7578063bb7fde71146101ca578063c370b042146101dd578063e985e9c5146101e5578063f242432a146101f8578063f5298aca1461020b57600080fd5b8062fdd58e146100ce57806301ffc9a7146100f45780630e89341c146101175780632eb2c2d614610137578063442890d51461014c5780634e1273f414610171578063880cdc3114610191575b600080fd5b6100e16100dc3660046114e1565b61021e565b6040519081526020015b60405180910390f35b610107610102366004611682565b6102b5565b60405190151581526020016100eb565b61012a6101253660046116bc565b610307565b6040516100eb9190611841565b61014a610145366004611398565b6103a9565b005b6003546001600160a01b03165b6040516001600160a01b0390911681526020016100eb565b61018461017f3660046115b2565b610440565b6040516100eb9190611800565b61014a61019f36600461134a565b610569565b600354610159906001600160a01b031681565b61014a6101c53660046114a5565b6105fc565b61014a6101d836600461153e565b61060b565b61012a61066f565b6101076101f3366004611365565b610701565b61014a610206366004611441565b61072f565b61014a61021936600461150b565b6107b6565b60006001600160a01b03831661028f5760405162461bcd60e51b815260206004820152602b60248201527f455243313135353a2062616c616e636520717565727920666f7220746865207a60448201526a65726f206164647265737360a81b60648201526084015b60405180910390fd5b506000908152602081815260408083206001600160a01b03949094168352929052205490565b60006001600160e01b03198216636cdb3d1360e11b14806102e657506001600160e01b031982166303a24d0760e21b145b8061030157506301ffc9a760e01b6001600160e01b03198316145b92915050565b60008181526005602052604090208054606091906103249061199d565b80601f01602080910402602001604051908101604052809291908181526020018280546103509061199d565b801561039d5780601f106103725761010080835404028352916020019161039d565b820191906000526020600020905b81548152906001019060200180831161038057829003601f168201915b50505050509050919050565b6001600160a01b0385163314806103c557506103c58533610701565b61042c5760405162461bcd60e51b815260206004820152603260248201527f455243313135353a207472616e736665722063616c6c6572206973206e6f74206044820152711bdddb995c881b9bdc88185c1c1c9bdd995960721b6064820152608401610286565b61043985858585856108d8565b5050505050565b606081518351146104a55760405162461bcd60e51b815260206004820152602960248201527f455243313135353a206163636f756e747320616e6420696473206c656e677468604482015268040dad2e6dac2e8c6d60bb1b6064820152608401610286565b600083516001600160401b038111156104c0576104c0611a4b565b6040519080825280602002602001820160405280156104e9578160200160208202803683370190505b50905060005b84518110156105615761053485828151811061050d5761050d611a35565b602002602001015185838151811061052757610527611a35565b602002602001015161021e565b82828151811061054657610546611a35565b602090810291909101015261055a81611a04565b90506104ef565b509392505050565b6003546001600160a01b031633146105935760405162461bcd60e51b81526004016102869061192b565b600354604080516001600160a01b03928316815291831660208301527fe2c7d1c4da37855e682bde14f17826d185497973b73fba7554daa6da467058d9910160405180910390a1600380546001600160a01b0319166001600160a01b0392909216919091179055565b610607338383610ab5565b5050565b6003546001600160a01b031633146106355760405162461bcd60e51b81526004016102869061192b565b61065084848460405180602001604052806000815250610b96565b600083815260056020908152604090912082516104399284019061119a565b60606004805461067e9061199d565b80601f01602080910402602001604051908101604052809291908181526020018280546106aa9061199d565b80156106f75780601f106106cc576101008083540402835291602001916106f7565b820191906000526020600020905b8154815290600101906020018083116106da57829003601f168201915b5050505050905090565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205460ff1690565b6001600160a01b03851633148061074b575061074b8533610701565b6107a95760405162461bcd60e51b815260206004820152602960248201527f455243313135353a2063616c6c6572206973206e6f74206f776e6572206e6f7260448201526808185c1c1c9bdd995960ba1b6064820152608401610286565b6104398585858585610c98565b6003546001600160a01b031633146107e05760405162461bcd60e51b81526004016102869061192b565b806107eb848461021e565b10156108545760405162461bcd60e51b815260206004820152603260248201527f566f7563686572206f776e657220646f6573206e6f74206861766520656e6f7560448201527133b4103130b630b731b2903a3790313ab93760711b6064820152608401610286565b6003546001600160a01b031661086a8482610701565b6108c75760405162461bcd60e51b815260206004820152602860248201527f566f7563686572206f776e657220646964206e6f7420617070726f766520746f60448201526735b2b710313ab93760c11b6064820152608401610286565b6108d2848484610db0565b50505050565b815183511461093a5760405162461bcd60e51b815260206004820152602860248201527f455243313135353a2069647320616e6420616d6f756e7473206c656e677468206044820152670dad2e6dac2e8c6d60c31b6064820152608401610286565b6001600160a01b0384166109605760405162461bcd60e51b81526004016102869061189c565b3360005b8451811015610a4757600085828151811061098157610981611a35565b60200260200101519050600085838151811061099f5761099f611a35565b602090810291909101810151600084815280835260408082206001600160a01b038e1683529093529190912054909150818110156109ef5760405162461bcd60e51b8152600401610286906118e1565b6000838152602081815260408083206001600160a01b038e8116855292528083208585039055908b16825281208054849290610a2c908490611985565b9250508190555050505080610a4090611a04565b9050610964565b50846001600160a01b0316866001600160a01b0316826001600160a01b03167f4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb8787604051610a97929190611813565b60405180910390a4610aad818787878787610f1a565b505050505050565b816001600160a01b0316836001600160a01b03161415610b295760405162461bcd60e51b815260206004820152602960248201527f455243313135353a2073657474696e6720617070726f76616c20737461747573604482015268103337b91039b2b63360b91b6064820152608401610286565b6001600160a01b03838116600081815260016020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b6001600160a01b038416610bf65760405162461bcd60e51b815260206004820152602160248201527f455243313135353a206d696e7420746f20746865207a65726f206164647265736044820152607360f81b6064820152608401610286565b336000610c0285611085565b90506000610c0f85611085565b90506000868152602081815260408083206001600160a01b038b16845290915281208054879290610c41908490611985565b909155505060408051878152602081018790526001600160a01b03808a169260009291871691600080516020611b20833981519152910160405180910390a4610c8f836000898989896110d0565b50505050505050565b6001600160a01b038416610cbe5760405162461bcd60e51b81526004016102869061189c565b336000610cca85611085565b90506000610cd785611085565b90506000868152602081815260408083206001600160a01b038c16845290915290205485811015610d1a5760405162461bcd60e51b8152600401610286906118e1565b6000878152602081815260408083206001600160a01b038d8116855292528083208985039055908a16825281208054889290610d57908490611985565b909155505060408051888152602081018890526001600160a01b03808b16928c82169291881691600080516020611b20833981519152910160405180910390a4610da5848a8a8a8a8a6110d0565b505050505050505050565b6001600160a01b038316610e125760405162461bcd60e51b815260206004820152602360248201527f455243313135353a206275726e2066726f6d20746865207a65726f206164647260448201526265737360e81b6064820152608401610286565b336000610e1e84611085565b90506000610e2b84611085565b60408051602080820183526000918290528882528181528282206001600160a01b038b1683529052205490915084811015610eb45760405162461bcd60e51b8152602060048201526024808201527f455243313135353a206275726e20616d6f756e7420657863656564732062616c604482015263616e636560e01b6064820152608401610286565b6000868152602081815260408083206001600160a01b038b81168086529184528285208a8703905582518b81529384018a9052909290881691600080516020611b20833981519152910160405180910390a4604080516020810190915260009052610c8f565b6001600160a01b0384163b15610aad5760405163bc197c8160e01b81526001600160a01b0385169063bc197c8190610f5e908990899088908890889060040161175d565b602060405180830381600087803b158015610f7857600080fd5b505af1925050508015610fa8575060408051601f3d908101601f19168201909252610fa59181019061169f565b60015b61105557610fb4611a61565b806308c379a01415610fee5750610fc9611a7d565b80610fd45750610ff0565b8060405162461bcd60e51b81526004016102869190611841565b505b60405162461bcd60e51b815260206004820152603460248201527f455243313135353a207472616e7366657220746f206e6f6e20455243313135356044820152732932b1b2b4bb32b91034b6b83632b6b2b73a32b960611b6064820152608401610286565b6001600160e01b0319811663bc197c8160e01b14610c8f5760405162461bcd60e51b815260040161028690611854565b604080516001808252818301909252606091600091906020808301908036833701905050905082816000815181106110bf576110bf611a35565b602090810291909101015292915050565b6001600160a01b0384163b15610aad5760405163f23a6e6160e01b81526001600160a01b0385169063f23a6e619061111490899089908890889088906004016117bb565b602060405180830381600087803b15801561112e57600080fd5b505af192505050801561115e575060408051601f3d908101601f1916820190925261115b9181019061169f565b60015b61116a57610fb4611a61565b6001600160e01b0319811663f23a6e6160e01b14610c8f5760405162461bcd60e51b815260040161028690611854565b8280546111a69061199d565b90600052602060002090601f0160209004810192826111c8576000855561120e565b82601f106111e157805160ff191683800117855561120e565b8280016001018555821561120e579182015b8281111561120e5782518255916020019190600101906111f3565b5061121a92915061121e565b5090565b5b8082111561121a576000815560010161121f565b60006001600160401b0383111561124c5761124c611a4b565b604051611263601f8501601f1916602001826119d8565b80915083815284848401111561127857600080fd5b83836020830137600060208583010152509392505050565b80356001600160a01b03811681146112a757600080fd5b919050565b600082601f8301126112bd57600080fd5b813560206112ca82611962565b6040516112d782826119d8565b8381528281019150858301600585901b870184018810156112f757600080fd5b60005b85811015611316578135845292840192908401906001016112fa565b5090979650505050505050565b600082601f83011261133457600080fd5b61134383833560208501611233565b9392505050565b60006020828403121561135c57600080fd5b61134382611290565b6000806040838503121561137857600080fd5b61138183611290565b915061138f60208401611290565b90509250929050565b600080600080600060a086880312156113b057600080fd5b6113b986611290565b94506113c760208701611290565b935060408601356001600160401b03808211156113e357600080fd5b6113ef89838a016112ac565b9450606088013591508082111561140557600080fd5b61141189838a016112ac565b9350608088013591508082111561142757600080fd5b5061143488828901611323565b9150509295509295909350565b600080600080600060a0868803121561145957600080fd5b61146286611290565b945061147060208701611290565b9350604086013592506060860135915060808601356001600160401b0381111561149957600080fd5b61143488828901611323565b600080604083850312156114b857600080fd5b6114c183611290565b9150602083013580151581146114d657600080fd5b809150509250929050565b600080604083850312156114f457600080fd5b6114fd83611290565b946020939093013593505050565b60008060006060848603121561152057600080fd5b61152984611290565b95602085013595506040909401359392505050565b6000806000806080858703121561155457600080fd5b61155d85611290565b9350602085013592506040850135915060608501356001600160401b0381111561158657600080fd5b8501601f8101871361159757600080fd5b6115a687823560208401611233565b91505092959194509250565b600080604083850312156115c557600080fd5b82356001600160401b03808211156115dc57600080fd5b818501915085601f8301126115f057600080fd5b813560206115fd82611962565b60405161160a82826119d8565b8381528281019150858301600585901b870184018b101561162a57600080fd5b600096505b848710156116545761164081611290565b83526001969096019591830191830161162f565b509650508601359250508082111561166b57600080fd5b50611678858286016112ac565b9150509250929050565b60006020828403121561169457600080fd5b813561134381611b06565b6000602082840312156116b157600080fd5b815161134381611b06565b6000602082840312156116ce57600080fd5b5035919050565b600081518084526020808501945080840160005b83811015611705578151875295820195908201906001016116e9565b509495945050505050565b6000815180845260005b818110156117365760208185018101518683018201520161171a565b81811115611748576000602083870101525b50601f01601f19169290920160200192915050565b6001600160a01b0386811682528516602082015260a060408201819052600090611789908301866116d5565b828103606084015261179b81866116d5565b905082810360808401526117af8185611710565b98975050505050505050565b6001600160a01b03868116825285166020820152604081018490526060810183905260a0608082018190526000906117f590830184611710565b979650505050505050565b60208152600061134360208301846116d5565b60408152600061182660408301856116d5565b828103602084015261183881856116d5565b95945050505050565b6020815260006113436020830184611710565b60208082526028908201527f455243313135353a204552433131353552656365697665722072656a656374656040820152676420746f6b656e7360c01b606082015260800190565b60208082526025908201527f455243313135353a207472616e7366657220746f20746865207a65726f206164604082015264647265737360d81b606082015260800190565b6020808252602a908201527f455243313135353a20696e73756666696369656e742062616c616e636520666f60408201526939103a3930b739b332b960b11b606082015260800190565b6020808252601c908201527f6f6e6c79206f776e65722063616e206d616b65207468652063616c6c00000000604082015260600190565b60006001600160401b0382111561197b5761197b611a4b565b5060051b60200190565b6000821982111561199857611998611a1f565b500190565b600181811c908216806119b157607f821691505b602082108114156119d257634e487b7160e01b600052602260045260246000fd5b50919050565b601f8201601f191681016001600160401b03811182821017156119fd576119fd611a4b565b6040525050565b6000600019821415611a1857611a18611a1f565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b600060033d1115611a7a5760046000803e5060005160e01c5b90565b600060443d1015611a8b5790565b6040516003193d81016004833e81513d6001600160401b038160248401118184111715611aba57505050505090565b8285019150815181811115611ad25750505050505090565b843d8701016020828501011115611aec5750505050505090565b611afb602082860101876119d8565b509095945050505050565b6001600160e01b031981168114611b1c57600080fd5b5056fec3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62a2646970667358221220a32ec34d7a33806923fe667deb7d902165f58ac0c402426f98477440247861a764736f6c63430008070033",
}

// TNT1155VoucherContractABI is the input ABI used to generate the binding from.
// Deprecated: Use TNT1155VoucherContractMetaData.ABI instead.
var TNT1155VoucherContractABI = TNT1155VoucherContractMetaData.ABI

// TNT1155VoucherContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TNT1155VoucherContractMetaData.Bin instead.
var TNT1155VoucherContractBin = TNT1155VoucherContractMetaData.Bin

// DeployTNT1155VoucherContract deploys a new Ethereum contract, binding an instance of TNT1155VoucherContract to it.
func DeployTNT1155VoucherContract(auth *bind.TransactOpts, backend bind.ContractBackend, owner_ common.Address, demon_ string) (common.Address, *types.Transaction, *TNT1155VoucherContract, error) {
	parsed, err := TNT1155VoucherContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TNT1155VoucherContractBin), backend, owner_, demon_)
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

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenID) view returns(string)
func (_TNT1155VoucherContract *TNT1155VoucherContractCaller) Uri(opts *bind.CallOpts, tokenID *big.Int) (string, error) {
	var out []interface{}
	err := _TNT1155VoucherContract.contract.Call(opts, &out, "uri", tokenID)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenID) view returns(string)
func (_TNT1155VoucherContract *TNT1155VoucherContractSession) Uri(tokenID *big.Int) (string, error) {
	return _TNT1155VoucherContract.Contract.Uri(&_TNT1155VoucherContract.CallOpts, tokenID)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenID) view returns(string)
func (_TNT1155VoucherContract *TNT1155VoucherContractCallerSession) Uri(tokenID *big.Int) (string, error) {
	return _TNT1155VoucherContract.Contract.Uri(&_TNT1155VoucherContract.CallOpts, tokenID)
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

// Mint is a paid mutator transaction binding the contract method 0xbb7fde71.
//
// Solidity: function mint(address voucherReceiver, uint256 tokenID, uint256 mintedAmount, string tokenUri) returns()
func (_TNT1155VoucherContract *TNT1155VoucherContractTransactor) Mint(opts *bind.TransactOpts, voucherReceiver common.Address, tokenID *big.Int, mintedAmount *big.Int, tokenUri string) (*types.Transaction, error) {
	return _TNT1155VoucherContract.contract.Transact(opts, "mint", voucherReceiver, tokenID, mintedAmount, tokenUri)
}

// Mint is a paid mutator transaction binding the contract method 0xbb7fde71.
//
// Solidity: function mint(address voucherReceiver, uint256 tokenID, uint256 mintedAmount, string tokenUri) returns()
func (_TNT1155VoucherContract *TNT1155VoucherContractSession) Mint(voucherReceiver common.Address, tokenID *big.Int, mintedAmount *big.Int, tokenUri string) (*types.Transaction, error) {
	return _TNT1155VoucherContract.Contract.Mint(&_TNT1155VoucherContract.TransactOpts, voucherReceiver, tokenID, mintedAmount, tokenUri)
}

// Mint is a paid mutator transaction binding the contract method 0xbb7fde71.
//
// Solidity: function mint(address voucherReceiver, uint256 tokenID, uint256 mintedAmount, string tokenUri) returns()
func (_TNT1155VoucherContract *TNT1155VoucherContractTransactorSession) Mint(voucherReceiver common.Address, tokenID *big.Int, mintedAmount *big.Int, tokenUri string) (*types.Transaction, error) {
	return _TNT1155VoucherContract.Contract.Mint(&_TNT1155VoucherContract.TransactOpts, voucherReceiver, tokenID, mintedAmount, tokenUri)
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
