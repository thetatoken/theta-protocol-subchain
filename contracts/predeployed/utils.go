package predeployed

import (
	"math/big"

	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/theta/crypto"
	"github.com/thetatoken/theta/ledger/types"
)

var mintVouchersFuncSelector = crypto.Keccak256([]byte("mintVouchers(string,address,bytes)"))[:4]

func constructProxySmartContractTx(proposerAddr common.Address, smartContractAddress common.Address, callData []byte) (*types.SmartContractTx, error) {
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
		Data:     callData,
	}

	return sctx, nil
}
