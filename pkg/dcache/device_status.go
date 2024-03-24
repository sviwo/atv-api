package dcache

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"sviwo/internal/consts"
	"sviwo/internal/model"
	"sviwo/internal/queues"
	"sviwo/internal/service"
	"sviwo/pkg/cache"
	"sviwo/pkg/iotModel"
	"time"
)

// UpdateStatus 更新设备状态
func UpdateStatus(ctx context.Context, device *model.DeviceOutput) {
	timeout := 0
	if device.OnlineTimeout > 0 {
		timeout = device.OnlineTimeout
	}
	if timeout == 0 {
		//设备默认超时时间 后期从配置里面取
		timeout = gconv.Int(30)
	}
	deviceStatus := GetDeviceStatus(ctx, device.DeviceName)

	if deviceStatus == consts.DeviceStatueOnline {
		//正常数据上报，更新设备在线的缓存时间
		_, err := cache.Instance().UpdateExpire(ctx, consts.DeviceStatusPrefix+device.DeviceName, time.Duration(timeout+1)*time.Second)
		if err != nil {
			g.Log().Debug(ctx, device.DeviceName, "更新设备在线的缓存时间失败")
		}

	} else {
		var deviceStatusLog = new(iotModel.DeviceStatusLog)
		deviceStatusLog.Status = 2
		deviceStatusLog.Timestamp = time.Now()
		deviceStatusLog.DeviceKey = device.DeviceName
		//设备首次上线，设置设备在线的缓存时间
		if err := cache.Instance().Set(ctx, consts.DeviceStatusPrefix+device.DeviceName, deviceStatusLog, time.Duration(timeout)*time.Second); err != nil {
			g.Log().Debug(ctx, device.DeviceName, "设置设备在线的缓存时间失败")
		}

		//添加延时下线消息判断
		go func(deviceKey string, timeOutSecond time.Duration) {
			if err := online(ctx, device); err != nil {
				g.Log().Debug(ctx, device.DeviceName, "设备上线处理失败")
			}
			for {
				ds := GetDeviceStatus(ctx, deviceKey)
				if ds == consts.DeviceStatueOffline { //离线
					break
				}
				time.Sleep(timeOutSecond * time.Second)
			}

			//设备下线
			err := offline(ctx, device)
			if err != nil {
				return
			}
		}(device.DeviceName, time.Duration(timeout)*time.Second)
	}
}

// pushDeviceStatus 推送设备状态
func pushDeviceStatus(deviceKey string, status int) {
	var deviceStatusLog = new(iotModel.DeviceStatusLog)
	deviceStatusLog.Status = status
	deviceStatusLog.Timestamp = time.Now()
	deviceStatusLog.DeviceKey = deviceKey
	data, _ := json.Marshal(deviceStatusLog)
	err := queues.DeviceStatusInfoUpdateWorker.Push(context.Background(), consts.QueueDeviceStatusInfoUpdate, data, 10)
	if err != nil {
		g.Log().Debug(context.Background(), "Push DeviceStatusInfoUpdateWorker: %v", err)
	}
}

// online 设备上线
func online(ctx context.Context, device *model.DeviceOutput) (err error) {
	//插入设备上线日志
	InertDeviceLog(ctx, consts.MsgTypeOnline, device.DeviceName, iotModel.DeviceOnlineMessage{
		Timestamp: time.Now().UnixMilli(),
		Desc:      "",
	})
	pushDeviceStatus(device.DeviceName, 2)

	//告警处理
	go func() {
		// 上线告警提醒
		data := iotModel.ReportStatusData{
			Status:     "online",
			CreateTime: gtime.Now().Unix(),
		}
		err = service.AlarmRule().Check(ctx, device.Product.ProductKey, device.DeviceName, consts.AlarmTriggerTypeOnline, data)
		if err != nil {
			g.Log().Errorf(ctx, "告警检测失败: %s", err.Error())
		}
	}()

	return
}

// offline 设备下线
func offline(ctx context.Context, device *model.DeviceOutput) (err error) {
	InertDeviceLog(ctx, consts.MsgTypeOffline, device.DeviceName, iotModel.DeviceOfflineMessage{
		Timestamp: time.Now().UnixMilli(),
		Desc:      "",
	})
	pushDeviceStatus(device.DeviceName, 1)

	// 离线告警提醒
	data := iotModel.ReportStatusData{
		Status:     "offline",
		CreateTime: gtime.Now().Unix(),
	}
	if err == nil {
		err = service.AlarmRule().Check(ctx, device.ProductKey, device.DeviceName, consts.AlarmTriggerTypeOffline, data)
	}
	return
}
