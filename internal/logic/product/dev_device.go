package product

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"strings"
	"sviwo/internal/consts"
	"sviwo/internal/consts/enums"
	"sviwo/internal/dao"
	"sviwo/internal/model"
	"sviwo/internal/model/entity"
	"sviwo/internal/service"
	"sviwo/pkg/cache"
	"sviwo/pkg/dcache"
	"sviwo/pkg/iotModel"
	"sviwo/pkg/tsd"
	"sviwo/pkg/tsd/comm"
	"sviwo/pkg/utility"
)

func init() {
	service.RegisterDevDevice(DeviceNew())
}

func DeviceNew() *sDevDevice {
	return &sDevDevice{}
}

type sDevDevice struct{}

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
	}).WithAll().Where(dao.Device.Columns().DeviceName, deviceCode).Scan(&out)
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

func (s *sDevDevice) Detail(ctx context.Context, key string) (out *model.DeviceOutput, err error) {
	err = dao.Device.Ctx(ctx).WithAll().Where(dao.Device.Columns().DeviceName, key).Scan(&out)
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

	//如果未设置，获取系统设置的默认超时时间 默认写30s 后期扩展从配置中取
	out.OnlineTimeout = 30

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
		m = m.WhereLike(dao.Device.Columns().DeviceName, "%"+keyWord+"%").WhereOrLike(dao.Device.Columns().DeviceName, "%"+keyWord+"%")
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
			WhereIn(dao.Device.Columns().DeviceName, offLineDeviceKeyList).
			Update()
	}

	return
}

func (s *sDevDevice) CheckDeviceBind(ctx context.Context, deviceName string) {
	s.checkDeviceInfo(ctx, deviceName)
}

func (s *sDevDevice) checkDeviceInfo(ctx context.Context, deviceName string) (device *entity.Device) {
	result, err := dao.Device.Ctx(ctx).
		Where(dao.Device.Columns().DeviceName, deviceName).
		Where(dao.Device.Columns().IsDelete, consts.DeleteOn).
		One()
	if err != nil {
		panic(err)
	}
	if result.IsEmpty() {
		panic(gerror.NewCode(enums.IllegalDevice))
	}
	udCnt, err := dao.UserDevice.Ctx(ctx).
		Where(dao.UserDevice.Columns().DeviceId, result.GMap().GetVar(dao.Device.Columns().DeviceId).Int64()).
		Count()
	if err != nil {
		panic(err)
	}
	if udCnt > 0 {
		panic(gerror.NewCode(enums.CarHaveMaster))
	}
	if err = result.Struct(&device); err != nil {
		panic(err)
	}
	return
}

func (s *sDevDevice) GetDeviceSecret(ctx context.Context, deviceName string) (
	out *model.DeviceSecretOutput) {
	device := s.checkDeviceInfo(ctx, deviceName)
	/*if device.Status != consts.DeviceStatueDisable {
		if err := gconv.Struct(device, &out); err != nil {
			panic(err)
		}
		out.MqttHostUrl = g.Cfg().MustGet(ctx, "aliyun.iot.mqtt.host").String()
		return
	}
	data, err := aliyun.RegisterDevice(ctx, device.ProductKey, device.DeviceName)
	if err != nil {
		panic(err)
	}
	if _, err = dao.Device.Ctx(ctx).
		Data(dao.Device.Columns().RegistryTime, gtime.Now(),
			dao.Device.Columns().Status, consts.DeviceStatueOffline,
			dao.Device.Columns().DeviceSecret, data.DeviceSecret,
		).Where(dao.Device.Columns().DeviceId, device.DeviceId).Update(); err != nil {
		if e := aliyun.DeleteDevice(ctx, device.ProductKey, device.DeviceName); e != nil {
			err = e
		}
		panic(err)
	}*/
	if err := gconv.Struct(device, &out); err != nil {
		panic(err)
	}
	if _, err := dao.Device.Ctx(ctx).
		Data(dao.Device.Columns().RegistryTime, gtime.Now(),
			dao.Device.Columns().Status, consts.DeviceStatueOffline,
		).Where(dao.Device.Columns().DeviceId, device.DeviceId).Update(); err != nil {
		panic(err)
	}
	out.MqttHostUrl = g.Cfg().MustGet(ctx, "aliyun.iot.mqtt.host").String()
	return
}

