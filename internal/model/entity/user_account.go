// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-01-16 21:22:22
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserAccount is the golang structure for table user_account.
type UserAccount struct {
	AccountId  uint64      `json:"accountId"  description:""`
	UserId     int64       `json:"userId"     description:""`
	Balance    float64     `json:"balance"    description:"账户余额"`
	MyIntegral uint        `json:"myIntegral" description:"我的积分"`
	CreateTime *gtime.Time `json:"createTime" description:""`
	UpdateTime *gtime.Time `json:"updateTime" description:""`
	IsDelete   bool        `json:"isDelete"   description:"是否删除：true=已删除，false=正常"`
}
