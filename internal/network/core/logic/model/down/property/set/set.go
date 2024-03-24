package set

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/util/guid"
	"sviwo/internal/consts"
	"sviwo/internal/network/core/logic/baseLogic"
	"sviwo/pkg/aliyun"
	"sviwo/pkg/iotModel/sviwoProtocol"
	"sviwo/pkg/iotModel/topicModel"
)

// 设备属性设置
func PropertySet(ctx context.Context, request topicModel.TopicDownHandlerData) (map[string]interface{}, error) {
	if request.DeviceDetail == nil {
		return nil, fmt.Errorf("device detail is nil")
	}
	if request.DeviceDetail.TSL == nil {
		return nil, fmt.Errorf("device tsl is nil")
	}

	requestDataMap := make(map[string]interface{})
	originRequestData := make(map[string]interface{})
	if err := json.Unmarshal(request.PayLoad, &originRequestData); err != nil {
		return nil, err
	}
	for k, v := range originRequestData {
		for _, property := range request.DeviceDetail.TSL.Properties {
			if property.Key == k {
				requestDataMap[k] = property.ValueType.ConvertValue(v)
			}
		}
	}
	r := sviwoProtocol.PropertySetRequest{
		Id:      guid.S(),
		Version: "1.0.0",
		Params:  requestDataMap,
		Method:  "thing.service.property.set",
	}
	//下发消息
	requestData, _ := json.Marshal(r.Params)
	err := aliyun.SetDevicePropertyRequest(ctx, request.DeviceDetail.ProductKey, request.DeviceDetail.DeviceName, string(requestData))
	if err != nil {
		return nil, err
	}

	baseLogic.InertTdLog(ctx, consts.MsgTypePropertySet, request.DeviceDetail.DeviceName, r)
	response, err := baseLogic.SyncRequest(ctx, r.Id, "SetProperty", r, 0)
	if err != nil {
		return nil, err
	} else if res, covertOk := response.(map[string]interface{}); !covertOk {
		return nil, fmt.Errorf("set property  failed,response: %+v", response)
	} else {
		return res, nil
	}
}
