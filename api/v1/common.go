package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

/*
VftCodeReq 获取验证码
*/
type VftCodeReq struct {
	g.Meta       `path:"/getVftCode" method:"get" tags:"公共接口" summary:"获取验证码"`
	EmailVftCode string `json:"emailVftCode" v:"required"   dc:"请输入邮箱"`
}

type VftCodeRes struct{}

/*
ImgUploadReq 上传图片
*/
type ImgUploadReq struct {
	g.Meta `path:"/img/upload" method:"post" mime:"multipart/form-data" tags:"公共接口" summary:"上传图片"`
	File   *ghttp.UploadFile `json:"file" type:"file" v:"required"   dc:"请选择上传文件"`
}

type ImgUploadRes struct {
	Uri string `json:"uri"    dc:"Uri"`
}
