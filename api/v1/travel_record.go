package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type TravelRecordBase struct {
	TravelRecordId string `json:"travelRecordId" dc:""`
	DeviceId       uint64 `json:"deviceId"       dc:""`
	StartPoint     string `json:"startPoint"     dc:"起点"`
	EndPoint       string `json:"endPoint"       dc:"终点"`
	MileageDriven  int    `json:"mileageDriven"  dc:"行驶里程，单位（m）"`
	Duration       int    `json:"duration"       dc:"时长，单位（min）"`
	AvgSpeed       string `json:"avgSpeed"       dc:"平均时速，单位（m）"`
	Consumption    int    `json:"consumption"    dc:"使用电量"`
}

type TravelRecordQueryReq struct {
	g.Meta   `path:"/travelRecord/list/get" method:"get" tags:"行程相关" sm:"获取行程列表"`
	DeviceId uint64 `json:"deviceId"       dc:""`
	CommonPaginationReq
}

type TravelRecordRes struct {
	Result []*TravelRecordBase `json:"result" dc:"列表数据"`
	CommonPaginationRes
}

type TravelRecordDeleteReq struct {
	g.Meta         `path:"/travelRecord/delete" method:"post" tags:"行程相关" sm:"删除行程"`
	TravelRecordId int64 `json:"travelRecordId"       dc:"" v:"required"`
}
