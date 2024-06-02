// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-05-10 15:47:49
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Battery is the golang structure for table battery.
type Battery struct {
	BatteryId   uint64      `json:"batteryId"   orm:"battery_id"   description:""`
	CarId       int64       `json:"carId"       orm:"car_id"       description:"车辆id"`
	BatteryCode string      `json:"batteryCode" orm:"battery_code" description:"电池编号"`
	BatteryTemp uint        `json:"batteryTemp" orm:"battery_temp" description:"电池温度"`
	CreateTime  *gtime.Time `json:"createTime"  orm:"create_time"  description:""`
	UpdateTime  *gtime.Time `json:"updateTime"  orm:"update_time"  description:""`
	IsDelete    bool        `json:"isDelete"    orm:"is_delete"    description:"是否删除：true=已删除，false=正常"`
}
