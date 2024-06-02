package model

type AppMediaTreeOutput struct {
	Id          int64                 `json:"id"          description:""`
	ParentId    int64                 `json:"parentId"    description:""`
	DisplayType int                   `json:"displayType" description:"显示类型：0=直接显示，1=跳转外链"`
	Title       string                `json:"title"       description:"标题"`
	MediaDesc   string                `json:"mediaDesc"   description:"简介"`
	SmallImg    string                `json:"smallImg"    description:"缩略图"`
	Icon        string                `json:"icon"        description:"icon"`
	Content     string                `json:"content"     description:"内容"`
	Children    []*AppMediaTreeOutput `json:"children"    description:""`
}
