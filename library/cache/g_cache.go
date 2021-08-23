package cache

import (
	"github.com/gogf/gf/os/gcache"
)

type GCache struct {
	cache *gcache.Cache
	//duration time.Duration
}

func NewGCache() *GCache {
	return &GCache{
		cache: gcache.New(),
	}
}

func (c *GCache) Get(key interface{}) (interface{}, error) {
	return c.cache.Get(key)
}

func (c *GCache) Set(key, value interface{}, options *Options) error {
	return c.cache.Set(key, value, options.Duration)
}

func (c *GCache) Delete(key interface{}) error {
	_, err := c.cache.Remove(key)
	return err
}

func (c *GCache) Clear() error {
	c.cache.Close()
	return nil
}
