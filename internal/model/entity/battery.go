// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-04-01 16:15:05
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Battery is the golang structure for table battery.
type Battery struct {
	BatteryId   uint64      `json:"batteryId"   dc:""`
	CarId       int64       `json:"carId"       dc:"车辆id"`
	BatteryCode string      `json:"batteryCode" dc:"电池编号"`
	BatteryTemp uint        `json:"batteryTemp" dc:"电池温度"`
	CreateTime  *gtime.Time `json:"createTime"  dc:""`
	UpdateTime  *gtime.Time `json:"updateTime"  dc:""`
	IsDelete    bool        `json:"isDelete"    dc:"是否删除：true=已删除，false=正常"`
}
