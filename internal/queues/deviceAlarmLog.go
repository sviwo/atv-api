package queues

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"sviwo/internal/consts"
	"sviwo/internal/model"
	"sviwo/internal/service"
	"sviwo/pkg/worker"
)

var ScheduledDeviceAlarmLog = new(worker.Scheduled)

func ScheduledDeviceAlarmLogRun() {
	ScheduledDeviceAlarmLog = worker.RegisterProcess(DeviceAlarmLog)
}

// DeviceAlarmLog 设备告警日志
var DeviceAlarmLog = &qDeviceAlarmLog{}

type qDeviceAlarmLog struct{}

// GetTopic 主题
func (q *qDeviceAlarmLog) GetTopic() string {
	return consts.QueueDeviceAlarmLogTopic
}

// Handle 处理消息
func (q *qDeviceAlarmLog) Handle(ctx context.Context, p worker.Payload) (err error) {
	g.Log().Debugf(ctx, "alarm_Handle: Payload(%s)", p.Payload)

	if p.Payload == nil || q.GetTopic() != p.Group {
		return nil
	}
	var data model.AlarmLogAddInput
	if err = json.Unmarshal(p.Payload, &data); err != nil {
		return err
	}
	//真正写日志
	_, err = service.AlarmLog().Add(ctx, &data)
	return
}
