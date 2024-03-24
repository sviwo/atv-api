package analysis

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"sviwo/internal/dao"
	"sviwo/internal/model/entity"
)

// getDeviceKeys 获取所有设备的key，用于过滤数据权限
func getDeviceKeys(ctx context.Context, productKey string) (deviceKeys []string) {
	var deviceList []*entity.Device
	m := dao.Device.Ctx(ctx)

	if productKey != "" {
		m = m.Where(dao.Device.Columns().ProductKey, productKey)
	}
	err := m.Scan(&deviceList)
	if err != nil {
		g.Log().Debug(ctx, err.Error())
		return
	}
	if err != nil {
		return
	}
	for _, device := range deviceList {
		deviceKeys = append(deviceKeys, device.DeviceName)
	}
	return
}
