package core

import (
	"context"
	"fmt"
	"math/big"
	"time"

	ec "github.com/thetatoken/thetasubchain/eth/ethclient"
)

// chainIDQueryTimeout bounds the identity check so a stalled endpoint fails startup
// rather than hanging it.
const chainIDQueryTimeout = 15 * time.Second

// VerifyChainID confirms an RPC endpoint really serves the chain we think it does.
//
// Both the witness and the orchestrator previously dialled whatever URL the config
// named and trusted it. Predeploy addresses are identical across subchains, so a URL
// pointing at the wrong subchain yields TokenBank contracts that answer plausibly at
// the expected addresses. Events read from it would then be labelled with *this*
// chain's ID, corroborated against that same wrong endpoint, and relayed as if genuine
// -- a misconfiguration that produces forged-looking cross-chain transfers without any
// attacker involved.
//
// The check is cheap, happens once at startup, and turns that class of misconfiguration
// into a refusal to start.
func VerifyChainID(client *ec.Client, expected *big.Int, label string) error {
	if client == nil {
		return fmt.Errorf("%v: no ETH RPC client to verify", label)
	}
	if expected == nil || expected.Sign() <= 0 {
		return fmt.Errorf("%v: no valid expected chain ID configured (got %v)", label, expected)
	}

	ctx, cancel := context.WithTimeout(context.Background(), chainIDQueryTimeout)
	defer cancel()

	actual, err := client.ChainID(ctx)
	if err != nil {
		return fmt.Errorf("%v: failed to read eth_chainId: %v", label, err)
	}
	if actual == nil || actual.Cmp(expected) != 0 {
		return fmt.Errorf("%v: RPC endpoint serves chain %v but %v was expected -- "+
			"refusing to start, as reading inter-chain events from the wrong chain would "+
			"relay them under this chain's identity", label, actual, expected)
	}
	return nil
}
