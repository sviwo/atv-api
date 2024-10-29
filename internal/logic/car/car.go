package car

import (
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"fmt"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gbase64"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/gogf/gf/v2/util/gutil"
	"golang.org/x/sync/errgroup"
	"sviwo/internal/consts"
	"sviwo/internal/consts/enums"
	"sviwo/internal/dao"
	"sviwo/internal/model"
	"sviwo/internal/model/do"
	"sviwo/internal/model/entity"
	"sviwo/internal/service"
	"sviwo/pkg/utility"
)

func init() {
	service.RegisterCar(New())
}

func New() *sCar {
	return &sCar{}
}

type sCar struct{}

func (s sCar) GetCarList(ctx context.Context) (out []*model.QueryCarOutput) {
	var (
		udTable = dao.UserDevice.Table()
		udCls   = dao.UserDevice.Columns()
		dTable  = dao.Device.Table()
		dCls    = dao.Device.Columns()
	)
	if err := dao.UserDevice.Ctx(ctx).FieldsPrefix(udTable, udCls.IsSelect, udCls.UserDeviceType).
		FieldsPrefix(dTable, dCls.DeviceId, dCls.Nickname, dCls.DeviceName).
		LeftJoinOnField(dTable, dCls.DeviceId).
		WherePrefix(dTable, dCls.IsDelete, consts.DeleteOn).
		WherePrefix(udTable, udCls.UserId, service.BizCtx().Get(ctx).Data.Get(consts.ContextKeyUserId)).
		OrderDesc(udCls.IsSelect).Scan(&out); err != nil {
		panic(err)
	}
	if out == nil {
		return
	}
	keys := make([]string, 0)
	keys = append(keys, consts.MileageStr)
	for _, ot := range out {
		res, err := service.DevDevice().GetProperty(ctx, &model.DeviceGetPropertyInput{
			DeviceKey:    ot.DeviceName,
			PropertyKeys: keys,
		})
		if err != nil {
			panic(err)
		}
		if !gutil.IsEmpty(res) {
			ot.Mileage = res[0].Value.Float32()
		}
	}
	return
}

func (s sCar) GetCarDetail(ctx context.Context, deviceId *int64) (out *model.UserDeviceOutput) {
	withContext, _ := errgroup.WithContext(ctx)
	defer func() {
		if err := withContext.Wait(); err != nil {
			panic(err)
		}
	}()
	withContext.Go(func() error {
		device := s.findDeviceInfo(ctx, deviceId)
		if device == nil {
			return nil
		}
		if err := gconv.Scan(device, &out); err != nil {
			return err
		}
		out.ActivateTime = device.ActivateTime.Format("n/d/Y")
		out.WarrantyTime = device.ActivateTime.AddDate(1, 0, 0).Format("n/d/Y")
		out.UserCarKeyList = s.findCarKeyList(ctx, device.DeviceId)
		keys := make([]string, 0)
		keys = append(keys, consts.LimitStr)
		keys = append(keys, consts.MileageStr)
		res, err := service.DevDevice().GetProperty(ctx, &model.DeviceGetPropertyInput{
			DeviceKey:    device.DeviceName,
			PropertyKeys: keys,
		})

		if err != nil {
			return err
		}
		for _, re := range res {
			switch re.Key {
			case consts.LimitStr:
				out.TopSpeedHour = re.Value.Int()
			case consts.MileageStr:
				out.Mileage = re.Value.Float32()
			default:
				break
			}
		}
		return nil
	})
	withContext.Go(func() error {
		userDevice := s.findUserDevice(ctx, nil)
		if userDevice != nil {
			if err := gconv.Scan(userDevice, &out); err != nil {
				return err
			}
		}
		return nil
	})
	return
}

