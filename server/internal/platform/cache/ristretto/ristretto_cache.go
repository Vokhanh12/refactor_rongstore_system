package ristretto

import (
	"fmt"

	"github.com/dgraph-io/ristretto/v2"
)

type RistrettoCache struct {
	cache *ristretto.Cache[string, any]
}

func NewRistrettoCache() (*RistrettoCache, error) {
	c, err := ristretto.NewCache(&ristretto.Config[string, any]{
		NumCounters: 1e4,
		MaxCost:     1 << 26, // ~64 MB
		BufferItems: 64,
	})

	if err != nil {
		return nil, err
	}

	return &RistrettoCache{
		cache: c,
	}, nil
}

func (c *RistrettoCache) Set(
	key string,
	value any,
	cost int64,
) error {

	ok := c.cache.Set(key, value, cost)

	if !ok {
		return fmt.Errorf("ristretto.set: cache rejected item")
	}

	return nil
}
func (c *RistrettoCache) Get(
	key string,
) (any, bool) {

	value, found := c.cache.Get(key)

	return value, found
}
