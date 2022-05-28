// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/thetatoken/thetasubchain/eth"
	"github.com/thetatoken/thetasubchain/eth/abi"
	"github.com/thetatoken/thetasubchain/eth/abi/bind"
	"github.com/thetatoken/thetasubchain/eth/core/types"
	"github.com/thetatoken/thetasubchain/eth/event"

	"github.com/thetatoken/theta/common"
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

// MainchainTFuelTokenBankMetaData contains all meta data concerning the MainchainTFuelTokenBank contract.
var MainchainTFuelTokenBankMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"registerContractAddress_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"targetChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"mainchainTokenSender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"subchainTokenReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"name\":\"TFeulTokenLocked\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allDenoms\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allVouchers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"denomToVoucherLookup\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voucherAddress\",\"type\":\"address\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voucherContractAddr\",\"type\":\"address\"}],\"name\":\"getDenom\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"name\":\"getVoucher\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"voucherAddressToDenomLookup\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"targetChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"subchainTokenReceiver\",\"type\":\"address\"}],\"name\":\"sendToChain\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_sig\",\"type\":\"bytes\"}],\"name\":\"unlock\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subchainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"hasSufficientLockedAmount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040526040518060600160405280602a8152602001620023a5602a91396007908051906020019062000035929190620000b1565b503480156200004357600080fd5b50604051620023cf380380620023cf833981810160405281019062000069919062000178565b80600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505062000262565b828054620000bf90620001de565b90600052602060002090601f016020900481019282620000e357600085556200012f565b82601f10620000fe57805160ff19168380011785556200012f565b828001600101855582156200012f579182015b828111156200012e57825182559160200191906001019062000111565b5b5090506200013e919062000142565b5090565b5b808211156200015d57600081600090555060010162000143565b5090565b600081519050620001728162000248565b92915050565b60006020828403121562000191576200019062000243565b5b6000620001a18482850162000161565b91505092915050565b6000620001b782620001be565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006002820490506001821680620001f757607f821691505b602082108114156200020e576200020d62000214565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600080fd5b6200025381620001aa565b81146200025f57600080fd5b50565b61213380620002726000396000f3fe60806040526004361061009c5760003560e01c80635d9c701d116100645780635d9c701d146101b257806360569b5e146101ef5780638d59151f1461022d578063a2cc698114610249578063ebda996214610286578063f6a3d24e146102c35761009c565b80631527b14d146100a1578063261a323e146100df57806327ca4df11461011c5780632ee015d714610159578063588b140814610175575b600080fd5b3480156100ad57600080fd5b506100c860048036038101906100c39190611488565b610300565b6040516100d692919061190a565b60405180910390f35b3480156100eb57600080fd5b5061010660048036038101906101019190611488565b610367565b6040516101139190611986565b60405180910390f35b34801561012857600080fd5b50610143600480360381019061013e91906114d1565b6103ac565b60405161015091906118ef565b60405180910390f35b610173600480360381019061016e91906114fe565b6103eb565b005b34801561018157600080fd5b5061019c600480360381019061019791906114d1565b6106df565b6040516101a991906119c5565b60405180910390f35b3480156101be57600080fd5b506101d960048036038101906101d4919061153e565b61078b565b6040516101e69190611986565b60405180910390f35b3480156101fb57600080fd5b50610216600480360381019061021191906113ad565b6107ea565b6040516102249291906119e7565b60405180910390f35b61024760048036038101906102429190611407565b6108a3565b005b34801561025557600080fd5b50610270600480360381019061026b9190611488565b610d20565b60405161027d91906118ef565b60405180910390f35b34801561029257600080fd5b506102ad60048036038101906102a891906113ad565b610df0565b6040516102ba91906119c5565b60405180910390f35b3480156102cf57600080fd5b506102ea60048036038101906102e591906113ad565b610f1d565b6040516102f79190611986565b60405180910390f35b6002818051602081018201805184825260208301602085012081835280955050505050506000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060000160149054906101000a900460ff16905082565b60008061037383610f76565b9050600281604051610385919061188d565b908152602001604051809103902060000160149054906101000a900460ff16915050919050565b600481815481106103bc57600080fd5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637d1daa96846040518263ffffffff1660e01b81526004016104489190611ab7565b60206040518083038186803b15801561046057600080fd5b505afa158015610474573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061049891906113da565b9050806104da576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104d190611a77565b60405180910390fd5b600034905060003390506104ed85610f88565b816008600087815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600061068261055046610fb5565b6040518060400160405280600181526020017f2f000000000000000000000000000000000000000000000000000000000000008152506040518060400160405280600181526020017f30000000000000000000000000000000000000000000000000000000000000008152506040518060400160405280600181526020017f2f00000000000000000000000000000000000000000000000000000000000000815250600780546105ff90611daf565b80601f016020809104026020016040519081016040528092919081815260200182805461062b90611daf565b80156106785780601f1061064d57610100808354040283529160200191610678565b820191906000526020600020905b81548152906001019060200180831161065b57829003601f168201915b5050505050611116565b90507fb0d72a5b25cc833270260cccd76510dddd5e3a0aa8e49018804c93ed3b256838868387866000808c815260200190815260200160002054866040516106cf96959493929190611ad2565b60405180910390a1505050505050565b600581815481106106ef57600080fd5b90600052602060002001600091509050805461070a90611daf565b80601f016020809104026020016040519081016040528092919081815260200182805461073690611daf565b80156107835780601f1061075857610100808354040283529160200191610783565b820191906000526020600020905b81548152906001019060200180831161076657829003601f168201915b505050505081565b6000816008600086815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054101590509392505050565b600360205280600052604060002060009150905080600001805461080d90611daf565b80601f016020809104026020016040519081016040528092919081815260200182805461083990611daf565b80156108865780601f1061085b57610100808354040283529160200191610886565b820191906000526020600020905b81548152906001019060200180831161086957829003601f168201915b5050505050908060010160009054906101000a900460ff16905082565b60008060008060608060008a8a8101906108bd9190611591565b809a50819b50829750839850849950859c50869d50505050505050506000600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166365ef160d89338a6040518463ffffffff1660e01b815260040161093a93929190611b3a565b60206040518083038186803b15801561095257600080fd5b505afa158015610966573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061098a91906113da565b9050806109cc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109c390611a17565b60405180910390fd5b60008c8c6040516020016109e1929190611874565b6040516020818303038152906040528051906020012090506000806000600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a7bb58038f8f6040518363ffffffff1660e01b8152600401610a5b9291906119a1565b60606040518083038186803b158015610a7357600080fd5b505afa158015610a87573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610aab919061166b565b8094508195508293505050506000600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663990366d433878588886040518663ffffffff1660e01b8152600401610b1c959493929190611933565b60206040518083038186803b158015610b3457600080fd5b505afa158015610b48573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b6c91906113da565b905080610bae576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ba590611a97565b60405180910390fd5b60003073ffffffffffffffffffffffffffffffffffffffff1663261a323e8a6040518263ffffffff1660e01b8152600401610be991906119c5565b60206040518083038186803b158015610c0157600080fd5b505afa158015610c15573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c3991906113da565b905080610c7b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c7290611a37565b60405180910390fd5b610c868e898e61078b565b610cc5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cbc90611a57565b60405180910390fd5b8773ffffffffffffffffffffffffffffffffffffffff166108fc349081150290604051600060405180830381858888f19350505050158015610d0b573d6000803e3d6000fd5b50505050505050505050505050505050505050565b600080610d2c83610f76565b90506000600282604051610d40919061188d565b90815260200160405180910390206040518060400160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160149054906101000a900460ff1615151515815250509050806020015115610de457806000015192505050610deb565b6000925050505b919050565b60606000600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020604051806040016040529081600082018054610e4e90611daf565b80601f0160208091040260200160405190810160405280929190818152602001828054610e7a90611daf565b8015610ec75780601f10610e9c57610100808354040283529160200191610ec7565b820191906000526020600020905b815481529060010190602001808311610eaa57829003601f168201915b505050505081526020016001820160009054906101000a900460ff1615151515815250509050806020015115610f04578060000151915050610f18565b604051806020016040528060008152509150505b919050565b6000600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900460ff169050919050565b6060610f818261114b565b9050919050565b60016000808381526020019081526020016000206000828254610fab9190611c0a565b9250508190555050565b60606000821415610ffd576040518060400160405280600181526020017f30000000000000000000000000000000000000000000000000000000000000008152509050611111565b600082905060005b6000821461102f57808061101890611e12565b915050600a826110289190611c97565b9150611005565b60008167ffffffffffffffff81111561104b5761104a611f48565b5b6040519080825280601f01601f19166020018201604052801561107d5781602001600182028036833780820191505090505b5090505b6000851461110a576001826110969190611cc8565b9150600a856110a59190611e5b565b60306110b19190611c0a565b60f81b8183815181106110c7576110c6611f19565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a856111039190611c97565b9450611081565b8093505050505b919050565b606085858585856040516020016111319594939291906118a4565b604051602081830303815290604052905095945050505050565b6060600082905060005b81518110156111d95761118482828151811061117457611173611f19565b5b602001015160f81c60f81b6111e3565b82828151811061119757611196611f19565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535080806111d190611e12565b915050611155565b5080915050919050565b6000604160f81b827effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916101580156112415750605a60f81b827effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191611155b156112605760208260f81c6112569190611c60565b60f81b9050611264565b8190505b919050565b600061127c61127784611b96565b611b71565b90508281526020810184848401111561129857611297611f86565b5b6112a3848285611d6d565b509392505050565b6000813590506112ba81612073565b92915050565b6000813590506112cf8161208a565b92915050565b6000815190506112e4816120a1565b92915050565b6000815190506112f9816120b8565b92915050565b60008083601f84011261131557611314611f7c565b5b8235905067ffffffffffffffff81111561133257611331611f77565b5b60208301915083600182028301111561134e5761134d611f81565b5b9250929050565b600082601f83011261136a57611369611f7c565b5b813561137a848260208601611269565b91505092915050565b600081359050611392816120cf565b92915050565b6000815190506113a7816120e6565b92915050565b6000602082840312156113c3576113c2611f90565b5b60006113d1848285016112ab565b91505092915050565b6000602082840312156113f0576113ef611f90565b5b60006113fe848285016112d5565b91505092915050565b6000806000806040858703121561142157611420611f90565b5b600085013567ffffffffffffffff81111561143f5761143e611f8b565b5b61144b878288016112ff565b9450945050602085013567ffffffffffffffff81111561146e5761146d611f8b565b5b61147a878288016112ff565b925092505092959194509250565b60006020828403121561149e5761149d611f90565b5b600082013567ffffffffffffffff8111156114bc576114bb611f8b565b5b6114c884828501611355565b91505092915050565b6000602082840312156114e7576114e6611f90565b5b60006114f584828501611383565b91505092915050565b6000806040838503121561151557611514611f90565b5b600061152385828601611383565b9250506020611534858286016112ab565b9150509250929050565b60008060006060848603121561155757611556611f90565b5b600061156586828701611383565b9350506020611576868287016112ab565b925050604061158786828701611383565b9150509250925092565b600080600080600080600060e0888a0312156115b0576115af611f90565b5b60006115be8a828b01611383565b97505060206115cf8a828b01611383565b965050604088013567ffffffffffffffff8111156115f0576115ef611f8b565b5b6115fc8a828b01611355565b955050606088013567ffffffffffffffff81111561161d5761161c611f8b565b5b6116298a828b01611355565b945050608061163a8a828b016112c0565b93505060a061164b8a828b01611383565b92505060c061165c8a828b01611383565b91505092959891949750929550565b60008060006060848603121561168457611683611f90565b5b600061169286828701611398565b93505060206116a3868287016112ea565b92505060406116b4868287016112ea565b9150509250925092565b6116c781611cfc565b82525050565b6116d681611d20565b82525050565b6116e581611d2c565b82525050565b60006116f78385611bd2565b9350611704838584611d6d565b61170d83611f95565b840190509392505050565b60006117248385611be3565b9350611731838584611d6d565b82840190509392505050565b600061174882611bc7565b6117528185611bee565b9350611762818560208601611d7c565b61176b81611f95565b840191505092915050565b600061178182611bc7565b61178b8185611bff565b935061179b818560208601611d7c565b80840191505092915050565b60006117b4601f83611bee565b91506117bf82611fa6565b602082019050919050565b60006117d7601483611bee565b91506117e282611fcf565b602082019050919050565b60006117fa601a83611bee565b915061180582611ff8565b602082019050919050565b600061181d601383611bee565b915061182882612021565b602082019050919050565b6000611840601f83611bee565b915061184b8261204a565b602082019050919050565b61185f81611d56565b82525050565b61186e81611d60565b82525050565b6000611881828486611718565b91508190509392505050565b60006118998284611776565b915081905092915050565b60006118b08288611776565b91506118bc8287611776565b91506118c88286611776565b91506118d48285611776565b91506118e08284611776565b91508190509695505050505050565b600060208201905061190460008301846116be565b92915050565b600060408201905061191f60008301856116be565b61192c60208301846116cd565b9392505050565b600060a08201905061194860008301886116be565b61195560208301876116dc565b6119626040830186611865565b61196f60608301856116dc565b61197c60808301846116dc565b9695505050505050565b600060208201905061199b60008301846116cd565b92915050565b600060208201905081810360008301526119bc8184866116eb565b90509392505050565b600060208201905081810360008301526119df818461173d565b905092915050565b60006040820190508181036000830152611a01818561173d565b9050611a1060208301846116cd565b9392505050565b60006020820190508181036000830152611a30816117a7565b9050919050565b60006020820190508181036000830152611a50816117ca565b9050919050565b60006020820190508181036000830152611a70816117ed565b9050919050565b60006020820190508181036000830152611a9081611810565b9050919050565b60006020820190508181036000830152611ab081611833565b9050919050565b6000602082019050611acc6000830184611856565b92915050565b600060c082019050611ae76000830189611856565b611af460208301886116be565b611b0160408301876116be565b611b0e6060830186611856565b611b1b6080830185611856565b81810360a0830152611b2d818461173d565b9050979650505050505050565b6000606082019050611b4f6000830186611856565b611b5c60208301856116be565b611b696040830184611856565b949350505050565b6000611b7b611b8c565b9050611b878282611de1565b919050565b6000604051905090565b600067ffffffffffffffff821115611bb157611bb0611f48565b5b611bba82611f95565b9050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b600081905092915050565b6000611c1582611d56565b9150611c2083611d56565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611c5557611c54611e8c565b5b828201905092915050565b6000611c6b82611d60565b9150611c7683611d60565b92508260ff03821115611c8c57611c8b611e8c565b5b828201905092915050565b6000611ca282611d56565b9150611cad83611d56565b925082611cbd57611cbc611ebb565b5b828204905092915050565b6000611cd382611d56565b9150611cde83611d56565b925082821015611cf157611cf0611e8c565b5b828203905092915050565b6000611d0782611d36565b9050919050565b6000611d1982611d36565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b83811015611d9a578082015181840152602081019050611d7f565b83811115611da9576000848401525b50505050565b60006002820490506001821680611dc757607f821691505b60208210811415611ddb57611dda611eea565b5b50919050565b611dea82611f95565b810181811067ffffffffffffffff82111715611e0957611e08611f48565b5b80604052505050565b6000611e1d82611d56565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415611e5057611e4f611e8c565b5b600182019050919050565b6000611e6682611d56565b9150611e7183611d56565b925082611e8157611e80611ebb565b5b828206905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f546869732073656e646572206973206e6f7420612076616c696461746f722100600082015250565b7f44656e6f6d20646f6573206e6f74206578697374000000000000000000000000600082015250565b7f496e73756666696369656e74206c6f636b656420616d6f756e74000000000000600082015250565b7f496e76616c696420537562636861696e49442100000000000000000000000000600082015250565b7f496e76616c6964207369676e61747572652c2077686f2061726520796f753f00600082015250565b61207c81611cfc565b811461208757600080fd5b50565b61209381611d0e565b811461209e57600080fd5b50565b6120aa81611d20565b81146120b557600080fd5b50565b6120c181611d2c565b81146120cc57600080fd5b50565b6120d881611d56565b81146120e357600080fd5b50565b6120ef81611d60565b81146120fa57600080fd5b5056fea26469706673582212206f0a7cfd73cf01c97ba951adf1b119a09b5cb41cdab24d68b3380d7f2934c73864736f6c63430008070033307830303030303030303030303030303030303030303030303030303030303030303030303030303030",
}

