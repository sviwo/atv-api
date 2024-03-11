package model

type AppTextTreeOutput struct {
	Id        int64                `json:"id"      description:""`
	ParentId  int64                `json:"parentId"    description:""`
	TextTitle string               `json:"textTitle"   description:"文本标题"`
	Children  []*AppTextTreeOutput `json:"children"    description:""`
}

type AppTextDetailOutput struct {
	TextContent string `json:"textContent" description:"文本内容"`
}
