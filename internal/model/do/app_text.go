// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-03-07 16:22:32
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
	ParentId    interface{} //
	TextTitle   interface{} // 文本标题
	TextContent interface{} // 文本内容
	CreateTime  *gtime.Time //
	UpdateTime  *gtime.Time //
	IsDelete    interface{} // 是否删除：true=已删除，false=正常
}
