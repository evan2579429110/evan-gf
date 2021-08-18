package service

import (
	"errors"
	"evan-gf/app/dao"
	"evan-gf/app/model"
	"fmt"
)

var User = userService{}

type userService struct{}

// 用户注册
func (s *userService) SignUp(r *model.UserServiceSignUpReq) error {
	if r.Nickname == "" {
		r.Nickname = r.Password
	}

	if dao.User.ExistByNickName(r.Nickname) {
		return errors.New(fmt.Sprintf("账号 %s 已存在", r.Passport))
	}
	if _, err := dao.User.Save(r); err != nil {
		return err
	}
	return nil
}
