// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-03-15 14:53:40
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ProductDao is the data access object for table sw_product.
type ProductDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns ProductColumns // columns contains all the column names of Table for convenient usage.
}

// ProductColumns defines and stores column names for table sw_product.
type ProductColumns struct {
	ProductId     string //
	ProductName   string // 产品名称
	ProductModel  string // 产品型号
	Status        string // 发布状态：0=未发布，1=已发布
	CreateTime    string //
	UpdateTime    string //
	IsDelete      string // 是否删除：true=已删除，false=正常
	ProductKey    string
	MetaData      string
	MetadataTable string //是否生成物模型表：0=否，1=是
}

// productColumns holds the columns for table sw_product.
var productColumns = ProductColumns{
	ProductId:     "product_id",
	ProductName:   "product_name",
	ProductKey:    "product_key",
	ProductModel:  "product_model",
	Status:        "status",
	MetaData:      "metadata",
	CreateTime:    "create_time",
	UpdateTime:    "update_time",
	MetadataTable: "metadata_table",
	IsDelete:      "is_delete",
}

// NewProductDao creates and returns a new DAO object for table data access.
func NewProductDao() *ProductDao {
	return &ProductDao{
		group:   "default",
		table:   "sw_product",
		columns: productColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ProductDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ProductDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ProductDao) Columns() ProductColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ProductDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ProductDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ProductDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
