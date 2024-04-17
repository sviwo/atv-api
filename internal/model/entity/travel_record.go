// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-04-01 16:15:05
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TravelRecord is the golang structure for table travel_record.
type TravelRecord struct {
	TravelRecordId int64       `json:"travelRecordId" dc:""`
	UserId         int64       `json:"userId"         dc:"用户ID"`
	DeviceId       int64       `json:"deviceId"       dc:"设备ID"`
	StartPoint     string      `json:"startPoint"     dc:"起点"`
	EndPoint       string      `json:"endPoint"       dc:"终点"`
	MileageDriven  int         `json:"mileageDriven"  dc:"行驶里程，单位（m）"`
	StartTime      *gtime.Time `json:"startTime"      dc:"行程开始时间"`
	EndTime        *gtime.Time `json:"endTime"        dc:"行程结束时间"`
	AvgSpeed       float64     `json:"avgSpeed"       dc:"平均时速，单位（km/h）"`
	Consumption    int         `json:"consumption"    dc:"使用电量"`
	CreateTime     *gtime.Time `json:"createTime"     dc:""`
	UpdateTime     *gtime.Time `json:"updateTime"     dc:""`
	IsDelete       bool        `json:"isDelete"       dc:"是否删除：true=已删除，false=正常"`
}
