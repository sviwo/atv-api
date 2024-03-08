package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

/*
EmptyFieldRes 无字段返回结构体，所有无字段返回的Res均使用此结构体
*/
type EmptyFieldRes struct{}

type VftCodeReq struct {
	g.Meta `path:"/common/getVftCode" method:"get" tags:"公共接口" sm:"获取验证码"`
	Email  string `json:"email" v:"required|email"   dc:"请输入邮箱"`
}

type ImgUploadReq struct {
	g.Meta `path:"/common/img/upload" method:"post" mime:"multipart/form-data" tags:"公共接口" sm:"上传图片"`
	File   *ghttp.UploadFile `json:"file" type:"file" v:"required"   dc:"请选择上传文件"`
}

type ImgUploadRes struct {
	Uri string `json:"uri"    dc:"Uri"`
}

type EccPublicKeyReq struct {
	g.Meta `path:"/common/getEccPublicKey" method:"get" tags:"公共接口" sm:"获取eccK公钥"`
}

type EccPublicKeyRes struct {
	PublicKey  string `json:"publicKey"    dc:""`
	PublicCode string `json:"publicCode"    dc:""`
}
