package version

import (
	"context"
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

func (s sAppText) GetAppTextList(ctx context.Context) (out []*model.AppTextListOutput) {
	if err := dao.AppText.Ctx(ctx).Where("is_delete", consts.DeleteOn).Scan(&out); err != nil {
		panic(err)
	}
	if err := gconv.Structs(gutil.Values(buildTree(out)), &out); err != nil {
		panic(err)
	}
	return
}

func buildTree(allMenus []*model.AppTextListOutput) map[int64]*model.AppTextListOutput {
	temp := map[int64]*model.AppTextListOutput{}
	result := map[int64]*model.AppTextListOutput{}
	for _, menu := range allMenus {
		temp[menu.TextId] = menu
	}
	for _, node := range temp {
		if temp[node.ParentId] == nil {
			result[node.TextId] = node
		} else {
			temp[node.ParentId].Children = append(temp[node.ParentId].Children, node)
		}
	}
	return result
}

func (s sAppText) GetAppTextDetail(ctx context.Context, textId int64) (out *model.AppTextDetailOutput) {
	if err := dao.AppText.Ctx(ctx).
		Where("text_id", textId).
		Where("is_delete", consts.DeleteOn).
		Scan(&out); err != nil {
		panic(err)
	}
	return
}
