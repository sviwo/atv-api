package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"sviwo/api/v1"
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
func (c cCar) EnabledMobileKey(ctx context.Context, req *v1.DelCarReq) (res *v1.EmptyFieldRes, err error) {
	service.Car().EnabledMobileKey(ctx, req.CarId)
	return
}

/*
EnabledSpeedLimit 开启/关闭速度限制
*/
func (c cCar) EnabledSpeedLimit(ctx context.Context, req *v1.DelCarReq) (res *v1.EmptyFieldRes, err error) {
	service.Car().EnabledSpeedLimit(ctx, req.CarId)
	return
}
