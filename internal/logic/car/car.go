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
	if err := dao.UserCar.Ctx(ctx).Where(
		"user_id", service.BizCtx().Get(ctx).Data.Get(consts.ContextKeyUserId),
	).ScanList(&out, "UserCarOutput"); err != nil {
		panic(err)
	}
	if out != nil {
		if err := dao.Car.Ctx(ctx).Where(
			"car_id", gdb.ListItemValuesUnique(out, "UserCarOutput", "CarId"),
		).ScanList(
			&out, "CarBase", "UserCarOutput", "car_id:CarId",
		); err != nil {
			panic(err)
		}
	}
	return
}

func (s sCar) DelCar(ctx context.Context, carId int64) {
	if _, err := dao.UserCar.Ctx(ctx).
		Where(`user_id`, service.BizCtx().Get(ctx).Data.Get(consts.ContextKeyUserId)).
		Where(`car_id`, carId).Delete(); err != nil {
		panic(err)
	}
}

func (s sCar) CtlLamp(ctx context.Context, carId int64) {
}

func (s sCar) CtlHorn(ctx context.Context, carId int64) {
}

func (s sCar) CtlSwitchDriveType(ctx context.Context, in model.CtlSwitchDTInput) {
}

func (s sCar) CtlSwitchEnergyRecoveryType(ctx context.Context, in model.CtlSwitchERTInput) {
}

func (s sCar) EnabledMobileKey(ctx context.Context, carId int64) {
	userCar := findUserCar(ctx, carId)
	var mobileKey = true
	if userCar.MobileKey {
		mobileKey = consts.CarMobileKeyNo
	}
	if err, _ := dao.UserCar.Ctx(ctx).Data("mobile_key", mobileKey).
		Where("user_car_id", userCar.UserCarId).Update(); err != nil {
		panic(err)
	}
}

func findUserCar(ctx context.Context, carId int64) entity.UserCar {
	userCar := *new(entity.UserCar)
	if err := dao.UserCar.Ctx(ctx).Where("car_id", carId).Where(
		"user_id", service.BizCtx().Get(ctx).Data.Get(consts.ContextKeyUserId),
	).Scan(&userCar); err != nil {
		panic(err)
	}
	if gutil.IsEmpty(userCar) {
		panic(gerror.NewCode(enums.IllegalOperation))
	}
	return userCar
}

func (s sCar) EnabledSpeedLimit(ctx context.Context, carId int64) {
	userCar := findUserCar(ctx, carId)
	var speedLimit = true
	if userCar.SpeedLimit {
		speedLimit = consts.CarSpeedLimitNo
	}
	if err, _ := dao.UserCar.Ctx(ctx).Data("speed_limit", speedLimit).
		Where("user_car_id", userCar.UserCarId).Update(); err != nil {
		panic(err)
	}
}

func (s sCar) BindingCar(ctx context.Context, carFrameCode string) {
	car := *new(entity.Car)
	if err := dao.Car.Ctx(ctx).Where("car_frame_code", carFrameCode).Scan(&car); err != nil {
		panic(err)
	}
	if gutil.IsEmpty(car) {
		panic(gerror.NewCode(enums.CarNotExists))
	}
	userId := service.BizCtx().Get(ctx).Data.Get(consts.ContextKeyUserId)
	count, err := dao.UserCar.Ctx(ctx).Count("user_id", userId)
	if err != nil {
		panic(err)
	}
	userCar := do.UserCar{
		UserCarId: boot.GID.Generate().Int64(), UserId: userId, CarId: car.CarId, CreateTime: gtime.Now(),
	}
	if count == 0 {
		userCar.IsSelect = consts.CarSelectYes
	}
	if _, err = dao.UserCar.Ctx(ctx).Data(userCar).Insert(); err != nil {
		panic(err)
	}
}
