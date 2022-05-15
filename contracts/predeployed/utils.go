package predeployed

import (
	"encoding/hex"
	"fmt"
	"math/big"

	log "github.com/sirupsen/logrus"

	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/theta/ledger/types"
)

const wordSizeInBytes uint = 32

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

func packUintParam(param uint) []byte {
	paramBytes := common.LeftPadBytes([]byte(fmt.Sprintf("%032x", param)), int(wordSizeInBytes))
	return paramBytes
}

func packBigIntParam(param *big.Int) []byte {
	paramBytes, err := hex.DecodeString(fmt.Sprintf("%032x", param))
	if err != nil {
		panic(fmt.Sprintf("Failed to encode *big.Int param %v: %v", fmt.Sprintf("%x", param), err)) // should never happen
	}

	return paramBytes
}

func packStringParam(param string) ([]byte, uint) {
	hexParam := hex.EncodeToString([]byte(param))
	lenHexStr := uint(len(hexParam) / 2)

	// first pack the length of the string
	paramBytes := packUintParam(lenHexStr)

	// next, pack the string itself
	for i := uint(0); i*wordSizeInBytes < lenHexStr; i++ {
		start := i * wordSizeInBytes
		end := (i + 1) * wordSizeInBytes
		paramBytes = append(paramBytes, common.RightPadBytes([]byte(hexParam[start:end]), int(wordSizeInBytes))...)
	}

	worldSize := uint(len(paramBytes)) / wordSizeInBytes
	if uint(len(paramBytes))%wordSizeInBytes != 0 {
		worldSize++
	}

	return paramBytes, uint(worldSize)
}
