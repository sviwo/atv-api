package queues

import (
	"context"
	"sviwo/internal/consts"
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

	return
}
