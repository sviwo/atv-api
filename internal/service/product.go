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
	IDevDeviceProperty interface {
		// Set 设备属性设置
		Set(ctx context.Context, in *model.DevicePropertyInput) (out *model.DevicePropertyOutput, err error)
		// 获取注册设备到指定产品下所需要的证书
		GetDeviceSecret(ctx context.Context, deviceCode string) (out *model.DeviceSecretOutput)
	}
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
	IDevDataReport interface {
		// Event 设备事件上报
		Event(ctx context.Context, deviceKey string, data model.ReportEventData, subKey ...string) error
		// Property 设备属性上报
		Property(ctx context.Context, deviceKey string, data model.ReportPropertyData, subKey ...string) error
	}
	IDevInit interface {
		// InitProductForTd 产品表结构初始化
		InitProductForTd(ctx context.Context) (err error)
		// InitDeviceForTd 设备表结构初始化
		InitDeviceForTd(ctx context.Context) (err error)
	}
	IDevDevice interface {
		// Get 获取设备详情
		Get(ctx context.Context, key string) (out *model.DeviceOutput, err error)
		// CacheDeviceDetailList 缓存所有设备详情数据
		CacheDeviceDetailList(ctx context.Context) (err error)
		List(ctx context.Context, productKey string, keyWord string) (list []*model.DeviceOutput, err error)
		BatchUpdateDeviceStatusInfo(ctx context.Context, deviceStatusLogList []iotModel.DeviceStatusLog) (err error)
	}
)

var (
	localDevTSLParse       IDevTSLParse
	localDevProduct        IDevProduct
	localDevDevice         IDevDevice
	localDevInit           IDevInit
	localDevDataReport     IDevDataReport
	localDevDeviceProperty IDevDeviceProperty
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

func DevInit() IDevInit {
	if localDevInit == nil {
		panic("implement not found for interface IDevInit, forgot register?")
	}
	return localDevInit
}

func RegisterDevInit(i IDevInit) {
	localDevInit = i
}

func DevDataReport() IDevDataReport {
	if localDevDataReport == nil {
		panic("implement not found for interface IDevDataReport, forgot register?")
	}
	return localDevDataReport
}

func RegisterDevDataReport(i IDevDataReport) {
	localDevDataReport = i
}

func DevDeviceProperty() IDevDeviceProperty {
	if localDevDeviceProperty == nil {
		panic("implement not found for interface IDevDeviceProperty, forgot register?")
	}
	return localDevDeviceProperty
}

func RegisterDevDeviceProperty(i IDevDeviceProperty) {
	localDevDeviceProperty = i
}
