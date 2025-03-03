package model

type UserDeviceOutput struct {
	DeviceId        int64              `json:"deviceId"        dc:""`
	DeviceName      string             `json:"deviceName"      dc:"设备名称（同于车架号）"`
	Nickname        string             `json:"nickname"        dc:"产品昵称（目前只有ATV，则等同于车辆昵称）"`
	ActivateTime    string             `json:"activateTime"    dc:"激活时间"`
	WarrantyTime    string             `json:"warrantyTime"    dc:"保修时间"`
	MobileKey       bool               `json:"mobileKey"       dc:"手机钥匙开关：false=关，true=开"`
	SpeedLimit      bool               `json:"speedLimit"      dc:"速度限制开关：false=关，true=开"`
	UserDeviceType  int                `json:"userDeviceType"  dc:"设备用户类型：0=主用户，1=从用户"`
	DrivingModeType int                `json:"drivingModeType"     dc:"驾驶模式：0=ECO模式，1=运动模式，2=狂暴模式"`
	EnergyRecovery  int                `json:"energyRecovery"  dc:"动能回收类型：0=无，1=中，2=强"`
	Mileage         float32            `json:"mileage"         dc:"行驶里程（km）"`
	TopSpeedHour    int                `json:"topSpeedHour"    dc:"最高时速"`
	UserCarKeyList  []UserCarKeyOutput `json:"userCarKeyList"  dc:"车辆钥匙组"`
}

type UserCarKeyOutput struct {
	UserDeviceId int64  `json:"userDeviceId"        dc:""`
	Name         string `json:"name"                dc:""`
	HeadImg      string `json:"headImg"             dc:""`
}
