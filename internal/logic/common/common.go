package common

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
	"sviwo/internal/consts"
	"sviwo/internal/service"
	"sviwo/pkg/utility"
	ecc2 "sviwo/pkg/utility/encrypt"
)

func init() {
	service.RegisterCommon(New())
}

func New() *sCommon {
	return &sCommon{}
}

type sCommon struct{}

func (s *sCommon) GetVftCode(ctx context.Context, email string) {
	utility.MethodReqLimit(ctx, "GetVftCode", email, 60)
	VftCode := grand.Digits(6)
	if err := g.Redis().SetEX(ctx, fmt.Sprintf(consts.RedisEmailVftCode, email), VftCode, 180); err != nil {
		panic(err)
	}
	vftCodeErr := utility.SendEmail("验证码", VftCode, email)
	//验证码发送失败则取消接口频繁访问限制
	if vftCodeErr != nil {
		if _, err := g.Redis().Del(ctx, fmt.Sprintf(consts.RedisMethodReqLimit, email)); err != nil {
			panic(err)
		}
		panic(vftCodeErr)
	}
}

func (s *sCommon) SendEmail(ctx context.Context, ip, email, content string) {
	utility.MethodReqLimit(ctx, "SendEmail", ip, 30)
	content = fmt.Sprintf("Contact Email: %s   Content: %s", email, content)
	receiveEmail := g.Cfg().MustGet(context.TODO(), "email.username").String()
	vftCodeErr := utility.SendEmail("Inquiry Email", content, receiveEmail)
	//验证码发送失败则取消接口频繁访问限制
	if vftCodeErr != nil {
		if _, err := g.Redis().Del(ctx, fmt.Sprintf(consts.RedisMethodReqLimit+"SendEmail", ip)); err != nil {
			panic(err)
		}
		panic(vftCodeErr)
	}
}

func (s *sCommon) GetEccPublicKey(ctx context.Context) (publicKey, publicCode string) {
	key, err := ecc2.GenerateEccKeyHex()
	if err != nil {
		panic(err)
	}
	publicCode = utility.GID.Generate().String()
	if err = g.Redis().SetEX(
		ctx, fmt.Sprintf(consts.RedisEccPrivateKey, publicCode), key.PrivateKey, 120,
	); err != nil {
		panic(err)
	}
	publicKey = key.PublicKey
	return
}
