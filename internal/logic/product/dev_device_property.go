package product

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"sviwo/internal/consts"
	"sviwo/internal/consts/enums"
	"sviwo/internal/dao"
	"sviwo/internal/model"
	"sviwo/internal/model/entity"
	dset "sviwo/internal/network/core/logic/model/down/property/set"
	"sviwo/internal/service"
	"sviwo/pkg/aliyun"
	"sviwo/pkg/dcache"
	"sviwo/pkg/iotModel/topicModel"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

type sDevDeviceProperty struct{}

func init() {
	service.RegisterDevDeviceProperty(devDeviceProperty())
}

func devDeviceProperty() *sDevDeviceProperty {
	return &sDevDeviceProperty{}
}

// Set 设备属性设置
func (s *sDevDeviceProperty) Set(ctx context.Context, in *model.DevicePropertyInput) (out *model.DevicePropertyOutput, err error) {
	device, err := dcache.GetDeviceDetailInfo(in.DeviceKey)
	if dcache.GetDeviceStatus(ctx, in.DeviceKey) != model.DeviceStatusOn {
		err = gerror.New("设备不在线")
		return
	}

	var params []byte
	if len(in.Params) > 0 {
		if params, err = json.Marshal(in.Params); err != nil {
			return
		}
	}
	request := topicModel.TopicDownHandlerData{
		DeviceDetail: device,
		PayLoad:      params,
	}

	out = &model.DevicePropertyOutput{}
	if out.Data, err = dset.PropertySet(ctx, request); err != nil {
		return
	}

	// 写日志
	logData := &model.TdLogAddInput{
		Ts:      gtime.Now(),
		Device:  in.DeviceKey,
		Type:    consts.MsgTypePropertyWrite,
		Content: string(params),
	}
	err = service.TdLogTable().Insert(ctx, logData)

	return
}

func (s *sDevDeviceProperty) GetDeviceSecret(ctx context.Context, deviceCode string) (
	out *model.DeviceSecretOutput) {
	count, err := dao.Device.Ctx(ctx).
		Where(dao.Device.Columns().DeviceName, deviceCode).
		Where(dao.Device.Columns().IsDelete, consts.DeleteOn).
		Count()
	if err != nil {
		panic(err)
	}
	if count == 0 {
		panic(gerror.NewCode(enums.IllegalDevice))
		return
	}

	device := new(entity.Device)
	if err := dao.Device.Ctx(ctx).
		Where(dao.Device.Columns().DeviceName, deviceCode).
		Where(dao.Device.Columns().IsDelete, consts.DeleteOn).
		Scan(&device); err != nil {
		panic(err)
	}
	if device.Status != consts.DeviceStatueRegister {
		if err := gconv.Struct(device, &out); err != nil {
			panic(err)
		}
		return
	}
	if err := g.DB().Transaction(context.TODO(), func(ctx context.Context, tx gdb.TX) error {
		data, err := aliyun.RegisterDevice(ctx, device.ProductKey, device.DeviceName)
		if err != nil {
			return err
		}
		if _, err := dao.Device.Ctx(ctx).
			Data(dao.Device.Columns().RegistryTime, gtime.Now(),
				dao.Device.Columns().Status, consts.DeviceStatueDisable,
				dao.Device.Columns().DeviceSecret, data.DeviceSecret,
			).Where(dao.Device.Columns().DeviceId, device.DeviceId).Update(); err != nil {
			return err
		}
		if err := gconv.Struct(data, &out); err != nil {
			return err
		}
		return nil
	}); err != nil {
		if e := aliyun.DeleteDevice(ctx, device.ProductKey, device.DeviceName); err != nil {
			panic(e)
		}
		panic(err)
	}
	return
}
