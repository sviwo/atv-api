package onoffline

import (
	"context"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
	"sviwo/internal/consts"
	"sviwo/internal/network/core"
	"sviwo/pkg/dcache"
	"sviwo/pkg/gpool"
	"sviwo/pkg/iotModel/sviwoProtocol"
	"sviwo/pkg/iotModel/topicModel"
)

func Init() (err error) {
	//  /sys/${productKey}/${devicekey}/thing/event/${tsl.event.identifier}/post
	if err = core.RegisterSubTopicHandler(sviwoProtocol.PropertyOnOfflineTopic, consts.MsgTypeEvent, OnOff); err != nil {
		return err
	}
	return nil
}

// 控制携程数量
var gPool = gpool.NewGPool(1000)

// 事件上报
func OnOff(ctx context.Context, data topicModel.TopicHandlerData) error {
	glog.Printf(ctx, "-----------设备上下线---tpoic:%s---内容：%s--", data.Topic, string(data.PayLoad))
	gPool.Go(func(ctx context.Context) error {
		payloadMap := gconv.Map(data.PayLoad)
		status := payloadMap["status"]
		if status == "offline" {
			dcache.Offline(ctx, data.DeviceDetail)
		}
		return nil
	})
	return nil
}
