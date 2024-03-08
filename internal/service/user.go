package service

import (
	"context"
	"sviwo/internal/model"
)

type (
	IUser interface {
		Login(ctx context.Context, in model.LoginInput) *uint64
		Register(ctx context.Context, in model.RegisterInput)
		Info(ctx context.Context) (out *model.UserInfoOutput)
		UpdatePassword(ctx context.Context, in model.UpdatePasswordInput)
		EditInfo(ctx context.Context, in model.EditInfoInput)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
