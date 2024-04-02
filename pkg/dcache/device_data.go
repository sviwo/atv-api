package dcache

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"sviwo/internal/service"
	"sviwo/pkg/iotModel"
)

// InertDeviceLog 插入设备日志
func InertDeviceLog(ctx context.Context, logType, deviceKey string, obj interface{}) {
	str, strIsOk := obj.(string)
	content := str
	if !strIsOk {
		objStr, _ := json.Marshal(obj)
		content = string(objStr)
	}
	// 向设备缓存数据库插入数据
	if err := DB().InsertData(context.Background(), deviceKey, iotModel.DeviceLog{
		Ts:      gtime.Now(),
		Device:  deviceKey,
		Type:    logType,
		Content: content,
	}); err != nil {
		g.Log().Debugf(ctx, "Failed to insert data: %v\n", err)
	}
}

// GetDeviceDetailDataByLatest 获取设备解析后的最新一条数据
func GetDeviceDetailDataByLatest(ctx context.Context, deviceKey string) (res iotModel.ReportPropertyData) {
	// 从设备缓存数据库获取数据
	data, err := DB().GetDataByLatest(context.Background(), deviceKey)
	if err != nil || data == "" {
		g.Log().Debugf(ctx, "Failed to get data: %v", err)
		return
	}

	var value = iotModel.DeviceLog{}
	if err := json.Unmarshal([]byte(data), &value); err != nil {
		g.Log().Debugf(ctx, "Failed to unmarshal data: %v", err)
		return
	}

	// 基于物模型解析数据
	res, err = service.DevTSLParse().ParseData(ctx, deviceKey, []byte(value.Content))
	if err != nil {
		g.Log().Debugf(ctx, "Failed to parse data: %v", err)
	}
	return
}
