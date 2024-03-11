// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-03-11 13:14:30
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AppVideos is the golang structure for table app_videos.
type AppVideos struct {
	Id          int64       `json:"id"          description:""`
	ParentId    int64       `json:"parentId"    description:""`
	Enable      bool        `json:"enable"      description:"显示或屏蔽：true=显示，false=屏蔽"`
	VideosTitle string      `json:"videosTitle" description:"视频标题"`
	SmallImg    string      `json:"smallImg"    description:"缩略图"`
	VideosDesc  string      `json:"videosDesc"  description:"视频简介"`
	VideosUrl   string      `json:"videosUrl"   description:"视频链接"`
	Orders      int         `json:"orders"      description:"排序"`
	CreateTime  *gtime.Time `json:"createTime"  description:""`
	UpdateTime  *gtime.Time `json:"updateTime"  description:""`
	IsDelete    bool        `json:"isDelete"    description:"是否删除：true=已删除，false=正常"`
}
