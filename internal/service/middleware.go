package service

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IMiddleware interface {
		CORSHandler(r *ghttp.Request)
		CtxHandler(r *ghttp.Request)
		I18NHandler(r *ghttp.Request)
		ResponseHandler(r *ghttp.Request)
		ErrorHandler(r *ghttp.Request)
		DecodeDataHandler(r *ghttp.Request)
	}
)

var (
	localMiddleware IMiddleware
)

func Middleware() IMiddleware {
	if localMiddleware == nil {
		panic("implement not found for interface IMiddleware, forgot register?")
	}
	return localMiddleware
}

func RegisterMiddleware(i IMiddleware) {
	localMiddleware = i
}
