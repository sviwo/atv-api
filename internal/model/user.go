package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type UserInfoBase struct {
	UserId            uint64      `json:"userId"            description:""`
	FirstName         string      `json:"firstName"         description:""`
	LastName          string      `json:"lastName"          description:""`
	Nickname          string      `json:"nickname"          description:"昵称"`
	HeadImg           string      `json:"headImg"           description:""`
	MobilePhone       string      `json:"mobilePhone"       description:"手机号"`
	PersonalSignature string      `json:"personalSignature" description:"个性签名"`
	Birthday          *gtime.Time `json:"birthday"          description:"生日"`
	Gender            int         `json:"gender"            description:"性别：0=未知，1=男，2=女"`
}

/*
*用户登陆传参
 */
type LoginInput struct {
	Username  string `json:"username"     description:"用户名"`
	VftCode   string `json:"VftCode"     description:"验证码"`
	Password  string `json:"password"     description:"密码"`
	LoginType uint8  `json:"LoginType"     description:"登陆类型：1=验证码，2=密码"`
}

type LogoutInput struct {
	UserId uint64 `json:"userId"            description:""`
}

/*
*用户注册传参
 */
type RegisterInput struct {
	Username        string `json:"username"          description:"用户名"`
	Password        string `json:"password"          description:"密码"`
	ConfirmPassword string `json:"confirmPassword"          description:"确认密码"`
	//MobilePhone       string      `json:"mobilePhone"       description:"手机号"`
}

/*
获取用户信息传参
*/
type UserInfoInput struct {
	Username string `json:"username"          description:""`
}

/*
获取用户信息返参
*/
type UserInfoOutput struct {
	Username string `json:"username"          description:""`
	VipLevel int    `json:"vipLevel"          description:"vip级别"`
	IpRegion string `json:"ipRegion"          description:"Ip属地"`
	UserInfoBase
}

/*
*用户修改密码传参
 */
type UpdatePasswordInput struct {
	UserId   uint64 `json:"userId"       description:""`
	Password string `json:"password"     description:""`
}

/*
*编辑用户资料传参
 */
type EditInfoInput struct {
	UserInfoBase
}
