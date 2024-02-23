package common

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/grand"
	"sviwo/internal/consts"
	"sviwo/internal/logic/biz/enums"
	"sviwo/internal/service"
)

func init() {
	service.RegisterCommon(New())
}

func New() *sCommon {
	return &sCommon{}
}

type sCommon struct{}

/*
GetVftCode 获取邮箱验证码
*/
func (s *sCommon) GetVftCode(ctx context.Context, email string) error {
	value, err := g.Redis().Get(ctx, fmt.Sprintf(consts.MethodUserServiceRedisKey, email))
	if err != nil {
		panic(err)
	}
	if !value.IsEmpty() {
		return gerror.NewCode(rcode.IllegalOperation)
	}
	VftCode := grand.Digits(6)
	err = g.Redis().SetEX(ctx, fmt.Sprintf(consts.MethodUserServiceRedisKey, email), VftCode, 60)
	//TODO 生成验证码存入redis并发送到用户邮箱
	glog.Info(ctx, "验证码====", VftCode)
	if err != nil {
		panic(err)
	}
	return nil
}
