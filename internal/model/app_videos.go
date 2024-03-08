package model

type AppVideosListOutput struct {
	VideosId    int64                  `json:"videosId"    description:""`
	ParentId    int64                  `json:"parentId"    description:""`
	VideosTitle string                 `json:"videosTitle" description:"视频标题"`
	SmallImg    string                 `json:"smallImg"    description:"缩略图"`
	VideosDesc  string                 `json:"videosDesc"  description:"视频简介"`
	VideosUrl   string                 `json:"videosUrl"   description:"视频链接"`
	Children    []*AppVideosListOutput `json:"children"    description:""`
}
