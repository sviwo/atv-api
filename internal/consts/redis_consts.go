package consts

/**
 * redis key值常量，建议使用冒号分割
 * 例如：user：%s:
 */
const (
	RedisUserWebVerify  = "u:%s:web"
	RedisEmailVftCode   = "eVftCode:%s"
	RedisUserLoginToken = "u:%s:token"
	RedisTaskRemindMsg  = "t:%s:msg"
	RedisEccPrivateKey  = "ecc:%s:private"
	RedisCarKey         = "car:%s:key"

	/**
	 * 缓存资源数据
	 **/
	RedisMethodReqLimit = "method:%s:"
)
