// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-01-16 21:22:22
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Car is the golang structure of table sw_car for DAO operations like Where/Data.
type Car struct {
	g.Meta             `orm:"table:sw_car, do:true"`
	CarId              interface{} //
	UserId             interface{} // 用户id
	CarName            interface{} // 车辆名称（厂家定义）
	CarNickname        interface{} // 车辆昵称（用户自定义）
	CarImg             interface{} //
	CarPosition        interface{} // 车辆定位
	ResidueElectricity interface{} // 剩余电量
	ResidueKm          interface{} // 剩余公里数
	CarFrameCode       interface{} // 车架编号
	CarMotorCode       interface{} // 车辆电机编号
	CreateTime         *gtime.Time //
	UpdateTime         *gtime.Time //
	IsDelete           interface{} // 是否删除：true=已删除，false=正常
}
