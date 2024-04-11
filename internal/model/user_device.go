package model

import "github.com/gogf/gf/v2/os/gtime"

type UserDeviceOutput struct {
	DeviceId       uint64      `json:"deviceId"        description:""`
	DeviceName     string      `json:"deviceName"      description:"设备名称（同于车架号）"`
	MobileKey      int         `json:"mobileKey"       description:"手机钥匙开关：0=关，1=开"`
	SpeedLimit     int         `json:"speedLimit"      description:"速度限制开关：0=关，1=开"`
	DrivingMode    int         `json:"drivingMode"     description:"驾驶模式：0=ECO模式，1=运动模式，2=狂暴模式"`
	EnergyRecovery int         `json:"energyRecovery"  description:"动能回收类型：0=无，1=中，2=强"`
	ActivateTime   *gtime.Time `json:"activateTime"    description:"激活时间"`
	WarrantyTime   *gtime.Time `json:"warrantyTime"    description:"保修时间"`
	Mileage        float32     `json:"mileage"         description:"行驶里程（km）"`
}
