package car

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gutil"
	"sviwo/internal/boot"
	"sviwo/internal/consts"
	"sviwo/internal/consts/enums"
	"sviwo/internal/dao"
	"sviwo/internal/model"
	"sviwo/internal/model/do"
	"sviwo/internal/model/entity"
	"sviwo/internal/service"
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
		"user_id", service.BizCtx().Get(ctx).Data.Get(consts.ContextKeyUserId),
	).ScanList(&out, "UserDeviceOutput"); err != nil {
		panic(err)
	}
	if out != nil {
		if err := dao.Device.Ctx(ctx).Where(
			"device_id", gdb.ListItemValuesUnique(out, "UserDeviceOutput", "DeviceId"),
		).ScanList(
			&out, "DeviceBase", "UserDeviceOutput", "device_id:DeviceId",
		); err != nil {
			panic(err)
		}
	}
	return
}

func (s sCar) DelCar(ctx context.Context, deviceId int64) {
	if _, err := dao.UserDevice.Ctx(ctx).
		Where(`user_id`, service.BizCtx().Get(ctx).Data.Get(consts.ContextKeyUserId)).
		Where(`device_id`, deviceId).Delete(); err != nil {
		panic(err)
	}
}

func (s sCar) CtlLamp(ctx context.Context, deviceId int64) {
}

func (s sCar) CtlHorn(ctx context.Context, deviceId int64) {
}

func (s sCar) CtlSwitchDT(ctx context.Context, in model.CtlSwitchDTInput) {
}

func (s sCar) CtlSwitchERT(ctx context.Context, in model.CtlSwitchERTInput) {
}

func (s sCar) EnabledMobileKey(ctx context.Context, deviceId int64) {
	userDevice := findUserDevice(ctx, deviceId)
	var mobileKey = true
	if userDevice.MobileKey {
		mobileKey = consts.CarMobileKeyNo
	}
	if err, _ := dao.UserDevice.Ctx(ctx).Data("mobile_key", mobileKey).
		Where("id", userDevice.Id).Update(); err != nil {
		panic(err)
	}
}

func findUserDevice(ctx context.Context, deviceId int64) entity.UserDevice {
	userDevice := *new(entity.UserDevice)
	if err := dao.UserDevice.Ctx(ctx).Where("device_id", deviceId).Where(
		"user_id", service.BizCtx().Get(ctx).Data.Get(consts.ContextKeyUserId),
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
	if err, _ := dao.UserDevice.Ctx(ctx).Data("speed_limit", speedLimit).
		Where("id", userDevice.Id).Update(); err != nil {
		panic(err)
	}
}

func (s sCar) BindingCar(ctx context.Context, carFrameCode string) {
	car := *new(entity.Device)
	if err := dao.Device.Ctx(ctx).Where("car_frame_code", carFrameCode).Scan(&car); err != nil {
		panic(err)
	}
	if gutil.IsEmpty(car) {
		panic(gerror.NewCode(enums.CarNotExists))
	}
	userId := service.BizCtx().Get(ctx).Data.Get(consts.ContextKeyUserId)
	count, err := dao.UserDevice.Ctx(ctx).Count("user_id", userId)
	if err != nil {
		panic(err)
	}
	userDevice := do.UserDevice{
		Id: boot.GID.Generate().Int64(), UserId: userId, DeviceId: car.DeviceId, CreateTime: gtime.Now(),
	}
	if count == 0 {
		userDevice.IsSelect = consts.CarSelectYes
	}
	if _, err = dao.UserDevice.Ctx(ctx).Data(userDevice).Insert(); err != nil {
		panic(err)
	}
}
