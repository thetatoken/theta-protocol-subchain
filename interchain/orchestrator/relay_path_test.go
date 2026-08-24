package orchestrator

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	scom "github.com/thetatoken/thetasubchain/common"
	score "github.com/thetatoken/thetasubchain/core"
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
