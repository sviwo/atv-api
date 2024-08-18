// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-08-18 15:21:45
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserThirdAccount is the golang structure of table sw_user_third_account for DAO operations like Where/Data.
type UserThirdAccount struct {
	g.Meta       `orm:"table:sw_user_third_account, do:true"`
	Id           interface{} //
	UserId       interface{} // 本地用户id
	ThirdUserId  interface{} // 第三方用户id
	ProviderType interface{} // 第三方平台类型：0=apple，1=facebook，2=google
	CreateTime   *gtime.Time //
	UpdateTime   *gtime.Time //
}
