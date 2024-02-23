package service

import (
	"context"
)

type (
	ICommon interface {
		GetVftCode(ctx context.Context, email string) error
	}
)

var (
	localCommon ICommon
)

func Common() ICommon {
	if localCommon == nil {
		panic("implement not found for interface ICommon, forgot register?")
	}
	return localCommon
}

func RegisterCommon(i ICommon) {
	localCommon = i
}
