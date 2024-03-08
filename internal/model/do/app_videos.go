// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-03-08 08:43:11
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AppVideos is the golang structure of table sw_app_videos for DAO operations like Where/Data.
type AppVideos struct {
	g.Meta      `orm:"table:sw_app_videos, do:true"`
	VideosId    interface{} //
	ParentId    interface{} //
	Enable      interface{} // 显示或屏蔽：true=显示，false=屏蔽
	VideosTitle interface{} // 视频标题
	SmallImg    interface{} // 缩略图
	VideosDesc  interface{} // 视频简介
	VideosUrl   interface{} // 视频链接
	Orders      interface{} // 排序
	CreateTime  *gtime.Time //
	UpdateTime  *gtime.Time //
	IsDelete    interface{} // 是否删除：true=已删除，false=正常
}
