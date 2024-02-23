// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-01-16 21:22:22
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table sw_user for DAO operations like Where/Data.
type User struct {
	g.Meta            `orm:"table:sw_user, do:true"`
	UserId            interface{} //
	Username          interface{} //
	Password          interface{} //
	PwdSalt           interface{} // 密码盐值
	PwdEncryNum       interface{} // 密码加密次数
	FirstName         interface{} //
	LastName          interface{} //
	Nickname          interface{} // 昵称
	Enable            interface{} // 账号是否可用：true=正常，false=停用
	HeadImg           interface{} //
	CertificateCode   interface{} // 身份证编号
	VipLevel          interface{} // vip级别
	IpRegion          interface{} // Ip属地
	MobilePhone       interface{} // 手机号
	PersonalSignature interface{} // 个性签名
	Birthday          *gtime.Time // 生日
	Gender            interface{} // 性别：0=未知，1=男，2=女
	CreateTime        *gtime.Time //
	UpdateTime        *gtime.Time //
	IsDelete          interface{} // 是否删除：true=已删除，false=正常
}
