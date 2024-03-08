// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-03-07 16:22:32
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AppText is the golang structure for table app_text.
type AppText struct {
	TextId      int64       `json:"textId"      description:""`
	ParentId    int64       `json:"parentId"    description:""`
	TextTitle   string      `json:"textTitle"   description:"文本标题"`
	TextContent string      `json:"textContent" description:"文本内容"`
	CreateTime  *gtime.Time `json:"createTime"  description:""`
	UpdateTime  *gtime.Time `json:"updateTime"  description:""`
	IsDelete    bool        `json:"isDelete"    description:"是否删除：true=已删除，false=正常"`
}