func (s *sDevDevice) ActivationSuccess(ctx context.Context, in *model.ActivationSuccessInput) {
	device := s.checkDeviceInfo(ctx, in.DeviceName)
	userId := service.BizCtx().Get(ctx).Data.Get(consts.ContextKeyUserId)
	if err := g.DB().Transaction(context.TODO(), func(ctx context.Context, tx gdb.TX) error {
		if _, err := dao.UserDevice.Ctx(ctx).
			Data(dao.UserDevice.Columns().IsSelect, consts.CarSelectNo).
			Where(dao.UserDevice.Columns().IsSelect, consts.CarSelectYes).
			Where(dao.UserDevice.Columns().UserId, userId).
			Update(); err != nil {
			return err
		}
		if _, err := dao.UserDevice.Ctx(ctx).Data(
			dao.UserDevice.Columns().Id, utility.GID.Generate().Int64(),
			dao.UserDevice.Columns().UserId, userId,
			dao.UserDevice.Columns().DeviceId, device.DeviceId,
			dao.UserDevice.Columns().IsSelect, consts.CarSelectYes,
			dao.UserDevice.Columns().UserDeviceType, consts.UserDeviceTypeMain,
			dao.UserDevice.Columns().CreateTime, gtime.Now(),
		).Insert(); err != nil {
			return err
		}
		if _, err := dao.Device.Ctx(ctx).Data(
			dao.Device.Columns().BluetoothSecretKey, gstr.Trim(in.BluetoothSecretKey),
			dao.Device.Columns().BluetoothAddress, gstr.Trim(in.BluetoothAddress),
			//dao.Device.Columns().SimId, gstr.Trim(in.SimID),
			dao.Device.Columns().UpdateTime, gtime.Now(),
		).Where(dao.Device.Columns().DeviceName, in.DeviceName).Update(); err != nil {
			return err
		}
		return nil
	}); err != nil {
		panic(err)
	}
	return
}

// GetLatestProperty 获取设备最新的属性值
func (s *sDevDevice) GetLatestProperty(ctx context.Context, key string) (list []model.DeviceLatestProperty, err error) {
	p, err := s.Get(ctx, key)
	if err != nil {
		return
	}
	if p.Status == model.DeviceStatusNoEnable {
		return
	}

	deviceTable := comm.DeviceTableName(p.DeviceName)

	tsdDb := tsd.DB()
	defer tsdDb.Close()

	for _, v := range p.TSL.Properties {
		ckey := comm.TsdColumnName(v.Key)

		// 获取属性最近有效值
		sql := "select last(?) as ? from ?"
		rs, err := tsdDb.GetTableDataOne(ctx, sql, ckey, ckey, deviceTable)
		if err != nil {
			return nil, err
		}
		value := rs[strings.ToLower(v.Key)]
		if value.IsEmpty() {
			continue
		}

		unit := ""
		if v.ValueType.TSLParam.Unit != nil {
			unit = *v.ValueType.TSLParam.Unit
		}

		pro := model.DeviceLatestProperty{
			Key:   v.Key,
			Name:  v.Name,
			Type:  v.ValueType.Type,
			Unit:  unit,
			Value: value,
		}
		list = append(list, pro)
	}
	return
}

// GetProperty 获取指定属性值
func (s *sDevDevice) GetProperty(ctx context.Context, input *model.DeviceGetPropertyInput) (list []model.DeviceLatestProperty, err error) {
	p, err := s.Detail(ctx, input.DeviceKey)
	if err != nil {
		return
	}
	if p.Status == model.DeviceStatusNoEnable {
		err = errors.New("设备未启用")
		return
	}

	deviceTable := comm.DeviceTableName(p.DeviceName)

	tsdDb := tsd.DB()
	defer tsdDb.Close()

	for _, i := range input.PropertyKeys {
		for _, v := range p.TSL.Properties {
			if v.Key != i {
				continue
			}

			ckey := comm.TsdColumnName(v.Key)

			// 获取属性最近有效值
			sql := "select last(?) as ? from ?"
			rs, err := tsdDb.GetTableDataOne(ctx, sql, ckey, ckey, deviceTable)
			if err != nil {
				return nil, err
			}
			value := rs[strings.ToLower(v.Key)]
			if value.IsEmpty() {
				continue
			}

			unit := ""
			if v.ValueType.TSLParam.Unit != nil {
				unit = *v.ValueType.TSLParam.Unit
			}

			pro := model.DeviceLatestProperty{
				Key:   v.Key,
				Name:  v.Name,
				Type:  v.ValueType.Type,
				Unit:  unit,
				Value: value,
			}
			list = append(list, pro)

			break
		}
	}
	return
}
