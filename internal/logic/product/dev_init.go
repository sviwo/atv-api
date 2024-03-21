package product

import (
	"context"
	"encoding/json"
	"sviwo/internal/consts"
	"sviwo/internal/dao"
	"sviwo/internal/model"
	"sviwo/internal/model/entity"
	"sviwo/internal/service"
	"sviwo/pkg/tsd/comm"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type sDevInit struct{}

func init() {
	service.RegisterDevInit(DevInitNew())
}

func DevInitNew() *sDevInit {
	return &sDevInit{}
}

// InitProductForTd 产品表结构初始化
func (s *sDevInit) InitProductForTd(ctx context.Context) (err error) {
	// 资源锁
	lockKey := "tdLock:initProductTable"
	lockVal, err := g.Redis().Do(ctx, "SET", lockKey, gtime.Now().Unix(), "NX", "EX", "3600")
	if err != nil {
		return
	}
	if lockVal.String() != "OK" {
		return
	}
	defer func() {
		_, err = g.Redis().Do(ctx, "DEL", lockKey)
	}()

	var list []*entity.Product
	c := dao.Product.Columns()
	err = dao.Product.Ctx(ctx).Where(c.Status, model.ProductStatusOn).Where(c.IsDelete, consts.DeleteOn).Scan(&list)
	if err != nil || len(list) == 0 {
		return
	}

	// 检测td表结构是否存在，不存在则创建
	for _, p := range list {
		stable := comm.ProductTableName(p.ProductKey)
		b, _ := service.TSLTable().CheckStable(ctx, stable)

		if b {
			continue
		}

		var tsl *model.TSL
		err = json.Unmarshal([]byte(p.Metadata), &tsl)
		if err != nil {
			g.Log().Error(ctx, err)
			continue
		}
		if len(tsl.Properties) == 0 {
			g.Log().Errorf(ctx, "产品 %s 物模型数据异常", p.ProductKey)
			continue
		}
		tsl.ProductName = p.ProductName
		tsl.ProductKey = p.ProductKey
		err = service.TSLTable().CreateStable(ctx, tsl)
		if err != nil {
			g.Log().Error(ctx, err)
			continue
		}
	}

	return nil
}

// InitDeviceForTd 设备表结构初始化
func (s *sDevInit) InitDeviceForTd(ctx context.Context) (err error) {
	// 资源锁
	lockKey := "tdLock:initDeviceTable"
	lockVal, err := g.Redis().Do(ctx, "SET", lockKey, gtime.Now().Unix(), "NX", "EX", "3600")
	if err != nil {
		return
	}
	if lockVal.String() != "OK" {
		return
	}
	defer func() {
		_, err = g.Redis().Do(ctx, "DEL", lockKey)
	}()

	var list []*entity.Device
	c := dao.Device.Columns()
	err = dao.Device.Ctx(ctx).WhereGT(c.Status, model.DeviceStatusNoEnable).Where(c.IsDelete, consts.DeleteOn).Scan(&list)
	if err != nil || len(list) == 0 {
		return
	}

	// 检测td表结构是否存在，不存在则创建
	for _, d := range list {

		// 检测设备表是否创建TD表的标识
		/*if d.MetadataTable == 1 {
			continue
		}*/

		pd, err := service.DevProduct().Detail(ctx, d.ProductKey)
		if err != nil {
			g.Log().Error(ctx, err)
			continue
		}
		if pd == nil {
			g.Log().Errorf(ctx, "设备 %s 所属产品不存在", d.DeviceName)
			continue
		}

		table := comm.DeviceTableName(d.DeviceName)
		b, _ := service.TSLTable().CheckTable(ctx, table)

		if b {
			continue
		}

		d := d
		go func() {
			err = service.TSLTable().CreateTable(ctx, pd.ProductKey, d.DeviceName)
			if err != nil {
				g.Log().Errorf(ctx, "设备 %s(%s) 建表失败: %s", d.DeviceName, pd.ProductKey, err.Error())
			}
		}()
	}

	return nil
}
