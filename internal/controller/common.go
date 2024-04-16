package controller

import (
	"context"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"sviwo/api/v1"
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
	uri, err := file.UploadFile(req.File)
	if err != nil {
		panic(err)
	}
	res = &v1.ImgUploadRes{Uri: uri}
	return
}

func (cCommon) GetEccPublicKey(ctx context.Context, req *v1.EccPublicKeyReq) (res *v1.EccPublicKeyRes, err error) {
	key, code := service.Common().GetEccPublicKey(ctx)
	res = &v1.EccPublicKeyRes{PublicKey: key, PublicCode: code}
	return
}
