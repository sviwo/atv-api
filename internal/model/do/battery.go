// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-03-15 14:53:40
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Battery is the golang structure of table sw_battery for DAO operations like Where/Data.
type Battery struct {
	g.Meta      `orm:"table:sw_battery, do:true"`
	BatteryId   interface{} //
	CarId       interface{} // 车辆id
	BatteryCode interface{} // 电池编号
	BatteryTemp interface{} // 电池温度
	CreateTime  *gtime.Time //
	UpdateTime  *gtime.Time //
	IsDelete    interface{} // 是否删除：true=已删除，false=正常
}
