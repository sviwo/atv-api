package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"sviwo/api/v1"
	"sviwo/internal/service"
)

var Version = cVersion{}

type cVersion struct{}

/*
GetNewVersion 获取最新版本信息
*/
func (c cVersion) GetNewVersion(ctx context.Context, req *v1.VersionReq) (res []*v1.VersionRes, err error) {
	if err = gconv.Structs(service.Version().GetNewVersion(ctx), &res); err != nil {
		panic(err)
	}
	return
}
