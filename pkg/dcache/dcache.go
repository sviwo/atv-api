package dcache

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"strings"
	"sviwo/internal/consts"
	"sviwo/internal/model"
	"sviwo/internal/service"
	"sviwo/pkg/cache"
)

// GetDeviceStatus 获取指定的设备状态
func GetDeviceStatus(ctx context.Context, deviceKey string) (res int) {
	data, err := cache.Instance().Get(ctx, consts.DeviceStatusPrefix+deviceKey)
	if err != nil || data == nil {
		return 1
	}
	if data.Val() != nil {
		return 2
	}
	return 1
}

// GetOnlineDeviceList 获取在线设备列表
func GetOnlineDeviceList() (list g.Slice, err error) {
	dataList, err := SearchKey(consts.DeviceStatusPrefix)
	if err != nil {
		return
	}
	for _, value := range dataList {
		list = append(list, strings.TrimPrefix(value, consts.DeviceStatusPrefix))
	}
	return
}

// GetAllDeviceList 获取缓存中的所有设备列表
func GetAllDeviceList() (outList []*model.DeviceOutput, err error) {
	deviceKeyList, err := SearchKey(consts.DeviceDetailInfoPrefix)
	if err != nil {
		return
	}
	for _, key := range deviceKeyList {
		data, err := cache.Instance().Get(context.Background(), key)
		if err != nil || data.Val() == nil {
			continue
		}
		var out model.DeviceOutput
		if err = gconv.Scan(data.Val(), &out); err != nil {
			continue
		}
		outList = append(outList, &out)
	}
	return
}

// CountDeviceOnlineNum 统计在线设备数量
func CountDeviceOnlineNum() (num int) {
	data, err := SearchKey(consts.DeviceStatusPrefix)
	if err != nil {
		return 0
	}
	num = len(data)
	return
}

// GetDeviceDetailInfo 获取设备详情缓存
func GetDeviceDetailInfo(deviceKey string) (out *model.DeviceOutput, err error) {
	data, err := cache.Instance().Get(context.Background(), consts.DeviceDetailInfoPrefix+deviceKey)
	if err != nil || data.Val() == nil {
		return
	}
	if err = gconv.Scan(data.Val(), &out); err != nil {
		return
	}

	return
}

// SetDeviceDetailInfo 设置设备详情缓存
func SetDeviceDetailInfo(deviceKey string, data *model.DeviceOutput) (err error) {
	if data == nil || data.Product == nil {
		return
	}
	//设备启用后，更新缓存数据
	if data.Product.Metadata != "" {
		_ = json.Unmarshal([]byte(data.Product.Metadata), &data.TSL)
		data.Product.Metadata = ""
	}
	err = cache.Instance().Set(context.Background(), consts.DeviceDetailInfoPrefix+deviceKey, data, 0)
	return
}

// GetProductDetailInfo 获取产品详情缓存
func GetProductDetailInfo(productKey string) (out *model.DetailProductOutput, err error) {
	data, err := cache.Instance().Get(context.Background(), consts.ProductDetailInfoPrefix+productKey)
	if err != nil || data.Val() == nil {
		return
	}
	if err = gconv.Scan(data.Val(), &out); err != nil {
		return
	}
	productDetailInfo, err := service.DevProduct().Detail(context.Background(), productKey)
	if err != nil {
		return
	}
	err = SetProductDetailInfo(productKey, productDetailInfo)
	return
}

// SetProductDetailInfo 设置产品详情缓存
func SetProductDetailInfo(productKey string, data *model.DetailProductOutput) (err error) {
	if data == nil {
		return
	}
	//设备启用后，更新缓存数据
	if data.Metadata != "" {
		_ = json.Unmarshal([]byte(data.Metadata), &data.TSL)
		data.Metadata = ""
	}
	err = cache.Instance().Set(context.Background(), consts.ProductDetailInfoPrefix+productKey, data, 0)
	return
}

// SearchKey 搜索指定的key
func SearchKey(keyword string) (keys []string, err error) {
	data, err := cache.Instance().Keys(context.Background())
	if err != nil {
		return nil, err
	}
	for _, item := range data {
		if strings.Contains(item.(string), keyword) {
			keys = append(keys, item.(string))
		}
	}
	return
}
