package predeployed

import (
	"encoding/hex"

	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/theta/crypto"
	"github.com/thetatoken/theta/ledger/types"
	"github.com/thetatoken/theta/rlp"

	"github.com/thetatoken/thetasubchain/core"
	slst "github.com/thetatoken/thetasubchain/ledger/state"
)

// Bytecode of the smart contracts hardcoded in the genesis block through pre-deployment
const TNT20TokenBankContractBytecode = "608060405234801561001057600080fd5b50613f9b806100206000396000f3fe6080604052600436106200009e5760003560e01c8063a2cc69811162000061578063a2cc698114620001f9578063b83f2901146200023d578063ba2c329c146200025d578063ebda9962146200028b578063f6a3d24e14620002cf576200009e565b80631527b14d14620000a3578063261a323e14620000e857806327ca4df1146200012c578063588b1408146200017057806360569b5e14620001b4575b600080fd5b348015620000b057600080fd5b50620000cf6004803603810190620000c991906200143a565b62000313565b604051620000df92919062001836565b60405180910390f35b348015620000f557600080fd5b506200011460048036038101906200010e91906200143a565b6200037a565b60405162000123919062001902565b60405180910390f35b3480156200013957600080fd5b50620001586004803603810190620001529190620015ff565b620003b3565b60405162000167919062001819565b60405180910390f35b3480156200017d57600080fd5b506200019c6004803603810190620001969190620015ff565b620003f3565b604051620001ab91906200191f565b60405180910390f35b348015620001c157600080fd5b50620001e06004803603810190620001da9190620013a4565b620004a8565b604051620001f092919062001943565b60405180910390f35b3480156200020657600080fd5b506200022560048036038101906200021f91906200143a565b62000569565b60405162000234919062001819565b60405180910390f35b6200025b60048036038101906200025591906200148b565b6200062e565b005b3480156200026a57600080fd5b5062000289600480360381019062000283919062001506565b6200082a565b005b3480156200029857600080fd5b50620002b76004803603810190620002b19190620013a4565b62000b54565b604051620002c691906200191f565b60405180910390f35b348015620002dc57600080fd5b50620002fb6004803603810190620002f59190620013a4565b62000c8b565b6040516200030a919062001902565b60405180910390f35b6000818051602081018201805184825260208301602085012081835280955050505050506000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060000160149054906101000a900460ff16905082565b600080826040516200038d919062001800565b908152602001604051809103902060000160149054906101000a900460ff169050919050565b60028181548110620003c457600080fd5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600381815481106200040457600080fd5b906000526020600020016000915090508054620004219062001c6f565b80601f01602080910402602001604051908101604052809291908181526020018280546200044f9062001c6f565b8015620004a05780601f106200047457610100808354040283529160200191620004a0565b820191906000526020600020905b8154815290600101906020018083116200048257829003601f168201915b505050505081565b6001602052806000526040600020600091509050806000018054620004cd9062001c6f565b80601f0160208091040260200160405190810160405280929190818152602001828054620004fb9062001c6f565b80156200054c5780601f1062000520576101008083540402835291602001916200054c565b820191906000526020600020905b8154815290600101906020018083116200052e57829003601f168201915b5050505050908060010160009054906101000a900460ff16905082565b6000806000836040516200057e919062001800565b90815260200160405180910390206040518060400160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160149054906101000a900460ff16151515158152505090508060200151156200062357806000015191505062000629565b60009150505b919050565b3073ffffffffffffffffffffffffffffffffffffffff1663261a323e846040518263ffffffff1660e01b81526004016200066991906200191f565b60206040518083038186803b1580156200068257600080fd5b505afa15801562000697573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620006bd919062001408565b620006ff576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620006f690620019bb565b60405180910390fd5b60003073ffffffffffffffffffffffffffffffffffffffff1663a2cc6981856040518263ffffffff1660e01b81526004016200073c91906200191f565b60206040518083038186803b1580156200075557600080fd5b505afa1580156200076a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620007909190620013d6565b90506000339050620007a482828562000ce4565b8373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1686604051620007e2919062001800565b60405180910390207f2875eb0bd4ce0a02cd43de28971f5475143104a7e6d2af8e11214433a0d711be866040516200081b919062001a21565b60405180910390a45050505050565b600080602067ffffffffffffffff8111156200084b576200084a62001db6565b5b6040519080825280601f01601f1916602001820160405280156200087e5781602001600182028036833780820191505090505b50905060008060b573ffffffffffffffffffffffffffffffffffffffff1683604051620008ac9190620017e7565b6000604051808303816000865af19150503d8060008114620008eb576040519150601f19603f3d011682016040523d82523d6000602084013e620008f0565b606091505b50915091508162000938576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200092f9062001999565b60405180910390fd5b6000620009458262000d60565b90506001811494508462000990576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200098790620019ff565b60405180910390fd5b3073ffffffffffffffffffffffffffffffffffffffff1663261a323e8c6040518263ffffffff1660e01b8152600401620009cb91906200191f565b60206040518083038186803b158015620009e457600080fd5b505afa158015620009f9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000a1f919062001408565b62000a3e5762000a3d8b62000a378d8d8d8d62000e59565b62000eb1565b5b60003073ffffffffffffffffffffffffffffffffffffffff1663a2cc69818d6040518263ffffffff1660e01b815260040162000a7b91906200191f565b60206040518083038186803b15801562000a9457600080fd5b505afa15801562000aa9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000acf9190620013d6565b905062000ade81898962001179565b8773ffffffffffffffffffffffffffffffffffffffff168c60405162000b05919062001800565b60405180910390207f4447df32d078a22540f11de3ca36c2145245056e49a8c3ea40135fa7d66873138960405162000b3e919062001a21565b60405180910390a3505050505050505050505050565b60606000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060405180604001604052908160008201805462000bb49062001c6f565b80601f016020809104026020016040519081016040528092919081815260200182805462000be29062001c6f565b801562000c335780601f1062000c075761010080835404028352916020019162000c33565b820191906000526020600020905b81548152906001019060200180831162000c1557829003601f168201915b505050505081526020016001820160009054906101000a900460ff161515151581525050905080602001511562000c7257806000015191505062000c86565b604051806020016040528060008152509150505b919050565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900460ff169050919050565b60008390508073ffffffffffffffffffffffffffffffffffffffff16639dc29fac84846040518363ffffffff1660e01b815260040162000d26929190620018d5565b600060405180830381600087803b15801562000d4157600080fd5b505af115801562000d56573d6000803e3d6000fd5b5050505050505050565b600080600083519050602081111562000db0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040162000da790620019dd565b60405180910390fd5b60005b8181101562000e4b5760088183602062000dce919062001b98565b62000dda919062001ada565b62000de6919062001b37565b60ff60f81b86838151811062000e015762000e0062001d87565b5b602001015160f81c60f81b167effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916901c83179250808062000e429062001cdb565b91505062000db3565b508160001c92505050919050565b6000803090506000818787878760405162000e7490620011f5565b62000e8495949392919062001863565b604051809103906000f08015801562000ea1573d6000803e3d6000fd5b5090508092505050949350505050565b3073ffffffffffffffffffffffffffffffffffffffff1663261a323e836040518263ffffffff1660e01b815260040162000eec91906200191f565b60206040518083038186803b15801562000f0557600080fd5b505afa15801562000f1a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000f40919062001408565b1562000f83576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040162000f7a9062001977565b60405180910390fd5b60405180604001604052808273ffffffffffffffffffffffffffffffffffffffff1681526020016001151581525060008360405162000fc3919062001800565b908152602001604051809103902060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548160ff021916908315150217905550905050604051806040016040528083815260200160011515815250600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000019080519060200190620010b092919062001203565b5060208201518160010160006101000a81548160ff0219169083151502179055509050506002819080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506003829080600181540180825580915050600190039060005260206000200160009091909190915090805190602001906200117492919062001203565b505050565b60008390508073ffffffffffffffffffffffffffffffffffffffff166340c10f1984846040518363ffffffff1660e01b8152600401620011bb929190620018d5565b600060405180830381600087803b158015620011d657600080fd5b505af1158015620011eb573d6000803e3d6000fd5b5050505050505050565b611f1c806200204a83390190565b828054620012119062001c6f565b90600052602060002090601f01602090048101928262001235576000855562001281565b82601f106200125057805160ff191683800117855562001281565b8280016001018555821562001281579182015b828111156200128057825182559160200191906001019062001263565b5b50905062001290919062001294565b5090565b5b80821115620012af57600081600090555060010162001295565b5090565b6000620012ca620012c48462001a67565b62001a3e565b905082815260208101848484011115620012e957620012e862001dea565b5b620012f684828562001c2a565b509392505050565b6000813590506200130f8162001fe1565b92915050565b600081519050620013268162001fe1565b92915050565b6000815190506200133d8162001ffb565b92915050565b600082601f8301126200135b576200135a62001de5565b5b81356200136d848260208601620012b3565b91505092915050565b600081359050620013878162002015565b92915050565b6000813590506200139e816200202f565b92915050565b600060208284031215620013bd57620013bc62001df4565b5b6000620013cd84828501620012fe565b91505092915050565b600060208284031215620013ef57620013ee62001df4565b5b6000620013ff8482850162001315565b91505092915050565b60006020828403121562001421576200142062001df4565b5b600062001431848285016200132c565b91505092915050565b60006020828403121562001453576200145262001df4565b5b600082013567ffffffffffffffff81111562001474576200147362001def565b5b620014828482850162001343565b91505092915050565b600080600060608486031215620014a757620014a662001df4565b5b600084013567ffffffffffffffff811115620014c857620014c762001def565b5b620014d68682870162001343565b9350506020620014e986828701620012fe565b9250506040620014fc8682870162001376565b9150509250925092565b60008060008060008060c0878903121562001526576200152562001df4565b5b600087013567ffffffffffffffff81111562001547576200154662001def565b5b6200155589828a0162001343565b965050602087013567ffffffffffffffff81111562001579576200157862001def565b5b6200158789828a0162001343565b955050604087013567ffffffffffffffff811115620015ab57620015aa62001def565b5b620015b989828a0162001343565b9450506060620015cc89828a016200138d565b9350506080620015df89828a01620012fe565b92505060a0620015f289828a0162001376565b9150509295509295509295565b60006020828403121562001618576200161762001df4565b5b6000620016288482850162001376565b91505092915050565b6200163c8162001bd3565b82525050565b6200164d8162001be7565b82525050565b6000620016608262001a9d565b6200166c818562001ab3565b93506200167e81856020860162001c39565b80840191505092915050565b6000620016978262001aa8565b620016a3818562001abe565b9350620016b581856020860162001c39565b620016c08162001df9565b840191505092915050565b6000620016d88262001aa8565b620016e4818562001acf565b9350620016f681856020860162001c39565b80840191505092915050565b60006200171160398362001abe565b91506200171e8262001e0a565b604082019050919050565b600062001738604d8362001abe565b9150620017458262001e59565b606082019050919050565b60006200175f60448362001abe565b91506200176c8262001ece565b606082019050919050565b60006200178660258362001abe565b9150620017938262001f43565b604082019050919050565b6000620017ad60388362001abe565b9150620017ba8262001f92565b604082019050919050565b620017d08162001c13565b82525050565b620017e18162001c1d565b82525050565b6000620017f5828462001653565b915081905092915050565b60006200180e8284620016cb565b915081905092915050565b600060208201905062001830600083018462001631565b92915050565b60006040820190506200184d600083018562001631565b6200185c602083018462001642565b9392505050565b600060a0820190506200187a600083018862001631565b81810360208301526200188e81876200168a565b90508181036040830152620018a481866200168a565b90508181036060830152620018ba81856200168a565b9050620018cb6080830184620017d6565b9695505050505050565b6000604082019050620018ec600083018562001631565b620018fb6020830184620017c5565b9392505050565b600060208201905062001919600083018462001642565b92915050565b600060208201905081810360008301526200193b81846200168a565b905092915050565b600060408201905081810360008301526200195f81856200168a565b905062001970602083018462001642565b9392505050565b60006020820190508181036000830152620019928162001702565b9050919050565b60006020820190508181036000830152620019b48162001729565b9050919050565b60006020820190508181036000830152620019d68162001750565b9050919050565b60006020820190508181036000830152620019f88162001777565b9050919050565b6000602082019050818103600083015262001a1a816200179e565b9050919050565b600060208201905062001a386000830184620017c5565b92915050565b600062001a4a62001a5d565b905062001a58828262001ca5565b919050565b6000604051905090565b600067ffffffffffffffff82111562001a855762001a8462001db6565b5b62001a908262001df9565b9050602081019050919050565b600081519050919050565b600081519050919050565b600081905092915050565b600082825260208201905092915050565b600081905092915050565b600062001ae78262001c13565b915062001af48362001c13565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111562001b2c5762001b2b62001d29565b5b828201905092915050565b600062001b448262001c13565b915062001b518362001c13565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161562001b8d5762001b8c62001d29565b5b828202905092915050565b600062001ba58262001c13565b915062001bb28362001c13565b92508282101562001bc85762001bc762001d29565b5b828203905092915050565b600062001be08262001bf3565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b8381101562001c5957808201518184015260208101905062001c3c565b8381111562001c69576000848401525b50505050565b6000600282049050600182168062001c8857607f821691505b6020821081141562001c9f5762001c9e62001d58565b5b50919050565b62001cb08262001df9565b810181811067ffffffffffffffff8211171562001cd25762001cd162001db6565b5b80604052505050565b600062001ce88262001c13565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82141562001d1e5762001d1d62001d29565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f566f75636865724d61702e616464566f75636865723a20616e20766f7563686560008201527f7220636f6e747261637420616c72656164792065786973747300000000000000602082015250565b7f537562636861696e544675656c546f6b656e42616e6b2e70726976696c65676560008201527f644163636573734f6e6c793a206661696c656420746f20636865636b2074686560208201527f20616363657373206c6576656c00000000000000000000000000000000000000604082015250565b7f537562636861696e544e543230546f6b656e42616e6b2e6275726e566f75636860008201527f6572733a20766f756368657220636f6e747261637420646f6573206e6f74206560208201527f7869737400000000000000000000000000000000000000000000000000000000604082015250565b7f627974657320746f2075696e7432353620636f6e76657273696f6e206f76657260008201527f666c6f7773000000000000000000000000000000000000000000000000000000602082015250565b7f566f75636865724d61702e70726976696c656765644163636573734f6e6c793a60008201527f20696e73756666696369656e742070726976696c656765730000000000000000602082015250565b62001fec8162001bd3565b811462001ff857600080fd5b50565b620020068162001be7565b81146200201257600080fd5b50565b620020208162001c13565b81146200202c57600080fd5b50565b6200203a8162001c1d565b81146200204657600080fd5b5056fe60806040523480156200001157600080fd5b5060405162001f1c38038062001f1c83398181016040528101906200003791906200024b565b848383816003908051906020019062000052929190620000ef565b5080600490805190602001906200006b929190620000ef565b50505080600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550508360069080519060200190620000c8929190620000ef565b5080600760006101000a81548160ff021916908360ff160217905550505050505062000529565b828054620000fd9062000406565b90600052602060002090601f0160209004810192826200012157600085556200016d565b82601f106200013c57805160ff19168380011785556200016d565b828001600101855582156200016d579182015b828111156200016c5782518255916020019190600101906200014f565b5b5090506200017c919062000180565b5090565b5b808211156200019b57600081600090555060010162000181565b5090565b6000620001b6620001b08462000359565b62000330565b905082815260208101848484011115620001d557620001d4620004d5565b5b620001e2848285620003d0565b509392505050565b600081519050620001fb81620004f5565b92915050565b600082601f830112620002195762000218620004d0565b5b81516200022b8482602086016200019f565b91505092915050565b60008151905062000245816200050f565b92915050565b600080600080600060a086880312156200026a5762000269620004df565b5b60006200027a88828901620001ea565b955050602086015167ffffffffffffffff8111156200029e576200029d620004da565b5b620002ac8882890162000201565b945050604086015167ffffffffffffffff811115620002d057620002cf620004da565b5b620002de8882890162000201565b935050606086015167ffffffffffffffff811115620003025762000301620004da565b5b620003108882890162000201565b9250506080620003238882890162000234565b9150509295509295909350565b60006200033c6200034f565b90506200034a82826200043c565b919050565b6000604051905090565b600067ffffffffffffffff821115620003775762000376620004a1565b5b6200038282620004e4565b9050602081019050919050565b60006200039c82620003a3565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600060ff82169050919050565b60005b83811015620003f0578082015181840152602081019050620003d3565b8381111562000400576000848401525b50505050565b600060028204905060018216806200041f57607f821691505b6020821081141562000436576200043562000472565b5b50919050565b6200044782620004e4565b810181811067ffffffffffffffff82111715620004695762000468620004a1565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b62000500816200038f565b81146200050c57600080fd5b50565b6200051a81620003c3565b81146200052657600080fd5b50565b6119e380620005396000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c806370a0823111610097578063a457c2d711610066578063a457c2d714610288578063a9059cbb146102b8578063c370b042146102e8578063dd62ed3e14610306576100f5565b806370a08231146102005780638da5cb5b1461023057806395d89b411461024e5780639dc29fac1461026c576100f5565b806323b872dd116100d357806323b872dd14610166578063313ce5671461019657806339509351146101b457806340c10f19146101e4576100f5565b806306fdde03146100fa578063095ea7b31461011857806318160ddd14610148575b600080fd5b610102610336565b60405161010f9190611349565b60405180910390f35b610132600480360381019061012d9190611100565b6103c8565b60405161013f919061132e565b60405180910390f35b6101506103eb565b60405161015d91906114ab565b60405180910390f35b610180600480360381019061017b91906110ad565b6103f5565b60405161018d919061132e565b60405180910390f35b61019e610424565b6040516101ab91906114c6565b60405180910390f35b6101ce60048036038101906101c99190611100565b61043b565b6040516101db919061132e565b60405180910390f35b6101fe60048036038101906101f99190611100565b610472565b005b61021a60048036038101906102159190611040565b6104da565b60405161022791906114ab565b60405180910390f35b610238610522565b6040516102459190611313565b60405180910390f35b610256610548565b6040516102639190611349565b60405180910390f35b61028660048036038101906102819190611100565b6105da565b005b6102a2600480360381019061029d9190611100565b610642565b6040516102af919061132e565b60405180910390f35b6102d260048036038101906102cd9190611100565b6106b9565b6040516102df919061132e565b60405180910390f35b6102f06106dc565b6040516102fd9190611349565b60405180910390f35b610320600480360381019061031b919061106d565b61076e565b60405161032d91906114ab565b60405180910390f35b6060600380546103459061160f565b80601f01602080910402602001604051908101604052809291908181526020018280546103719061160f565b80156103be5780601f10610393576101008083540402835291602001916103be565b820191906000526020600020905b8154815290600101906020018083116103a157829003601f168201915b5050505050905090565b6000806103d36107f5565b90506103e08185856107fd565b600191505092915050565b6000600254905090565b6000806104006107f5565b905061040d8582856109c8565b610418858585610a54565b60019150509392505050565b6000600760009054906101000a900460ff16905090565b6000806104466107f5565b9050610467818585610458858961076e565b61046291906114fd565b6107fd565b600191505092915050565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146104cc57600080fd5b6104d68282610cd5565b5050565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6060600480546105579061160f565b80601f01602080910402602001604051908101604052809291908181526020018280546105839061160f565b80156105d05780601f106105a5576101008083540402835291602001916105d0565b820191906000526020600020905b8154815290600101906020018083116105b357829003601f168201915b5050505050905090565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461063457600080fd5b61063e8282610e35565b5050565b60008061064d6107f5565b9050600061065b828661076e565b9050838110156106a0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106979061146b565b60405180910390fd5b6106ad82868684036107fd565b60019250505092915050565b6000806106c46107f5565b90506106d1818585610a54565b600191505092915050565b6060600680546106eb9061160f565b80601f01602080910402602001604051908101604052809291908181526020018280546107179061160f565b80156107645780601f1061073957610100808354040283529160200191610764565b820191906000526020600020905b81548152906001019060200180831161074757829003601f168201915b5050505050905090565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141561086d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108649061144b565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156108dd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108d4906113ab565b60405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040516109bb91906114ab565b60405180910390a3505050565b60006109d4848461076e565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610a4e5781811015610a40576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a37906113cb565b60405180910390fd5b610a4d84848484036107fd565b5b50505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610ac4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610abb9061142b565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610b34576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b2b9061136b565b60405180910390fd5b610b3f83838361100c565b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811015610bc5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bbc906113eb565b60405180910390fd5b8181036000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610c5891906114fd565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610cbc91906114ab565b60405180910390a3610ccf848484611011565b50505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610d45576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d3c9061148b565b60405180910390fd5b610d516000838361100c565b8060026000828254610d6391906114fd565b92505081905550806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610db891906114fd565b925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610e1d91906114ab565b60405180910390a3610e3160008383611011565b5050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610ea5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e9c9061140b565b60405180910390fd5b610eb18260008361100c565b60008060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811015610f37576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f2e9061138b565b60405180910390fd5b8181036000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508160026000828254610f8e9190611553565b92505081905550600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610ff391906114ab565b60405180910390a361100783600084611011565b505050565b505050565b505050565b6000813590506110258161197f565b92915050565b60008135905061103a81611996565b92915050565b6000602082840312156110565761105561169f565b5b600061106484828501611016565b91505092915050565b600080604083850312156110845761108361169f565b5b600061109285828601611016565b92505060206110a385828601611016565b9150509250929050565b6000806000606084860312156110c6576110c561169f565b5b60006110d486828701611016565b93505060206110e586828701611016565b92505060406110f68682870161102b565b9150509250925092565b600080604083850312156111175761111661169f565b5b600061112585828601611016565b92505060206111368582860161102b565b9150509250929050565b61114981611587565b82525050565b61115881611599565b82525050565b6000611169826114e1565b61117381856114ec565b93506111838185602086016115dc565b61118c816116a4565b840191505092915050565b60006111a46023836114ec565b91506111af826116b5565b604082019050919050565b60006111c76022836114ec565b91506111d282611704565b604082019050919050565b60006111ea6022836114ec565b91506111f582611753565b604082019050919050565b600061120d601d836114ec565b9150611218826117a2565b602082019050919050565b60006112306026836114ec565b915061123b826117cb565b604082019050919050565b60006112536021836114ec565b915061125e8261181a565b604082019050919050565b60006112766025836114ec565b915061128182611869565b604082019050919050565b60006112996024836114ec565b91506112a4826118b8565b604082019050919050565b60006112bc6025836114ec565b91506112c782611907565b604082019050919050565b60006112df601f836114ec565b91506112ea82611956565b602082019050919050565b6112fe816115c5565b82525050565b61130d816115cf565b82525050565b60006020820190506113286000830184611140565b92915050565b6000602082019050611343600083018461114f565b92915050565b60006020820190508181036000830152611363818461115e565b905092915050565b6000602082019050818103600083015261138481611197565b9050919050565b600060208201905081810360008301526113a4816111ba565b9050919050565b600060208201905081810360008301526113c4816111dd565b9050919050565b600060208201905081810360008301526113e481611200565b9050919050565b6000602082019050818103600083015261140481611223565b9050919050565b6000602082019050818103600083015261142481611246565b9050919050565b6000602082019050818103600083015261144481611269565b9050919050565b600060208201905081810360008301526114648161128c565b9050919050565b60006020820190508181036000830152611484816112af565b9050919050565b600060208201905081810360008301526114a4816112d2565b9050919050565b60006020820190506114c060008301846112f5565b92915050565b60006020820190506114db6000830184611304565b92915050565b600081519050919050565b600082825260208201905092915050565b6000611508826115c5565b9150611513836115c5565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561154857611547611641565b5b828201905092915050565b600061155e826115c5565b9150611569836115c5565b92508282101561157c5761157b611641565b5b828203905092915050565b6000611592826115a5565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b60005b838110156115fa5780820151818401526020810190506115df565b83811115611609576000848401525b50505050565b6000600282049050600182168061162757607f821691505b6020821081141561163b5761163a611670565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600080fd5b6000601f19601f8301169050919050565b7f45524332303a207472616e7366657220746f20746865207a65726f206164647260008201527f6573730000000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60008201527f6365000000000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a20617070726f766520746f20746865207a65726f20616464726560008201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a20696e73756666696369656e7420616c6c6f77616e6365000000600082015250565b7f45524332303a207472616e7366657220616d6f756e742065786365656473206260008201527f616c616e63650000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360008201527f7300000000000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a207472616e736665722066726f6d20746865207a65726f20616460008201527f6472657373000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760008201527f207a65726f000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a206d696e7420746f20746865207a65726f206164647265737300600082015250565b61198881611587565b811461199357600080fd5b50565b61199f816115c5565b81146119aa57600080fd5b5056fea26469706673582212204555a2c6fb885b29dbbe7c23d5e8cc93009afdbd7d4699d00508dc55b11f4d5164736f6c63430008070033a2646970667358221220878e4a54c99215f75456a20540ec7fdda6651a68fd1c7319fd677395ff52dcdd64736f6c63430008070033"

