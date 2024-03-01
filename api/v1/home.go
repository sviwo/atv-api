package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type HomeDataReq struct {
	g.Meta `path:"home/get/data" method:"get" tags:"首页相关" summary:"获取首页数据"`
}

type HomeDataRes struct {
	TravelKm           int           `json:"travelKm"           dc:"行驶公里数"`
	ResidueElectricity int           `json:"residueElectricity" dc:"剩余电量"`
	Version            []*VersionRes `json:"version"            dc:"新版本信息"`
}
