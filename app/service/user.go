package service

import (
	"errors"
	"evan-gf/app/dao"
	"evan-gf/app/model"
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"golang.org/x/crypto/bcrypt"
)

var User = userService{}

type userService struct{}

// 登录
func (s *userService) Login(serviceReq *model.ServiceLoginReq) (*model.User, error) {
	var q *model.User
	err := dao.User.Where("name = ? and status = ? ", serviceReq.Name, model.Normal).Scan(&q)
	if err != nil {
		return nil, errors.New("该用户不存在，请检查")
	}
	if err = bcrypt.CompareHashAndPassword([]byte(q.Password), []byte(serviceReq.Password)); err != nil {
		return nil, errors.New("用户/密码错误,请检查")
	}
	return q, err
}

// 注册
func (s *userService) Register(serviceReq *model.ServiceRegisterReq) error {
	var q *model.User
	dao.User.Where("name", serviceReq.Name).Scan(&q)
	if q != nil {
		return errors.New("该用户已存在")
	}

	p, _ := bcrypt.GenerateFromPassword([]byte(serviceReq.Password), bcrypt.MinCost)
	dao.User.Data(model.User{
		Name:     serviceReq.Name,
		NickName: serviceReq.NickName,
		Password: string(p),
		Status:   model.Normal,
	}).Insert()

	return nil
}

func (s *userService) Profile(r *ghttp.Request) interface{} {
	fmt.Println(r.Get("id"), r.Get("user"), r.GetBodyString())
	return r.Get("user")
}
