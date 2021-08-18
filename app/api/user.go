package api

import (
	"evan-gf/app/model"
	"evan-gf/app/service"
	"evan-gf/library/response"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

var User = new(userApi)

type userApi struct{}


// @summary 用户注册接口
// @tags    用户服务
// @produce json
// @param   entity  body model.UserApiSignUpReq true "注册请求"
// @router  /user/sign-up [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *userApi) SignUp(r *ghttp.Request) {
	var (
		apiReq     *model.UserApiSignUpReq
		serviceReq *model.UserServiceSignUpReq
	)
	if err := r.Parse(&apiReq); err != nil {
		response.RetFail(r, response.ERR_FAIL, err.Error())
	}

	if err := gconv.Struct(apiReq, &serviceReq); err != nil {
		response.RetFail(r, response.ERR_STRUCT, err.Error())
	}

	if err := service.User.SignUp(serviceReq); err != nil {
		response.RetFail(r, response.ERR_AUTH_SIGN_UP, err.Error())
	} else {
		response.RetSuccess(r, "注册成功")
	}

}
