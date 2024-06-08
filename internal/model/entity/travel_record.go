// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-06-08 20:37:43
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TravelRecord is the golang structure for table travel_record.
type TravelRecord struct {
	TravelRecordId int64       `json:"travelRecordId" orm:"travel_record_id" description:""`
	UserId         int64       `json:"userId"         orm:"user_id"          description:"用户ID"`
	DeviceId       int64       `json:"deviceId"       orm:"device_id"        description:"设备ID"`
	StartPoint     string      `json:"startPoint"     orm:"start_point"      description:"起点"`
	EndPoint       string      `json:"endPoint"       orm:"end_point"        description:"终点"`
	MileageDriven  int         `json:"mileageDriven"  orm:"mileage_driven"   description:"行驶里程，单位（m）"`
	StartTime      *gtime.Time `json:"startTime"      orm:"start_time"       description:"行程开始时间"`
	EndTime        *gtime.Time `json:"endTime"        orm:"end_time"         description:"行程结束时间"`
	AvgSpeed       float64     `json:"avgSpeed"       orm:"avg_speed"        description:"平均时速，单位（km/h）"`
	Consumption    int         `json:"consumption"    orm:"consumption"      description:"使用电量"`
	CreateTime     *gtime.Time `json:"createTime"     orm:"create_time"      description:""`
	UpdateTime     *gtime.Time `json:"updateTime"     orm:"update_time"      description:""`
	IsDelete       bool        `json:"isDelete"       orm:"is_delete"        description:"是否删除：true=已删除，false=正常"`
}
