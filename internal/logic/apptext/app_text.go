package apptext

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"sviwo/internal/consts"
	"sviwo/internal/dao"
	"sviwo/internal/model"
	"sviwo/internal/service"
	"sviwo/utility"
)

func init() {
	service.RegisterAppText(New())
}

func New() *sAppText {
	return &sAppText{}
}

type sAppText struct{}

func (s sAppText) GetAppTextTree(ctx context.Context) (out []*model.AppTextTreeOutput) {
	var array []*model.AppTextTreeOutput
	if err := dao.AppText.Ctx(ctx).
		Where("enable", consts.EnableDisplay).
		Where("is_delete", consts.DeleteOn).
		OrderAsc("orders").
		Scan(&array); err != nil {
		panic(err)
	}
	if err := gconv.Structs(utility.BuildTree(gconv.Maps(array)), &out); err != nil {
		panic(err)
	}
	return
}

func (s sAppText) GetAppTextDetail(ctx context.Context, textId int64) (out *model.AppTextDetailOutput) {
	if err := dao.AppText.Ctx(ctx).
		Where("text_id", textId).
		Where("enable", consts.EnableDisplay).
		Where("is_delete", consts.DeleteOn).
		Scan(&out); err != nil {
		panic(err)
	}
	return
}
