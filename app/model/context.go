package model

import (
	"github.com/gogf/gf/i18n/gi18n"
)

const (
	// 上下文变量存储键名
	ContextKey = "ContextKey"
)

// 请求上下文结构
type Context struct {
	I18n *gi18n.Manager // 当前i18n管理对象
}

// 请求上下文中的
