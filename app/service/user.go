package service

import (
	"errors"
	"evan-gf/app/dao"
	"evan-gf/app/model"
	"golang.org/x/crypto/bcrypt"
)

var User = userService{}

type userService struct{}

// 登录
func (s *userService) Login(serviceReq *model.ServiceLoginReq) (*model.User, error) {
	var q model.User
	err := dao.User.Where("name = ? and status = ? ", serviceReq.Name, model.Normal).Scan(&q)
	if err != nil {
		return nil, errors.New("该用户不存在，请检查")
	}
	if err = bcrypt.CompareHashAndPassword([]byte(q.Password), []byte(serviceReq.Password)); err != nil {
		return nil, errors.New("用户/密码错误,请检查")
	}
	return &q, err
}