func (s sCar) findCarKeyList(ctx context.Context, deviceId uint64) (userCarKeys []model.UserCarKeyOutput) {
	all, err := dao.UserDevice.Ctx(ctx).
		Where(dao.UserDevice.Columns().DeviceId, deviceId).
		Where(dao.UserDevice.Columns().UserDeviceType, consts.UserDeviceChild).
		All()
	if err != nil {
		panic(err)
	}
	if !all.IsEmpty() {
		userCarKeys = make([]model.UserCarKeyOutput, all.Len())
		for i, record := range all {
			userCarKeys[i].UserDeviceId = record.GMap().GetVar(dao.UserDevice.Columns().Id).Int64()
		}
		users, err := dao.User.Ctx(ctx).
			WhereIn(dao.User.Columns().UserId, all.Array(dao.UserDevice.Columns().UserId)).
			All()
		if err != nil {
			panic(err)
		}
		for i, user := range users {
			userCarKeys[i].Name = user.GMap().GetVar(dao.User.Columns().LastName).String() +
				" " + user.GMap().GetVar(dao.User.Columns().FirstName).String()
			headImg := user.GMap().GetVar(dao.User.Columns().HeadImg)
			if !gutil.IsEmpty(headImg) {
				userCarKeys[i].HeadImg = g.Cfg().MustGet(ctx, "aliyun.oss.fileUrlPrefix").String() +
					headImg.String()
			}
		}
	}
	return
}

func (s sCar) SwitchCar(ctx context.Context, deviceId int64) {
	if err := g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		userId := service.BizCtx().Get(ctx).Data.Get(consts.ContextKeyUserId)
		result, err := dao.UserDevice.Ctx(ctx).Data(dao.UserDevice.Columns().IsSelect, consts.CarSelectNo).
			Where(dao.UserDevice.Columns().UserId, userId).
			Where(dao.UserDevice.Columns().IsSelect, consts.CarSelectYes).
			WhereNot(dao.UserDevice.Columns().DeviceId, deviceId).
			Update()
		if err != nil {
			return err
		}
		if affected, _ := result.RowsAffected(); affected > 0 {
			if _, err = dao.UserDevice.Ctx(ctx).Data(dao.UserDevice.Columns().IsSelect, consts.CarSelectYes).
				Where(dao.UserDevice.Columns().UserId, userId).
				Where(dao.UserDevice.Columns().DeviceId, deviceId).
				Update(); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		panic(err)
	}
}

func (s sCar) EditCarNickname(ctx context.Context, nickname string) {
	userDevice := s.findUserDevice(ctx, nil)
	if userDevice == nil {
		panic(gerror.NewCode(enums.IllegalOperation))
	}
	if _, err := dao.Device.Ctx(ctx).Data(dao.Device.Columns().Nickname, nickname).
		Where(dao.Device.Columns().DeviceId, userDevice.DeviceId).Update(); err != nil {
		panic(err)
	}
}

func (s sCar) RemoveCar(ctx context.Context, userDeviceId, deviceId *int64) {
	if userDeviceId != nil {
		if _, err := dao.UserDevice.Ctx(ctx).
			Where(dao.UserDevice.Columns().Id, userDeviceId).
			Where(dao.UserDevice.Columns().UserDeviceType, consts.UserDeviceChild).
			Delete(); err != nil {
			panic(err)
		}
		return
	}
	userDevice := s.findUserDevice(ctx, deviceId)
	if userDevice == nil {
		panic(gerror.NewCode(enums.CarNotExists))
	}
	if userDevice.UserDeviceType == consts.UserDeviceTypeMain {
		panic(gerror.NewCode(enums.UnbindCarError))
	}
	userId := service.BizCtx().Get(ctx).Data.Get(consts.ContextKeyUserId)
	if _, err := dao.UserDevice.Ctx(ctx).
		Where(dao.UserDevice.Columns().UserId, userId).
		Where(dao.UserDevice.Columns().DeviceId, userDevice.DeviceId).
		Delete(); err != nil {
		panic(err)
	}

	count, err := dao.UserDevice.Ctx(ctx).
		Where(dao.UserDevice.Columns().IsSelect, consts.CarSelectYes).
		Where(dao.UserDevice.Columns().UserId, userId).
		Count()
	if err != nil {
		panic(err)
	}
	if count == 0 {
		if _, err = dao.UserDevice.Ctx(ctx).
			Data(dao.UserDevice.Columns().IsSelect, consts.CarSelectYes).
			Where(dao.UserDevice.Columns().UserId, userId).
			Limit(1).
			Update(); err != nil {
			panic(err)
		}
	}
}

