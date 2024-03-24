package product

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"sviwo/internal/consts"
	"sviwo/internal/dao"
	"sviwo/internal/model"
	"sviwo/internal/service"
	"sviwo/pkg/cache"
	"sviwo/pkg/dcache"
	"sviwo/pkg/iotModel"
)

type sDevDevice struct{}

func init() {
	service.RegisterDevDevice(DeviceNew())
}

func DeviceNew() *sDevDevice {
	return &sDevDevice{}
}

// CacheDeviceDetailList 缓存所有设备详情数据
func (s *sDevDevice) CacheDeviceDetailList(ctx context.Context) (err error) {
	productList, err := service.DevProduct().List(context.Background())
	if err != nil {
		return
	}
	for _, p := range productList {
		deviceList, err := service.DevDevice().List(context.Background(), p.ProductKey, "")
		if err != nil {
			g.Log().Error(ctx, err.Error())
		}

		//缓存产品详细信息
		var detailProduct = new(model.DetailProductOutput)
		err = gconv.Scan(p, detailProduct)
		if err == nil {
			err = dcache.SetProductDetailInfo(p.ProductKey, detailProduct)
		} else {
			g.Log().Error(ctx, err.Error())
		}

		for _, d := range deviceList {
			if d.Product.Metadata != "" {
				err = json.Unmarshal([]byte(d.Product.Metadata), &d.TSL)
				d.Product.Metadata = ""
				if err != nil {
					continue
				}
			}
			//缓存设备详细信息
			err := cache.Instance().Set(context.Background(), consts.DeviceDetailInfoPrefix+d.DeviceName, d, 0)
			if err != nil {
				g.Log().Error(ctx, err.Error())
			}
		}
	}
	return

}

// Get 获取设备详情
func (s *sDevDevice) Get(ctx context.Context, deviceCode string) (out *model.DeviceOutput, err error) {
	err = dao.Device.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: 0,
		Name:     consts.GetDetailDeviceOutput + deviceCode,
		Force:    false,
	}).WithAll().Where(dao.Device.Columns().DeviceCode, deviceCode).Scan(&out)
	if err != nil {
		return
	}
	if out == nil {
		err = errors.New("设备不存在")
		return
	}
	if out.Status != 0 {
		out.Status = dcache.GetDeviceStatus(ctx, out.DeviceName) //查询设备状态
	}
	if out.Product != nil {
		out.ProductName = out.Product.ProductName
		if out.Product.Metadata != "" {
			err = json.Unmarshal([]byte(out.Product.Metadata), &out.TSL)
		}
	}
	return
}

// List 已发布产品的设备列表
func (s *sDevDevice) List(ctx context.Context, productKey string, keyWord string) (list []*model.DeviceOutput, err error) {
	m := dao.Device.Ctx(ctx).
		Where(dao.Device.Columns().Status+" > ?", model.DeviceStatusNoEnable)
	if productKey != "" {
		m = m.Where(dao.Device.Columns().ProductKey, productKey)
	}
	if keyWord != "" {
		m = m.WhereLike(dao.Device.Columns().DeviceCode, "%"+keyWord+"%").WhereOrLike(dao.Device.Columns().DeviceName, "%"+keyWord+"%")
	}

	err = m.WhereIn(dao.Device.Columns().ProductKey,
		dao.Product.Ctx(ctx).
			Fields(dao.Product.Columns().ProductKey).
			Where(dao.Product.Columns().Status, model.ProductStatusOn)).
		WithAll().
		OrderDesc(dao.Device.Columns().DeviceId).
		Scan(&list)
	if err != nil {
		return
	}

	for i, v := range list {
		if v.Product != nil {
			list[i].ProductName = v.Product.ProductName
		}
	}

	return
}

// BatchUpdateDeviceStatusInfo 批量更新设备状态信息，设备上线、离线、注册
func (s *sDevDevice) BatchUpdateDeviceStatusInfo(ctx context.Context, deviceStatusLogList []iotModel.DeviceStatusLog) (err error) {

	onLineData := g.Map{}
	onlineDeviceKeyList := make([]string, 0)

	offLineData := g.Map{}
	offLineDeviceKeyList := make([]string, 0)

	registryData := g.Map{}
	registryDeviceKeyList := make([]string, 0)

	for _, statusLog := range deviceStatusLogList {
		switch statusLog.Status {
		case consts.DeviceStatueOnline:
			device, err := dcache.GetDeviceDetailInfo(statusLog.DeviceKey)
			if err != nil {
				continue
			}

			if device.RegistryTime == nil {
				registryData[dao.Device.Columns().RegistryTime] = statusLog.Timestamp
				registryDeviceKeyList = append(registryDeviceKeyList, statusLog.DeviceKey)
			}

			onLineData[dao.Device.Columns().LastOnlineTime] = statusLog.Timestamp
			onLineData[dao.Device.Columns().Status] = consts.DeviceStatueOnline
			onlineDeviceKeyList = append(onlineDeviceKeyList, statusLog.DeviceKey)
		case consts.DeviceStatueOffline:
			offLineData[dao.Device.Columns().LastOnlineTime] = statusLog.Timestamp
			offLineData[dao.Device.Columns().Status] = consts.DeviceStatueOffline
			offLineDeviceKeyList = append(offLineDeviceKeyList, statusLog.DeviceKey)
		}
	}
	if len(registryDeviceKeyList) > 0 {
		_, err = dao.Device.Ctx(ctx).
			Data(registryData).
			WhereIn(dao.Device.Columns().DeviceName, registryDeviceKeyList).
			Update()
	}
	if len(onlineDeviceKeyList) > 0 {
		_, err = dao.Device.Ctx(ctx).
			Data(onLineData).
			WhereIn(dao.Device.Columns().DeviceName, onlineDeviceKeyList).
			Update()
	}
	if len(offLineDeviceKeyList) > 0 {
		_, err = dao.Device.Ctx(ctx).
			Data(offLineData).
			WhereIn(dao.Device.Columns().DeviceName, onlineDeviceKeyList).
			Update()
	}

	return
}
