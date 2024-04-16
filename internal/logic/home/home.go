package home

import (
	"context"
	"sviwo/internal/consts"
	"sviwo/internal/dao"
	"sviwo/internal/model"
	"sviwo/internal/model/entity"
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
	out = &model.HomeDataOutput{Version: service.Version().GetNewVersion(ctx)}
	result, err := dao.UserDevice.Ctx(ctx).
		Where(dao.UserDevice.Columns().UserId, service.BizCtx().Get(ctx).Data.Get(consts.ContextKeyUserId)).
		Where(dao.UserDevice.Columns().IsSelect, consts.CarSelectYes).One()
	if err != nil {
		panic(err)
	}
	if result.IsEmpty() {
		return
	}
	userDevice := new(entity.UserDevice)
	if err = result.Struct(&userDevice); err != nil {
		panic(err)
	}
	out.UserDeviceType = userDevice.UserDeviceType

	device := new(entity.Device)
	if err = dao.Device.Ctx(ctx).Fields(dao.Device.Columns().Nickname, dao.Device.Columns().DeviceName).
		Where(dao.Device.Columns().DeviceId, userDevice.DeviceId).
		Where(dao.Device.Columns().IsDelete, consts.DeleteOn).Scan(&device); err != nil {
		panic(err)
	}
	out.IsHavingCar = true
	out.Nickname = device.Nickname
	//根据物模型获取所有属性数据 根据情况选择
	keys := make([]string, 0)
	keys = append(keys, "RemainMile")
	keys = append(keys, "Electricity")
	keys = append(keys, "BatteryStatus")
	keys = append(keys, "LockedStatus")
	keys = append(keys, "GeoLocation")
	res, err := service.DevDevice().GetProperty(ctx, &model.DeviceGetPropertyInput{
		DeviceKey:    device.DeviceName,
		PropertyKeys: keys,
	})
	if err != nil {
		panic(err)
	}
	for _, re := range res {
		switch re.Key {
		case consts.RemainMile:
			out.RemainMile = re.Value.Float32()
		case consts.Electricity:
			out.Electricity = re.Value.Int()
		case consts.BatteryStatus:
			out.BatteryStatus = re.Value.Int()
		case consts.LockedStatus:
			out.LockedStatus = re.Value.Int()
		case consts.GeoLocation:
			out.GeoLocation = re.Value.Map()
		default:
			break
		}
	}
	return
}
