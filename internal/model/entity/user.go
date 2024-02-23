// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-01-16 21:22:22
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	UserId            uint64      `json:"userId"            description:""`
	Username          string      `json:"username"          description:""`
	Password          string      `json:"password"          description:""`
	PwdSalt           string      `json:"pwdSalt"           description:"密码盐值"`
	PwdEncryNum       int         `json:"pwdEncryNum"       description:"密码加密次数"`
	FirstName         string      `json:"firstName"         description:""`
	LastName          string      `json:"lastName"          description:""`
	Nickname          string      `json:"nickname"          description:"昵称"`
	Enable            bool        `json:"enable"            description:"账号是否可用：true=正常，false=停用"`
	HeadImg           string      `json:"headImg"           description:""`
	CertificateCode   string      `json:"certificateCode"   description:"身份证编号"`
	VipLevel          int         `json:"vipLevel"          description:"vip级别"`
	IpRegion          string      `json:"ipRegion"          description:"Ip属地"`
	MobilePhone       string      `json:"mobilePhone"       description:"手机号"`
	PersonalSignature string      `json:"personalSignature" description:"个性签名"`
	Birthday          *gtime.Time `json:"birthday"          description:"生日"`
	Gender            int         `json:"gender"            description:"性别：0=未知，1=男，2=女"`
	CreateTime        *gtime.Time `json:"createTime"        description:""`
	UpdateTime        *gtime.Time `json:"updateTime"        description:""`
	IsDelete          bool        `json:"isDelete"          description:"是否删除：true=已删除，false=正常"`
}
