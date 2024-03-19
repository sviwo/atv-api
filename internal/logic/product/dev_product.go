package product

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/database/gdb"
	"sviwo/internal/consts"
	"sviwo/internal/dao"
	"sviwo/internal/model"
	"sviwo/internal/service"
	"time"
)

type sDevProduct struct{}

func init() {
	service.RegisterDevProduct(productNew())
}

func productNew() *sDevProduct {
	return &sDevProduct{}
}

func (s *sDevProduct) Detail(ctx context.Context, key string) (out *model.DetailProductOutput, err error) {
	err = dao.Product.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: time.Second * 30,
		Name:     consts.GetDetailProductOutput + key,
		Force:    false,
	}).WithAll().Where(dao.Product.Columns().ProductKey, key).Scan(&out)
	if err != nil || out == nil {
		return
	}

	if out.Metadata != "" {
		err = json.Unmarshal([]byte(out.Metadata), &out.TSL)
		if err != nil {
			return
		}
	}

	if out.Category != nil {
		out.CategoryName = out.Category.Name
	}

	return
}

func (s *sDevProduct) List(ctx context.Context) (list []*model.ProductOutput, err error) {
	m := dao.Product.Ctx(ctx)
	err = m.WithAll().
		Where(dao.Product.Columns().Status, model.ProductStatusOn).
		OrderDesc(dao.Product.Columns().ProductId).
		Scan(&list)
	if err != nil || len(list) == 0 {
		return
	}

	dLen := len(list)
	var productIds = make([]int64, dLen)
	for i, v := range list {
		productIds[i] = v.ProductId
		list[i].Metadata = ""
	}

	return
}
