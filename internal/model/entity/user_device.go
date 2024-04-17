// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-04-14 22:19:23
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserDevice is the golang structure for table user_device.
type UserDevice struct {
	Id             uint64      `json:"id"             dc:""`
	DeviceId       int64       `json:"deviceId"       dc:""`
	UserId         int64       `json:"userId"         dc:""`
	IsSelect       bool        `json:"isSelect"       dc:"是否选定：false=未选定，true=已选定"`
	UserDeviceType int         `json:"userDeviceType" dc:"设备用户类型：0=主用户，1=从用户"`
	MobileKey      bool        `json:"mobileKey"      dc:"手机钥匙开关：false=关，true=开"`
	SpeedLimit     bool        `json:"speedLimit"     dc:"速度限制开关：false=关，true=开"`
	DrivingMode    int         `json:"drivingMode"    dc:"驾驶模式：0=ECO模式，1=运动模式，2=狂暴模式"`
	EnergyRecovery int         `json:"energyRecovery" dc:"动能回收类型：0=无，1=中，2=强"`
	CreateTime     *gtime.Time `json:"createTime"     dc:""`
	UpdateTime     *gtime.Time `json:"updateTime"     dc:""`
}
