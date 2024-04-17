// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-04-01 16:15:05
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// Product is the golang structure for table product.
type Product struct {
	gmeta.Meta    `orm:"table:sw_product"`
	ProductId     int64       `json:"productId"     dc:""`
	ProductName   string      `json:"productName"   dc:"产品名称"`
	ProductKey    string      `json:"productKey"    dc:"产品key"`
	ProductModel  string      `json:"productModel"  dc:"产品型号"`
	Status        int         `json:"status"        dc:"发布状态：0=未发布，1=已发布"`
	Metadata      string      `json:"metadata"      dc:"物模型"`
	CreateTime    *gtime.Time `json:"createTime"    dc:""`
	UpdateTime    *gtime.Time `json:"updateTime"    dc:""`
	IsDelete      bool        `json:"isDelete"      dc:"是否删除：true=已删除，false=正常"`
	MetadataTable int         `json:"metadataTable" dc:"是否生成物模型表：0=否，1=是"`
}
