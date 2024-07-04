// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-06-08 20:37:43
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AlarmLog is the golang structure for table alarm_log.
type AlarmLog struct {
	Id          int64       `json:"id"          orm:"id"           description:""`
	Type        uint        `json:"type"        orm:"type"         description:"告警类型：1=规则告警，2=设备自主告警"`
	Data        string      `json:"data"        orm:"data"         description:"触发告警的数据"`
	ProductKey  string      `json:"productKey"  orm:"product_key"  description:"产品标识"`
	DeviceKey   string      `json:"deviceKey"   orm:"device_key"   description:"设备标识"`
	Status      int         `json:"status"      orm:"status"       description:"告警状态：0=未处理，1=已处理"`
	Content     string      `json:"content"     orm:"content"      description:"处理意见"`
	UpdatedBy   uint        `json:"updatedBy"   orm:"updated_by"   description:"告警处理人员"`
	CreatedTime *gtime.Time `json:"createdTime" orm:"created_time" description:"告警时间"`
	UpdatedTime *gtime.Time `json:"updatedTime" orm:"updated_time" description:"处理时间"`
}