func (s sCar) GetCarKey(ctx context.Context) (carKey string) {
	userId := service.BizCtx().Get(ctx).Data.Get(consts.ContextKeyUserId)
	utility.MethodReqLimit(ctx, "GetCarKey", userId, 5)

	result, err := dao.UserDevice.Ctx(ctx).Fields(dao.UserDevice.Columns().DeviceId).
		Where(dao.UserDevice.Columns().UserId, userId).
		Where(dao.UserDevice.Columns().UserDeviceType, consts.UserDeviceTypeMain).
		Where(dao.UserDevice.Columns().IsSelect, consts.CarSelectYes).
		One()
	if err != nil {
		panic(err)
	}
	if result.IsEmpty() {
		panic(gerror.NewCode(enums.ShareCarKeyError))
	}
	deviceId := result.GMap().GetVar(dao.Device.Columns().DeviceId).Int64()
	s.checkCarKeyLimit(ctx, deviceId)

	val, err := g.Redis().Get(ctx, fmt.Sprintf(consts.RedisCarKey, gconv.String(deviceId)))
	if err != nil {
		panic(err)
	}
	if !val.IsEmpty() {
		carKey = val.String()
	} else {
		carKey = consts.CarKeyPrefix + grand.S(32)
		if err = g.Redis().SetEX(ctx, fmt.Sprintf(consts.RedisCarKey, carKey), deviceId, 60*60); err != nil {
			panic(err)
		}
		if err = g.Redis().SetEX(ctx, fmt.Sprintf(consts.RedisCarKey, gconv.String(deviceId)),
			carKey, 60*60); err != nil {
			panic(err)
		}
	}
	return
}

func (s sCar) checkCarKeyLimit(ctx context.Context, deviceId int64) {
	cot, err := dao.UserDevice.Ctx(ctx).
		Count(dao.UserDevice.Columns().DeviceId, deviceId)
	if err != nil {
		panic(err)
	}
	if cot >= 4 {
		panic(gerror.NewCode(enums.CarKeyLimitError))
	}
}

func (s sCar) InviteBindCar(ctx context.Context, carKey string) {
	deviceId, err := g.Redis().Get(
		ctx, fmt.Sprintf(consts.RedisCarKey, carKey),
	)
	if err != nil {
		panic(err)
	}
	if deviceId.IsEmpty() {
		panic(gerror.NewCode(enums.CarKeyInvalidError))
	}
	s.checkCarKeyLimit(ctx, deviceId.Int64())
	userId := service.BizCtx().Get(ctx).Data.Get(consts.ContextKeyUserId)
	cot, err := dao.UserDevice.Ctx(ctx).
		Where(dao.UserDevice.Columns().DeviceId, deviceId).
		Where(dao.UserDevice.Columns().UserId, userId).
		Count()
	if err != nil {
		panic(err)
	}
	if cot > 0 {
		panic(gerror.NewCode(enums.BindCarError))
	}

	userDevice := do.UserDevice{
		Id:             utility.GID.Generate().Int64(),
		UserId:         userId,
		DeviceId:       deviceId,
		UserDeviceType: consts.UserDeviceChild,
		CreateTime:     gtime.Now(),
	}
	count, err := dao.UserDevice.Ctx(ctx).Count(dao.UserDevice.Columns().UserId, userId)
	if err != nil {
		panic(err)
	}
	if count == 0 {
		userDevice.IsSelect = consts.CarSelectYes
	}
	if _, err = dao.UserDevice.Ctx(ctx).Data(userDevice).Insert(); err != nil {
		panic(err)
	}
	if _, err = g.Redis().Del(ctx, fmt.Sprintf(consts.RedisCarKey, carKey)); err != nil {
		panic(err)
	}
	if _, err = g.Redis().Del(ctx, fmt.Sprintf(consts.RedisCarKey, deviceId.String())); err != nil {
		panic(err)
	}
}

func (s sCar) CtlCar(ctx context.Context, instructions int) {
	device := s.findDeviceInfo(ctx, nil)
	if device == nil {
		panic(gerror.NewCode(enums.CarNotExists))
	}
	mp := make(map[string]any)
	switch instructions {
	case consts.LightCode:
		mp = map[string]any{consts.LightStr: 1}
	case consts.SpeakerCode:
		mp = map[string]any{consts.SpeakerStr: 1}
	default:
		panic(gerror.NewCode(enums.IllegalOperation))
	}
	in := &model.DevicePropertyInput{
		DeviceKey: device.DeviceName,
		Params:    mp,
	}
	if _, err := service.DevDeviceProperty().Set(ctx, in); err != nil {
		panic(err)
	}
}

