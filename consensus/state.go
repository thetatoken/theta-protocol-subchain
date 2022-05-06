package consensus

import (
	"fmt"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/theta/store"
	sbc "github.com/thetatoken/thetasubchain/blockchain"
	score "github.com/thetatoken/thetasubchain/core"
)

type StateStub struct {
	Root                       common.Hash
	HighestCCBlock             common.Hash
	LastFinalizedBlock         common.Hash
	LastProposal               score.Proposal
	LastVote                   score.Vote
	LastInterChainMessageEvent score.InterChainMessageEvent
	Epoch                      uint64
}

const (
	DBStateStubKey      = "cs/ss"
	DBVoteByBlockPrefix = "cs/vbb/"
	DBEpochVotesKey     = "cs/ev"
)

type State struct {
	mu *sync.RWMutex

	db    store.Store
	chain *sbc.Chain

	highestCCBlock     common.Hash
	lastFinalizedBlock common.Hash

	LastProposal               score.Proposal
	LastVote                   score.Vote
	LastInterChainMessageEvent score.InterChainMessageEvent
	epoch                      uint64
}

func NewState(db store.Store, chain *sbc.Chain) *State {
	s := &State{
		mu:                 &sync.RWMutex{},
		db:                 db,
		chain:              chain,
		highestCCBlock:     chain.Root().Hash(),
		lastFinalizedBlock: chain.Root().Hash(),
		epoch:              chain.Root().Epoch,
	}
	err := s.Load()
	if err != nil {
		log.Panic(err)
	}
	return s
}

func (s *State) String() string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	highestCCBlockStr := s.highestCCBlock.Hex()
	lastFinalizedBlockStr := s.lastFinalizedBlock.Hex()

	return fmt.Sprintf("State{highestCCBlock: %v, lastFinalizedBlock: %v,  epoch: %d, LastProposal: %v, LastVote: %v}",
		highestCCBlockStr, lastFinalizedBlockStr, s.epoch, s.LastProposal, s.LastVote)
}

func (s *State) GetSummary() *StateStub {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.getSummary()
}

func (s *State) getSummary() *StateStub {
	stub := &StateStub{
		LastVote:     s.LastVote,
		LastProposal: s.LastProposal,
		Epoch:        s.epoch,
		Root:         s.chain.Root().Hash(),
	}
	stub.HighestCCBlock = s.highestCCBlock
	stub.LastFinalizedBlock = s.lastFinalizedBlock
	return stub
}

func (s *State) commit() error {
	stub := s.getSummary()
	key := []byte(DBStateStubKey)

	return s.db.Put(key, stub)
}

func (s *State) Load() (err error) {
	key := []byte(DBStateStubKey)
	stub := &StateStub{}
	s.db.Get(key, stub)

	if stub.Root != s.chain.Root().Hash() {
		logger.WithFields(log.Fields{
			"stub.Root":  stub.Root.Hex(),
			"chain.Root": s.chain.Root().Hash().Hex(),
		}).Info("Ignoring previous consensus state since it is on a different root")
		return
	}

	s.LastProposal = stub.LastProposal
	s.LastVote = stub.LastVote
	s.epoch = stub.Epoch
	s.lastFinalizedBlock = stub.LastFinalizedBlock
	s.highestCCBlock = stub.HighestCCBlock
	return
}

func (s *State) GetEpoch() uint64 {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.epoch
}

func (s *State) SetEpoch(epoch uint64) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.epoch = epoch
	return s.commit()
}

func (s *State) GetLastProposal() score.Proposal {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.LastProposal
}

func (s *State) SetLastProposal(p score.Proposal) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.LastProposal = p
	return s.commit()
}

func (s *State) GetLastVote() score.Vote {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.LastVote
}

func (s *State) SetLastVote(v score.Vote) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.LastVote = v
	return s.commit()
}

func (s *State) GetHighestCCBlock() *score.ExtendedBlock {
	s.mu.RLock()
	defer s.mu.RUnlock()

	ret, err := s.chain.FindBlock(s.highestCCBlock)
	if err != nil {
		log.Fatal("Failed to load highest CC block")
	}
	return ret
}

func (s *State) SetHighestCCBlock(block *score.ExtendedBlock) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.highestCCBlock = block.Hash()
	return s.commit()
}

func (s *State) GetLastFinalizedBlock() *score.ExtendedBlock {
	s.mu.RLock()
	defer s.mu.RUnlock()

	ret, err := s.chain.FindBlock(s.lastFinalizedBlock)
	if err != nil {
		log.Fatal("Failed to load last finalized block")
	}
	return ret
}

func (s *State) SetLastFinalizedBlock(block *score.ExtendedBlock) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.lastFinalizedBlock = block.Hash()
	return s.commit()
}

func (s *State) AddVote(vote *score.Vote) error {
	if err := s.AddEpochVote(vote); err != nil {
		return err
	}
	s.chain.AddVoteToIndex(*vote)
	return nil
}

func (s *State) GetEpochVotes() (*score.VoteSet, error) {
	key := []byte(DBEpochVotesKey)
	ret := score.NewVoteSet()
	err := s.db.Get(key, ret)
	return ret, err
}

func (s *State) AddEpochVote(vote *score.Vote) error {
	voteset, err := s.GetEpochVotes()
	if err != nil {
		voteset = score.NewVoteSet()
	}
	voteset.AddVote(*vote)
	voteset = voteset.UniqueVoter()

	key := []byte(DBEpochVotesKey)
	return s.db.Put(key, voteset)
}
