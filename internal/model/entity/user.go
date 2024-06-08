// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-06-08 20:37:43
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	UserId      uint64      `json:"userId"      orm:"user_id"       description:""`
	Username    string      `json:"username"    orm:"username"      description:""`
	Password    string      `json:"password"    orm:"password"      description:""`
	PwdSalt     string      `json:"pwdSalt"     orm:"pwd_salt"      description:"密码盐值"`
	PwdEncryNum int         `json:"pwdEncryNum" orm:"pwd_encry_num" description:"密码加密次数"`
	FirstName   string      `json:"firstName"   orm:"first_name"    description:""`
	LastName    string      `json:"lastName"    orm:"last_name"     description:""`
	Enable      bool        `json:"enable"      orm:"enable"        description:"账号是否可用：true=正常，false=停用"`
	HeadImg     string      `json:"headImg"     orm:"head_img"      description:""`
	MobilePhone string      `json:"mobilePhone" orm:"mobile_phone"  description:"手机号"`
	UserAddress string      `json:"userAddress" orm:"user_address"  description:"用户地址"`
	CreateTime  *gtime.Time `json:"createTime"  orm:"create_time"   description:""`
	UpdateTime  *gtime.Time `json:"updateTime"  orm:"update_time"   description:""`
	IsDelete    bool        `json:"isDelete"    orm:"is_delete"     description:"是否删除：true=已删除，false=正常"`
}
