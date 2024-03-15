package travelrecord

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"sviwo/internal/consts"
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

func (s sTravelRecord) GetTravelRecordList(ctx context.Context, in model.TravelRecordQueryInput) (
	total int, out []*model.TravelRecordOutput, err error) {
	userId := service.BizCtx().Get(ctx).Data.Get(consts.ContextKeyUserId)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.TravelRecord.Ctx(ctx)
		m = m.Where("user_id", userId)
		if in.DeviceId != 0 {
			m = m.Where("device_id", in.DeviceId)
		}
		m = m.Where("is_delete", consts.DeleteOn)
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

func (s sTravelRecord) Delete(ctx context.Context, in model.TravelRecordInput) (err error) {
	userId := service.BizCtx().Get(ctx).Data.Get(consts.ContextKeyUserId)
	_, err = dao.TravelRecord.Ctx(ctx).
		Data(g.Map{"is_delete": consts.DeleteYes}).
		Where(g.Map{"travel_record_id": in.TravelRecordId, "user_id": userId}).
		Update()
	return
}
