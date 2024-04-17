// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-04-01 16:15:05
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AppText is the golang structure for table app_text.
type AppText struct {
	Id          int64       `json:"id"          dc:""`
	ParentId    int64       `json:"parentId"    dc:""`
	Enable      bool        `json:"enable"      dc:"显示或屏蔽：true=显示，false=屏蔽"`
	TextTitle   string      `json:"textTitle"   dc:"文本标题"`
	TextContent string      `json:"textContent" dc:"文本内容"`
	Orders      int         `json:"orders"      dc:"排序"`
	CreateTime  *gtime.Time `json:"createTime"  dc:""`
	UpdateTime  *gtime.Time `json:"updateTime"  dc:""`
	IsDelete    bool        `json:"isDelete"    dc:"是否删除：true=已删除，false=正常"`
}
