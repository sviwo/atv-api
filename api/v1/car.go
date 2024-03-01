package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type GetCarInfoReq struct {
	g.Meta `path:"/car/get/list" method:"get" tags:"车辆相关" summary:"获取车辆列表"`
}

type CarInfo struct {
	CarNickname    string      `json:"carNickname"        dc:"车辆昵称（用户自定义）"`
	TravelKm       int         `json:"travelKm"           dc:"行驶公里数"`
	CarFrameCode   string      `json:"carFrameCode"       dc:"车架编号"`
	CarMotorCode   string      `json:"carMotorCode"       dc:"车辆电机编号"`
	AfterSalesTime *gtime.Time `json:"afterSalesTime"     dc:"保修日期"`
	ActivationTime *gtime.Time `json:"activationTime"     dc:"激活时间"`
}

type UserCarInfo struct {
	CarId           string `json:"carId"           dc:""`
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
	g.Meta       `path:"/car/binding" method:"post" tags:"车辆相关" summary:"绑定车辆"`
	CarFrameCode string `json:"carFrameCode"       dc:"车架号"          v:"required"`
}

type DelCarReq struct {
	g.Meta `path:"/car/del" method:"post" tags:"车辆相关" summary:"删除（解绑）车辆"`
	CarId  int64 `json:"carId"           dc:""          v:"required"`
}

type EnabledMobileKeyReq struct {
	g.Meta `path:"/car/enabled/mobileKey" method:"post" tags:"车辆相关" summary:"开启/关闭蓝牙钥匙"`
	CarId  int64 `json:"carId"           dc:""          v:"required"`
}

type EnabledSpeedLimitReq struct {
	g.Meta `path:"/car/enabled/speedLimit" method:"post" tags:"车辆相关" summary:"开启/关闭速度限制"`
	CarId  int64 `json:"carId"           dc:""          v:"required"`
}

type ControlCarLampReq struct {
	g.Meta `path:"/car/control/lamp" method:"post" tags:"车辆相关" summary:"控制车灯"`
	CarId  int64 `json:"carId"           dc:""          v:"required"`
}

type ControlCarHornReq struct {
	g.Meta `path:"/car/control/horn" method:"post" tags:"车辆相关" summary:"控制喇叭"`
	CarId  int64 `json:"carId"           dc:""          v:"required"`
}

type ControlCarSwitchDriveModeReq struct {
	g.Meta `path:"/car/control/switch/driveMode" method:"post" tags:"车辆相关" summary:"切换驾驶模式"`
	CarId  int64 `json:"carId"           dc:""          v:"required"`
}

type ControlCarSwitchEnergyRecoveryTypeReq struct {
	g.Meta `path:"/car/Control/switch/energyRecoveryType" method:"post" tags:"车辆相关" summary:"切换动能回收模式"`
	CarId  int64 `json:"carId"           dc:""          v:"required"`
}
