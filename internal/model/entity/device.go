// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-03-15 14:53:40
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Device is the golang structure for table device.
type Device struct {
	DeviceId       uint64      `json:"deviceId"    description:""`
	ProductKey     string      `json:"ProductKey"   description:"产品Key"`
	DeviceCode     string      `json:"deviceCode"  description:"设备编号（同于车架号）"`
	DeviceName     string      `json:"deviceName"  description:"设备名称"`
	DeviceModel    string      `json:"deviceModel" description:"设备型号"`
	OnlineTimeout  int         `json:"onlineTimeout"  description:"设备在线超时设置，单位：秒"`
	RegistryTime   string      // 激活时间
	Version        string      // 固件版本号
	LastOnlineTime string      // 最后上线时间
	Status         int         `json:"status"         description:"状态：0=未启用,1=离线,2=在线"`
	Nickname       string      `json:"nickname"    description:"产品昵称（目前只有ATV，则等同于车辆昵称）"`
	CreateTime     *gtime.Time `json:"createTime"  description:""`
	UpdateTime     *gtime.Time `json:"updateTime"  description:""`
	IsDelete       bool        `json:"isDelete"    description:"是否删除：true=已删除，false=正常"`
	MetadataTable  string      `json:"metadata_table"    description:"是否生成物模型子表：0=否，1=是"`
}
