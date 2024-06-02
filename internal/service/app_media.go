package service

import (
	"context"
	"sviwo/internal/model"
)

type (
	IAppMedia interface {
		GetAppMediaTree(ctx context.Context, pageType int) (out []*model.AppMediaTreeOutput)
		GetAppMediaDetail(ctx context.Context, Id int64) (content string)
	}
)

var (
	localAppMedia IAppMedia
)

func AppMedia() IAppMedia {
	if localAppMedia == nil {
		panic("implement not found for interface IAppMedia, forgot register?")
	}
	return localAppMedia
}

func RegisterAppMedia(i IAppMedia) {
	localAppMedia = i
}
