package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type TravelRecordQueryReq struct {
	g.Meta `path:"/travelRecord/list/get" method:"get" tags:"行程相关" summary:"获取行程列表"`
	UserId int64 `json:"userId"         description:"用户ID"     v:"required"`
	CarId  int64 `json:"carId"         description:"车辆ID"     `
	CommonPaginationReq
}

type TravelRecordRes struct {
	Result interface{} `dc:"列表数据"`
	CommonPaginationRes
}
