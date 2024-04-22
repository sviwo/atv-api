package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type VersionRes struct {
	VersionId         string `json:"versionId"         dc:""`
	VersionCode       string `json:"versionCode"       dc:"版本编码"`
	VersionType       int    `json:"versionType"       dc:"版本类型：0=APP更新，1=固件升级"`
	VersionUpdateType int    `json:"versionUpdateType" dc:"版本更新类型：0=弱更新，1=强更新"`
	VersionUrl        string `json:"versionUrl"        dc:"版本链接"`
	VersionDesc       string `json:"versionDesc"       dc:"版本描述，用于app显示的新版本信息"`
}

type NewVersionReq struct {
	g.Meta `path:"/version/info" method:"get" tags:"版本相关" sm:"获取最新版本信息"`
}

type VersionListReq struct {
	g.Meta `path:"/version/list" method:"get" tags:"版本相关" sm:"获取发行说明"`
	CommonPaginationReq
}

type VersionListRes struct {
	Result []*VersionRes `json:"result" dc:"列表数据"`
	CommonPaginationRes
}
