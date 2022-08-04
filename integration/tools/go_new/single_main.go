package main

const CHAIN_ID_OFFSET int64 = 360

func main() {
	//oneAcoountStake(1)
	claimStake()
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
