package orchestrator

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/theta/crypto"
	scom "github.com/thetatoken/thetasubchain/common"
	score "github.com/thetatoken/thetasubchain/core"
	"github.com/thetatoken/thetasubchain/eth/abi"
	"github.com/thetatoken/thetasubchain/eth/abi/bind"
	"github.com/thetatoken/thetasubchain/eth/core/types"
	ec "github.com/thetatoken/thetasubchain/eth/ethclient"
	scta "github.com/thetatoken/thetasubchain/interchain/contracts/accessors"
)

//
// A minimal JSON-RPC stand-in. The guards go through the real abigen bindings and
// the real eth client, so the ABI encoding and the RPC plumbing are exercised too --
// only the node at the far end is faked.
//

type rpcRequest struct {
	ID     json.RawMessage   `json:"id"`
	Method string            `json:"method"`
	Params []json.RawMessage `json:"params"`
}

type fakeNode struct {
	server *httptest.Server
	// handler returns (result, rpcError). A non-empty rpcError is returned as a
	// JSON-RPC error object, which is how a revert surfaces.
	handler  func(method string, params []json.RawMessage) (string, string)
	delay    time.Duration
	calls    int64
	lastData string
}

func newFakeNode(t *testing.T, handler func(method string, params []json.RawMessage) (string, string)) *fakeNode {
	t.Helper()
	fn := &fakeNode{handler: handler}
	fn.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req rpcRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		atomic.AddInt64(&fn.calls, 1)
		if len(req.Params) > 0 {
			fn.lastData = string(req.Params[0])
		}
		if fn.delay > 0 {
			// Mimic the adaptor's retry-with-sleep behaviour on a reverting call.
			select {
			case <-time.After(fn.delay):
			case <-r.Context().Done():
				return // the caller gave up, which is exactly what we want to observe
			}
		}
		result, rpcErr := fn.handler(req.Method, req.Params)
		w.Header().Set("Content-Type", "application/json")
		if rpcErr != "" {
			fmt.Fprintf(w, `{"jsonrpc":"2.0","id":%s,"error":{"code":-32000,"message":%q}}`, req.ID, rpcErr)
			return
		}
		fmt.Fprintf(w, `{"jsonrpc":"2.0","id":%s,"result":%q}`, req.ID, result)
	}))
	t.Cleanup(fn.server.Close)
	return fn
}

func (fn *fakeNode) client(t *testing.T) *ec.Client {
	t.Helper()
	client, err := ec.Dial(fn.server.URL)
	if err != nil {
		t.Fatalf("failed to dial the fake node: %v", err)
	}
	t.Cleanup(client.Close)
	return client
}

// word encodes a uint256 return value.
func word(v *big.Int) string {
	return "0x" + fmt.Sprintf("%064x", v)
}

// addrWord encodes an address return value.
func addrWord(a common.Address) string {
	return "0x" + fmt.Sprintf("%064s", strings.TrimPrefix(a.Hex(), "0x"))
}

const (
	testMainchainID = 361
	testSubchainID  = 360890
	// Denoms are "<originatedChainID>/<tokenType>/<contractAddress>"; the token type
	// is the standard's number, so TNT20 is 20 rather than 2.
	testTNT20Denom   = "361/20/0xaf537fb7e4c77c97403de94ce141b7edb9f7fcf0"
	testTNT721Denom  = "361/721/0xaf537fb7e4c77c97403de94ce141b7edb9f7fcf0"
	testTNT1155Denom = "361/1155/0xaf537fb7e4c77c97403de94ce141b7edb9f7fcf0"
)

var (
	testTokenAddr = common.HexToAddress("0xaf537fb7e4c77c97403de94ce141b7edb9f7fcf0")
	testBankAddr  = common.HexToAddress("0x0c99f551ac2d69510db41f8c8d3348ca55ef099d")
)

// newTestOrchestrator wires an Orchestrator whose main chain RPC points at the fake
// node. The collateral guards are enabled unless a test turns them off.
func newTestOrchestrator(t *testing.T, fn *fakeNode) *Orchestrator {
	t.Helper()
	viper.Set(scom.CfgSubchainEnforceUnlockCollateral, true)
	t.Cleanup(func() { viper.Set(scom.CfgSubchainEnforceUnlockCollateral, true) })

	return &Orchestrator{
		mainchainID:                   big.NewInt(testMainchainID),
		subchainID:                    big.NewInt(testSubchainID),
		mainchainEthRpcClient:         fn.client(t),
		mainchainTNT20TokenBankAddr:   testBankAddr,
		mainchainTNT721TokenBankAddr:  testBankAddr,
		mainchainTNT1155TokenBankAddr: testBankAddr,
	}
}

