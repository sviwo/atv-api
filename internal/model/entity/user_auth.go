// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-02-23 10:50:21
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserAuth is the golang structure for table user_auth.
type UserAuth struct {
	AuthId              uint64      `json:"authId"              description:""`
	UserId              int64       `json:"userId"              description:""`
	AuthFirstName       string      `json:"authFirstName"       description:""`
	AuthLastName        string      `json:"authLastName"        description:""`
	CertificateFrontImg string      `json:"certificateFrontImg" description:"证件正面照片"`
	CertificateBackImg  string      `json:"certificateBackImg"  description:"证件背面照片"`
	AuthStatus          uint        `json:"authStatus"          description:"认证状态：0=未认证，1=认证中，2=认证成功，3=认证失败"`
	AuthFailReason      string      `json:"authFailReason"      description:"认证失败原因"`
	AuthTime            *gtime.Time `json:"authTime"            description:"认证时间"`
	VerifyTime          *gtime.Time `json:"verifyTime"          description:"审核时间"`
	CreateTime          *gtime.Time `json:"createTime"          description:""`
	UpdateTime          *gtime.Time `json:"updateTime"          description:""`
}
