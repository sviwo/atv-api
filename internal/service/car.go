package service

import (
	"context"
	"sviwo/internal/model"
)

type (
	ICar interface {
		GetCarList(ctx context.Context) (out []*model.QueryCarOutput)
		DelCar(ctx context.Context, carId int64)
		/*
			ControlCarLamp 控制车灯
		*/
		ControlCarLamp(ctx context.Context, carId int64)
		/*
			ControlCarHorn 控制喇叭
		*/
		ControlCarHorn(ctx context.Context, carId int64)
		/*
			ControlCarSwitchMode 切换模式
		*/
		ControlCarSwitchMode(ctx context.Context, in model.ControlCarSwitchModeInput)
		/*
			EnabledMobileKey 开启/关闭蓝牙钥匙
		*/
		EnabledMobileKey(ctx context.Context, carId int64)
		/*
			EnabledSpeedLimit 开启/关闭速度限制
		*/
		EnabledSpeedLimit(ctx context.Context, carId int64)
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
