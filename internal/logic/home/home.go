package home

import (
	"context"
	"sviwo/internal/consts"
	"sviwo/internal/dao"
	"sviwo/internal/model"
	"sviwo/internal/service"
	"sviwo/pkg/gpool"
)

func init() {
	service.RegisterHome(New())
}

func New() *sHome {
	return &sHome{}
}

type sHome struct{}

func (s sHome) GetHomeData(ctx context.Context) (out *model.HomeDataOutput) {
	goPool := gpool.NewSyncParallelGPool(100, ctx)
	defer func() {
		// 关闭工作池并处理错误
		goPool.Shutdown()
		// 从错误通道处理和打印错误
		for err := range goPool.ErrChan() {
			panic(err)
		}
	}()
	userId := service.BizCtx().Get(ctx).Data.Get(consts.ContextKeyUserId)
	out = &model.HomeDataOutput{}
	goPool.Go(func(ctx context.Context) error {
		out.Version = service.Version().GetNewVersion(ctx)
		return nil
	})
	goPool.Go(func(ctx context.Context) error {
		userAuthStatus, err := dao.UserAuth.Ctx(ctx).Fields(dao.UserAuth.Columns().AuthStatus).
			One(dao.UserAuth.Columns().UserId, userId)
		if err != nil {
			return err
		}
		if !userAuthStatus.IsEmpty() {
			out.AuthStatus = userAuthStatus.GMap().GetVar(dao.UserAuth.Columns().AuthStatus).Int()
		}
		return err
	})
	goPool.Go(func(ctx context.Context) error {
		var (
			udTable = dao.UserDevice.Table()
			udCls   = dao.UserDevice.Columns()
			dTable  = dao.Device.Table()
			dCls    = dao.Device.Columns()
		)
		result, err := dao.UserDevice.Ctx(ctx).
			FieldsPrefix(udTable, udCls).
			FieldsPrefix(dTable, dCls.Nickname, dCls.DeviceName).
			LeftJoinOnField(dTable, dCls.DeviceId).
			WherePrefix(dTable, dCls.IsDelete, consts.DeleteOn).
			WherePrefix(udTable, udCls.UserId, userId).
			WherePrefix(udTable, udCls.IsSelect, consts.CarSelectYes).One()
		if err != nil {
			return err
		}
		if result.IsEmpty() {
			return nil
		}
		if err = result.Struct(&out); err != nil {
			return err
		}
		if consts.UserDeviceChild == out.UserDeviceType {
			out.SpeedLimit = nil
			out.MobileKey = nil
		}
		out.IsHavingCar = true
		s.findTDDeviceInfo(ctx, result.GMap().GetVar(dao.Device.Columns().DeviceName).String(), out)
		return nil
	})
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
	return
}
