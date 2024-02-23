// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-02-23 10:50:21
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserCar is the golang structure of table sw_user_car for DAO operations like Where/Data.
type UserCar struct {
	g.Meta          `orm:"table:sw_user_car, do:true"`
	UserCarId       interface{} //
	CarId           interface{} //
	UserId          interface{} //
	MobileKey       interface{} // 手机钥匙开关：0=关，1=开
	SpeedLimit      interface{} // 速度限制开关：0=关，1=开
	DrivingModeType interface{} // 驾驶模式：0=ECO模式，1=运动模式，2=狂暴模式，3=脱困模式
	CreateTime      *gtime.Time //
}
