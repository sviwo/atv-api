package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CarCleanBindReq struct {
	g.Meta `path:"/car/clean/bind" method:"get" tags:"车辆相关" sm:"清除车辆绑定（临时接口）"`
}

type GetCarInfoReq struct {
	g.Meta `path:"/car/get/list" method:"get" tags:"车辆相关" sm:"获取车辆列表"`
}

type GetCarListRes struct {
	DeviceId       string  `json:"deviceId"           dc:""`
	Nickname       string  `json:"nickname"           dc:"产品昵称（目前只有ATV，则等同于车辆昵称）"`
	Mileage        float32 `json:"mileage"            dc:"行驶里程（km）"`
	DeviceName     string  `json:"deviceName"         dc:"设备名称（同于车架号）"`
	IsSelect       bool    `json:"isSelect"           dc:"是否选定：false=未选定，true=已选定"`
	UserDeviceType int     `json:"userDeviceType"  dc:"设备用户类型：0=主用户，1=从用户"`
}

type GetCarDetailReq struct {
	g.Meta   `path:"/car/get/detail" method:"get" tags:"车辆相关" sm:"获取车况信息" dc:"车况信息、动能模式、安全性"`
	DeviceId *int64 `json:"deviceId"           dc:""`
}

type GetCarDetailRes struct {
	DeviceId        string          `json:"deviceId"        dc:""`
	DeviceName      string          `json:"deviceName"      dc:"设备名称（同于车架号）"`
	Nickname        string          `json:"nickname"        dc:"产品昵称（目前只有ATV，则等同于车辆昵称）"`
	MobileKey       bool            `json:"mobileKey"       dc:"手机钥匙开关：false=关，true=开"`
	SpeedLimit      bool            `json:"speedLimit"      dc:"速度限制开关：false=关，true=开"`
	DrivingModeType int             `json:"drivingModeType"     dc:"驾驶模式：0=ECO模式，1=运动模式，2=狂暴模式"`
	EnergyRecovery  int             `json:"energyRecovery"  dc:"动能回收类型：0=无，1=中，2=强"`
	UserDeviceType  int             `json:"userDeviceType"  dc:"设备用户类型：0=主用户，1=从用户"`
	ActivateTime    string          `json:"activateTime"    dc:"激活时间"`
	WarrantyTime    string          `json:"warrantyTime"    dc:"保修时间"`
	Mileage         float32         `json:"mileage"         dc:"行驶里程（km）"`
	TopSpeedHour    int             `json:"topSpeedHour"    dc:"最高时速"`
	UserCarKeyList  []UserCarKeyRes `json:"userCarKeyList"  dc:"车辆钥匙组"`
}

type UserCarKeyRes struct {
	UserDeviceId string `json:"userDeviceId"        dc:""`
	Name         string `json:"name"                dc:""`
	HeadImg      string `json:"headImg"             dc:""`
}

type CarKeyReq struct {
	g.Meta `path:"/car/get/key" method:"get" tags:"车辆相关" sm:"获取车钥匙" dc:"此接口限制30秒访问一次"`
}

type CarKeyRes struct {
	CarKey string `json:"carKey"    dc:"车钥匙（有效期一小时）"`
}

type InviteBindCarReq struct {
	g.Meta `path:"/car/invite/bind" method:"post" tags:"车辆相关" sm:"邀请绑定车辆"`
	CarKey string `json:"carKey"    dc:"车钥匙（长度限制38位）"  v:"required|max-length:38"`
}

type SwitchCarReq struct {
	g.Meta   `path:"/car/switch" method:"post" tags:"车辆相关" sm:"切换车辆"`
	DeviceId int64 `json:"deviceId"           dc:""          v:"required"`
}

type RemoveCarReq struct {
	g.Meta       `path:"/car/remove" method:"post" tags:"车辆相关" sm:"删除（解绑）车辆"`
	UserDeviceId *int64 `json:"userDeviceId"           dc:"钥匙id"`
	DeviceId     *int64 `json:"deviceId"               dc:""`
}

type EditCarNicknameReq struct {
	g.Meta   `path:"/car/edit/nickname" method:"post" tags:"车辆相关" sm:"编辑车辆昵称"`
	Nickname string `json:"nickname"        dc:"产品昵称（目前只有ATV，则等同于车辆昵称）"`
}

type EnabledMobileKeyReq struct {
	g.Meta `path:"/car/enabled/mobileKey" method:"post" tags:"车辆相关" sm:"开启/关闭蓝牙钥匙"`
}

type EnabledSpeedLimitReq struct {
	g.Meta `path:"/car/enabled/speedLimit" method:"post" tags:"车辆相关" sm:"开启/关闭速度限制"`
}

type CtlCarReq struct {
	g.Meta       `path:"/car/control/lamp" method:"post" tags:"车辆相关" sm:"控制车辆"`
	Instructions int `json:"instructions"  dc:"指令：0=灯光，1=鸣笛（不可输入其他指令）"   v:"required|between:0,1"`
}

type CtlSwitchDTReq struct {
	g.Meta          `path:"/car/control/switch/dt" method:"post" tags:"车辆相关" sm:"切换驾驶模式"`
	DrivingModeType int `json:"drivingModeType"    dc:"驾驶模式：0=ECO模式，1=运动模式，2=狂暴模式"`
}

type CtlSwitchERTReq struct {
	g.Meta             `path:"/car/control/switch/ert" method:"post" tags:"车辆相关" sm:"切换动能回收模式"`
	EnergyRecoveryType int `json:"energyRecoveryType"    dc:"动能回收类型：0=无，1=中，2=强"`
}

type SimDataTrafficReq struct {
	g.Meta `path:"/car/sim/data/traffic" method:"get" tags:"车辆相关" sm:"查询车辆sim卡流量"`
}

type SimDataTrafficRes struct {
	TotalDataTraffic   string `json:"totalDataTraffic"          dc:"总流量"`
	ConsumeDataTraffic string `json:"consumeDataTraffic"        dc:"已消耗流量"`
	SurplusDataTraffic string `json:"surplusDataTraffic"        dc:"剩余流量"`
}
