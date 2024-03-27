package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type DeviceSecretReq struct {
	g.Meta     `path:"/device/get/secret" method:"get" tags:"设备相关" sm:"获取注册设备到指定产品下所需要的证书"`
	DeviceCode string `json:"deviceCode" dc:"设备的唯一标识" v:"required"`
}

type DeviceSecretRes struct {
	ProductKey   string `json:"productKey" dc:"对应物联网平台产品的ProductKey"`
	DeviceName   string `json:"deviceName" dc:"对应物联网平台颁发的设备证书的DeviceName"`
	DeviceSecret string `json:"deviceSecret" dc:"对应物联网平台颁发的设备证书的DeviceSecret"`
}
