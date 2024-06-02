package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gutil"
	"sviwo/api/v1"
	"sviwo/internal/service"
)

var AppMedia = cAppMedia{}

type cAppMedia struct{}

func (c cAppMedia) GetAppMediaTree(ctx context.Context, req *v1.AppMediaTreeReq) (res []*v1.AppMediaTreeRes, err error) {
	if err = gconv.Structs(service.AppMedia().GetAppMediaTree(ctx, req.PageType), &res); err != nil {
		panic(err)
	}
	return
}

func (c cAppMedia) GetAppMediaDetail(ctx context.Context, req *v1.AppMediaDetailReq) (res *v1.AppMediaDetailRes, err error) {
	content := service.AppMedia().GetAppMediaDetail(ctx, req.Id)
	if gutil.IsEmpty(content) {
		return
	}
	if err = gconv.Scan(content, &res); err != nil {
		panic(err)
	}
	return
}
