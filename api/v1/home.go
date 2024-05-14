package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type HomeDataReq struct {
	g.Meta `path:"/home/get/data" method:"get" tags:"首页相关" sm:"获取首页数据"`
}

type HomeDataRes struct {
	Nickname           string                 `json:"nickname"            dc:"车辆昵称"`
	RemainMile         int                    `json:"remainMile"          dc:"剩余里程（km）"`
	Electricity        int                    `json:"electricity"         dc:"电池电量（%）"`
	BatteryStatus      int                    `json:"batteryStatus"       dc:"电池状态：0=放电，1=充电"`
	LockedStatus       int                    `json:"lockedStatus"        dc:"锁车状态：0=关机，1=开机"`
	IsHavingCar        bool                   `json:"isHavingCar"         dc:"是否有车：false=没有，true=有"`
	UserDeviceType     int                    `json:"userDeviceType"      dc:"设备用户类型：0=主用户，1=从用户"`
	AuthStatus         int                    `json:"authStatus"          dc:"认证状态：0=未认证，1=认证中，2=认证成功，3=认证失败"`
	MobileKey          *bool                  `json:"mobileKey"           dc:"手机钥匙开关：false=关，true=开"`
	SpeedLimit         *bool                  `json:"speedLimit"          dc:"速度限制开关：false=关，true=开"`
	DrivingMode        int                    `json:"drivingMode"         dc:"驾驶模式：0=ECO模式，1=运动模式，2=狂暴模式"`
	EnergyRecovery     int                    `json:"energyRecovery"      dc:"动能回收类型：0=无，1=中，2=强"`
	ServicePhone       string                 `json:"servicePhone"        dc:"服务电话号码"`
	BluetoothAddress   string                 `json:"bluetoothAddress"    dc:"蓝牙MAC地址"`
	BluetoothSecretKey string                 `json:"bluetoothSecretKey"  dc:"蓝牙握手密钥"`
	GeoLocation        map[string]interface{} `json:"geoLocation"         dc:"地理位置"`
	Version            []*VersionRes          `json:"version"             dc:"新版本信息"`
}
