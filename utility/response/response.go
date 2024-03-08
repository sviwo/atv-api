package response

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/net/ghttp"
	"sviwo/internal/consts/enums"
)

// JsonRes 数据返回通用JSON数据结构
type JsonRes struct {
	Code    int         `json:"code"` // 错误码((0:成功, 1:失败, >1:错误码))
	Message string      `json:"msg"`  // 提示信息
	Data    interface{} `json:"data"` // 返回数据(业务接口定义具体数据结构)
}

// Json 返回标准JSON数据。
func Json(r *ghttp.Request, gc gcode.Code, data ...interface{}) {
	var responseData interface{}
	if len(data) > 0 {
		responseData = data[0]
	} else {
		responseData = g.Map{}
	}
	r.Response.WriteJson(JsonRes{
		Code:    gc.Code(),
		Message: gi18n.T(r.GetCtx(), gc.Message()),
		Data:    responseData,
	})
}

// JsonExit 返回标准JSON数据并退出当前HTTP执行函数。
func JsonExit(r *ghttp.Request, gc gcode.Code, data interface{}) {
	Json(r, gc, data)
	r.Exit()
}

func SuccessMsg(r *ghttp.Request, data interface{}) {
	Json(r, enums.Success, data)
}

func FailMsg(r *ghttp.Request) {
	Json(r, enums.Fail)
}
