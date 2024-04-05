package home

import (
	"context"
	"github.com/gogf/gf/v2/os/glog"
	"sviwo/internal/model"
	"sviwo/internal/service"
)

func init() {
	service.RegisterHome(New())
}

func New() *sHome {
	return &sHome{}
}

type sHome struct{}

func (s sHome) GetHomeData(ctx context.Context) (out *model.HomeDataOutput) {
	//todo 等待调用硬件电量和公里数
	out = new(model.HomeDataOutput)
	res, _ := service.DevDevice().GetLatestProperty(ctx, "sviwo_atv")
	glog.Print(ctx, res)
	out.Version = service.Version().GetNewVersion(ctx)
	return
}
