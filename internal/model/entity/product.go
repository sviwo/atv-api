// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-06-08 20:37:43
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Product is the golang structure for table product.
type Product struct {
	g.Meta        `orm:"table:sw_product"`
	ProductId     int64       `json:"productId"     orm:"product_id"     description:""`
	ProductName   string      `json:"productName"   orm:"product_name"   description:"产品名称"`
	ProductKey    string      `json:"productKey"    orm:"product_key"    description:"产品key"`
	ProductModel  string      `json:"productModel"  orm:"product_model"  description:"产品型号"`
	Status        int         `json:"status"        orm:"status"         description:"发布状态：0=未发布，1=已发布"`
	Metadata      string      `json:"metadata"      orm:"metadata"       description:"物模型"`
	MetadataTable int         `json:"metadataTable" orm:"metadata_table" description:"是否生成物模型表：0=否，1=是"`
	CreateTime    *gtime.Time `json:"createTime"    orm:"create_time"    description:""`
	UpdateTime    *gtime.Time `json:"updateTime"    orm:"update_time"    description:""`
	IsDelete      bool        `json:"isDelete"      orm:"is_delete"      description:"是否删除：true=已删除，false=正常"`
}
