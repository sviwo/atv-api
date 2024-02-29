package model

import "github.com/gogf/gf/v2/os/gtime"

type UserAuthInput struct {
	AuthFirstName       string `json:"authFirstName"       description:""`
	AuthLastName        string `json:"authLastName"        description:""`
	CertificateFrontImg string `json:"certificateFrontImg" description:"证件正面照片"`
	CertificateBackImg  string `json:"certificateBackImg"  description:"证件背面照片"`
}

type UserAuthOutput struct {
	AuthFirstName       string      `json:"authFirstName"       description:""`
	AuthLastName        string      `json:"authLastName"        description:""`
	CertificateFrontImg string      `json:"certificateFrontImg" description:"证件正面照片"`
	CertificateBackImg  string      `json:"certificateBackImg"  description:"证件背面照片"`
	AuthStatus          uint        `json:"authStatus"          description:"认证状态：0=未认证，1=认证中，2=认证成功，3=认证失败"`
	AuthFailReason      string      `json:"authFailReason"      description:"认证失败原因"`
	AuthTime            *gtime.Time `json:"authTime"            description:"认证时间"`
	VerifyTime          *gtime.Time `json:"verifyTime"          description:"审核时间"`
}
