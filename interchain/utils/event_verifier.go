package core

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/spf13/viper"
	"github.com/thetatoken/thetasubchain/eth/abi/bind"

	scom "github.com/thetatoken/thetasubchain/common"
	score "github.com/thetatoken/thetasubchain/core"
)

// defaultVerifierTimeout bounds the corroboration read when no value is configured.
const defaultVerifierTimeout = 5 * time.Second

// verifierCallOpts bounds the corroboration read.
//
// Theta's eth_call retries internally with a one-block sleep between attempts, so an
// unbounded read against an unresponsive node holds the caller for tens of seconds.
// Both callers sit on a sequential loop, so that delay is paid by every event behind
// this one. A read that times out surfaces as ErrEventVerificationUnavailable rather
// than as a contradiction, which is the distinction the callers act on.
func verifierCallOpts() (*bind.CallOpts, context.CancelFunc) {
	timeout := time.Duration(viper.GetInt(scom.CfgSubchainRelayDryRunTimeoutInSeconds)) * time.Second
	if timeout <= 0 {
		timeout = defaultVerifierTimeout
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	return &bind.CallOpts{Context: ctx}, cancel
}

var (
	// ErrEventNotCorroborated indicates that the source chain's TokenBank contract
	// has no record of the event that eth_getLogs reported. Such an event MUST NOT
	// be relayed.
	ErrEventNotCorroborated = errors.New("inter-chain event is not corroborated by the source chain TokenBank state")

	// ErrEventVerificationUnavailable indicates that the cross-check could not be
	// performed (e.g. the source chain RPC is unreachable). The event is neither
	// proven nor disproven, so the caller should retry rather than relay or drop it.
	ErrEventVerificationUnavailable = errors.New("unable to verify the inter-chain event against the source chain state")
)

// TokenBankEventHeightReader is the subset of the generated TokenBank accessors
// needed to corroborate an inter-chain event. All four TokenBank flavors
// (TFuel/TNT20/TNT721/TNT1155) satisfy it.
type TokenBankEventHeightReader interface {
	GetTokenLockEventHeight(opts *bind.CallOpts, chainID *big.Int, eventNonce *big.Int) (*big.Int, error)
	GetVoucherBurnEventHeight(opts *bind.CallOpts, chainID *big.Int, eventNonce *big.Int) (*big.Int, error)
}

// VerifyEventAgainstTokenBankState cross-checks an event obtained from eth_getLogs
// against the emitting TokenBank's own storage.
//
// Rationale: a log is not by itself proof that the state transition it describes
// was committed. Relaying an event whose state changes did not survive would
// release tokens on the target chain against nothing on the source chain, so the
// event is confirmed against contract state before it is acted upon.
//
// Every path that increments a token lock or voucher burn nonce also writes
// block.number into the corresponding event height map (_incrementTokenLockNonce
// and _incrementVoucherBurnNonce in TokenBank.sol). Those writes live in the
// state trie. Requiring the recorded height to match the height of the block the
// log was found in therefore confirms the event against committed state rather
// than trusting the log alone.
//
// The check is nonce-anchored as well: the map is keyed by (targetChainID,
// nonce), and a nonce is assigned at most once.
func VerifyEventAgainstTokenBankState(bank TokenBankEventHeightReader, event *score.InterChainMessageEvent) error {
	if bank == nil {
		return fmt.Errorf("%w: no TokenBank accessor for event type %v", ErrEventVerificationUnavailable, event.Type)
	}
	if event.Nonce == nil || event.BlockHeight == nil || event.TargetChainID == nil {
		return fmt.Errorf("%w: event is missing the nonce, the block height, or the target chain ID", ErrEventNotCorroborated)
	}

	var recordedHeight *big.Int
	var err error

	opts, cancel := verifierCallOpts()
	defer cancel()

	switch event.Type {
	case score.IMCEventTypeCrossChainTokenLockTFuel, score.IMCEventTypeCrossChainTokenLockTNT20,
		score.IMCEventTypeCrossChainTokenLockTNT721, score.IMCEventTypeCrossChainTokenLockTNT1155:
		// tokenLockEventHeightMap is keyed by the target chain, matching the
		// targetChainID carried by the TokenLocked event.
		recordedHeight, err = bank.GetTokenLockEventHeight(opts, event.TargetChainID, event.Nonce)

	case score.IMCEventTypeCrossChainVoucherBurnTFuel, score.IMCEventTypeCrossChainVoucherBurnTNT20,
		score.IMCEventTypeCrossChainVoucherBurnTNT721, score.IMCEventTypeCrossChainVoucherBurnTNT1155:
		// voucherBurnEventHeightMap is keyed by the chain the tokens return to,
		// which the witness derives from the denom of the VoucherBurned event.
		recordedHeight, err = bank.GetVoucherBurnEventHeight(opts, event.TargetChainID, event.Nonce)

	default:
		// Voucher mint and token unlock events are never relayed onwards, so
		// there is nothing to corroborate.
		return nil
	}

	if err != nil {
		return fmt.Errorf("%w: %v", ErrEventVerificationUnavailable, err)
	}
	if recordedHeight == nil || recordedHeight.Sign() == 0 {
		return fmt.Errorf("%w: the TokenBank on chain %v has no record of a type %v event with nonce %v (log claims block %v)",
			ErrEventNotCorroborated, event.SourceChainID, event.Type, event.Nonce, event.BlockHeight)
	}
	if recordedHeight.Cmp(event.BlockHeight) != 0 {
		return fmt.Errorf("%w: the TokenBank on chain %v recorded the type %v event with nonce %v at block %v, but the log was found at block %v",
			ErrEventNotCorroborated, event.SourceChainID, event.Type, event.Nonce, recordedHeight, event.BlockHeight)
	}

	return nil
}

// IsEventNotCorroborated reports whether the error means the event was
// contradicted by the source chain state, as opposed to the check being
// unavailable.
func IsEventNotCorroborated(err error) bool {
	return errors.Is(err, ErrEventNotCorroborated)
}
