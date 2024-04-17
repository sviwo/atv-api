package model

import "github.com/gogf/gf/v2/os/gtime"

type UserAuthInput struct {
	AuthFirstName       string  `json:"authFirstName"       dc:""`
	AuthLastName        string  `json:"authLastName"        dc:""`
	CertificateFrontImg *string `json:"certificateFrontImg" dc:"证件正面照片"`
	CertificateBackImg  *string `json:"certificateBackImg"  dc:"证件背面照片"`
}

type UserAuthOutput struct {
	AuthFirstName       string      `json:"authFirstName"       dc:""`
	AuthLastName        string      `json:"authLastName"        dc:""`
	CertificateFrontImg string      `json:"certificateFrontImg" dc:"证件正面照片"`
	CertificateBackImg  string      `json:"certificateBackImg"  dc:"证件背面照片"`
	AuthStatus          uint        `json:"authStatus"          dc:"认证状态：0=未认证，1=认证中，2=认证成功，3=认证失败"`
	AuthFailReason      string      `json:"authFailReason"      dc:"认证失败原因"`
	AuthTime            *gtime.Time `json:"authTime"            dc:"认证时间"`
	VerifyTime          *gtime.Time `json:"verifyTime"          dc:"审核时间"`
}
