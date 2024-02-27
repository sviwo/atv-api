package network

import (
	"context"
	"sviwo/internal/network/core"
)

func ReloadNetwork(c context.Context) (err error) {
	for _, f := range []func(ctx context.Context) error{core.StartSubscriber} {
		if err = f(c); err != nil {
			return err
		}
	}
	return
}
