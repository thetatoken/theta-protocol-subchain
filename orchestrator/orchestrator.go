package orchestrator

import (
	log "github.com/sirupsen/logrus"
	// "github.com/thetatoken/theta/crypto"
	// "github.com/thetatoken/thetasubchain/eth/abi/bind"
)

// // SubchainRegisterSendToSubchainEvent
var logger *log.Entry = log.WithFields(log.Fields{"prefix": "orchestrator"})

// type UnprocessedVoucherContractBurnTx struct {
// 	blockHash common.Hash
// }

// type Orchestrator struct {
// 	mainChainID *big.Int
// 	subchainID  *big.Int

// 	privateKey *crypto.PrivateKey

// 	client                                *ec.Client
// 	collectTicker                         *time.Ticker

// 	registerContractAddr common.Address
// 	ercContractAddr      common.Address
// 	tfuelTokenBankContractAddr common.Address
// 	tnt20TokenBankContractAddr common.Address
// 	registerContract     *scta.SubchainRegister
// 	ercContract          *scta.SubchainERC
// 	tfuelTokenBankContract *scta.MainchainTFuelTokenBank
// 	tnt20TokenBankContract *scta.MainchainTNT20TokenBank

// 	collect int

// 	// Life cycle
// 	wg     *sync.WaitGroup
// 	ctx    context.Context
// 	cancel context.CancelFunc
// }

// // NewOrchestrator creates a new Orchestrator
// func NewOrchestrator(
// 	ethClientAddress string,
// 	subchainID *big.Int,
// 	registerContractAddr common.Address,
// 	ercContractAddr common.Address,
// 	tfuelTokenBankContractAddr common.Address,
// 	tnt20TokenBankContractAddr common.Address,
// 	collectInterval int,
// ) *Orchestrator {
// 	client, err := ec.Dial(ethClientAddress)
// 	if err != nil {
// 		logger.Fatalf("the eth client failed to connect %v\n", err)
// 	}
// 	subchainRegisterContract, err := scta.NewSubchainRegister(registerContractAddr, client)
// 	if err != nil {
// 		logger.Fatalf("failed to create subchain register contract %v\n", err)
// 	}
// 	subchainERCContract, err := scta.NewSubchainERC(ercContractAddr, client)
// 	if err != nil {
// 		logger.Fatalf("failed to create erc contract %v\n", err)
// 	}
// 	tfuelTokenBankContract, err := scta.NewMainchainTFuelTokenBank(tfuelTokenBankContractAddr, client)
// 	if err != nil {
// 		logger.Fatalf("failed to create TFuel token bank contract %v\n", err)
// 	}
// 	tnt20TokenBankContract, err := scta.NewMainchainTNT20TokenBank(tnt20TokenBankContractAddr, client)
// 	if err != nil {
// 		logger.Fatalf("failed to create TNT20 token bank contract %v\n", err)
// 	}

// 	mainChainID, err := client.ChainID(context.Background())
// 	if err != nil {
// 		logger.Fatalf("failed to get the chainID of the main chain %v\n", err)
// 	}
// 	logger.Printf("Create transfer validator for chain %d\n", mainChainID)

// 	oc := &Orchestrator{
// 		mainChainID:       mainChainID,
// 		subchainID:        subchainID,
// 		witnessedDynasty:  nil, // will be updated in the first collectVoucherBurnEvent() call
// 		validatorSetCache: make(map[string]*score.ValidatorSet),
// 		client:            client,

// 		registerContractAddr: registerContractAddr,
// 		ercContractAddr:      ercContractAddr,
// 		registerContract:     subchainRegisterContract,
// 		ercContract:          subchainERCContract,
// 		tokenBankContractAddr:

// 		collectInterval: collectInterval,

// 		crossChainEventCache: crossChainEventCache,

// 		lastQueryedMainChainHeight: big.NewInt(0),

// 		wg: &sync.WaitGroup{},
// 	}
// 	return oc
// }

// func (oc *Orchestrator) Start(ctx context.Context) {
// 	c, cancel := context.WithCancel(ctx)
// 	oc.ctx = c
// 	oc.cancel = cancel

// 	oc.wg.Add(1)
// 	go oc.mainloop(ctx)
// }

// func (oc *Orchestrator) Stop() {
// 	if oc.collectTicker != nil {
// 		oc.collectTicker.Stop()
// 	}
// 	oc.cancel()
// }

// func (oc *Orchestrator) Wait() {
// 	oc.wg.Wait()
// }

// func (oc *Orchestrator) mainloop(ctx context.Context) {
// 	oc.collectTicker = time.NewTicker(time.Duration(oc.collectInterval) * time.Millisecond)
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			return
// 		case <-oc.collectTicker.C:
// 			oc.collectVoucherBurnEvent()
// 		}
// 	}
// }

// func (oc *Orchestrator) collectVoucherBurnEvent() {
// 	从上次处理的，到lastFinalizedBlock
// }

// func (oc *Orchestrator) submitVoucherBurnToMainchain() {

// }

// // Get the end nonce that the orchestrator should process, start nonce should be fetched from DB.
// func (oc *Orchestrator) submitVoucherBurnToMainchainTillNonce() *big.Int{

// }

// func (oc *Orchestrator) buildTxOpts() *bind.TransactOpts {
// 	gasPrice, err := client.SuggestGasPrice(context.Background())
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	txOpts, err := bind.NewKeyedTransactorWithChainID(oc.privateKey, oc.mainChainID)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	txOpts.Nonce = big.NewInt(int64(nonce))
// 	txOpts.Value = big.NewInt(0)       // in wei
// 	txOpts.GasLimit = uint64(10000000) // in units
// 	txOpts.GasPrice = gasPrice

// 	return txOpts
// }

// func (oc *Orchestrator) AddToVoucherBurnTxQueue(event score.InterChainMessageEvent) (ok bool) {
// 	oc.Queue.insert()

// }
// func (oc *Orchestrator) HasVoucherBurnToProcess() (yes bool) {

// }

// func (oc *Orchestrator) CallVourcherBurnOnMainchain(type) (common.Hash, error) {

// 	subchainID, dynasty,
// 	TxHash, denom, receiver, amount,
// 	TFuel amount
// 	validatorSig
// 	return proxySctx, nil
// }
