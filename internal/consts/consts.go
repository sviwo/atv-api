package consts

/*
代码业务常量管理
*/
const (
	ProjectName              = "sviwo"
	ProjectUsage             = "sviwo app server"
	ProjectBrief             = "start http server"
	Version                  = "v1.0.0"           // 当前服务版本(用于模板展示)
	CaptchaName              = "UserVftCode"      // 验证码存储空间名称
	ContextKey               = "ContextKey"       // 上下文变量存储键名，前后端系统共享
	ContextKeyUserId         = "ContextKeyUserId" // 上下文变量存储键名，前后端系统共享
	CarKeyPrefix             = "SVIWO_"           // 车钥匙邀请码前缀
	FileMaxUploadCountMinute = 10                 // 同一用户1分钟之内最大上传数量

	//登陆类型：1=密码，2=第三方
	LoginTypePwd   = 1
	LoginTypeThird = 2

	//认证状态：0=未认证，1=认证中，2=认证成功，3=认证失败
	UserAuthStatusNot     = 0
	UserAuthStatusIn      = 1
	UserAuthStatusSuccess = 2
	UserAuthStatusFail    = 3

	//删除状态：true(1)=已删除，false(0)=正常
	DeleteYes = true
	DeleteOn  = false

	//车辆是否选定：false=未选定，true=已选定
	CarSelectYes = true
	CarSelectNo  = false

	//手机钥匙开关：false=关，true=开
	CarMobileKeyYes = true
	CarMobileKeyNo  = false

	//速度限制开关：false=关，true=开
	CarSpeedLimitYes = true
	CarSpeedLimitNo  = false

	//驾驶模式：0=ECO模式，1=运动模式，2=狂暴模式
	drivingModeECO    = 0
	drivingModeSports = 1
	drivingModeRage   = 2

	//动能回收类型：0=无，1=中，2=强
	energyRecoveryNot    = 0
	energyRecoveryMedium = 1
	energyRecoveryStrong = 2

	//显示或屏蔽：true=显示，false=屏蔽
	EnableDisplay = true
	EnableShield  = false

	//是否生成物模型表或子表：1，0 否
	MetadataTable   = 1
	DeMetadataTable = 0

	//设备用户类型：0=主用户，1=从用户
	UserDeviceTypeMain = 0
	UserDeviceChild    = 1

	//版本发布状态：0=待发布，1=已发布，2=已过期
	VersionReleaseStatusWait    = 0
	VersionReleaseStatusYes     = 1
	VersionReleaseStatusExpired = 2

	//版本类型：0=APP更新，1=固件升级
	VersionTypeApp      = 0
	VersionTypeFirmware = 1

	//页面类型：0=帮助页，1=注册用户页，2=服务页
	PageTypeHelp         = 0
	PageTypeRegisterUser = 1
	PageTypeServer       = 2

	//显示类型：0=直接显示，1=跳转外链
	DisplayTypeDirect   = 0
	DisplayTypeRedirect = 1
)
