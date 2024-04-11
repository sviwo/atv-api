package service

import (
	"context"
	"sviwo/internal/model"
)

type (
	ICar interface {
		GetCarList(ctx context.Context) (out []*model.QueryCarOutput)

		GetCarDetail(ctx context.Context, deviceId *int64) (out *model.UserDeviceOutput)

		SwitchCar(ctx context.Context, deviceId int64)

		DelCar(ctx context.Context, deviceId int64)
		/*
			CtlLamp 控车
		*/
		CtlCar(ctx context.Context, instructions int)
		/*
			CtlSwitchDT 切换驾驶模式
		*/
		CtlSwitchDT(ctx context.Context, in model.CtlSwitchDTInput)
		/*
			CtlSwitchERT 切换动能回收模式
		*/
		CtlSwitchERT(ctx context.Context, in model.CtlSwitchERTInput)
		/*
			EnabledMobileKey 开启/关闭蓝牙钥匙
		*/
		EnabledMobileKey(ctx context.Context)
		/*
			EnabledSpeedLimit 开启/关闭速度限制
		*/
		EnabledSpeedLimit(ctx context.Context)
		BindingCar(ctx context.Context, userId int64, deviceName string)
	}
)

var (
	localCar ICar
)

func Car() ICar {
	if localCar == nil {
		panic("implement not found for interface ICar, forgot register?")
	}
	return localCar
}

func RegisterCar(i ICar) {
	localCar = i
}
