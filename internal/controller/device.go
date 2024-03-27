package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"sviwo/api/v1"
	"sviwo/internal/service"
)

var Device = cDevice{}

type cDevice struct{}

func (c cDevice) GetDeviceSecret(ctx context.Context, req *v1.DeviceSecretReq) (res *v1.DeviceSecretRes, err error) {
	if err = gconv.Struct(service.DevDevice().GetDeviceSecret(ctx, req.DeviceCode), &res); err != nil {
		panic(err)
	}
	return
}
