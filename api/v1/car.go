package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type GetCarInfoReq struct {
	g.Meta `path:"/car/get/list" method:"get" tags:"车辆相关" sm:"获取车辆列表"`
}

type CarInfo struct {
	Nickname       string      `json:"nickname"           dc:"产品昵称（目前只有ATV，则等同于车辆昵称）"`
	TravelKm       int         `json:"travelKm"           dc:"行驶公里数"`
	DeviceCode     string      `json:"deviceCode"         dc:"设备编号（同于车架号）"`
	AfterSalesTime *gtime.Time `json:"afterSalesTime"     dc:"保修日期"`
	ActivationTime *gtime.Time `json:"activationTime"     dc:"激活时间"`
}

type UserCarInfo struct {
	DeviceId        string `json:"deviceId"           dc:""`
	IsSelect        bool   `json:"isSelect"        dc:"是否选定：false=未选定，true=已选定"`
	MobileKey       bool   `json:"mobileKey"       dc:"手机钥匙开关：0=关，1=开"`
	SpeedLimit      bool   `json:"speedLimit"      dc:"速度限制开关：0=关，1=开"`
	DrivingModeType int    `json:"drivingModeType" dc:"驾驶模式：0=ECO模式，1=运动模式，2=狂暴模式，3=脱困模式"`
}

type GetCarInfoRes struct {
	*CarInfo
	*UserCarInfo
}

type BindingCarReq struct {
	g.Meta     `path:"/car/binding" method:"post" tags:"车辆相关" sm:"绑定车辆"`
	DeviceCode string `json:"deviceCode"       dc:"车架号（最大长度20）"          v:"required|max-length:20"`
}

type DelCarReq struct {
	g.Meta   `path:"/car/del" method:"post" tags:"车辆相关" sm:"删除（解绑）车辆"`
	DeviceId int64 `json:"deviceId"           dc:""          v:"required"`
}

type EnabledMobileKeyReq struct {
	g.Meta   `path:"/car/enabled/mobileKey" method:"post" tags:"车辆相关" sm:"开启/关闭蓝牙钥匙"`
	DeviceId int64 `json:"deviceId"           dc:""          v:"required"`
}

type EnabledSpeedLimitReq struct {
	g.Meta   `path:"/car/enabled/speedLimit" method:"post" tags:"车辆相关" sm:"开启/关闭速度限制"`
	DeviceId int64 `json:"deviceId"           dc:""          v:"required"`
}

type CtlLampReq struct {
	g.Meta   `path:"/car/control/lamp" method:"post" tags:"车辆相关" sm:"控制车灯"`
	DeviceId int64 `json:"deviceId"           dc:""          v:"required"`
}

type CtlHornReq struct {
	g.Meta   `path:"/car/control/horn" method:"post" tags:"车辆相关" sm:"控制喇叭"`
	DeviceId int64 `json:"deviceId"           dc:""          v:"required"`
}

type CtlSwitchDTReq struct {
	g.Meta   `path:"/car/control/switch/dt" method:"post" tags:"车辆相关" sm:"切换驾驶模式"`
	DeviceId int64 `json:"deviceId"           dc:""          v:"required"`
}

type CtlSwitchERTReq struct {
	g.Meta   `path:"/car/control/switch/ert" method:"post" tags:"车辆相关" sm:"切换动能回收模式"`
	DeviceId int64 `json:"deviceId"           dc:""          v:"required"`
}
