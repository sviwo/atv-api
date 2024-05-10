// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-05-10 15:47:49
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AppMedia is the golang structure of table sw_app_media for DAO operations like Where/Data.
type AppMedia struct {
	g.Meta      `orm:"table:sw_app_media, do:true"`
	Id          interface{} //
	ParentId    interface{} //
	Enable      interface{} // 显示或屏蔽：true=显示，false=屏蔽
	PageType    interface{} // 页面类型：0=帮助页，1=注册用户页，2=视屏教程页
	DisplayType interface{} // 显示类型：0=直接显示，1=跳转外链
	Title       interface{} // 标题
	MediaDesc   interface{} // 简介
	SmallImg    interface{} // 缩略图
	Content     interface{} // 内容
	Orders      interface{} // 排序
	CreateTime  *gtime.Time //
	UpdateTime  *gtime.Time //
	IsDelete    interface{} // 是否删除：true=已删除，false=正常
}
