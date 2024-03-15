// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-03-15 14:53:40
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AppVideosDao is the data access object for table sw_app_videos.
type AppVideosDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns AppVideosColumns // columns contains all the column names of Table for convenient usage.
}

// AppVideosColumns defines and stores column names for table sw_app_videos.
type AppVideosColumns struct {
	Id          string //
	ParentId    string //
	Enable      string // 显示或屏蔽：true=显示，false=屏蔽
	VideosTitle string // 视频标题
	SmallImg    string // 缩略图
	VideosDesc  string // 视频简介
	VideosUrl   string // 视频链接
	Orders      string // 排序
	CreateTime  string //
	UpdateTime  string //
	IsDelete    string // 是否删除：true=已删除，false=正常
}

// appVideosColumns holds the columns for table sw_app_videos.
var appVideosColumns = AppVideosColumns{
	Id:          "id",
	ParentId:    "parent_id",
	Enable:      "enable",
	VideosTitle: "videos_title",
	SmallImg:    "small_img",
	VideosDesc:  "videos_desc",
	VideosUrl:   "videos_url",
	Orders:      "orders",
	CreateTime:  "create_time",
	UpdateTime:  "update_time",
	IsDelete:    "is_delete",
}

// NewAppVideosDao creates and returns a new DAO object for table data access.
func NewAppVideosDao() *AppVideosDao {
	return &AppVideosDao{
		group:   "default",
		table:   "sw_app_videos",
		columns: appVideosColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AppVideosDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AppVideosDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AppVideosDao) Columns() AppVideosColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AppVideosDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AppVideosDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AppVideosDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