var mintTNT20VouchersFuncSelector = crypto.Keccak256([]byte("mintVouchers(string,string,string,uint8,address,uint)"))[:4]

// TNT20TokenBank implements the TokenBank interface.
type TNT20TokenBank struct {
}

func NewTNT20TokenBank() *TNT20TokenBank {
	return &TNT20TokenBank{}
}

// Mint vouchers for the token transferred cross-chain. If the voucher contract for the token does not yet exist, the
// TokenBank contract deploys the the vouncher contract first and then mints the vouchers in the same call.
func (tb *TNT20TokenBank) GenerateMintVouchersProxySctx(blockProposer common.Address, view *slst.StoreView, ccte *core.CrossChainTNT20TransferEvent) (*types.SmartContractTx, error) {
	voucherReceiver := ccte.Receiver
	denom := ccte.Denom
	name := ccte.Name
	symbol := ccte.Symbol
	decimals := ccte.Decimals
	amount := ccte.Amount
	callData, err := rlp.EncodeToBytes([]interface{}{
		mintTNT20VouchersFuncSelector,
		denom,
		name,
		symbol,
		decimals,
		hex.EncodeToString(voucherReceiver.Bytes()),
		amount,
	})
	if err != nil {
		return nil, err
	}

	tnt20TokenBankContractAddr := view.GetTNT20TokenBankContractAddress()
	sctx, err := constructProxySmartContractTx(blockProposer, *tnt20TokenBankContractAddr, callData)
	if err != nil {
		return nil, err
	}

	return sctx, nil
}