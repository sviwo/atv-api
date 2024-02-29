package model

import "github.com/gogf/gf/v2/os/gtime"

type UserInfoBase struct {
	FirstName   string  `json:"firstName"         description:""`
	LastName    string  `json:"lastName"          description:""`
	HeadImg     *string `json:"headImg"           description:""`
	MobilePhone string  `json:"mobilePhone"       description:""`
	UserAddress string  `json:"userAddress"       description:""`
}

/*
*用户登陆
 */
type LoginInput struct {
	Username  string `json:"username"     description:"用户名"`
	Password  string `json:"password"     description:"密码"`
	LoginType uint8  `json:"LoginType"    description:"登陆类型：1=账号+密码，2=第三方"`
}

/*
*用户注册
 */
type RegisterInput struct {
	Username        string `json:"username"          description:""`
	Password        string `json:"password"          description:""`
	ConfirmPassword string `json:"confirmPassword"   description:""`
	EmailVftCode    string `json:"emailVftCode"      description:""`
}

/*
获取用户信息
*/
type UserInfoOutput struct {
	Username string `json:"username"          description:""`
	UserInfoBase
}

/*
*用户修改密码
 */
type UpdatePasswordInput struct {
	Username        string `json:"username"       description:""`
	NewPassword     string `json:"newPassword"     description:""`
	ConfirmPassword string `json:"confirmPassword"     description:""`
	EmailVftCode    string `json:"emailVftCode"     description:""`
}

/*
*编辑用户资料
 */
type EditInfoInput struct {
	UserInfoBase
	UpdateTime *gtime.Time
}
