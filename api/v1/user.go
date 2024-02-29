package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type UserInfoBase struct {
	FirstName   string            `json:"firstName"            dc:""  v:"max-length:100"`
	LastName    string            `json:"lastName"             dc:""  v:"max-length:100"`
	HeadImg     *ghttp.UploadFile `json:"headImg" type:"file"  dc:"头像，请选择上传文件"`
	MobilePhone string            `json:"mobilePhone"          dc:""  v:"integer"`
	UserAddress string            `json:"userAddress"          dc:""  v:"max-length:100"`
}

type RegisterReq struct {
	g.Meta          `path:"/register" method:"post" tags:"用户相关" summary:"用户注册"`
	Username        string `json:"username"         dc:"用户名"     v:"required|email"`
	Password        string `json:"password"         dc:"密码"       v:"required|password3"`
	ConfirmPassword string `json:"confirmPassword"  dc:"确认密码"    v:"required|eq:Password"`
	EmailVftCode    string `json:"emailVftCode"     dc:"邮箱验证码"   v:"required|size:6"`
}

type UserInfoReq struct {
	g.Meta `path:"/user/info" method:"get" tags:"用户相关" summary:"当前登录用户信息"`
}

type UserInfoRes struct {
	Username string `json:"username"          dc:""`
	UserInfoBase
}

type UpdatePasswordReq struct {
	g.Meta          `path:"/update/password" method:"post" tags:"用户相关" summary:"修改密码"`
	Username        string `json:"username"          dc:""      v:"required|email"`
	NewPassword     string `json:"newPassword"       dc:""      v:"required|password3"`
	ConfirmPassword string `json:"confirmPassword"   dc:""      v:"required|eq:NewPassword"`
	EmailVftCode    string `json:"emailVftCode"      dc:""      v:"required|size:6"`
}

type EditInfoReq struct {
	g.Meta `path:"/edit/user/info" method:"post" tags:"用户相关" summary:"编辑用户信息"`
	UserInfoBase
}
