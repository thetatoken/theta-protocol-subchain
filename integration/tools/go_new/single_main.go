package main

import "math/big"

// rg "chainRegistrarOnMainchain" // for demo
// func main1() {
// 	client, err := ethclient.Dial("http://localhost:18888/rpc")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	instanceChainRegistrar, err := ct.NewChainRegistrarOnMainchain(registerOnMainchainAddress, client)
// 	if err != nil {
// 		log.Fatal("hhh", err)
// 	}
// 	tx3, err := instanceChainRegistrar.GetStakeSnapshot(nil, big.NewInt(360777), big.NewInt(463))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(tx3)
// }

// func mint() {

// 	// oneAccountRegister()
// 	// oneAcoountStake()
// 	client, _ := ethclient.Dial("http://localhost:19888/rpc")

// 	auth := subchainSelectAccount(client, 1)
// 	instance, _ := ct.NewTestSubchainTNT20TokenBank(subchaintnt20TokenBankAddress, client)
// 	tx, err := instance.MintVouchers(auth, "366/20/"+"0x7f1C87Bd3a22159b8a2E5D195B1a3283D10ea895", "TNT20Name", "vTNT20Symbol", uint8(1), accountList[1].fromAddress, big.NewInt(20), big.NewInt(139), big.NewInt(1))
// 	fmt.Println("-------------")
// 	fmt.Println(tx.Hash().Hex(), err)
// 	height, _ := client.BlockNumber(context.Background())
// 	fmt.Println(height)
// }

// // const RawABI = `
// // [
// // 	{
// // 		"inputs": [
// // 			{
// // 				"components": [
// // 					{
// // 						"internalType": "address",
// // 						"name": "validator",
// // 						"type": "address"
// // 					},
// // 					{
// // 						"internalType": "uint256",
// // 						"name": "shareAmount",
// // 						"type": "uint256"
// // 					}
// // 				],
// // 				"internalType": "struct TestRegister.ValidatorAddrSharePair[]",
// // 				"name": "a",
// // 				"type": "tuple[]"
// // 			}
// // 		],
// // 		"stateMutability": "nonpayable",
// // 		"type": "constructor"
// // 	}
// // ]`

// // type hhh struct {
// // 	Validator   common.Address
// // 	ShareAmount *big.Int
// // }
// func subchainBurnTNT20() {
// 	client, err := ethclient.Dial("http://localhost:19888/rpc")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	authUser := subchainSelectAccount(client, 1)
// 	subchainTNT20VoucherAddress := common.HexToAddress("0x7D7e270b7E279C94b265A535CdbC00Eb62E6e68f")
// 	testSubchainTNT20TokenBankInstance, err := ct.NewTestSubchainTNT20TokenBank(subchaintnt20TokenBankAddress, client)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	subchainTNT20VoucherAddressInstance, _ := ct.NewTNT20VoucherContract(subchainTNT20VoucherAddress, client)
// 	// tx, err := subchainTNT20VoucherAddressInstance.BalanceOf(nil, accountList[1].fromAddress)
// 	//await TNT20Token.approve(TNT20TokenBank1.address, 20, { from: valGuarantor1 })
// 	tx, err := subchainTNT20VoucherAddressInstance.Approve(authUser, subchaintnt20TokenBankAddress, big.NewInt(10))

// 	authUser = subchainSelectAccount(client, 1)
// 	tx, err = testSubchainTNT20TokenBankInstance.BurnVouchers(authUser, subchainTNT20VoucherAddress, accountList[6].fromAddress, big.NewInt(10))
// 	//tx, err := testSubchainTNT20TokenBankInstance.GetDenom(nil, subchainTNT20VoucherAddress)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	//fmt.Println(tx)
// 	fmt.Println(tx.Hash().Hex())
// }

// func subchainBurnTNT721() {
// 	client, err := ethclient.Dial("http://localhost:19888/rpc")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	authUser := subchainSelectAccount(client, 1)
// 	subchainTNT721VoucherAddress := common.HexToAddress("0xd5125d7bB9c4Fb222C522c4b1922cabC631E52D7")
// 	TNT721TokenBankInstance, err := ct.NewTNT721TokenBank(tnt721TokenBankAddress, client)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	subchainTNT721VoucherAddressInstance, _ := ct.NewTNT721VoucherContract(subchainTNT721VoucherAddress, client)
// 	// tx, err := subchainTNT721VoucherAddressInstance.BalanceOf(nil, accountList[1].fromAddress)
// 	//await TNT721Token.approve(TNT721TokenBank1.address, 20, { from: valGuarantor1 })
// 	tx, err := subchainTNT721VoucherAddressInstance.Approve(authUser, tnt721TokenBankAddress, big.NewInt(1))
// 	fmt.Println("721 apporve: ", tx.Hash().Hex())
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	authUser = subchainSelectAccount(client, 1)
// 	tx, err = TNT721TokenBankInstance.BurnVouchers(authUser, subchainTNT721VoucherAddress, accountList[6].fromAddress, big.NewInt(1))
// 	//tx, err := testSubchainTNT721TokenBankInstance.GetDenom(nil, subchainTNT721VoucherAddress)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	//fmt.Println(tx)
// 	fmt.Println(tx.Hash().Hex())
// }

// func getSubchainBlockHeight() {
// 	client, err := ethclient.Dial("http://localhost:19888/rpc")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	height, _ := client.BlockNumber(context.Background())
// 	fmt.Println(big.NewInt(int64(height)))
// }
// func mainchainBurnTNT20() {

// }
func main() {
	mainchainTNT721Lock(big.NewInt(33))
	subchainTNT721Burn(big.NewInt(33))

}
