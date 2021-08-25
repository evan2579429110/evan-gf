package service

import (
	"evan-gf/app/model"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/guid"
)

// 中间件管理服务
var Middleware = middlewareService{}

type middlewareService struct{}

// 自定义上下文对象
func (s *middlewareService) Ctx(r *ghttp.Request) {
	// 初始化，务必最开始执行
	customCtx := &model.Context{
		I18n: nil,
	}
	Context.Init(r, customCtx)
	// 执行下一步请求逻辑
	r.Middleware.Next()
}

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

// log
func (s *middlewareService) Log(r *ghttp.Request) {
	if r.GetCtxVar("key").String() == "" {
		r.SetCtxVar("key", "[ tranceId "+guid.S()+" ] ["+r.RequestURI+"]")
	}

	if r.GetBodyString() != "" {
		g.Log().Info(r.GetCtxVar("key"), " [Request] ", r.GetBodyString())
	}

	r.Middleware.Next()

	if r.Response.BufferLength() > 0 {
		g.Log().Info(r.GetCtxVar("key"), " [Response] ", r.Response.BufferString())
	}
}
