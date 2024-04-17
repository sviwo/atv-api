// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-04-01 16:15:05
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AlarmLog is the golang structure for table alarm_log.
type AlarmLog struct {
	Id          int64       `json:"id"          dc:""`
	Type        uint        `json:"type"        dc:"告警类型：1=规则告警，2=设备自主告警"`
	Data        string      `json:"data"        dc:"触发告警的数据"`
	ProductKey  string      `json:"productKey"  dc:"产品标识"`
	DeviceKey   string      `json:"deviceKey"   dc:"设备标识"`
	Status      int         `json:"status"      dc:"告警状态：0=未处理，1=已处理"`
	CreatedTime *gtime.Time `json:"createdTime" dc:"告警时间"`
	UpdatedBy   uint        `json:"updatedBy"   dc:"告警处理人员"`
	UpdatedTime *gtime.Time `json:"updatedTime" dc:"处理时间"`
	Content     string      `json:"content"     dc:"处理意见"`
}
