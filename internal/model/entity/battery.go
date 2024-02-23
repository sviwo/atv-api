// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-02-23 10:50:21
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Battery is the golang structure for table battery.
type Battery struct {
	BatteryId   uint64      `json:"batteryId"   description:""`
	CarId       int64       `json:"carId"       description:"车辆id"`
	BatteryCode string      `json:"batteryCode" description:"电池编号"`
	BatteryTemp uint        `json:"batteryTemp" description:"电池温度"`
	CreateTime  *gtime.Time `json:"createTime"  description:""`
	UpdateTime  *gtime.Time `json:"updateTime"  description:""`
	IsDelete    bool        `json:"isDelete"    description:"是否删除：true=已删除，false=正常"`
}
