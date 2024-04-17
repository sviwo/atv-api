package model

type AppTextTreeOutput struct {
	Id        int64                `json:"id"      dc:""`
	ParentId  int64                `json:"parentId"    dc:""`
	TextTitle string               `json:"textTitle"   dc:"文本标题"`
	Children  []*AppTextTreeOutput `json:"children"    dc:""`
}

type AppTextDetailOutput struct {
	TextContent string `json:"textContent" dc:"文本内容"`
}
