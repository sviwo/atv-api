// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-02-23 10:50:21
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TravelRecordDao is the data access object for table sw_travel_record.
type TravelRecordDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns TravelRecordColumns // columns contains all the column names of Table for convenient usage.
}

// TravelRecordColumns defines and stores column names for table sw_travel_record.
type TravelRecordColumns struct {
	TravelRecordId string //
	StartPoint     string // 起点
	EndPoint       string // 终点
	MileageDriven  string // 行驶里程，单位（m）
	StartTime      string // 行程开始时间
	EndTime        string // 行程结束时间
	AvgSpeed       string // 平均时速，单位（m）
	Consumption    string // 使用电量
	CreateTime     string //
	UpdateTime     string //
	IsDelete       string // 是否删除：true=已删除，false=正常
}

// travelRecordColumns holds the columns for table sw_travel_record.
var travelRecordColumns = TravelRecordColumns{
	TravelRecordId: "travel_record_id",
	StartPoint:     "start_point",
	EndPoint:       "end_point",
	MileageDriven:  "mileage_driven",
	StartTime:      "start_time",
	EndTime:        "end_time",
	AvgSpeed:       "avg_speed",
	Consumption:    "consumption",
	CreateTime:     "create_time",
	UpdateTime:     "update_time",
	IsDelete:       "is_delete",
}

// NewTravelRecordDao creates and returns a new DAO object for table data access.
func NewTravelRecordDao() *TravelRecordDao {
	return &TravelRecordDao{
		group:   "default",
		table:   "sw_travel_record",
		columns: travelRecordColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TravelRecordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TravelRecordDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TravelRecordDao) Columns() TravelRecordColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TravelRecordDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TravelRecordDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TravelRecordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