//
// TNT20 -- the standard that was actually drained on 2026-08-09.
//

// The ordinary case: the bank holds more than the unlock asks for.
func TestTNT20CollateralAcceptsSufficientBalance(t *testing.T) {
	fn := newFakeNode(t, func(method string, _ []json.RawMessage) (string, string) {
		return word(big.NewInt(10_000)), ""
	})
	oc := newTestOrchestrator(t, fn)

	err := oc.verifyTNT20UnlockCollateral(big.NewInt(testMainchainID), testTNT20Denom, big.NewInt(9_999))
	assert.Nil(t, err)
}

// The incident case. A forged voucher burn asked for far more than the bank held;
// TFuelTokenBank rejected the equivalent on-chain, TNT20TokenBank did not. This is
// the check that has to catch it, and it must not be satisfiable by the dry run,
// since TNT20 unlockTokens() swallows a failed transfer in a try/catch.
func TestTNT20CollateralRejectsShortfall(t *testing.T) {
	assert := assert.New(t)

	fn := newFakeNode(t, func(_ string, _ []json.RawMessage) (string, string) {
		return word(big.NewInt(50)), "" // what the bank actually holds
	})
	oc := newTestOrchestrator(t, fn)

	err := oc.verifyTNT20UnlockCollateral(big.NewInt(testMainchainID), testTNT20Denom, big.NewInt(8_200))
	assert.NotNil(err)
	assert.ErrorIs(err, ErrUnlockUncollateralized)
	assert.Contains(err.Error(), "8200") // the amount asked for is in the alarm
}

// Exactly-equal collateral is sufficient: the guard must not be off by one and
// block a legitimate full withdrawal.
func TestTNT20CollateralAcceptsExactBalance(t *testing.T) {
	fn := newFakeNode(t, func(_ string, _ []json.RawMessage) (string, string) {
		return word(big.NewInt(8_200)), ""
	})
	oc := newTestOrchestrator(t, fn)

	assert.Nil(t, oc.verifyTNT20UnlockCollateral(big.NewInt(testMainchainID), testTNT20Denom, big.NewInt(8_200)))
}

// If the balance cannot be read the guard must fail closed. Skipping a round only
// delays a genuine transfer; voting blind is what let the drain through.
func TestTNT20CollateralFailsClosedWhenUnreadable(t *testing.T) {
	fn := newFakeNode(t, func(_ string, _ []json.RawMessage) (string, string) {
		return "", "connection refused"
	})
	oc := newTestOrchestrator(t, fn)

	assert.NotNil(t, oc.verifyTNT20UnlockCollateral(big.NewInt(testMainchainID), testTNT20Denom, big.NewInt(1)))
}

// The denom rides along with the event, so a malformed one must be rejected rather
// than resolving to the zero address and reading some unrelated balance.
func TestTNT20CollateralRejectsMalformedDenom(t *testing.T) {
	assert := assert.New(t)

	fn := newFakeNode(t, func(_ string, _ []json.RawMessage) (string, string) {
		return word(big.NewInt(1_000_000)), "" // would pass if the denom were parsed loosely
	})
	oc := newTestOrchestrator(t, fn)

	for _, denom := range []string{"361/20/not-an-address", "361/20", ""} {
		err := oc.verifyTNT20UnlockCollateral(big.NewInt(testMainchainID), denom, big.NewInt(1))
		assert.ErrorIs(err, ErrUnlockUncollateralized, "denom %q", denom)
	}
	assert.Zero(fn.calls, "a malformed denom must be rejected before any RPC call")
}

// A missing amount must be rejected rather than treated as zero, which would
// otherwise pass against any balance.
func TestTNT20CollateralRejectsNilAmount(t *testing.T) {
	fn := newFakeNode(t, func(_ string, _ []json.RawMessage) (string, string) {
		return word(big.NewInt(0)), ""
	})
	oc := newTestOrchestrator(t, fn)

	assert.ErrorIs(t, oc.verifyTNT20UnlockCollateral(big.NewInt(testMainchainID), testTNT20Denom, nil), ErrUnlockUncollateralized)
}

