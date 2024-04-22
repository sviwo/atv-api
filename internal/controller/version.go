package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gutil"
	v1 "sviwo/api/v1"
	"sviwo/internal/model"
	"sviwo/internal/service"
)

var Version = cVersion{}

type cVersion struct{}

func (c cVersion) GetVersionList(ctx context.Context, req *v1.VersionListReq) (res *v1.VersionListRes, err error) {
	in := new(model.VersionInput)
	if err = gconv.Scan(req, &in); err != nil {
		panic(err)
	}
	count, out := service.Version().GetVersionList(ctx, in)
	if gutil.IsEmpty(out) {
		return
	}
	res = &v1.VersionListRes{}
	if err = gconv.Scan(out, &res.Result); err != nil {
		panic(err)
	}
	res.Total = count
	res.CurrentPage = req.PageNum
	return
}

func (c cVersion) GetNewVersion(ctx context.Context, req *v1.NewVersionReq) (res []*v1.VersionRes, err error) {
	if err = gconv.Scan(service.Version().GetNewVersion(ctx), &res); err != nil {
		panic(err)
	}
	return
}
