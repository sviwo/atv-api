package model

import "github.com/gogf/gf/v2/os/gtime"

type TravelRecordQueryInput struct {
	UserId int64 `json:"userId"         description:"用户ID"     v:"required"`
	CarId  int64 `json:"carId"         description:"车辆ID"     `
	CommonPaginationInput
}

type TravelRecordOutput struct {
	TravelRecordId int64       `json:"travelRecordId" description:""`
	UserId         int64       `json:"userId"         description:"用户ID"`
	CarId          int64       `json:"carId"          description:"车辆ID"`
	StartPoint     string      `json:"startPoint"     description:"起点"`
	EndPoint       string      `json:"endPoint"       description:"终点"`
	MileageDriven  int         `json:"mileageDriven"  description:"行驶里程，单位（m）"`
	StartTime      *gtime.Time `json:"startTime"      description:"行程开始时间"`
	EndTime        *gtime.Time `json:"endTime"        description:"行程结束时间"`
	AvgSpeed       string      `json:"avgSpeed"       description:"平均时速，单位（m）"`
	Consumption    int         `json:"consumption"    description:"使用电量"`
}