// The escape hatch for chains bridging an elastic-supply token, which is why the
// on-chain version was commented out in the first place.
func TestTNT20CollateralCanBeDisabled(t *testing.T) {
	assert := assert.New(t)

	fn := newFakeNode(t, func(_ string, _ []json.RawMessage) (string, string) {
		return word(big.NewInt(0)), ""
	})
	oc := newTestOrchestrator(t, fn)
	viper.Set(scom.CfgSubchainEnforceUnlockCollateral, false)

	assert.Nil(oc.verifyTNT20UnlockCollateral(big.NewInt(testMainchainID), testTNT20Denom, big.NewInt(8_200)))
	assert.Zero(fn.calls, "the disabled guard must not make RPC calls")
}

// The guard must read the bank on the chain the tokens are being released on.
func TestTNT20CollateralUsesTargetChainBank(t *testing.T) {
	fn := newFakeNode(t, func(_ string, _ []json.RawMessage) (string, string) {
		return word(big.NewInt(10_000)), ""
	})
	oc := newTestOrchestrator(t, fn)
	oc.subchainEthRpcClient = oc.mainchainEthRpcClient
	oc.subchainTNT20TokenBankAddr = common.HexToAddress("0x47e9fbef8c83a1714f1951f142132e6e90f5fa5d")

	assert.Nil(t, oc.verifyTNT20UnlockCollateral(big.NewInt(testSubchainID), testTNT20Denom, big.NewInt(1)))
	assert.Contains(t, strings.ToLower(fn.lastData), "47e9fbef8c83a1714f1951f142132e6e90f5fa5d",
		"balanceOf must be called against the subchain bank when the subchain is the target")
}

//
// TNT721 -- ownership rather than balance.
//

func TestTNT721CollateralRequiresBankOwnership(t *testing.T) {
	assert := assert.New(t)

	// Owned by someone else: the unlock has nothing behind it.
	fn := newFakeNode(t, func(_ string, _ []json.RawMessage) (string, string) {
		return addrWord(common.HexToAddress("0x8556d1a34013feb0cba9349636895d94831528a0")), ""
	})
	oc := newTestOrchestrator(t, fn)
	err := oc.verifyTNT721UnlockCollateral(big.NewInt(testMainchainID), testTNT721Denom, big.NewInt(7))
	assert.ErrorIs(err, ErrUnlockUncollateralized)

	// Owned by the bank: legitimate.
	fn2 := newFakeNode(t, func(_ string, _ []json.RawMessage) (string, string) {
		return addrWord(testBankAddr), ""
	})
	oc2 := newTestOrchestrator(t, fn2)
	assert.Nil(oc2.verifyTNT721UnlockCollateral(big.NewInt(testMainchainID), testTNT721Denom, big.NewInt(7)))
}

func TestTNT721CollateralRejectsNilTokenID(t *testing.T) {
	fn := newFakeNode(t, func(_ string, _ []json.RawMessage) (string, string) {
		return addrWord(testBankAddr), ""
	})
	oc := newTestOrchestrator(t, fn)

	assert.ErrorIs(t, oc.verifyTNT721UnlockCollateral(big.NewInt(testMainchainID), testTNT721Denom, nil), ErrUnlockUncollateralized)
}

//
// TNT1155 -- balance, but per token ID.
//

func TestTNT1155CollateralChecksPerTokenBalance(t *testing.T) {
	assert := assert.New(t)

	fn := newFakeNode(t, func(_ string, _ []json.RawMessage) (string, string) {
		return word(big.NewInt(3)), ""
	})
	oc := newTestOrchestrator(t, fn)

	assert.Nil(oc.verifyTNT1155UnlockCollateral(big.NewInt(testMainchainID), testTNT1155Denom, big.NewInt(1), big.NewInt(3)))

	err := oc.verifyTNT1155UnlockCollateral(big.NewInt(testMainchainID), testTNT1155Denom, big.NewInt(1), big.NewInt(4))
	assert.ErrorIs(err, ErrUnlockUncollateralized)
}

//
// F1 -- the dry run must be bounded, and a timeout must not look like a revert.
//

