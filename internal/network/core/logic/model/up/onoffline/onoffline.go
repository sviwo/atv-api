package onoffline

import (
	"context"
	"fmt"
	"sviwo/internal/consts"
	"sviwo/internal/network/core"
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

// 事件上报
func OnOff(ctx context.Context, data topicModel.TopicHandlerData) error {
	fmt.Println("-----------设备上下线-------", data)
	return nil
}
