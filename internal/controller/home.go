package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"sviwo/api/v1"
	"sviwo/internal/service"
)

var Home = cHome{}

type cHome struct{}

/*
GetHomeData 获取首页数据
*/
func (c cHome) GetHomeData(ctx context.Context, req *v1.HomeDataReq) (res v1.HomeDataRes, err error) {
	if err = gconv.Structs(service.Home().GetHomeData(ctx), &res); err != nil {
		panic(err)
	}
	return
}
