// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-06-08 20:37:43
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Product is the golang structure of table sw_product for DAO operations like Where/Data.
type Product struct {
	g.Meta        `orm:"table:sw_product, do:true"`
	ProductId     interface{} //
	ProductName   interface{} // 产品名称
	ProductKey    interface{} // 产品key
	ProductModel  interface{} // 产品型号
	Status        interface{} // 发布状态：0=未发布，1=已发布
	Metadata      interface{} // 物模型
	MetadataTable interface{} // 是否生成物模型表：0=否，1=是
	CreateTime    *gtime.Time //
	UpdateTime    *gtime.Time //
	IsDelete      interface{} // 是否删除：true=已删除，false=正常
}
