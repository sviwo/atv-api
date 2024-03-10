package dcache

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"sviwo/internal/consts"
	"sviwo/internal/model"
	"sviwo/pkg/cache"
)

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
