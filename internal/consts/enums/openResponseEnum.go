package enums

import "github.com/gogf/gf/v2/errors/gcode"

/*
错误提示码管理
*/
type OpenResponseEnum struct {
	code    int
	message string
}

func (c OpenResponseEnum) Code() int {
	return c.code
}

func (c OpenResponseEnum) Message() string {
	return c.message
}

func (c OpenResponseEnum) Detail() interface{} {
	return nil
}

var (
	// 系统错误，httpStatus = 500
	Fail = New(0, "系统错误，请联系管理员")

	// httpStatus = 200
	Success = New(1, "SUCCESS")

	// httpStatus = 401
	LoginOverdue = New(999, "登录过期，请重新登录")

	/*
		   公共异常 1000～1999
			httpStatus = 200
	*/
	IllegalOperation          = New(1000, "非法操作")
	RequestMissingParam       = New(1001, "请求参数缺失")
	RequestParamTypeError     = New(1002, "请求参数错误")
	RequestMethodTypeError    = New(1003, "请求方式错误")
	RequestThirdInterFaceFail = New(1004, "请求第三方接口失败")
	MethodReqLimitError       = New(1005, "请勿频繁请求")
	TransactionUserError      = New(1999, "用户重复请求或伪造请求")

	/*
		   app错误码 2000～2999
			httpStatus = 200
	*/
	UserLoginFailed    = New(2000, "登录失败，用户名或密码错误")
	UserExists         = New(2001, "此用户已存在")
	UserNotExists      = New(2002, "此用户不存在")
	UserAcctFrozen     = New(2003, "此账户已被冻结")
	VftCodeError       = New(2004, "验证码错误")
	VftCodeOverdue     = New(2005, "验证码已过期")
	VftCodeSendFailed  = New(2006, "验证码发送失败，请检查网络")
	CarNotExists       = New(2007, "车辆不存在")
	DeviceOffline      = New(2008, "设备不在线")
	IllegalDevice      = New(2009, "非法设备")
	UnbindCarError     = New(2010, "您是此车的主用户，不可解绑")
	BindCarError       = New(2011, "您已绑定过此车辆")
	CarKeyInvalidError = New(2012, "这把车钥匙已失效")
	ShareCarKeyError   = New(2013, "只能分享自己车的钥匙")
	CarKeyLimitError   = New(2014, "此车的钥匙已达上限")

	/*
		   第三方错误码 3000～3999
			httpStatus = 200
	*/
)

func New(code int, message string) gcode.Code {
	return OpenResponseEnum{code: code, message: message}
}
