package queues

import (
	"context"
	"sviwo/internal/consts"
	"testing"
)

func TestScheduledBluetoothData(t *testing.T) {
	b := make([]byte, 0)
	Run()
	ScheduledBluetoothData.Push(context.Background(), consts.QueueBlueToothTopic, b, 10)
}
