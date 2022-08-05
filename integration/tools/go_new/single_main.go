package main

import (
	"fmt"
	"log"

	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/thetasubchain/eth/ethclient"
	ct "github.com/thetatoken/thetasubchain/integration/tools/go_new/accessors"
)

const CHAIN_ID_OFFSET int64 = 360

func config() []string {
	wthetaAddress = common.HexToAddress("0x7d73424a8256C0b2BA245e5d5a3De8820E45F390")
	registerOnMainchainAddress = common.HexToAddress("0x08425D9Df219f93d5763c3e85204cb5B4cE33aAa")
	governanceTokenAddress = common.HexToAddress("0x6E05f58eEddA592f34DD9105b1827f252c509De0")
	tnt20VoucherContractAddress = common.HexToAddress("0x4fb87c52Bb6D194f78cd4896E3e574028fedBAB9")
	tnt20TokenBankAddress = common.HexToAddress("0x2Ce636d6240f8955d085a896e12429f8B3c7db26")
	subchaintnt20TokenBankAddress = common.HexToAddress("0x47e9Fbef8C83A1714F1951F142132E6e90F5fa5D")

	tnt721TokenBankAddress = common.HexToAddress("0x47eb28D8139A188C5686EedE1E9D8EDE3Afdd543")
	tnt721VoucherContractAddress = common.HexToAddress("0x52d2878492EF30d625fc54EC52c4dB7f010d471e")
	Subchaintnt721TokenBankAddress = common.HexToAddress("0x8Be503bcdEd90ED42Eff31f56199399B2b0154CA")

	tfuelTokenbankAddress = common.HexToAddress("0x560A0c0CA6B0A67895024dae77442C5fd3DC473e")
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

	// var dec18 = new(big.Int)
	// dec18.SetString("1000000000000000000", 10)
	// //amount := new(big.Int).Mul(dec18, big.NewInt(200000))
	// client, err := ethclient.Dial("http://localhost:18888/rpc")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// nonce, err := client.PendingNonceAt(context.Background(), accountList[9].fromAddress)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// queryStr := fmt.Sprintf(`{"jsonrpc":"2.0","method":"thetacli.Send","params":[{"chain_id":"privatenet", "from":"%v", "to":"%v", "thetawei":"9900000", "tfuelwei":"88000000", "fee":"100000000", "sequence":"%v", "async":true}],"id":1}`, accountList[1].fromAddress, accountList[10].fromAddress, fmt.Sprintf("%d", nonce))
	// var jsonData = []byte(queryStr)

	// request, err := http.NewRequest("POST", "http://localhost:18888/rpc", bytes.NewBuffer(jsonData))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// request.Header.Set("Content-Type", "application/json")

	// client1 := &http.Client{}
	// response, err := client1.Do(request)
	// if err != nil {
	// 	log.Fatalf("response error : %v", err)
	// }
	// defer response.Body.Close()
	// fmt.Println(response)
	//oneAcoountStake(10)
	//oneAcoountStake(12)
	//claimStake()
	//oneAcoountStake(1)
	//mainchainTfuelLock(big.NewInt(33))

	// mainchainTNT721Lock(big.NewInt(33))
	// subchainTNT721Burn(big.NewInt(33))
	client, err := ethclient.Dial("http://localhost:18888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	TfuelTokenBankInstance, err := ct.NewTFuelTokenBank(tfuelTokenbankAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	//auth := subchainSelectAccount(client, 1)
	tx, _ := TfuelTokenBankInstance.Test(nil)
	fmt.Println(tx)
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
