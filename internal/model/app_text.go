package model

type AppTextListOutput struct {
	TextId    int64                `json:"textId"      description:""`
	ParentId  int64                `json:"parentId"    description:""`
	TextTitle string               `json:"textTitle"   description:"文本标题"`
	Children  []*AppTextListOutput `json:"children"    description:""`
}

type AppTextDetailOutput struct {
	TextContent string `json:"textContent" description:"文本内容"`
}
