package router

import (
	"evan-gf/app/api"
	"evan-gf/app/service"
	"evan-gf/library/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/", func(g *ghttp.RouterGroup) {
		g.Middleware(service.Middleware.CORS, service.Middleware.Ctx, service.Middleware.Log)
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

	s.BindStatusHandlerByMap(map[int]ghttp.HandlerFunc{
		403: func(r *ghttp.Request) { response.RetFail(r, 403, service.I18n.GetC(r.Context(), "AccessDenied")) },
		404: func(r *ghttp.Request) { response.RetFail(r, 404, service.I18n.GetC(r.Context(), "NotFound")) },
		500: func(r *ghttp.Request) {
			response.RetFail(r, 500, service.I18n.GetC(r.Context(), "ServerInternalError"))
		},
	})

}

// authHook is the HOOK function implements JWT logistics.
func middlewareAuth(r *ghttp.Request) {
	api.Auth.MiddlewareFunc()(r)
	r.Middleware.Next()
}
