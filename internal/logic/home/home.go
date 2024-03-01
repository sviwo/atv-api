package home

import (
	"context"
	"sviwo/internal/consts"
	"sviwo/internal/dao"
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
	out = new(model.HomeDataOutput)
	one, err := dao.UserCar.Ctx(ctx).Where(
		"user_id", service.BizCtx().Get(ctx).Data.Get(consts.ContextKeyUserId),
	).Where("is_select", consts.CarSelectYes).One()
	if err != nil {
		panic(err)
	}
	if !one.IsEmpty() {
		if err := dao.Car.Ctx(ctx).Where("car_id", one.GMap().Get("carId")).
			Where("is_delete", consts.DeleteOn).Scan(&out); err != nil {
			panic(err)
		}
	}
	//todo 等待调用mqtt电量数据
	out.Version = service.Version().GetNewVersion(ctx)
	return
}
