package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "sviwo/api/v1"
	"sviwo/internal/model"
	"sviwo/internal/service"
)

var TravelRecord = cTravelRecord{}

type cTravelRecord struct{}

func (c cTravelRecord) GetTravelRecordList(ctx context.Context, req *v1.TravelRecordQueryReq) (res *v1.TravelRecordRes, err error) {
	tData := model.TravelRecordQueryInput{}
	err = gconv.Struct(req, &tData)
	if err != nil {
		panic(err)
	}
	total, out, err := service.TravelRecord().GetTravelRecordList(ctx, tData)
	if err != nil {
		panic(err)
	}
	res = new(v1.TravelRecordRes)
	res.Total = total
	res.CurrentPage = req.PageNum
	res.Result = out
	return
}

func (c cTravelRecord) Delete(ctx context.Context, req *v1.TravelRecordReq) (res *v1.CommonRes, err error) {
	tData := model.TravelRecordInput{}
	err = gconv.Struct(req, &tData)
	if err != nil {
		panic(err)
	}
	service.TravelRecord().Delete(ctx, tData)
	return
}
