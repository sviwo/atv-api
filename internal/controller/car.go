package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"sviwo/api/v1"
	"sviwo/internal/dao"
	"sviwo/internal/model"
	"sviwo/internal/service"
)

var Car = cCar{}

type cCar struct{}

func (c cCar) CarCleanBind(ctx context.Context, req *v1.CarCleanBindReq) (res []*v1.EmptyFieldRes, err error) {
	dao.UserDevice.Ctx(ctx).Delete(dao.UserDevice.Columns().DeviceId, 10000)
	return
}

func (c cCar) GetCarList(ctx context.Context, req *v1.GetCarInfoReq) (res []*v1.GetCarListRes, err error) {
	if err = gconv.Structs(service.Car().GetCarList(ctx), &res); err != nil {
		panic(err)
	}
	return
}

func (c cCar) GetCarDetail(ctx context.Context, req *v1.GetCarDetailReq) (res *v1.GetCarDetailRes, err error) {
	out := service.Car().GetCarDetail(ctx, req.DeviceId)
	if out != nil {
		if err = gconv.Struct(out, &res); err != nil {
			panic(err)
		}
	}
	return
}

func (c cCar) GetCarKey(ctx context.Context, req *v1.CarKeyReq) (res *v1.CarKeyRes, err error) {
	res = &v1.CarKeyRes{CarKey: service.Car().GetCarKey(ctx)}
	return
}

func (c cCar) InviteBindCar(ctx context.Context, req *v1.InviteBindCarReq) (res *v1.EmptyFieldRes, err error) {
	service.Car().InviteBindCar(ctx, req.CarKey)
	return
}

func (c cCar) SwitchCar(ctx context.Context, req *v1.SwitchCarReq) (res *v1.EmptyFieldRes, err error) {
	service.Car().SwitchCar(ctx, req.DeviceId)
	return
}

func (c cCar) RemoveCar(ctx context.Context, req *v1.RemoveCarReq) (res *v1.EmptyFieldRes, err error) {
	service.Car().RemoveCar(ctx, req.UserDeviceId, req.DeviceId)
	return
}

func (c cCar) EditCarNickname(ctx context.Context, req *v1.EditCarNicknameReq) (res *v1.EmptyFieldRes, err error) {
	service.Car().EditCarNickname(ctx, req.Nickname)
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

func (c cCar) CtlSwitchDT(ctx context.Context, req *v1.CtlSwitchDTReq) (res *v1.EmptyFieldRes, err error) {
	data := model.CtlSwitchDTInput{}
	if err = gconv.Struct(req, &data); err != nil {
		panic(err)
	}
	service.Car().CtlSwitchDT(ctx, data)
	return
}

func (c cCar) CtlSwitchERT(ctx context.Context, req *v1.CtlSwitchERTReq) (res *v1.EmptyFieldRes, err error) {
	data := model.CtlSwitchERTInput{}
	if err = gconv.Struct(req, &data); err != nil {
		panic(err)
	}
	service.Car().CtlSwitchERT(ctx, data)
	return
}
func (c cCar) GetSimDataTraffic(ctx context.Context, req *v1.SimDataTrafficReq) (res *v1.SimDataTrafficRes, err error) {
	if err = gconv.Scan(service.Car().GetSimDataTraffic(ctx), &res); err != nil {
		panic(err)
	}
	return
}
