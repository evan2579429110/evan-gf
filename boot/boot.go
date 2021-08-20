package boot

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/swagger"
)

func init() {
	if g.Cfg().GetBool("swagger.debug") {
		s := g.Server()
		s.Plugin(&swagger.Swagger{})
	}
}
