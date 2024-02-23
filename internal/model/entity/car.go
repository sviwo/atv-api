// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-01-16 21:22:22
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Car is the golang structure for table car.
type Car struct {
	CarId              uint64      `json:"carId"              description:""`
	UserId             int64       `json:"userId"             description:"用户id"`
	CarName            string      `json:"carName"            description:"车辆名称（厂家定义）"`
	CarNickname        string      `json:"carNickname"        description:"车辆昵称（用户自定义）"`
	CarImg             string      `json:"carImg"             description:""`
	CarPosition        string      `json:"carPosition"        description:"车辆定位"`
	ResidueElectricity int         `json:"residueElectricity" description:"剩余电量"`
	ResidueKm          int         `json:"residueKm"          description:"剩余公里数"`
	CarFrameCode       string      `json:"carFrameCode"       description:"车架编号"`
	CarMotorCode       string      `json:"carMotorCode"       description:"车辆电机编号"`
	CreateTime         *gtime.Time `json:"createTime"         description:""`
	UpdateTime         *gtime.Time `json:"updateTime"         description:""`
	IsDelete           bool        `json:"isDelete"           description:"是否删除：true=已删除，false=正常"`
}
