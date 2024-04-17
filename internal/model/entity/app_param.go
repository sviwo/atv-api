// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-04-01 16:15:05
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AppParam is the golang structure for table app_param.
type AppParam struct {
	ParamId    int64       `json:"paramId"    dc:""`
	ParentId   int64       `json:"parentId"   dc:""`
	ParamName  string      `json:"paramName"  dc:"参数名称"`
	ParamValue string      `json:"paramValue" dc:"参数值"`
	CreateTime *gtime.Time `json:"createTime" dc:""`
	UpdateTime *gtime.Time `json:"updateTime" dc:""`
	IsDelete   bool        `json:"isDelete"   dc:"是否删除：true=已删除，false=正常"`
}
