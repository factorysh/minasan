package cache

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/coreos/bbolt"
)

const (
	expirationTime = 5 * time.Minute
)

type callback func(key string) (interface{}, error)

type Cache struct {
	Time    time.Time
	Content interface{}
}

func setupDB() (*bbolt.DB, error) {
	db, err := bbolt.Open("cache.db", 0600, nil)
	if err != nil {
		return nil, fmt.Errorf("can't open db, %v", err)
	}
	err = db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("DB"))
		if err != nil {
			return fmt.Errorf("can't create DB bucket: %v", err)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("set up bucket error, %v", err)
	}
	return db, nil
}

func setCache(db *bbolt.DB, value interface{}) error {
	err := db.Update(func(tx *bbolt.Tx) error {
		cache := Cache{time.Now(), value}
		encoded, err := json.Marshal(cache)
		if err != nil {
			return fmt.Errorf("can't encode in json: %v", err)
		}
		err = tx.Bucket([]byte("DB")).Put([]byte("CACHE"), encoded)
		if err != nil {
			return fmt.Errorf("can't set cache: %v", err)
		}
		return nil
	})
	return err
}

func getCache(db *bbolt.DB, expTime time.Duration) (*Cache, bool, error) {
	cache := new(Cache)
	exp := false
	err := db.View(func(tx *bbolt.Tx) error {
		bk := tx.Bucket([]byte("DB"))
		if bk == nil {
			return fmt.Errorf("failed to get bucket DB")
		}
		encoded := bk.Get([]byte("CACHE"))
		if encoded == nil {
			return fmt.Errorf("CACHE not found in the db")
		}
		err := json.Unmarshal(encoded, &cache)
		if err != nil {
			return fmt.Errorf("can't decode json: %v", err)
		}
		t := time.Since(cache.Time)
		if t > expTime {
			exp = true
		}
		return nil
	})
	if err != nil {
		return nil, false, err
	}
	return cache, exp, err
}

// Get an item from the cache. If the item is not found or expired, get it from the callback function and set it in the cache, or take the expired cache if fail
func GetWithCallback(key string, fn callback) (interface{}, error) {
	db, err := setupDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	cached, exp, err := getCache(db, expirationTime)
	if cached == nil || (cached != nil && exp) {
		result, err := fn(key)
		if err != nil && cached == nil {
			return nil, err
		} else if err == nil {
			err = setCache(db, result)
			if err != nil {
				return result, err
			}
			return result, nil
		}
	}
	return cached.Content, nil
}
