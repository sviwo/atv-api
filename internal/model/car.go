package model

import "github.com/gogf/gf/v2/os/gtime"

type CarBase struct {
	CarId          uint64      `json:"carId"              description:""`
	CarNickname    string      `json:"carNickname"        description:"车辆昵称（用户自定义）"`
	TravelKm       int         `json:"travelKm"           description:"行驶公里数"`
	CarFrameCode   string      `json:"carFrameCode"       description:"车架编号"`
	AfterSalesTime *gtime.Time `json:"afterSalesTime"     description:"保修日期"`
	ActivationTime *gtime.Time `json:"activationTime"     description:"激活时间"`
}

/*
查询车辆
*/
type QueryCarInput struct {
	CarNickname string `json:"carNickname"        description:"车辆昵称"`
}

type QueryCarOutput struct {
	CarBase
}

/*
控制车辆切换模式
*/
type ControlCarSwitchModeInput struct {
	CarId           uint64 `json:"carId"              description:""`
	DrivingModeType int    `json:"drivingModeType"    description:"驾驶模式：0=ECO模式，1=运动模式，2=狂暴模式，3=脱困模式"`
}

/*
绑定车辆
*/
type BindingCarInput struct {
	UserId string `json:"userId"          description:""`
}