func newSimulateTestOrchestrator(t *testing.T, fn *fakeNode) *Orchestrator {
	t.Helper()
	privKey, _, err := crypto.GenerateKeyPair()
	if err != nil {
		t.Fatalf("failed to generate a key pair: %v", err)
	}
	return &Orchestrator{
		mainchainID:           big.NewInt(testMainchainID),
		subchainID:            big.NewInt(testSubchainID),
		mainchainEthRpcClient: fn.client(t),
		privateKey:            privKey,
	}
}

func testRelayTx() *types.Transaction {
	return types.NewTransaction(1, testBankAddr, big.NewInt(0), 1_000_000, big.NewInt(1), []byte{0x01, 0x02, 0x03, 0x04})
}

// The regression this guards: Theta's eth_call retries internally with a one-block
// sleep between attempts, so an unbounded dry run of a reverting relay blocks the
// orchestrator for the better part of a minute. The tick is sequential across all
// four asset classes in both directions, so that starves every relay behind it.
func TestSimulateAndSendTimesOutInsteadOfBlocking(t *testing.T) {
	assert := assert.New(t)

	fn := newFakeNode(t, func(_ string, _ []json.RawMessage) (string, string) {
		return "0x", ""
	})
	fn.delay = 30 * time.Second // a node that will not answer in time
	oc := newSimulateTestOrchestrator(t, fn)

	viper.Set(scom.CfgSubchainRelayDryRunTimeoutInSeconds, 1)
	defer viper.Set(scom.CfgSubchainRelayDryRunTimeoutInSeconds, 5)

	start := time.Now()
	_, err := oc.simulateAndSend(big.NewInt(testMainchainID), testRelayTx())
	elapsed := time.Since(start)

	assert.NotNil(err)
	assert.ErrorIs(err, ErrDryRunTimedOut)
	assert.NotErrorIs(err, ErrTxWouldRevert, "a timeout must not be reported as a revert")
	assert.Less(elapsed, 10*time.Second, "the dry run must give up rather than block the processing tick")
}

// A revert and a timeout mean different things -- "this can never succeed" versus
// "we do not know" -- and during an incident that distinction is what separates a
// dead relay path from an RPC outage.
func TestSimulateAndSendDistinguishesRevertFromTimeout(t *testing.T) {
	assert := assert.New(t)

	fn := newFakeNode(t, func(_ string, _ []json.RawMessage) (string, string) {
		return "", "execution reverted: Cannot unlock the requested amount of TFuel"
	})
	oc := newSimulateTestOrchestrator(t, fn)

	_, err := oc.simulateAndSend(big.NewInt(testMainchainID), testRelayTx())
	assert.NotNil(err)
	assert.ErrorIs(err, ErrTxWouldRevert)
	assert.NotErrorIs(err, ErrDryRunTimedOut)
	assert.Contains(err.Error(), "Cannot unlock the requested amount of TFuel",
		"the revert reason must survive into the log")
}

// A relay that would revert must never be broadcast -- that is what turned the
// permanently stuck nonce into a rebroadcast loop that burned gas every retry.
func TestSimulateAndSendDoesNotBroadcastARevertingRelay(t *testing.T) {
	var sent int64
	fn := newFakeNode(t, func(method string, _ []json.RawMessage) (string, string) {
		if strings.Contains(method, "sendRawTransaction") || strings.Contains(method, "sendTransaction") {
			atomic.AddInt64(&sent, 1)
			return "0x", ""
		}
		return "", "execution reverted"
	})
	oc := newSimulateTestOrchestrator(t, fn)

	_, err := oc.simulateAndSend(big.NewInt(testMainchainID), testRelayTx())
	assert.NotNil(t, err)
	assert.Zero(t, atomic.LoadInt64(&sent), "a reverting relay must not be broadcast")
}

// The happy path still broadcasts and reports the hash.
func TestSimulateAndSendBroadcastsWhenTheDryRunSucceeds(t *testing.T) {
	assert := assert.New(t)

	var sent int64
	fn := newFakeNode(t, func(method string, _ []json.RawMessage) (string, string) {
		if strings.Contains(method, "sendRawTransaction") || strings.Contains(method, "sendTransaction") {
			atomic.AddInt64(&sent, 1)
		}
		return "0x", ""
	})
	oc := newSimulateTestOrchestrator(t, fn)

	tx := testRelayTx()
	txHash, err := oc.simulateAndSend(big.NewInt(testMainchainID), tx)
	assert.Nil(err)
	assert.Equal(tx.Hash(), txHash)
	assert.EqualValues(1, atomic.LoadInt64(&sent), "the relay must be broadcast exactly once")
}

