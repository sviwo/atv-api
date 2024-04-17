package model

type AppVideosTreeOutput struct {
	Id          string                 `json:"id"    dc:""`
	ParentId    string                 `json:"parentId"    dc:""`
	VideosTitle string                 `json:"videosTitle" dc:"视频标题"`
	SmallImg    string                 `json:"smallImg"    dc:"缩略图"`
	VideosDesc  string                 `json:"videosDesc"  dc:"视频简介"`
	VideosUrl   string                 `json:"videosUrl"   dc:"视频链接"`
	Children    []*AppVideosTreeOutput `json:"children"    dc:""`
}
