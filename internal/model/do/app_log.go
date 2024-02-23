// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-02-23 10:50:21
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AppLog is the golang structure of table sw_app_log for DAO operations like Where/Data.
type AppLog struct {
	g.Meta     `orm:"table:sw_app_log, do:true"`
	AppLogId   interface{} //
	UserId     interface{} //
	BusinessId interface{} // 业务id，与类型字段组合查询
	LogType    interface{} // 操作类型：0=其他，1=车辆，2=账户
	LogContent interface{} // 操作内容
	CreateTime *gtime.Time //
	UpdateTime *gtime.Time //
	IsDelete   interface{} // 是否删除：true=已删除，false=正常
}
