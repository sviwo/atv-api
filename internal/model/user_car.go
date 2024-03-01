package model

type UserCarOutput struct {
	CarId           uint64 `json:"carId"           description:""`
	IsSelect        bool   `json:"isSelect"        description:"是否选定：false=未选定，true=已选定"`
	MobileKey       bool   `json:"mobileKey"       description:"手机钥匙开关：0=关，1=开"`
	SpeedLimit      bool   `json:"speedLimit"      description:"速度限制开关：0=关，1=开"`
	DrivingModeType int    `json:"drivingModeType" description:"驾驶模式：0=ECO模式，1=运动模式，2=狂暴模式，3=脱困模式"`
}
