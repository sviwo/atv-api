// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-05-10 15:47:49
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserAuth is the golang structure for table user_auth.
type UserAuth struct {
	AuthId              int64       `json:"authId"              orm:"auth_id"               description:""`
	UserId              int64       `json:"userId"              orm:"user_id"               description:""`
	AuthFirstName       string      `json:"authFirstName"       orm:"auth_first_name"       description:""`
	AuthLastName        string      `json:"authLastName"        orm:"auth_last_name"        description:""`
	CertificateFrontImg string      `json:"certificateFrontImg" orm:"certificate_front_img" description:"证件正面照片"`
	CertificateBackImg  string      `json:"certificateBackImg"  orm:"certificate_back_img"  description:"证件背面照片"`
	AuthStatus          uint        `json:"authStatus"          orm:"auth_status"           description:"认证状态：0=未认证，1=认证中，2=认证成功，3=认证失败"`
	AuthFailReason      string      `json:"authFailReason"      orm:"auth_fail_reason"      description:"认证失败原因"`
	AuthTime            *gtime.Time `json:"authTime"            orm:"auth_time"             description:"认证时间"`
	VerifyTime          *gtime.Time `json:"verifyTime"          orm:"verify_time"           description:"审核时间"`
	CreateTime          *gtime.Time `json:"createTime"          orm:"create_time"           description:""`
	UpdateTime          *gtime.Time `json:"updateTime"          orm:"update_time"           description:""`
}
