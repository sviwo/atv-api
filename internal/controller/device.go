package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"sviwo/api/v1"
	"sviwo/internal/model"
	"sviwo/internal/service"
)

var Device = cDevice{}

type cDevice struct{}

func (c cDevice) GetDeviceSecret(ctx context.Context, req *v1.DeviceSecretReq) (res *v1.DeviceSecretRes, err error) {
	if err = gconv.Struct(service.DevDevice().GetDeviceSecret(ctx, req.DeviceName), &res); err != nil {
		panic(err)
	}
	return
}

func (c cDevice) CheckDeviceBind(ctx context.Context, req *v1.CheckDeviceBindReq) (res *v1.EmptyFieldRes, err error) {
	service.DevDevice().CheckDeviceBind(ctx, req.DeviceName)
	return
}

func (c cDevice) ActivationSuccess(ctx context.Context, req *v1.ActivationSuccessReq) (res *v1.EmptyFieldRes, err error) {
	successInput := model.ActivationSuccessInput{}
	if err = gconv.Scan(req, &successInput); err != nil {
		panic(err)
	}
	service.DevDevice().ActivationSuccess(ctx, &successInput)
	return
}
