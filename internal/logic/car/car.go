package car

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/gogf/gf/v2/util/gutil"
	"sort"
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
	all, err := dao.UserDevice.Ctx(ctx).
		Where(dao.UserDevice.Columns().UserId, service.BizCtx().Get(ctx).Data.Get(consts.ContextKeyUserId)).
		All()
	if err != nil {
		panic(err)
	}
	if all.IsEmpty() {
		return
	}
	if err = all.Structs(&out); err != nil {
		panic(err)
	}
	if err = dao.Device.Ctx(ctx).
		WhereIn(dao.Device.Columns().DeviceId, all.Array(dao.Device.Columns().DeviceId)).
		Where(dao.Device.Columns().IsDelete, consts.DeleteOn).Scan(&out); err != nil {
		panic(err)
	}
	for _, ot := range out {
		ot.Mileage = s.findMileage(ctx, ot.DeviceName)
	}
	sort.Slice(out, func(i, j int) bool {
		return gconv.Int(out[i].IsSelect) > gconv.Int(out[j].IsSelect)
	})
	return
}

func (s sCar) findMileage(ctx context.Context, deviceName string) float32 {
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
	device := s.findDeviceInfo(ctx, deviceId)
	if device == nil {
		return
	}
	if err := gconv.Struct(device, &out); err != nil {
		panic(err)
	}
	all, err := dao.UserDevice.Ctx(ctx).
		Where(dao.UserDevice.Columns().DeviceId, device.DeviceId).
		Where(dao.UserDevice.Columns().UserDeviceType, consts.UserDeviceChild).
		Where(dao.UserDevice.Columns().MobileKey, consts.CarMobileKeyYes).
		All()
	if err != nil {
		panic(err)
	}
	if !all.IsEmpty() {
		userCarKeyList := make([]model.UserCarKeyOutput, all.Len())
		for i, record := range all {
			userCarKeyList[i].UserDeviceId = record.GMap().GetVar(dao.UserDevice.Columns().Id).Int64()
		}
		users, err := dao.User.Ctx(ctx).
			WhereIn(dao.User.Columns().UserId, all.Array(dao.UserDevice.Columns().UserId)).
			All()
		if err != nil {
			panic(err)
		}
		for i, user := range users {
			userCarKeyList[i].FirstName = user.GMap().GetVar(dao.User.Columns().FirstName).String()
			userCarKeyList[i].LastName = user.GMap().GetVar(dao.User.Columns().LastName).String()
			userCarKeyList[i].HeadImg = user.GMap().GetVar(dao.User.Columns().HeadImg).String()
		}
		out.UserCarKeyList = userCarKeyList
	}
	out.Mileage = s.findMileage(ctx, device.DeviceName)
	out.WarrantyTime = device.ActivateTime.AddDate(1, 0, 0)

	keys := make([]string, 0)
	keys = append(keys, "Limit")
	res, err := service.DevDevice().GetProperty(ctx, &model.DeviceGetPropertyInput{
		DeviceKey:    device.DeviceName,
		PropertyKeys: keys,
	})
	if err != nil {
		panic(err)
	}
	if !res[0].Value.IsEmpty() {
		out.TopSpeedHour = res[0].Value.Int()
	}
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

func (s sCar) EditCarNickname(ctx context.Context, nickname string) {
	userDevice := s.findUserDevice(ctx, nil)
	if userDevice == nil {
		panic(gerror.NewCode(enums.IllegalOperation))
	}
	if _, err := dao.Device.Ctx(ctx).Data(dao.Device.Columns().Nickname, nickname).
		Where(dao.Device.Columns().DeviceId, userDevice.DeviceId).Update(); err != nil {
		panic(err)
	}
}

func (s sCar) RemoveCar(ctx context.Context, userDeviceId, deviceId *int64) {
	if userDeviceId != nil {
		if _, err := dao.UserDevice.Ctx(ctx).
			Where(dao.UserDevice.Columns().Id, userDeviceId).
			Where(dao.UserDevice.Columns().UserDeviceType, consts.UserDeviceChild).
			Delete(); err != nil {
			panic(err)
		}
		return
	}
	userDevice := s.findUserDevice(ctx, deviceId)
	if userDevice == nil {
		panic(gerror.NewCode(enums.CarNotExists))
	}
	if userDevice.UserDeviceType == consts.UserDeviceTypeMain {
		panic(gerror.NewCode(enums.UnbindCarError))
	}
	userId := service.BizCtx().Get(ctx).Data.Get(consts.ContextKeyUserId)
	if _, err := dao.UserDevice.Ctx(ctx).
		Where(dao.UserDevice.Columns().UserId, userId).
		Where(dao.UserDevice.Columns().DeviceId, userDevice.DeviceId).
		Delete(); err != nil {
		panic(err)
	}

	count, err := dao.UserDevice.Ctx(ctx).
		Where(dao.UserDevice.Columns().IsSelect, consts.CarSelectYes).
		Where(dao.UserDevice.Columns().UserId, userId).
		Count()
	if err != nil {
		panic(err)
	}
	if count == 0 {
		if _, err = dao.UserDevice.Ctx(ctx).
			Data(dao.UserDevice.Columns().IsSelect, consts.CarSelectYes).
			Where(dao.UserDevice.Columns().UserId, userId).
			Limit(1).
			Update(); err != nil {
			panic(err)
		}
	}
}

func (s sCar) GetCarKey(ctx context.Context) (carKey string) {
	userId := service.BizCtx().Get(ctx).Data.Get(consts.ContextKeyUserId)
	utility.MethodReqLimit(ctx, "GetCarKey", userId, 30)

	result, err := dao.UserDevice.Ctx(ctx).Fields(dao.UserDevice.Columns().DeviceId).
		Where(dao.UserDevice.Columns().UserId, userId).
		Where(dao.UserDevice.Columns().UserDeviceType, consts.UserDeviceTypeMain).
		Where(dao.UserDevice.Columns().IsSelect, consts.CarSelectYes).
		One()
	if err != nil {
		panic(err)
	}
	if result.IsEmpty() {
		panic(gerror.NewCode(enums.ShareCarKeyError))
	}
	deviceId := result.GMap().GetVar(dao.Device.Columns().DeviceId).Int64()
	s.checkCarKeyLimit(ctx, deviceId)

	carKey = consts.CarKeyPrefix + grand.S(32)
	if err = g.Redis().SetEX(
		ctx,
		fmt.Sprintf(consts.RedisCarKey, carKey),
		deviceId,
		60*60,
	); err != nil {
		panic(err)
	}
	return
}

func (s sCar) checkCarKeyLimit(ctx context.Context, deviceId int64) {
	cot, err := dao.UserDevice.Ctx(ctx).
		Count(dao.UserDevice.Columns().DeviceId, deviceId)
	if err != nil {
		panic(err)
	}
	if cot > 4 {
		panic(gerror.NewCode(enums.CarKeyLimitError))
	}
}

func (s sCar) InviteBindCar(ctx context.Context, carKey string) {
	result, err := g.Redis().Get(
		ctx,
		fmt.Sprintf(consts.RedisCarKey, carKey),
	)
	if err != nil {
		panic(err)
	}
	if result.IsEmpty() {
		panic(gerror.NewCode(enums.CarKeyInvalidError))
	}

	deviceId := result.Int64()
	s.checkCarKeyLimit(ctx, deviceId)

	userId := service.BizCtx().Get(ctx).Data.Get(consts.ContextKeyUserId)
	cot, err := dao.UserDevice.Ctx(ctx).
		Where(dao.UserDevice.Columns().DeviceId, deviceId).
		Where(dao.UserDevice.Columns().UserId, userId).
		Count()
	if err != nil {
		panic(err)
	}
	if cot > 0 {
		panic(gerror.NewCode(enums.BindCarError))
	}

	userDevice := do.UserDevice{
		Id:             utility.GID.Generate().Int64(),
		UserId:         userId,
		DeviceId:       deviceId,
		UserDeviceType: consts.UserDeviceChild,
		CreateTime:     gtime.Now(),
	}
	count, err := dao.UserDevice.Ctx(ctx).Count(dao.UserDevice.Columns().UserId, userId)
	if err != nil {
		panic(err)
	}
	if count == 0 {
		userDevice.IsSelect = consts.CarSelectYes
	}
	if _, err = dao.UserDevice.Ctx(ctx).Data(userDevice).Insert(); err != nil {
		panic(err)
	}
	if _, err = g.Redis().Del(ctx, fmt.Sprintf(consts.RedisCarKey, carKey)); err != nil {
		panic(err)
	}
}

func (s sCar) CtlCar(ctx context.Context, instructions int) {
	device := s.findDeviceInfo(ctx, nil)
	if device == nil {
		panic(gerror.NewCode(enums.CarNotExists))
	}
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
	if _, err := service.DevDeviceProperty().Set(ctx, in); err != nil {
		panic(err)
	}
}

func (s sCar) findDeviceInfo(ctx context.Context, deviceId *int64) (device *entity.Device) {
	if deviceId == nil {
		userDevice := s.findUserDevice(ctx, nil)
		if userDevice == nil {
			return
		}
		deviceId = &userDevice.DeviceId
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
	userDevice := s.findUserDevice(ctx, nil)
	if userDevice == nil {
		panic(gerror.NewCode(enums.IllegalOperation))
	}
	var mobileKey = true
	if userDevice.MobileKey {
		mobileKey = consts.CarMobileKeyNo
	}
	if _, err := dao.UserDevice.Ctx(ctx).Data(dao.UserDevice.Columns().MobileKey, mobileKey).
		Where(dao.UserDevice.Columns().Id, userDevice.Id).Update(); err != nil {
		panic(err)
	}
}

func (s sCar) findUserDevice(ctx context.Context, deviceId *int64) (userDevice *entity.UserDevice) {
	sql := dao.UserDevice.Ctx(ctx)
	if deviceId == nil {
		sql.Where(dao.UserDevice.Columns().IsSelect, consts.CarSelectYes)
	} else {
		sql.Where(dao.UserDevice.Columns().DeviceId, deviceId)
	}
	result, err := sql.Where(
		dao.UserDevice.Columns().UserId, service.BizCtx().Get(ctx).Data.Get(consts.ContextKeyUserId),
	).One()
	if err != nil {
		panic(err)
	}
	if !result.IsEmpty() {
		if err = result.Struct(&userDevice); err != nil {
			panic(err)
		}
	}
	return
}

func (s sCar) EnabledSpeedLimit(ctx context.Context) {
	userDevice := s.findUserDevice(ctx, nil)
	if userDevice == nil {
		panic(gerror.NewCode(enums.IllegalOperation))
	}
	mp := make(map[string]any)
	mp = map[string]any{"Limit": 10}
	var speedLimit = true
	if userDevice.SpeedLimit {
		mp = map[string]any{"Limit": 60}
		speedLimit = consts.CarSpeedLimitNo
	}
	in := &model.DevicePropertyInput{
		DeviceKey: s.findDeviceInfo(ctx, &userDevice.DeviceId).DeviceName,
		Params:    mp,
	}
	if _, err := service.DevDeviceProperty().Set(ctx, in); err != nil {
		panic(err)
	}
	if _, err := dao.UserDevice.Ctx(ctx).Data(dao.UserDevice.Columns().SpeedLimit, speedLimit).
		Where(dao.UserDevice.Columns().Id, userDevice.Id).Update(); err != nil {
		panic(err)
	}
}
