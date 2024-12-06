package controller

import (
	"context"
	"fmt"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"sviwo/api/v1"
	"sviwo/internal/consts"
	"sviwo/internal/service"
	"sviwo/pkg/utility/file"
)

var Common = cCommon{}

type cCommon struct{}

func (cCommon) GetVftCode(ctx context.Context, req *v1.VftCodeReq) (res *v1.EmptyFieldRes, err error) {
	service.Common().GetVftCode(ctx, req.Email)
	return
}

func (cCommon) ImgUpload(ctx context.Context, req *v1.ImgUploadReq) (res *v1.ImgUploadRes, err error) {
	res = &v1.ImgUploadRes{Uri: file.UploadFile(req.File)}
	return
}

func (cCommon) GetEccPublicKey(ctx context.Context, req *v1.EccPublicKeyReq) (res *v1.EccPublicKeyRes, err error) {
	key, code := service.Common().GetEccPublicKey(ctx)
	res = &v1.EccPublicKeyRes{PublicKey: key, PublicCode: code}
	return
}

func (cCommon) SendEmail(ctx context.Context, req *v1.SendEmailReq) (res *v1.EmptyFieldRes, err error) {
	ip := service.BizCtx().Get(ctx).Data.Get(consts.ClientIp)
	go service.Common().SendEmail(ctx, fmt.Sprint(ip), req.Email, req.Content)
	return
}
