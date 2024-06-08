// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-06-08 20:37:43
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AppMedia is the golang structure for table app_media.
type AppMedia struct {
	Id          int64       `json:"id"          orm:"id"           description:""`
	ParentId    int64       `json:"parentId"    orm:"parent_id"    description:""`
	Enable      bool        `json:"enable"      orm:"enable"       description:"显示或屏蔽：true=显示，false=屏蔽"`
	PageType    int         `json:"pageType"    orm:"page_type"    description:"页面类型：0=帮助页，1=注册用户页，2=服务页"`
	DisplayType int         `json:"displayType" orm:"display_type" description:"显示类型：0=直接显示，1=跳转外链"`
	Title       string      `json:"title"       orm:"title"        description:"标题"`
	MediaDesc   string      `json:"mediaDesc"   orm:"media_desc"   description:"简介"`
	SmallImg    string      `json:"smallImg"    orm:"small_img"    description:"缩略图"`
	Icon        string      `json:"icon"        orm:"icon"         description:""`
	Content     string      `json:"content"     orm:"content"      description:"内容"`
	Orders      int         `json:"orders"      orm:"orders"       description:"排序"`
	CreateTime  *gtime.Time `json:"createTime"  orm:"create_time"  description:""`
	UpdateTime  *gtime.Time `json:"updateTime"  orm:"update_time"  description:""`
	IsDelete    bool        `json:"isDelete"    orm:"is_delete"    description:"是否删除：true=已删除，false=正常"`
}
