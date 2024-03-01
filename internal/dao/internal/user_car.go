// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-03-01 17:48:00
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserCarDao is the data access object for table sw_user_car.
type UserCarDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns UserCarColumns // columns contains all the column names of Table for convenient usage.
}

// UserCarColumns defines and stores column names for table sw_user_car.
type UserCarColumns struct {
	UserCarId       string //
	CarId           string //
	UserId          string //
	IsSelect        string // 是否选定：false=未选定，true=已选定
	MobileKey       string // 手机钥匙开关：false=关，true=开
	SpeedLimit      string // 速度限制开关：false=关，true=开
	DrivingModeType string // 驾驶模式：0=ECO模式，1=运动模式，2=狂暴模式，3=脱困模式
	CreateTime      string //
}

// userCarColumns holds the columns for table sw_user_car.
var userCarColumns = UserCarColumns{
	UserCarId:       "user_car_id",
	CarId:           "car_id",
	UserId:          "user_id",
	IsSelect:        "is_select",
	MobileKey:       "mobile_key",
	SpeedLimit:      "speed_limit",
	DrivingModeType: "driving_mode_type",
	CreateTime:      "create_time",
}

// NewUserCarDao creates and returns a new DAO object for table data access.
func NewUserCarDao() *UserCarDao {
	return &UserCarDao{
		group:   "default",
		table:   "sw_user_car",
		columns: userCarColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserCarDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserCarDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserCarDao) Columns() UserCarColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserCarDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserCarDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserCarDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
