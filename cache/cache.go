package cache

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/coreos/bbolt"
)

type callback func(key string) (interface{}, error)

type Cache struct {
	Time    time.Time
	Content interface{}
}

type Cachedb struct {
	db       *bbolt.DB
	duration time.Duration
}

const (
	DBNAME = "DB"
)

var (
	BUCKET = []byte(DBNAME)
)

func New(d time.Duration, path string) (*Cachedb, error) {
	db, err := bbolt.Open(path, 0600, nil)
	if err != nil {
		return nil, fmt.Errorf("can't open db, %v", err)
	}
	err = db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(BUCKET)
		if err != nil {
			return fmt.Errorf("can't create DB bucket: %v", err)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("set up bucket error, %v", err)
	}
	return &Cachedb{
		db:       db,
		duration: d,
	}, nil
}

func (c *Cachedb) Close() {
	c.db.Close()
}

func (c *Cachedb) set(key string, value interface{}) error {
	cache := Cache{
		Time:    time.Now().Add(c.duration),
		Content: value,
	}
	encoded, err := json.Marshal(cache)
	if err != nil {
		return fmt.Errorf("can't encode in json: %v", err)
	}
	err = c.db.Update(func(tx *bbolt.Tx) error {
		err = tx.Bucket([]byte("DB")).Put([]byte(key), encoded)
		if err != nil {
			return fmt.Errorf("can't set cache: %v", err)
		}
		return nil
	})
	return err
}

// Get key return value, is expired, error
func (c *Cachedb) get(key string) (interface{}, bool, error) {
	now := time.Now()
	var encoded []byte
	err := c.db.View(func(tx *bbolt.Tx) error {
		bk := tx.Bucket(BUCKET)
		if bk == nil {
			return fmt.Errorf("failed to get bucket DB")
		}
		encoded = bk.Get([]byte(key))
		return nil
	})
	if err != nil {
		return nil, false, err
	}
	if encoded == nil {
		return nil, false, nil
	}
	var cache Cache
	err = json.Unmarshal(encoded, &cache)
	if err != nil {
		return nil, false, err
	}
	return cache.Content, cache.Time.Before(now), nil
}

// LazyGet an item from the cache. If the item is not found or expired, get it from the callback function and set it in the cache, or take the expired cache if fail
func (c *Cachedb) LazyGet(key string, fn callback) (interface{}, error) {
	cached, exp, err := c.get(key)
	if err != nil {
		return nil, err
	}
	if cached == nil || (cached != nil && exp) {
		result, err := fn(key)
		if err != nil && cached == nil {
			return nil, err
		} else if err == nil {
			err = c.set(key, result)
			if err != nil {
				return result, err
			}
			return result, nil
		}
	}
	return cached, nil
}
