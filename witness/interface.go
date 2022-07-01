package witness

import (
	"context"
	"math/big"

	score "github.com/thetatoken/thetasubchain/core"
)

type ChainWitness interface {
	Start(ctx context.Context)
	Stop()
	Wait()
	GetMainchainBlockHeight() (*big.Int, error)
	GetValidatorSetByDynasty(dynasty *big.Int) (*score.ValidatorSet, error)
	GetInterChainEventCache() *score.InterChainEventCache
}
