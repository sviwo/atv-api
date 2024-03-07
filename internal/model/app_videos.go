package model

type AppVideos struct {
	VideosId   int64  `json:"videosId"   description:""`
	VideosName string `json:"videosName" description:"视频名称"`
	VideosDesc string `json:"videosDesc" description:"视频简介"`
	VideosType int    `json:"videosType" description:"视频类型：0=充电，1=车辆驾驶，2=控制和设置，3=锁定和解锁，4=账户"`
	VideosUrl  string `json:"videosUrl"  description:"视频链接"`
}
