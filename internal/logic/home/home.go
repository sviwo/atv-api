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
	//根据物模型获取所有属性数据 根据情况选择
	res, _ := service.DevDevice().GetLatestProperty(ctx, "sviwo_atv")
	//根据物模型获取单个属性数据 根据情况选择
	//res, _ := service.DevDevice().GetProperty(ctx, &model.DeviceGetPropertyInput{
	//	DeviceKey:   "sviwo_atv",
	//	PropertyKey: "Vehspeed",
	//})
	glog.Print(ctx, res)
	out.Version = service.Version().GetNewVersion(ctx)
	return
}
