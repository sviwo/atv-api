// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-06-08 20:37:43
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserDevice is the golang structure for table user_device.
type UserDevice struct {
	Id             uint64      `json:"id"             orm:"id"               description:""`
	DeviceId       int64       `json:"deviceId"       orm:"device_id"        description:""`
	UserId         int64       `json:"userId"         orm:"user_id"          description:""`
	IsSelect       bool        `json:"isSelect"       orm:"is_select"        description:"是否选定：false=未选定，true=已选定"`
	UserDeviceType int         `json:"userDeviceType" orm:"user_device_type" description:"设备用户类型：0=主用户，1=从用户"`
	MobileKey      bool        `json:"mobileKey"      orm:"mobile_key"       description:"手机钥匙开关：false=关，true=开"`
	SpeedLimit     bool        `json:"speedLimit"     orm:"speed_limit"      description:"速度限制开关：false=关，true=开"`
	DrivingMode    int         `json:"drivingMode"    orm:"driving_mode"     description:"驾驶模式：0=ECO模式，1=运动模式，2=狂暴模式"`
	EnergyRecovery int         `json:"energyRecovery" orm:"energy_recovery"  description:"动能回收类型：0=无，1=中，2=强"`
	CreateTime     *gtime.Time `json:"createTime"     orm:"create_time"      description:""`
	UpdateTime     *gtime.Time `json:"updateTime"     orm:"update_time"      description:""`
}
