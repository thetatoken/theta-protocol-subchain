package core

import (
	"errors"
	"math/big"
	"strconv"
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/thetatoken/theta/common"
	ts "github.com/thetatoken/theta/store"
	"github.com/thetatoken/theta/store/database"
	"github.com/thetatoken/theta/store/kvstore"
	score "github.com/thetatoken/thetasubchain/core"
)

var logger *log.Entry = log.WithFields(log.Fields{"prefix": "interchain"})

// ------------------------------------ Inter-Chain Event Cache ----------------------------------------------

var (
	// ErrInterChainMessageEventrNotFound for ID is not found in crosschain transfer event set.
	ErrInterChainMessageEventNotFound      = errors.New("InterChainMessageEventNotFound")
	ErrInterChainMessageEventExisted       = errors.New("InterChainMessageEventrExisted")
	ErrInterChainMessageEventPersistFailed = errors.New("InterChainMessageEventPersistFailed")
)

// InterChainEventIndexKey constructs the DB key for the given block hash.
func InterChainEventIndexKey(sourceChainID *big.Int, icmeType score.InterChainMessageEventType, nonce *big.Int) common.Bytes {
	return common.Bytes("ice/" + sourceChainID.String() + "/" + strconv.FormatUint(uint64(icmeType), 10) + "/" + nonce.String())
}

type InterChainEventCache struct {
	mutex *sync.Mutex // mutex to for concurrency protection, e.g., the witness thread and consensus thread may access it concurrently
	db    database.Database
}

// NewInterChainEventCache creates a new crosschain transfer event cache instance.
func NewInterChainEventCache(db database.Database) *InterChainEventCache {
	cache := &InterChainEventCache{
		mutex: &sync.Mutex{},
		db:    db,
	}
	return cache
}

func (c *InterChainEventCache) Insert(event *score.InterChainMessageEvent) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	store := kvstore.NewKVStore(c.db)
	key := InterChainEventIndexKey(event.SourceChainID, event.Type, event.Nonce)
	reportNonceCollision(store, key, event)
	err := store.Put(key, event)
	return err // the caller should handle the error
}

func (c *InterChainEventCache) InsertList(events []*score.InterChainMessageEvent) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	store := kvstore.NewKVStore(c.db)
	for _, event := range events {
		key := InterChainEventIndexKey(event.SourceChainID, event.Type, event.Nonce)
		reportNonceCollision(store, key, event)
		err := store.Put(key, event)
		if err != nil {
			return err // the caller should handle the error
		}
	}
	return nil
}

// reportNonceCollision raises an alarm when two materially different events claim
// the same (source chain, event type, nonce) slot.
//
// The TokenBank assigns each nonce exactly once, so this is not expected to
// happen. If it does, at most one of the two events reflects committed state. The
// cache is keyed by nonce, so the second event silently displaces the first; make
// that visible rather than quiet.
//
// This is a detection aid only. The authoritative check happens in the witness
// and the orchestrator, which confirm every event against the emitting
// TokenBank's own state before relaying it.
func reportNonceCollision(store ts.Store, key common.Bytes, event *score.InterChainMessageEvent) {
	existing := score.InterChainMessageEvent{}
	if err := store.Get(key, &existing); err != nil {
		return // nothing cached under this key yet, which is the normal case
	}
	if existing.Equals(event) {
		return // the same event seen twice, e.g. after a re-scan
	}

	logger.Errorf("Two different inter-chain events claim source chain %v, type %v, nonce %v. "+
		"At most one of them can reflect committed state. Cached: %v. Incoming: %v",
		event.SourceChainID, event.Type, event.Nonce, existing.String(), event.String())
}

func (c *InterChainEventCache) Delete(sourceChainID *big.Int, imceType score.InterChainMessageEventType, nonce *big.Int) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	store := kvstore.NewKVStore(c.db)
	err := store.Delete(InterChainEventIndexKey(sourceChainID, imceType, nonce))
	return err // the caller should handle the error
}

func (c *InterChainEventCache) Get(sourceChainID *big.Int, imceType score.InterChainMessageEventType, nonce *big.Int) (*score.InterChainMessageEvent, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	event := score.InterChainMessageEvent{}
	store := kvstore.NewKVStore(c.db)
	err := store.Get(InterChainEventIndexKey(sourceChainID, imceType, nonce), &event)
	return &event, err // the caller should handle the error
}

func (c *InterChainEventCache) Exists(sourceChainID *big.Int, imceType score.InterChainMessageEventType, nonce *big.Int) (bool, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	event := score.InterChainMessageEvent{}
	store := kvstore.NewKVStore(c.db)
	err := store.Get(InterChainEventIndexKey(sourceChainID, imceType, nonce), &event)
	if err == nil {
		return true, nil
	}

	if err == ts.ErrKeyNotFound {
		return false, nil
	}

	return false, err // the caller should handle the error
}
