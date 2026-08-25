package orchestrator

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/thetatoken/theta/store/database/backend"
	scom "github.com/thetatoken/thetasubchain/common"
	score "github.com/thetatoken/thetasubchain/core"
	scta "github.com/thetatoken/thetasubchain/interchain/contracts/accessors"
	siu "github.com/thetatoken/thetasubchain/interchain/utils"
)

// Every relayed event type must map to an asset name, otherwise the operator's
// switch for that asset would silently do nothing.
func TestRelayAssetNameCoversEveryRelayedEventType(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range []struct {
		eventType score.InterChainMessageEventType
		asset     string
	}{
		{score.IMCEventTypeCrossChainTokenLockTFuel, "tfuel"},
		{score.IMCEventTypeCrossChainVoucherBurnTFuel, "tfuel"},
		{score.IMCEventTypeCrossChainTokenLockTNT20, "tnt20"},
		{score.IMCEventTypeCrossChainVoucherBurnTNT20, "tnt20"},
		{score.IMCEventTypeCrossChainTokenLockTNT721, "tnt721"},
		{score.IMCEventTypeCrossChainVoucherBurnTNT721, "tnt721"},
		{score.IMCEventTypeCrossChainTokenLockTNT1155, "tnt1155"},
		{score.IMCEventTypeCrossChainVoucherBurnTNT1155, "tnt1155"},
	} {
		assert.Equal(tc.asset, relayAssetName(tc.eventType), "event type %v", tc.eventType)
	}
}

// An unconfigured path stays enabled, so upgrading a node without touching its
// config must not silently stop it relaying.
func TestRelayPathEnabledByDefault(t *testing.T) {
	oc := &Orchestrator{mainchainID: big.NewInt(361), subchainID: big.NewInt(360890)}
	assert.True(t, oc.relayPathEnabled(big.NewInt(360890), score.IMCEventTypeCrossChainVoucherBurnTFuel))
}

// Suspending one asset in one direction must leave every other path running --
// that is the whole point of the switch being per-path.
func TestRelayPathSuspendsOnlyTheConfiguredPath(t *testing.T) {
	assert := assert.New(t)

	mainchainID, subchainID := big.NewInt(361), big.NewInt(360890)
	oc := &Orchestrator{mainchainID: mainchainID, subchainID: subchainID}

	key := fmt.Sprintf("%v.tfuel.outbound", scom.CfgSubchainRelayPathPrefix)
	viper.Set(key, false)
	defer viper.Set(key, true)

	// TFuel subchain -> main chain is closed.
	assert.False(oc.relayPathEnabled(subchainID, score.IMCEventTypeCrossChainVoucherBurnTFuel))
	// The opposite direction for the same asset is untouched.
	assert.True(oc.relayPathEnabled(mainchainID, score.IMCEventTypeCrossChainTokenLockTFuel))
	// Other asset classes in the same direction are untouched.
	assert.True(oc.relayPathEnabled(subchainID, score.IMCEventTypeCrossChainVoucherBurnTNT20))
	assert.True(oc.relayPathEnabled(subchainID, score.IMCEventTypeCrossChainVoucherBurnTNT721))
}

//
// Zero-RPC regression.
//
// The tests above call relayPathEnabled() directly. That is exactly the weakness that
// let the gate sit in the wrong place unnoticed: it was inside processNextEvent(),
// but every path reads GetMaxProcessed...Nonce before calling it, so a disabled path
// still spent an RPC round trip and its timeout on every tick. Asserting on the
// predicate proves nothing about where it is invoked from -- only driving the real
// entry point does.
//

func newRelayPathOrchestrator(t *testing.T, fn *fakeNode) *Orchestrator {
	t.Helper()
	client := fn.client(t)

	tfuelBank, err := scta.NewTFuelTokenBank(testBankAddr, client)
	if err != nil {
		t.Fatalf("failed to bind the TFuelTokenBank: %v", err)
	}
	tnt20Bank, err := scta.NewTNT20TokenBank(testBankAddr, client)
	if err != nil {
		t.Fatalf("failed to bind the TNT20TokenBank: %v", err)
	}

	return &Orchestrator{
		mainchainID:             big.NewInt(testMainchainID),
		subchainID:              big.NewInt(testSubchainID),
		mainchainEthRpcClient:   client,
		subchainEthRpcClient:    client,
		mainchainTFuelTokenBank: tfuelBank,
		subchainTFuelTokenBank:  tfuelBank,
		mainchainTNT20TokenBank: tnt20Bank,
		subchainTNT20TokenBank:  tnt20Bank,
		eventProcessedTime:      make(map[string]time.Time),
		firstContradictedTime:   make(map[string]time.Time),
		interChainEventCache:    siu.NewInterChainEventCache(backend.NewMemDatabase()),
	}
}

// A suspended path must not touch the network at all.
func TestDisabledRelayPathMakesNoRpcCalls(t *testing.T) {
	assert := assert.New(t)

	fn := newFakeNode(t, func(_ string, _ []json.RawMessage) (string, string) {
		return "0x" + strings.Repeat("0", 64), ""
	})
	oc := newRelayPathOrchestrator(t, fn)

	key := fmt.Sprintf("%v.tfuel.outbound", scom.CfgSubchainRelayPathPrefix)
	viper.Set(key, false)
	defer viper.Set(key, true)

	oc.processNextTFuelVoucherBurnEvent(big.NewInt(testSubchainID), big.NewInt(testMainchainID))

	assert.Zero(atomic.LoadInt64(&fn.calls),
		"a disabled relay path must not issue any RPC call, not even the max-processed-nonce read")
}

// Positive control: with the path enabled the same entry point does reach the
// network. Without this, the test above would still pass if the whole path were
// broken rather than merely gated.
func TestEnabledRelayPathDoesReachTheNetwork(t *testing.T) {
	fn := newFakeNode(t, func(_ string, _ []json.RawMessage) (string, string) {
		return "0x" + strings.Repeat("0", 64), ""
	})
	oc := newRelayPathOrchestrator(t, fn)

	key := fmt.Sprintf("%v.tfuel.outbound", scom.CfgSubchainRelayPathPrefix)
	viper.Set(key, true)
	defer viper.Set(key, true)

	oc.processNextTFuelVoucherBurnEvent(big.NewInt(testSubchainID), big.NewInt(testMainchainID))

	assert.NotZero(t, atomic.LoadInt64(&fn.calls),
		"an enabled relay path must still query the target chain")
}

// Every other asset must be unaffected by one path being suspended.
func TestDisablingOnePathLeavesOthersReachingTheNetwork(t *testing.T) {
	fn := newFakeNode(t, func(_ string, _ []json.RawMessage) (string, string) {
		return "0x" + strings.Repeat("0", 64), ""
	})
	oc := newRelayPathOrchestrator(t, fn)

	key := fmt.Sprintf("%v.tfuel.outbound", scom.CfgSubchainRelayPathPrefix)
	viper.Set(key, false)
	defer viper.Set(key, true)

	oc.processNextTNT20VoucherBurnEvent(big.NewInt(testSubchainID), big.NewInt(testMainchainID))

	assert.NotZero(t, atomic.LoadInt64(&fn.calls),
		"suspending TFuel must not suspend TNT20")
}
