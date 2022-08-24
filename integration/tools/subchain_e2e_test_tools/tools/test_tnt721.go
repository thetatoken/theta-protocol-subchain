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

func MainchainTNT721Lock(tokenID *big.Int) {
	mainchainClient, err := ethclient.Dial("http://localhost:18888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	subchainClient, err := ethclient.Dial("http://localhost:19888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Preparing for TNT721 cross-chain transfer...\n")

	sender := accountList[1].fromAddress
	receiver := accountList[1].fromAddress

	tnt721ContractAddress := tnt721VoucherContractAddress // FIXME: should instantiate a mock TNT721 instead of using the Voucher contract (which causes confusion)
	tnt721VoucherContract, err := ct.NewTNT721VoucherContract(tnt721ContractAddress, mainchainClient)
	if err != nil {
		log.Fatal(err)
	}
	tnt721TokenBankContract, err := ct.NewTNT721TokenBank(tnt721TokenBankAddress, mainchainClient)
	if err != nil {
		log.Fatal(err)
	}
	authAccount0 := mainchainSelectAccount(mainchainClient, 0)
	_, err = tnt721VoucherContract.Mint(authAccount0, sender, tokenID, tokenID.String())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Mainchain TNT721 contract address: %v\n", tnt721ContractAddress)
	// fmt.Printf("Mint TNT721 tx hash (Mainchain): %v\n", tx.Hash().Hex())
	time.Sleep(12 * time.Second) // This is needed otherwise the "Approve tx" below might fail
	// authUser := mainchainSelectAccount(client, 1)
	// tx, err = tnt721VoucherContract.SetApprovalForAll(authUser, tnt721TokenBankAddress, true)
	authUser := mainchainSelectAccount(mainchainClient, 1)
	_, err = tnt721VoucherContract.Approve(authUser, tnt721TokenBankAddress, tokenID)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("Approve TNT721 tx hash (Mainchain): %v\n", tx.Hash().Hex())

	mainchainNFTOwner, _ := tnt721VoucherContract.OwnerOf(nil, tokenID)
	if mainchainNFTOwner != sender {
		log.Fatalf("mainchain token owner and sender mismatch: %v vs. %v\n", mainchainNFTOwner, sender)
	}
	fmt.Printf("Mainchain NFT sender : %v, tokenID: %v\n", sender, tokenID)
	fmt.Printf("Subchain NFT receiver: %v, tokenID: %v\n\n", receiver, tokenID)

	authUser = mainchainSelectAccount(mainchainClient, 1)
	lockTx, err := tnt721TokenBankContract.LockTokens(authUser, subchainID, tnt721VoucherContractAddress, receiver, tokenID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("TNT721 Token Lock tx hash (Mainchain): %v\n", lockTx.Hash().Hex())
	fmt.Printf("Transfering TNT721 NFT (tokenID: %v) from the Mainchain to Subchain %v...\n\n", tokenID, subchainID)

	receipt, err := mainchainClient.TransactionReceipt(context.Background(), lockTx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	if receipt.Status != 1 {
		fmt.Printf("lock error\n")
		return
	}

	fromHeight, _ := subchainClient.BlockNumber(context.Background())
	var subchainVoucherAddress common.Address
	for {
		time.Sleep(6 * time.Second)
		toHeight, _ := subchainClient.BlockNumber(context.Background())
		result := getSubchainTNT721VoucherMintLogs(int(fromHeight), int(toHeight), subchainTNT721TokenBankAddress, receiver)
		if result != nil {
			subchainVoucherAddress = *result
			break
		}
		fmt.Printf("waiting for cross-chain transfer completion (scanning Subchain from height %v to %v)...\n", fromHeight, toHeight)
	}
	fmt.Printf("Cross-chain transfer completed.\n\n")

	mainchainNFTOwner, _ = tnt721VoucherContract.OwnerOf(nil, tokenID)
	if mainchainNFTOwner != tnt721TokenBankAddress {
		log.Fatalf("mainchain token owner should be the TNT721TokenBank contract: %v vs. %v\n", mainchainNFTOwner, tnt721TokenBankAddress)
	}

	subchainVoucherContract, _ := ct.NewTNT721VoucherContract(subchainVoucherAddress, subchainClient)
	subchainNFTOwner, _ := subchainVoucherContract.OwnerOf(nil, tokenID)
	if subchainNFTOwner != receiver {
		log.Fatalf("subchain token owner and receiver mismatch: %v vs. %v\n", subchainNFTOwner, receiver)
	}

	fmt.Printf("Subchain TNT721 Voucher contract address: %v\n", subchainVoucherAddress)
	fmt.Printf("Mainchain NFT owner: %v, tokenID: %v (Note: the mainchain NFT owner is the TNT721TokenBank contract since the NFT has been locked)\n", mainchainNFTOwner, tokenID)
	fmt.Printf("Subchain NFT owner : %v, tokenID: %v\n\n", subchainNFTOwner, tokenID)
}

func SubchainTNT721Burn(tokenID *big.Int) {
	mainchainClient, err := ethclient.Dial("http://localhost:18888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	subchainClient, err := ethclient.Dial("http://localhost:19888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Preparing for TNT721 cross-chain transfer...\n")

	sender := accountList[1].fromAddress
	receiver := accountList[6].fromAddress

	subchainTNT721TokenBank, _ := ct.NewTNT721TokenBank(subchainTNT721TokenBankAddress, subchainClient)

	subchainTNT721VoucherAddress := common.HexToAddress("0xd5125d7bB9c4Fb222C522c4b1922cabC631E52D7")
	subchainTNT721Voucher, _ := ct.NewTNT721VoucherContract(subchainTNT721VoucherAddress, subchainClient)
	authUser := subchainSelectAccount(subchainClient, 1)
	_, err = subchainTNT721Voucher.Approve(authUser, subchainTNT721TokenBankAddress, tokenID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Subchain TNT721 Voucher contract address: %v\n", subchainTNT721VoucherAddress)

	subchainNFTOwner, _ := subchainTNT721Voucher.OwnerOf(nil, tokenID)
	if subchainNFTOwner != sender {
		log.Fatalf("subchain token owner and sender mismatch: %v vs. %v\n", subchainNFTOwner, sender)
	}
	fmt.Printf("Subchain NFT Voucher sender: %v, tokenID: %v\n", sender, tokenID)
	fmt.Printf("Mainchain NFT receiver     : %v, tokenID: %v\n\n", receiver, tokenID)

	mainchainTNT721ContractAddress := tnt721VoucherContractAddress
	mainchainTNT721Contract, _ := ct.NewTNT721VoucherContract(mainchainTNT721ContractAddress, mainchainClient)
	mainchainTNT721Owner, _ := mainchainTNT721Contract.OwnerOf(nil, tokenID)

	authUser = subchainSelectAccount(subchainClient, 1)
	burnTx, err := subchainTNT721TokenBank.BurnVouchers(authUser, subchainTNT721VoucherAddress, receiver, tokenID)
	if err != nil {
		log.Fatal(err)
	}
	receipt, err := subchainClient.TransactionReceipt(context.Background(), burnTx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	if receipt.Status != 1 {
		fmt.Println("burn error")
		return
	}

	fmt.Printf("TNT721 Token Voucher burn tx hash (Subchain): %v\n", burnTx.Hash().Hex())
	fmt.Printf("Burn TNT721 NFT Voucher (tokenID: %v) on Subchain %v to recover the authentic token on the Mainchain...\n\n", tokenID, subchainID)

	for {
		time.Sleep(6 * time.Second)
		updatedOwner, _ := mainchainTNT721Contract.OwnerOf(nil, tokenID)
		if updatedOwner != mainchainTNT721Owner {
			break
		}
		fmt.Printf("waiting for cross-chain transfer completion...\n")
	}
	fmt.Printf("Cross-chain transfer completed.\n\n")

	mainchainNFTOwner, _ := mainchainTNT721Contract.OwnerOf(nil, tokenID)
	if mainchainNFTOwner != receiver {
		log.Fatalf("mainchain token owner and receiver mismatch: %v vs. %v\n", mainchainNFTOwner, receiver)
	}

	fmt.Printf("Mainchain TNT721 Voucher contract address: %v\n", mainchainTNT721ContractAddress)
	fmt.Printf("Mainchain NFT owner: %v, tokenID: %v\n\n", mainchainNFTOwner, tokenID)
}

func SubchainTNT721Lock(tokenID *big.Int) {
	mainchainClient, err := ethclient.Dial("http://localhost:18888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	subchainClient, err := ethclient.Dial("http://localhost:19888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Preparing for TNT721 cross-chain transfer...\n")

	sender := accountList[1].fromAddress
	receiver := accountList[6].fromAddress

	subchainTNT721VoucherAddress := common.HexToAddress("0xd5125d7bB9c4Fb222C522c4b1922cabC631E52D7") // FIXME: should instantiate a mock TNT721 instead of using the Voucher contract (which causes confusion)
	subchainTNT721VoucherInstance, _ := ct.NewTNT721VoucherContract(subchainTNT721VoucherAddress, subchainClient)
	subchainTNT721TokenBankinstance, err := ct.NewTNT721TokenBank(subchainTNT721TokenBankAddress, subchainClient)
	if err != nil {
		log.Fatal(err)
	}
	auth := subchainSelectAccount(subchainClient, 1)
	_, err = subchainTNT721VoucherInstance.Approve(auth, subchainTNT721TokenBankAddress, tokenID)
	if err != nil {
		log.Fatal(err)
	}

	subchainNFTOwner, _ := subchainTNT721VoucherInstance.OwnerOf(nil, tokenID)
	if subchainNFTOwner != sender {
		log.Fatalf("mainchain token owner and sender mismatch: %v vs. %v\n", subchainNFTOwner, sender)
	}
	fmt.Printf("Subchain NFT sender   : %v, tokenID: %v\n", sender, tokenID)
	fmt.Printf("Mainchain NFT receiver: %v, tokenID: %v\n\n", receiver, tokenID)

	auth = subchainSelectAccount(subchainClient, 1)
	lockTx, err := subchainTNT721TokenBankinstance.LockTokens(auth, big.NewInt(366), subchainTNT721VoucherAddress, accountList[6].fromAddress, tokenID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("TNT721 Token Lock tx hash (Subchain): %v\n", lockTx.Hash().Hex())
	fmt.Printf("Transfering TNT721 NFT (tokenID: %v) from Subchain %v to the Mainchain...\n\n", tokenID, subchainID)

	receipt, err := subchainClient.TransactionReceipt(context.Background(), lockTx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	if receipt.Status != 1 {
		fmt.Println("burn error")
		return
	}

	// secondOnwer, _ := subchainTNT721VoucherInstance.OwnerOf(nil, tokenID)
	// fmt.Println("The subchain owner of ", tokenID, "is", secondOnwer)

	fromHeight, _ := mainchainClient.BlockNumber(context.Background())
	var mainchainVoucherAddress common.Address
	for {
		time.Sleep(6 * time.Second)
		toHeight, _ := mainchainClient.BlockNumber(context.Background())
		result := getMainchainTNT721VoucherMintLogs(int(fromHeight), int(toHeight), tnt721TokenBankAddress, accountList[6].fromAddress)
		if result != nil {
			mainchainVoucherAddress = *result
			break
		}
		fmt.Printf("waiting for cross-chain transfer completion (scanning Mainchain from height %v to %v)...\n", fromHeight, toHeight)
	}
	fmt.Printf("Cross-chain transfer completed.\n\n")

	subchainNFTOwner, _ = subchainTNT721VoucherInstance.OwnerOf(nil, tokenID)
	if subchainNFTOwner != subchainTNT721TokenBankAddress {
		log.Fatalf("subchain token owner should be the TNT721TokenBank contract: %v vs. %v\n", subchainNFTOwner, subchainTNT721TokenBankAddress)
	}

	mainchainVoucherContract, _ := ct.NewTNT721VoucherContract(mainchainVoucherAddress, mainchainClient)
	mainchainNFTOwner, _ := mainchainVoucherContract.OwnerOf(nil, tokenID)
	if mainchainNFTOwner != receiver {
		log.Fatalf("mainchain token owner and receiver mismatch: %v vs. %v\n", mainchainNFTOwner, receiver)
	}

	fmt.Printf("Mainchain TNT721 Voucher contract address: %v\n", mainchainVoucherAddress)
	fmt.Printf("Subchain NFT owner : %v, tokenID: %v (Note: the subchain NFT owner is the TNT721TokenBank contract since the NFT has been locked)\n", subchainNFTOwner, tokenID)
	fmt.Printf("Mainchain NFT owner: %v, tokenID: %v\n\n", mainchainNFTOwner, tokenID)
}

func MainchainTNT721Burn(tokenID *big.Int) {
	mainchainClient, err := ethclient.Dial("http://localhost:18888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	subchainClient, err := ethclient.Dial("http://localhost:19888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Preparing for TNT721 cross-chain transfer...\n")

	sender := accountList[6].fromAddress
	receiver := accountList[1].fromAddress

	mainchainTNT721TokenBank, _ := ct.NewTNT721TokenBank(tnt721TokenBankAddress, mainchainClient)

	mainchainTNT721VoucherAddr := common.HexToAddress("0x20B2897c4f95df71a5a8A62Ae812f5843AA92E25")
	mainchainTNT721VoucherContract, _ := ct.NewTNT721VoucherContract(mainchainTNT721VoucherAddr, mainchainClient)
	auth := mainchainSelectAccount(mainchainClient, 6)
	_, err = mainchainTNT721VoucherContract.Approve(auth, tnt721TokenBankAddress, tokenID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Mainchain TNT721 Voucher contract address: %v\n", mainchainTNT721VoucherAddr)

	mainchainNFTOwner, _ := mainchainTNT721VoucherContract.OwnerOf(nil, tokenID)
	if mainchainNFTOwner != sender {
		log.Fatalf("mainchain token owner and sender mismatch: %v vs. %v\n", mainchainNFTOwner, sender)
	}
	fmt.Printf("Mainchain NFT Voucher sender: %v, tokenID: %v\n", sender, tokenID)
	fmt.Printf("Subchain NFT receiver       : %v, tokenID: %v\n\n", receiver, tokenID)

	subchainTNT721ContractAddress := common.HexToAddress("0xd5125d7bB9c4Fb222C522c4b1922cabC631E52D7")
	subchainTNT721Contract, _ := ct.NewTNT721VoucherContract(subchainTNT721ContractAddress, subchainClient)
	subchainTNT721Owner, _ := subchainTNT721Contract.OwnerOf(nil, tokenID)

	//
	auth = mainchainSelectAccount(mainchainClient, 6)
	burnTx, err := mainchainTNT721TokenBank.BurnVouchers(auth, mainchainTNT721VoucherAddr, accountList[1].fromAddress, tokenID)
	if err != nil {
		log.Fatal(err)
	}
	receipt, err := mainchainClient.TransactionReceipt(context.Background(), burnTx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	if receipt.Status != 1 {
		fmt.Println("burn error")
		return
	}

	fmt.Printf("TNT721 Token Voucher burn tx hash (Mainchain): %v\n", burnTx.Hash().Hex())
	fmt.Printf("Burn TNT721 NFT Voucher (tokenID: %v) on the Mainchain to recover the authentic token on Subchain %v...\n\n", tokenID, subchainID)

	for {
		time.Sleep(6 * time.Second)
		updatedOwner, _ := subchainTNT721Contract.OwnerOf(nil, tokenID)
		if updatedOwner != subchainTNT721Owner {
			break
		}
		fmt.Printf("waiting for cross-chain transfer completion...\n")
	}
	fmt.Printf("Cross-chain transfer completed.\n\n")

	subchainNFTOwner, _ := subchainTNT721Contract.OwnerOf(nil, tokenID)
	if subchainNFTOwner != receiver {
		log.Fatalf("subchain token owner and receiver mismatch: %v vs. %v\n", mainchainNFTOwner, receiver)
	}

	fmt.Printf("Subchain TNT721 Voucher contract address: %v\n", subchainTNT721ContractAddress)
	fmt.Printf("Subchain NFT owner: %v, tokenID: %v\n\n", subchainNFTOwner, tokenID)
}
