package version

import (
	"context"
	"github.com/gogf/gf/v2/container/glist"
	"github.com/gogf/gf/v2/util/gconv"
	"sviwo/internal/consts"
	"sviwo/internal/dao"
	"sviwo/internal/model"
	"sviwo/internal/service"
)

func init() {
	service.RegisterAppVideos(New())
}

func New() *sAppVideos {
	return &sAppVideos{}
}

type sAppVideos struct{}

func (s sAppVideos) GetAppVideosTree(ctx context.Context) (out []*model.AppVideosListOutput) {
	var array []*model.AppVideosListOutput
	if err := dao.AppVideos.Ctx(ctx).
		Where("enable", consts.EnableDisplay).
		Where("is_delete", consts.DeleteOn).
		OrderAsc("orders").Scan(&array); err != nil {
		panic(err)
	}
	if err := gconv.Structs(buildTree(array), &out); err != nil {
		panic(err)
	}
	return
}

func buildTree(array []*model.AppVideosListOutput) (outputs []*model.AppVideosListOutput) {
	temp := map[interface{}]*model.AppVideosListOutput{}
	list := glist.New()
	for _, data := range array {
		temp[data.VideosId] = data
		list.PushBack(data.VideosId)
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
