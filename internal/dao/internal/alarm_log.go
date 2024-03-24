// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AlarmLogDao is the data access object for table alarm_log.
type AlarmLogDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns AlarmLogColumns // columns contains all the column names of Table for convenient usage.
}

// AlarmLogColumns defines and stores column names for table alarm_log.
type AlarmLogColumns struct {
	Id         string //
	Type       string // 告警类型：1=规则告警，2=设备自主告警
	Data       string // 触发告警的数据
	ProductKey string // 产品标识
	DeviceKey  string // 设备标识
	Status     string // 告警状态：0=未处理，1=已处理
	CreatedTime  string // 告警时间
	UpdatedBy  string // 告警处理人员
	UpdatedTime  string // 处理时间
	Content    string // 处理意见
}

// alarmLogColumns holds the columns for table alarm_log.
var alarmLogColumns = AlarmLogColumns{
	Id:         "id",
	Type:       "type",
	Data:       "data",
	ProductKey: "product_key",
	DeviceKey:  "device_key",
	Status:     "status",
	CreatedTime:  "created_time",
	UpdatedBy:  "updated_by",
	UpdatedTime:  "updated_time",
	Content:    "content",
}

// NewAlarmLogDao creates and returns a new DAO object for table data access.
func NewAlarmLogDao() *AlarmLogDao {
	return &AlarmLogDao{
		group:   "default",
		table:   "sw_alarm_log",
		columns: alarmLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AlarmLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AlarmLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AlarmLogDao) Columns() AlarmLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AlarmLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AlarmLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AlarmLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
