package consts

/**
 * redis key值常量，建议使用冒号分割
 * 例如：user：%s:
 */
const (
	UserWebVerify   = "u:%s:web"
	UserLoginVerify = "u:%s:verify"
	UserLoginToken  = "u:%s:token"
	TaskRemindMsg   = "t:%s:msg"

	/**
	 * 缓存资源数据
	 **/
	MethodUserServiceRedisKey = "method:userService:getvftcode:%s"
)
