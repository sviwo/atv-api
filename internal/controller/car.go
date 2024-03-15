package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"sviwo/api/v1"
	"sviwo/internal/model"
	"sviwo/internal/service"
)

var Car = cCar{}

type cCar struct{}

/*
GetCarList 获取用户车辆列表
*/
func (c cCar) GetCarList(ctx context.Context, req *v1.GetCarInfoReq) (res []*v1.GetCarInfoRes, err error) {
	if err = gconv.Structs(service.Car().GetCarList(ctx), &res); err != nil {
		panic(err)
	}
	return
}

/*
BindingCar 绑定车辆
*/
func (c cCar) BindingCar(ctx context.Context, req *v1.BindingCarReq) (res *v1.EmptyFieldRes, err error) {
	service.Car().BindingCar(ctx, req.CarFrameCode)
	return
}

/*
DelCar 删除（解绑）车辆
*/
func (c cCar) DelCar(ctx context.Context, req *v1.DelCarReq) (res *v1.EmptyFieldRes, err error) {
	service.Car().DelCar(ctx, req.CarId)
	return
}

/*
EnabledMobileKey 开启/关闭蓝牙钥匙
*/
func (c cCar) EnabledMobileKey(ctx context.Context, req *v1.EnabledMobileKeyReq) (res *v1.EmptyFieldRes, err error) {
	service.Car().EnabledMobileKey(ctx, req.CarId)
	return
}

/*
EnabledSpeedLimit 开启/关闭速度限制
*/
func (c cCar) EnabledSpeedLimit(ctx context.Context, req *v1.EnabledSpeedLimitReq) (res *v1.EmptyFieldRes, err error) {
	service.Car().EnabledSpeedLimit(ctx, req.CarId)
	return
}

/*
CtlLamp 控制车灯
*/
func (c cCar) CtlLamp(ctx context.Context, req *v1.CtlLampReq) (res *v1.EmptyFieldRes, err error) {
	service.Car().CtlLamp(ctx, req.CarId)
	return
}

/*
CtlHorn 控制喇叭
*/
func (c cCar) CtlHorn(ctx context.Context, req *v1.CtlHornReq) (res *v1.EmptyFieldRes, err error) {
	service.Car().CtlHorn(ctx, req.CarId)
	return
}

/*
CtlSwitchDriveType 切换驾驶模式
*/
func (c cCar) CtlSwitchDriveType(ctx context.Context, req *v1.CtlSwitchDTReq) (res *v1.EmptyFieldRes, err error) {
	data := model.CtlSwitchDTInput{}
	if err = gconv.Struct(req, &data); err != nil {
		panic(err)
	}
	service.Car().CtlSwitchDriveType(ctx, data)
	return
}

/*
CtlSwitchEnergyRecoveryType 切换动能回收模式
*/
func (c cCar) CtlSwitchEnergyRecoveryType(ctx context.Context, req *v1.CtlSwitchERTReq) (
	res *v1.EmptyFieldRes, err error) {
	data := model.CtlSwitchERTInput{}
	if err = gconv.Struct(req, &data); err != nil {
		panic(err)
	}
	service.Car().CtlSwitchEnergyRecoveryType(ctx, data)
	return
}
