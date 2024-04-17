// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-04-07 19:40:01
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Version is the golang structure for table version.
type Version struct {
	VersionId         int64       `json:"versionId"         dc:""`
	VersionCode       string      `json:"versionCode"       dc:"版本编码"`
	VersionType       int         `json:"versionType"       dc:"版本类型：0=APP更新，1=固件升级"`
	VersionUpdateType int         `json:"versionUpdateType" dc:"版本更新类型：0=弱更新，1=强更新"`
	VersionStatus     int         `json:"versionStatus"     dc:"版本发布状态：0=待发布，1=已发布，2=已过期"`
	VersionUrl        string      `json:"versionUrl"        dc:"版本链接"`
	VersionDesc       string      `json:"versionDesc"       dc:"版本描述，用于app显示的新版本信息"`
	CreateTime        *gtime.Time `json:"createTime"        dc:""`
	UpdateTime        *gtime.Time `json:"updateTime"        dc:""`
	IsDelete          bool        `json:"isDelete"          dc:"是否删除：true=已删除，false=正常"`
}
