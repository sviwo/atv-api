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
	userId := service.BizCtx().Get(ctx).Data.Get(consts.ContextKeyUserId)
	result, err := dao.UserDevice.Ctx(ctx).
		Where(dao.UserDevice.Columns().UserId, userId).
		Where(dao.UserDevice.Columns().IsSelect, consts.CarSelectYes).One()
	if err != nil {
		panic(err)
	}
	if result.IsEmpty() {
		return
	}
	if err = result.Struct(&out); err != nil {
		panic(err)
	}
	if consts.UserDeviceChild == out.UserDeviceType {
		out.SpeedLimit = nil
		out.MobileKey = nil
	}

	userAuthStatus, err := dao.UserAuth.Ctx(ctx).Fields(dao.UserAuth.Columns().AuthStatus).
		One(dao.UserAuth.Columns().UserId, userId)
	if err != nil {
		panic(err)
	}
	if !userAuthStatus.IsEmpty() {
		out.AuthStatus = userAuthStatus.GMap().GetVar(dao.UserAuth.Columns().AuthStatus).Int()
	}

	device := new(entity.Device)
	if err = dao.Device.Ctx(ctx).Fields(dao.Device.Columns().Nickname, dao.Device.Columns().DeviceName).
		Where(dao.Device.Columns().DeviceId, result.GMap().GetVar(dao.UserDevice.Columns().DeviceId).Int64()).
		Where(dao.Device.Columns().IsDelete, consts.DeleteOn).Scan(&device); err != nil {
		panic(err)
	}
	out.IsHavingCar = true
	out.Nickname = device.Nickname

	s.findTDDeviceInfo(ctx, device.DeviceName, out)
	return
}

func (s sHome) findTDDeviceInfo(ctx context.Context, deviceName string, out *model.HomeDataOutput) {
	//根据物模型获取所有属性数据 根据情况选择
	keys := make([]string, 0)
	keys = append(keys, consts.RemainMileStr)
	keys = append(keys, consts.ElectricityStr)
	keys = append(keys, consts.BatteryStatusStr)
	keys = append(keys, consts.LockedStatusStr)
	keys = append(keys, consts.GeoLocationStr)
	res, err := service.DevDevice().GetProperty(ctx, &model.DeviceGetPropertyInput{
		DeviceKey:    deviceName,
		PropertyKeys: keys,
	})
	if err != nil {
		panic(err)
	}
	for _, re := range res {
		switch re.Key {
		case consts.RemainMileStr:
			out.RemainMile = re.Value.Float32()
		case consts.ElectricityStr:
			out.Electricity = re.Value.Int()
		case consts.BatteryStatusStr:
			out.BatteryStatus = re.Value.Int()
		case consts.LockedStatusStr:
			out.LockedStatus = re.Value.Int()
		case consts.GeoLocationStr:
			out.GeoLocation = re.Value.Map()
		default:
			break
		}
	}
}
