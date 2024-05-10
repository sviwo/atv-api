// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-05-10 15:47:49
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AppParam is the golang structure for table app_param.
type AppParam struct {
	ParamId    int64       `json:"paramId"    orm:"param_id"    description:""`
	ParentId   int64       `json:"parentId"   orm:"parent_id"   description:""`
	ParamName  string      `json:"paramName"  orm:"param_name"  description:"参数名称"`
	ParamValue string      `json:"paramValue" orm:"param_value" description:"参数值"`
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:""`
	UpdateTime *gtime.Time `json:"updateTime" orm:"update_time" description:""`
	IsDelete   bool        `json:"isDelete"   orm:"is_delete"   description:"是否删除：true=已删除，false=正常"`
}
