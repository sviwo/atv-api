package common

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
	"sviwo/internal/consts"
	"sviwo/internal/consts/enums"
	"sviwo/internal/service"
	"sviwo/pkg/utility"
)

func init() {
	service.RegisterCommon(New())
}

func New() *sCommon {
	return &sCommon{}
}

type sCommon struct{}

/*
GetVftCode
获取邮箱验证码
*/
func (s *sCommon) GetVftCode(ctx context.Context, email string) error {
	restrictFrequentRequest(ctx, email)
	VftCode := grand.Digits(6)
	err := g.Redis().SetEX(ctx, fmt.Sprintf(consts.RedisEmailVftCode, email), VftCode, 180)
	if err != nil {
		panic(err)
	}
	vftCodeErr := utility.SendEmail("验证码", VftCode, email)
	//验证码发送失败则取消接口频繁访问限制
	if vftCodeErr != nil {
		_, err = g.Redis().Del(ctx, fmt.Sprintf(consts.RedisMethodCommonServiceGetvftcode, email))
		if err != nil {
			panic(err)
		}
	}
	return vftCodeErr
}

/*
restrictFrequentRequest
接口频繁访问限制
*/
func restrictFrequentRequest(ctx context.Context, email string) {
	value, err := g.Redis().Get(ctx, fmt.Sprintf(consts.RedisMethodCommonServiceGetvftcode, email))
	if err != nil {
		panic(err)
	}
	if !value.IsEmpty() {
		panic(gerror.NewCode(enums.IllegalOperation))
	}
	err = g.Redis().SetEX(ctx, fmt.Sprintf(consts.RedisMethodCommonServiceGetvftcode, email), email, 60)
	if err != nil {
		panic(err)
	}
}
