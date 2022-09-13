package tools

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/thetasubchain/eth/ethclient"
	ct "github.com/thetatoken/thetasubchain/interchain/contracts/accessors"
)

func MainchainTNT1155Lock() {
	user := accountList[1].fromAddress
	mainchainClient, err := ethclient.Dial("http://localhost:18888/rpc")
	subchainClient, err := ethclient.Dial("http://localhost:19888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	mainchainTNT1155ContractAddress := common.HexToAddress("0x47eb28D8139A188C5686EedE1E9D8EDE3Afdd543")
	instanceTNT1155VoucherContract, err := ct.NewMockTNT1155(mainchainTNT1155ContractAddress, mainchainClient)
	if err != nil {
		log.Fatal(err)
	}
	instanceTNT1155TokenBank, err := ct.NewTNT1155TokenBank(mainchainTNT1155TokenBankAddress, mainchainClient)
	if err != nil {
		log.Fatal(err)
	}
	authAccount0 := mainchainSelectAccount(mainchainClient, 0)
	tx, err := instanceTNT1155VoucherContract.Mint(authAccount0, user, big.NewInt(1), big.NewInt(1), []byte(""))
	if err != nil {
		log.Fatal(err)
	}
	authAccount1 := mainchainSelectAccount(mainchainClient, 1)
	tx, err = instanceTNT1155VoucherContract.SetApprovalForAll(authAccount1, mainchainTNT1155TokenBankAddress, true)
	if err != nil {
		log.Fatal(err)
	}
	authAccount1 = mainchainSelectAccount(mainchainClient, 1)
	tx, err = instanceTNT1155TokenBank.LockTokens(authAccount1, subchainID, mainchainTNT1155ContractAddress, user, big.NewInt(1), big.NewInt(1), []byte(""), []byte(""))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tx.Hash().Hex())
	fromHeight, _ := subchainClient.BlockNumber(context.Background())
	receipt, err := mainchainClient.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	if receipt.Status != 1 {
		fmt.Println("burn error")
		return
	}

	// secondOnwer, _ := subchainTNT1155VoucherInstance.OwnerOf(nil, tokenID)
	// fmt.Println("The subchain owner of ", tokenID, "is", secondOnwer)

	var mainchainVoucherAddress common.Address
	for {
		time.Sleep(6 * time.Second)
		toHeight, _ := subchainClient.BlockNumber(context.Background())
		result := getSubchainTNT1155VoucherMintLogs(int(fromHeight), int(toHeight), subchainTNT1155TokenBankAddress, accountList[1].fromAddress)
		if result != nil {
			mainchainVoucherAddress = *result
			break
		}
		//fmt.Printf("waiting for cross-chain transfer completion (scanning Mainchain from height %v to %v)...\n", fromHeight, toHeight)
	}
	fmt.Println(mainchainVoucherAddress, "Cross-chain transfer completed.\n\n")
}

