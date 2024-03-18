package boot

import (
	"context"
	"sviwo/pkg/cache"
)

func AllSystemInit(ctx context.Context) error {
	// 设置缓存适配器
	cache.SetAdapter(ctx)

	return nil
}
