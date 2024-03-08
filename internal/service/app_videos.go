package service

import (
	"context"
	"sviwo/internal/model"
)

type (
	IAppVideos interface {
		GetAppVideosTree(ctx context.Context) (out []*model.AppVideosListOutput)
	}
)

var (
	localAppVideos IAppVideos
)

func AppVideos() IAppVideos {
	if localAppVideos == nil {
		panic("implement not found for interface IAppVideos, forgot register?")
	}
	return localAppVideos
}

func RegisterAppVideos(i IAppVideos) {
	localAppVideos = i
}
