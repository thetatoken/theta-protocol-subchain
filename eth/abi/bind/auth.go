// Copyright 2016 The go-ethereum Authors
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

package bind

import (
	"context"
	"errors"
	"math/big"

	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/theta/crypto"
	"github.com/thetatoken/thetasubchain/eth/core/types"
)

// ErrNoChainID is returned whenever the user failed to specify a chain id.
var ErrNoChainID = errors.New("no chain id specified")

// ErrNotAuthorized is returned when an account is not properly unlocked.
var ErrNotAuthorized = errors.New("not authorized to sign this account")

// NewKeyedTransactorWithChainID is a utility method to easily create a transaction signer
// from a single private key.
func NewKeyedTransactorWithChainID(key *crypto.PrivateKey, chainID *big.Int) (*TransactOpts, error) {
	keyAddr := key.PublicKey().Address()
	if chainID == nil {
		return nil, ErrNoChainID
	}
	signer := types.LatestSignerForChainID(chainID)
	return &TransactOpts{
		From: keyAddr,
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			if address != keyAddr {
				return nil, ErrNotAuthorized
			}
			signature, err := crypto.Sign(signer.Hash(tx).Bytes(), crypto.PrivKeyToECDSA(key))
			if err != nil {
				return nil, err
			}
			return tx.WithSignature(signer, signature)
		},
		Context: context.Background(),
	}, nil
}
