package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type UserInfoBase struct {
	UserId      uint64 `json:"userId"            description:""  v:"required"`
	FirstName   string `json:"firstName"         description:""  v:"max-length:100"`
	LastName    string `json:"lastName"          description:""  v:"max-length:100"`
	HeadImg     string `json:"headImg"           description:""`
	MobilePhone string `json:"mobilePhone"       description:""  v:"integer"`
	UserAddress string `json:"userAddress"       description:""  v:"max-length:100"`
}

/*
RegisterReq 用户注册
*/
type RegisterReq struct {
	g.Meta          `path:"/register" method:"post" tags:"用户相关" summary:"用户注册"`
	Username        string `json:"username"         description:"用户名"     v:"required|email"`
	Password        string `json:"password"         description:"密码"       v:"required|password3"`
	ConfirmPassword string `json:"confirmPassword"  description:"确认密码"    v:"required|eq:Password"`
	EmailVftCode    string `json:"emailVftCode"     description:"邮箱验证码"   v:"required|size:6"`
}

type RegisterRes struct{}

/*
UserInfoReq 登陆用户信息
*/
type UserInfoReq struct {
	g.Meta `path:"/user/info" method:"get" tags:"用户相关" summary:"当前登录用户信息"`
	UserId string `json:"userId"           description:"" v:"required"`
}

type UserInfoRes struct {
	Username string `json:"username"          description:""`
	UserInfoBase
}

/*
UpdatePasswordReq 修改密码
*/
type UpdatePasswordReq struct {
	g.Meta          `path:"/update/password" method:"post" tags:"用户相关" summary:"修改密码"`
	Username        string `json:"username"          description:""      v:"required|email"`
	NewPassword     string `json:"newPassword"       description:""      v:"required|password3"`
	ConfirmPassword string `json:"confirmPassword"   description:""      v:"required|eq:NewPassword"`
	EmailVftCode    string `json:"emailVftCode"      description:""      v:"required|size:6"`
}

type UpdatePasswordRes struct{}

type EditInfoReq struct {
	g.Meta `path:"/edit/user/info" method:"post" tags:"用户相关" summary:"编辑用户信息"`
	UserInfoBase
}

type EditInfoRes struct{}
