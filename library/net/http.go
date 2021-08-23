package net

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

type Http struct {
	client *ghttp.Client
}

func NewServer() *Http {
	return &Http{
		client: g.Client(),
	}
}

// data 目前支持struct与map,其他类型不支持
func (h *Http) Get(host string, data ...interface{}) (interface{}, error) {
	uri := host
	// todo: 增加格式转换成a=b&c=d
	//var d map[string]interface{}
	//switch data.(type) {
	//	case struct{}:
	//		d = gconv.Map(data)
	//	case map[string]interface{}:
	//		d = data.(map[string]interface{})
	//}
	if r, err := h.client.Get(uri); err != nil {
		return nil, err
	} else {
		defer r.Close()
		g.Log().Info(r.Raw())
		return r.ReadAllString(), nil
	}
}

func (h *Http) Post(host string, data interface{}) (interface{}, error) {
	if r, err := g.Client().Post(host, data); err != nil {
		return nil, err
	} else {
		defer r.Close()
		g.Log().Info(r.Raw())
		return r.ReadAllString(), nil
	}
}

func (h *Http) SetCookie(key, value string) {
	h.client.SetCookie(key, value)
}

func (h *Http) SetHeader(key, value string) {
	h.client.SetHeader(key, value)
}
