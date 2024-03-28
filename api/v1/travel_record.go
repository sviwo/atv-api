package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type TravelRecordBase struct {
	TravelRecordId string      `json:"travelRecordId" description:""`
	DeviceId       uint64      `json:"deviceId"       description:""`
	StartPoint     string      `json:"startPoint"     description:"起点"`
	EndPoint       string      `json:"endPoint"       description:"终点"`
	MileageDriven  int         `json:"mileageDriven"  description:"行驶里程，单位（m）"`
	StartTime      *gtime.Time `json:"startTime"      description:"行程开始时间"`
	EndTime        *gtime.Time `json:"endTime"        description:"行程结束时间"`
	AvgSpeed       string      `json:"avgSpeed"       description:"平均时速，单位（m）"`
	Consumption    int         `json:"consumption"    description:"使用电量"`
}

type TravelRecordQueryReq struct {
	g.Meta   `path:"/travelRecord/list/get" method:"get" tags:"行程相关" sm:"获取行程列表"`
	DeviceId uint64 `json:"deviceId"       description:""`
	CommonPaginationReq
}

type TravelRecordRes struct {
	Result []*TravelRecordBase `json:"result" dc:"列表数据"`
	CommonPaginationRes
}

type TravelRecordDeleteReq struct {
	g.Meta         `path:"/travelRecord/delete" method:"post" tags:"行程相关" sm:"删除行程"`
	TravelRecordId int64 `json:"travelRecordId"       description:"" v:"required"`
}
