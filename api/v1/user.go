package v1

import (
	"github.com/gogf/gf/v2/frame/g"
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
RegisterReq 用户注册
*/
type RegisterReq struct {
	g.Meta            `path:"/register" method:"post" tags:"用户相关" summary:"用户注册"`
	Username          string `json:"username"    v:"email"     description:"用户名" v:"required"`
	Password          string `json:"password"    v:"password3"  description:"密码"`
	HeadImg           string `json:"headImg"       description:"头像"`
	UserSalt          string `json:"userSalt"     description:"加密盐 生成密码用"`
	Gender            int    `json:"gender"          description:"1男 2女"`
	Enable            int    `json:"enable"       description:"1正常 2拉黑冻结"`
	PersonalSignature string `json:"personalSignature"         description:"个性签名"`
}

type RegisterRes struct{}

/*
UserInfoReq 登陆用户信息
*/
type UserInfoReq struct {
	g.Meta   `path:"/user/info" method:"get" tags:"用户相关" summary:"当前登录用户信息"`
	Username string `json:"username"      v:"required"     description:""`
}

type UserInfoRes struct {
	Username string `json:"username"          description:""`
	VipLevel int    `json:"vipLevel"          description:"vip级别"`
	IpRegion string `json:"ipRegion"          description:"Ip属地"`
	UserInfoBase
}

/*
UpdatePasswordReq 修改密码
*/
type UpdatePasswordReq struct {
	g.Meta   `path:"/update/password" method:"post" tags:"用户相关" summary:"修改密码"`
	UserId   uint64 `json:"userId"       v:"required"     description:""`
	Password string `json:"password"     v:"required"     description:""`
}

type UpdatePasswordRes struct{}

type EditInfoReq struct {
	g.Meta `path:"/edit/user/info" method:"post" tags:"用户相关" summary:"修改密码"`
	UserInfoBase
}

type EditInfoRes struct{}
