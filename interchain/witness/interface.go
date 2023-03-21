package witness

import (
	"context"
	"math/big"

	score "github.com/thetatoken/thetasubchain/core"
	siu "github.com/thetatoken/thetasubchain/interchain/utils"
)

type ChainWitness interface {
	Start(ctx context.Context)
	Stop()
	Wait()
	GetMainchainBlockHeight() (*big.Int, error)
	GetValidatorSetByDynasty(dynasty *big.Int) (*score.ValidatorSet, error)
	GetSubchainRegistrationHeight() (*big.Int, error)
	GetInterChainEventCache() *siu.InterChainEventCache
}
