package service

import (
	"context"
	"sviwo/internal/model"
)

type (
	IUserAuth interface {
		GetUserAuthInfo(ctx context.Context) (out *model.UserAuthOutput)
		SubmitUserAuth(ctx context.Context, in model.UserAuthInput)
	}
)

var (
	localUserAuth IUserAuth
)

func UserAuth() IUserAuth {
	if localUserAuth == nil {
		panic("implement not found for interface IUserAuth, forgot register?")
	}
	return localUserAuth
}

func RegisterUserAuth(i IUserAuth) {
	localUserAuth = i
}
