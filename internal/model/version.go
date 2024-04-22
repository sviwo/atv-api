package model

type VersionInput struct {
	CommonPaginationInput
}

type VersionOutput struct {
	VersionId         int64  `json:"versionId"         dc:""`
	VersionCode       string `json:"versionCode"       dc:"版本编码"`
	VersionType       int    `json:"versionType"       dc:"版本类型：0=APP更新，1=固件升级"`
	VersionUpdateType int    `json:"versionUpdateType" dc:"版本更新类型：0=弱更新，1=强更新"`
	VersionUrl        string `json:"versionUrl"        dc:"版本链接"`
	VersionDesc       string `json:"versionDesc"       dc:"版本描述，用于app显示的新版本信息"`
}
