package appmedia

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"sviwo/internal/consts"
	"sviwo/internal/dao"
	"sviwo/internal/model"
	"sviwo/internal/service"
	"sviwo/pkg/utility"
)

func init() {
	service.RegisterAppMedia(New())
}

func New() *sAppMedia {
	return &sAppMedia{}
}

type sAppMedia struct{}

func (s sAppMedia) GetAppMediaTree(ctx context.Context, pageType int) (out []*model.AppMediaTreeOutput) {
	var array []*model.AppMediaTreeOutput
	if err := dao.AppMedia.Ctx(ctx).Union(
		dao.AppMedia.Ctx(ctx).
			Fields(
				dao.AppMedia.Columns().Id,
				dao.AppMedia.Columns().ParentId,
				dao.AppMedia.Columns().DisplayType,
				dao.AppMedia.Columns().Title,
				dao.AppMedia.Columns().MediaDesc,
				dao.AppMedia.Columns().SmallImg,
				dao.AppMedia.Columns().Orders,
				" '' AS content ",
			).
			Where(dao.AppMedia.Columns().DisplayType, consts.DisplayTypeDirect).
			Where(dao.AppMedia.Columns().PageType, pageType).
			Where(dao.AppMedia.Columns().Enable, consts.EnableDisplay).
			Where(dao.AppMedia.Columns().IsDelete, consts.DeleteOn),
		dao.AppMedia.Ctx(ctx).
			Fields(
				dao.AppMedia.Columns().Id,
				dao.AppMedia.Columns().ParentId,
				dao.AppMedia.Columns().DisplayType,
				dao.AppMedia.Columns().Title,
				dao.AppMedia.Columns().MediaDesc,
				dao.AppMedia.Columns().SmallImg,
				dao.AppMedia.Columns().Orders,
				dao.AppMedia.Columns().Content,
			).
			Where(dao.AppMedia.Columns().DisplayType, consts.DisplayTypeRedirect).
			Where(dao.AppMedia.Columns().PageType, pageType).
			Where(dao.AppMedia.Columns().Enable, consts.EnableDisplay).
			Where(dao.AppMedia.Columns().IsDelete, consts.DeleteOn),
	).OrderAsc(dao.AppMedia.Columns().Orders).
		Scan(&array); err != nil {
		panic(err)
	}

	if err := gconv.Scan(utility.BuildTree(gconv.Maps(array)), &out); err != nil {
		panic(err)
	}
	return
}

func (s sAppMedia) GetAppMediaDetail(ctx context.Context, Id int64) (content string) {
	result, err := dao.AppMedia.Ctx(ctx).
		Fields(dao.AppMedia.Columns().Content).
		Where(dao.AppMedia.Columns().Id, Id).
		Where(dao.AppMedia.Columns().DisplayType, consts.DisplayTypeDirect).
		WhereNot(dao.AppMedia.Columns().PageType, consts.PageTypeVideoTutorial).
		Where(dao.AppMedia.Columns().Enable, consts.EnableDisplay).
		Where(dao.AppMedia.Columns().IsDelete, consts.DeleteOn).
		One()
	if err != nil {
		panic(err)
	}
	if !result.IsEmpty() {
		return result.GMap().String()
	}
	return
}
