package controller

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"sviwo/api/v1"
	"sviwo/internal/model"
	"sviwo/internal/service"
	"sviwo/pkg/utility/file"
)

var User = cUser{}

type cUser struct{}

/*
Register 用户注册
*/
func (c cUser) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.EmptyFieldRes, err error) {
	data := model.RegisterInput{}
	if err = gconv.Struct(req, &data); err != nil {
		panic(err)
	}
	service.User().Register(ctx, data)
	return
}

/*
Info 获取用户信息
*/
func (c cUser) Info(ctx context.Context, req *v1.UserInfoReq) (res *v1.UserInfoRes, err error) {
	if err = gconv.Struct(service.User().Info(ctx), &res); err != nil {
		panic(err)
	}
	return
}

/*
UpdatePassword 修改密码
*/
func (c cUser) UpdatePassword(ctx context.Context, req *v1.UpdatePasswordReq) (res *v1.EmptyFieldRes, err error) {
	data := model.UpdatePasswordInput{}
	if err = gconv.Struct(req, &data); err != nil {
		panic(err)
	}
	service.User().UpdatePassword(ctx, data)
	return
}

/*
EditInfo 编辑用户信息
*/
func (c cUser) EditInfo(ctx context.Context, req *v1.EditInfoReq) (res *v1.EmptyFieldRes, err error) {
	data := model.EditInfoInput{UpdateTime: gtime.Now()}
	if err = gconv.Struct(req, &data); err != nil {
		panic(err)
	}
	if req.HeadImg != nil {
		uri, err2 := file.UploadFile(req.HeadImg)
		if err2 != nil {
			panic(err2)
		}
		data.HeadImg = &uri
	}
	service.User().EditInfo(ctx, data)
	return
}
