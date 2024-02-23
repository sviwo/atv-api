// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-02-23 10:50:21
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AppText is the golang structure for table app_text.
type AppText struct {
	TextId      int64       `json:"textId"      description:""`
	TextName    string      `json:"textName"    description:"文本名称"`
	TextType    int         `json:"textType"    description:"文本类型：0=账户，1=车辆功能和充电，2=售后服务和维修保养，3=其他"`
	TextContent string      `json:"textContent" description:"文本内容"`
	CreateTime  *gtime.Time `json:"createTime"  description:""`
	UpdateTime  *gtime.Time `json:"updateTime"  description:""`
	IsDelete    bool        `json:"isDelete"    description:"是否删除：true=已删除，false=正常"`
}