// A missing or nonsensical timeout must fall back to the default rather than
// degenerating into no timeout at all.
func TestSimulateAndSendFallsBackToDefaultTimeout(t *testing.T) {
	fn := newFakeNode(t, func(_ string, _ []json.RawMessage) (string, string) {
		return "0x", ""
	})
	oc := newSimulateTestOrchestrator(t, fn)

	viper.Set(scom.CfgSubchainRelayDryRunTimeoutInSeconds, 0)
	defer viper.Set(scom.CfgSubchainRelayDryRunTimeoutInSeconds, 5)

	_, err := oc.simulateAndSend(big.NewInt(testMainchainID), testRelayTx())
	assert.Nil(t, err, "a zero configured timeout must fall back to the default, not fail")
}

//
// Wiring. The checks above exercise the guards directly; these confirm the guards
// are actually installed on the vote paths. A guard nothing calls is worthless, and
// nothing else in the suite would notice its removal.
//

// stubLedger supplies the dynasty and nothing else. The embedded nil interface
// panics if the code under test reaches for anything further, which is the signal
// we want rather than a silent zero value.
type stubLedger struct {
	score.Ledger
	dynasty *big.Int
}

func (s *stubLedger) GetDynasty() *big.Int { return s.dynasty }

func tnt20BurnEvent(t *testing.T, denom string, amount *big.Int) *score.InterChainMessageEvent {
	t.Helper()
	contractAbi, err := abi.JSON(strings.NewReader(string(scta.TNT20TokenBankABI)))
	if err != nil {
		t.Fatalf("failed to parse the TNT20TokenBank ABI: %v", err)
	}
	data, err := contractAbi.Events["TNT20VoucherBurned"].Inputs.Pack(
		denom,
		common.HexToAddress("0x8556d1a34013feb0cba9349636895d94831528a0"),
		common.HexToAddress("0x4c4f17bdde7dacc6b118b159ed602915a2031da3"),
		amount,
		big.NewInt(3698),
	)
	if err != nil {
		t.Fatalf("failed to encode the TNT20VoucherBurned event: %v", err)
	}
	return &score.InterChainMessageEvent{
		Type:          score.IMCEventTypeCrossChainVoucherBurnTNT20,
		SourceChainID: big.NewInt(testSubchainID),
		TargetChainID: big.NewInt(testMainchainID),
		Nonce:         big.NewInt(3698),
		BlockHeight:   big.NewInt(49741664),
		Data:          data,
	}
}

func newWiringTestOrchestrator(t *testing.T, fn *fakeNode) (*Orchestrator, *bind.TransactOpts) {
	t.Helper()
	oc := newTestOrchestrator(t, fn)
	privKey, _, err := crypto.GenerateKeyPair()
	if err != nil {
		t.Fatalf("failed to generate a key pair: %v", err)
	}
	oc.privateKey = privKey
	oc.ledger = &stubLedger{dynasty: big.NewInt(1)}

	bank, err := scta.NewTNT20TokenBank(testBankAddr, oc.mainchainEthRpcClient)
	if err != nil {
		t.Fatalf("failed to bind the TNT20TokenBank: %v", err)
	}
	oc.mainchainTNT20TokenBank = bank

	txOpts, err := bind.NewKeyedTransactorWithChainID(privKey, big.NewInt(testMainchainID))
	if err != nil {
		t.Fatalf("failed to build tx opts: %v", err)
	}
	txOpts.GasLimit = 1_000_000
	txOpts.GasPrice = big.NewInt(1)
	txOpts.Value = big.NewInt(0)
	txOpts.Nonce = big.NewInt(0)
	txOpts.NoSend = true
	return oc, txOpts
}

