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

func packAddressParam(param common.Address) []byte {
	paramBytes := common.LeftPadBytes(param.Bytes(), int(wordSizeInBytes))
	return paramBytes
}

func packUintParam(param uint) []byte {
	paramBytes, err := hex.DecodeString(fmt.Sprintf("%064x", param))
	if err != nil {
		panic(fmt.Sprintf("Failed to encode uint param %v: %v", fmt.Sprintf("%x", param), err)) // should never happen
	}
	paramBytes = common.LeftPadBytes(paramBytes, int(wordSizeInBytes))
	return paramBytes
}

func packBigIntParam(param *big.Int) []byte {
	paramBytes, err := hex.DecodeString(fmt.Sprintf("%064x", param))
	if err != nil {
		panic(fmt.Sprintf("Failed to encode *big.Int param %v: %v", fmt.Sprintf("%x", param), err)) // should never happen
	}
	paramBytes = common.LeftPadBytes(paramBytes, int(wordSizeInBytes))
	return paramBytes
}

func packStringParam(param string) ([]byte, uint) {
	hexParam := hex.EncodeToString([]byte(param))
	hexParamBytes, err := hex.DecodeString(hexParam)
	if err != nil {
		panic(fmt.Sprintf("Failed to encode string param %v: %v", param, err)) // should never happen
	}

	lenHexStr := uint(len(hexParam) / 2)

	// first pack the length of the string
	paramBytes := packUintParam(lenHexStr)

	// next, pack the string itself
	for i := uint(0); i*wordSizeInBytes < lenHexStr; i++ {
		start := i * wordSizeInBytes
		end := (i + 1) * wordSizeInBytes
		if end > lenHexStr {
			end = lenHexStr
		}
		paramBytes = append(paramBytes, common.RightPadBytes(hexParamBytes[start:end], int(wordSizeInBytes))...)
	}

	worldSize := uint(len(paramBytes)) / wordSizeInBytes
	if uint(len(paramBytes))%wordSizeInBytes != 0 {
		worldSize++
	}

	return paramBytes, uint(worldSize * wordSizeInBytes)
}

func PrepareTFuelCalldata(subchainID *big.Int, dynasty *big.Int, voucherOwner common.Address, mainchainTokenReceiver common.Address, amount *big.Int, nonce *big.Int, denom string) []byte {
	calldata := append([]byte{})
	denomData, _ := packStringParam(denom)
	denomParamOffset := 7 * wordSizeInBytes // 7: the offsets of seven parameters

	// offset
	calldata = append(calldata, packBigIntParam(subchainID)...)
	calldata = append(calldata, packBigIntParam(dynasty)...)
	calldata = append(calldata, packAddressParam(voucherOwner)...)
	calldata = append(calldata, packAddressParam(mainchainTokenReceiver)...)
	calldata = append(calldata, packBigIntParam(amount)...)
	calldata = append(calldata, packBigIntParam(nonce)...)
	calldata = append(calldata, packUintParam(denomParamOffset)...)
	calldata = append(calldata, denomData...)

	res := "0x" + hex.EncodeToString(calldata)
	logger.Infof("calldata bytes to string : %v", res)
	return calldata
}
