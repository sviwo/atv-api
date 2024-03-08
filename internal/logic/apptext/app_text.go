package version

import (
	"context"
	"github.com/gogf/gf/v2/container/glist"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gutil"
	"sviwo/internal/consts"
	"sviwo/internal/dao"
	"sviwo/internal/model"
	"sviwo/internal/service"
)

func init() {
	service.RegisterAppText(New())
}

func New() *sAppText {
	return &sAppText{}
}

type sAppText struct{}

func (s sAppText) GetAppTextTree(ctx context.Context) (out []*model.AppTextListOutput) {
	if err := dao.AppText.Ctx(ctx).
		Where("enable", consts.EnableDisplay).
		Where("is_delete", consts.DeleteOn).
		OrderAsc("orders").
		Scan(&out); err != nil {
		panic(err)
	}
	if err := gconv.Structs(gutil.Values(buildTree(out)), &out); err != nil {
		panic(err)
	}
	return
}

func buildTree(array []*model.AppTextListOutput) (outputs []*model.AppTextListOutput) {
	temp := map[interface{}]*model.AppTextListOutput{}
	list := glist.New()
	for _, data := range array {
		temp[data.TextId] = data
		list.PushBack(data.TextId)
	}
	for range temp {
		output := temp[list.PopFront()]
		if temp[output.ParentId] == nil {
			outputs = append(outputs, output)
		} else {
			temp[output.ParentId].Children = append(temp[output.ParentId].Children, output)
		}
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
