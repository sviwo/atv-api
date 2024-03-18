// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"sviwo/internal/model"
	"sviwo/pkg/iotModel"
	"sviwo/pkg/iotModel/sviwoProtocol"
)

type (
	IDevTSLParse interface {
		// ParseData 基于物模型解析上报数据
		ParseData(ctx context.Context, deviceKey string, data []byte) (res iotModel.ReportPropertyData, err error)
		// HandleProperties 处理属性
		HandleProperties(ctx context.Context, device *model.DeviceOutput, properties map[string]interface{}) (reportDataInfo iotModel.ReportPropertyData, err error)
		// HandleEvents 处理事件上报
		HandleEvents(ctx context.Context, device *model.DeviceOutput, events map[string]sviwoProtocol.EventNode) (res []iotModel.ReportEventData, err error)
	}
	IDevProduct interface {
		Detail(ctx context.Context, key string) (out *model.DetailProductOutput, err error)
		List(ctx context.Context) (list []*model.ProductOutput, err error)
	}
	IDevDevice interface {
		// Get 获取设备详情
		Get(ctx context.Context, key string) (out *model.DeviceOutput, err error)
		// CacheDeviceDetailList 缓存所有设备详情数据
		CacheDeviceDetailList(ctx context.Context) (err error)
		List(ctx context.Context, productKey string, keyWord string) (list []*model.DeviceOutput, err error)

	}
)

var (
	localDevTSLParse       IDevTSLParse
	localDevProduct        IDevProduct
	localDevDevice         IDevDevice
)

func DevTSLParse() IDevTSLParse {
	if localDevTSLParse == nil {
		panic("implement not found for interface IDevTSLParse, forgot register?")
	}
	return localDevTSLParse
}

func RegisterDevTSLParse(i IDevTSLParse) {
	localDevTSLParse = i
}

func DevProduct() IDevProduct {
	if localDevProduct == nil {
		panic("implement not found for interface IDevProduct, forgot register?")
	}
	return localDevProduct
}

func RegisterDevProduct(i IDevProduct) {
	localDevProduct = i
}

func DevDevice() IDevDevice {
	if localDevDevice == nil {
		panic("implement not found for interface IDevDevice, forgot register?")
	}
	return localDevDevice
}

func RegisterDevDevice(i IDevDevice) {
	localDevDevice = i
}