// MainchainTFuelTokenBankABI is the input ABI used to generate the binding from.
// Deprecated: Use MainchainTFuelTokenBankMetaData.ABI instead.
var MainchainTFuelTokenBankABI = MainchainTFuelTokenBankMetaData.ABI

// MainchainTFuelTokenBankBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MainchainTFuelTokenBankMetaData.Bin instead.
var MainchainTFuelTokenBankBin = MainchainTFuelTokenBankMetaData.Bin

// DeployMainchainTFuelTokenBank deploys a new Ethereum contract, binding an instance of MainchainTFuelTokenBank to it.
func DeployMainchainTFuelTokenBank(auth *bind.TransactOpts, backend bind.ContractBackend, registerContractAddress_ common.Address) (common.Address, *types.Transaction, *MainchainTFuelTokenBank, error) {
	parsed, err := MainchainTFuelTokenBankMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MainchainTFuelTokenBankBin), backend, registerContractAddress_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MainchainTFuelTokenBank{MainchainTFuelTokenBankCaller: MainchainTFuelTokenBankCaller{contract: contract}, MainchainTFuelTokenBankTransactor: MainchainTFuelTokenBankTransactor{contract: contract}, MainchainTFuelTokenBankFilterer: MainchainTFuelTokenBankFilterer{contract: contract}}, nil
}

