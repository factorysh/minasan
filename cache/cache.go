package cache

import (
	"time"

	_cache "github.com/patrickmn/go-cache"
)

type Cache struct {
	*_cache.Cache
}

type callback func(key string) (interface{}, error)

func NewCache() *Cache {
	cache := _cache.New(5*time.Minute, 0)
	return &Cache{cache}
}

// Get an item form the cache. If the item is not found or expired, get it from the callback function and set it in the cache, or take the expired cache if fail
func (c *Cache) GetWithCallback(key string, fn callback) (interface{}, error) {
	cached, expTime, found := c.GetWithExpiration(key)
	if !found || (found && _cache.Item{cached, expTime.Unix()}.Expired()) {
		result, err := fn(key)
		if err != nil && !found {
			return nil, err
		} else if err == nil {
			c.SetDefault(key, result)
			return result, nil
		}
	}
	return cached, nil
}
