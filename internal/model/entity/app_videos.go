// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-02-23 10:50:21
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AppVideos is the golang structure for table app_videos.
type AppVideos struct {
	VideosId   int64       `json:"videosId"   description:""`
	VideosName string      `json:"videosName" description:"视频名称"`
	VideosDesc string      `json:"videosDesc" description:"视频简介"`
	VideosType int         `json:"videosType" description:"视频类型：0=充电，1=车辆驾驶，2=控制和设置，3=锁定和解锁，4=账户"`
	VideosUrl  string      `json:"videosUrl"  description:"视频链接"`
	CreateTime *gtime.Time `json:"createTime" description:""`
	UpdateTime *gtime.Time `json:"updateTime" description:""`
	IsDelete   bool        `json:"isDelete"   description:"是否删除：true=已删除，false=正常"`
}
