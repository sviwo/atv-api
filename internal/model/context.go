package model

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Context 请求上下文结构
type Context struct {
	Data g.Map // 自定KV变量，业务模块根据需要设置，不固定
}