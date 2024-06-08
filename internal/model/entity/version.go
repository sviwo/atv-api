// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-06-08 20:37:43
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Version is the golang structure for table version.
type Version struct {
	VersionId         int64       `json:"versionId"         orm:"version_id"          description:""`
	VersionCode       string      `json:"versionCode"       orm:"version_code"        description:"版本编码"`
	VersionType       int         `json:"versionType"       orm:"version_type"        description:"版本类型：0=APP更新，1=固件升级"`
	VersionUpdateType int         `json:"versionUpdateType" orm:"version_update_type" description:"版本更新类型：0=弱更新，1=强更新"`
	VersionStatus     int         `json:"versionStatus"     orm:"version_status"      description:"版本发布状态：0=待发布，1=已发布，2=已过期"`
	VersionUrl        string      `json:"versionUrl"        orm:"version_url"         description:"版本链接"`
	VersionDesc       string      `json:"versionDesc"       orm:"version_desc"        description:"版本描述，用于app显示的新版本信息"`
	CreateTime        *gtime.Time `json:"createTime"        orm:"create_time"         description:""`
	UpdateTime        *gtime.Time `json:"updateTime"        orm:"update_time"         description:""`
	IsDelete          bool        `json:"isDelete"          orm:"is_delete"           description:"是否删除：true=已删除，false=正常"`
}
