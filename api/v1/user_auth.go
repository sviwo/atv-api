package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
)

type GetUserAuthReq struct {
	g.Meta `path:"/user/get/auth" method:"get" tags:"用户实名相关" sm:"获取用户实名信息"`
}

type GetUserAuthRes struct {
	AuthFirstName       string      `json:"authFirstName"       dc:""`
	AuthLastName        string      `json:"authLastName"        dc:""`
	CertificateFrontImg string      `json:"certificateFrontImg" dc:"证件正面照片"`
	CertificateBackImg  string      `json:"certificateBackImg"  dc:"证件背面照片"`
	AuthStatus          uint        `json:"authStatus"          dc:"认证状态：0=未认证，1=认证中，2=认证成功，3=认证失败"`
	AuthFailReason      string      `json:"authFailReason"      dc:"认证失败原因"`
	AuthTime            *gtime.Time `json:"authTime"            dc:"认证时间"`
	VerifyTime          *gtime.Time `json:"verifyTime"          dc:"审核时间"`
}

type SubmitUserAuthReq struct {
	g.Meta              `path:"/user/submit/auth" method:"post" tags:"用户实名相关" sm:"提交用户实名信息"`
	AuthFirstName       string            `json:"authFirstName"       dc:""                      v:"required"`
	AuthLastName        string            `json:"authLastName"        dc:""                      v:"required"`
	CertificateFrontImg *ghttp.UploadFile `json:"certificateFrontImg" type:"file"   dc:"证件正面"  v:"required"`
	CertificateBackImg  *ghttp.UploadFile `json:"certificateBackImg"  type:"file"   dc:"证件背面"  v:"required"`
}
