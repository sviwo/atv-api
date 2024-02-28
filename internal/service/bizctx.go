package service

import (
	"context"
	"github.com/gogf/gf/v2/container/gmap"
	"sviwo/internal/model"

	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IBizCtx interface {
		Init(r *ghttp.Request, customCtx *model.Context)
		Get(ctx context.Context) *model.Context
		SetData(ctx context.Context, data *gmap.Map)
	}
)

var (
	localBizCtx IBizCtx
)

func BizCtx() IBizCtx {
	if localBizCtx == nil {
		panic("implement not found for interface IBizCtx, forgot register?")
	}
	return localBizCtx
}

func RegisterBizCtx(i IBizCtx) {
	localBizCtx = i
}
