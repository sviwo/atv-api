package service

import (
	"context"
	"sviwo/internal/model"
)

type (
	ITravelRecord interface {
		GetTravelRecordList(ctx context.Context, in model.TravelRecordQueryInput) (total int, out []*model.TravelRecordOutput, err error)
		Delete(ctx context.Context, travelRecordId int64) (err error)
		CreateOnline(ctx context.Context, in model.TravelRecordOnline) (err error)
		UpdateOnlineToOffline(ctx context.Context, in model.TravelRecordOnline) (err error)
	}
)

var (
	localTravelRecord ITravelRecord
)

func TravelRecord() ITravelRecord {
	if localTravelRecord == nil {
		panic("implement not found for interface ITravelRecord, forgot register?")
	}
	return localTravelRecord
}

func RegisterTravelRecord(i ITravelRecord) {
	localTravelRecord = i
}
