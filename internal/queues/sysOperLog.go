package queues

import (
	"context"
	"sviwo/internal/consts"
	"sviwo/pkg/worker"
)

var ScheduledSysOperLog = new(worker.Scheduled)

func ScheduledSysOperLogRun() {
	ScheduledSysOperLog = worker.RegisterProcess(SysOperLog)
}

// SysOperLog 系统日志
var SysOperLog = &qSysOperLog{}

type qSysOperLog struct{}

// GetTopic 主题
func (q *qSysOperLog) GetTopic() string {
	return consts.QueueRequestLogTopic
}

// Handle 处理消息
func (q *qSysOperLog) Handle(ctx context.Context, p worker.Payload) (err error) {

	return nil
}
