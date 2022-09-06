// Copyright 2014 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package vm

import (
	"crypto/sha256"
	"errors"
	"math/big"
	"strings"

	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/theta/common/math"
	"github.com/thetatoken/theta/crypto"
	"github.com/thetatoken/theta/crypto/bn256"
	"github.com/thetatoken/theta/ledger/vm/params"
	"github.com/thetatoken/thetasubchain/eth/abi"
	"golang.org/x/crypto/ripemd160"
)

// PrecompiledContract is the basic interface for native Go contracts. The implementation
// requires a deterministic gas count based on the input size of the Run method of the
// contract.
type PrecompiledContract interface {
	RequiredGas(input []byte, blockHeight uint64) uint64                   // RequiredPrice calculates the contract gas use
	Run(evm *EVM, input []byte, callerAddr common.Address) ([]byte, error) // Run runs the precompiled contract
}

var PrecompiledContracts = map[common.Address]PrecompiledContract{
	common.BytesToAddress([]byte{1}): &ecrecover{},
	common.BytesToAddress([]byte{2}): &sha256hash{},
	common.BytesToAddress([]byte{3}): &ripemd160hash{},
	common.BytesToAddress([]byte{4}): &dataCopy{},
	common.BytesToAddress([]byte{5}): &bigModExp{},
	common.BytesToAddress([]byte{6}): &bn256Add{},
	common.BytesToAddress([]byte{7}): &bn256ScalarMul{},
	common.BytesToAddress([]byte{8}): &bn256Pairing{},

	common.BytesToAddress([]byte{180}): &dynasty{},
	common.BytesToAddress([]byte{181}): &validatorSet{},
	common.BytesToAddress([]byte{182}): &mintTFuel{},
	common.BytesToAddress([]byte{183}): &burnTFuel{},

	// common.BytesToAddress([]byte{201}): &thetaBalance{},
	// common.BytesToAddress([]byte{202}): &thetaStake{},
	// common.BytesToAddress([]byte{203}): &transferTheta{},
}

// RunPrecompiledContract runs and evaluates the output of a precompiled contract.
func RunPrecompiledContract(evm *EVM, p PrecompiledContract, input []byte, contract *Contract) (ret []byte, err error) {
	blockHeight := evm.StateDB.GetBlockHeight()
	gas := p.RequiredGas(input, blockHeight)
	if contract.UseGas(gas) {
		callerAddr := contract.CallerAddress
		return p.Run(evm, input, callerAddr)
	}
	return nil, ErrOutOfGas
}

// ECRECOVER implemented as a native contract.
type ecrecover struct{}

func (c *ecrecover) RequiredGas(input []byte, blockHeight uint64) uint64 {
	return params.EcrecoverGas
}

func (c *ecrecover) Run(evm *EVM, input []byte, callerAddr common.Address) ([]byte, error) {
	const ecRecoverInputLength = 128

	input = common.RightPadBytes(input, ecRecoverInputLength)
	// "input" is (hash, v, r, s), each 32 bytes
	// but for ecrecover we want (r, s, v)

	r := new(big.Int).SetBytes(input[64:96])
	s := new(big.Int).SetBytes(input[96:128])
	v := input[63] - 27

	// tighter sig s values input homestead only apply to tx sigs
	if !allZero(input[32:63]) || !crypto.ValidateSignatureValues(v, r, s, false) {
		return nil, nil
	}
	// v needs to be at the end for libsecp256k1
	pubKey, err := crypto.Ecrecover(input[:32], append(input[64:128], v))
	// make sure the public key is a valid one
	if err != nil {
		return nil, nil
	}

	// the first byte of pubkey is bitcoin heritage
	return common.LeftPadBytes(crypto.Keccak256(pubKey[1:])[12:], 32), nil
}

// SHA256 implemented as a native contract.
type sha256hash struct{}

// RequiredGas returns the gas required to execute the pre-compiled contract.
//
// This method does not require any overflow checking as the input size gas costs
// required for anything significant is so high it's impossible to pay for.
func (c *sha256hash) RequiredGas(input []byte, blockHeight uint64) uint64 {
	return uint64(len(input)+31)/32*params.Sha256PerWordGas + params.Sha256BaseGas
}
func (c *sha256hash) Run(evm *EVM, input []byte, callerAddr common.Address) ([]byte, error) {
	h := sha256.Sum256(input)
	return h[:], nil
}

// RIPEMD160 implemented as a native contract.
type ripemd160hash struct{}