func SubchainTNT1155Burn(burnAmount *big.Int) {
	mainchainClient, err := ethclient.Dial("http://localhost:18888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	subchainClient, err := ethclient.Dial("http://localhost:19888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Preparing for TNT1155 cross-chain transfer...\n")

	sender := accountList[1].fromAddress
	receiver := accountList[1].fromAddress

	subchainTNT1155TokenBank, _ := ct.NewTNT1155TokenBank(subchainTNT1155TokenBankAddress, subchainClient)

	subchainTNT1155VoucherAddress := common.HexToAddress("0x0ede92cac9161f6c397a604de508dcd1e6f43e61")
	subchainTNT1155VoucherContract, _ := ct.NewTNT1155VoucherContract(subchainTNT1155VoucherAddress, subchainClient)
	senderSubchainTNT1155VoucherBalance, _ := subchainTNT1155VoucherContract.BalanceOf(nil, sender, big.NewInt(1))

	mainchainTNT1155ContractAddress := common.HexToAddress("")
	mainchainTNT1155Contract, _ := ct.NewTNT1155VoucherContract(mainchainTNT1155ContractAddress, mainchainClient)
	receiverMainchainTNT1155TokenBalance, _ := mainchainTNT1155Contract.BalanceOf(nil, receiver, big.NewInt(1))

	fmt.Printf("Subchain TNT1155 Voucher address: %v\n", subchainTNT1155VoucherAddress)
	fmt.Printf("Subchain sender   : %v, TNT1155 Voucher balance on Subchain: %v\n", sender, senderSubchainTNT1155VoucherBalance)
	fmt.Printf("Mainchain receiver: %v, TNT1155 Token balance on Mainchin  : %v\n\n", receiver, receiverMainchainTNT1155TokenBalance)

	authUser := subchainSelectAccount(subchainClient, 1)
	subchainTNT1155VoucherContract.SetApprovalForAll(authUser, subchainTNT1155TokenBankAddress, true)

	authUser = subchainSelectAccount(subchainClient, 1)
	burnTx, err := subchainTNT1155TokenBank.BurnVouchers(authUser, subchainTNT1155VoucherAddress, accountList[1].fromAddress, big.NewInt(1), big.NewInt(1), []byte(""))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("TNT1155 Voucher Burn tx hash (Subchain): %v\n", burnTx.Hash().Hex())
	fmt.Printf("Burn %v TNT1155 Vouchers (Wei) on Subchain %v to recover the authentic tokens on the Mainchain...\n\n", burnAmount, subchainID)

	fmt.Printf("Start transfer, timestamp        : %v\n", time.Now())
	receipt, err := subchainClient.TransactionReceipt(context.Background(), burnTx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	if receipt.Status != 1 {
		log.Fatal("lock error")
	}
	fmt.Printf("Voucher burn confirmed, timestamp: %v\n", time.Now())

	for {
		time.Sleep(1 * time.Second)
		updatedBalance, _ := mainchainTNT1155Contract.BalanceOf(nil, receiver, big.NewInt(1))
		if receiverMainchainTNT1155TokenBalance.Cmp(updatedBalance) != 0 {
			break
		}
		fmt.Printf("waiting for cross-chain transfer completion...\n")
	}
	fmt.Printf("End transfer, timestamp          : %v\n", time.Now())
	fmt.Printf("Cross-chain transfer completed.\n\n")

	// senderSubchainTNT1155VoucherBalance, _ = subchainTNT1155VoucherContract.BalanceOf(nil, sender, big.NewInt(1))
	// receiverMainchainTNT1155TokenBalance, _ = mainchainTNT1155Contract.BalanceOf(nil, receiver, big.NewInt(1))

	// fmt.Printf("Mainchain TNT1155 token contract address: %v\n", mainchainTNT1155ContractAddress)
	// fmt.Printf("Subchain sender   : %v, TNT1155 Voucher balance on Subchain: %v\n", sender, senderSubchainTNT1155VoucherBalance)
	// fmt.Printf("Mainchain receiver: %v, TNT1155 Token balance on Mainchain  : %v\n\n", receiver, receiverMainchainTNT1155TokenBalance)
}

func SubchainTNT1155Lock(tokenID *big.Int) {
	mainchainClient, err := ethclient.Dial("http://localhost:18888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	subchainClient, err := ethclient.Dial("http://localhost:19888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Preparing for TNT1155 cross-chain transfer...\n")

	//sender := accountList[1].fromAddress
	receiver := accountList[6].fromAddress

	subchainTNT1155VoucherAddress := common.HexToAddress("0x0ede92cac9161f6c397a604de508dcd1e6f43e61") // FIXME: should instantiate a mock TNT1155 instead of using the Voucher contract (which causes confusion)
	subchainTNT1155VoucherInstance, _ := ct.NewTNT1155VoucherContract(subchainTNT1155VoucherAddress, subchainClient)
	subchainTNT1155TokenBankinstance, err := ct.NewTNT1155TokenBank(subchainTNT1155TokenBankAddress, subchainClient)
	if err != nil {
		log.Fatal(err)
	}
	auth := subchainSelectAccount(subchainClient, 1)
	_, err = subchainTNT1155VoucherInstance.SetApprovalForAll(auth, subchainTNT1155TokenBankAddress, true)
	if err != nil {
		log.Fatal(err)
	}

	// subchainNFTOwner, _ := subchainTNT1155VoucherInstance.BalanceOf(nil, accountList[1].fromAddress, big.NewInt(1))
	// if subchainNFTOwner.Cmp(big.NewInt(1)) != 0 {
	// 	log.Fatalf("mainchain token owner and sender mismatch: %v vs. %v\n", subchainNFTOwner, sender)
	// }
	// fmt.Printf("Subchain NFT sender   : %v, tokenID: %v\n", sender, tokenID)
	// fmt.Printf("Mainchain NFT receiver: %v, tokenID: %v\n\n", receiver, tokenID)

	auth = subchainSelectAccount(subchainClient, 1)
	lockTx, err := subchainTNT1155TokenBankinstance.LockTokens(auth, big.NewInt(366), subchainTNT1155VoucherAddress, receiver, big.NewInt(1), big.NewInt(1), []byte(""), []byte(""))
	if err != nil {
		log.Fatal(err)
	}
	tx1, _ := subchainTNT1155TokenBankinstance.TotalLockedAmounts(nil, big.NewInt(366), common.HexToAddress("0x0ede92cac9161f6c397a604de508dcd1e6f43e61"), big.NewInt(1))
	fmt.Println(tx1)
	fmt.Printf("TNT1155 Token Lock tx hash (Subchain): %v\n", lockTx.Hash().Hex())
	fmt.Printf("Transfering TNT1155 NFT (tokenID: %v) from Subchain %v to the Mainchain...\n\n", tokenID, subchainID)
	fromHeight, _ := mainchainClient.BlockNumber(context.Background())
	receipt, err := subchainClient.TransactionReceipt(context.Background(), lockTx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	if receipt.Status != 1 {
		fmt.Println("burn error")
		return
	}

	// secondOnwer, _ := subchainTNT1155VoucherInstance.OwnerOf(nil, tokenID)
	// fmt.Println("The subchain owner of ", tokenID, "is", secondOnwer)

	var mainchainVoucherAddress common.Address
	for {
		time.Sleep(2 * time.Second)
		toHeight, _ := mainchainClient.BlockNumber(context.Background())
		result := getMainchainTNT1155VoucherMintLogs(int(fromHeight), int(toHeight), mainchainTNT1155TokenBankAddress, accountList[6].fromAddress)
		if result != nil {
			mainchainVoucherAddress = *result
			break
		}
		//fmt.Printf("waiting for cross-chain transfer completion (scanning Mainchain from height %v to %v)...\n", fromHeight, toHeight)
	}
	fmt.Println(mainchainVoucherAddress, "Cross-chain transfer completed.\n\n")

	// subchainNFTOwner, _ = subchainTNT1155VoucherInstance.OwnerOf(nil, tokenID)
	// if subchainNFTOwner != subchainTNT1155TokenBankAddress {
	// 	log.Fatalf("subchain token owner should be the TNT1155TokenBank contract: %v vs. %v\n", subchainNFTOwner, subchainTNT1155TokenBankAddress)
	// }

	// mainchainVoucherContract, _ := ct.NewTNT1155VoucherContract(mainchainVoucherAddress, mainchainClient)
	// mainchainNFTOwner, _ := mainchainVoucherContract.OwnerOf(nil, tokenID)
	// if mainchainNFTOwner != receiver {
	// 	log.Fatalf("mainchain token owner and receiver mismatch: %v vs. %v\n", mainchainNFTOwner, receiver)
	// }

	// fmt.Printf("Mainchain TNT1155 Voucher contract address: %v\n", mainchainVoucherAddress)
	// fmt.Printf("Subchain NFT owner : %v, tokenID: %v (Note: the subchain NFT owner is the TNT1155TokenBank contract since the NFT has been locked)\n", subchainNFTOwner, tokenID)
	// fmt.Printf("Mainchain NFT owner: %v, tokenID: %v\n\n", mainchainNFTOwner, tokenID)
}

func MainchainTNT1155Burn(burnAmount *big.Int) {
	mainchainClient, err := ethclient.Dial("http://localhost:18888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	subchainClient, err := ethclient.Dial("http://localhost:19888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Preparing for TNT1155 cross-chain transfer...\n")

	sender := accountList[6].fromAddress
	receiver := accountList[1].fromAddress

	mainchainTNT1155TokenBankInstance, _ := ct.NewTNT1155TokenBank(mainchainTNT1155TokenBankAddress, mainchainClient)

	mainchainTNT1155VoucherAddress := common.HexToAddress("0xCFbb5680AAbcaA297bC1b9233677098308BadB82")
	mainchainTNT1155VoucherContract, _ := ct.NewTNT1155VoucherContract(mainchainTNT1155VoucherAddress, mainchainClient)
	senderMainchainTNT1155VoucherBalance, _ := mainchainTNT1155VoucherContract.BalanceOf(nil, sender, big.NewInt(1))

	subchainTNT1155TokenAddress := common.HexToAddress("0x0ede92cac9161f6c397a604de508dcd1e6f43e61") // FIXME: should instantiate a mock TNT1155 instead of using the Voucher contract (which causes confusion)
	subchainTNT1155TokenContract, _ := ct.NewTNT1155VoucherContract(subchainTNT1155TokenAddress, subchainClient)
	receiverSubchainTNT1155TokenBalance, _ := subchainTNT1155TokenContract.BalanceOf(nil, receiver, big.NewInt(1))

	fmt.Printf("Mainchain TNT1155 Voucher address: %v\n", mainchainTNT1155VoucherAddress)
	fmt.Printf("Mainchain sender : %v, TNT1155 Voucher balance on Mainchain: %v\n", sender, senderMainchainTNT1155VoucherBalance)
	fmt.Printf("Subchain receiver: %v, TNT1155 Token balance on Subchain   : %v\n\n", receiver, receiverSubchainTNT1155TokenBalance)

	authUser := mainchainSelectAccount(mainchainClient, 6)
	mainchainTNT1155VoucherContract.SetApprovalForAll(authUser, mainchainTNT1155TokenBankAddress, true)

	authUser = mainchainSelectAccount(mainchainClient, 6)
	burnTx, err := mainchainTNT1155TokenBankInstance.BurnVouchers(authUser, mainchainTNT1155VoucherAddress, accountList[1].fromAddress, big.NewInt(1), big.NewInt(1), []byte(""))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("TNT1155 Voucher Burn tx hash (Mainchain): %v\n", burnTx.Hash().Hex())
	fmt.Printf("Burn %v TNT1155 Vouchers (Wei) on the Mainchain to recover the authentic tokens on Subchain %v...\n\n", burnAmount, subchainID)

	fmt.Printf("Start transfer, timestamp        : %v\n", time.Now())
	receipt, err := mainchainClient.TransactionReceipt(context.Background(), burnTx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	if receipt.Status != 1 {
		log.Fatal("lock error")
		return
	}
	fmt.Printf("Voucher burn confirmed, timestamp: %v\n", time.Now())

	for {
		time.Sleep(1 * time.Second)
		updatedBalance, _ := subchainTNT1155TokenContract.BalanceOf(nil, receiver, big.NewInt(1))
		if receiverSubchainTNT1155TokenBalance.Cmp(updatedBalance) != 0 {
			break
		}
		fmt.Printf("waiting for cross-chain transfer completion...\n")
	}
	fmt.Printf("End transfer, timestamp          : %v\n", time.Now())
	fmt.Printf("Cross-chain transfer completed.\n\n")

	// senderMainchainTNT1155VoucherBalance, _ = mainchainTNT1155VoucherContract.BalanceOf(nil, sender)
	// receiverSubchainTNT1155TokenBalance, _ = subchainTNT1155TokenContract.BalanceOf(nil, receiver)

	// fmt.Printf("Subchain TNT1155 token contract address: %v\n", subchainTNT1155TokenAddress)
	// fmt.Printf("Mainchain sender : %v, TNT1155 Voucher balance on Mainchain: %v\n", sender, senderMainchainTNT1155VoucherBalance)
	// fmt.Printf("Subchain receiver: %v, TNT1155 Token balance on Subchain   : %v\n\n", receiver, receiverSubchainTNT1155TokenBalance)
}
