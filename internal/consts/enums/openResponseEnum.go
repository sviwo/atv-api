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
	Fail    = New(0, "系统错误，请联系管理员")
	Success = New(1, "SUCCESS")
	/*
	   公共异常 1000～1999
	*/
	ApiException           = New(1000, "网络开小差了,请稍后再试")
	IllegalOperation       = New(1001, "非法操作")
	IllegalArgument        = New(1002, "请求参数缺失")
	RequestParamTypeError  = New(1003, "请求参数错误")
	RequestMethodTypeError = New(1004, "请求方式错误")
	TransactionUserError   = New(1999, "用户重复请求或伪造请求")

	/*
	   app错误码 2000～2999
	*/
	UserLoginFailed   = New(2000, "登录失败，用户名或密码错误")
	UserExists        = New(2001, "此用户已存在")
	UserNotExists     = New(2002, "此用户不存在")
	UserAcctFrozen    = New(2003, "此账户已被冻结")
	VftCodeError      = New(2004, "验证码错误")
	VftCodeOverdue    = New(2005, "验证码已过期")
	VftCodeSendFailed = New(2006, "验证码发送失败，请检查网络")
	CarNotExists      = New(2007, "此车辆不存在")
)

func New(code int, message string) gcode.Code {
	return OpenResponseEnum{code: code, message: message}
}
