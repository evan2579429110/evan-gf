package response

import "github.com/gogf/gf/net/ghttp"

type JsonResponse struct {
	Code    int         `json:"code"`    // 错误码((0:成功, 1:失败, >1:错误码))
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 返回数据(业务接口定义具体数据结构)
}

func Json(r *ghttp.Request, code int, message string, data ...interface{}) {
	res := interface{}(nil)
	if len(data) > 0 {
		res = data[0]
	}
	r.Response.WriteJson(JsonResponse{
		Code:    code,
		Message: message,
		Data:    res,
	})
}

func JsonExit(r *ghttp.Request, err int, msg string, data ...interface{}) {
	Json(r, err, msg, data...)
	r.Exit()
}

// 成功返回
func RetSuccess(r *ghttp.Request, data interface{}) {
	JsonExit(r, ERR_OK, "", data)
}

// 失败返回
func RetFail(r *ghttp.Request, err int, msg string) {
	JsonExit(r, err, msg, nil)
}
