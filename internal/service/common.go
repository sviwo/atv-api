package service

import (
	"context"
)

type (
	ICommon interface {
		/*
			GetVftCode获取验证码
		*/
		GetVftCode(ctx context.Context, email string)

		/*
			发送邮箱信息
		*/
		SendEmail(ctx context.Context, ip, email, content string)

		GetEccPublicKey(ctx context.Context) (publicKey, publicCode string)
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
