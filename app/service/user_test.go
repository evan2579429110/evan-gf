package service

import (
	"evan-gf/app/model"
	"fmt"
	"github.com/gogf/gf/frame/g"
	"testing"
)

func TestSignUp(t *testing.T) {
	c := g.Config()
	_ = c.AddPath("../../config")
	err := User.SignUp(&model.UserServiceSignUpReq{
		Passport: "123",
		Password: "123456",
		Nickname: "evan1",
	})
	fmt.Println(err)

}
