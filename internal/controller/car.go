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

func (c cCar) GetCarList(ctx context.Context, req *v1.GetCarInfoReq) (res []*v1.GetCarListRes, err error) {
	if err = gconv.Structs(service.Car().GetCarList(ctx), &res); err != nil {
		panic(err)
	}
	return
}

func (c cCar) GetCarDetail(ctx context.Context, req *v1.GetCarDetailReq) (res *v1.GetCarDetailRes, err error) {
	if err = gconv.Struct(service.Car().GetCarDetail(ctx, req.DeviceId), &res); err != nil {
		panic(err)
	}
	return
}

func (c cCar) SwitchCar(ctx context.Context, req *v1.SwitchCarReq) (res *v1.EmptyFieldRes, err error) {
	service.Car().SwitchCar(ctx, req.DeviceId)
	return
}

func (c cCar) BindingCar(ctx context.Context, req *v1.BindingCarReq) (res *v1.EmptyFieldRes, err error) {
	service.Car().BindingCar(ctx, req.UserId, req.DeviceName)
	return
}

func (c cCar) DelCar(ctx context.Context, req *v1.DelCarReq) (res *v1.EmptyFieldRes, err error) {
	service.Car().DelCar(ctx, req.DeviceId)
	return
}

func (c cCar) EnabledMobileKey(ctx context.Context, req *v1.EnabledMobileKeyReq) (res *v1.EmptyFieldRes, err error) {
	service.Car().EnabledMobileKey(ctx)
	return
}

func (c cCar) EnabledSpeedLimit(ctx context.Context, req *v1.EnabledSpeedLimitReq) (res *v1.EmptyFieldRes, err error) {
	service.Car().EnabledSpeedLimit(ctx)
	return
}

func (c cCar) CtlCar(ctx context.Context, req *v1.CtlCarReq) (res *v1.EmptyFieldRes, err error) {
	service.Car().CtlCar(ctx, req.Instructions)
	return
}

func (c cCar) CtlSwitchDriveType(ctx context.Context, req *v1.CtlSwitchDTReq) (res *v1.EmptyFieldRes, err error) {
	data := model.CtlSwitchDTInput{}
	if err = gconv.Struct(req, &data); err != nil {
		panic(err)
	}
	service.Car().CtlSwitchDT(ctx, data)
	return
}

func (c cCar) CtlSwitchEnergyRecoveryType(ctx context.Context, req *v1.CtlSwitchERTReq) (
	res *v1.EmptyFieldRes, err error) {
	data := model.CtlSwitchERTInput{}
	if err = gconv.Struct(req, &data); err != nil {
		panic(err)
	}
	service.Car().CtlSwitchERT(ctx, data)
	return
}
