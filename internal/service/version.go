package service

import (
	"context"
	"sviwo/internal/model"
)

type (
	IVersion interface {
		GetVersionList(ctx context.Context, in *model.VersionInput) (totalCount int, out []*model.VersionOutput)
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
