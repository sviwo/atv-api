package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type UserInfoBase struct {
	FirstName   string `json:"firstName"            dc:"最大长度100"  v:"max-length:100"`
	LastName    string `json:"lastName"             dc:"最大长度100"  v:"max-length:100"`
	MobilePhone string `json:"mobilePhone"          dc:"必须是数字"   v:"integer"`
	UserAddress string `json:"userAddress"          dc:"最大长度100"  v:"max-length:100"`
}

type RegisterReq struct {
	g.Meta          `path:"/user/register" method:"post" tags:"用户相关" sm:"用户注册"`
	Username        string `json:"username"        dc:"用户名（标准邮箱格式）"                       v:"required|email"`
	Password        string `json:"password"        dc:"密码（长度在6~18之间，必须包含大小写字母和数字）" v:"required|password2"`
	ConfirmPassword string `json:"confirmPassword" dc:"确认密码（必须与密码一致）"                    v:"required|eq:Password"`
	EmailVftCode    string `json:"emailVftCode"    dc:"邮箱验证码（长度必须6位）"                     v:"required|integer|size:6"`
}

type UserInfoReq struct {
	g.Meta `path:"/user/info" method:"get" tags:"用户相关" sm:"当前登录用户信息"`
}

type UserInfoRes struct {
	Username string `json:"username"          dc:""`
	UserInfoBase
	HeadImg string `json:"headImg"           description:""`
}

type UpdatePasswordReq struct {
	g.Meta          `path:"/user/update/password" method:"post" tags:"用户相关" sm:"修改密码"`
	Username        string `json:"username"          dc:"用户名（标准邮箱格式）"      v:"required|email"`
	NewPassword     string `json:"newPassword"       dc:"长度在6~18之间，必须包含大小写字母、数字和特殊字符" v:"required|password3"`
	ConfirmPassword string `json:"confirmPassword"   dc:"必须与密码一致"             v:"required|eq:NewPassword"`
	EmailVftCode    string `json:"emailVftCode"      dc:"邮箱验证码（长度必须6位）"    v:"required|size:6"`
}

type EditInfoReq struct {
	g.Meta `path:"/user/edit/info" method:"post" tags:"用户相关" sm:"编辑用户信息"`
	UserInfoBase
	HeadImg *ghttp.UploadFile `json:"headImg" type:"file"  dc:"头像，请选择上传文件（最大10MB）"`
}
