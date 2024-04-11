package car

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gutil"
	"sviwo/internal/consts"
	"sviwo/internal/consts/enums"
	"sviwo/internal/dao"
	"sviwo/internal/model"
	"sviwo/internal/model/do"
	"sviwo/internal/model/entity"
	"sviwo/internal/service"
	"sviwo/pkg/utility"
)

func init() {
	service.RegisterCar(New())
}

func New() *sCar {
	return &sCar{}
}

type sCar struct{}

func (s sCar) GetCarList(ctx context.Context) (out []*model.QueryCarOutput) {
	if err := dao.UserDevice.Ctx(ctx).Where(
		dao.UserDevice.Columns().UserId, service.BizCtx().Get(ctx).Data.Get(consts.ContextKeyUserId),
	).ScanList(&out, "UserDeviceOutput"); err != nil {
		panic(err)
	}
	if out != nil {
		if err := dao.Device.Ctx(ctx).Where(
			dao.Device.Columns().DeviceId, gdb.ListItemValuesUnique(out, "UserDeviceOutput", "DeviceId"),
		).ScanList(
			&out, "DeviceBase", "UserDeviceOutput", "device_id:DeviceId",
		); err != nil {
			panic(err)
		}
	}
	return
}

func (s sCar) DelCar(ctx context.Context, deviceId int64) {
	count, err := dao.UserDevice.Ctx(ctx).Count(dao.UserDevice.Columns().DeviceId, deviceId)
	if err != nil {
		panic(err)
	}
	if count <= 1 {
		panic(gerror.NewCode(enums.UnbindCarError))
	}
	if _, err = dao.UserDevice.Ctx(ctx).
		Where(dao.UserDevice.Columns().UserId, service.BizCtx().Get(ctx).Data.Get(consts.ContextKeyUserId)).
		Where(dao.UserDevice.Columns().DeviceId, deviceId).Delete(); err != nil {
		panic(err)
	}
}

func (s sCar) CtlCar(ctx context.Context, instructions int) {
	device := findDeviceInfo(ctx)
	mp := make(map[string]any)
	switch instructions {
	case consts.Light:
		mp = map[string]any{"Light": 1}
	case consts.Speaker:
		mp = map[string]any{"speaker": 1}
	default:
		panic(gerror.NewCode(enums.IllegalOperation))
	}
	in := &model.DevicePropertyInput{
		DeviceKey: device.DeviceName,
		Params:    mp,
	}
	_, err := service.DevDeviceProperty().Set(ctx, in)
	if err != nil {
		panic(err)
	}
}

func findDeviceInfo(ctx context.Context) (device *entity.Device) {
	result, err := dao.UserDevice.Ctx(ctx).Fields(dao.UserDevice.Columns().DeviceId).
		Where(dao.UserDevice.Columns().UserId, service.BizCtx().Get(ctx).Data.Get(consts.ContextKeyUserId)).
		Where(dao.UserDevice.Columns().IsSelect, consts.CarSelectYes).One()
	if err != nil {
		panic(err)
	}
	if result.IsEmpty() {
		panic(gerror.NewCode(enums.IllegalOperation))
	}
	if err = result.Struct(&device); err != nil {
		panic(err)
	}
	if err = dao.Device.Ctx(ctx).Fields(dao.Device.Columns().DeviceName).
		Where(dao.Device.Columns().DeviceId, device.DeviceId).
		Where(dao.Device.Columns().IsDelete, consts.DeleteOn).Scan(&device); err != nil {
		panic(err)
	}
	return
}

func (s sCar) CtlSwitchDT(ctx context.Context, in model.CtlSwitchDTInput) {
	if _, err := dao.UserDevice.Ctx(ctx).
		Data(
			dao.UserDevice.Columns().DrivingMode, in.DrivingModeType,
			dao.UserDevice.Columns().UpdateTime, gtime.Now(),
		).Where(dao.UserDevice.Columns().UserId, service.BizCtx().Get(ctx).Data.Get(consts.ContextKeyUserId)).
		Where(dao.UserDevice.Columns().IsSelect, consts.CarSelectYes).
		Update(); err != nil {
		panic(err)
	}
}

func (s sCar) CtlSwitchERT(ctx context.Context, in model.CtlSwitchERTInput) {
	if _, err := dao.UserDevice.Ctx(ctx).
		Data(
			dao.UserDevice.Columns().EnergyRecovery, in.EnergyRecoveryType,
			dao.UserDevice.Columns().UpdateTime, gtime.Now(),
		).Where(dao.UserDevice.Columns().UserId, service.BizCtx().Get(ctx).Data.Get(consts.ContextKeyUserId)).
		Where(dao.UserDevice.Columns().IsSelect, consts.CarSelectYes).
		Update(); err != nil {
		panic(err)
	}
}

func (s sCar) EnabledMobileKey(ctx context.Context, deviceId int64) {
	userDevice := findUserDevice(ctx, deviceId)
	var mobileKey = true
	if userDevice.MobileKey {
		mobileKey = consts.CarMobileKeyNo
	}
	if _, err := dao.UserDevice.Ctx(ctx).Data(dao.UserDevice.Columns().MobileKey, mobileKey).
		Where(dao.UserDevice.Columns().Id, userDevice.Id).Update(); err != nil {
		panic(err)
	}
}

func findUserDevice(ctx context.Context, deviceId int64) entity.UserDevice {
	userDevice := *new(entity.UserDevice)
	if err := dao.UserDevice.Ctx(ctx).Where(dao.UserDevice.Columns().DeviceId, deviceId).Where(
		dao.UserDevice.Columns().UserId, service.BizCtx().Get(ctx).Data.Get(consts.ContextKeyUserId),
	).Scan(&userDevice); err != nil {
		panic(err)
	}
	if gutil.IsEmpty(userDevice) {
		panic(gerror.NewCode(enums.IllegalOperation))
	}
	return userDevice
}

func (s sCar) EnabledSpeedLimit(ctx context.Context, deviceId int64) {
	userDevice := findUserDevice(ctx, deviceId)
	var speedLimit = true
	if userDevice.SpeedLimit {
		speedLimit = consts.CarSpeedLimitNo
	}
	if _, err := dao.UserDevice.Ctx(ctx).Data(dao.UserDevice.Columns().SpeedLimit, speedLimit).
		Where(dao.UserDevice.Columns().Id, userDevice.Id).Update(); err != nil {
		panic(err)
	}
}

func (s sCar) BindingCar(ctx context.Context, userId int64, deviceName string) {
	cot, err := dao.User.Ctx(ctx).Count(dao.User.Columns().UserId, userId)
	if err != nil {
		panic(err)
	}
	if cot == 0 {
		panic(gerror.NewCode(enums.UserNotExists))
	}
	device := *new(entity.Device)
	if err = dao.Device.Ctx(ctx).Where(dao.Device.Columns().DeviceName, deviceName).Scan(&device); err != nil {
		panic(err)
	}
	if gutil.IsEmpty(device) {
		panic(gerror.NewCode(enums.CarNotExists))
	}
	count, err := dao.UserDevice.Ctx(ctx).Count(dao.UserDevice.Columns().UserId, userId)
	if err != nil {
		panic(err)
	}
	userDevice := do.UserDevice{
		Id: utility.GID.Generate().Int64(), UserId: userId, DeviceId: device.DeviceId, CreateTime: gtime.Now(),
	}
	if count == 0 {
		userDevice.IsSelect = consts.CarSelectYes
	}
	if _, err = dao.UserDevice.Ctx(ctx).Data(userDevice).Insert(); err != nil {
		panic(err)
	}
}
