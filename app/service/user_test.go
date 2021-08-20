package service

import (
	"evan-gf/app/model"
	"fmt"
	"github.com/gogf/gf/frame/g"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestLogin(t *testing.T) {
	c := g.Config()
	_ = c.AddPath("../../config")

	if err := User.Register(&model.ServiceRegisterReq{
		Name:     "admin",
		NickName: "admin-name",
		Password: "123456",
	}); err != nil {
		fmt.Println("注册错误:", err.Error())
		return
	}

	if res, err := User.Login(&model.ServiceLoginReq{
		Name:     "admin",
		Password: "123456",
	}); err != nil {
		fmt.Println("登录错误:", err.Error())
	} else {
		fmt.Println("登录成功:", res)
	}

}

func TestPassword(t *testing.T) {
	fmt.Println(bcrypt.CompareHashAndPassword([]byte("$2a$04$PP1HO6oHiYjAdyXjNYolDORwh7f5oxWdrsTCfqq1UYGlOwMH4gIgK"), []byte("123456")))
}
