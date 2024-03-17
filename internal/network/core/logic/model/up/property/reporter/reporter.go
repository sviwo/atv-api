package reporter

import (
	"context"
	"fmt"
	"strings"
	"sviwo/internal/consts"
	"sviwo/internal/network/core"
	"sviwo/pkg/iotModel/sviwoProtocol"
	"sviwo/pkg/iotModel/topicModel"
)

func Init() (err error) {
	if err = core.RegisterSubTopicHandler(sviwoProtocol.PropertySubRequestTopic, consts.MsgTypeEvent, ReportProperty); err != nil {
		return err
	}
	return nil
}

// ReportProperty 属性上报
func ReportProperty(ctx context.Context, data topicModel.TopicHandlerData) error {
	//对存在转义字符的进行全量替换
	payLoad := strings.ReplaceAll(string(data.PayLoad), "\\", "")
	fmt.Println("-----------属性上报-------", payLoad)
	return nil

}
