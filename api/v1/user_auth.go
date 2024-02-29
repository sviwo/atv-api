package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type GetUserAuthReq struct {
	g.Meta `path:"get/user/auth" method:"get" tags:"用户实名相关" summary:"获取用户实名信息"`
}

type GetUserAuthRes struct {
	AuthFirstName       string      `json:"authFirstName"       description:""`
	AuthLastName        string      `json:"authLastName"        description:""`
	CertificateFrontImg string      `json:"certificateFrontImg" description:"证件正面照片"`
	CertificateBackImg  string      `json:"certificateBackImg"  description:"证件背面照片"`
	AuthStatus          uint        `json:"authStatus"          description:"认证状态：0=未认证，1=认证中，2=认证成功，3=认证失败"`
	AuthFailReason      string      `json:"authFailReason"      description:"认证失败原因"`
	AuthTime            *gtime.Time `json:"authTime"            description:"认证时间"`
	VerifyTime          *gtime.Time `json:"verifyTime"          description:"审核时间"`
}

type SubmitUserAuthReq struct {
	g.Meta              `path:"submit/user/auth" method:"post" tags:"用户实名相关" summary:"提交用户实名信息"`
	AuthFirstName       string `json:"authFirstName"       description:""`
	AuthLastName        string `json:"authLastName"        description:""`
	CertificateFrontImg string `json:"certificateFrontImg" description:"证件正面照片"`
	CertificateBackImg  string `json:"certificateBackImg"  description:"证件背面照片"`
}
