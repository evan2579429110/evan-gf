package cache

import (
	"fmt"
	"testing"
)

func TestGCacheGet(t *testing.T) {
	c := NewCache(NewGCache())
	err := c.Set("sss", "fff", &Options{})
	c.Clear()
	d, er := c.Get("sss")
	fmt.Println(err)
	fmt.Println(d, er)

}