// This is the incident, replayed through the real vote path: a burn event asking
// for 8,200 against a bank holding 50. It must be refused before any transaction is
// built, let alone broadcast.
func TestUnlockTNT20TokensRefusesAnUncollateralizedRelay(t *testing.T) {
	assert := assert.New(t)

	var broadcasts int64
	fn := newFakeNode(t, func(method string, _ []json.RawMessage) (string, string) {
		if strings.Contains(method, "sendRawTransaction") || strings.Contains(method, "sendTransaction") {
			atomic.AddInt64(&broadcasts, 1)
			return "0x", ""
		}
		return word(big.NewInt(50)), "" // the bank's actual holdings
	})
	oc, txOpts := newWiringTestOrchestrator(t, fn)

	event := tnt20BurnEvent(t, testTNT20Denom, big.NewInt(8_200))
	_, err := oc.unlockTNT20Tokens(txOpts, big.NewInt(testMainchainID), event)

	assert.NotNil(err)
	assert.ErrorIs(err, ErrUnlockUncollateralized,
		"the collateral guard must be installed on the TNT20 vote path, not merely defined")
	assert.Zero(atomic.LoadInt64(&broadcasts), "an uncollateralized relay must never be broadcast")
}

// The mirror image: a properly collateralized relay must still get through, so the
// guard cannot be blamed for a stalled pipeline.
func TestUnlockTNT20TokensAllowsACollateralizedRelay(t *testing.T) {
	fn := newFakeNode(t, func(_ string, _ []json.RawMessage) (string, string) {
		return word(big.NewInt(100_000)), ""
	})
	oc, txOpts := newWiringTestOrchestrator(t, fn)

	event := tnt20BurnEvent(t, testTNT20Denom, big.NewInt(8_200))
	_, err := oc.unlockTNT20Tokens(txOpts, big.NewInt(testMainchainID), event)

	// The fake node is not a real chain, so the relay may still fail further down;
	// what matters is that it was not stopped by the collateral guard.
	if err != nil {
		assert.NotErrorIs(t, err, ErrUnlockUncollateralized,
			"a fully collateralized relay must not be refused by the guard")
	}
}

//
// Bounded hot-path reads. The dry run was bounded first, but the collateral and
// nonce reads run *before* it on every tick, so leaving those unbounded left the
// starvation they were meant to prevent fully intact.
//

// The guard must give up on an unresponsive node rather than hold the sequential
// processing tick open for as long as that node cares to stay silent.
func TestCollateralGuardIsBoundedAgainstAnUnresponsiveNode(t *testing.T) {
	assert := assert.New(t)

	fn := newFakeNode(t, func(_ string, _ []json.RawMessage) (string, string) {
		return word(big.NewInt(10_000)), ""
	})
	fn.delay = 30 * time.Second
	oc := newTestOrchestrator(t, fn)

	viper.Set(scom.CfgSubchainRelayDryRunTimeoutInSeconds, 1)
	defer viper.Set(scom.CfgSubchainRelayDryRunTimeoutInSeconds, 5)

	start := time.Now()
	err := oc.verifyTNT20UnlockCollateral(big.NewInt(testMainchainID), testTNT20Denom, big.NewInt(1))
	elapsed := time.Since(start)

	assert.NotNil(err, "an unreadable balance must fail closed")
	assert.Less(elapsed, 10*time.Second, "the collateral read must be bounded, not merely the dry run")
}

// Same for the first RPC call made on every relay path. If this one is unbounded the
// tick never even reaches the guard or the dry run.
func TestMaxProcessedNonceReadIsBounded(t *testing.T) {
	assert := assert.New(t)

	fn := newFakeNode(t, func(_ string, _ []json.RawMessage) (string, string) {
		return word(big.NewInt(1)), ""
	})
	fn.delay = 30 * time.Second
	oc := newTestOrchestrator(t, fn)

	bank, err := scta.NewTFuelTokenBank(testBankAddr, oc.mainchainEthRpcClient)
	if err != nil {
		t.Fatalf("failed to bind the TFuelTokenBank: %v", err)
	}

	viper.Set(scom.CfgSubchainRelayDryRunTimeoutInSeconds, 1)
	defer viper.Set(scom.CfgSubchainRelayDryRunTimeoutInSeconds, 5)

	opts, cancel := boundedCallOpts()
	defer cancel()

	start := time.Now()
	_, err = bank.GetMaxProcessedVoucherBurnNonce(opts, big.NewInt(testSubchainID))
	elapsed := time.Since(start)

	assert.NotNil(err)
	assert.Less(elapsed, 10*time.Second, "boundedCallOpts must carry a deadline into the binding")
}
