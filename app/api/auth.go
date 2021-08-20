package api

import (
	"evan-gf/app/model"
	"evan-gf/app/service"
	"evan-gf/library/response"
	jwt "github.com/gogf/gf-jwt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
	"time"
)

var (
	// The underlying JWT middleware.
	Auth        *jwt.GfJWTMiddleware
	IdentityKey = "id"
	DetailKey   = "user"
)

// Initialization function,
// rewrite this function to customized your own JWT settings.
func init() {
	authMiddleware, err := jwt.New(&jwt.GfJWTMiddleware{
		Realm:           "test zone",
		Key:             []byte("secret key"),
		Timeout:         time.Minute * 5,
		MaxRefresh:      time.Minute * 5,
		IdentityKey:     IdentityKey,
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
		Authenticator:   Authenticator,
		LoginResponse:   LoginResponse,
		RefreshResponse: RefreshResponse,
		LogoutResponse:  LogoutResponse,
		Unauthorized:    Unauthorized,
		PayloadFunc:     PayloadFunc,
		IdentityHandler: IdentityHandler,
	})
	if err != nil {
		glog.Fatal("JWT Error:" + err.Error())
	}
	Auth = authMiddleware
}

// PayloadFunc is a callback function that will be called during login.
// Using this function it is possible to add additional payload data to the webtoken.
// The data is then made available during requests via c.Get("JWT_PAYLOAD").
// Note that the payload is not encrypted.
// The attributes mentioned on jwt.io can't be used as keys for the map.
// Optional, by default no additional data will be set.
func PayloadFunc(data interface{}) jwt.MapClaims {
	claims := jwt.MapClaims{}
	params := data.(map[string]interface{})

	if len(params) > 0 {
		for k, v := range params {
			claims[k] = v
		}
	}
	return claims
}

// IdentityHandler get the identity from JWT and set the identity for every request
// Using this function, by r.GetParam("id") get identity
func IdentityHandler(r *ghttp.Request) interface{} {
	claims := jwt.ExtractClaims(r)
	return claims[Auth.IdentityKey]
}

// Unauthorized is used to define customized Unauthorized callback function.
func Unauthorized(r *ghttp.Request, code int, message string) {
	response.RetFail(r, code, message)
}

type LoginRes struct {
	Token  string `json:"token"`
	Expire string `json:"expire"`
}

// LoginResponse is used to define customized login-successful callback function.
func LoginResponse(r *ghttp.Request, code int, token string, expire time.Time) {
	response.RetSuccess(r, &LoginRes{
		token,
		expire.Format(time.RFC3339),
	})
}

// RefreshResponse is used to get a new token no matter current token is expired or not.
func RefreshResponse(r *ghttp.Request, code int, token string, expire time.Time) {
	response.RetSuccess(r, &LoginRes{
		token,
		expire.Format(time.RFC3339),
	})
}

// LogoutResponse is used to set token blacklist.
func LogoutResponse(r *ghttp.Request, code int) {
	response.RetSuccess(r, "")
}

// Authenticator is used to validate login parameters.
// It must return user data as user identifier, it will be stored in Claim Array.
// if your identityKey is 'id', your user data must have 'id'
// Check error (e) to determine the appropriate error message.
func Authenticator(r *ghttp.Request) (interface{}, error) {
	// todo:登录次数重试
	var (
		apiReq     *model.ApiLoginReq
		serviceReq *model.ServiceLoginReq
	)
	if err := r.Parse(&apiReq); err != nil {
		return "", err
	}
	if err := gconv.Struct(apiReq, &serviceReq); err != nil {
		return "", err
	}

	user, err := service.User.Login(serviceReq)
	if err != nil {
		return nil, err
	}
	return g.Map{
		IdentityKey: user.Id,
		DetailKey: model.UserLoginInfo{
			Id:       user.Id,
			Name:     user.Name,
			NickName: user.NickName,
		},
	}, err
}
