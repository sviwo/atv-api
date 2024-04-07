package controller

import (
	"context"
	v1 "sviwo/api/v1"
	"sviwo/internal/service"
)

var DeviceProperty = cDeviceProperty{}

type cDeviceProperty struct{}

// Set 设备属性设置
func (c *cDeviceProperty) Set(ctx context.Context, req *v1.DevicePropertyReq) (res *v1.DevicePropertyRes, err error) {
	out, err := service.DevDeviceProperty().Set(ctx, req.DevicePropertyInput)
	res = &v1.DevicePropertyRes{
		DevicePropertyOutput: out,
	}
	return
}
