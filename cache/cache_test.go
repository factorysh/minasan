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

// Good callback function
func CallbackGood(key string) (interface{}, error) {
	tab := []string{"test", "blabla", "coucou"}
	return tab, nil
}

// Bad callback function
func CallbackBad(key string) (interface{}, error) {
	return nil, fmt.Errorf("callback bad error")
}

func TestNewCache(t *testing.T) {
	c, err := New(5*time.Minute, "test_cache_minasan.db")
	assert.NoError(t, err)
	assert.NotNil(t, *c)
	assert.NotNil(t, *c.db)
	assert.Equal(t, 5*time.Minute, c.duration, "The times should be the same")
	err = c.set("testnew", "test_string")
	assert.NoError(t, err)
	c.Close()
	err = c.set("testnew", "test_string")
	assert.Error(t, err)
	// Remove the DB file
	os.Remove("test_cache_minasan.db")
}

// Test for the Get function
func TestGetCache(t *testing.T) {
	// Setup the DB
	c, err := New(5*time.Minute, "test_cache_minasan.db")
	assert.NoError(t, err)
	defer c.Close()
	// Create an artificial cache in the DB
	err = c.db.Update(func(tx *bbolt.Tx) error {
		cache := Cache{
			Time:    time.Now().Add(5 * time.Minute),
			Content: "test_get_string",
		}
		encoded, err := json.Marshal(cache)
		if err != nil {
			return fmt.Errorf("can't encode in json: %v", err)
		}
		err = tx.Bucket([]byte(BUCKET)).Put([]byte("testget"), encoded)
		if err != nil {
			return fmt.Errorf("can't set cache: %v", err)
		}
		return nil
	})
	assert.NoError(t, err)
	// get the cache with normal expiration time -> should be good
	cached, exp, err := c.get("testget")
	assert.NoError(t, err)
	assert.Equal(t, "test_get_string", cached, "The two interfaces should be the same")
	assert.False(t, exp, "The expiration time should be false")
	// Remove the DB file
	os.Remove("test_cache_minasan.db")
}

// Test for the Set function
func TestSetCache(t *testing.T) {
	// Setup the DB
	c, err := New(5*time.Minute, "test_cache_minasan.db")
	assert.NoError(t, err)
	defer c.Close()

	err = c.set("testset", "test_set_string")
	assert.NoError(t, err)
	// Get the cache with the normal (previously tested) Get
	cached, exp, err := c.get("testset")
	assert.NoError(t, err)
	assert.Equal(t, "test_set_string", cached, "The two interfaces should be the same")
	assert.False(t, exp, "The expiration time should be false")
	// Remove the DB file
	os.Remove("test_cache_minasan.db")
}

// Test the expiration time
func TestExpirationCache(t *testing.T) {
	// Setup the DB
	c, err := New(time.Nanosecond, "test_cache_minasan.db")
	assert.NoError(t, err)
	defer c.Close()

	err = c.set("testexp", "test_exp")
	assert.NoError(t, err)
	cached, exp, err := c.get("testexp")
	assert.NoError(t, err)
	assert.Equal(t, "test_exp", cached, "The two interfaces should be the same")
	assert.True(t, exp, "The expiration time should be true")
	// Remove the DB file
	os.Remove("test_cache_minasan.db")
}

// General test for the normal use of the cache
func TestGeneralCache(t *testing.T) {
	// New Cache
	c, err := New(5*time.Minute, "test_cache_minasan.db")
	assert.NoError(t, err)
	defer c.Close()
	// Callback function return an error
	_, err = c.LazyGet("test", CallbackBad)
	assert.Error(t, err)
	// get the good tab from the callback function for compare
	goodtab, _ := CallbackGood("test")
	// Callback function return good
	result, err := c.LazyGet("test", CallbackGood)
	assert.NoError(t, err)
	assert.ElementsMatch(t, goodtab, result, "The two interfaces should be the same")
	// Callback function return an error, but shoul return the previous result from cache
	result2, err2 := c.LazyGet("test", CallbackBad)
	assert.NoError(t, err2)
	assert.ElementsMatch(t, goodtab, result2, "The two interfaces should be the same")
	// Remove the DB file
	os.Remove("test_cache_minasan.db")
}
