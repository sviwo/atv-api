package product

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"sviwo/internal/model"
	"sviwo/internal/service"
	"sviwo/pkg/dcache"
	"sviwo/pkg/iotModel"
	"sviwo/pkg/iotModel/sviwoProtocol"
	"time"
)

type sDevTSLParse struct {
}

func init() {
	service.RegisterDevTSLParse(New())
}

func New() *sDevTSLParse {
	return &sDevTSLParse{}
}

// ParseData 基于物模型解析上报数据
func (s *sDevTSLParse) ParseData(ctx context.Context, deviceKey string, data []byte) (res iotModel.ReportPropertyData, err error) {
	if data == nil || len(data) == 0 {
		return nil, errors.New("data is empty")
	}
	var reportData sviwoProtocol.ReportPropertyReq
	if err = json.Unmarshal(data, &reportData); err != nil {
		g.Log().Errorf(ctx, "parse data error: %s, message:%s, message ignored", err, string(data))
		return
	}
	if reportData.Params == nil || len(reportData.Params) == 0 {
		return nil, errors.New("report data is empty")
	}
	device, err := dcache.GetDeviceDetailInfo(deviceKey)
	if err != nil {
		return
	}
	if device == nil {
		g.Log().Errorf(ctx, "device not found, deviceKey:%s", deviceKey)
		return nil, errors.New("device not found")
	}
	res, err = s.HandleProperties(ctx, device, reportData.Params)
	if err != nil {
		return
	}
	return
}

// HandleProperties 处理属性
func (s *sDevTSLParse) HandleProperties(ctx context.Context, device *model.DeviceOutput, properties map[string]interface{}) (reportDataInfo iotModel.ReportPropertyData, err error) {
	reportDataInfo = make(iotModel.ReportPropertyData)
	nowTime := time.Now()
	for k, v := range properties {
		for _, property := range device.TSL.Properties {
			if property.Key == k {
				var createTimestamp int64
				var value interface{}

				if mapInfo, ok := v.(map[string]interface{}); ok {
					// 处理带时间戳的属性值
					if timeValue, timeOK := mapInfo["time"].(float64); timeOK && mapInfo["value"] != nil {
						createTimestamp = int64(timeValue)
						value = property.ValueType.ConvertValue(mapInfo["value"])
					}
				} else {
					// 处理不带时间戳的属性值
					createTimestamp = nowTime.Unix()
					value = property.ValueType.ConvertValue(v)
				}

				// 构建数据
				reportDataInfo[k] = iotModel.ReportPropertyNode{
					CreateTime: createTimestamp,
					Value:      value,
				}

				break
			}
		}
	}
	return
}

// HandleEvents 处理事件上报
func (s *sDevTSLParse) HandleEvents(ctx context.Context, device *model.DeviceOutput, events map[string]sviwoProtocol.EventNode) (res []iotModel.ReportEventData, err error) {
	res = make([]iotModel.ReportEventData, 0, len(events))
	for eventKey, eventData := range events {
		// 查找对应的事件定义
		for _, event := range device.TSL.Events {
			if event.Key == eventKey {
				reportEventData := iotModel.ReportEventData{
					Key: eventKey,
					Param: iotModel.ReportEventParam{
						Value:      make(map[string]interface{}),
						CreateTime: eventData.CreateTime,
					},
				}

				// 遍历事件中的每个输出参数
				for _, output := range event.Outputs {
					if value, exists := eventData.Value[output.Name]; exists {
						reportEventData.Param.Value[output.Name] = output.ValueType.ConvertValue(value)
					}
				}

				// 上报事件
				if len(reportEventData.Param.Value) > 0 {
					res = append(res, reportEventData)
				}
				break
			}
		}
	}
	return
}
