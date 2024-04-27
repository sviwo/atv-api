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
		m = m.Where(dao.TravelRecord.Columns().UserId, userId)
		if in.DeviceId != 0 {
			m = m.Where(dao.TravelRecord.Columns().DeviceId, in.DeviceId)
		}
		m = m.Where(dao.TravelRecord.Columns().IsDelete, consts.DeleteOn)
		total, err = m.Count()
		if err != nil {
			panic(err)
		}

		err = m.Page(in.PageNum, in.PageSize).OrderDesc(dao.TravelRecord.Columns().CreateTime).Scan(&out)
		if err != nil {
			panic(err)
		}
	})
	return
}

func (s sTravelRecord) Delete(ctx context.Context, travelRecordId int64) (err error) {
	userId := service.BizCtx().Get(ctx).Data.Get(consts.ContextKeyUserId)
	_, err = dao.TravelRecord.Ctx(ctx).
		Data(g.Map{dao.TravelRecord.Columns().IsDelete: consts.DeleteYes}).
		Where(g.Map{dao.TravelRecord.Columns().TravelRecordId: travelRecordId, dao.TravelRecord.Columns().UserId: userId}).
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
		dao.TravelRecord.Columns().DeviceId: p.DeviceId,
		dao.TravelRecord.Columns().UserId:   userDevice.UserId,
		dao.TravelRecord.Columns().IsDelete: consts.DeleteOn,
		dao.TravelRecord.Columns().EndTime:  nil,
	}).One()
	if !trRes.IsEmpty() {
		dao.TravelRecord.Ctx(ctx).
			Data(g.Map{dao.TravelRecord.Columns().IsDelete: consts.DeleteYes}).
			Where(g.Map{dao.TravelRecord.Columns().TravelRecordId: trRes.GMap().Get(dao.TravelRecord.Columns().TravelRecordId)}).
			Update()
	}

	tsdDb := tsd.DB()
	defer tsdDb.Close()
	if err != nil {
		return
	}
	deviceTable := comm.DeviceTableName(p.DeviceName)
	ckey := comm.TsdColumnName(consts.GeoLocationStr)
	sql := "select last(?) as ? from ?"
	res, err := tsdDb.GetTableDataOne(ctx, sql, ckey, ckey, deviceTable)
	if err != nil {
		panic(err)
	}
	value := res[strings.ToLower(consts.GeoLocationStr)]
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
		dao.TravelRecord.Columns().DeviceId: p.DeviceId,
		dao.TravelRecord.Columns().UserId:   userDevice.UserId,
		dao.TravelRecord.Columns().IsDelete: consts.DeleteOn,
		dao.TravelRecord.Columns().EndTime:  nil,
	}).Scan(&travelRecord)
	if err != nil {
		panic(err)
	}
	tsdDb := tsd.DB()
	defer tsdDb.Close()
	// 获取最近的用车数据
	deviceTable := comm.DeviceTableName(p.DeviceName)
	geoLocation := comm.TsdColumnName(consts.GeoLocationStr)
	remainMile := comm.TsdColumnName(consts.RemainMileStr)
	electricity := comm.TsdColumnName(consts.ElectricityStr)
	sql := fmt.Sprintf("select last(%s) as %s,last(%s) as %s,last(%s) as %s from %s", geoLocation, geoLocation, remainMile, remainMile, electricity, electricity, deviceTable)
	res, err := tsdDb.GetTableDataOne(ctx, sql)
	geoValue := res[strings.ToLower(consts.GeoLocationStr)].String()
	remainMileValue := res[strings.ToLower(consts.RemainMileStr)].Int()
	electricityValue := res[strings.ToLower(consts.ElectricityStr)].Int()

	// 获取开机第一条数据
	firstSql := fmt.Sprintf("select %s,%s,%s from %s where ts > %d order by ts limit 1", geoLocation, remainMile, electricity, deviceTable, travelRecord.StartTime.Time.UnixMilli())
	fRes, err := tsdDb.GetTableDataOne(ctx, firstSql)
	fRemainMileValue := fRes[strings.ToLower(consts.RemainMileStr)].Int()
	fElectricityValue := fRes[strings.ToLower(consts.ElectricityStr)].Int()

	travelRecord.EndTime = gtime.Now()
	travelRecord.UpdateTime = gtime.Now()
	travelTimeInMinutes := travelRecord.EndTime.Sub(travelRecord.StartTime).Minutes()
	if travelTimeInMinutes <= 1 {
		dao.TravelRecord.Ctx(ctx).
			Data(g.Map{dao.TravelRecord.Columns().IsDelete: consts.DeleteYes}).
			Where(g.Map{dao.TravelRecord.Columns().TravelRecordId: travelRecord.TravelRecordId}).
			Update()
		return
	}
	travelRecord.MileageDriven = fRemainMileValue - remainMileValue
	if travelRecord.MileageDriven == 0 {
		dao.TravelRecord.Ctx(ctx).
			Data(g.Map{dao.TravelRecord.Columns().IsDelete: consts.DeleteYes}).
			Where(g.Map{dao.TravelRecord.Columns().TravelRecordId: travelRecord.TravelRecordId}).
			Update()
		return
	}
	travelRecord.Consumption = fElectricityValue - electricityValue
	avgSpeed := float64(travelRecord.MileageDriven) / float64(travelTimeInMinutes)
	integerAverageSpeed := math.Floor(avgSpeed * 60)
	//fmt.Println(integerAverageSpeed)
	travelRecord.EndPoint = geoValue
	//travelRecord.AvgSpeed = strconv.FormatFloat(integerAverageSpeed, 'f', 2, 32)
	travelRecord.AvgSpeed = integerAverageSpeed
	if _, err = dao.TravelRecord.Ctx(ctx).Data(travelRecord).Where(dao.TravelRecord.Columns().TravelRecordId, travelRecord.TravelRecordId).Update(); err != nil {
		panic(err)
	}
	return
}
