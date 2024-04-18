package version

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gutil"
	"sviwo/internal/consts"
	"sviwo/internal/consts/enums"
	"sviwo/internal/dao"
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
	imgPrefix := g.Cfg().MustGet(ctx, "aliyun.oss.fileUrlPrefix").String()
	out.CertificateFrontImg = imgPrefix + out.CertificateFrontImg
	out.CertificateBackImg = imgPrefix + out.CertificateBackImg
	return
}

func findUserAuthByUserId(ctx context.Context) (userAuth *entity.UserAuth) {
	if err := dao.UserAuth.Ctx(ctx).Where(
		dao.UserAuth.Columns().UserId, service.BizCtx().Get(ctx).Data.Get(consts.ContextKeyUserId),
	).Scan(&userAuth); err != nil {
		panic(err)
	}
	return
}

func (s sUserAuth) SubmitUserAuth(ctx context.Context, in model.UserAuthInput) {
	userAuth := findUserAuthByUserId(ctx)
	if userAuth == nil {
		panic(gerror.NewCode(enums.UserNotExists))
	}
	if consts.UserAuthStatusIn == userAuth.AuthStatus {
		panic(gerror.NewCode(enums.IllegalOperation))
	}
	if err := g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		if err := gconv.Scan(in, &userAuth); err != nil {
			return err
		}
		userAuth.AuthTime = gtime.Now()
		userAuth.AuthStatus = consts.UserAuthStatusIn
		if result, err := dao.UserAuth.Ctx(ctx).Update(
			&userAuth, dao.User.Columns().UserId, service.BizCtx().Get(ctx).Data.Get(consts.ContextKeyUserId),
		); err != nil {
			return err
		} else {
			if affected, _ := result.RowsAffected(); affected > 0 {
				user := new(entity.User)
				if err = dao.User.Ctx(ctx).Where(dao.User.Columns().UserId, userAuth.UserId).Scan(&user); err != nil {
					return err
				}
				if gutil.IsEmpty(user.FirstName) && gutil.IsEmpty(user.LastName) {
					user.FirstName = in.AuthFirstName
					user.LastName = in.AuthLastName
					if _, err = dao.User.Ctx(ctx).Update(&user, dao.User.Columns().UserId, user.UserId); err != nil {
						return err
					}
				}
			}
		}
		return nil
	}); err != nil {
		panic(err)
	}
}
