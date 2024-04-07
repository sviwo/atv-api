package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type HomeDataReq struct {
	g.Meta `path:"/home/get/data" method:"get" tags:"首页相关" sm:"获取首页数据"`
}

type HomeDataRes struct {
	Nickname      string        `json:"nickname"         dc:"车辆昵称"`
	RemainMile    int           `json:"remainMile"       dc:"剩余里程（km）"`
	Electricity   int           `json:"electricity"      dc:"电池电量（%）"`
	BatteryStatus int           `json:"batteryStatus"    dc:"电池状态：0=放电，1=充电"`
	LockedStatus  int           `json:"lockedStatus"     dc:"锁车状态：0=关机，1=开机"`
	GeoLocation   string        `json:"geoLocation"      dc:"地理位置"`
	Version       []*VersionRes `json:"version"          dc:"新版本信息"`
}
