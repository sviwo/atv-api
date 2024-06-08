// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-06-08 20:37:43
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Device is the golang structure for table device.
type Device struct {
	DeviceId           uint64      `json:"deviceId"           orm:"device_id"            description:""`
	ProductId          int64       `json:"productId"          orm:"product_id"           description:"所属产品"`
	SimId              string      `json:"simId"              orm:"sim_id"               description:""`
	ProductKey         string      `json:"productKey"         orm:"product_key"          description:"对应物联网平台产品的ProductKey"`
	DeviceName         string      `json:"deviceName"         orm:"device_name"          description:"对应物联网平台颁发的设备证书的DeviceName"`
	DeviceSecret       string      `json:"deviceSecret"       orm:"device_secret"        description:"对应物联网平台颁发的设备证书的DeviceSecret"`
	DeviceModel        string      `json:"deviceModel"        orm:"device_model"         description:"设备型号"`
	Nickname           string      `json:"nickname"           orm:"nickname"             description:"产品昵称（目前只有ATV，则等同于车辆昵称）"`
	Status             int         `json:"status"             orm:"status"               description:"状态：0=未启用，1=离线，2=在线"`
	OnlineTimeout      int         `json:"onlineTimeout"      orm:"online_timeout"       description:"设备在线超时设置，单位：秒"`
	ActivateTime       *gtime.Time `json:"activateTime"       orm:"activate_time"        description:"激活时间"`
	RegistryTime       *gtime.Time `json:"registryTime"       orm:"registry_time"        description:"注册时间"`
	Version            string      `json:"version"            orm:"version"              description:"固件版本号"`
	LastOnlineTime     *gtime.Time `json:"lastOnlineTime"     orm:"last_online_time"     description:"最后上线时间"`
	MetadataTable      bool        `json:"metadataTable"      orm:"metadata_table"       description:"是否生成物模型表：0=否，1=是"`
	BluetoothAddress   string      `json:"bluetoothAddress"   orm:"bluetooth_address"    description:"蓝牙地址"`
	BluetoothSecretKey string      `json:"bluetoothSecretKey" orm:"bluetooth_secret_key" description:"蓝牙握手密钥"`
	CreateTime         *gtime.Time `json:"createTime"         orm:"create_time"          description:""`
	UpdateTime         *gtime.Time `json:"updateTime"         orm:"update_time"          description:""`
	IsDelete           bool        `json:"isDelete"           orm:"is_delete"            description:"是否删除：true=已删除，false=正常"`
}
