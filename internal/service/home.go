package service

import (
	"context"
	"sviwo/internal/model"
)

type (
	IHome interface {
		GetHomeData(ctx context.Context) (out *model.HomeDataOutput)
	}
)

var (
	localHome IHome
)

func Home() IHome {
	if localHome == nil {
		panic("implement not found for interface IHome, forgot register?")
	}
	return localHome
}

func RegisterHome(i IHome) {
	localHome = i
}
