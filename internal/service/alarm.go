// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"sviwo/internal/model"
)

type (
	IAlarmLog interface {
		Detail(ctx context.Context, id uint64) (out *model.AlarmLogOutput, err error)
		Add(ctx context.Context, in *model.AlarmLogAddInput) (id uint64, err error)
		List(ctx context.Context, in *model.AlarmLogListInput) (out *model.AlarmLogListOutput, err error)
		// ClearLogByDays 按日期删除日志
		ClearLogByDays(ctx context.Context, days int) (err error)
	}
	IAlarmRule interface {
		// Check 告警检测
		Check(ctx context.Context, productKey string, deviceKey string, triggerType int, param any, subKey ...string) (err error)
	}
)

var (
	localAlarmLog   IAlarmLog
	localAlarmRule  IAlarmRule
)

func AlarmLog() IAlarmLog {
	if localAlarmLog == nil {
		panic("implement not found for interface IAlarmLog, forgot register?")
	}
	return localAlarmLog
}

func RegisterAlarmLog(i IAlarmLog) {
	localAlarmLog = i
}

func AlarmRule() IAlarmRule {
	if localAlarmRule == nil {
		panic("implement not found for interface IAlarmRule, forgot register?")
	}
	return localAlarmRule
}

func RegisterAlarmRule(i IAlarmRule) {
	localAlarmRule = i
}