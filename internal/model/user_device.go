package model

type UserDeviceOutput struct {
	DeviceId   uint64 `json:"deviceId"           description:""`
	IsSelect   bool   `json:"isSelect"        description:"是否选定：false=未选定，true=已选定"`
	MobileKey  bool   `json:"mobileKey"       description:"手机钥匙开关：0=关，1=开"`
	SpeedLimit bool   `json:"speedLimit"      description:"速度限制开关：0=关，1=开"`
}
