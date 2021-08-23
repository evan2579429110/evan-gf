package model

type ApiLoginReq struct {
	Name     string `v:"required#请输入名称" json:"name"`
	Password string `v:"required#请输入密码" json:"password"`
}

/**
Service
*/
type ServiceLoginReq struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
