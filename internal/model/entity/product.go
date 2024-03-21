// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-03-15 14:53:40
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// Product is the golang structure for table product.
type Product struct {
	gmeta.Meta    `orm:"table:sw_product"`
	ProductId     int64       `json:"productId"    description:""`
	ProductName   string      `json:"productName"  description:"产品名称"`
	ProductKey    string      `json:"productKey"  description:"产品key"`
	ProductModel  string      `json:"productModel" description:"产品型号"`
	Status        int         `json:"status"         description:"状态：发布状态：0=未发布，1=已发布"`
	CreateTime    *gtime.Time `json:"createTime"   description:""`
	UpdateTime    *gtime.Time `json:"updateTime"   description:""`
	IsDelete      bool        `json:"isDelete"     description:"是否删除：true=已删除，false=正常"`
	Metadata      string      `json:"metadata"          description:"物模型"`
	MetadataTable string      `json:"metadata_table"    description:"是否生成物模型子表：0=否，1=是"`
}