// MainchainTFuelTokenBank is an auto generated Go binding around an Ethereum contract.
type MainchainTFuelTokenBank struct {
	MainchainTFuelTokenBankCaller     // Read-only binding to the contract
	MainchainTFuelTokenBankTransactor // Write-only binding to the contract
	MainchainTFuelTokenBankFilterer   // Log filterer for contract events
}

// MainchainTFuelTokenBankCaller is an auto generated read-only Go binding around an Ethereum contract.
type MainchainTFuelTokenBankCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainchainTFuelTokenBankTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MainchainTFuelTokenBankTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainchainTFuelTokenBankFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MainchainTFuelTokenBankFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainchainTFuelTokenBankSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MainchainTFuelTokenBankSession struct {
	Contract     *MainchainTFuelTokenBank // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// MainchainTFuelTokenBankCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MainchainTFuelTokenBankCallerSession struct {
	Contract *MainchainTFuelTokenBankCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// MainchainTFuelTokenBankTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MainchainTFuelTokenBankTransactorSession struct {
	Contract     *MainchainTFuelTokenBankTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// MainchainTFuelTokenBankRaw is an auto generated low-level Go binding around an Ethereum contract.
type MainchainTFuelTokenBankRaw struct {
	Contract *MainchainTFuelTokenBank // Generic contract binding to access the raw methods on
}

// MainchainTFuelTokenBankCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MainchainTFuelTokenBankCallerRaw struct {
	Contract *MainchainTFuelTokenBankCaller // Generic read-only contract binding to access the raw methods on
}

// MainchainTFuelTokenBankTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MainchainTFuelTokenBankTransactorRaw struct {
	Contract *MainchainTFuelTokenBankTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMainchainTFuelTokenBank creates a new instance of MainchainTFuelTokenBank, bound to a specific deployed contract.
func NewMainchainTFuelTokenBank(address common.Address, backend bind.ContractBackend) (*MainchainTFuelTokenBank, error) {
	contract, err := bindMainchainTFuelTokenBank(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MainchainTFuelTokenBank{MainchainTFuelTokenBankCaller: MainchainTFuelTokenBankCaller{contract: contract}, MainchainTFuelTokenBankTransactor: MainchainTFuelTokenBankTransactor{contract: contract}, MainchainTFuelTokenBankFilterer: MainchainTFuelTokenBankFilterer{contract: contract}}, nil
}

// NewMainchainTFuelTokenBankCaller creates a new read-only instance of MainchainTFuelTokenBank, bound to a specific deployed contract.
func NewMainchainTFuelTokenBankCaller(address common.Address, caller bind.ContractCaller) (*MainchainTFuelTokenBankCaller, error) {
	contract, err := bindMainchainTFuelTokenBank(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MainchainTFuelTokenBankCaller{contract: contract}, nil
}

// NewMainchainTFuelTokenBankTransactor creates a new write-only instance of MainchainTFuelTokenBank, bound to a specific deployed contract.
func NewMainchainTFuelTokenBankTransactor(address common.Address, transactor bind.ContractTransactor) (*MainchainTFuelTokenBankTransactor, error) {
	contract, err := bindMainchainTFuelTokenBank(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MainchainTFuelTokenBankTransactor{contract: contract}, nil
}

// NewMainchainTFuelTokenBankFilterer creates a new log filterer instance of MainchainTFuelTokenBank, bound to a specific deployed contract.
func NewMainchainTFuelTokenBankFilterer(address common.Address, filterer bind.ContractFilterer) (*MainchainTFuelTokenBankFilterer, error) {
	contract, err := bindMainchainTFuelTokenBank(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MainchainTFuelTokenBankFilterer{contract: contract}, nil
}

// bindMainchainTFuelTokenBank binds a generic wrapper to an already deployed contract.
func bindMainchainTFuelTokenBank(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MainchainTFuelTokenBankABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MainchainTFuelTokenBank.Contract.MainchainTFuelTokenBankCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MainchainTFuelTokenBank.Contract.MainchainTFuelTokenBankTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MainchainTFuelTokenBank.Contract.MainchainTFuelTokenBankTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MainchainTFuelTokenBank.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MainchainTFuelTokenBank.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MainchainTFuelTokenBank.Contract.contract.Transact(opts, method, params...)
}

// AllDenoms is a free data retrieval call binding the contract method 0x588b1408.
//
// Solidity: function allDenoms(uint256 ) view returns(string)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankCaller) AllDenoms(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _MainchainTFuelTokenBank.contract.Call(opts, &out, "allDenoms", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AllDenoms is a free data retrieval call binding the contract method 0x588b1408.
//
// Solidity: function allDenoms(uint256 ) view returns(string)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankSession) AllDenoms(arg0 *big.Int) (string, error) {
	return _MainchainTFuelTokenBank.Contract.AllDenoms(&_MainchainTFuelTokenBank.CallOpts, arg0)
}

// AllDenoms is a free data retrieval call binding the contract method 0x588b1408.
//
// Solidity: function allDenoms(uint256 ) view returns(string)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankCallerSession) AllDenoms(arg0 *big.Int) (string, error) {
	return _MainchainTFuelTokenBank.Contract.AllDenoms(&_MainchainTFuelTokenBank.CallOpts, arg0)
}

// AllVouchers is a free data retrieval call binding the contract method 0x27ca4df1.
//
// Solidity: function allVouchers(uint256 ) view returns(address)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankCaller) AllVouchers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MainchainTFuelTokenBank.contract.Call(opts, &out, "allVouchers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllVouchers is a free data retrieval call binding the contract method 0x27ca4df1.
//
// Solidity: function allVouchers(uint256 ) view returns(address)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankSession) AllVouchers(arg0 *big.Int) (common.Address, error) {
	return _MainchainTFuelTokenBank.Contract.AllVouchers(&_MainchainTFuelTokenBank.CallOpts, arg0)
}

// AllVouchers is a free data retrieval call binding the contract method 0x27ca4df1.
//
// Solidity: function allVouchers(uint256 ) view returns(address)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankCallerSession) AllVouchers(arg0 *big.Int) (common.Address, error) {
	return _MainchainTFuelTokenBank.Contract.AllVouchers(&_MainchainTFuelTokenBank.CallOpts, arg0)
}

