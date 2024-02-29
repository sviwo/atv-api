package controller

import (
	"context"
	"fmt"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
	"sviwo/api/v1"
	"sviwo/internal/boot"
	"sviwo/internal/consts"
	"sviwo/internal/service"
	ecc "sviwo/utility/encrypt"
	"sviwo/utility/file"
)

var Common = cCommon{}

type cCommon struct{}

/*
GetVftCode 获取邮箱验证码
*/
func (cCommon) GetVftCode(ctx context.Context, req *v1.VftCodeReq) (res *v1.EmptyFieldRes, err error) {
	return res, service.Common().GetVftCode(ctx, req.Email)
}

/*
ImgUpload 图片上传
*/
func (cCommon) ImgUpload(ctx context.Context, req *v1.ImgUploadReq) (res *v1.ImgUploadRes, err error) {
	uri, err := file.UploadFile(req.File)
	res = new(v1.ImgUploadRes)
	res = &v1.ImgUploadRes{Uri: uri}
	return
}

/*
GetEccPublicKey 获取ecc公钥
*/
func (cCommon) GetEccPublicKey(ctx context.Context, req *v1.EccPublicKeyReq) (res *v1.EccPublicKeyRes, err error) {
	key, err := ecc.GenerateEccKeyHex()
	if err != nil {
		panic(err)
	}
	publicCode := boot.GID.Generate().String()
	if err = g.Redis().SetEX(
		ctx, fmt.Sprintf(consts.RedisEccPrivateKey, publicCode), key.PrivateKey, 120,
	); err != nil {
		panic(err)
	}
	publicKey := key.PublicKey
	res = &v1.EccPublicKeyRes{PublicKey: publicKey, PublicCode: publicCode}
	text, _ := ecc.EccEncryptToHex(gconv.Bytes("{\"email\":\"821317143@qq.com\"}"), publicKey)
	glog.Info(ctx, "text=====", text)
	return
}
