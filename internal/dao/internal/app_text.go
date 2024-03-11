// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-03-11 13:14:36
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AppTextDao is the data access object for table sw_app_text.
type AppTextDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns AppTextColumns // columns contains all the column names of Table for convenient usage.
}

// AppTextColumns defines and stores column names for table sw_app_text.
type AppTextColumns struct {
	Id          string //
	ParentId    string //
	Enable      string // 显示或屏蔽：true=显示，false=屏蔽
	TextTitle   string // 文本标题
	TextContent string // 文本内容
	Orders      string // 排序
	CreateTime  string //
	UpdateTime  string //
	IsDelete    string // 是否删除：true=已删除，false=正常
}

// appTextColumns holds the columns for table sw_app_text.
var appTextColumns = AppTextColumns{
	Id:          "id",
	ParentId:    "parent_id",
	Enable:      "enable",
	TextTitle:   "text_title",
	TextContent: "text_content",
	Orders:      "orders",
	CreateTime:  "create_time",
	UpdateTime:  "update_time",
	IsDelete:    "is_delete",
}

// NewAppTextDao creates and returns a new DAO object for table data access.
func NewAppTextDao() *AppTextDao {
	return &AppTextDao{
		group:   "default",
		table:   "sw_app_text",
		columns: appTextColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AppTextDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AppTextDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AppTextDao) Columns() AppTextColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AppTextDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AppTextDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AppTextDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