// DenomToVoucherLookup is a free data retrieval call binding the contract method 0x1527b14d.
//
// Solidity: function denomToVoucherLookup(string ) view returns(address contractAddress, bool exists)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankCaller) DenomToVoucherLookup(opts *bind.CallOpts, arg0 string) (struct {
	ContractAddress common.Address
	Exists          bool
}, error) {
	var out []interface{}
	err := _MainchainTFuelTokenBank.contract.Call(opts, &out, "denomToVoucherLookup", arg0)

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
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankSession) DenomToVoucherLookup(arg0 string) (struct {
	ContractAddress common.Address
	Exists          bool
}, error) {
	return _MainchainTFuelTokenBank.Contract.DenomToVoucherLookup(&_MainchainTFuelTokenBank.CallOpts, arg0)
}

// DenomToVoucherLookup is a free data retrieval call binding the contract method 0x1527b14d.
//
// Solidity: function denomToVoucherLookup(string ) view returns(address contractAddress, bool exists)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankCallerSession) DenomToVoucherLookup(arg0 string) (struct {
	ContractAddress common.Address
	Exists          bool
}, error) {
	return _MainchainTFuelTokenBank.Contract.DenomToVoucherLookup(&_MainchainTFuelTokenBank.CallOpts, arg0)
}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(string denom) view returns(bool)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankCaller) Exists(opts *bind.CallOpts, denom string) (bool, error) {
	var out []interface{}
	err := _MainchainTFuelTokenBank.contract.Call(opts, &out, "exists", denom)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(string denom) view returns(bool)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankSession) Exists(denom string) (bool, error) {
	return _MainchainTFuelTokenBank.Contract.Exists(&_MainchainTFuelTokenBank.CallOpts, denom)
}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(string denom) view returns(bool)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankCallerSession) Exists(denom string) (bool, error) {
	return _MainchainTFuelTokenBank.Contract.Exists(&_MainchainTFuelTokenBank.CallOpts, denom)
}

