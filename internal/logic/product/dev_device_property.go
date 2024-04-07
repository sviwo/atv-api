package product

import (
	"context"
	"encoding/json"
	"sviwo/internal/consts"
	"sviwo/internal/consts/enums"
	"sviwo/internal/model"
	dset "sviwo/internal/network/core/logic/model/down/property/set"
	"sviwo/internal/service"
	"sviwo/pkg/dcache"
	"sviwo/pkg/gpool"
	"sviwo/pkg/iotModel/topicModel"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

type sDevDeviceProperty struct{}

// 控制携程数量
var gSetPool = gpool.NewGPool(5000)

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
		panic(gerror.NewCode(enums.DeviceOffline))
	}
	gSetPool.Go(func(ctx context.Context) error {
		var params []byte
		if len(in.Params) > 0 {
			if params, err = json.Marshal(in.Params); err != nil {
				return err
			}
		}
		request := topicModel.TopicDownHandlerData{
			DeviceDetail: device,
			PayLoad:      params,
		}
		if _, err = dset.PropertySet(ctx, request); err != nil {
			return err
		}
		// 写日志
		logData := &model.TdLogAddInput{
			Ts:      gtime.Now(),
			Device:  in.DeviceKey,
			Type:    consts.MsgTypePropertyWrite,
			Content: string(params),
		}
		err = service.TdLogTable().Insert(ctx, logData)
		return nil
	})
	return
}
