package ledger

import (
	"context"
	"fmt"
	"math/big"
	"strconv"
	"sync"

	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/theta/crypto"
	dp "github.com/thetatoken/theta/dispatcher"
	"github.com/thetatoken/theta/ledger/types"
	"github.com/thetatoken/theta/p2p"
	p2psim "github.com/thetatoken/theta/p2p/simulation"
	"github.com/thetatoken/theta/p2pl"
	"github.com/thetatoken/theta/store/database"
	"github.com/thetatoken/theta/store/database/backend"
	"github.com/thetatoken/theta/store/kvstore"

	sbc "github.com/thetatoken/thetasubchain/blockchain"
	scom "github.com/thetatoken/thetasubchain/common"
	sconsensus "github.com/thetatoken/thetasubchain/consensus"
	score "github.com/thetatoken/thetasubchain/core"
	sexec "github.com/thetatoken/thetasubchain/ledger/execution"
	slst "github.com/thetatoken/thetasubchain/ledger/state"
	stypes "github.com/thetatoken/thetasubchain/ledger/types"
	smp "github.com/thetatoken/thetasubchain/mempool"
)

type mockSnapshot struct {
	block *score.Block
	vsp   *score.ValidatorSet
}

type execSim struct {
	chainID   string
	chain     *sbc.Chain
	state     *slst.LedgerState
	consensus *sconsensus.ConsensusEngine
	executor  *sexec.Executor
}

func newExecSim(chainID string, db database.Database, snapshot mockSnapshot, valPrivAcc *types.PrivAccount) *execSim {
	initHeight := snapshot.block.Height

	sv := slst.NewStoreView(initHeight, common.Hash{}, db)
	chainIDInt := scom.MapChainID(chainID)
	sv.UpdateValidatorSet(chainIDInt, snapshot.vsp)

	store := kvstore.NewKVStore(db)
	chain := sbc.NewChain(chainID, store, snapshot.block)

	p2psimnet := p2psim.NewSimnetWithHandler(nil)
	messenger := p2psimnet.AddEndpoint("peerID0")

	dispatcher := dp.NewDispatcher(messenger, nil)

	valMgr := sconsensus.NewFixedValidatorManager()
	consensus := sconsensus.NewConsensusEngine(valPrivAcc.PrivKey, store, chain, dispatcher, valMgr, nil)
	valMgr.SetConsensusEngine(consensus)

	mempool := smp.CreateMempool(dispatcher, consensus)

	ledgerState := slst.NewLedgerState(chainID, db, nil)
	//ledgerState.ResetState(initHeight, snapshot.block.StateHash)
	ledgerState.ResetState(snapshot.block)

	ledger := &Ledger{
		consensus: consensus,
		valMgr:    valMgr,
		mempool:   mempool,
		mu:        &sync.RWMutex{},
		state:     ledgerState,
	}
	executor := sexec.NewExecutor(db, chain, ledgerState, consensus, valMgr, ledger, nil)
	ledger.SetExecutor(executor)

	consensus.SetLedger(ledger)

	es := &execSim{
		chainID:   chainID,
		chain:     chain,
		state:     ledgerState,
		consensus: consensus,
		executor:  executor,
	}

	return es
}

func (es *execSim) addBlock(block *score.Block) {
	es.chain.AddBlock(block)
}

func (es *execSim) finalizePreviousBlocks(blockHash common.Hash) {
	es.chain.FinalizePreviousBlocks(blockHash)
}

func (es *execSim) getTipBlock() *score.ExtendedBlock {
	return es.consensus.GetTip(true)
}

func (es *execSim) findBlocksByHeight(height uint64) []*score.ExtendedBlock {
	return es.chain.FindBlocksByHeight(height)
}

func newTestLedger() (chainID string, ledger *Ledger, mempool *smp.Mempool) {
	chainID = "test_chain_id"
	peerID := "peer0"
	proposerSeed := "proposer"

	db := backend.NewMemDatabase()
	chain := &sbc.Chain{ChainID: chainID}
	consensus := sexec.NewTestConsensusEngine(proposerSeed)
	valMgr := newTesetValidatorManager(consensus)
	p2psimnet := p2psim.NewSimnetWithHandler(nil)
	messenger := p2psimnet.AddEndpoint(peerID)
	mempool = newTestMempool(peerID, messenger, nil)
	ledger = NewLedger(chainID, db, nil, chain, consensus, valMgr, mempool, nil)
	mempool.SetLedger(ledger)

	ctx := context.Background()
	messenger.Start(ctx)
	mempool.Start(ctx)

	initHeight := uint64(1)
	initRootHash := common.Hash{}

	initBlock := &score.Block{
		BlockHeader: &score.BlockHeader{
			ChainID:   chainID,
			Height:    initHeight,
			StateHash: initRootHash,
		},
	}
	//ledger.ResetState(initHeight, initRootHash)
	ledger.ResetState(initBlock)

	return chainID, ledger, mempool
}

