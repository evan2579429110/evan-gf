package router

import (
	"evan-gf/app/api"
	"evan-gf/app/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/", func(g *ghttp.RouterGroup) {
		g.Middleware(service.Middleware.CORS)
		g.POST("/login", api.Auth.LoginHandler)
		g.POST("/refresh", api.Auth.RefreshHandler)
		g.POST("/logout", api.Auth.LogoutHandler)
		// todo:注册

		g.Group("/user", func(g *ghttp.RouterGroup) {
			g.Middleware(middlewareAuth)
			//g.GET("info",api.User.Info)

		})

		//group.Group("/", func(group *ghttp.RouterGroup) {
		//	group.Middleware(service.Middleware.Auth)
		//	//group.ALL("/user/profile", api.User.Profile)
		//})

	})

}

// authHook is the HOOK function implements JWT logistics.
func middlewareAuth(r *ghttp.Request) {
	api.Auth.MiddlewareFunc()(r)
	r.Middleware.Next()
}
