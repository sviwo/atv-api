package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type AppMediaTreeReq struct {
	g.Meta   `path:"/app/media/get/tree" method:"get" tags:"APP媒体相关" sm:"获取APP媒体标题树"`
	PageType int `json:"pageType"    v:"required"        dc:"页面类型：0=帮助页，1=注册用户页，2=视屏教程页"`
}

type AppMediaTreeRes struct {
	Id          string             `json:"id"          description:""`
	ParentId    string             `json:"parentId"    description:""`
	DisplayType int                `json:"displayType" description:"显示类型：0=直接显示，1=跳转外链"`
	Title       string             `json:"title"       description:"标题"`
	MediaDesc   string             `json:"mediaDesc"   description:"简介"`
	SmallImg    string             `json:"smallImg"    description:"缩略图"`
	Content     string             `json:"content"     description:"内容（当显示类型为0的时候content为空，为1的时content为链接地址）"`
	Children    []*AppMediaTreeRes `json:"children"    dc:""`
}

type AppMediaDetailReq struct {
	g.Meta `path:"/app/media/get/detail" method:"get" tags:"APP媒体相关" sm:"获取APP文本内容"`
	Id     int64 `json:"id"            dc:""               v:"required"`
}

type AppMediaDetailRes struct {
	Content string `json:"content" dc:"文本内容"`
}
