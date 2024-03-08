// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-03-04 17:02:28
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Car is the golang structure for table car.
type Car struct {
	CarId            uint64      `json:"carId"            description:""`
	CarNickname      string      `json:"carNickname"      description:"车辆昵称（用户自定义）"`
	CarFinalPosition string      `json:"carFinalPosition" description:"车辆最后定位"`
	TravelKm         int         `json:"travelKm"         description:"行驶公里数"`
	CarFrameCode     string      `json:"carFrameCode"     description:"车架编号"`
	CarMotorCode     string      `json:"carMotorCode"     description:"车辆电机编号"`
	AfterSalesTime   *gtime.Time `json:"afterSalesTime"   description:"保修日期"`
	ActivationTime   *gtime.Time `json:"activationTime"   description:"激活时间"`
	CreateTime       *gtime.Time `json:"createTime"       description:""`
	UpdateTime       *gtime.Time `json:"updateTime"       description:""`
	IsDelete         bool        `json:"isDelete"         description:"是否删除：true=已删除，false=正常"`
}
