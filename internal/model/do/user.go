// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-03-15 14:53:40
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table sw_user for DAO operations like Where/Data.
type User struct {
	g.Meta      `orm:"table:sw_user, do:true"`
	UserId      interface{} //
	Username    interface{} //
	Password    interface{} //
	PwdSalt     interface{} // 密码盐值
	PwdEncryNum interface{} // 密码加密次数
	FirstName   interface{} //
	LastName    interface{} //
	Enable      interface{} // 账号是否可用：true=正常，false=停用
	HeadImg     interface{} //
	MobilePhone interface{} // 手机号
	UserAddress interface{} // 用户地址
	CreateTime  *gtime.Time //
	UpdateTime  *gtime.Time //
	IsDelete    interface{} // 是否删除：true=已删除，false=正常
}
