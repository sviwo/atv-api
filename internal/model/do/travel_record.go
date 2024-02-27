// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-02-26 19:31:45
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TravelRecord is the golang structure of table sw_travel_record for DAO operations like Where/Data.
type TravelRecord struct {
	g.Meta         `orm:"table:sw_travel_record, do:true"`
	TravelRecordId interface{} //
	UserId         interface{} // 用户ID
	CarId          interface{} // 车辆ID
	StartPoint     interface{} // 起点
	EndPoint       interface{} // 终点
	MileageDriven  interface{} // 行驶里程，单位（m）
	StartTime      *gtime.Time // 行程开始时间
	EndTime        *gtime.Time // 行程结束时间
	AvgSpeed       interface{} // 平均时速，单位（m）
	Consumption    interface{} // 使用电量
	CreateTime     *gtime.Time //
	UpdateTime     *gtime.Time //
	IsDelete       interface{} // 是否删除：true=已删除，false=正常
}
