package appvideos

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"sviwo/internal/consts"
	"sviwo/internal/dao"
	"sviwo/internal/model"
	"sviwo/internal/service"
	"sviwo/pkg/utility"
)

func init() {
	service.RegisterAppVideos(New())
}

func New() *sAppVideos {
	return &sAppVideos{}
}

type sAppVideos struct{}

func (s sAppVideos) GetAppVideosTree(ctx context.Context) (out []*model.AppVideosTreeOutput) {
	var array []*model.AppVideosTreeOutput
	if err := dao.AppVideos.Ctx(ctx).
		Where("enable", consts.EnableDisplay).
		Where("is_delete", consts.DeleteOn).
		OrderAsc("orders").Scan(&array); err != nil {
		panic(err)
	}
	if err := gconv.Structs(utility.BuildTree(gconv.Maps(array)), &out); err != nil {
		panic(err)
	}
	return
}
