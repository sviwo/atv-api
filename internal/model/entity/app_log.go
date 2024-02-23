// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-02-23 10:50:21
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AppLog is the golang structure for table app_log.
type AppLog struct {
	AppLogId   int64       `json:"appLogId"   description:""`
	UserId     int64       `json:"userId"     description:""`
	BusinessId int64       `json:"businessId" description:"业务id，与类型字段组合查询"`
	LogType    int         `json:"logType"    description:"操作类型：0=其他，1=车辆，2=账户"`
	LogContent string      `json:"logContent" description:"操作内容"`
	CreateTime *gtime.Time `json:"createTime" description:""`
	UpdateTime *gtime.Time `json:"updateTime" description:""`
	IsDelete   bool        `json:"isDelete"   description:"是否删除：true=已删除，false=正常"`
}
