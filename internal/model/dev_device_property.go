package model

type DevicePropertyInput struct {
	DeviceKey string         `json:"deviceKey" dc:"设备标识" v:"required#设备标识不能为空"`
	Params    map[string]any `json:"params" dc:"设备属性设置"`
}
type DevicePropertyOutput struct {
	Data map[string]any `json:"data" dc:"设备属性设置输出"`
}

type DeviceSecretOutput struct {
	ProductKey   string `json:"productKey" dc:"对应物联网平台产品的ProductKey"`
	DeviceName   string `json:"deviceName" dc:"对应物联网平台颁发的设备证书的DeviceName"`
	DeviceSecret string `json:"deviceSecret" dc:"对应物联网平台颁发的设备证书的DeviceSecret"`
}
