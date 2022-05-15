package predeployed

import (
	"math/big"

	log "github.com/sirupsen/logrus"

	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/theta/ledger/types"
)

var logger *log.Entry = log.WithFields(log.Fields{"prefix": "predeployed"})

func constructProxySmartContractTx(proposerAddr common.Address, smartContractAddress common.Address, calldata []byte) (*types.SmartContractTx, error) {
	from := types.TxInput{
		Address: proposerAddr,
		Coins: types.Coins{
			ThetaWei: new(big.Int).SetUint64(0),
			TFuelWei: new(big.Int).SetUint64(0),
		},
		Sequence: 0,
	}

	to := types.TxOutput{
		Address: smartContractAddress,
	}

	dummyGasLimit := uint64(10000000)
	dummyGasPrice := big.NewInt(1)
	sctx := &types.SmartContractTx{
		From:     from,
		To:       to,
		GasLimit: dummyGasLimit,
		GasPrice: dummyGasPrice,
		Data:     calldata,
	}

	return sctx, nil
}
