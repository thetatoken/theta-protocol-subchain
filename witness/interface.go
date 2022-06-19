package witness

import (
	"context"
	"math/big"

	score "github.com/thetatoken/thetasubchain/core"
	"github.com/thetatoken/thetasubchain/eth/abi/bind"
)

type ChainWitness interface {
	Start(ctx context.Context)
	Stop()
	Wait()
	GetMainchainBlockNumber() (*big.Int, error)
	GetMainchainBlockNumberUint() (uint64, error)
	GetValidatorSetByDynasty(dynasty *big.Int) (*score.ValidatorSet, error)
	GetInterChainEventCache() *score.InterChainEventCache
	CallMintOnMainchain(auth *bind.TransactOpts)
}
