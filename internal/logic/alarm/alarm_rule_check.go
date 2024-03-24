package alarm

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"sviwo/internal/consts"
	"sviwo/internal/model"
	"sviwo/internal/queues"
	"sviwo/pkg/cache"
	"sviwo/pkg/iotModel"
	"time"
)

// Check 告警检测
func (s *sAlarmRule) Check(ctx context.Context, productKey string, deviceKey string, triggerType int, param any, subKey ...string) (err error) {
	//g.Log().Debugf(ctx, "alarm_rule_check: deviceKey(%s), productKey(%s), triggerType(%d), param(%v)", deviceKey, productKey, triggerType, param)
	//重组param数据
	eventKey, data, err := s.getParamData(ctx, param)
	if data == nil || err != nil {
		return
	}
	logData, err := json.Marshal(param)
	if err != nil {
		g.Log().Errorf(ctx, "告警规则数据 -%s -%s - %s：%s", eventKey, productKey, deviceKey, err)
		return
	}
	// 填充告警触发时间
	if ts, ok := data["CreateTime"]; ok {
		switch tt := ts.(type) {
		case string:
			if t, err := time.Parse("2006-01-02 15:04:05", tt); err == nil {
				data["CreateTime"] = t.Unix()
			} else {
				data["CreateTime"] = time.Now().Unix()
			}
		case int, int32, int64, float32, float64:
			data["CreateTime"] = tt.(int64)
		}
	} else {
		data["CreateTime"] = time.Now().Unix()
	}
	go func() {
		// 写告警日志
		log := model.AlarmLogAddInput{
			Type:       1,
			Data:       string(logData),
			ProductKey: productKey,
			DeviceKey:  deviceKey,
		}
		if log.ProductKey == "" {
			return
		}
		// 写入队列
		logData, _ := json.Marshal(log)
		err = queues.ScheduledDeviceAlarmLog.Push(ctx, consts.QueueDeviceAlarmLogTopic, logData, 10)

		key := consts.DeviceAlarmLogPrefix + productKey + deviceKey
		err = cache.Instance().Set(ctx, key, log, 60*time.Second)
		if err != nil {
			return
		}
	}()
	return
}

// getParamData 获取Param数据 eventKey 事件标识
func (s *sAlarmRule) getParamData(ctx context.Context, param any) (eventKey string, res map[string]any, err error) {
	var (
		data = make(map[string]any) // 上传数据
	)
	// 上传数据重组
	switch pd := param.(type) {
	case model.ReportPropertyData:
		for k, v := range pd {
			vv := gconv.String(v.Value)
			if gstr.IsNumeric(vv) {
				data[k] = gconv.Float64(v.Value)
			} else if gt, err := gtime.StrToTime(vv); err == nil {
				data[k] = gt.Unix()
			} else {
				data[k] = v.Value
			}
			data[k+"_time"] = v.CreateTime
		}
	case model.ReportEventData:
		for k, v := range pd.Param.Value {
			vv := gconv.String(v)
			if gstr.IsNumeric(vv) {
				data[k] = gconv.Float64(v)
			} else if gt, err := gtime.StrToTime(vv); err == nil {
				data[k] = gt.Unix()
			} else {
				data[k] = v
			}
		}
		data["CreateTime"] = pd.Param.CreateTime
		eventKey = pd.Key
	case iotModel.ReportStatusData:
		data["Status"] = pd.Status
		data["CreateTime"] = pd.CreateTime
	default:
		return "", nil, gerror.New("数据格式错误")
	}
	return eventKey, data, nil
}
