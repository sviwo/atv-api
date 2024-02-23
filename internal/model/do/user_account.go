// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-01-16 21:22:22
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserAccount is the golang structure of table sw_user_account for DAO operations like Where/Data.
type UserAccount struct {
	g.Meta     `orm:"table:sw_user_account, do:true"`
	AccountId  interface{} //
	UserId     interface{} //
	Balance    interface{} // 账户余额
	MyIntegral interface{} // 我的积分
	CreateTime *gtime.Time //
	UpdateTime *gtime.Time //
	IsDelete   interface{} // 是否删除：true=已删除，false=正常
}
