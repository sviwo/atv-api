package event

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
	if err = core.RegisterSubTopicHandler(sagooProtocol.PropertySubRequestTopic, consts.MsgTypeEvent, ReportEvent); err != nil {
		return err
	}
	return nil
}

// 事件上报
func ReportEvent(ctx context.Context, data topicModel.TopicHandlerData) error {
	fmt.Println("-----------处理事件上报-------", string(data.PayLoad))
	return nil
}
