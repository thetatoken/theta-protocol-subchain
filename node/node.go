package node

import (
	"context"
	"log"
	"math/big"
	"reflect"
	"sync"

	"github.com/spf13/viper"
	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/theta/crypto"
	dp "github.com/thetatoken/theta/dispatcher"
	"github.com/thetatoken/theta/p2p"
	"github.com/thetatoken/theta/p2pl"
	"github.com/thetatoken/theta/store"
	"github.com/thetatoken/theta/store/database"
	"github.com/thetatoken/theta/store/kvstore"

	sbc "github.com/thetatoken/thetasubchain/blockchain"
	scom "github.com/thetatoken/thetasubchain/common"
	sconsensus "github.com/thetatoken/thetasubchain/consensus"
	score "github.com/thetatoken/thetasubchain/core"
	"github.com/thetatoken/thetasubchain/interchain/orchestrator"
	siu "github.com/thetatoken/thetasubchain/interchain/utils"
	"github.com/thetatoken/thetasubchain/interchain/witness"

	sld "github.com/thetatoken/thetasubchain/ledger"
	smp "github.com/thetatoken/thetasubchain/mempool"
	snsync "github.com/thetatoken/thetasubchain/netsync"
	srpc "github.com/thetatoken/thetasubchain/rpc"
	ssnst "github.com/thetatoken/thetasubchain/snapshot"
	srollingdb "github.com/thetatoken/thetasubchain/store/rollingdb"
)

type Node struct {
	Store                store.Store
	Chain                *sbc.Chain
	Consensus            *sconsensus.ConsensusEngine
	ValidatorManager     score.ValidatorManager
	SyncManager          *snsync.SyncManager
	Dispatcher           *dp.Dispatcher
	Ledger               score.Ledger
	Mempool              *smp.Mempool
	RPC                  *srpc.ThetaRPCServer
	InterChainEventCache *siu.InterChainEventCache
	MainchainWitness     witness.ChainWitness
	Orchestrator         orchestrator.ChainOrchestrator

	// reporter *srp.Reporter

	// Life cycle
	wg      *sync.WaitGroup
	quit    chan struct{}
	ctx     context.Context
	cancel  context.CancelFunc
	stopped bool
}

type Params struct {
	ChainID             string
	GasPriceLimit       *big.Int
	PrivateKey          *crypto.PrivateKey
	Root                *score.Block
	NetworkOld          p2p.Network
	Network             p2pl.Network
	DB                  database.Database
	RollingDB           *srollingdb.RollingDB
	SnapshotPath        string
	ChainImportDirPath  string
	ChainCorrectionPath string
}

func NewNode(params *Params) *Node {
	store := kvstore.NewKVStore(params.DB)
	chain := sbc.NewChain(params.ChainID, store, params.Root)
	params.RollingDB.SetChain(chain)

	validatorManager := sconsensus.NewRotatingValidatorManager()
	dispatcher := dp.NewDispatcher(params.NetworkOld, params.Network)

	interChainEventCache := siu.NewInterChainEventCache(params.DB)

	// // For testing...
	// metachainWitness := witness.NewSimulatedMetachainWitness(
	// 	"privatenet",
	// 	params.ChainID,
	// 	interChainEventCache,
	// 	0)
	metachainWitness := witness.NewMetachainWitness(
		params.DB,
		viper.GetInt(scom.CfgSubchainUpdateIntervalInMilliseconds),
		interChainEventCache)
	orchestrator := orchestrator.NewOrchestrator(
		params.DB,
		viper.GetInt(scom.CfgSubchainUpdateIntervalInMilliseconds),
		interChainEventCache,
		metachainWitness,
		params.PrivateKey,
	)

	consensus := sconsensus.NewConsensusEngine(params.PrivateKey, store, chain, dispatcher, validatorManager, metachainWitness)
	// reporter := srp.NewReporter(dispatcher, consensus, chain)

	syncMgr := snsync.NewSyncManager(chain, consensus, params.NetworkOld, params.Network, dispatcher, consensus)
	mempool := smp.CreateMempool(dispatcher, consensus)
	ledger := sld.NewLedger(params.ChainID, params.RollingDB, params.RollingDB, chain, consensus, validatorManager, mempool, metachainWitness)

	validatorManager.SetConsensusEngine(consensus)
	consensus.SetLedger(ledger)
	mempool.SetLedger(ledger)

	txMsgHandler := smp.CreateMempoolMessageHandler(mempool)

	if !reflect.ValueOf(params.Network).IsNil() {
		params.Network.RegisterMessageHandler(txMsgHandler)
	}
	if !reflect.ValueOf(params.NetworkOld).IsNil() {
		params.NetworkOld.RegisterMessageHandler(txMsgHandler)
	}

	currentHeight := consensus.GetLastFinalizedBlock().Height
	if currentHeight <= params.Root.Height {
		snapshotPath := params.SnapshotPath
		chainImportDirPath := params.ChainImportDirPath
		chainCorrectionPath := params.ChainCorrectionPath
		var lastCC *score.ExtendedBlock
		var err error
		if _, lastCC, err = ssnst.ImportSnapshot(snapshotPath, chainImportDirPath, chainCorrectionPath, chain, params.DB, ledger); err != nil {
			log.Fatalf("Failed to load snapshot: %v, err: %v", snapshotPath, err)
		}
		if lastCC != nil {
			state := consensus.State()
			state.SetLastFinalizedBlock(lastCC)
			state.SetHighestCCBlock(lastCC)
			state.SetLastVote(score.Vote{})
			state.SetLastProposal(score.Proposal{})
		}
	}
	metachainWitness.SetSubchainTokenBanks(ledger)
	orchestrator.SetLedgerAndSubchainTokenBanks(ledger)
	node := &Node{
		Store:                store,
		Chain:                chain,
		Consensus:            consensus,
		ValidatorManager:     validatorManager,
		SyncManager:          syncMgr,
		Dispatcher:           dispatcher,
		Ledger:               ledger,
		Mempool:              mempool,
		InterChainEventCache: interChainEventCache,
		MainchainWitness:     metachainWitness,
		Orchestrator:         orchestrator,
		// reporter:             reporter,
	}

	if viper.GetBool(common.CfgRPCEnabled) {
		node.RPC = srpc.NewThetaRPCServer(mempool, ledger, dispatcher, chain, consensus)
	}
	return node
}

// Start starts sub components and kick off the main loop.
func (n *Node) Start(ctx context.Context) {
	c, cancel := context.WithCancel(ctx)
	n.ctx = c
	n.cancel = cancel

	n.Consensus.Start(n.ctx)
	n.SyncManager.Start(n.ctx)
	n.Dispatcher.Start(n.ctx)
	n.Mempool.Start(n.ctx)
	// n.reporter.Start(n.ctx)
	n.MainchainWitness.Start(n.ctx)
	n.Orchestrator.Start(n.ctx)

	if viper.GetBool(common.CfgRPCEnabled) {
		n.RPC.Start(n.ctx)
	}
}

// Stop notifies all sub components to stop without blocking.
func (n *Node) Stop() {
	n.cancel()
}

// Wait blocks until all sub components stop.
func (n *Node) Wait() {
	n.Consensus.Wait()
	n.SyncManager.Wait()
	n.MainchainWitness.Wait()

	if n.RPC != nil {
		n.RPC.Wait()
	}
}
