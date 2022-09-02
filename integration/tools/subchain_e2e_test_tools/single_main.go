package main

import (
	"fmt"
	"log"
	"math/big"

	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/thetasubchain/eth/ethclient"
	ct "github.com/thetatoken/thetasubchain/interchain/contracts/accessors"
)

const CHAIN_ID_OFFSET int64 = 360

func config() []string {
	subchainID = big.NewInt(360777)

	wthetaAddress = common.HexToAddress("0x7d73424a8256C0b2BA245e5d5a3De8820E45F390")
	registerOnMainchainAddress = common.HexToAddress("0x08425D9Df219f93d5763c3e85204cb5B4cE33aAa")
	governanceTokenAddress = common.HexToAddress("0x6E05f58eEddA592f34DD9105b1827f252c509De0")
	tnt20VoucherContractAddress = common.HexToAddress("0x59AF421cB35fc23aB6C8ee42743e6176040031f4")
	tnt20TokenBankAddress = common.HexToAddress("0x2Ce636d6240f8955d085a896e12429f8B3c7db26")
	subchaintnt20TokenBankAddress = common.HexToAddress("0x47e9Fbef8C83A1714F1951F142132E6e90F5fa5D")

	tnt721TokenBankAddress = common.HexToAddress("0xEd8d61f42dC1E56aE992D333A4992C3796b22A74")
	tnt721VoucherContractAddress = common.HexToAddress("0x47eb28D8139A188C5686EedE1E9D8EDE3Afdd543")
	Subchaintnt721TokenBankAddress = common.HexToAddress("0x8Be503bcdEd90ED42Eff31f56199399B2b0154CA")

	tfuelTokenbankAddress = common.HexToAddress("0x7f1C87Bd3a22159b8a2E5D195B1a3283D10ea895")
	subchainTfuelTokenBank = common.HexToAddress("0x5a443704dd4B594B382c22a083e2BD3090A6feF3")

	mainchainTNT1155TokenBankAddress = common.HexToAddress("0x27e5Ee255a177D1902D7FF48D66f950ed9408867")
	mainchainTNT1155VoucherContractAddress = common.HexToAddress("0xc7253857256391E518c4aF60aDa5Eb5972Dd6Dbc")
	subchainTNT1155TokenBankAddress = common.HexToAddress("0x47c5e40890bcE4a473A49D7501808b9633F29782")
	subchainTNT1155VoucherContractAddress = common.HexToAddress("0x0ede92cac9161f6c397a604de508dcd1e6f43e61")

	var map1 []string
	map1 = append(map1, "1111111111111111111111111111111111111111111111111111111111111111")
	map1 = append(map1, "93a90ea508331dfdf27fb79757d4250b4e84954927ba0073cd67454ac432c737")
	map1 = append(map1, "3333333333333333333333333333333333333333333333333333333333333333")
	map1 = append(map1, "4444444444444444444444444444444444444444444444444444444444444444")
	map1 = append(map1, "5555555555555555555555555555555555555555555555555555555555555555")
	map1 = append(map1, "6666666666666666666666666666666666666666666666666666666666666666")
	map1 = append(map1, "7777777777777777777777777777777777777777777777777777777777777777")
	map1 = append(map1, "8888888888888888888888888888888888888888888888888888888888888888")
	map1 = append(map1, "9999999999999999999999999999999999999999999999999999999999999999")
	map1 = append(map1, "1000000000000000000000000000000000000000000000000000000000000000")

	map1 = append(map1, "a249a82c42a282e87b2ddef63404d9dfcf6ea501dcaf5d447761765bd74f666d") //10
	map1 = append(map1, "d0d53ac0b4cd47d0ce0060dddc179d04145fea2ee2e0b66c3ee1699c6b492013") //11
	map1 = append(map1, "83f0bb8655139cef4657f90db64a7bb57847038a9bd0ccd87c9b0828e9cbf76d")
	return map1
}
func main() {

	//oneAccountRegister() //register subchain , its id is 360777
	//oneAcoountStake(1)   // use the 10th account to deposit and become validator

	// //Tfuel
	// mainchainTfuelLock(big.NewInt(10)) //lock some tfuel tokens on mainchain and transfer to subchain 360777
	// subchainTfuelBurn(big.NewInt(10))

	// //TNT20
	//mainchainTNT20Lock(big.NewInt(100)) //mainchainTNT20Lock(tokenAmount) last address printed in console is subchain voucher contract address
	//subchainTNT20Lock(big.NewInt(10)) //last address printed in console is mainchain voucher contract address
	// mainchainTNT20Burn(big.NewInt(10))
	// subchainTNT20Burn(big.NewInt(10))

	// //TNT721
	// mainchainTNT721Lock(big.NewInt(13)) //mainchainTNT721Lock(tokenid) last address printed in console is subchain voucher contract address
	// subchainTNT721Lock(big.NewInt(12))  //last address printed in console is mainchain voucher contract address
	// mainchainTNT721Burn(big.NewInt(12))
	// subchainTNT721Burn(big.NewInt(13))
	//MainchainTNT1155Lock()
	//SubchainTNT1155Lock(big.NewInt(1))
	//MainchainTNT1155Burn(big.NewInt(1))
	//SubchainTNT1155Burn(big.NewInt(1))
	//getSubchainTNT1155VoucherMintLog(39,45)
	// mintLockAmount := big.NewInt(100)
	// client, err := ethclient.Dial("http://localhost:18888/rpc")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// var dec18 = new(big.Int)
	// dec18.SetString("1000000000000000000", 10)
	// user := accountList[1].fromAddress

	// instanceTNT20VoucherContract, err := ct.NewTNT20VoucherContract(tnt20VoucherContractAddress, client)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// authAccount0 := mainchainSelectAccount(client, 0)
	// _, err = instanceTNT20VoucherContract.Mint(authAccount0, user, mintLockAmount)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// authUser := mainchainSelectAccount(client, 1)
	// _, err = instanceTNT20VoucherContract.Approve(authUser, tnt20TokenBankAddress, mintLockAmount)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// chs := make([]chan int, 10)
	// for i := 0; i < 10; i++ {
	// 	chs[i] = make(chan int)
	// 	go parallelLock(chs[i])
	// }
	// for _, ch := range chs {
	// 	<-ch
	// }
	query()
}
func query() {
	client, err := ethclient.Dial("http://localhost:19888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	subchainAddress := common.HexToAddress("0x7D7e270b7E279C94b265A535CdbC00Eb62E6e68f")
	subchainTNT20VoucherAddressInstance, _ := ct.NewMockTNT20(subchainAddress, client)
	tx, err := subchainTNT20VoucherAddressInstance.BalanceOf(nil, accountList[1].fromAddress)
	fmt.Println(tx)
	if err != nil {
		log.Fatal(err)
	}
}
func deploy_contracts() {
	subchainClient, err := ethclient.Dial("http://localhost:19888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	auth := subchainSelectAccount(subchainClient, 1)
	address, tx, _, err := ct.DeployMockTNT20(auth, subchainClient)
	fmt.Println("tnt20", address)
	fmt.Println(tx.Hash().Hex())
	if err != nil {
		log.Fatal(err)
	}
	subchainTNT20TokenAddress = address
	auth = subchainSelectAccount(subchainClient, 1)
	address, tx, _, err = ct.DeployMockTNT721(auth, subchainClient)
	fmt.Println("tnt721", address)
	fmt.Println(tx.Hash().Hex())
	if err != nil {
		log.Fatal(err)
	}
	subchainTNT721TokenAddress = address
	auth = subchainSelectAccount(subchainClient, 1)
	address, tx, _, err = ct.DeployMockTNT1155(auth, subchainClient)
	fmt.Println("tnt1155", address)
	fmt.Println(tx.Hash().Hex())
	if err != nil {
		log.Fatal(err)
	}
	subchainTNT1155TokenAddress = address
}
