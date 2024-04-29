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

/*func (s sHome) GetHomeData(ctx context.Context) (out *model.HomeDataOutput) {
	wg := sync.WaitGroup{}
	wg.Add(3)
	errChan := make(chan error, 4)
	defer func() {
		wg.Wait()
		select {
		case err := <-errChan:
			panic(err)
		default:
			close(errChan)
		}
	}()
	userId := service.BizCtx().Get(ctx).Data.Get(consts.ContextKeyUserId)
	out = &model.HomeDataOutput{}
	go func(ctx context.Context) {
		defer wg.Done()
		out.Version = service.Version().GetNewVersion(ctx)
	}(ctx)
	go func(ctx context.Context) {
		defer wg.Done()
		userAuthStatus, err := dao.UserAuth.Ctx(ctx).Fields(dao.UserAuth.Columns().AuthStatus).
			One(dao.UserAuth.Columns().UserId, userId)
		if err != nil {
			errChan <- err
		}
		if !userAuthStatus.IsEmpty() {
			out.AuthStatus = userAuthStatus.GMap().GetVar(dao.UserAuth.Columns().AuthStatus).Int()
		}
	}(ctx)
	go func(ctx context.Context) {
		defer wg.Done()
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
			errChan <- err
		}
		if !result.IsEmpty() {
			if err = result.Struct(&out); err != nil {
				errChan <- err
			}
			if consts.UserDeviceChild == out.UserDeviceType {
				out.SpeedLimit = nil
				out.MobileKey = nil
			}
			out.IsHavingCar = true
			if err = s.findTDDeviceInfo(
				ctx, result.GMap().GetVar(dao.Device.Columns().DeviceName).String(), out,
			); err != nil {
				errChan <- err
			}
		}
	}(ctx)
	return
}*/

func (s sHome) GetHomeData(ctx context.Context) (out *model.HomeDataOutput) {
	gPool := gpool.NewGPool2(4, ctx)
	defer func() {
		if err := gPool.Shutdown2(); err != nil {
			panic(err)
		}
	}()
	userId := service.BizCtx().Get(ctx).Data.Get(consts.ContextKeyUserId)
	out = &model.HomeDataOutput{}
	gPool.Go(func(ctx context.Context) error {
		out.Version = service.Version().GetNewVersion(ctx)
		return nil
	})
	gPool.Go(func(ctx context.Context) error {
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
	gPool.Go(func(ctx context.Context) error {
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
				//ctx, result.GMap().GetVar(dao.Device.Columns().DeviceName).String(), out,
				ctx, "sdasda", out,
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
