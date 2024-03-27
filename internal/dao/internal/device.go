// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-03-27 15:02:22
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DeviceDao is the data access object for table sw_device.
type DeviceDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns DeviceColumns // columns contains all the column names of Table for convenient usage.
}

// DeviceColumns defines and stores column names for table sw_device.
type DeviceColumns struct {
	DeviceId       string //
	ProductId      string // 所属产品
	ProductKey     string // 对应物联网平台产品的ProductKey
	DeviceName     string // 对应物联网平台颁发的设备证书的DeviceName
	DeviceSecret   string // 对应物联网平台颁发的设备证书的DeviceSecret
	DeviceModel    string // 设备型号
	Nickname       string // 产品昵称（目前只有ATV，则等同于车辆昵称）
	Status         string // 状态：0=未启用，1=离线，2=在线
	OnlineTimeout  string // 设备在线超时设置，单位：秒
	ActivateTime   string // 激活时间
	RegistryTime   string // 注册时间
	Version        string // 固件版本号
	LastOnlineTime string // 最后上线时间
	CreateTime     string //
	UpdateTime     string //
	MetadataTable  string // 是否生成物模型表：0=否，1=是
	IsDelete       string // 是否删除：true=已删除，false=正常
}

// deviceColumns holds the columns for table sw_device.
var deviceColumns = DeviceColumns{
	DeviceId:       "device_id",
	ProductId:      "product_id",
	ProductKey:     "product_key",
	DeviceName:     "device_name",
	DeviceSecret:   "device_secret",
	DeviceModel:    "device_model",
	Nickname:       "nickname",
	Status:         "status",
	OnlineTimeout:  "online_timeout",
	ActivateTime:   "activate_time",
	RegistryTime:   "registry_time",
	Version:        "version",
	LastOnlineTime: "last_online_time",
	CreateTime:     "create_time",
	UpdateTime:     "update_time",
	MetadataTable:  "metadata_table",
	IsDelete:       "is_delete",
}

// NewDeviceDao creates and returns a new DAO object for table data access.
func NewDeviceDao() *DeviceDao {
	return &DeviceDao{
		group:   "default",
		table:   "sw_device",
		columns: deviceColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *DeviceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *DeviceDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *DeviceDao) Columns() DeviceColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *DeviceDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *DeviceDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *DeviceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
