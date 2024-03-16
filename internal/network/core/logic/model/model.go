package model

import (
	"context"
	"sviwo/internal/network/core/logic/model/up/event"
	"sviwo/internal/network/core/logic/model/up/onoffline"
	"sviwo/internal/network/core/logic/model/up/property/reporter"
)

func InitCoreLogic(ctx context.Context) error {
	for _, v := range []func() error{
		event.Init,     //事件
		reporter.Init,  //属性上报
		onoffline.Init, //上下线
	} {
		if err := v(); err != nil {
			return err
		}
	}
	return nil
}
