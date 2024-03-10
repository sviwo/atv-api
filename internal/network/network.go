package network

import (
	"context"
	"sviwo/internal/network/core"
)

var reloadNetWorkFunc = []func(ctx context.Context) error{
	// 开启主题订阅
	core.StartSubscriber,
}

func ReloadNetwork(c context.Context) (err error) {
	for _, f := range reloadNetWorkFunc {
		if err = f(c); err != nil {
			return err
		}
	}
	return
}
