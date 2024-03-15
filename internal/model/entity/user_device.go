// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-03-15 14:53:40
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserDevice is the golang structure for table user_device.
type UserDevice struct {
	Id         uint64      `json:"id"         description:""`
	DeviceId   int64       `json:"deviceId"   description:""`
	UserId     int64       `json:"userId"     description:""`
	IsSelect   bool        `json:"isSelect"   description:"是否选定：false=未选定，true=已选定"`
	MobileKey  bool        `json:"mobileKey"  description:"手机钥匙开关：false=关，true=开"`
	SpeedLimit bool        `json:"speedLimit" description:"速度限制开关：false=关，true=开"`
	CreateTime *gtime.Time `json:"createTime" description:""`
}
