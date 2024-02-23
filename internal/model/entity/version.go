// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-02-23 10:50:21
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Version is the golang structure for table version.
type Version struct {
	VersionId         int64       `json:"versionId"         description:""`
	VersionNumber     int         `json:"versionNumber"     description:"版本号，用于app版本比较判断"`
	VersionCode       string      `json:"versionCode"       description:"版本编码：app显示当前版本号使用，例如：V1.1.1"`
	VersionType       int         `json:"versionType"       description:"版本类型：0=APP更新，1=固件升级"`
	VersionUpdateType int         `json:"versionUpdateType" description:"版本更新类型：0=弱更新，1=强更新"`
	VersionStatus     int         `json:"versionStatus"     description:"版本发布状态：0=待发布，1=已发布，2=已过期"`
	VersionUrl        string      `json:"versionUrl"        description:"版本链接"`
	VersionDesc       string      `json:"versionDesc"       description:"版本描述，用于app显示的新版本信息"`
	CreateTime        *gtime.Time `json:"createTime"        description:""`
	UpdateTime        *gtime.Time `json:"updateTime"        description:""`
	IsDelete          bool        `json:"isDelete"          description:"是否删除：true=已删除，false=正常"`
}
