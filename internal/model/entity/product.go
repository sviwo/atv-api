// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-03-15 14:53:40
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Product is the golang structure for table product.
type Product struct {
	ProductId    int64       `json:"productId"    description:""`
	ProductName  string      `json:"productName"  description:"产品名称"`
	ProductModel string      `json:"productModel" description:"产品型号"`
	CreateTime   *gtime.Time `json:"createTime"   description:""`
	UpdateTime   *gtime.Time `json:"updateTime"   description:""`
	IsDelete     bool        `json:"isDelete"     description:"是否删除：true=已删除，false=正常"`
}
