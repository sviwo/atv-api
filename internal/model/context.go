package model

import "github.com/gogf/gf/v2/container/gmap"

// Context 请求上下文结构
type Context struct {
	Data *gmap.Map // 自定KV变量，业务模块根据需要设置，不固定
}
