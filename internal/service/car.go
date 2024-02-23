package service

import (
	"context"
	"sviwo/internal/model"
)

type (
	ICar interface {
		GetCarList(ctx context.Context, in model.QueryCarInput) (out *model.QueryCarOutput)
		DelCar(ctx context.Context, carId uint64) error
		ControlCarLamp(ctx context.Context, carId uint64) error
		ControlCarHorn(ctx context.Context, carId uint64) error
		ControlCarSwitchMode(ctx context.Context, in model.ControlCarSwitchModeInput) error
		EnabledMobileKey(ctx context.Context, carId uint64) error
		EnabledSpeedLimit(ctx context.Context, carId uint64) error
		BindingCar(ctx context.Context, in model.BindingCarInput) error
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