// Exists0 is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address voucherAddress) view returns(bool)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankCaller) Exists0(opts *bind.CallOpts, voucherAddress common.Address) (bool, error) {
	var out []interface{}
	err := _MainchainTFuelTokenBank.contract.Call(opts, &out, "exists0", voucherAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists0 is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address voucherAddress) view returns(bool)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankSession) Exists0(voucherAddress common.Address) (bool, error) {
	return _MainchainTFuelTokenBank.Contract.Exists0(&_MainchainTFuelTokenBank.CallOpts, voucherAddress)
}

// Exists0 is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address voucherAddress) view returns(bool)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankCallerSession) Exists0(voucherAddress common.Address) (bool, error) {
	return _MainchainTFuelTokenBank.Contract.Exists0(&_MainchainTFuelTokenBank.CallOpts, voucherAddress)
}

// GetDenom is a free data retrieval call binding the contract method 0xebda9962.
//
// Solidity: function getDenom(address voucherContractAddr) view returns(string)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankCaller) GetDenom(opts *bind.CallOpts, voucherContractAddr common.Address) (string, error) {
	var out []interface{}
	err := _MainchainTFuelTokenBank.contract.Call(opts, &out, "getDenom", voucherContractAddr)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetDenom is a free data retrieval call binding the contract method 0xebda9962.
//
// Solidity: function getDenom(address voucherContractAddr) view returns(string)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankSession) GetDenom(voucherContractAddr common.Address) (string, error) {
	return _MainchainTFuelTokenBank.Contract.GetDenom(&_MainchainTFuelTokenBank.CallOpts, voucherContractAddr)
}

// GetDenom is a free data retrieval call binding the contract method 0xebda9962.
//
// Solidity: function getDenom(address voucherContractAddr) view returns(string)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankCallerSession) GetDenom(voucherContractAddr common.Address) (string, error) {
	return _MainchainTFuelTokenBank.Contract.GetDenom(&_MainchainTFuelTokenBank.CallOpts, voucherContractAddr)
}

// GetVoucher is a free data retrieval call binding the contract method 0xa2cc6981.
//
// Solidity: function getVoucher(string denom) view returns(address)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankCaller) GetVoucher(opts *bind.CallOpts, denom string) (common.Address, error) {
	var out []interface{}
	err := _MainchainTFuelTokenBank.contract.Call(opts, &out, "getVoucher", denom)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVoucher is a free data retrieval call binding the contract method 0xa2cc6981.
//
// Solidity: function getVoucher(string denom) view returns(address)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankSession) GetVoucher(denom string) (common.Address, error) {
	return _MainchainTFuelTokenBank.Contract.GetVoucher(&_MainchainTFuelTokenBank.CallOpts, denom)
}

