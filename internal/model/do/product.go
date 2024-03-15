// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-03-15 14:53:40
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Product is the golang structure of table sw_product for DAO operations like Where/Data.
type Product struct {
	g.Meta       `orm:"table:sw_product, do:true"`
	ProductId    interface{} //
	ProductName  interface{} // 产品名称
	ProductModel interface{} // 产品型号
	CreateTime   *gtime.Time //
	UpdateTime   *gtime.Time //
	IsDelete     interface{} // 是否删除：true=已删除，false=正常
}
