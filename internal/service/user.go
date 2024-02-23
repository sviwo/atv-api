package service

import (
	"context"
	"sviwo/internal/model"
)

type (
	IUser interface {
		Login(ctx context.Context, in model.LoginInput) (err error, userId *uint64)
		Register(ctx context.Context, in model.RegisterInput) error
		Info(ctx context.Context, in model.UserInfoInput) (out *model.UserInfoOutput)
		UpdatePassword(ctx context.Context, in model.UpdatePasswordInput) error
		EditInfo(ctx context.Context, in model.EditInfoInput) error
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
