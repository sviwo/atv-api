// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-05-09 14:07:28
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AppMedia is the golang structure for table app_media.
type AppMedia struct {
	Id          int64       `json:"id"          description:""`
	ParentId    int64       `json:"parentId"    description:""`
	Enable      bool        `json:"enable"      description:"显示或屏蔽：true=显示，false=屏蔽"`
	PageType    int         `json:"pageType"    description:"页面类型：0=帮助页，1=注册用户页，2=视屏教程页"`
	DisplayType int         `json:"displayType" description:"显示类型：0=直接显示，1=跳转外链"`
	Title       string      `json:"title"       description:"标题"`
	MediaDesc   string      `json:"mediaDesc"   description:"简介"`
	SmallImg    string      `json:"smallImg"    description:"缩略图"`
	Content     string      `json:"content"     description:"内容"`
	Orders      int         `json:"orders"      description:"排序"`
	CreateTime  *gtime.Time `json:"createTime"  description:""`
	UpdateTime  *gtime.Time `json:"updateTime"  description:""`
	IsDelete    bool        `json:"isDelete"    description:"是否删除：true=已删除，false=正常"`
}
