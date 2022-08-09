package main

import (
	"math/big"

	"github.com/thetatoken/theta/common"
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
	mainchainTfuelLock(big.NewInt(10)) //lock some tfuel tokens on mainchain and transfer to subchain 360777
	 //subchainTfuelBurn(big.NewInt(10))

	// //TNT20
	//mainchainTNT20Lock(big.NewInt(100)) // last address printed in console is subchain voucher contract address
	//subchainTNT20Lock(big.NewInt(10)) //last address printed in console is mainchain voucher contract address
	//mainchainTNT20Burn(big.NewInt(10))
	// subchainTNT20Burn(big.NewInt(10))

	// //TNT721
	// mainchainTNT721Lock(big.NewInt(100)) //last address printed in console is subchain voucher contract address
	// subchainTNT721Burn(big.NewInt(10))
	// subchainTNT721Lock(big.NewInt(10)) //last address printed in console is mainchain voucher contract address
	// mainchainTNT721Burn(big.NewInt(10))

	//oneAcoountStake(12)
	//claimStake()
	//oneAcoountStake(1)
	//mainchainTfuelLock(big.NewInt(33))

	// mainchainTNT721Lock(big.NewInt(33))
	// subchainTNT721Burn(big.NewInt(33))
	// client, err := ethclient.Dial("http://localhost:19888/rpc")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// TfuelTokenBankInstance, err := ct.NewTFuelTokenBank(subchainTfuelTokenBank, client)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// auth := subchainSelectAccount(client, 1)
	// tx, _ := TfuelTokenBankInstance.Id(auth)
	// fmt.Println(tx.Hash().Hex())
	// getMintlog1(1, 9, common.HexToAddress("0x5a443704dd4b594b382c22a083e2bd3090a6fef3"))
	// chainID, err := client.ChainID(context.Background())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(chainID)
	// chainIDStr := "tsub_360777"
	// chainIDWithoutOffset := new(big.Int).Abs(crypto.Keccak256Hash(common.Bytes(chainIDStr)).Big())
	// chainID1 := big.NewInt(1).Add(big.NewInt(CHAIN_ID_OFFSET), chainIDWithoutOffset)
	// //chainID1.SetString("tsub_360777", 16)
	// fmt.Println(hex.EncodeToString(chainID1.Bytes()))
}
