package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"sviwo/api/v1"
	"sviwo/internal/model"
	"sviwo/internal/service"
)

var UserAuth = cUserAuth{}

type cUserAuth struct{}

/*
GetUserAuthInfo 获取用户实名信息
*/
func (c cUserAuth) GetUserAuthInfo(ctx context.Context, req *v1.GetUserAuthReq) (res *v1.GetUserAuthRes, err error) {
	if err = gconv.Struct(service.UserAuth().GetUserAuthInfo(ctx), &res); err != nil {
		panic(err)
	}
	return
}

/*
SubmitUserAuth 提交用户实名信息
*/
func (c cUserAuth) SubmitUserAuth(ctx context.Context, req *v1.SubmitUserAuthReq) (res *v1.EmptyFieldRes, err error) {
	data := model.UserAuthInput{}
	if err = gconv.Struct(req, &data); err != nil {
		panic(err)
	}
	service.UserAuth().SubmitUserAuth(ctx, data)
	return
}
