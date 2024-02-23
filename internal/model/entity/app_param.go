// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-02-23 10:50:21
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AppParam is the golang structure for table app_param.
type AppParam struct {
	ParamId    int64       `json:"paramId"    description:""`
	ParentId   int64       `json:"parentId"   description:""`
	ParamName  string      `json:"paramName"  description:"参数名称"`
	ParamValue string      `json:"paramValue" description:"参数值"`
	CreateTime *gtime.Time `json:"createTime" description:""`
	UpdateTime *gtime.Time `json:"updateTime" description:""`
	IsDelete   bool        `json:"isDelete"   description:"是否删除：true=已删除，false=正常"`
}
