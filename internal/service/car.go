package service

import (
	"context"
	"sviwo/internal/model"
)

type (
	ICar interface {
		GetCarList(ctx context.Context) (out []*model.QueryCarOutput)
		DelCar(ctx context.Context, deviceId int64)
		/*
			CtlLamp 控制车灯
		*/
		CtlLamp(ctx context.Context, deviceId int64)
		/*
			CtlHorn 控制喇叭
		*/
		CtlHorn(ctx context.Context, deviceId int64)
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
		EnabledMobileKey(ctx context.Context, deviceId int64)
		/*
			EnabledSpeedLimit 开启/关闭速度限制
		*/
		EnabledSpeedLimit(ctx context.Context, deviceId int64)
		BindingCar(ctx context.Context, carFrameCode string)
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
