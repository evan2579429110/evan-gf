package boot

import (
	"evan-gf/library/swagger"
)

func init() {
	swagger.InitSwagger()

	// 加载缓存
	//cache.NewCache(cache.NewRedisCache())
}
