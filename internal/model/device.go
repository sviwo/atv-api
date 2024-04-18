package model

type QueryCarOutput struct {
	DeviceId       uint64  `json:"deviceId"        dc:""`
	Mileage        float32 `json:"mileage"         dc:"行驶里程（km）"`
	DeviceName     string  `json:"deviceName"      dc:"设备名称（同于车架号）"`
	Nickname       string  `json:"nickname"        dc:"产品昵称（目前只有ATV，则等同于车辆昵称）"`
	IsSelect       bool    `json:"isSelect"        dc:"是否选定：false=未选定，true=已选定"`
	UserDeviceType int     `json:"userDeviceType"  dc:"设备用户类型：0=主用户，1=从用户"`
}

type CtlSwitchDTInput struct {
	DrivingModeType int `json:"drivingModeType"    dc:"驾驶模式：0=ECO模式，1=运动模式，2=狂暴模式"`
}

type CtlSwitchERTInput struct {
	EnergyRecoveryType int `json:"energyRecoveryType"    dc:"动能回收类型：0=无，1=中，2=强"`
}
