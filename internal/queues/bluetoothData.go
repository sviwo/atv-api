package queues

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"sviwo/internal/consts"
	"sviwo/pkg/worker"
)

var ScheduledBluetoothData = new(worker.Scheduled)

func ScheduledBluetoothDataRun() {
	ScheduledBluetoothData = worker.RegisterProcess(DeviceBluetoothData)
}

// DeviceBluetoothData 蓝牙数据
var DeviceBluetoothData = &qBluetoothData{}

type qBluetoothData struct{}

// GetTopic 主题
func (q *qBluetoothData) GetTopic() string {
	return consts.QueueBlueToothTopic
}

// Handle 处理消息
func (q *qBluetoothData) Handle(ctx context.Context, p worker.Payload) (err error) {
	g.Log().Debugf(ctx, "bluetooth_data_Handle: Payload(%s)", p.Payload)
	return
}
