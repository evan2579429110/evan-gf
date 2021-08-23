package cache

import (
	"context"
	"github.com/gogf/gf/database/gredis"
	"github.com/gogf/gf/frame/g"
)

type RedisCache struct {
	cache *gredis.Redis
	name  string
}

func NewRedisCache(name ...string) *RedisCache {
	return &RedisCache{
		cache: g.Redis(name...),
		name:  "cache-",
	}
}

func (c *RedisCache) Get(key interface{}) (interface{}, error) {
	conn := c.cache.Conn()
	defer conn.Close()
	return c.cache.DoVar("GET", c.name+key.(string))
}

func (c *RedisCache) Set(key, val interface{}, options *Options) error {
	conn := c.cache.Conn()
	defer conn.Close()
	_, err := c.cache.Do("SET", c.name+key.(string), val)
	return err
}

func (c *RedisCache) Delete(ctx context.Context, key interface{}) error {
	_, err := c.cache.Do("Del", c.name+key.(string))
	return err
}

func (c *RedisCache) Clear(ctx context.Context) error {
	panic("暂未开放该功能")
	//conn := c.cache.Conn()
	//defer conn.Close()
	//c.cache.DoVar("KEYS PATTERN",c.name)
	return nil

}