func (s sCar) findDeviceInfo(ctx context.Context, deviceId *int64) (device *entity.Device) {
	if deviceId == nil {
		userDevice := s.findUserDevice(ctx, nil)
		if userDevice == nil {
			return
		}
		deviceId = &userDevice.DeviceId
	}
	if err := dao.Device.Ctx(ctx).
		Where(dao.Device.Columns().DeviceId, &deviceId).
		Where(dao.Device.Columns().IsDelete, consts.DeleteOn).Scan(&device); err != nil {
		panic(err)
	}
	return
}

func (s sCar) CtlSwitchDT(ctx context.Context, in model.CtlSwitchDTInput) {
	if _, err := dao.UserDevice.Ctx(ctx).
		Data(
			dao.UserDevice.Columns().DrivingMode, in.DrivingModeType,
			dao.UserDevice.Columns().UpdateTime, gtime.Now(),
		).Where(dao.UserDevice.Columns().UserId, service.BizCtx().Get(ctx).Data.Get(consts.ContextKeyUserId)).
		Where(dao.UserDevice.Columns().IsSelect, consts.CarSelectYes).
		Update(); err != nil {
		panic(err)
	}
}

func (s sCar) CtlSwitchERT(ctx context.Context, in model.CtlSwitchERTInput) {
	if _, err := dao.UserDevice.Ctx(ctx).
		Data(
			dao.UserDevice.Columns().EnergyRecovery, in.EnergyRecoveryType,
			dao.UserDevice.Columns().UpdateTime, gtime.Now(),
		).Where(dao.UserDevice.Columns().UserId, service.BizCtx().Get(ctx).Data.Get(consts.ContextKeyUserId)).
		Where(dao.UserDevice.Columns().IsSelect, consts.CarSelectYes).
		Update(); err != nil {
		panic(err)
	}
}

func (s sCar) EnabledMobileKey(ctx context.Context) {
	userDevice := s.findUserDevice(ctx, nil)
	if userDevice == nil || consts.UserDeviceChild == userDevice.UserDeviceType {
		panic(gerror.NewCode(enums.IllegalOperation))
	}
	var mobileKey = true
	if userDevice.MobileKey {
		mobileKey = consts.CarMobileKeyNo
	}
	if _, err := dao.UserDevice.Ctx(ctx).Data(dao.UserDevice.Columns().MobileKey, mobileKey).
		Where(dao.UserDevice.Columns().DeviceId, userDevice.DeviceId).Update(); err != nil {
		panic(err)
	}
}

func (s sCar) findUserDevice(ctx context.Context, deviceId *int64) (userDevice *entity.UserDevice) {
	sql := dao.UserDevice.Ctx(ctx).Where(
		dao.UserDevice.Columns().UserId, service.BizCtx().Get(ctx).Data.Get(consts.ContextKeyUserId),
	)
	if deviceId == nil {
		sql = sql.Where(dao.UserDevice.Columns().IsSelect, consts.CarSelectYes)
	} else {
		sql = sql.Where(dao.UserDevice.Columns().DeviceId, deviceId)
	}
	result, err := sql.One()
	if err != nil {
		panic(err)
	}
	if !result.IsEmpty() {
		if err = result.Struct(&userDevice); err != nil {
			panic(err)
		}
	}
	return
}

func (s sCar) EnabledSpeedLimit(ctx context.Context) {
	userDevice := s.findUserDevice(ctx, nil)
	if userDevice == nil || consts.UserDeviceChild == userDevice.UserDeviceType {
		panic(gerror.NewCode(enums.IllegalOperation))
	}
	mp := make(map[string]any)
	mp = map[string]any{consts.LimitStr: 10}
	var speedLimit = true
	if userDevice.SpeedLimit {
		mp = map[string]any{consts.LimitStr: 60}
		speedLimit = consts.CarSpeedLimitNo
	}
	in := &model.DevicePropertyInput{
		DeviceKey: s.findDeviceInfo(ctx, &userDevice.DeviceId).DeviceName,
		Params:    mp,
	}
	if _, err := service.DevDeviceProperty().Set(ctx, in); err != nil {
		panic(err)
	}
	if _, err := dao.UserDevice.Ctx(ctx).Data(dao.UserDevice.Columns().SpeedLimit, speedLimit).
		Where(dao.UserDevice.Columns().DeviceId, userDevice.DeviceId).Update(); err != nil {
		panic(err)
	}
}

