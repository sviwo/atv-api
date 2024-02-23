// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-02-23 10:50:21
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AppVideos is the golang structure of table sw_app_videos for DAO operations like Where/Data.
type AppVideos struct {
	g.Meta     `orm:"table:sw_app_videos, do:true"`
	VideosId   interface{} //
	VideosName interface{} // 视频名称
	VideosDesc interface{} // 视频简介
	VideosType interface{} // 视频类型：0=充电，1=车辆驾驶，2=控制和设置，3=锁定和解锁，4=账户
	VideosUrl  interface{} // 视频链接
	CreateTime *gtime.Time //
	UpdateTime *gtime.Time //
	IsDelete   interface{} // 是否删除：true=已删除，false=正常
}
