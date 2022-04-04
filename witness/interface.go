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
	GetMainchainBlockNumber() (*big.Int, error)
	GetMainchainBlockNumberUint() (uint64, error)
	GetValidatorSetByDynasty(dynasty *big.Int) (*score.ValidatorSet, error)
}
