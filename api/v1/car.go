package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type GetCarInfoReq struct {
	g.Meta `path:"/car/get/list" method:"get" tags:"车辆相关" sm:"获取车辆列表"`
}

type GetCarListRes struct {
	DeviceId   string  `json:"deviceId"           dc:""`
	Nickname   string  `json:"nickname"           dc:"产品昵称（目前只有ATV，则等同于车辆昵称）"`
	Mileage    float32 `json:"mileage"            dc:"行驶里程（km）"`
	DeviceName string  `json:"deviceName"         dc:"设备名称（同于车架号）"`
}

type GetCarDetailReq struct {
	g.Meta   `path:"/car/get/detail" method:"get" tags:"车辆相关" sm:"获取车况信息"`
	DeviceId *int64 `json:"deviceId"           dc:""`
}

type GetCarDetailRes struct {
	DeviceId       string      `json:"deviceId"        dc:""`
	DeviceName     string      `json:"deviceName"      dc:"设备名称（同于车架号）"`
	MobileKey      int         `json:"mobileKey"       dc:"手机钥匙开关：0=关，1=开"`
	SpeedLimit     int         `json:"speedLimit"      dc:"速度限制开关：0=关，1=开"`
	DrivingMode    int         `json:"drivingMode"     dc:"驾驶模式：0=ECO模式，1=运动模式，2=狂暴模式"`
	EnergyRecovery int         `json:"energyRecovery"  dc:"动能回收类型：0=无，1=中，2=强"`
	ActivateTime   *gtime.Time `json:"activateTime"    dc:"激活时间"`
	WarrantyTime   *gtime.Time `json:"warrantyTime"    dc:"保修时间"`
	Mileage        float32     `json:"mileage"         dc:"行驶里程（km）"`
}

type BindingCarReq struct {
	g.Meta     `path:"/car/binding" method:"post" tags:"车辆相关" sm:"绑定车辆"`
	UserId     int64  `json:"userId"           dc:"用户id"                   v:"required"`
	DeviceName string `json:"deviceName"       dc:"车架号（最大长度20）"        v:"required|max-length:20"`
}

type SwitchCarReq struct {
	g.Meta   `path:"/car/switch" method:"post" tags:"车辆相关" sm:"切换车辆"`
	DeviceId int64 `json:"deviceId"           dc:""          v:"required"`
}

type DelCarReq struct {
	g.Meta   `path:"/car/del" method:"post" tags:"车辆相关" sm:"删除（解绑）车辆"`
	DeviceId int64 `json:"deviceId"           dc:""          v:"required"`
}

type EnabledMobileKeyReq struct {
	g.Meta `path:"/car/enabled/mobileKey" method:"post" tags:"车辆相关" sm:"开启/关闭蓝牙钥匙"`
}

type EnabledSpeedLimitReq struct {
	g.Meta `path:"/car/enabled/speedLimit" method:"post" tags:"车辆相关" sm:"开启/关闭速度限制"`
}

type CtlCarReq struct {
	g.Meta       `path:"/car/control/lamp" method:"post" tags:"车辆相关" sm:"控制车灯"`
	Instructions int `json:"instructions"  dc:"指令：0=灯光，1=鸣笛（不可输入其他指令）"   v:"required|between:0,1"`
}

type CtlSwitchDTReq struct {
	g.Meta          `path:"/car/control/switch/dt" method:"post" tags:"车辆相关" sm:"切换驾驶模式"`
	DrivingModeType int `json:"drivingModeType"    dc:"驾驶模式：0=ECO模式，1=运动模式，2=狂暴模式"`
}

type CtlSwitchERTReq struct {
	g.Meta             `path:"/car/control/switch/ert" method:"post" tags:"车辆相关" sm:"切换动能回收模式"`
	EnergyRecoveryType int `json:"energyRecoveryType"    description:"动能回收类型：0=无，1=中，2=强"`
}
