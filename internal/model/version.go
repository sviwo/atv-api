package model

type VersionOutput struct {
	VersionId         int64  `json:"versionId"         description:""`
	VersionCode       string `json:"versionCode"       description:"版本编码"`
	VersionType       int    `json:"versionType"       description:"版本类型：0=APP更新，1=固件升级"`
	VersionUpdateType int    `json:"versionUpdateType" description:"版本更新类型：0=弱更新，1=强更新"`
	VersionUrl        string `json:"versionUrl"        description:"版本链接"`
	VersionDesc       string `json:"versionDesc"       description:"版本描述，用于app显示的新版本信息"`
}
