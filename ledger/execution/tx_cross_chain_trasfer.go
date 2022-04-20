
func (exec *CrossChainTransferExecutor) sanityCheck(chainID string, view *slst.StoreView, transaction types.Tx) result.Result {
	tx := transaction.(*stypes.CrossChainTransferTx)

	// basic checks, and verify the tx proposer is a validator
	xxxx // see tx_validator_set_upadate.go

	crossChainTransferID := exec.state.GetCrossChainTransferID(tx)
	if view.ChainTransferHasBeenProcessed(crossChainTransferID) {
		return result.Error("cross-chain chain transfer %v has already been processed", crossChainTransferID)
	}

	smartContracTx, err := constructSmartContractTx(tx) // construct a smart contract tx from the special tx (proposed and siged by the leader)
	if err != nil {
		return result.Error("error constructing smart contract tx: %v", err)
	}

	// Other checks, e.g. signature check..
	xxxx
	
	_, result := exec.smartContractExecutor.sanityCheck(chainID, view, smartContractTx)
	if result.IsError() {
		return result
	}
}

func (exec *CrossChainTransferExecutor) process(chainID string, view *slst.StoreView, transaction types.Tx) (common.Hash, result.Result) {
	tx := transaction.(*stypes.CrossChainTransferTx)
	crossChainTransferID := exec.state.GetCrossChainTransferID(tx)
	xferDetails, err := exec.mainchainWitness.GetCrossChainTransfer(crossChainTransferID)
	if err != nil || xferDetails == nil {
		return common.Hash{}, result.UndecidedWith(...) // not seen on mainchain yet
	}

	if view.ChainTransferHasBeenProcessed(crossChainTransferID) {
		return common.Hash{}, result.Error("cross-chain chain transfer %v has already been processed", crossChainTransferID)
	}

	// the leader could be malicious, so we need to verify if the cross-chain xfer proposed by the leader is consistent with the query results
	if !verifyTransferDetails(xferDetails, tx) {
		return common.Hash{}, result.Error("cross-chain transfer verification failed for ID %v", crossChainTransferID)
	}

	smartContracTx, err := constructSmartContractTx(tx) // construct a smart contract tx from the special tx (proposed and siged by the leader)
	if err != nil {
		return common.Hash{}, result.Error("error constructing smart contract tx: %v", err)
	}

	mainchainTokenContractAddress := tx.GetSourceTokenContractAddrOnMainChain()
	isContractDeployment := (smartContracTx.To.Address == commmon.Address{})
	tokenVoucherContractAddress := view.GetTokenVoucherContract(mainchainTokenContractAddress)
	if isContractDeployment && (tokenVoucherContractAddress != commmon.Address{}) {
		// the token voucher contract should not have been deployed, otherwise the tx is malicious
		return result.Error("token voucher contract for %v has already been deployed at %v, the tx is invalid", mainchainTokenContractAddress.Hex(), tokenVoucherContractAddress.Hex())
	}
	if !isContractDeployment && (tokenVoucherContractAddress != smartContracTx.To) {
		return result.Error("token voucher contract mismatch, %v vs %v", tokenVoucherContractAddress.Hex(), smartContracTx.To.Hex())
	}

	_, result := exec.smartContractExecutor.process(chainID, view, smartContractTx)
	if result.IsError() {
		return result
	}

	if isContractDeployment {
		tokenVoucherContractAddress = (result.Info[contractAddrInfoKey]).(commmon.Address)
		view.SetTokenVoucherContract(mainchainTokenContractAddress, tokenVoucherContractAddress)
	}

	sourceTokenContractAddrOnMainChain := xferDetails.SourceTokenContractAddrOnMainChain
	if sourceTokenContractAddrOnMainChain == nil

	view.MarkCrossChainTransferAsProcessed(crossChainTransferID)
}

// to be called by Ledger to compose the special tx for cross-chain transfers
func GenerateCrossChainTransferTx(proposerAddress common.Address, view *slst.StoreView) (*types.SmartContractTx, error) {
	xferDetails, err := exec.mainchainWitness.GetCrossChainTransfer(crossChainTransferID)
	if err != nil || xferDetails == nil {
		return nil, err
	}

	proposerSeq := getProposerSequence(proposerAddress)

	from := types.TxInput{
		Address: proposerAddr,
		Coins: types.Coins{
			ThetaWei: new(big.Int).SetUint64(0),
			TFuelWei: new(big.Int).SetUint64(0),
		},
		Sequence: proposerSeq + 1,
	}

	sourceTokenContractAddrOnMainChain := xferDetails.SourceTokenContractAddrOnMainChain
	tokenVoucherContract := view.GetTokenVoucherContract(sourceTokenContractAddrOnMainChain)
	var to types.TxOutput
	var data []byte
	if tokenVoucherContract == nil { // token voucher contract not yet exists, should deploy and mint
		// generate a tx with "deploy and mint" data field
		to.Address = commmon.Address{}
		data = xxxx // need to handle both TFuel and TNT tokens, will add a precomiled contract for minting TFuel
	} else { // token voucher contract already exists, should mint the token voucher
		// generate a tx with "mint" data field
		to.Address = tokenVoucherContract.Address
		data = xxxx // need to handle both TFuel and TNT tokens
	}

	sctx := &types.SmartContractTx{
		From:     from,
		To:       to,
		GasLimit: gasLimit,
		GasPrice: gasPrice,
		Data:     data,
	}

	return sctx, nil
}