// RequiredGas returns the gas required to execute the pre-compiled contract.
//
// This method does not require any overflow checking as the input size gas costs
// required for anything significant is so high it's impossible to pay for.
func (c *ripemd160hash) RequiredGas(input []byte, blockHeight uint64) uint64 {
	return uint64(len(input)+31)/32*params.Ripemd160PerWordGas + params.Ripemd160BaseGas
}
func (c *ripemd160hash) Run(evm *EVM, input []byte, callerAddr common.Address) ([]byte, error) {
	ripemd := ripemd160.New()
	ripemd.Write(input)
	return common.LeftPadBytes(ripemd.Sum(nil), 32), nil
}

// data copy implemented as a native contract.
type dataCopy struct{}

// RequiredGas returns the gas required to execute the pre-compiled contract.
//
// This method does not require any overflow checking as the input size gas costs
// required for anything significant is so high it's impossible to pay for.
func (c *dataCopy) RequiredGas(input []byte, blockHeight uint64) uint64 {
	return uint64(len(input)+31)/32*params.IdentityPerWordGas + params.IdentityBaseGas
}
func (c *dataCopy) Run(evm *EVM, in []byte, callerAddr common.Address) ([]byte, error) {
	return in, nil
}

// bigModExp implements a native big integer exponential modular operation.
type bigModExp struct{}

var (
	big1      = big.NewInt(1)
	big4      = big.NewInt(4)
	big8      = big.NewInt(8)
	big16     = big.NewInt(16)
	big32     = big.NewInt(32)
	big64     = big.NewInt(64)
	big96     = big.NewInt(96)
	big480    = big.NewInt(480)
	big1024   = big.NewInt(1024)
	big3072   = big.NewInt(3072)
	big199680 = big.NewInt(199680)
)

// RequiredGas returns the gas required to execute the pre-compiled contract.
func (c *bigModExp) RequiredGas(input []byte, blockHeight uint64) uint64 {
	var (
		baseLen = new(big.Int).SetBytes(getData(input, 0, 32))
		expLen  = new(big.Int).SetBytes(getData(input, 32, 32))
		modLen  = new(big.Int).SetBytes(getData(input, 64, 32))
	)
	if len(input) > 96 {
		input = input[96:]
	} else {
		input = input[:0]
	}
	// Retrieve the head 32 bytes of exp for the adjusted exponent length
	var expHead *big.Int
	if big.NewInt(int64(len(input))).Cmp(baseLen) <= 0 {
		expHead = new(big.Int)
	} else {
		if expLen.Cmp(big32) > 0 {
			expHead = new(big.Int).SetBytes(getData(input, baseLen.Uint64(), 32))
		} else {
			expHead = new(big.Int).SetBytes(getData(input, baseLen.Uint64(), expLen.Uint64()))
		}
	}
	// Calculate the adjusted exponent length
	var msb int
	if bitlen := expHead.BitLen(); bitlen > 0 {
		msb = bitlen - 1
	}
	adjExpLen := new(big.Int)
	if expLen.Cmp(big32) > 0 {
		adjExpLen.Sub(expLen, big32)
		adjExpLen.Mul(big8, adjExpLen)
	}
	adjExpLen.Add(adjExpLen, big.NewInt(int64(msb)))

	// Calculate the gas cost of the operation
	gas := new(big.Int).Set(math.BigMax(modLen, baseLen))
	switch {
	case gas.Cmp(big64) <= 0:
		gas.Mul(gas, gas)
	case gas.Cmp(big1024) <= 0:
		gas = new(big.Int).Add(
			new(big.Int).Div(new(big.Int).Mul(gas, gas), big4),
			new(big.Int).Sub(new(big.Int).Mul(big96, gas), big3072),
		)
	default:
		gas = new(big.Int).Add(
			new(big.Int).Div(new(big.Int).Mul(gas, gas), big16),
			new(big.Int).Sub(new(big.Int).Mul(big480, gas), big199680),
		)
	}
	gas.Mul(gas, math.BigMax(adjExpLen, big1))
	gas.Div(gas, new(big.Int).SetUint64(params.ModExpQuadCoeffDiv))

	if gas.BitLen() > 64 {
		return math.MaxUint64Val
	}
	return gas.Uint64()
}

func (c *bigModExp) Run(evm *EVM, input []byte, callerAddr common.Address) ([]byte, error) {
	var (
		baseLen = new(big.Int).SetBytes(getData(input, 0, 32)).Uint64()
		expLen  = new(big.Int).SetBytes(getData(input, 32, 32)).Uint64()
		modLen  = new(big.Int).SetBytes(getData(input, 64, 32)).Uint64()
	)
	if len(input) > 96 {
		input = input[96:]
	} else {
		input = input[:0]
	}
	// Handle a special case when both the base and mod length is zero
	if baseLen == 0 && modLen == 0 {
		return []byte{}, nil
	}
	// Retrieve the operands and execute the exponentiation
	var (
		base = new(big.Int).SetBytes(getData(input, 0, baseLen))
		exp  = new(big.Int).SetBytes(getData(input, baseLen, expLen))
		mod  = new(big.Int).SetBytes(getData(input, baseLen+expLen, modLen))
	)
	if mod.BitLen() == 0 {
		// Modulo 0 is undefined, return zero
		return common.LeftPadBytes([]byte{}, int(modLen)), nil
	}
	return common.LeftPadBytes(base.Exp(base, exp, mod).Bytes(), int(modLen)), nil
}

