package alarm

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"sviwo/internal/dao"
	"sviwo/internal/model"
	"sviwo/internal/model/do"
	"sviwo/internal/service"
)

type sAlarmLog struct{}

func init() {
	service.RegisterAlarmLog(alarmLogNew())
}

func alarmLogNew() *sAlarmLog {
	return &sAlarmLog{}
}

func (s *sAlarmLog) Detail(ctx context.Context, id uint64) (out *model.AlarmLogOutput, err error) {
	err = dao.AlarmLog.Ctx(ctx).WithAll().Where(dao.AlarmLog.Columns().Id, id).Scan(&out)
	return
}

func (s *sAlarmLog) Add(ctx context.Context, in *model.AlarmLogAddInput) (id uint64, err error) {
	rs, err := dao.AlarmLog.Ctx(ctx).Data(do.AlarmLog{
		Type:        in.Type,
		Data:        in.Data,
		ProductKey:  in.ProductKey,
		DeviceKey:   in.DeviceKey,
		Status:      0,
		CreatedTime: gtime.Now(),
		Content:     "",
	}).Insert()
	if err != nil {
		return
	}
	newId, err := rs.LastInsertId()
	id = uint64(newId)
	return
}

func (s *sAlarmLog) List(ctx context.Context, in *model.AlarmLogListInput) (out *model.AlarmLogListOutput, err error) {
	out = new(model.AlarmLogListOutput)
	c := dao.AlarmLog.Columns()
	m := dao.AlarmLog.Ctx(ctx).WithAll().
		OrderDesc(c.Id)

	if len(in.DateRange) > 0 {
		m = m.WhereBetween(c.CreatedTime, in.DateRange[0], in.DateRange[1])
	}

	if len(in.Status) > 0 {
		m = m.Where(dao.AlarmLog.Columns().Status, in.Status)
	}

	out.Total, _ = m.Count()
	out.CurrentPage = in.PageNum
	err = m.Page(in.PageNum, in.PageSize).Scan(&out.List)
	return
}

// ClearLogByDays 按日期删除日志
func (s *sAlarmLog) ClearLogByDays(ctx context.Context, days int) (err error) {
	_, err = dao.AlarmLog.Ctx(ctx).Delete("to_days(now())-to_days(`created_at`) > ?", days+1)
	return
}
