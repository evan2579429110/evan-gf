package service

import (
	"context"
	"github.com/gogf/gf/i18n/gi18n"
)

var I18n = i18nService{}

type i18nService struct {
}

func (i *i18nService) GetInstance(ctx context.Context) *gi18n.Manager {
	if Context.Get(ctx).I18n == nil {
		i18 := gi18n.New()
		i18.SetLanguage("zh-cn")
		Context.Get(ctx).I18n = i18
	}
	return Context.Get(ctx).I18n
}

func (i *i18nService) Get(ctx context.Context, s string, data ...interface{}) string {
	return i.GetInstance(ctx).Tf(ctx, `{#OrderPaid}`, 865271654, 99.8)
}
