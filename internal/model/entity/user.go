// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-03-15 14:53:40
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	UserId      uint64      `json:"userId"      description:""`
	Username    string      `json:"username"    description:""`
	Password    string      `json:"password"    description:""`
	PwdSalt     string      `json:"pwdSalt"     description:"密码盐值"`
	PwdEncryNum int         `json:"pwdEncryNum" description:"密码加密次数"`
	FirstName   string      `json:"firstName"   description:""`
	LastName    string      `json:"lastName"    description:""`
	Enable      bool        `json:"enable"      description:"账号是否可用：true=正常，false=停用"`
	HeadImg     string      `json:"headImg"     description:""`
	MobilePhone string      `json:"mobilePhone" description:"手机号"`
	UserAddress string      `json:"userAddress" description:"用户地址"`
	CreateTime  *gtime.Time `json:"createTime"  description:""`
	UpdateTime  *gtime.Time `json:"updateTime"  description:""`
	IsDelete    bool        `json:"isDelete"    description:"是否删除：true=已删除，false=正常"`
}
