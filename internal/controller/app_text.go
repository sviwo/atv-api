package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"sviwo/api/v1"
	"sviwo/internal/service"
)

var AppText = cAppText{}

type cAppText struct{}

/*
GetAppTextTree 获取app文本标题树
*/
func (c cAppText) GetAppTextTree(ctx context.Context, req *v1.AppTextTreeReq) (res []*v1.AppTextTreeRes, err error) {
	if err = gconv.Structs(service.AppText().GetAppTextTree(ctx), &res); err != nil {
		panic(err)
	}
	return
}

func (c cAppText) GetAppTextDetail(ctx context.Context, req *v1.AppTextDetailReq) (res *v1.AppTextDetailRes, err error) {
	if err = gconv.Structs(service.AppText().GetAppTextDetail(ctx, req.TextId), &res); err != nil {
		panic(err)
	}
	return
}
