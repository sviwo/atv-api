// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-02-23 10:50:21
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AppText is the golang structure of table sw_app_text for DAO operations like Where/Data.
type AppText struct {
	g.Meta      `orm:"table:sw_app_text, do:true"`
	TextId      interface{} //
	TextName    interface{} // 文本名称
	TextType    interface{} // 文本类型：0=账户，1=车辆功能和充电，2=售后服务和维修保养，3=其他
	TextContent interface{} // 文本内容
	CreateTime  *gtime.Time //
	UpdateTime  *gtime.Time //
	IsDelete    interface{} // 是否删除：true=已删除，false=正常
}
