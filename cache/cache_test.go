package cache

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/coreos/bbolt"
	"github.com/stretchr/testify/assert"
)

func cleanDB(db *bbolt.DB) error {
	err := db.Update(func(tx *bbolt.Tx) error {
		err := tx.Bucket([]byte("DB")).Delete([]byte("CACHE"))
		if err != nil {
			return fmt.Errorf("can't delete cache bucket: %v", err)
		}
		err = tx.DeleteBucket([]byte("DB"))
		if err != nil {
			return fmt.Errorf("can't delete db: %v", err)
		}
		return nil
	})
	return err
}

// Good callback function
func CallbackGood(key string) (interface{}, error) {
	tab := []string{"test", "blabla", "coucou"}
	return tab, nil
}

// Bad callback function
func CallbackBad(key string) (interface{}, error) {
	return nil, fmt.Errorf("error")
}

// Test for the SetupDB
func TestSetupDB(t *testing.T) {
	db, err := setupDB()
	assert.NoError(t, err)
	defer db.Close()
	assert.FileExists(t, "cache.db", "file cache.db should exist")
	err = cleanDB(db)
	assert.NoError(t, err)
}

func TestSetCache(t *testing.T) {
	// Setup the DB
	db, err := setupDB()
	assert.NoError(t, err)
	defer db.Close()
	// Set the test cache
	err = setCache(db, "test_set_string")
	assert.NoError(t, err)
	// get the cache manualy
	cache := new(Cache)
	err = db.View(func(tx *bbolt.Tx) error {
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
		return nil
	})
	assert.NoError(t, err)
	// Check if the cache is good
	assert.Equal(t, "test_set_string", cache.Content, "The two interfaces should be the same")
	// Clean the DB
	err = cleanDB(db)
	assert.NoError(t, err)
}

func TestGetCache(t *testing.T) {
	// Setup the DB
	db, err := setupDB()
	assert.NoError(t, err)
	defer db.Close()
	// Create an artificial cache in the DB
	err = db.Update(func(tx *bbolt.Tx) error {
		cache := Cache{time.Now(), "test_get_string"}
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
	assert.NoError(t, err)
	// get the cache with normal expiration time -> should be good
	cached, exp, err := getCache(db, expirationTime)
	assert.NoError(t, err)
	assert.Equal(t, "test_get_string", cached.Content, "The two interfaces should be the same")
	assert.False(t, exp, "The expiration time should be false")
	// get the cache with bad expiration time -> should be expired
	_, exp, err = getCache(db, 1)
	assert.NoError(t, err)
	assert.True(t, exp, "The expiration time should be true")
	// Clean the DB
	err = cleanDB(db)
	assert.NoError(t, err)
}

func TestGeneralCache(t *testing.T) {
	// Callback function return an error
	_, err := GetWithCallback("test_bad", CallbackBad)
	assert.Error(t, err)
	// get the good tab from the callback function for compare
	goodtab, _ := CallbackGood("test")
	// Callback function return good
	result, err := GetWithCallback("test_good", CallbackGood)
	assert.NoError(t, err)
	assert.ElementsMatch(t, goodtab, result, "The two interfaces should be the same")
	// Callback function return an error, but shoul return the previous result from cache
	result2, err2 := GetWithCallback("test_bad_but_good", CallbackBad)
	assert.NoError(t, err2)
	assert.ElementsMatch(t, goodtab, result2, "The two interfaces should be the same")
	// Remove the DB file
	os.Remove("cache.db")
}
