// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-04-01 16:15:05
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserAuth is the golang structure for table user_auth.
type UserAuth struct {
	AuthId              int64       `json:"authId"              dc:""`
	UserId              int64       `json:"userId"              dc:""`
	AuthFirstName       string      `json:"authFirstName"       dc:""`
	AuthLastName        string      `json:"authLastName"        dc:""`
	CertificateFrontImg string      `json:"certificateFrontImg" dc:"证件正面照片"`
	CertificateBackImg  string      `json:"certificateBackImg"  dc:"证件背面照片"`
	AuthStatus          uint        `json:"authStatus"          dc:"认证状态：0=未认证，1=认证中，2=认证成功，3=认证失败"`
	AuthFailReason      string      `json:"authFailReason"      dc:"认证失败原因"`
	AuthTime            *gtime.Time `json:"authTime"            dc:"认证时间"`
	VerifyTime          *gtime.Time `json:"verifyTime"          dc:"审核时间"`
	CreateTime          *gtime.Time `json:"createTime"          dc:""`
	UpdateTime          *gtime.Time `json:"updateTime"          dc:""`
}
