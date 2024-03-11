package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type AppVideosTreeReq struct {
	g.Meta `path:"/app/videos/tree" method:"get" tags:"APP视屏相关" sm:"获取APP视屏树"`
}

type AppVideosTreeRes struct {
	Id          string              `json:"id"    dc:""`
	ParentId    string              `json:"parentId"    dc:""`
	VideosTitle string              `json:"videosTitle" dc:"视频标题"`
	SmallImg    string              `json:"smallImg"    dc:"缩略图"`
	VideosDesc  string              `json:"videosDesc"  dc:"视频简介"`
	VideosUrl   string              `json:"videosUrl"   dc:"视频链接"`
	Children    []*AppVideosTreeRes `json:"children"    dc:""`
}