// newCurvePoint unmarshals a binary blob into a bn256 elliptic curve point,
// returning it, or an error if the point is invalid.
func newCurvePoint(blob []byte) (*bn256.G1, error) {
	p := new(bn256.G1)
	if _, err := p.Unmarshal(blob); err != nil {
		return nil, err
	}
	return p, nil
}

// newTwistPoint unmarshals a binary blob into a bn256 elliptic curve point,
// returning it, or an error if the point is invalid.
func newTwistPoint(blob []byte) (*bn256.G2, error) {
	p := new(bn256.G2)
	if _, err := p.Unmarshal(blob); err != nil {
		return nil, err
	}
	return p, nil
}

// bn256Add implements a native elliptic curve point addition.
type bn256Add struct{}

// RequiredGas returns the gas required to execute the pre-compiled contract.
func (c *bn256Add) RequiredGas(input []byte, blockHeight uint64) uint64 {
	if blockHeight < common.HeightJune2021FeeAdjustment {
		return params.Bn256AddGas
	}
	return params.Bn256AddGasIstanbul
}

func (c *bn256Add) Run(evm *EVM, input []byte, callerAddr common.Address) ([]byte, error) {
	x, err := newCurvePoint(getData(input, 0, 64))
	if err != nil {
		return nil, err
	}
	y, err := newCurvePoint(getData(input, 64, 64))
	if err != nil {
		return nil, err
	}
	res := new(bn256.G1)
	res.Add(x, y)
	return res.Marshal(), nil
}

// bn256ScalarMul implements a native elliptic curve scalar multiplication.
type bn256ScalarMul struct{}

// RequiredGas returns the gas required to execute the pre-compiled contract.
func (c *bn256ScalarMul) RequiredGas(input []byte, blockHeight uint64) uint64 {
	if blockHeight < common.HeightJune2021FeeAdjustment {
		return params.Bn256ScalarMulGas
	}

	return params.Bn256ScalarMulGasIstanbul
}

func (c *bn256ScalarMul) Run(evm *EVM, input []byte, callerAddr common.Address) ([]byte, error) {
	p, err := newCurvePoint(getData(input, 0, 64))
	if err != nil {
		return nil, err
	}
	res := new(bn256.G1)
	res.ScalarMult(p, new(big.Int).SetBytes(getData(input, 64, 32)))
	return res.Marshal(), nil
}

var (
	// true32Byte is returned if the bn256 pairing check succeeds.
	true32Byte = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}

	// false32Byte is returned if the bn256 pairing check fails.
	false32Byte = make([]byte, 32)

	// errBadPairingInput is returned if the bn256 pairing input is invalid.
	errBadPairingInput = errors.New("bad elliptic curve pairing size")
)

// bn256Pairing implements a pairing pre-compile for the bn256 curve
type bn256Pairing struct{}

// RequiredGas returns the gas required to execute the pre-compiled contract.
func (c *bn256Pairing) RequiredGas(input []byte, blockHeight uint64) uint64 {
	if blockHeight < common.HeightJune2021FeeAdjustment {
		return params.Bn256PairingBaseGas + uint64(len(input)/192)*params.Bn256PairingPerPointGas
	}

	return params.Bn256PairingBaseGasIstanbul + uint64(len(input)/192)*params.Bn256PairingPerPointGasIstanbul
}

func (c *bn256Pairing) Run(evm *EVM, input []byte, callerAddr common.Address) ([]byte, error) {
	// Handle some corner cases cheaply
	if len(input)%192 > 0 {
		return nil, errBadPairingInput
	}
	// Convert the input into a set of coordinates
	var (
		cs []*bn256.G1
		ts []*bn256.G2
	)
	for i := 0; i < len(input); i += 192 {
		c, err := newCurvePoint(input[i : i+64])
		if err != nil {
			return nil, err
		}
		t, err := newTwistPoint(input[i+64 : i+192])
		if err != nil {
			return nil, err
		}
		cs = append(cs, c)
		ts = append(ts, t)
	}
	// Execute the pairing checks and return the results
	if bn256.PairingCheck(cs, ts) {
		return true32Byte, nil
	}
	return false32Byte, nil
}

