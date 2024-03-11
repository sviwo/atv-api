package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type AppTextTreeReq struct {
	g.Meta `path:"/app/text/tree" method:"get" tags:"APP文本相关" sm:"获取APP文本标题树"`
}

type AppTextTreeRes struct {
	Id        string            `json:"id"      dc:""`
	ParentId  string            `json:"parentId"    dc:""`
	TextTitle string            `json:"textTitle"   dc:"文本标题"`
	Children  []*AppTextTreeRes `json:"children"    dc:""`
}

type AppTextDetailReq struct {
	g.Meta `path:"/app/text/detail" method:"get" tags:"APP文本相关" sm:"获取APP文本内容"`
	TextId int64 `json:"textId"            dc:""               v:"required"`
}

type AppTextDetailRes struct {
	TextContent string `json:"textContent" dc:"文本内容"`
}
