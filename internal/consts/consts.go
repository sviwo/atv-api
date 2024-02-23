package consts

/*
代码业务常量管理
*/
const (
	ProjectName              = "sviwo"
	ProjectUsage             = "sviwo app server"
	ProjectBrief             = "start http server"
	Version                  = "v1.0.0"      // 当前服务版本(用于模板展示)
	CaptchaName              = "UserVftCode" // 验证码存储空间名称
	ContextKey               = "ContextKey"  // 上下文变量存储键名，前后端系统共享
	FileMaxUploadCountMinute = 10            // 同一用户1分钟之内最大上传数量

	//登陆类型：1=密码，2=第三方
	LOGIN_TYPE_PWD   = 1
	LOGIN_TYPE_THIRD = 2
)
