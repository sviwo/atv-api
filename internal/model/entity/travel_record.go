// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-03-15 14:53:40
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TravelRecord is the golang structure for table travel_record.
type TravelRecord struct {
	TravelRecordId int64       `json:"travelRecordId" description:""`
	UserId         int64       `json:"userId"         description:"用户ID"`
	DeviceId       int64       `json:"deviceId"       description:"设备ID"`
	StartPoint     string      `json:"startPoint"     description:"起点"`
	EndPoint       string      `json:"endPoint"       description:"终点"`
	MileageDriven  int         `json:"mileageDriven"  description:"行驶里程，单位（m）"`
	StartTime      *gtime.Time `json:"startTime"      description:"行程开始时间"`
	EndTime        *gtime.Time `json:"endTime"        description:"行程结束时间"`
	AvgSpeed       string      `json:"avgSpeed"       description:"平均时速，单位（m）"`
	Consumption    int         `json:"consumption"    description:"使用电量"`
	CreateTime     *gtime.Time `json:"createTime"     description:""`
	UpdateTime     *gtime.Time `json:"updateTime"     description:""`
	IsDelete       bool        `json:"isDelete"       description:"是否删除：true=已删除，false=正常"`
}
