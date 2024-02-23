// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-01-16 21:22:22
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserDao is the data access object for table sw_user.
type UserDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns UserColumns // columns contains all the column names of Table for convenient usage.
}

// UserColumns defines and stores column names for table sw_user.
type UserColumns struct {
	UserId            string //
	Username          string //
	Password          string //
	PwdSalt           string // 密码盐值
	PwdEncryNum       string // 密码加密次数
	FirstName         string //
	LastName          string //
	Nickname          string // 昵称
	Enable            string // 账号是否可用：true=正常，false=停用
	HeadImg           string //
	CertificateCode   string // 身份证编号
	VipLevel          string // vip级别
	IpRegion          string // Ip属地
	MobilePhone       string // 手机号
	PersonalSignature string // 个性签名
	Birthday          string // 生日
	Gender            string // 性别：0=未知，1=男，2=女
	CreateTime        string //
	UpdateTime        string //
	IsDelete          string // 是否删除：true=已删除，false=正常
}

// userColumns holds the columns for table sw_user.
var userColumns = UserColumns{
	UserId:            "user_id",
	Username:          "username",
	Password:          "password",
	PwdSalt:           "pwd_salt",
	PwdEncryNum:       "pwd_encry_num",
	FirstName:         "first_name",
	LastName:          "last_name",
	Nickname:          "nickname",
	Enable:            "enable",
	HeadImg:           "head_img",
	CertificateCode:   "certificate_code",
	VipLevel:          "vip_level",
	IpRegion:          "ip_region",
	MobilePhone:       "mobile_phone",
	PersonalSignature: "personal_signature",
	Birthday:          "birthday",
	Gender:            "gender",
	CreateTime:        "create_time",
	UpdateTime:        "update_time",
	IsDelete:          "is_delete",
}

// NewUserDao creates and returns a new DAO object for table data access.
func NewUserDao() *UserDao {
	return &UserDao{
		group:   "default",
		table:   "sw_user",
		columns: userColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserDao) Columns() UserColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
