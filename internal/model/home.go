package model

type QueryHomeDataInput struct {
	VersionUpdateType int    `json:"versionUpdateType" description:"版本更新类型：0=弱更新，1=强更新"`
	VersionUrl        string `json:"versionUrl"        description:"版本链接"`
	VersionDesc       string `json:"versionDesc"       description:"版本描述，用于app显示的新版本信息"`
}

type HomeDataOutput struct {
	VersionUpdateType int    `json:"versionUpdateType" description:"版本更新类型：0=弱更新，1=强更新"`
	VersionUrl        string `json:"versionUrl"        description:"版本链接"`
	VersionDesc       string `json:"versionDesc"       description:"版本描述，用于app显示的新版本信息"`
}
