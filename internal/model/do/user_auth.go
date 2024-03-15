// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-03-15 14:53:40
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserAuth is the golang structure of table sw_user_auth for DAO operations like Where/Data.
type UserAuth struct {
	g.Meta              `orm:"table:sw_user_auth, do:true"`
	AuthId              interface{} //
	UserId              interface{} //
	AuthFirstName       interface{} //
	AuthLastName        interface{} //
	CertificateFrontImg interface{} // 证件正面照片
	CertificateBackImg  interface{} // 证件背面照片
	AuthStatus          interface{} // 认证状态：0=未认证，1=认证中，2=认证成功，3=认证失败
	AuthFailReason      interface{} // 认证失败原因
	AuthTime            *gtime.Time // 认证时间
	VerifyTime          *gtime.Time // 审核时间
	CreateTime          *gtime.Time //
	UpdateTime          *gtime.Time //
}
