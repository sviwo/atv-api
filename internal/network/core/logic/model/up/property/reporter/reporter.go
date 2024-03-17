package reporter

import (
	"context"
	"github.com/gogf/gf/v2/os/glog"
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
	glog.Printf(ctx, "-----------属性上报---tpoic:%s---内容：%s--", data.Topic, payLoad)
	return nil

}
