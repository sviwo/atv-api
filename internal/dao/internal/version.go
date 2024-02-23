// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-02-23 10:50:21
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// VersionDao is the data access object for table sw_version.
type VersionDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns VersionColumns // columns contains all the column names of Table for convenient usage.
}

// VersionColumns defines and stores column names for table sw_version.
type VersionColumns struct {
	VersionId         string //
	VersionNumber     string // 版本号，用于app版本比较判断
	VersionCode       string // 版本编码：app显示当前版本号使用，例如：V1.1.1
	VersionType       string // 版本类型：0=APP更新，1=固件升级
	VersionUpdateType string // 版本更新类型：0=弱更新，1=强更新
	VersionStatus     string // 版本发布状态：0=待发布，1=已发布，2=已过期
	VersionUrl        string // 版本链接
	VersionDesc       string // 版本描述，用于app显示的新版本信息
	CreateTime        string //
	UpdateTime        string //
	IsDelete          string // 是否删除：true=已删除，false=正常
}

// versionColumns holds the columns for table sw_version.
var versionColumns = VersionColumns{
	VersionId:         "version_id",
	VersionNumber:     "version_number",
	VersionCode:       "version_code",
	VersionType:       "version_type",
	VersionUpdateType: "version_update_type",
	VersionStatus:     "version_status",
	VersionUrl:        "version_url",
	VersionDesc:       "version_desc",
	CreateTime:        "create_time",
	UpdateTime:        "update_time",
	IsDelete:          "is_delete",
}

// NewVersionDao creates and returns a new DAO object for table data access.
func NewVersionDao() *VersionDao {
	return &VersionDao{
		group:   "default",
		table:   "sw_version",
		columns: versionColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *VersionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *VersionDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *VersionDao) Columns() VersionColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *VersionDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *VersionDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *VersionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
