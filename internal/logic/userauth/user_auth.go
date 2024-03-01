package version

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"sviwo/internal/consts"
	"sviwo/internal/dao"
	rcode "sviwo/internal/logic/biz/enums"
	"sviwo/internal/model"
	"sviwo/internal/model/entity"
	"sviwo/internal/service"
)

func init() {
	service.RegisterUserAuth(New())
}

func New() *sUserAuth {
	return &sUserAuth{}
}

type sUserAuth struct{}

func (s sUserAuth) GetUserAuthInfo(ctx context.Context) (out *model.UserAuthOutput) {
	if err := gconv.Struct(findUserAuthByUserId(ctx), &out); err != nil {
		panic(err)
	}
	return
}

func findUserAuthByUserId(ctx context.Context) (userAuth *entity.UserAuth) {
	if err := dao.UserAuth.Ctx(ctx).Where(
		"user_id", service.BizCtx().Get(ctx).Data.Get(consts.ContextKeyUserId),
	).Scan(&userAuth); err != nil {
		panic(err)
	}
	return
}

func (s sUserAuth) SubmitUserAuth(ctx context.Context, in model.UserAuthInput) {
	userAuth := findUserAuthByUserId(ctx)
	if userAuth == nil {
		panic(gerror.NewCode(rcode.UserNotExists))
	}
	if consts.UserAuthStatusIn == userAuth.AuthStatus {
		panic(gerror.NewCode(rcode.IllegalOperation))
	}
	if _, err := dao.UserAuth.Ctx(ctx).Where(
		"user_id", service.BizCtx().Get(ctx).Data.Get(consts.ContextKeyUserId),
	).Where("auth_status", consts.UserAuthStatusIn).
		Where("auth_time", gtime.Now()).Update(in); err != nil {
		panic(err)
	}
}
