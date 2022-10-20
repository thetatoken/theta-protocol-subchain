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

	mainchainTNT721ContractAddress := common.HexToAddress("0xEd8d61f42dC1E56aE992D333A4992C3796b22A74")
	mainchainTNT721Contract, err := ct.NewMockTNT721(mainchainTNT721ContractAddress, mainchainClient)
	if err != nil {
		log.Fatal(err)
	}
	tnt721TokenBankContract, err := ct.NewTNT721TokenBank(mainchainTNT721TokenBankAddress, mainchainClient)
	if err != nil {
		log.Fatal(err)
	}
	authAccount0 := mainchainSelectAccount(mainchainClient, 0)
	_, err = mainchainTNT721Contract.Mint(authAccount0, sender, tokenID)
	if err != nil {
		log.Fatal(err)
	}

	mainchainTNT721Name, _ := mainchainTNT721Contract.Name(nil)
	mainchainTNT721Symbol, _ := mainchainTNT721Contract.Symbol(nil)
	mainchainTNT721TokenURI, _ := mainchainTNT721Contract.TokenURI(nil, tokenID)
	fmt.Printf("Mainchain TNT721 contract address: %v, Name: %v, Symbol: %v, TokenURI: %v\n", mainchainTNT721ContractAddress, mainchainTNT721Name, mainchainTNT721Symbol, mainchainTNT721TokenURI)
	// fmt.Printf("Mint TNT721 tx hash (Mainchain): %v\n", tx.Hash().Hex())
	time.Sleep(12 * time.Second) // This is needed otherwise the "Approve tx" below might fail
	// authUser := mainchainSelectAccount(client, 1)
	// tx, err = mainchainTNT721Contract.SetApprovalForAll(authUser, mainchainTNT721TokenBankAddress, true)
	authUser := mainchainSelectAccount(mainchainClient, 1)
	_, err = mainchainTNT721Contract.Approve(authUser, mainchainTNT721TokenBankAddress, tokenID)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("Approve TNT721 tx hash (Mainchain): %v\n", tx.Hash().Hex())

	mainchainNFTOwner, _ := mainchainTNT721Contract.OwnerOf(nil, tokenID)
	if mainchainNFTOwner != sender {
		log.Fatalf("mainchain token owner and sender mismatch: %v vs. %v\n", mainchainNFTOwner, sender)
	}
	fmt.Printf("Mainchain NFT sender         : %v, tokenID: %v\n", sender, tokenID)
	fmt.Printf("Subchain NFT Voucher receiver: %v, tokenID: %v\n\n", receiver, tokenID)

	// printTNT721TokenBankLockedAmount(mainchainTNT721TokenBankAddress, mainchainTNT721ContractAddress, subchainID, tokenID, mainchainClient)

	authUser = mainchainSelectAccount(mainchainClient, 1)
	authUser.Value.Set(crossChainFee)
	lockTx, err := tnt721TokenBankContract.LockTokens(authUser, subchainID, mainchainTNT721ContractAddress, receiver, tokenID)
	authUser.Value.Set(common.Big0)
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

	// printTNT721TokenBankLockedAmount(mainchainTNT721TokenBankAddress, mainchainTNT721ContractAddress, subchainID, tokenID, mainchainClient)

	mainchainNFTOwner, _ = mainchainTNT721Contract.OwnerOf(nil, tokenID)
	if mainchainNFTOwner != mainchainTNT721TokenBankAddress {
		log.Fatalf("mainchain token owner should be the TNT721TokenBank contract: %v vs. %v\n", mainchainNFTOwner, mainchainTNT721TokenBankAddress)
	}

	subchainVoucherContract, _ := ct.NewTNT721VoucherContract(subchainVoucherAddress, subchainClient)
	subchainNFTOwner, _ := subchainVoucherContract.OwnerOf(nil, tokenID)
	if subchainNFTOwner != receiver {
		log.Fatalf("subchain token owner and receiver mismatch: %v vs. %v\n", subchainNFTOwner, receiver)
	}
	subchainTNT721Name, _ := subchainVoucherContract.Name(nil)
	subchainTNT721Symbol, _ := subchainVoucherContract.Symbol(nil)
	subchainTNT721TokenURI, _ := subchainVoucherContract.TokenURI(nil, tokenID)

	fmt.Printf("Subchain TNT721 Voucher contract address: %v, Name: %v, Symbol: %v, TokenURI: %v\n", subchainVoucherAddress, subchainTNT721Name, subchainTNT721Symbol, subchainTNT721TokenURI)
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
	fmt.Printf("Subchain TNT721 Voucher contract address: %v, Name: %v, Symbol: %v, TokenURI: %v\n", subchainTNT721VoucherAddress, subchainTNT721Name, subchainTNT721Symbol, subchainTNT721TokenURI)

	subchainNFTOwner, _ := subchainTNT721Voucher.OwnerOf(nil, tokenID)
	if subchainNFTOwner != sender {
		log.Fatalf("subchain token owner and sender mismatch: %v vs. %v\n", subchainNFTOwner, sender)
	}
	fmt.Printf("Subchain NFT Voucher sender: %v, tokenID: %v\n", sender, tokenID)
	fmt.Printf("Mainchain NFT receiver     : %v, tokenID: %v\n\n", receiver, tokenID)

	mainchainTNT721ContractAddress := common.HexToAddress("0xEd8d61f42dC1E56aE992D333A4992C3796b22A74")
	mainchainTNT721Contract, _ := ct.NewTNT721VoucherContract(mainchainTNT721ContractAddress, mainchainClient)
	mainchainTNT721Owner, _ := mainchainTNT721Contract.OwnerOf(nil, tokenID)

	// printTNT721TokenBankLockedAmount(mainchainTNT721TokenBankAddress, mainchainTNT721ContractAddress, subchainID, tokenID, mainchainClient)

	authUser = subchainSelectAccount(subchainClient, 1)
	authUser.Value.Set(crossChainFee)
	burnTx, err := subchainTNT721TokenBank.BurnVouchers(authUser, subchainTNT721VoucherAddress, receiver, tokenID)
	authUser.Value.Set(common.Big0)
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

	// printTNT721TokenBankLockedAmount(mainchainTNT721TokenBankAddress, mainchainTNT721ContractAddress, subchainID, tokenID, mainchainClient)

	mainchainNFTOwner, _ := mainchainTNT721Contract.OwnerOf(nil, tokenID)
	if mainchainNFTOwner != receiver {
		log.Fatalf("mainchain token owner and receiver mismatch: %v vs. %v\n", mainchainNFTOwner, receiver)
	}
	mainchainTNT721Name, _ := mainchainTNT721Contract.Name(nil)
	mainchainTNT721Symbol, _ := mainchainTNT721Contract.Symbol(nil)
	mainchainTNT721TokenURI, _ := mainchainTNT721Contract.TokenURI(nil, tokenID)
	fmt.Printf("Mainchain TNT721 token contract address: %v, Name: %v, Symbol: %v, TokenURI: %v\n", mainchainTNT721ContractAddress, mainchainTNT721Name, mainchainTNT721Symbol, mainchainTNT721TokenURI)
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

	subchainTNT721Address := common.HexToAddress("0x0293801741ceF9465b2cf717578e57255863E8B2")
	subchainTNT721Instance, _ := ct.NewMockTNT721(subchainTNT721Address, subchainClient)
	subchainTNT721TokenBankinstance, err := ct.NewTNT721TokenBank(subchainTNT721TokenBankAddress, subchainClient)
	if err != nil {
		log.Fatal(err)
	}

	auth := subchainSelectAccount(subchainClient, 1)
	mintTx, _ := subchainTNT721Instance.Mint(auth, sender, tokenID)
	time.Sleep(2 * time.Second)
	fmt.Printf("Mint token tx: %v\n", mintTx.Hash().Hex())

	auth = subchainSelectAccount(subchainClient, 1)
	_, err = subchainTNT721Instance.Approve(auth, subchainTNT721TokenBankAddress, tokenID)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(2 * time.Second)

	subchainNFTOwner, _ := subchainTNT721Instance.OwnerOf(nil, tokenID)
	if subchainNFTOwner != sender {
		log.Fatalf("subchain token owner and sender mismatch: %v vs. %v\n", subchainNFTOwner, sender)
	}
	subchainTNT721Name, _ := subchainTNT721Instance.Name(nil)
	subchainTNT721Symbol, _ := subchainTNT721Instance.Symbol(nil)
	subchainTNT721TokenURI, _ := subchainTNT721Instance.TokenURI(nil, tokenID)
	fmt.Printf("Subchain TNT721 contract address: %v, Name: %v, Symbol: %v, TokenURI: %v\n", subchainTNT721Address, subchainTNT721Name, subchainTNT721Symbol, subchainTNT721TokenURI)

	fmt.Printf("Subchain NFT sender           : %v, tokenID: %v\n", sender, tokenID)
	fmt.Printf("Mainchain NFT Voucher receiver: %v, tokenID: %v\n\n", receiver, tokenID)

	// mainchainID, _ := mainchainClient.ChainID(context.Background())
	// printTNT721TokenBankLockedAmount(subchainTNT721TokenBankAddress, subchainTNT721Address, mainchainID, tokenID, subchainClient)

	auth = subchainSelectAccount(subchainClient, 1)
	auth.Value.Set(crossChainFee)
	lockTx, err := subchainTNT721TokenBankinstance.LockTokens(auth, big.NewInt(366), subchainTNT721Address, accountList[6].fromAddress, tokenID)
	auth.Value.Set(common.Big0)
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
		result := getMainchainTNT721VoucherMintLogs(int(fromHeight), int(toHeight), mainchainTNT721TokenBankAddress, accountList[6].fromAddress)
		if result != nil {
			mainchainVoucherAddress = *result
			break
		}
		fmt.Printf("waiting for cross-chain transfer completion (scanning Mainchain from height %v to %v)...\n", fromHeight, toHeight)
	}
	fmt.Printf("End transfer, timestamp        : %v\n", time.Now())
	fmt.Printf("Cross-chain transfer completed.\n\n")

	// printTNT721TokenBankLockedAmount(subchainTNT721TokenBankAddress, subchainTNT721Address, mainchainID, tokenID, subchainClient)

	subchainNFTOwner, _ = subchainTNT721Instance.OwnerOf(nil, tokenID)
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
	fmt.Printf("Mainchain TNT721 Voucher contract address: %v, Name: %v, Symbol: %v, TokenURI: %v\n", mainchainVoucherAddress, mainchainTNT721Name, mainchainTNT721Symbol, mainchainTNT721TokenURI)
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

	mainchainTNT721TokenBank, _ := ct.NewTNT721TokenBank(mainchainTNT721TokenBankAddress, mainchainClient)

	mainchainTNT721VoucherAddr := common.HexToAddress("0xBcd7d42CA0F1B43355823Cda5328aa8a5BF9038d")
	mainchainTNT721VoucherContract, _ := ct.NewTNT721VoucherContract(mainchainTNT721VoucherAddr, mainchainClient)
	auth := mainchainSelectAccount(mainchainClient, 6)
	_, err = mainchainTNT721VoucherContract.Approve(auth, mainchainTNT721TokenBankAddress, tokenID)
	if err != nil {
		log.Fatal(err)
	}
	mainchainTNT721Name, _ := mainchainTNT721VoucherContract.Name(nil)
	mainchainTNT721Symbol, _ := mainchainTNT721VoucherContract.Symbol(nil)
	mainchainTNT721TokenURI, _ := mainchainTNT721VoucherContract.TokenURI(nil, tokenID)
	fmt.Printf("Mainchain TNT721 Voucher contract address: %v, Name: %v, Symbol: %v, TokenURI: %v\n", mainchainTNT721VoucherAddr, mainchainTNT721Name, mainchainTNT721Symbol, mainchainTNT721TokenURI)

	mainchainNFTOwner, _ := mainchainTNT721VoucherContract.OwnerOf(nil, tokenID)
	if mainchainNFTOwner != sender {
		log.Fatalf("mainchain token owner and sender mismatch: %v vs. %v\n", mainchainNFTOwner, sender)
	}
	fmt.Printf("Mainchain NFT Voucher sender: %v, tokenID: %v\n", sender, tokenID)
	fmt.Printf("Subchain NFT receiver       : %v, tokenID: %v\n\n", receiver, tokenID)

	subchainTNT721ContractAddress := common.HexToAddress("0x0293801741ceF9465b2cf717578e57255863E8B2")
	// subchainTNT721ContractAddress := subchainTNT721TokenAddress
	subchainTNT721Contract, _ := ct.NewTNT721VoucherContract(subchainTNT721ContractAddress, subchainClient)
	subchainTNT721Owner, _ := subchainTNT721Contract.OwnerOf(nil, tokenID)

	// mainchainID, _ := mainchainClient.ChainID(context.Background())
	// printTNT721TokenBankLockedAmount(subchainTNT721TokenBankAddress, subchainTNT721ContractAddress, mainchainID, tokenID, subchainClient)

	auth = mainchainSelectAccount(mainchainClient, 6)
	auth.Value.Set(crossChainFee)
	burnTx, err := mainchainTNT721TokenBank.BurnVouchers(auth, mainchainTNT721VoucherAddr, accountList[1].fromAddress, tokenID)
	auth.Value.Set(common.Big0)
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

	// printTNT721TokenBankLockedAmount(subchainTNT721TokenBankAddress, subchainTNT721ContractAddress, mainchainID, tokenID, subchainClient)

	subchainNFTOwner, _ := subchainTNT721Contract.OwnerOf(nil, tokenID)
	if subchainNFTOwner != receiver {
		log.Fatalf("subchain token owner and receiver mismatch: %v vs. %v\n", mainchainNFTOwner, receiver)
	}
	subchainTNT721Name, _ := subchainTNT721Contract.Name(nil)
	subchainTNT721Symbol, _ := subchainTNT721Contract.Symbol(nil)
	subchainTNT721TokenURI, _ := subchainTNT721Contract.TokenURI(nil, tokenID)
	fmt.Printf("Subchain TNT721 Voucher contract address: %v, Name: %v, Symbol: %v, TokenURI: %v\n", subchainTNT721ContractAddress, subchainTNT721Name, subchainTNT721Symbol, subchainTNT721TokenURI)
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

// func printTNT721TokenBankLockedAmount(tokenBankAddr common.Address, tokenAddr common.Address, chainID *big.Int, tokenID *big.Int, ethclient *ethclient.Client) {
// 	tnt721TokenBank, _ := ct.NewTNT721TokenBank(tokenBankAddr, ethclient)
// 	tokenBankLockedAmount, _ := tnt721TokenBank.TotalLockedAmounts(nil, chainID, tokenAddr, tokenID)
// 	fmt.Printf("TNT721TokenBank locked amount for tokenID %v: %v\n\n", tokenID, tokenBankLockedAmount)
// }
