package boot

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"sviwo/pkg/cache"
)

func AllSystemInit(ctx context.Context) {
	// 默认上海时区
	if err := gtime.SetTimeZone("Asia/Shanghai"); err != nil {
		g.Log().Fatalf(ctx, "时区设置异常 err：%+v", err)
		return
	}
	// 设置缓存适配器
	cache.SetAdapter(ctx)
}
