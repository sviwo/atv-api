package onoffline

import (
	"context"
	"fmt"
	"sviwo/internal/consts"
	"sviwo/internal/network/core"
	sagooProtocol "sviwo/pkg/iotModel/sviwoProtocol"
	"sviwo/pkg/iotModel/topicModel"
)

func Init() (err error) {
	//  /sys/${productKey}/${devicekey}/thing/event/${tsl.event.identifier}/post
	if err = core.RegisterSubTopicHandler(sagooProtocol.PropertyOnOfflineTopic, consts.MsgTypeEvent, OnOff); err != nil {
		return err
	}
	return nil
}

// 事件上报
func OnOff(ctx context.Context, data topicModel.TopicHandlerData) error {
	fmt.Println("-----------设备上下线-------", string(data.PayLoad))
	return nil
}