func newTesetValidatorManager(consensus score.ConsensusEngine) score.ValidatorManager {
	proposerAddressStr := consensus.PrivateKey().PublicKey().Address().String()
	propser := score.NewValidator(proposerAddressStr, new(big.Int).SetUint64(999))

	_, val2PubKey, err := crypto.TEST_GenerateKeyPairWithSeed("val2")
	if err != nil {
		panic(fmt.Sprintf("Failed to generate key pair with seed: %v", err))
	}
	val2 := score.NewValidator(val2PubKey.Address().String(), new(big.Int).SetUint64(100))

	valSet := score.NewValidatorSet(big.NewInt(0))
	valSet.AddValidator(propser)
	valSet.AddValidator(val2)
	valMgr := sexec.NewTestValidatorManager(propser, valSet)

	return valMgr
}

func newTestMempool(peerID string, messenger p2p.Network, messengerL p2pl.Network) *smp.Mempool {
	dispatcher := dp.NewDispatcher(messenger, nil)
	mempool := smp.CreateMempool(dispatcher, nil)
	txMsgHandler := smp.CreateMempoolMessageHandler(mempool)
	messenger.RegisterMessageHandler(txMsgHandler)
	return mempool
}

func prepareInitLedgerState(ledger *Ledger, numInAccs int) (accOut types.PrivAccount, accIns []types.PrivAccount) {
	txFee := getMinimumTxFee()
	validators := ledger.valMgr.GetValidatorSet(common.Hash{}).Validators()
	for _, val := range validators {
		valAccount := &types.Account{
			Address:                val.Address,
			LastUpdatedBlockHeight: 1,
			Balance:                types.NewCoins(100000000000, 1000),
		}
		ledger.state.Delivered().SetAccount(val.Address, valAccount)
	}

	accOut = types.MakeAccWithInitBalance("accOut", types.NewCoins(700000, 3))
	ledger.state.Delivered().SetAccount(accOut.Account.Address, &accOut.Account)

	for i := 0; i < numInAccs; i++ {
		secret := "in_secret_" + strconv.FormatInt(int64(i), 16)
		accIn := types.MakeAccWithInitBalance(secret, types.NewCoins(900000, 50000*txFee))
		accIns = append(accIns, accIn)
		ledger.state.Delivered().SetAccount(accIn.Account.Address, &accIn.Account)
	}

	ledger.state.Commit()

	return accOut, accIns
}

func newRawCoinbaseTx(chainID string, ledger *Ledger, sequence int) common.Bytes {
	vaList := ledger.valMgr.GetValidatorSet(common.Hash{}).Validators()
	if len(vaList) < 2 {
		panic("Insufficient number of validators")
	}
	outputs := []types.TxOutput{}
	for _, val := range vaList {
		output := types.TxOutput{val.Address, types.NewCoins(0, 0)}
		outputs = append(outputs, output)
	}

	proposerSk := ledger.consensus.PrivateKey()
	proposerPk := proposerSk.PublicKey()
	coinbaseTx := &types.CoinbaseTx{
		Proposer:    types.TxInput{Address: proposerPk.Address(), Sequence: uint64(sequence)},
		Outputs:     outputs,
		BlockHeight: 2,
	}

	signBytes := coinbaseTx.SignBytes(chainID)
	sig, err := proposerSk.Sign(signBytes)
	if err != nil {
		panic("Failed to sign the coinbase transaction")
	}
	if !coinbaseTx.SetSignature(proposerPk.Address(), sig) {
		panic("Failed to set signature for the coinbase transaction")
	}

	coinbaseTxBytes, err := stypes.TxToBytes(coinbaseTx)
	if err != nil {
		panic(err)
	}
	return coinbaseTxBytes
}

func newRawSendTx(chainID string, sequence int, addPubKey bool, accOut, accIn types.PrivAccount, injectFeeFluctuation bool) common.Bytes {
	delta := int64(0)
	var err error
	if injectFeeFluctuation {
		// inject so fluctuation into the txFee, so later we can test whether the
		// mempool orders the txs by txFee
		randint, err := strconv.ParseInt(string(accIn.Address.Hex()[2:9]), 16, 64)
		if randint < 0 {
			randint = -randint
		}
		delta = randint * int64(types.GasRegularTxJune2021)
		if err != nil {
			panic(err)
		}
	}
	txFee := getMinimumTxFee() + delta
	sendTx := &types.SendTx{
		Fee: types.NewCoins(0, txFee),
		Inputs: []types.TxInput{
			{
				Sequence: uint64(sequence),
				Address:  accIn.Address,
				Coins:    types.NewCoins(15, txFee),
			},
		},
		Outputs: []types.TxOutput{
			{
				Address: accOut.Address,
				Coins:   types.NewCoins(15, 0),
			},
		},
	}

	signBytes := sendTx.SignBytes(chainID)
	inAccs := []types.PrivAccount{accIn}
	for idx, in := range sendTx.Inputs {
		inAcc := inAccs[idx]
		sig, err := inAcc.PrivKey.Sign(signBytes)
		if err != nil {
			panic("Failed to sign the coinbase transaction")
		}
		sendTx.SetSignature(in.Address, sig)
	}

	sendTxBytes, err := stypes.TxToBytes(sendTx)
	if err != nil {
		panic(err)
	}
	return sendTxBytes
}

func getMinimumTxFee() int64 {
	return int64(types.MinimumTransactionFeeTFuelWei)
}
