// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-08-18 15:21:45
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserThirdAccount is the golang structure for table user_third_account.
type UserThirdAccount struct {
	Id           uint64      `json:"id"           orm:"id"            description:""`
	UserId       int64       `json:"userId"       orm:"user_id"       description:"本地用户id"`
	ThirdUserId  string      `json:"thirdUserId"  orm:"third_user_id" description:"第三方用户id"`
	ProviderType int         `json:"providerType" orm:"provider_type" description:"第三方平台类型：0=apple，1=facebook，2=google"`
	CreateTime   *gtime.Time `json:"createTime"   orm:"create_time"   description:""`
	UpdateTime   *gtime.Time `json:"updateTime"   orm:"update_time"   description:""`
}
