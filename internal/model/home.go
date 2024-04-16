package model

type HomeDataOutput struct {
	Nickname       string                 `json:"nickname"         dc:"车辆昵称"`
	RemainMile     float32                `json:"remainMile"       dc:"剩余里程（km）"`
	Electricity    int                    `json:"electricity"      dc:"电池电量"`
	BatteryStatus  int                    `json:"batteryStatus"    dc:"电池状态：0=放电，1=充电"`
	LockedStatus   int                    `json:"lockedStatus"     dc:"锁车状态：0=关机，1=开机"`
	GeoLocation    map[string]interface{} `json:"geoLocation"      dc:"地理位置"`
	IsHavingCar    bool                   `json:"isHavingCar"      dc:"是否有车：false=没有，true=有"`
	UserDeviceType int                    `json:"userDeviceType"   dc:"设备用户类型：0=主用户，1=从用户"`
	Version        []*VersionOutput       `json:"version"          dc:"新版本信息"`
}
