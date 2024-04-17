// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-04-07 20:24:54
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Device is the golang structure for table device.
type Device struct {
	DeviceId       uint64      `json:"deviceId"       dc:""`
	ProductId      int64       `json:"productId"      dc:"所属产品"`
	ProductKey     string      `json:"productKey"     dc:"对应物联网平台产品的ProductKey"`
	DeviceName     string      `json:"deviceName"     dc:"对应物联网平台颁发的设备证书的DeviceName"`
	DeviceSecret   string      `json:"deviceSecret"   dc:"对应物联网平台颁发的设备证书的DeviceSecret"`
	DeviceModel    string      `json:"deviceModel"    dc:"设备型号"`
	Nickname       string      `json:"nickname"       dc:"产品昵称（目前只有ATV，则等同于车辆昵称）"`
	Status         int         `json:"status"         dc:"状态：0=未启用，1=离线，2=在线"`
	OnlineTimeout  int         `json:"onlineTimeout"  dc:"设备在线超时设置，单位：秒"`
	ActivateTime   *gtime.Time `json:"activateTime"   dc:"激活时间"`
	RegistryTime   *gtime.Time `json:"registryTime"   dc:"注册时间"`
	Version        string      `json:"version"        dc:"固件版本号"`
	LastOnlineTime *gtime.Time `json:"lastOnlineTime" dc:"最后上线时间"`
	CreateTime     *gtime.Time `json:"createTime"     dc:""`
	UpdateTime     *gtime.Time `json:"updateTime"     dc:""`
	MetadataTable  bool        `json:"metadataTable"  dc:"是否生成物模型表：0=否，1=是"`
	IsDelete       bool        `json:"isDelete"       dc:"是否删除：true=已删除，false=正常"`
}
