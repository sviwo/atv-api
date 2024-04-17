package travelrecord

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"math"
	"strings"
	"sviwo/internal/consts"
	"sviwo/internal/dao"
	"sviwo/internal/model"
	"sviwo/internal/model/do"
	"sviwo/internal/model/entity"
	"sviwo/internal/service"
	"sviwo/pkg/tsd"
	"sviwo/pkg/tsd/comm"
	"sviwo/pkg/utility"
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

func (s sTravelRecord) Delete(ctx context.Context, travelRecordId int64) (err error) {
	userId := service.BizCtx().Get(ctx).Data.Get(consts.ContextKeyUserId)
	_, err = dao.TravelRecord.Ctx(ctx).
		Data(g.Map{"is_delete": consts.DeleteYes}).
		Where(g.Map{"travel_record_id": travelRecordId, "user_id": userId}).
		Update()
	return
}

func (s sTravelRecord) CreateOnline(ctx context.Context, in model.TravelRecordOnline) (err error) {
	//根据物模型获取所有属性数据 根据情况选择
	p, err := service.DevDevice().Detail(ctx, in.DeviceName)
	if err != nil {
		panic(err)
	}
	result, err := dao.UserDevice.Ctx(ctx).Fields(dao.UserDevice.Columns().UserId).
		Where(dao.UserDevice.Columns().DeviceId, p.DeviceId).
		Where(dao.UserDevice.Columns().UserDeviceType, consts.UserDeviceTypeMain).One()
	if err != nil {
		panic(err)
	}
	userDevice := new(entity.UserDevice)
	result.Struct(&userDevice)

	trRes, err := dao.TravelRecord.Ctx(ctx).Where(g.Map{
		"device_id": p.DeviceId,
		"user_id":   userDevice.UserId,
		"is_delete": consts.DeleteOn,
		"end_time":  nil,
	}).One()
	if trRes != nil {
		return
	}

	tsdDb := tsd.DB()
	defer tsdDb.Close()
	if err != nil {
		return
	}
	deviceTable := comm.DeviceTableName(p.DeviceName)
	ckey := comm.TsdColumnName(consts.GeoLocation)
	sql := "select last(?) as ? from ?"
	res, err := tsdDb.GetTableDataOne(ctx, sql, ckey, ckey, deviceTable)
	if err != nil {
		panic(err)
	}
	value := res[strings.ToLower(consts.GeoLocation)]
	travelRecord := do.TravelRecord{
		TravelRecordId: utility.GID.Generate().Int64(),
		UserId:         userDevice.UserId,
		DeviceId:       p.DeviceId,
		StartPoint:     value.String(),
		StartTime:      gtime.Now(),
		CreateTime:     gtime.Now(),
	}

	if _, err = dao.TravelRecord.Ctx(ctx).Data(travelRecord).Insert(); err != nil {
		panic(err)
	}
	return
}

func (s sTravelRecord) UpdateOnlineToOffline(ctx context.Context, in model.TravelRecordOnline) (err error) {
	p, err := service.DevDevice().Detail(ctx, in.DeviceName)
	if err != nil {
		panic(err)
	}

	result, err := dao.UserDevice.Ctx(ctx).Fields(dao.UserDevice.Columns().UserId).
		Where(dao.UserDevice.Columns().DeviceId, p.DeviceId).
		Where(dao.UserDevice.Columns().UserDeviceType, consts.UserDeviceTypeMain).One()
	if err != nil {
		panic(err)
	}
	userDevice := new(entity.UserDevice)
	result.Struct(&userDevice)

	travelRecord := entity.TravelRecord{}
	err = dao.TravelRecord.Ctx(ctx).Where(g.Map{
		"device_id": p.DeviceId,
		"user_id":   userDevice.UserId,
		"is_delete": consts.DeleteOn,
		"end_time":  nil,
	}).Scan(&travelRecord)
	if err != nil {
		panic(err)
	}
	tsdDb := tsd.DB()
	defer tsdDb.Close()
	// 获取最近的用车数据
	deviceTable := comm.DeviceTableName(p.DeviceName)
	geoLocation := comm.TsdColumnName(consts.GeoLocation)
	remainMile := comm.TsdColumnName(consts.RemainMile)
	electricity := comm.TsdColumnName(consts.Electricity)
	sql := fmt.Sprintf("select last(%s) as %s,last(%s) as %s,last(%s) as %s from %s", geoLocation, geoLocation, remainMile, remainMile, electricity, electricity, deviceTable)
	res, err := tsdDb.GetTableDataOne(ctx, sql)
	geoValue := res[strings.ToLower(consts.GeoLocation)].String()
	remainMileValue := res[strings.ToLower(consts.RemainMile)].Int()
	electricityValue := res[strings.ToLower(consts.Electricity)].Int()

	// 获取开机开机第一条数据
	firstSql := fmt.Sprintf("select %s,%s,%s from %s where ts > %d order by ts limit 1", geoLocation, remainMile, electricity, deviceTable, travelRecord.StartTime.Time.UnixMilli())
	fRes, err := tsdDb.GetTableDataOne(ctx, firstSql)
	fRemainMileValue := fRes[strings.ToLower(consts.RemainMile)].Int()
	fElectricityValue := fRes[strings.ToLower(consts.Electricity)].Int()

	travelRecord.EndTime = gtime.Now()
	travelRecord.UpdateTime = gtime.Now()
	travelTimeInMinutes := travelRecord.EndTime.Sub(travelRecord.StartTime).Minutes()
	if travelTimeInMinutes <= 1 {
		travelTimeInMinutes = 1
	}
	travelRecord.MileageDriven = fRemainMileValue - remainMileValue
	travelRecord.Consumption = fElectricityValue - electricityValue
	avgSpeed := float64(travelRecord.MileageDriven) / float64(travelTimeInMinutes)
	integerAverageSpeed := math.Floor(avgSpeed * 60)
	//fmt.Println(integerAverageSpeed)
	travelRecord.EndPoint = geoValue
	//travelRecord.AvgSpeed = strconv.FormatFloat(integerAverageSpeed, 'f', 2, 32)
	travelRecord.AvgSpeed = integerAverageSpeed
	if _, err = dao.TravelRecord.Ctx(ctx).Data(travelRecord).Where("travel_record_id", travelRecord.TravelRecordId).Update(); err != nil {
		panic(err)
	}
	return
}
