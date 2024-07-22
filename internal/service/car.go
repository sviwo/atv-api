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

		EditCarNickname(ctx context.Context, nickname string)

		RemoveCar(ctx context.Context, userDeviceId, deviceId *int64)
		/*
			GetCarKey 获取车辆钥匙
		*/
		GetCarKey(ctx context.Context) (carKey string)
		/*
			InviteBindCar 使用车辆钥匙绑定车辆
		*/
		InviteBindCar(ctx context.Context, carKey string)
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
		/*
			GetSimDataTraffic 获取sim卡数据流量
		*/
		GetSimDataTraffic(ctx context.Context) (out *model.SimDataTrafficOutput)
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
