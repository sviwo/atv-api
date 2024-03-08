package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"sviwo/api/v1"
	"sviwo/internal/service"
)

var AppVideos = cAppVideos{}

type cAppVideos struct{}

/*
GetAppVideosTree 获取APP视屏树
*/
func (c cAppVideos) GetAppVideosTree(ctx context.Context, req *v1.AppVideosTreeReq) (res []*v1.AppVideosTreeRes, err error) {
	if err = gconv.Structs(service.AppVideos().GetAppVideosTree(ctx), &res); err != nil {
		panic(err)
	}
	return
}
