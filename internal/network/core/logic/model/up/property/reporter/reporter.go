package reporter

import (
	"context"
	"fmt"
	"sviwo/internal/consts"
	"sviwo/internal/network/core"
	sagooProtocol "sviwo/pkg/iotModel/sviwoProtocol"
	"sviwo/pkg/iotModel/topicModel"
)

func Init() (err error) {
	if err = core.RegisterSubTopicHandler(sagooProtocol.PropertySubRequestTopic, consts.MsgTypeEvent, ReportProperty); err != nil {
		return err
	}
	return nil
}

// ReportProperty 属性上报
func ReportProperty(ctx context.Context, data topicModel.TopicHandlerData) error {
	fmt.Println("-----------属性上报-------", string(data.PayLoad))
	return nil

}
