// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-03-27 15:02:22
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Device is the golang structure for table device.
type Device struct {
	DeviceId       uint64      `json:"deviceId"       description:""`
	ProductId      int64       `json:"productId"      description:"所属产品"`
	ProductKey     string      `json:"productKey"     description:"对应物联网平台产品的ProductKey"`
	DeviceName     string      `json:"deviceName"     description:"对应物联网平台颁发的设备证书的DeviceName"`
	DeviceSecret   string      `json:"deviceSecret"   description:"对应物联网平台颁发的设备证书的DeviceSecret"`
	DeviceModel    string      `json:"deviceModel"    description:"设备型号"`
	Nickname       string      `json:"nickname"       description:"产品昵称（目前只有ATV，则等同于车辆昵称）"`
	Status         int         `json:"status"         description:"状态：0=未启用，1=离线，2=在线"`
	OnlineTimeout  int         `json:"onlineTimeout"  description:"设备在线超时设置，单位：秒"`
	ActivateTime   *gtime.Time `json:"activateTime"   description:"激活时间"`
	RegistryTime   *gtime.Time `json:"registryTime"   description:"注册时间"`
	Version        string      `json:"version"        description:"固件版本号"`
	LastOnlineTime *gtime.Time `json:"lastOnlineTime" description:"最后上线时间"`
	CreateTime     *gtime.Time `json:"createTime"     description:""`
	UpdateTime     *gtime.Time `json:"updateTime"     description:""`
	MetadataTable  bool        `json:"metadataTable"  description:"是否生成物模型表：0=否，1=是"`
	IsDelete       bool        `json:"isDelete"       description:"是否删除：true=已删除，false=正常"`
}
