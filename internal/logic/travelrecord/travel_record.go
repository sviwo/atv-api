package travelrecord

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"sviwo/internal/dao"
	"sviwo/internal/model"
	"sviwo/internal/service"
)

func init() {
	service.RegisterTravelRecord(New())
}

func New() *sTravelRecord {
	return &sTravelRecord{}
}

type sTravelRecord struct{}

func (s sTravelRecord) Delete(ctx context.Context, in model.TravelRecordInput) {
	dao.TravelRecord.Ctx(ctx).Data(g.Map{"is_delete": 0}).Where("ravel_record_id", in.TravelRecordId).Update()
}

func (s sTravelRecord) GetTravelRecordList(ctx context.Context, in model.TravelRecordQueryInput) (total int, out []*model.TravelRecordOutput, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.TravelRecord.Ctx(ctx)
		if in.UserId != 0 {
			m = m.Where("user_id", in.UserId)
		}
		if in.CarId != 0 {
			m = m.Where("car_id", in.CarId)
		}
		m.Where("is_delete", 0)
		total, err = m.Count()
		if err != nil {
			panic(err)
		}

		err = m.Page(in.PageNum, in.PageSize).OrderDesc("create_time").Scan(&out)
		if err != nil {
			panic(err)
		}
	})
	return
}
