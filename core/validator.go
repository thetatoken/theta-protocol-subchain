package core

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"math/big"
	"sort"

	log "github.com/sirupsen/logrus"

	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/theta/rlp"
)

var logger *log.Entry = log.WithFields(log.Fields{"prefix": "core"})

var (
	// ErrValidatorNotFound for ID is not found in validator set.
	ErrValidatorNotFound = errors.New("ValidatorNotFound")
)

// Validator contains the public information of a validator.
type Validator struct {
	Address common.Address
	Stake   *big.Int
}

// NewValidator creates a new validator instance.
func NewValidator(addressStr string, stake *big.Int) Validator {
	address := common.HexToAddress(addressStr)
	return Validator{address, stake}
}

// ID returns the ID of the validator, which is the string representation of its address.
func (v Validator) ID() common.Address {
	return v.Address
}

// Equals checks whether the validator is the same as another validator
func (v Validator) Equals(x Validator) bool {
	logger.Debugf("v.Equals(x) check, v.Addr: %v, x.Addr: %v, v.Stake: %v, x.Stake: %v\n", v.Address.Hex(), x.Address.Hex(), v.Stake, x.Stake)
	if v.Address != x.Address {
		return false
	}
	if v.Stake.Cmp(x.Stake) != 0 {
		return false
	}
	return true
}

// String represents the string representation of the validator
func (v Validator) String() string {
	return fmt.Sprintf("{ID: %v, Stake: %v}", v.ID(), v.Stake)
}

// ValidatorSet represents a set of validators.
type ValidatorSet struct {
	dynasty    *big.Int
	validators []Validator
}

// NewValidatorSet returns a new instance of ValidatorSet.
func NewValidatorSet(dynasty *big.Int) *ValidatorSet {
	return &ValidatorSet{
		dynasty:    dynasty,
		validators: []Validator{},
	}
}

// Dynasty returns the dynasty of the validator set.
func (s *ValidatorSet) Dynasty() *big.Int {
	return s.dynasty
}

// // SetValidators sets validators
// func (s *ValidatorSet) SetValidators(validators []Validator) {
// 	s.validators = validators
// }

// Copy creates a copy of this validator set.
func (s *ValidatorSet) Copy() *ValidatorSet {
	ret := NewValidatorSet(s.dynasty)
	for _, v := range s.Validators() {
		ret.AddValidator(v)
	}
	return ret
}

// Size returns the number of the validators in the validator set.
func (s *ValidatorSet) Size() int {
	return len(s.validators)
}

// Equals checks whether the validator set is the same as another validator set
func (s *ValidatorSet) Equals(t *ValidatorSet) bool {
	if s.dynasty.Cmp(t.dynasty) != 0 {
		return false
	}

	numVals := len(s.validators)
	if numVals != len(t.validators) {
		return false
	}
	for i := 0; i < numVals; i++ {
		if !s.validators[i].Equals(t.validators[i]) {
			return false
		}
	}
	return true
}

// EqualsDisregardingDynasty checks whether the validator set is the same as another validator set disregarding the dynasty
func (s *ValidatorSet) EqualsDisregardingDynasty(t *ValidatorSet) bool {
	numVals := len(s.validators)
	if numVals != len(t.validators) {
		return false
	}
	for i := 0; i < numVals; i++ {
		if !s.validators[i].Equals(t.validators[i]) {
			return false
		}
	}
	return true
}

// String represents the string representation of the validator set
func (s *ValidatorSet) String() string {
	return fmt.Sprintf("{Dynasty: %v, Validators: %v}", s.dynasty, s.validators)
}

// ByID implements sort.Interface for ValidatorSet based on ID.
type ByID []Validator

func (b ByID) Len() int           { return len(b) }
func (b ByID) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b ByID) Less(i, j int) bool { return bytes.Compare(b[i].ID().Bytes(), b[j].ID().Bytes()) < 0 }

// GetValidator returns a validator if a matching ID is found.
func (s *ValidatorSet) GetValidator(id common.Address) (Validator, error) {
	for _, v := range s.validators {
		if v.ID() == id {
			return v, nil
		}
	}
	return Validator{}, ErrValidatorNotFound
}

// AddValidator adds a validator to the validator set.
func (s *ValidatorSet) AddValidator(validator Validator) {
	s.validators = append(s.validators, validator)
	sort.Sort(ByID(s.validators))
}

// TotalStake returns the total stake of the validators in the set.
func (s *ValidatorSet) TotalStake() *big.Int {
	ret := new(big.Int).SetUint64(0)
	for _, v := range s.validators {
		ret = new(big.Int).Add(ret, v.Stake)
	}
	return ret
}

// HasMajorityVotes checks whether a vote set has reach majority.
func (s *ValidatorSet) HasMajorityVotes(votes []Vote) bool {
	votedStake := new(big.Int).SetUint64(0)
	for _, vote := range votes {
		validator, err := s.GetValidator(vote.ID)
		if err == nil {
			votedStake = new(big.Int).Add(votedStake, validator.Stake)
		}
	}

	three := new(big.Int).SetUint64(3)
	two := new(big.Int).SetUint64(2)
	lhs := new(big.Int)
	rhs := new(big.Int)

	//return votedStake*3 > s.TotalStake()*2
	return lhs.Mul(votedStake, three).Cmp(rhs.Mul(s.TotalStake(), two)) > 0
}

// HasMajority checks whether a vote set has reach majority.
func (s *ValidatorSet) HasMajority(votes *VoteSet) bool {
	return s.HasMajorityVotes(votes.Votes())
}

// Validators returns a slice of validators.
func (s *ValidatorSet) Validators() []Validator {
	return s.validators
}

var _ rlp.Encoder = (*ValidatorSet)(nil)

// EncodeRLP implements RLP Encoder interface.
func (vs *ValidatorSet) EncodeRLP(w io.Writer) error {
	if vs == nil {
		return rlp.Encode(w, &ValidatorSet{})
	}
	return rlp.Encode(w, []interface{}{
		vs.dynasty,
		vs.validators,
	})
}

// DecodeRLP implements RLP Decoder interface.
func (vs *ValidatorSet) DecodeRLP(stream *rlp.Stream) error {
	_, err := stream.List()
	if err != nil {
		return err
	}

	dynasty := big.NewInt(0)
	err = stream.Decode(dynasty)
	if err != nil {
		return err
	}
	vs.dynasty = dynasty

	validators := []Validator{}
	err = stream.Decode(&validators)
	if err != nil {
		return err
	}
	vs.validators = validators

	return stream.ListEnd()
}
