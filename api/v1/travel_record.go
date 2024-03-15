package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type TravelRecordQueryReq struct {
	g.Meta   `path:"/travelRecord/list/get" method:"get" tags:"行程相关" sm:"获取行程列表"`
	UserId   int64  `json:"userId"         dc:"用户ID"     `
	DeviceId uint64 `json:"deviceId"       description:""`
	CommonPaginationReq
}

type TravelRecordRes struct {
	Result interface{} `dc:"列表数据"`
	CommonPaginationRes
}

type TravelRecordDeleteReq struct {
	g.Meta `path:"/travelRecord/delete" method:"post" tags:"行程相关" sm:"删除行程"`
	TravelRecordReq
}
type TravelRecordReq struct {
	TravelRecordId int64       `json:"travelRecordId" dc:""`
	UserId         string      `json:"userId"         dc:"用户ID"`
	DeviceId       uint64      `json:"deviceId"       description:""`
	StartPoint     string      `json:"startPoint"     dc:"起点"`
	EndPoint       string      `json:"endPoint"       dc:"终点"`
	MileageDriven  int         `json:"mileageDriven"  dc:"行驶里程，单位（m）"`
	StartTime      *gtime.Time `json:"startTime"      dc:"行程开始时间"`
	EndTime        *gtime.Time `json:"endTime"        dc:"行程结束时间"`
	AvgSpeed       string      `json:"avgSpeed"       dc:"平均时速，单位（m）"`
	Consumption    int         `json:"consumption"    dc:"使用电量"`
}
