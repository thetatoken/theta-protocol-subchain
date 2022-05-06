package common

import (
	"math/big"
)

const NumMainchainBlocksPerDynasty int64 = 100

func CalculateDynasty(mainchainHeight *big.Int) *big.Int {
	return new(big.Int).Div(mainchainHeight, big.NewInt(NumMainchainBlocksPerDynasty))
}
