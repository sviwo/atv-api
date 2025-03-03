// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-06-08 20:37:43
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserDeviceDao is the data access object for table sw_user_device.
type UserDeviceDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns UserDeviceColumns // columns contains all the column names of Table for convenient usage.
}

// UserDeviceColumns defines and stores column names for table sw_user_device.
type UserDeviceColumns struct {
	Id              string //
	DeviceId        string //
	UserId          string //
	IsSelect        string // 是否选定：false=未选定，true=已选定
	UserDeviceType  string // 设备用户类型：0=主用户，1=从用户
	MobileKey       string // 手机钥匙开关：false=关，true=开
	SpeedLimit      string // 速度限制开关：false=关，true=开
	DrivingModeType string // 驾驶模式：0=ECO模式，1=运动模式，2=狂暴模式
	EnergyRecovery  string // 动能回收类型：0=无，1=中，2=强
	CreateTime      string //
	UpdateTime      string //
}

// userDeviceColumns holds the columns for table sw_user_device.
var userDeviceColumns = UserDeviceColumns{
	Id:              "id",
	DeviceId:        "device_id",
	UserId:          "user_id",
	IsSelect:        "is_select",
	UserDeviceType:  "user_device_type",
	MobileKey:       "mobile_key",
	SpeedLimit:      "speed_limit",
	DrivingModeType: "driving_mode_type",
	EnergyRecovery:  "energy_recovery",
	CreateTime:      "create_time",
	UpdateTime:      "update_time",
}

// NewUserDeviceDao creates and returns a new DAO object for table data access.
func NewUserDeviceDao() *UserDeviceDao {
	return &UserDeviceDao{
		group:   "default",
		table:   "sw_user_device",
		columns: userDeviceColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserDeviceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserDeviceDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserDeviceDao) Columns() UserDeviceColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserDeviceDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserDeviceDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserDeviceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
