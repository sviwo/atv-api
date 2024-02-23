package car

import (
	"context"
	"sviwo/internal/model"
	"sviwo/internal/service"
)

func init() {
	service.RegisterCar(New())
}

func New() *sCar {
	return &sCar{}
}

type sCar struct{}

func (s sCar) GetCarList(ctx context.Context, in model.QueryCarInput) (out *model.QueryCarOutput) {
	panic("implement me")
}

func (s sCar) DelCar(ctx context.Context, carId uint64) error {
	panic("implement me")
}

func (s sCar) ControlCarLamp(ctx context.Context, carId uint64) error {
	panic("implement me")
}

func (s sCar) ControlCarHorn(ctx context.Context, carId uint64) error {
	panic("implement me")
}

func (s sCar) ControlCarSwitchMode(ctx context.Context, in model.ControlCarSwitchModeInput) error {
	panic("implement me")
}

func (s sCar) EnabledMobileKey(ctx context.Context, carId uint64) error {
	panic("implement me")
}

func (s sCar) EnabledSpeedLimit(ctx context.Context, carId uint64) error {
	panic("implement me")
}

func (s sCar) BindingCar(ctx context.Context, in model.BindingCarInput) error {
	panic("implement me")
}
