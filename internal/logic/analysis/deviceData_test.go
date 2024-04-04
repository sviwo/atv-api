package analysis

import (
	"context"
	"sviwo/internal/model"
	"sviwo/internal/service"
	"sviwo/pkg/cache"
	"sviwo/pkg/general"
	"testing"
)

// 测试获取设备在线离线统计数据
func TestGetDeviceData(t *testing.T) {
	// 设置缓存适配器
	cache.SetAdapter(context.Background())

	var reqData model.DeviceDataReq
	reqData.DeviceKey = "t202200002"
	reqData.PageSize = 10
	reqData.PageNum = 1

	data, err := service.AnalysisDeviceData().GetDeviceData(context.Background(), reqData)
	if err != nil {
		t.Log(err)
	}
	t.Log(data)
}

func TestGetDeviceAlarmData(t *testing.T) {
	var reqData = new(general.SelectReq)
	reqData.KeyWords = "t202200002"
	data, err := service.AnalysisDeviceData().GetDeviceAlarmLogData(context.Background(), reqData)
	if err != nil {
		t.Log(err)
	}
	t.Log(data)
}
