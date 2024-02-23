// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-02-23 10:50:21
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CarDao is the data access object for table sw_car.
type CarDao struct {
	table   string     // table is the underlying table name of the DAO.
	group   string     // group is the database configuration group name of current DAO.
	columns CarColumns // columns contains all the column names of Table for convenient usage.
}

// CarColumns defines and stores column names for table sw_car.
type CarColumns struct {
	CarId              string //
	CarNickname        string // 车辆昵称（用户自定义）
	CarFinalPosition   string // 车辆最后定位
	TravelKm           string // 行驶公里数
	CarFrameCode       string // 车架编号
	CarMotorCode       string // 车辆电机编号
	ResidueElectricity string // 剩余电量
	ResidueKm          string // 剩余公里数
	AfterSalesTime     string // 保修日期
	ActivationTime     string // 激活时间
	CreateTime         string //
	UpdateTime         string //
	IsDelete           string // 是否删除：true=已删除，false=正常
}

// carColumns holds the columns for table sw_car.
var carColumns = CarColumns{
	CarId:              "car_id",
	CarNickname:        "car_nickname",
	CarFinalPosition:   "car_final_position",
	TravelKm:           "travel_km",
	CarFrameCode:       "car_frame_code",
	CarMotorCode:       "car_motor_code",
	ResidueElectricity: "residue_electricity",
	ResidueKm:          "residue_km",
	AfterSalesTime:     "after_sales_time",
	ActivationTime:     "activation_time",
	CreateTime:         "create_time",
	UpdateTime:         "update_time",
	IsDelete:           "is_delete",
}

// NewCarDao creates and returns a new DAO object for table data access.
func NewCarDao() *CarDao {
	return &CarDao{
		group:   "default",
		table:   "sw_car",
		columns: carColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CarDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CarDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CarDao) Columns() CarColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CarDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CarDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CarDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
