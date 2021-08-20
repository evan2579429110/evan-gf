package api

import (
	"evan-gf/app/model"
	"evan-gf/app/service"
	"evan-gf/library/response"
	"github.com/gogf/gf/net/ghttp"
)

var User = new(userApi)

type userApi struct{}

// @summary 用户注册接口
// @tags    用户服务
// @produce json
// @param   entity  body model.ServiceRegisterReq true "注册请求"
// @router  /user/register [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *userApi) Register(r *ghttp.Request) {
	var serviceReq *model.ServiceRegisterReq
	if err := r.Parse(&serviceReq); err != nil {
		response.RetFail(r, response.ERR_STRUCT, err.Error())
	}

	if err := service.User.Register(serviceReq); err != nil {
		response.RetFail(r, response.ERR_AUTH_SIGN_UP, err.Error())
	}
	response.RetSuccess(r, "")
}

// @summary 获取用户详情信息
// @tags    用户服务
// @router  /user/profile [GET]
// @success 200 {object} model.UserLoginInfo "用户信息"
func (a *userApi) Profile(r *ghttp.Request) {
	response.RetSuccess(r, service.User.Profile(r))
}
