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
		// todo: 1.考虑账号被删除jwt短时间也能登录的问题,2.登录次数重试
		// token续签问题由前端解决
		g.POST("/login", api.Auth.LoginHandler)
		g.POST("/refresh", api.Auth.RefreshHandler)
		g.POST("/logout", api.Auth.LogoutHandler)

		g.Group("/", func(g *ghttp.RouterGroup) {
			g.Middleware(middlewareAuth)
			g.POST("/user/register", api.User.Register)
			g.GET("/user/profile", api.User.Profile)

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
