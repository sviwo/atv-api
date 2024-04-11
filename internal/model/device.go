package model

type QueryCarOutput struct {
	DeviceId   uint64  `json:"deviceId"    description:""`
	Mileage    float32 `json:"mileage"     description:"行驶里程（km）"`
	DeviceName string  `json:"deviceName"  description:"设备名称（同于车架号）"`
	Nickname   string  `json:"nickname"    description:"产品昵称（目前只有ATV，则等同于车辆昵称）"`
}

type CtlSwitchDTInput struct {
	DrivingModeType int `json:"drivingModeType"    description:"驾驶模式：0=ECO模式，1=运动模式，2=狂暴模式"`
}

type CtlSwitchERTInput struct {
	EnergyRecoveryType int `json:"energyRecoveryType"    description:"动能回收类型：0=无，1=中，2=强"`
}
