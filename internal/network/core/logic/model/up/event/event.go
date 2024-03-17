package event

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sviwo/internal/consts"
	"sviwo/internal/network/core"
	"sviwo/internal/network/core/logic/model/up/property/reporter"
	sviwoProtocol "sviwo/pkg/iotModel/sviwoProtocol"
	"sviwo/pkg/iotModel/topicModel"
)

func Init() (err error) {
	//  /sys/${productKey}/${devicekey}/thing/event/${tsl.event.identifier}/post
	if err = core.RegisterSubTopicHandler(sviwoProtocol.EvnetTopic, consts.MsgTypeEvent, ReportEvent); err != nil {
		return err
	}
	return nil
}

// 事件上报
func ReportEvent(ctx context.Context, data topicModel.TopicHandlerData) error {
	if strings.HasSuffix(data.Topic, "property/post") {
		// 这段特殊的逻辑是为了处理topic通配符的问题
		return reporter.ReportProperty(ctx, data)
	}
	topicInfo := strings.Split(data.Topic, "/")
	eventKey := topicInfo[6]
	if eventKey == "property" {
		// 忽略属性上报信息
		return errors.New("ignore")
	}
	fmt.Println("-----------事件上报-------", string(data.PayLoad))
	return nil
}
