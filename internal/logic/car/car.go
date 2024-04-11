package car

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
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
	deviceIds, err := dao.UserDevice.Ctx(ctx).Fields(dao.UserDevice.Columns().DeviceId).
		Where(dao.UserDevice.Columns().UserId, service.BizCtx().Get(ctx).Data.Get(consts.ContextKeyUserId)).All()
	if err != nil {
		panic(err)
	}
	if deviceIds.IsEmpty() {
		return
	}
	if err = dao.Device.Ctx(ctx).
		WhereIn(dao.Device.Columns().DeviceId, deviceIds.Array(dao.Device.Columns().DeviceId)).
		Where(dao.Device.Columns().IsDelete, consts.DeleteOn).Scan(&out); err != nil {
		panic(err)
	}
	for _, ot := range out {
		ot.Mileage = findMileage(ctx, ot.DeviceName)
	}
	return
}

func findMileage(ctx context.Context, deviceName string) float32 {
	keys := make([]string, 0)
	keys = append(keys, "Mileage")
	res, err := service.DevDevice().GetProperty(ctx, &model.DeviceGetPropertyInput{
		DeviceKey:    deviceName,
		PropertyKeys: keys,
	})
	if err != nil {
		panic(err)
	}
	if !gutil.IsEmpty(res) {
		return res[0].Value.Float32()
	}
	return 0
}

func (s sCar) GetCarDetail(ctx context.Context, deviceId *int64) (out *model.UserDeviceOutput) {
	if err := gconv.Struct(findDeviceInfo(ctx, deviceId), &out); err != nil {
		panic(err)
	}
	out.Mileage = findMileage(ctx, out.DeviceName)
	out.WarrantyTime = out.ActivateTime.AddDate(1, 0, 0)
	return
}

func (s sCar) SwitchCar(ctx context.Context, deviceId int64) {
	if err := g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		userId := service.BizCtx().Get(ctx).Data.Get(consts.ContextKeyUserId)
		result, err := dao.UserDevice.Ctx(ctx).Data(dao.UserDevice.Columns().IsSelect, consts.CarSelectNo).
			Where(dao.UserDevice.Columns().UserId, userId).
			Where(dao.UserDevice.Columns().IsSelect, consts.CarSelectYes).
			WhereNot(dao.UserDevice.Columns().DeviceId, deviceId).
			Update()
		if err != nil {
			return err
		}
		if affected, _ := result.RowsAffected(); affected > 0 {
			if _, err = dao.UserDevice.Ctx(ctx).Data(dao.UserDevice.Columns().IsSelect, consts.CarSelectYes).
				Where(dao.UserDevice.Columns().UserId, userId).
				Where(dao.UserDevice.Columns().DeviceId, deviceId).
				Update(); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		panic(err)
	}
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
	device := findDeviceInfo(ctx, nil)
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

func findDeviceInfo(ctx context.Context, deviceId *int64) (device *entity.Device) {
	if deviceId == nil {
		result, err := dao.UserDevice.Ctx(ctx).Fields(dao.UserDevice.Columns().DeviceId).
			Where(dao.UserDevice.Columns().UserId, service.BizCtx().Get(ctx).Data.Get(consts.ContextKeyUserId)).
			Where(dao.UserDevice.Columns().IsSelect, consts.CarSelectYes).One()
		if err != nil {
			panic(err)
		}
		if result.IsEmpty() {
			return
		}
		getVar := result.GMap().GetVar(dao.Device.Columns().DeviceId).Int64()
		deviceId = &getVar
	}
	if err := dao.Device.Ctx(ctx).
		Where(dao.Device.Columns().DeviceId, &deviceId).
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

func (s sCar) EnabledMobileKey(ctx context.Context) {
	userDevice := findUserDevice(ctx)
	var mobileKey = true
	if userDevice.MobileKey {
		mobileKey = consts.CarMobileKeyNo
	}
	if _, err := dao.UserDevice.Ctx(ctx).Data(dao.UserDevice.Columns().MobileKey, mobileKey).
		Where(dao.UserDevice.Columns().Id, userDevice.Id).Update(); err != nil {
		panic(err)
	}
}

func findUserDevice(ctx context.Context) entity.UserDevice {
	userDevice := *new(entity.UserDevice)
	if err := dao.UserDevice.Ctx(ctx).
		Where(dao.UserDevice.Columns().IsSelect, consts.CarSelectYes).
		Where(dao.UserDevice.Columns().UserId, service.BizCtx().Get(ctx).Data.Get(consts.ContextKeyUserId)).
		Scan(&userDevice); err != nil {
		panic(err)
	}
	if gutil.IsEmpty(userDevice) {
		panic(gerror.NewCode(enums.IllegalOperation))
	}
	return userDevice
}

func (s sCar) EnabledSpeedLimit(ctx context.Context) {
	userDevice := findUserDevice(ctx)
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
