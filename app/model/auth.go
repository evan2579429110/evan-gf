package model

type ApiLoginReq struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

/**
Service
*/
type ServiceLoginReq struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
