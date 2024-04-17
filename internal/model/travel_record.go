package model

import "github.com/gogf/gf/v2/os/gtime"

type TravelRecordQueryInput struct {
	UserId   int64  `json:"userId"         dc:"用户ID"     v:"required"`
	DeviceId uint64 `json:"deviceId"       dc:""`
	CommonPaginationInput
}

type TravelRecordOutput struct {
	TravelRecordId int64             `json:"travelRecordId" dc:""`
	UserId         int64             `json:"userId"         dc:"用户ID"`
	DeviceId       uint64            `json:"deviceId"       dc:""`
	StartPoint     map[string]string `json:"startPoint"     dc:"起点"`
	EndPoint       map[string]string `json:"endPoint"       dc:"终点"`
	MileageDriven  int               `json:"mileageDriven"  dc:"行驶里程，单位（m）"`
	StartTime      *gtime.Time       `json:"startTime"      dc:"行程开始时间"`
	EndTime        *gtime.Time       `json:"endTime"        dc:"行程结束时间"`
	AvgSpeed       string            `json:"avgSpeed"       dc:"平均时速，单位（m）"`
	Consumption    int               `json:"consumption"    dc:"使用电量"`
}

type TravelRecordOnline struct {
	DeviceName string `json:"deviceName"       dc:"对应物联网平台颁发的设备证书的DeviceName"`
}
