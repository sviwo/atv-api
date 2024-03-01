// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-03-01 17:48:00
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserCar is the golang structure for table user_car.
type UserCar struct {
	UserCarId       uint64      `json:"userCarId"       description:""`
	CarId           int64       `json:"carId"           description:""`
	UserId          int64       `json:"userId"          description:""`
	IsSelect        bool        `json:"isSelect"        description:"是否选定：false=未选定，true=已选定"`
	MobileKey       bool        `json:"mobileKey"       description:"手机钥匙开关：false=关，true=开"`
	SpeedLimit      bool        `json:"speedLimit"      description:"速度限制开关：false=关，true=开"`
	DrivingModeType int         `json:"drivingModeType" description:"驾驶模式：0=ECO模式，1=运动模式，2=狂暴模式，3=脱困模式"`
	CreateTime      *gtime.Time `json:"createTime"      description:""`
}
