// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-02-23 10:50:21
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BatteryDao is the data access object for table sw_battery.
type BatteryDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns BatteryColumns // columns contains all the column names of Table for convenient usage.
}

// BatteryColumns defines and stores column names for table sw_battery.
type BatteryColumns struct {
	BatteryId   string //
	CarId       string // 车辆id
	BatteryCode string // 电池编号
	BatteryTemp string // 电池温度
	CreateTime  string //
	UpdateTime  string //
	IsDelete    string // 是否删除：true=已删除，false=正常
}

// batteryColumns holds the columns for table sw_battery.
var batteryColumns = BatteryColumns{
	BatteryId:   "battery_id",
	CarId:       "car_id",
	BatteryCode: "battery_code",
	BatteryTemp: "battery_temp",
	CreateTime:  "create_time",
	UpdateTime:  "update_time",
	IsDelete:    "is_delete",
}

// NewBatteryDao creates and returns a new DAO object for table data access.
func NewBatteryDao() *BatteryDao {
	return &BatteryDao{
		group:   "default",
		table:   "sw_battery",
		columns: batteryColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *BatteryDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *BatteryDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *BatteryDao) Columns() BatteryColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *BatteryDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *BatteryDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *BatteryDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
