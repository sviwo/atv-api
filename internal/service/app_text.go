package service

import (
	"context"
	"sviwo/internal/model"
)

type (
	IAppText interface {
		GetAppTextTree(ctx context.Context) (out []*model.AppTextListOutput)
		GetAppTextDetail(ctx context.Context, textId int64) (out *model.AppTextDetailOutput)
	}
)

var (
	localAppText IAppText
)

func AppText() IAppText {
	if localAppText == nil {
		panic("implement not found for interface IAppText, forgot register?")
	}
	return localAppText
}

func RegisterAppText(i IAppText) {
	localAppText = i
}