// GetVoucher is a free data retrieval call binding the contract method 0xa2cc6981.
//
// Solidity: function getVoucher(string denom) view returns(address)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankCallerSession) GetVoucher(denom string) (common.Address, error) {
	return _MainchainTFuelTokenBank.Contract.GetVoucher(&_MainchainTFuelTokenBank.CallOpts, denom)
}

// HasSufficientLockedAmount is a free data retrieval call binding the contract method 0x5d9c701d.
//
// Solidity: function hasSufficientLockedAmount(uint256 subchainID, address receiver, uint256 amount) view returns(bool)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankCaller) HasSufficientLockedAmount(opts *bind.CallOpts, subchainID *big.Int, receiver common.Address, amount *big.Int) (bool, error) {
	var out []interface{}
	err := _MainchainTFuelTokenBank.contract.Call(opts, &out, "hasSufficientLockedAmount", subchainID, receiver, amount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasSufficientLockedAmount is a free data retrieval call binding the contract method 0x5d9c701d.
//
// Solidity: function hasSufficientLockedAmount(uint256 subchainID, address receiver, uint256 amount) view returns(bool)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankSession) HasSufficientLockedAmount(subchainID *big.Int, receiver common.Address, amount *big.Int) (bool, error) {
	return _MainchainTFuelTokenBank.Contract.HasSufficientLockedAmount(&_MainchainTFuelTokenBank.CallOpts, subchainID, receiver, amount)
}

// HasSufficientLockedAmount is a free data retrieval call binding the contract method 0x5d9c701d.
//
// Solidity: function hasSufficientLockedAmount(uint256 subchainID, address receiver, uint256 amount) view returns(bool)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankCallerSession) HasSufficientLockedAmount(subchainID *big.Int, receiver common.Address, amount *big.Int) (bool, error) {
	return _MainchainTFuelTokenBank.Contract.HasSufficientLockedAmount(&_MainchainTFuelTokenBank.CallOpts, subchainID, receiver, amount)
}

// VoucherAddressToDenomLookup is a free data retrieval call binding the contract method 0x60569b5e.
//
// Solidity: function voucherAddressToDenomLookup(address ) view returns(string denom, bool exists)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankCaller) VoucherAddressToDenomLookup(opts *bind.CallOpts, arg0 common.Address) (struct {
	Denom  string
	Exists bool
}, error) {
	var out []interface{}
	err := _MainchainTFuelTokenBank.contract.Call(opts, &out, "voucherAddressToDenomLookup", arg0)

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
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankSession) VoucherAddressToDenomLookup(arg0 common.Address) (struct {
	Denom  string
	Exists bool
}, error) {
	return _MainchainTFuelTokenBank.Contract.VoucherAddressToDenomLookup(&_MainchainTFuelTokenBank.CallOpts, arg0)
}

// VoucherAddressToDenomLookup is a free data retrieval call binding the contract method 0x60569b5e.
//
// Solidity: function voucherAddressToDenomLookup(address ) view returns(string denom, bool exists)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankCallerSession) VoucherAddressToDenomLookup(arg0 common.Address) (struct {
	Denom  string
	Exists bool
}, error) {
	return _MainchainTFuelTokenBank.Contract.VoucherAddressToDenomLookup(&_MainchainTFuelTokenBank.CallOpts, arg0)
}

// SendToChain is a paid mutator transaction binding the contract method 0x2ee015d7.
//
// Solidity: function sendToChain(uint256 targetChainID, address subchainTokenReceiver) payable returns()
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankTransactor) SendToChain(opts *bind.TransactOpts, targetChainID *big.Int, subchainTokenReceiver common.Address) (*types.Transaction, error) {
	return _MainchainTFuelTokenBank.contract.Transact(opts, "sendToChain", targetChainID, subchainTokenReceiver)
}

// SendToChain is a paid mutator transaction binding the contract method 0x2ee015d7.
//
// Solidity: function sendToChain(uint256 targetChainID, address subchainTokenReceiver) payable returns()
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankSession) SendToChain(targetChainID *big.Int, subchainTokenReceiver common.Address) (*types.Transaction, error) {
	return _MainchainTFuelTokenBank.Contract.SendToChain(&_MainchainTFuelTokenBank.TransactOpts, targetChainID, subchainTokenReceiver)
}

// SendToChain is a paid mutator transaction binding the contract method 0x2ee015d7.
//
// Solidity: function sendToChain(uint256 targetChainID, address subchainTokenReceiver) payable returns()
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankTransactorSession) SendToChain(targetChainID *big.Int, subchainTokenReceiver common.Address) (*types.Transaction, error) {
	return _MainchainTFuelTokenBank.Contract.SendToChain(&_MainchainTFuelTokenBank.TransactOpts, targetChainID, subchainTokenReceiver)
}

