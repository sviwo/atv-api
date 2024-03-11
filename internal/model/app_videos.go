package model

type AppVideosTreeOutput struct {
	Id          string                 `json:"id"    dc:""`
	ParentId    string                 `json:"parentId"    dc:""`
	VideosTitle string                 `json:"videosTitle" description:"视频标题"`
	SmallImg    string                 `json:"smallImg"    description:"缩略图"`
	VideosDesc  string                 `json:"videosDesc"  description:"视频简介"`
	VideosUrl   string                 `json:"videosUrl"   description:"视频链接"`
	Children    []*AppVideosTreeOutput `json:"children"    dc:""`
}
