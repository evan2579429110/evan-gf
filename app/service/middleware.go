package service

import (
	"github.com/gogf/gf/net/ghttp"
)

// 中间件管理服务
var Middleware = middlewareService{}

type middlewareService struct{}

// 自定义上下文对象
//func (s *middlewareService) Ctx(r *ghttp.Request) {
//	// 初始化，务必最开始执行
//	customCtx := &model.Context{
//		Session: r.Session,
//	}
//	Context.Init(r, customCtx)
//	if user := Session.GetUser(r.Context()); user != nil {
//		customCtx.User = &model.ContextUser{
//			Id:       user.Id,
//			Passport: user.Passport,
//			Nickname: user.Nickname,
//		}
//	}
//	// 执行下一步请求逻辑
//	r.Middleware.Next()
//}

// 鉴权中间件，只有登录成功之后才能通过
//func (s *middlewareService) Auth(r *ghttp.Request) {
//	api.Auth.MiddlewareFunc()(r)
//	r.Middleware.Next()
//}

// 允许接口跨域请求
func (s *middlewareService) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
