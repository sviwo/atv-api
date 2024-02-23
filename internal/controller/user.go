package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"sviwo/api/v1"
	"sviwo/internal/model"
	"sviwo/internal/service"
)

var User = cUser{}

type cUser struct{}

/*
Register 用户注册
*/
func (c cUser) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error) {
	data := model.RegisterInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		panic(err)
	}
	return res, service.User().Register(ctx, data)
}

/*
Info 获取用户信息
*/
func (c cUser) Info(ctx context.Context, req *v1.UserInfoReq) (res *v1.UserInfoRes, err error) {
	iData := model.UserInfoInput{}
	err = gconv.Struct(req, &iData)
	if err != nil {
		panic(err)
	}
	err = gconv.Struct(service.User().Info(ctx, iData), &res)
	if err != nil {
		panic(err)
	}
	return res, err
}

/*
UpdatePassword 修改密码
*/
func (c cUser) UpdatePassword(ctx context.Context, req *v1.UpdatePasswordReq) (res *v1.UpdatePasswordRes, err error) {
	data := model.UpdatePasswordInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		panic(err)
	}
	return res, service.User().UpdatePassword(ctx, data)
}

/*
EditInfo 编辑用户信息
*/
func (c cUser) EditInfo(ctx context.Context, req *v1.EditInfoReq) (res *v1.EditInfoRes, err error) {
	data := model.EditInfoInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		panic(err)
	}
	return res, service.User().EditInfo(ctx, data)
}
