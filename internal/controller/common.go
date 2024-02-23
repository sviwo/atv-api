package controller

import (
	"context"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"sviwo/api/v1"
	"sviwo/internal/service"
	"sviwo/utility/file"
)

var Common = cCommon{}

type cCommon struct{}

/*
GetVftCode 获取邮箱验证码
*/
func (cCommon) GetVftCode(ctx context.Context, req *v1.VftCodeReq) (res *v1.VftCodeRes, err error) {
	return res, service.Common().GetVftCode(ctx, req.Email)
}

/*
ImgUpload 图片上传
*/
func (cCommon) ImgUpload(ctx context.Context, req *v1.ImgUploadReq) (res *v1.ImgUploadRes, err error) {
	uri, err := file.UploadFile(req.File)
	res = new(v1.ImgUploadRes)
	res.Uri = uri
	return res, err
}
