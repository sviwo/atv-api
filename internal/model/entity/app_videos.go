// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-04-01 16:15:05
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AppVideos is the golang structure for table app_videos.
type AppVideos struct {
	Id          int64       `json:"id"          dc:""`
	ParentId    int64       `json:"parentId"    dc:""`
	Enable      bool        `json:"enable"      dc:"显示或屏蔽：true=显示，false=屏蔽"`
	VideosTitle string      `json:"videosTitle" dc:"视频标题"`
	SmallImg    string      `json:"smallImg"    dc:"缩略图"`
	VideosDesc  string      `json:"videosDesc"  dc:"视频简介"`
	VideosUrl   string      `json:"videosUrl"   dc:"视频链接"`
	Orders      int         `json:"orders"      dc:"排序"`
	CreateTime  *gtime.Time `json:"createTime"  dc:""`
	UpdateTime  *gtime.Time `json:"updateTime"  dc:""`
	IsDelete    bool        `json:"isDelete"    dc:"是否删除：true=已删除，false=正常"`
}
