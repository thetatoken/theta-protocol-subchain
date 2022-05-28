package common

import (
	"math/big"

	log "github.com/sirupsen/logrus"
	"github.com/thetatoken/theta/ledger/types"
)

const NumMainchainBlocksPerDynasty int64 = 100

var logger *log.Entry = log.WithFields(log.Fields{"prefix": "common"})

func CalculateDynasty(mainchainHeight *big.Int) *big.Int {
	return new(big.Int).Div(mainchainHeight, big.NewInt(NumMainchainBlocksPerDynasty))
}

func MapChainID(chainIDStr string) *big.Int {
	mainChainBlockHeight := uint64(1000000000) // doesn't really matter for subchains, just set it to a sufficiently large number
	return types.MapChainID(chainIDStr, mainChainBlockHeight)
}