// Unlock is a paid mutator transaction binding the contract method 0x8d59151f.
//
// Solidity: function unlock(bytes _data, bytes _sig) payable returns()
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankTransactor) Unlock(opts *bind.TransactOpts, _data []byte, _sig []byte) (*types.Transaction, error) {
	return _MainchainTFuelTokenBank.contract.Transact(opts, "unlock", _data, _sig)
}

// Unlock is a paid mutator transaction binding the contract method 0x8d59151f.
//
// Solidity: function unlock(bytes _data, bytes _sig) payable returns()
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankSession) Unlock(_data []byte, _sig []byte) (*types.Transaction, error) {
	return _MainchainTFuelTokenBank.Contract.Unlock(&_MainchainTFuelTokenBank.TransactOpts, _data, _sig)
}

// Unlock is a paid mutator transaction binding the contract method 0x8d59151f.
//
// Solidity: function unlock(bytes _data, bytes _sig) payable returns()
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankTransactorSession) Unlock(_data []byte, _sig []byte) (*types.Transaction, error) {
	return _MainchainTFuelTokenBank.Contract.Unlock(&_MainchainTFuelTokenBank.TransactOpts, _data, _sig)
}

// MainchainTFuelTokenBankTFeulTokenLockedIterator is returned from FilterTFeulTokenLocked and is used to iterate over the raw logs and unpacked data for TFeulTokenLocked events raised by the MainchainTFuelTokenBank contract.
type MainchainTFuelTokenBankTFeulTokenLockedIterator struct {
	Event *MainchainTFuelTokenBankTFeulTokenLocked // Event containing the contract specifics and raw log

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
func (it *MainchainTFuelTokenBankTFeulTokenLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainchainTFuelTokenBankTFeulTokenLocked)
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
		it.Event = new(MainchainTFuelTokenBankTFeulTokenLocked)
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
func (it *MainchainTFuelTokenBankTFeulTokenLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainchainTFuelTokenBankTFeulTokenLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainchainTFuelTokenBankTFeulTokenLocked represents a TFeulTokenLocked event raised by the MainchainTFuelTokenBank contract.
type MainchainTFuelTokenBankTFeulTokenLocked struct {
	TargetChainID         *big.Int
	MainchainTokenSender  common.Address
	SubchainTokenReceiver common.Address
	LockedAmount          *big.Int
	Nonce                 *big.Int
	Denom                 string
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterTFeulTokenLocked is a free log retrieval operation binding the contract event 0xb0d72a5b25cc833270260cccd76510dddd5e3a0aa8e49018804c93ed3b256838.
//
// Solidity: event TFeulTokenLocked(uint256 targetChainID, address mainchainTokenSender, address subchainTokenReceiver, uint256 lockedAmount, uint256 nonce, string denom)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankFilterer) FilterTFeulTokenLocked(opts *bind.FilterOpts) (*MainchainTFuelTokenBankTFeulTokenLockedIterator, error) {

	logs, sub, err := _MainchainTFuelTokenBank.contract.FilterLogs(opts, "TFeulTokenLocked")
	if err != nil {
		return nil, err
	}
	return &MainchainTFuelTokenBankTFeulTokenLockedIterator{contract: _MainchainTFuelTokenBank.contract, event: "TFeulTokenLocked", logs: logs, sub: sub}, nil
}

// WatchTFeulTokenLocked is a free log subscription operation binding the contract event 0xb0d72a5b25cc833270260cccd76510dddd5e3a0aa8e49018804c93ed3b256838.
//
// Solidity: event TFeulTokenLocked(uint256 targetChainID, address mainchainTokenSender, address subchainTokenReceiver, uint256 lockedAmount, uint256 nonce, string denom)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankFilterer) WatchTFeulTokenLocked(opts *bind.WatchOpts, sink chan<- *MainchainTFuelTokenBankTFeulTokenLocked) (event.Subscription, error) {

	logs, sub, err := _MainchainTFuelTokenBank.contract.WatchLogs(opts, "TFeulTokenLocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainchainTFuelTokenBankTFeulTokenLocked)
				if err := _MainchainTFuelTokenBank.contract.UnpackLog(event, "TFeulTokenLocked", log); err != nil {
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

// ParseTFeulTokenLocked is a log parse operation binding the contract event 0xb0d72a5b25cc833270260cccd76510dddd5e3a0aa8e49018804c93ed3b256838.
//
// Solidity: event TFeulTokenLocked(uint256 targetChainID, address mainchainTokenSender, address subchainTokenReceiver, uint256 lockedAmount, uint256 nonce, string denom)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankFilterer) ParseTFeulTokenLocked(log types.Log) (*MainchainTFuelTokenBankTFeulTokenLocked, error) {
	event := new(MainchainTFuelTokenBankTFeulTokenLocked)
	if err := _MainchainTFuelTokenBank.contract.UnpackLog(event, "TFeulTokenLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}