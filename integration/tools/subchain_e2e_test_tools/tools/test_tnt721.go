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
	mainchainTNT721Name, _ := tnt721VoucherContract.Name(nil)
	mainchainTNT721Symbol, _ := tnt721VoucherContract.Symbol(nil)
	mainchainTNT721TokenURI, _ := tnt721VoucherContract.TokenURI(nil, tokenID)
	tnt721TokenBankContract, err := ct.NewTNT721TokenBank(tnt721TokenBankAddress, mainchainClient)
	if err != nil {
		log.Fatal(err)
	}
	authAccount0 := mainchainSelectAccount(mainchainClient, 0)
	_, err = tnt721VoucherContract.Mint(authAccount0, sender, tokenID, tokenID.String())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Mainchain TNT721 contract address: %v, Name:%v, Symblol:%v, TokenUri:%v\n", tnt721ContractAddress, mainchainTNT721Name, mainchainTNT721Symbol, mainchainTNT721TokenURI)
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
	fmt.Printf("Mainchain NFT sender         : %v, tokenID: %v\n", sender, tokenID)
	fmt.Printf("Subchain NFT Voucher receiver: %v, tokenID: %v\n\n", receiver, tokenID)

	authUser = mainchainSelectAccount(mainchainClient, 1)
	authUser.Value.Set(crossChainFee)
	lockTx, err := tnt721TokenBankContract.LockTokens(authUser, subchainID, tnt721VoucherContractAddress, receiver, tokenID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("TNT721 Token Lock tx hash (Mainchain): %v\n", lockTx.Hash().Hex())
	fmt.Printf("Transfering TNT721 NFT (tokenID: %v) from the Mainchain to Subchain %v...\n\n", tokenID, subchainID)

	fmt.Printf("Start transfer, timestamp      : %v\n", time.Now())
	receipt, err := mainchainClient.TransactionReceipt(context.Background(), lockTx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	if receipt.Status != 1 {
		fmt.Printf("lock error\n")
		return
	}
	fmt.Printf("Token lock confirmed, timestamp: %v\n", time.Now())

	fromHeight, _ := subchainClient.BlockNumber(context.Background())
	var subchainVoucherAddress common.Address
	for {
		time.Sleep(1 * time.Second)
		toHeight, _ := subchainClient.BlockNumber(context.Background())
		result := getSubchainTNT721VoucherMintLogs(int(fromHeight), int(toHeight), subchainTNT721TokenBankAddress, receiver)
		if result != nil {
			subchainVoucherAddress = *result
			break
		}
		fmt.Printf("waiting for cross-chain transfer completion (scanning Subchain from height %v to %v)...\n", fromHeight, toHeight)
	}
	fmt.Printf("End transfer, timestamp        : %v\n", time.Now())
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
	subchainTNT721Name, _ := subchainVoucherContract.Name(nil)
	subchainTNT721Symbol, _ := subchainVoucherContract.Symbol(nil)
	subchainTNT721TokenURI, _ := subchainVoucherContract.TokenURI(nil, tokenID)

	fmt.Printf("Subchain TNT721 Voucher contract address: %v, Name:%v, Symblol:%v, TokenUri:%v\n", subchainVoucherAddress, subchainTNT721Name, subchainTNT721Symbol, subchainTNT721TokenURI)
	fmt.Printf("Mainchain NFT owner          : %v, tokenID: %v (Note: the mainchain NFT owner is the TNT721TokenBank contract since the NFT has been locked)\n", mainchainNFTOwner, tokenID)
	fmt.Printf("Subchain NFT Voucher owner   : %v, tokenID: %v\n\n", subchainNFTOwner, tokenID)
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
	subchainTNT721Name, _ := subchainTNT721Voucher.Name(nil)
	subchainTNT721Symbol, _ := subchainTNT721Voucher.Symbol(nil)
	subchainTNT721TokenURI, _ := subchainTNT721Voucher.TokenURI(nil, tokenID)
	fmt.Printf("Subchain TNT721 Voucher contract address: %v, Name:%v, Symblol:%v, TokenUri:%v\n", subchainTNT721VoucherAddress, subchainTNT721Name, subchainTNT721Symbol, subchainTNT721TokenURI)

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

	fmt.Printf("TNT721 Token Voucher burn tx hash (Subchain): %v\n", burnTx.Hash().Hex())
	fmt.Printf("Burn TNT721 NFT Voucher (tokenID: %v) on Subchain %v to recover the authentic token on the Mainchain...\n\n", tokenID, subchainID)

	fmt.Printf("Start transfer, timestamp        : %v\n", time.Now())
	receipt, err := subchainClient.TransactionReceipt(context.Background(), burnTx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	if receipt.Status != 1 {
		fmt.Println("burn error")
		return
	}
	fmt.Printf("Voucher burn confirmed, timestamp: %v\n", time.Now())

	for {
		time.Sleep(1 * time.Second)
		updatedOwner, _ := mainchainTNT721Contract.OwnerOf(nil, tokenID)
		if updatedOwner != mainchainTNT721Owner {
			break
		}
		fmt.Printf("waiting for cross-chain transfer completion...\n")
	}
	fmt.Printf("End transfer, timestamp          : %v\n", time.Now())
	fmt.Printf("Cross-chain transfer completed.\n\n")

	mainchainNFTOwner, _ := mainchainTNT721Contract.OwnerOf(nil, tokenID)
	if mainchainNFTOwner != receiver {
		log.Fatalf("mainchain token owner and receiver mismatch: %v vs. %v\n", mainchainNFTOwner, receiver)
	}
	mainchainTNT721Name, _ := mainchainTNT721Contract.Name(nil)
	mainchainTNT721Symbol, _ := mainchainTNT721Contract.Symbol(nil)
	mainchainTNT721TokenURI, _ := mainchainTNT721Contract.TokenURI(nil, tokenID)
	fmt.Printf("Mainchain TNT721 Voucher contract address: %v, Name:%v, Symblol:%v, TokenUri:%v\n", mainchainTNT721ContractAddress, mainchainTNT721Name, mainchainTNT721Symbol, mainchainTNT721TokenURI)
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

	//subchainTNT721VoucherAddress := common.HexToAddress("0xd5125d7bB9c4Fb222C522c4b1922cabC631E52D7") // FIXME: should instantiate a mock TNT721 instead of using the Voucher contract (which causes confusion)
	subchainTNT721VoucherAddress := subchainTNT721TokenAddress
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
	subchainTNT721Name, _ := subchainTNT721VoucherInstance.Name(nil)
	subchainTNT721Symbol, _ := subchainTNT721VoucherInstance.Symbol(nil)
	subchainTNT721TokenURI, _ := subchainTNT721VoucherInstance.TokenURI(nil, tokenID)
	fmt.Printf("Subchain TNT721 Voucher contract address: %v, Name:%v, Symblol:%v, TokenUri:%v\n", subchainTNT721VoucherAddress, subchainTNT721Name, subchainTNT721Symbol, subchainTNT721TokenURI)

	fmt.Printf("Subchain NFT sender           : %v, tokenID: %v\n", sender, tokenID)
	fmt.Printf("Mainchain NFT Voucher receiver: %v, tokenID: %v\n\n", receiver, tokenID)

	auth = subchainSelectAccount(subchainClient, 1)
	lockTx, err := subchainTNT721TokenBankinstance.LockTokens(auth, big.NewInt(366), subchainTNT721VoucherAddress, accountList[6].fromAddress, tokenID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("TNT721 Token Lock tx hash (Subchain): %v\n", lockTx.Hash().Hex())
	fmt.Printf("Transfering TNT721 NFT (tokenID: %v) from Subchain %v to the Mainchain...\n\n", tokenID, subchainID)

	fmt.Printf("Start transfer, timestamp      : %v\n", time.Now())
	receipt, err := subchainClient.TransactionReceipt(context.Background(), lockTx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	if receipt.Status != 1 {
		fmt.Println("burn error")
		return
	}
	fmt.Printf("Token lock confirmed, timestamp: %v\n", time.Now())

	// secondOnwer, _ := subchainTNT721VoucherInstance.OwnerOf(nil, tokenID)
	// fmt.Println("The subchain owner of ", tokenID, "is", secondOnwer)

	fromHeight, _ := mainchainClient.BlockNumber(context.Background())
	var mainchainVoucherAddress common.Address
	for {
		time.Sleep(1 * time.Second)
		toHeight, _ := mainchainClient.BlockNumber(context.Background())
		result := getMainchainTNT721VoucherMintLogs(int(fromHeight), int(toHeight), tnt721TokenBankAddress, accountList[6].fromAddress)
		if result != nil {
			mainchainVoucherAddress = *result
			break
		}
		fmt.Printf("waiting for cross-chain transfer completion (scanning Mainchain from height %v to %v)...\n", fromHeight, toHeight)
	}
	fmt.Printf("End transfer, timestamp        : %v\n", time.Now())
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
	mainchainTNT721Name, _ := mainchainVoucherContract.Name(nil)
	mainchainTNT721Symbol, _ := mainchainVoucherContract.Symbol(nil)
	mainchainTNT721TokenURI, _ := mainchainVoucherContract.TokenURI(nil, tokenID)
	fmt.Printf("Mainchain TNT721 Voucher contract address: %v, Name:%v, Symblol:%v, TokenUri:%v\n", mainchainVoucherAddress, mainchainTNT721Name, mainchainTNT721Symbol, mainchainTNT721TokenURI)
	fmt.Printf("Subchain NFT owner            : %v, tokenID: %v (Note: the subchain NFT owner is the TNT721TokenBank contract since the NFT has been locked)\n", subchainNFTOwner, tokenID)
	fmt.Printf("Mainchain NFT Voucher owner   : %v, tokenID: %v\n\n", mainchainNFTOwner, tokenID)
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
	mainchainTNT721Name, _ := mainchainTNT721VoucherContract.Name(nil)
	mainchainTNT721Symbol, _ := mainchainTNT721VoucherContract.Symbol(nil)
	mainchainTNT721TokenURI, _ := mainchainTNT721VoucherContract.TokenURI(nil, tokenID)
	fmt.Printf("Mainchain TNT721 Voucher contract address: %v, Name:%v, Symblol:%v, TokenUri:%v\n", mainchainTNT721VoucherAddr, mainchainTNT721Name, mainchainTNT721Symbol, mainchainTNT721TokenURI)

	mainchainNFTOwner, _ := mainchainTNT721VoucherContract.OwnerOf(nil, tokenID)
	if mainchainNFTOwner != sender {
		log.Fatalf("mainchain token owner and sender mismatch: %v vs. %v\n", mainchainNFTOwner, sender)
	}
	fmt.Printf("Mainchain NFT Voucher sender: %v, tokenID: %v\n", sender, tokenID)
	fmt.Printf("Subchain NFT receiver       : %v, tokenID: %v\n\n", receiver, tokenID)

	//subchainTNT721ContractAddress := common.HexToAddress("0xd5125d7bB9c4Fb222C522c4b1922cabC631E52D7")
	subchainTNT721ContractAddress := subchainTNT721TokenAddress
	subchainTNT721Contract, _ := ct.NewTNT721VoucherContract(subchainTNT721ContractAddress, subchainClient)
	subchainTNT721Owner, _ := subchainTNT721Contract.OwnerOf(nil, tokenID)

	auth = mainchainSelectAccount(mainchainClient, 6)
	auth.Value.Set(crossChainFee)
	burnTx, err := mainchainTNT721TokenBank.BurnVouchers(auth, mainchainTNT721VoucherAddr, accountList[1].fromAddress, tokenID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("TNT721 Token Voucher burn tx hash (Mainchain): %v\n", burnTx.Hash().Hex())
	fmt.Printf("Burn TNT721 NFT Voucher (tokenID: %v) on the Mainchain to recover the authentic token on Subchain %v...\n\n", tokenID, subchainID)

	fmt.Printf("Start transfer, timestamp        : %v\n", time.Now())
	receipt, err := mainchainClient.TransactionReceipt(context.Background(), burnTx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	if receipt.Status != 1 {
		fmt.Println("burn error")
		return
	}
	fmt.Printf("Voucher burn confirmed, timestamp: %v\n", time.Now())

	for {
		time.Sleep(1 * time.Second)
		updatedOwner, _ := subchainTNT721Contract.OwnerOf(nil, tokenID)
		if updatedOwner != subchainTNT721Owner {
			break
		}
		fmt.Printf("waiting for cross-chain transfer completion...\n")
	}
	fmt.Printf("End transfer, timestamp          : %v\n", time.Now())
	fmt.Printf("Cross-chain transfer completed.\n\n")

	subchainNFTOwner, _ := subchainTNT721Contract.OwnerOf(nil, tokenID)
	if subchainNFTOwner != receiver {
		log.Fatalf("subchain token owner and receiver mismatch: %v vs. %v\n", mainchainNFTOwner, receiver)
	}
	subchainTNT721Name, _ := subchainTNT721Contract.Name(nil)
	subchainTNT721Symbol, _ := subchainTNT721Contract.Symbol(nil)
	subchainTNT721TokenURI, _ := subchainTNT721Contract.TokenURI(nil, tokenID)
	fmt.Printf("Subchain TNT721 Voucher contract address: %v, Name:%v, Symblol:%v, TokenUri:%v\n", subchainTNT721ContractAddress, subchainTNT721Name, subchainTNT721Symbol, subchainTNT721TokenURI)
	fmt.Printf("Subchain NFT owner: %v, tokenID: %v\n\n", subchainNFTOwner, tokenID)
}

func QueryTNT721(chainID int64, contractAddress string, tokenID *big.Int) {
	mainchainClient, err := ethclient.Dial("http://localhost:18888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	subchainClient, err := ethclient.Dial("http://localhost:19888/rpc")
	if err != nil {
		log.Fatal(err)
	}

	tnt721ContractAddress := common.HexToAddress(contractAddress) // FIXME: should instantiate a mock TNT721 instead of using the Voucher contract (which causes confusion)
	var instaceTNT721Contract *ct.TNT721VoucherContract
	if chainID == 366 {
		fmt.Printf("Preparing for TNT721 mainchain query...\n")
		instaceTNT721Contract, err = ct.NewTNT721VoucherContract(tnt721ContractAddress, mainchainClient)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Printf("Preparing for TNT721 subchain query...\n")
		instaceTNT721Contract, err = ct.NewTNT721VoucherContract(tnt721ContractAddress, subchainClient)
		if err != nil {
			log.Fatal(err)
		}
	}
	owner, _ := instaceTNT721Contract.OwnerOf(nil, tokenID)
	fmt.Println("TokenID", tokenID, " TNT721 owner in ", contractAddress, " is ", owner)

}

// func QueryTNT721(tokenID *big.Int) {
// 	mainchainClient, err := ethclient.Dial("http://localhost:18888/rpc")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	subchainClient, err := ethclient.Dial("http://localhost:19888/rpc")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	tnt721ContractAddress := tnt721VoucherContractAddress // FIXME: should instantiate a mock TNT721 instead of using the Voucher contract (which causes confusion)
// 	tnt721VoucherContract, _ := ct.NewTNT721VoucherContract(tnt721ContractAddress, mainchainClient)
// 	mainchainNFTOwner, _ := tnt721VoucherContract.OwnerOf(nil, tokenID)
// 	subchainVoucherAddress := common.HexToAddress("0xd5125d7bB9c4Fb222C522c4b1922cabC631E52D7")
// 	subchainVoucherContract, _ := ct.NewTNT721VoucherContract(subchainVoucherAddress, subchainClient)
// 	subchainNFTOwner, _ := subchainVoucherContract.OwnerOf(nil, tokenID)

// 	fmt.Printf("Subchain TNT721 Voucher contract address: %v\n", subchainVoucherAddress)
// 	fmt.Printf("Mainchain NFT owner        : %v, tokenID: %v\n", mainchainNFTOwner, tokenID)
// 	fmt.Printf("Subchain NFT Voucher owner : %v, tokenID: %v\n\n", subchainNFTOwner, tokenID)
// }
