// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-01-16 21:22:22
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AppLogDao is the data access object for table sw_app_log.
type AppLogDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns AppLogColumns // columns contains all the column names of Table for convenient usage.
}

// AppLogColumns defines and stores column names for table sw_app_log.
type AppLogColumns struct {
	AppLogId   string //
	UserId     string //
	BusinessId string // 业务id，与类型字段组合查询
	LogType    string // 操作类型：0=其他，1=车辆，2=账户
	LogContent string // 操作内容
	CreateTime string //
	UpdateTime string //
	IsDelete   string // 是否删除：true=已删除，false=正常
}

// appLogColumns holds the columns for table sw_app_log.
var appLogColumns = AppLogColumns{
	AppLogId:   "app_log_id",
	UserId:     "user_id",
	BusinessId: "business_id",
	LogType:    "log_type",
	LogContent: "log_content",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	IsDelete:   "is_delete",
}

// NewAppLogDao creates and returns a new DAO object for table data access.
func NewAppLogDao() *AppLogDao {
	return &AppLogDao{
		group:   "default",
		table:   "sw_app_log",
		columns: appLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AppLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AppLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AppLogDao) Columns() AppLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AppLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AppLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AppLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
