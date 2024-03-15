// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-03-15 14:53:40
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserDevice is the golang structure of table sw_user_device for DAO operations like Where/Data.
type UserDevice struct {
	g.Meta     `orm:"table:sw_user_device, do:true"`
	Id         interface{} //
	DeviceId   interface{} //
	UserId     interface{} //
	IsSelect   interface{} // 是否选定：false=未选定，true=已选定
	MobileKey  interface{} // 手机钥匙开关：false=关，true=开
	SpeedLimit interface{} // 速度限制开关：false=关，true=开
	CreateTime *gtime.Time //
}
