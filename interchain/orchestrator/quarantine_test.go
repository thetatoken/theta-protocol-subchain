package orchestrator

import (
	"math/big"
	"testing"
	"time"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	scom "github.com/thetatoken/thetasubchain/common"
	score "github.com/thetatoken/thetasubchain/core"
)

func quarantineTestEvent(nonce int64) *score.InterChainMessageEvent {
	return &score.InterChainMessageEvent{
		Type:          score.IMCEventTypeCrossChainVoucherBurnTFuel,
		SourceChainID: big.NewInt(testSubchainID),
		TargetChainID: big.NewInt(testMainchainID),
		Nonce:         big.NewInt(nonce),
		BlockHeight:   big.NewInt(1000),
	}
}

func newQuarantineOrchestrator() *Orchestrator {
	return &Orchestrator{
		mainchainID:           big.NewInt(testMainchainID),
		subchainID:            big.NewInt(testSubchainID),
		firstContradictedTime: make(map[string]time.Time),
	}
}

// A contradiction seen once must not discard the event. The corroboration read is
// served from latest state, so a lagging backend produces the same verdict as a
// forgery, and discarding is irreversible: the scan window has moved past the block
// the event came from and nothing rescans it.
func TestQuarantineDoesNotDiscardOnFirstContradiction(t *testing.T) {
	oc := newQuarantineOrchestrator()
	assert.False(t, oc.quarantineExpired(quarantineTestEvent(288)),
		"the first contradiction must start the clock, not discard the event")
}

// Sustained contradiction past the window is a forgery, not a blip, and may be
// discarded.
func TestQuarantineExpiresAfterTheConfiguredPeriod(t *testing.T) {
	assert := assert.New(t)
	oc := newQuarantineOrchestrator()
	event := quarantineTestEvent(288)

	viper.Set(scom.CfgSubchainUncorroboratedEventQuarantineInSeconds, 3600)
	defer viper.Set(scom.CfgSubchainUncorroboratedEventQuarantineInSeconds, 1800)

	assert.False(oc.quarantineExpired(event))
	assert.False(oc.quarantineExpired(event), "still inside the window")

	// Backdate the first sighting to just beyond the window.
	oc.firstContradictedTime[event.ID()] = time.Now().Add(-2 * time.Hour)
	assert.True(oc.quarantineExpired(event))
}

// The window must measure a *continuous* disagreement. A blip that resolves has to
// reset the clock, otherwise unrelated transient failures accumulate over hours and
// eventually discard a healthy event.
func TestQuarantineClockResetsWhenTheContradictionResolves(t *testing.T) {
	assert := assert.New(t)
	oc := newQuarantineOrchestrator()
	event := quarantineTestEvent(288)

	oc.quarantineExpired(event)
	oc.firstContradictedTime[event.ID()] = time.Now().Add(-29 * time.Minute)

	oc.clearContradicted(event) // the event verified successfully this round
	assert.Empty(oc.firstContradictedTime)

	// A later contradiction starts a fresh window rather than inheriting the old one.
	assert.False(oc.quarantineExpired(event))
}

// Quarantine state must be per event, so one bad event cannot age out a neighbour.
func TestQuarantineIsTrackedPerEvent(t *testing.T) {
	assert := assert.New(t)
	oc := newQuarantineOrchestrator()

	stale, fresh := quarantineTestEvent(288), quarantineTestEvent(289)
	oc.quarantineExpired(stale)
	oc.firstContradictedTime[stale.ID()] = time.Now().Add(-2 * time.Hour)

	assert.True(oc.quarantineExpired(stale))
	assert.False(oc.quarantineExpired(fresh), "a neighbouring nonce must start its own clock")
}

// A missing or nonsensical setting must fall back to the default rather than
// degenerating into "discard immediately".
func TestQuarantinePeriodFallsBackToDefault(t *testing.T) {
	oc := newQuarantineOrchestrator()
	viper.Set(scom.CfgSubchainUncorroboratedEventQuarantineInSeconds, 0)
	defer viper.Set(scom.CfgSubchainUncorroboratedEventQuarantineInSeconds, 1800)

	assert.Equal(t, defaultUncorroboratedEventQuarantine, oc.getQuarantinePeriod())
}
