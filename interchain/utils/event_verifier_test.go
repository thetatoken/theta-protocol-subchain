package core

import (
	"errors"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thetatoken/thetasubchain/eth/abi/bind"

	score "github.com/thetatoken/thetasubchain/core"
)

// fakeTokenBank stands in for the generated TokenBank accessors. The maps are
// keyed by "chainID/nonce", mirroring the event height maps in TokenBank.sol.
type fakeTokenBank struct {
	lockHeights map[string]*big.Int
	burnHeights map[string]*big.Int
	err         error
}

func heightKey(chainID *big.Int, nonce *big.Int) string {
	return chainID.String() + "/" + nonce.String()
}

func (f *fakeTokenBank) GetTokenLockEventHeight(opts *bind.CallOpts, chainID *big.Int, eventNonce *big.Int) (*big.Int, error) {
	if f.err != nil {
		return nil, f.err
	}
	if h, ok := f.lockHeights[heightKey(chainID, eventNonce)]; ok {
		return h, nil
	}
	return big.NewInt(0), nil // unset mapping entry
}

func (f *fakeTokenBank) GetVoucherBurnEventHeight(opts *bind.CallOpts, chainID *big.Int, eventNonce *big.Int) (*big.Int, error) {
	if f.err != nil {
		return nil, f.err
	}
	if h, ok := f.burnHeights[heightKey(chainID, eventNonce)]; ok {
		return h, nil
	}
	return big.NewInt(0), nil
}

const (
	testSubchainID  = 9001
	testMainchainID = 9000
)

func burnEvent(nonce int64, blockHeight int64) *score.InterChainMessageEvent {
	return &score.InterChainMessageEvent{
		Type:          score.IMCEventTypeCrossChainVoucherBurnTFuel,
		SourceChainID: big.NewInt(testSubchainID),
		TargetChainID: big.NewInt(testMainchainID),
		Nonce:         big.NewInt(nonce),
		BlockHeight:   big.NewInt(blockHeight),
	}
}

func lockEvent(nonce int64, blockHeight int64) *score.InterChainMessageEvent {
	return &score.InterChainMessageEvent{
		Type:          score.IMCEventTypeCrossChainTokenLockTFuel,
		SourceChainID: big.NewInt(testMainchainID),
		TargetChainID: big.NewInt(testSubchainID),
		Nonce:         big.NewInt(nonce),
		BlockHeight:   big.NewInt(blockHeight),
	}
}

// The genuine event: the TokenBank recorded the nonce at the very block the log
// was found in.
func TestVerifyAcceptsCorroboratedEvent(t *testing.T) {
	assert := assert.New(t)

	bank := &fakeTokenBank{burnHeights: map[string]*big.Int{"9000/42": big.NewInt(1000)}}
	assert.Nil(VerifyEventAgainstTokenBankState(bank, burnEvent(42, 1000)))

	bank = &fakeTokenBank{lockHeights: map[string]*big.Int{"9001/17": big.NewInt(2000)}}
	assert.Nil(VerifyEventAgainstTokenBankState(bank, lockEvent(17, 2000)))
}

// A log whose frame reverted: the nonce was never committed, so the height the
// bank recorded for it belongs to the burn that actually took that nonce, in a
// different block.
func TestVerifyRejectsEventFromRevertedFrame(t *testing.T) {
	assert := assert.New(t)

	bank := &fakeTokenBank{burnHeights: map[string]*big.Int{"9000/42": big.NewInt(1050)}}

	err := VerifyEventAgainstTokenBankState(bank, burnEvent(42, 1000))
	assert.NotNil(err)
	assert.True(IsEventNotCorroborated(err))
}

// A nonce the bank has no record of at all.
func TestVerifyRejectsUnknownNonce(t *testing.T) {
	assert := assert.New(t)

	err := VerifyEventAgainstTokenBankState(&fakeTokenBank{}, burnEvent(999, 1000))
	assert.NotNil(err)
	assert.True(IsEventNotCorroborated(err))
}

// A burn must be checked against the voucher burn map, not the token lock map:
// consulting the wrong one would let a burn borrow a lock's recorded height.
func TestVerifyDoesNotCrossTheLockAndBurnMaps(t *testing.T) {
	assert := assert.New(t)

	bank := &fakeTokenBank{lockHeights: map[string]*big.Int{"9000/42": big.NewInt(1000)}}

	err := VerifyEventAgainstTokenBankState(bank, burnEvent(42, 1000))
	assert.NotNil(err)
	assert.True(IsEventNotCorroborated(err))
}

// An unreachable RPC must be reported as "cannot check", never as "contradicted":
// the witness keeps such events and the orchestrator retries them later.
func TestVerifyReportsUnavailableSeparatelyFromContradicted(t *testing.T) {
	assert := assert.New(t)

	bank := &fakeTokenBank{err: errors.New("connection refused")}

	err := VerifyEventAgainstTokenBankState(bank, burnEvent(42, 1000))
	assert.NotNil(err)
	assert.False(IsEventNotCorroborated(err))
	assert.True(errors.Is(err, ErrEventVerificationUnavailable))

	err = VerifyEventAgainstTokenBankState(nil, burnEvent(42, 1000))
	assert.NotNil(err)
	assert.False(IsEventNotCorroborated(err))
	assert.True(errors.Is(err, ErrEventVerificationUnavailable))
}

// Voucher mint and token unlock events are terminal: they are not relayed
// anywhere, so there is no event height map to check them against.
func TestVerifySkipsNonRelayedEventTypes(t *testing.T) {
	assert := assert.New(t)

	event := burnEvent(42, 1000)
	event.Type = score.IMCEventTypeCrossChainVoucherMintTFuel
	assert.Nil(VerifyEventAgainstTokenBankState(&fakeTokenBank{}, event))

	event.Type = score.IMCEventTypeCrossChainTokenUnlockTNT20
	assert.Nil(VerifyEventAgainstTokenBankState(&fakeTokenBank{}, event))
}

// A malformed event (missing nonce, height, or target chain) must be rejected
// rather than silently skipped.
func TestVerifyRejectsIncompleteEvent(t *testing.T) {
	assert := assert.New(t)

	bank := &fakeTokenBank{burnHeights: map[string]*big.Int{"9000/42": big.NewInt(1000)}}

	event := burnEvent(42, 1000)
	event.TargetChainID = nil
	assert.True(IsEventNotCorroborated(VerifyEventAgainstTokenBankState(bank, event)))

	event = burnEvent(42, 1000)
	event.BlockHeight = nil
	assert.True(IsEventNotCorroborated(VerifyEventAgainstTokenBankState(bank, event)))

	event = burnEvent(42, 1000)
	event.Nonce = nil
	assert.True(IsEventNotCorroborated(VerifyEventAgainstTokenBankState(bank, event)))
}
