// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-04-01 16:15:05
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	UserId      uint64      `json:"userId"      dc:""`
	Username    string      `json:"username"    dc:""`
	Password    string      `json:"password"    dc:""`
	PwdSalt     string      `json:"pwdSalt"     dc:"密码盐值"`
	PwdEncryNum int         `json:"pwdEncryNum" dc:"密码加密次数"`
	FirstName   string      `json:"firstName"   dc:""`
	LastName    string      `json:"lastName"    dc:""`
	Enable      bool        `json:"enable"      dc:"账号是否可用：true=正常，false=停用"`
	HeadImg     string      `json:"headImg"     dc:""`
	MobilePhone string      `json:"mobilePhone" dc:"手机号"`
	UserAddress string      `json:"userAddress" dc:"用户地址"`
	CreateTime  *gtime.Time `json:"createTime"  dc:""`
	UpdateTime  *gtime.Time `json:"updateTime"  dc:""`
	IsDelete    bool        `json:"isDelete"    dc:"是否删除：true=已删除，false=正常"`
}
