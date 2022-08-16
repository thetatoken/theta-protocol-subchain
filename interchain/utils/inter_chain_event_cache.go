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
	err := store.Put(InterChainEventIndexKey(event.SourceChainID, event.Type, event.Nonce), event)
	return err // the caller should handle the error
}

func (c *InterChainEventCache) InsertList(events []*score.InterChainMessageEvent) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	store := kvstore.NewKVStore(c.db)
	for _, event := range events {
		err := store.Put(InterChainEventIndexKey(event.SourceChainID, event.Type, event.Nonce), event)
		if err != nil {
			return err // the caller should handle the error
		}
	}
	return nil
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
