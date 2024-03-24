package event

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"strings"
	"sviwo/internal/consts"
	"sviwo/internal/model"
	"sviwo/internal/network/core"
	"sviwo/internal/network/core/logic/model/up/property/reporter"
	"sviwo/internal/service"
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
	eventKey := topicInfo[5]
	if eventKey == "property" {
		// 忽略属性上报信息
		return errors.New("ignore")
	}
	glog.Printf(ctx, "-----------事件上报---tpoic:%s---内容：%s--", data.Topic, string(data.PayLoad))
	var reportData sviwoProtocol.ReportEventReq
	if reportDataErr := json.Unmarshal(data.PayLoad, &reportData); reportDataErr != nil {
		g.Log().Errorf(ctx, "parse data error: %v, topic:%s, message:%s, message ignored", reportDataErr, data.Topic, string(data.PayLoad))
		return reportDataErr
	}
	var reportEventData = model.ReportEventData{
		Key: eventKey,
		Param: model.ReportEventParam{
			Value:      map[string]any{},
			CreateTime: reportData.GmtCreate,
		},
	}
	for _, event := range data.DeviceDetail.TSL.Events {
		if event.Key == eventKey {
			for _, o := range event.Outputs {
				for k, v := range reportData.Value {
					if k == o.Key {
						reportEventData.Param.Value[k] = o.ValueType.ConvertValue(v)
					}
				}
			}
		}
	}
	if reportEventErr := service.DevDataReport().Event(ctx, data.DeviceKey, reportEventData); reportEventErr != nil {
		g.Log().Errorf(ctx, "report event error: %v, topic:%s, message:%s, message ignored", reportEventErr, data.Topic, string(data.PayLoad))
		return reportEventErr
	}
	return nil
}
