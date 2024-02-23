package service

import (
	"context"
	"sviwo/internal/model"
)

type (
	IVersion interface {
		GetNewVersion(ctx context.Context) (out []*model.VersionOutput)
	}
)

var (
	localVersion IVersion
)

func Version() IVersion {
	if localVersion == nil {
		panic("implement not found for interface IVersion, forgot register?")
	}
	return localVersion
}

func RegisterVersion(i IVersion) {
	localVersion = i
}
