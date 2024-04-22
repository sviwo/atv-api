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

func (s sVersion) GetVersionList(ctx context.Context, in *model.VersionInput) (totalCount int, out []*model.VersionOutput) {
	result, totalCount, err := dao.Version.Ctx(ctx).
		Where(dao.Version.Columns().VersionStatus, consts.VersionReleaseStatusYes).
		Where(dao.Version.Columns().IsDelete, consts.DeleteOn).
		Where(dao.Version.Columns().VersionType, consts.VersionTypeApp).
		Page(in.PageNum, in.PageSize).
		Order(dao.Version.Columns().CreateTime, in.OrderBy).
		AllAndCount(true)
	if err != nil {
		panic(err)
	}
	if result.IsEmpty() {
		return
	}
	if err = result.Structs(&out); err != nil {
		panic(err)
	}
	return
}

func (s sVersion) GetNewVersion(ctx context.Context) (out []*model.VersionOutput) {
	if err := dao.Version.Ctx(ctx).
		WhereIn(
			dao.Version.Columns().VersionCode,
			dao.Version.Ctx(ctx).
				Fields("MAX(version_code)").
				Where(dao.Version.Columns().VersionStatus, consts.VersionReleaseStatusYes).
				Where(dao.Version.Columns().IsDelete, consts.DeleteOn).
				Group(dao.Version.Columns().VersionType),
		).Limit(2).
		Scan(&out); err != nil {
		panic(err)
	}
	return
}
