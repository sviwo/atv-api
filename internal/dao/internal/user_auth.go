// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-03-15 14:53:40
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserAuthDao is the data access object for table sw_user_auth.
type UserAuthDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns UserAuthColumns // columns contains all the column names of Table for convenient usage.
}

// UserAuthColumns defines and stores column names for table sw_user_auth.
type UserAuthColumns struct {
	AuthId              string //
	UserId              string //
	AuthFirstName       string //
	AuthLastName        string //
	CertificateFrontImg string // 证件正面照片
	CertificateBackImg  string // 证件背面照片
	AuthStatus          string // 认证状态：0=未认证，1=认证中，2=认证成功，3=认证失败
	AuthFailReason      string // 认证失败原因
	AuthTime            string // 认证时间
	VerifyTime          string // 审核时间
	CreateTime          string //
	UpdateTime          string //
}

// userAuthColumns holds the columns for table sw_user_auth.
var userAuthColumns = UserAuthColumns{
	AuthId:              "auth_id",
	UserId:              "user_id",
	AuthFirstName:       "auth_first_name",
	AuthLastName:        "auth_last_name",
	CertificateFrontImg: "certificate_front_img",
	CertificateBackImg:  "certificate_back_img",
	AuthStatus:          "auth_status",
	AuthFailReason:      "auth_fail_reason",
	AuthTime:            "auth_time",
	VerifyTime:          "verify_time",
	CreateTime:          "create_time",
	UpdateTime:          "update_time",
}

// NewUserAuthDao creates and returns a new DAO object for table data access.
func NewUserAuthDao() *UserAuthDao {
	return &UserAuthDao{
		group:   "default",
		table:   "sw_user_auth",
		columns: userAuthColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserAuthDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserAuthDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserAuthDao) Columns() UserAuthColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserAuthDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserAuthDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserAuthDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
