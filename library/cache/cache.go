package cache

import (
	"context"
	"time"
)

var CacheManager *Cache

type CacheInterface interface {
	Get(key interface{}) (interface{}, error)
	Set(key, object interface{}, options *Options) error
	Delete(key interface{}) error
	Clear() error
}

type Options struct {
	Cost     int64
	Duration time.Duration
	Tags     []string
}

type Cache struct {
	ctx   context.Context
	cache CacheInterface
}

func NewCache(c CacheInterface) *Cache {
	ctx := context.Background()
	return &Cache{
		cache: c,
		ctx:   ctx,
	}
}

func (c *Cache) Get(key interface{}) (interface{}, error) {
	return c.cache.Get(key)
}

func (c *Cache) Set(key, value interface{}, Options *Options) error {
	return c.cache.Set(key, value, Options)
}

func (c *Cache) Delete(key interface{}) error {
	return c.cache.Delete(key)
}

func (c *Cache) Clear() error {
	return c.cache.Clear()
}
