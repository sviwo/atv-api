// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AlarmLog is the golang structure for table alarm_log.
type AlarmLog struct {
	Id         int64       `json:"id"         description:""`
	DeptId     int         `json:"deptId"     description:"部门ID"`
	Type       uint        `json:"type"       description:"告警类型：1=规则告警，2=设备自主告警"`
	Data       string      `json:"data"       description:"触发告警的数据"`
	ProductKey string      `json:"productKey" description:"产品标识"`
	DeviceKey  string      `json:"deviceKey"  description:"设备标识"`
	Status     int         `json:"status"     description:"告警状态：0=未处理，1=已处理"`
	CreatedTime  *gtime.Time `json:"createdTime"  description:"告警时间"`
	UpdatedBy  uint        `json:"updatedBy"  description:"告警处理人员"`
	UpdatedTime  *gtime.Time `json:"updatedTime"  description:"处理时间"`
	Content    string      `json:"content"    description:"处理意见"`
}
