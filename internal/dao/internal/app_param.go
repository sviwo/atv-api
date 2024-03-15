// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-03-15 14:53:40
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AppParamDao is the data access object for table sw_app_param.
type AppParamDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns AppParamColumns // columns contains all the column names of Table for convenient usage.
}

// AppParamColumns defines and stores column names for table sw_app_param.
type AppParamColumns struct {
	ParamId    string //
	ParentId   string //
	ParamName  string // 参数名称
	ParamValue string // 参数值
	CreateTime string //
	UpdateTime string //
	IsDelete   string // 是否删除：true=已删除，false=正常
}

// appParamColumns holds the columns for table sw_app_param.
var appParamColumns = AppParamColumns{
	ParamId:    "param_id",
	ParentId:   "parent_id",
	ParamName:  "param_name",
	ParamValue: "param_value",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	IsDelete:   "is_delete",
}

// NewAppParamDao creates and returns a new DAO object for table data access.
func NewAppParamDao() *AppParamDao {
	return &AppParamDao{
		group:   "default",
		table:   "sw_app_param",
		columns: appParamColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AppParamDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AppParamDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AppParamDao) Columns() AppParamColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AppParamDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AppParamDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AppParamDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
