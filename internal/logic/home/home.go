package home

import (
	"context"
	"sviwo/internal/model"
	"sviwo/internal/service"
)

func init() {
	service.RegisterHome(New())
}

func New() *sHome {
	return &sHome{}
}

type sHome struct{}

func (s sHome) GetHomeData(ctx context.Context, in model.QueryHomeDataInput) (out *model.HomeDataOutput) {
	panic("implement me")
}
