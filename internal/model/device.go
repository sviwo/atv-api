package model

type DeviceBase struct {
	DeviceId    uint64 `json:"deviceId"    description:""`
	ProductId   int64  `json:"productId"   description:"产品ID"`
	DeviceCode  string `json:"deviceCode"  description:"设备编号（同于车架号）"`
	DeviceName  string `json:"deviceName"  description:"设备名称"`
	DeviceModel string `json:"deviceModel" description:"设备型号"`
	Nickname    string `json:"nickname"    description:"产品昵称（目前只有ATV，则等同于车辆昵称）"`
}

type QueryCarOutput struct {
	*DeviceBase
	*UserDeviceOutput
}

type CtlSwitchDTInput struct {
	DeviceId        string `json:"deviceId"           dc:""`
	DrivingModeType int    `json:"drivingModeType"    description:"驾驶模式：0=ECO模式，1=运动模式，2=狂暴模式"`
}

type CtlSwitchERTInput struct {
	DeviceId        string `json:"deviceId"           dc:""`
	DrivingModeType int    `json:"drivingModeType"    description:"动能回收类型：0=无，1=中，2=强"`
}