// dynasty return the current dynasty
type dynasty struct {
}

func (c *dynasty) RequiredGas(input []byte, blockHeight uint64) uint64 {
	const dynastyGas = uint64(200)
	return dynastyGas
}

func (c *dynasty) Run(evm *EVM, input []byte, callerAddr common.Address) ([]byte, error) {
	dynasty := evm.StateDB.GetDynasty()

	dynastyBytes := dynasty.Bytes()
	dynastyBytes32 := common.LeftPadBytes(dynastyBytes[:], 32) // easier to convert bytes32 into uint256 in smart contracts
	return dynastyBytes32, nil
}

// validatorSet return the validator set of a chain for a given dynasty
type validatorSet struct {
}

func (c *validatorSet) RequiredGas(input []byte, blockHeight uint64) uint64 {
	const validatorSetGas = uint64(200)
	return validatorSetGas
}

func (c *validatorSet) Run(evm *EVM, input []byte, callerAddr common.Address) ([]byte, error) {
	const RawABI = `
[
	{
		"inputs": [
			{
				"components": [
					{
						"internalType": "address",
						"name": "validator",
						"type": "address"
					},
					{
						"internalType": "uint256",
						"name": "shareAmount",
						"type": "uint256"
					}
				],
				"internalType": "struct TestRegister.ValidatorAddrSharePair[]",
				"name": "a",
				"type": "tuple[]"
			}
		],
		"stateMutability": "nonpayable",
		"type": "constructor"
	}
]`
	type validatorSetInfo struct {
		Validator   common.Address
		ShareAmount *big.Int
	}
	parsed, err := abi.JSON(strings.NewReader(RawABI))
	if err != nil {
		panic(err)
	}
	var infoSet []validatorSetInfo

	chainID := new(big.Int).SetBytes(getData(input, 0, 32))
	dynasty := new(big.Int).SetBytes(getData(input, 32, 32))

	validatorSet := evm.StateDB.GetValidatorSetForChainDuringDynasty(chainID, dynasty)
	if validatorSet == nil {
		return common.Bytes{}, nil
	}

	validators := validatorSet.Validators()
	for _, val := range validators {
		c := &validatorSetInfo{
			Validator:   val.Address,
			ShareAmount: val.Stake,
		}
		infoSet = append(infoSet, *c)
	}
	valSetBytes, err := parsed.Pack("", infoSet)
	if err != nil {
		panic(err)
	}
	return valSetBytes, nil
}

// mintTFuel mints TFuel for a given account if the caller contract has previledged access
type mintTFuel struct {
}

func (c *mintTFuel) RequiredGas(input []byte, blockHeight uint64) uint64 {
	const mintTFuelGas = uint64(200)
	return mintTFuelGas
}

func (c *mintTFuel) Run(evm *EVM, input []byte, callerAddr common.Address) ([]byte, error) {
	if !evm.HasPreviledgedAccess(callerAddr) {
		return nil, errors.New("the execution context does not have previledged access")
	}

	recipient := common.BytesToAddress(getData(input, 0, 20))
	mintAmount := new(big.Int).SetBytes(getData(input, 20, 32))
	if mintAmount.Cmp(big.NewInt(0)) < 0 {
		return common.Bytes{}, errors.New("mint amount must not be negative")
	}

	db := evm.StateDB
	db.AddBalance(recipient, mintAmount)

	return common.Bytes{}, nil
}

// burnTFuel burns TFuel for a given smart contract account. Note that this precomplied contract does NOT
// need a current execution context with previledge access. A smart contract method that calls burnTFuel
// will deduct the specificed amount of TFuel from the contract's TFuel balance. The access control model
// is similar to TFuel transfer. It is up to the smart contract logic ensure that the caller has the right
// to burn the TFuel balance.
type burnTFuel struct {
}

func (c *burnTFuel) RequiredGas(input []byte, blockHeight uint64) uint64 {
	const burnTFuelGas = uint64(200)
	return burnTFuelGas
}

func (c *burnTFuel) Run(evm *EVM, input []byte, callerAddr common.Address) ([]byte, error) {
	burnAmount := new(big.Int).SetBytes(getData(input, 0, 32))
	if burnAmount.Cmp(big.NewInt(0)) < 0 {
		return common.Bytes{}, errors.New("burn amount must not be negative")
	}
	if !CanTransfer(evm.StateDB, callerAddr, burnAmount) { // callerAddr: the smart contract that called the burnTFuel precomplied contract
		return common.Bytes{}, ErrInsufficientBalance
	}
	db := evm.StateDB
	db.SubBalance(callerAddr, burnAmount)

	return common.Bytes{}, nil
}
