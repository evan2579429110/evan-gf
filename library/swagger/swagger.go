package swagger

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/swagger"
)

// 配置swagger
func InitSwagger() {
	if g.Cfg().GetBool("swagger.debug") {
		s := g.Server()
		s.Plugin(&swagger.Swagger{
			Info: swagger.SwaggerInfo{
				Title:       g.Cfg().GetString("swagger.title"),
				Version:     g.Cfg().GetString("swagger.version"),
				Description: g.Cfg().GetString("swagger.description") + "所有返回格式均为 { code: int , message : string , data : interface },以下展示内容为data",
			},
			Schemes: []string{"http", "https"},
			Host:    g.Cfg().GetString("swagger.host"),
		})

	}
}
