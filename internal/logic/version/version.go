package version

import (
	"context"
	"sviwo/internal/consts"
	"sviwo/internal/dao"
	"sviwo/internal/model"
	"sviwo/internal/service"
)

func init() {
	service.RegisterVersion(New())
}

func New() *sVersion {
	return &sVersion{}
}

type sVersion struct{}

func (s sVersion) GetNewVersion(ctx context.Context) (out []*model.VersionOutput) {
	err := dao.Version.Ctx(ctx).
		WhereIn(
			"version_number",
			dao.Version.Ctx(ctx).
				Fields("MAX(version_number)").
				Where("version_status", 1).
				Where("is_delete", consts.DeleteOn).
				Group("version_type"),
		).Limit(2).Scan(&out)
	if err != nil {
		panic(err)
	}
	return
}
