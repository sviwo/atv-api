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
GetAppTextList 获取app文本标题列表
*/
func (c cAppText) GetAppTextList(ctx context.Context, req *v1.AppTextListReq) (res []*v1.AppTextListRes, err error) {
	if err = gconv.Structs(service.AppText().GetAppTextList(ctx), &res); err != nil {
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
