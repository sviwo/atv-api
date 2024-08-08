package model

import "github.com/gogf/gf/v2/os/gtime"

type UserInfoBase struct {
	FirstName   string  `json:"firstName"         dc:""`
	LastName    string  `json:"lastName"          dc:""`
	HeadImg     *string `json:"headImg"           dc:""`
	MobilePhone string  `json:"mobilePhone"       dc:""`
	UserAddress string  `json:"userAddress"       dc:""`
}

type LoginInput struct {
	Username       string `json:"username"       dc:"用户名"`
	Password       string `json:"password"       dc:"密码"`
	UserIdentifier string `json:"userIdentifier" dc:"苹果的cliUserID"`
	IdentityToken  string `json:"identityToken"  dc:"苹果的token"`
	AccessToken    string `json:"accessToken"    dc:"facebook的token"`
	LoginType      int    `json:"loginType"      dc:"登陆类型：1=账号+密码，2=apple，3=facebook"`
}

type RegisterInput struct {
	Username        string `json:"username"          dc:""`
	Password        string `json:"password"          dc:""`
	ConfirmPassword string `json:"confirmPassword"   dc:""`
	EmailVftCode    string `json:"emailVftCode"      dc:""`
}

type UserInfoOutput struct {
	UserInfoBase
	Username   string  `json:"username"          dc:""`
	Nickname   string  `json:"nickname"          dc:"车辆昵称"`
	DeviceName string  `json:"deviceName"        dc:"车架号"`
	Mileage    float32 `json:"mileage"           dc:"行驶里程(km)"`
	AuthStatus int     `json:"authStatus"        dc:"认证状态：0=未认证，1=认证中，2=认证成功，3=认证失败"`
}

type UpdatePasswordInput struct {
	Username        string `json:"username"        dc:""`
	NewPassword     string `json:"newPassword"     dc:""`
	ConfirmPassword string `json:"confirmPassword" dc:""`
	EmailVftCode    string `json:"emailVftCode"    dc:""`
}

type EditInfoInput struct {
	UserInfoBase
	UpdateTime *gtime.Time
}
