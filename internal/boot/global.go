package boot

import (
	"context"
	"sviwo/internal/logic/analysis"
	"sviwo/pkg/cache"
	"sviwo/pkg/statistics"
)

func AllSystemInit(ctx context.Context) error {
	// 设置缓存适配器
	cache.SetAdapter(ctx)
	// 初始化统计设备数据
	statistics.InitCountDeviceData()

	//清除设备统计缓存
	analysis.RemoveDeviceStatusCountCache(ctx)

	return nil
}
