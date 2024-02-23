// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-02-23 10:50:21
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
	MobileKey       bool        `json:"mobileKey"       description:"手机钥匙开关：0=关，1=开"`
	SpeedLimit      bool        `json:"speedLimit"      description:"速度限制开关：0=关，1=开"`
	DrivingModeType int         `json:"drivingModeType" description:"驾驶模式：0=ECO模式，1=运动模式，2=狂暴模式，3=脱困模式"`
	CreateTime      *gtime.Time `json:"createTime"      description:""`
}
