// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-03-15 14:53:40
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Version is the golang structure of table sw_version for DAO operations like Where/Data.
type Version struct {
	g.Meta            `orm:"table:sw_version, do:true"`
	VersionId         interface{} //
	VersionNumber     interface{} // 版本号，用于app版本比较判断
	VersionCode       interface{} // 版本编码：app显示当前版本号使用，例如：V1.1.1
	VersionType       interface{} // 版本类型：0=APP更新，1=固件升级
	VersionUpdateType interface{} // 版本更新类型：0=弱更新，1=强更新
	VersionStatus     interface{} // 版本发布状态：0=待发布，1=已发布，2=已过期
	VersionUrl        interface{} // 版本链接
	VersionDesc       interface{} // 版本描述，用于app显示的新版本信息
	CreateTime        *gtime.Time //
	UpdateTime        *gtime.Time //
	IsDelete          interface{} // 是否删除：true=已删除，false=正常
}
