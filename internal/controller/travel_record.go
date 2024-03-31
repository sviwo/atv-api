package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gutil"
	v1 "sviwo/api/v1"
	"sviwo/internal/model"
	"sviwo/internal/service"
)

var TravelRecord = cTravelRecord{}

type cTravelRecord struct{}

func (c cTravelRecord) GetTravelRecordList(ctx context.Context, req *v1.TravelRecordQueryReq) (res *v1.TravelRecordRes, err error) {
	tData := model.TravelRecordQueryInput{}
	if err = gconv.Struct(req, &tData); err != nil {
		panic(err)
	}
	total, out, err := service.TravelRecord().GetTravelRecordList(ctx, tData)
	if err != nil {
		panic(err)
	}
	res = new(v1.TravelRecordRes)
	res.Total = total
	res.CurrentPage = req.PageNum
	if !gutil.IsEmpty(out) {
		for _, tr := range out {
			recordBase := new(v1.TravelRecordBase)
			if err = gconv.Scan(tr, &recordBase); err != nil {
				panic(err)
			}
			recordBase.Duration = int(tr.EndTime.Sub(tr.StartTime).Minutes())
			res.Result = append(res.Result, recordBase)
		}
	}
	return
}

func (c cTravelRecord) Delete(ctx context.Context, req *v1.TravelRecordDeleteReq) (res *v1.EmptyFieldRes, err error) {
	if err = service.TravelRecord().Delete(ctx, req.TravelRecordId); err != nil {
		panic(err)
	}
	return
}
