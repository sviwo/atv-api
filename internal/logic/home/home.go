package home

import (
	"context"
	"golang.org/x/sync/errgroup"
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
	withContext, _ := errgroup.WithContext(ctx)
	defer func() {
		if err := withContext.Wait(); err != nil {
			panic(err)
		}
	}()
	userId := service.BizCtx().Get(ctx).Data.Get(consts.ContextKeyUserId)
	out = &model.HomeDataOutput{}
	withContext.Go(func() error {
		out.Version = service.Version().GetNewVersion(ctx)
		return nil
	})
	withContext.Go(func() error {
		one, err := dao.AppParam.Ctx(ctx).
			Where(dao.AppParam.Columns().ParamConst, consts.ServicePhone).
			Where(dao.AppParam.Columns().IsDelete, consts.DeleteOn).One()
		if err != nil {
			return err
		}
		if !one.IsEmpty() {
			out.ServicePhone = one.GMap().GetVar(dao.AppParam.Columns().ParamValue).String()
		}
		return nil
	})
	withContext.Go(func() error {
		userAuthStatus, err := dao.UserAuth.Ctx(ctx).Fields(dao.UserAuth.Columns().AuthStatus).
			One(dao.UserAuth.Columns().UserId, userId)
		if err != nil {
			return err
		}
		if !userAuthStatus.IsEmpty() {
			out.AuthStatus = userAuthStatus.GMap().GetVar(dao.UserAuth.Columns().AuthStatus).Int()
		}
		return nil
	})
	withContext.Go(func() error {
		var (
			udTable = dao.UserDevice.Table()
			udCls   = dao.UserDevice.Columns()
			dTable  = dao.Device.Table()
			dCls    = dao.Device.Columns()
		)
		result, err := dao.UserDevice.Ctx(ctx).
			FieldsPrefix(udTable, udCls).
			FieldsPrefix(dTable, dCls.Nickname, dCls.DeviceName, dCls.BluetoothAddress, dCls.BluetoothSecretKey).
			LeftJoinOnField(dTable, dCls.DeviceId).
			WherePrefix(dTable, dCls.IsDelete, consts.DeleteOn).
			WherePrefix(udTable, udCls.UserId, userId).
			WherePrefix(udTable, udCls.IsSelect, consts.CarSelectYes).One()
		if err != nil {
			return err
		}
		if !result.IsEmpty() {
			if err = result.Struct(&out); err != nil {
				return err
			}
			if consts.UserDeviceChild == out.UserDeviceType {
				out.SpeedLimit = nil
				out.MobileKey = nil
			}
			out.IsHavingCar = true
			if err = s.findTDDeviceInfo(
				ctx, result.GMap().GetVar(dao.Device.Columns().DeviceName).String(), out,
			); err != nil {
				return err
			}
		}
		return nil
	})
	return
}

func (s sHome) findTDDeviceInfo(ctx context.Context, deviceName string, out *model.HomeDataOutput) error {
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
		return err
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
	return nil
}
