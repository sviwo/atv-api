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
	FileMaxUploadCountMinute = 10                 // 同一用户1分钟之内最大上传数量

	//登陆类型：1=密码，2=第三方
	LOGIN_TYPE_PWD   = 1
	LOGIN_TYPE_THIRD = 2

	//认证状态：0=未认证，1=认证中，2=认证成功，3=认证失败
	USER_AUTH_STATUS_NOT     = 0
	USER_AUTH_STATUS_IN      = 1
	USER_AUTH_STATUS_SUCCESS = 2
	USER_AUTH_STATUS_FAIL    = 3
)