func (s *sCar) GetSimDataTraffic(ctx context.Context) (out *model.SimDataTrafficOutput) {
	device := s.findDeviceInfo(ctx, nil)
	if device == nil {
		return
	}
	treeMap := gmap.NewTreeMap(gutil.ComparatorString)
	treeMap.Set("x-sign-uri", "/cube/v4/sims/"+device.SimId)
	num := gconv.String(grand.Intn(100))
	treeMap.Set("nonce", num)
	timestamp := gtime.Now().TimestampMilliStr()
	treeMap.Set("timestamp", timestamp)
	jsonString, err := gjson.EncodeString(treeMap)
	if err != nil {
		panic(err)
	}
	sign, err := generateSignature(g.Cfg().MustGet(ctx, "linkSim.privateKey").String(), []byte(jsonString))
	if err != nil {
		panic(err)
	}
	resp, err := httpClient(ctx, g.Cfg().MustGet(ctx, "linkSim.accessKey").String(), sign, device.SimId, timestamp, num)
	if err != nil {
		panic(err)
	}
	jsonMap := make(map[string]any, 1)
	if err = gjson.DecodeTo(resp, &jsonMap); err != nil {
		panic(err)
	}
	data := gconv.Map(jsonMap["data"])
	if gutil.IsEmpty(data) {
		return
	}
	simService := gconv.Map(data["sim_service"])
	if gutil.IsEmpty(simService) {
		return
	}
	bundles := gconv.Maps(simService["bundles"])
	if gutil.IsEmpty(bundles) {
		return
	}
	out = &model.SimDataTrafficOutput{}
	dataLimit := gconv.Float32(bundles[0]["data_limit"])
	out.TotalDataTraffic = fmt.Sprintf("%.1f", dataLimit/(1024*1024)) + "MB"
	currentCycleUsage := gconv.Float32(bundles[0]["current_cycle_usage"])
	out.ConsumeDataTraffic = fmt.Sprintf("%.1f", currentCycleUsage/(1024*1024)) + "MB"
	out.SurplusDataTraffic = fmt.Sprintf("%.1f", (dataLimit-currentCycleUsage)/(1024*1024)) + "MB"
	return
}

func httpClient(ctx context.Context, key, sign, simID, timestamp, num string) (string, error) {
	client := g.Client()
	client.SetHeader("Span-Id", "0.0.1")
	client.SetHeader("Trace-Id", "NBC56410N97LJ016FQA")
	client.SetHeader("Accept-Language", "zh-CN")
	client.SetHeader("Authorization", "LF "+key+"/"+sign)
	client.SetHeader("X-LF-Signature-Type", "2.0")
	client.SetHeader("timestamp", timestamp)
	client.SetHeader("nonce", num)
	client.SetHeader("Content-Type", "application/json")
	r, err := client.Get(ctx, "https://api.linksfield.net/cube/v4/sims/"+simID)
	if err != nil {
		return "", err
	}
	return r.ReadAllString(), nil
}

func generateSignature(privateKeyStr string, message []byte) (string, error) {
	privateKey, err := parsePrivateKey(privateKeyStr)
	if err != nil {
		return "", err
	}
	hasher := sha1.New()
	hasher.Write(message)
	hash := hasher.Sum(nil)
	sign, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA1, hash)
	if err != nil {
		return "", err
	}
	return gbase64.EncodeToString(sign), nil
}

func parsePrivateKey(privateKey string) (*rsa.PrivateKey, error) {
	decodeString, err := gbase64.DecodeString(privateKey)
	if err != nil {
		return nil, err
	}
	privKey, err := x509.ParsePKCS8PrivateKey(decodeString)
	if err != nil {
		return nil, err
	}
	rsaPrivKey, ok := privKey.(*rsa.PrivateKey)
	if !ok {
		return nil, err
	}
	return rsaPrivKey, nil
}
