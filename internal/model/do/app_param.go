// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-06-08 20:37:43
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AppParam is the golang structure of table sw_app_param for DAO operations like Where/Data.
type AppParam struct {
	g.Meta     `orm:"table:sw_app_param, do:true"`
	ParamId    interface{} //
	ParentId   interface{} //
	ParamName  interface{} // 参数名称
	ParamValue interface{} // 参数值
	ParamConst interface{} // 参数常量
	CreateTime *gtime.Time //
	UpdateTime *gtime.Time //
	IsDelete   interface{} // 是否删除：true=已删除，false=正常
}
