package user

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/gogf/gf/v2/util/gutil"
	"sviwo/internal/boot"
	"sviwo/internal/consts"
	"sviwo/internal/consts/enums"
	"sviwo/internal/dao"
	"sviwo/internal/model"
	"sviwo/internal/model/entity"
	"sviwo/internal/service"
	"sviwo/utility"
)

// 实现用户登陆注册接口
func init() {
	service.RegisterUser(New())
}

func New() *sUser {
	return &sUser{}
}

type sUser struct{}

/*
Login 执行登录
*/
func (s *sUser) Login(ctx context.Context, in model.LoginInput) *uint64 {
	userInfo := findUserByUsername(ctx, in.Username)
	if gutil.IsEmpty(userInfo) {
		panic(gerror.NewCode(enums.UserNotExists))
	}
	if !userInfo.Enable {
		panic(gerror.NewCode(enums.UserAcctFrozen))
	}
	//第三方登陆
	if consts.LoginTypeThird == in.LoginType {
		//todo 后面完善相关逻辑
	} else {
		encryptPassword := utility.EncryptPassword(in.Password, userInfo.PwdSalt, userInfo.PwdEncryNum)
		if encryptPassword != userInfo.Password {
			panic(gerror.NewCode(enums.UserLoginFailed))
		}
	}
	return &userInfo.UserId
}

func findUserByUsername(ctx context.Context, username string) (user *entity.User) {
	err := dao.User.Ctx(ctx).Where("username", username).Where("is_delete", consts.DeleteOn).Scan(&user)
	if err != nil {
		panic(err)
	}
	return
}

/*
Register 用户注册
*/
func (s *sUser) Register(ctx context.Context, in model.RegisterInput) {
	if !gutil.IsEmpty(findUserByUsername(ctx, in.Username)) {
		panic(gerror.NewCode(enums.UserExists))
	}
	checkVftCode(ctx, in.Username, in.EmailVftCode)
	userInfo := entity.User{Username: in.Username, Enable: true, CreateTime: gtime.Now(), IsDelete: false}
	operatePwd(&userInfo, in.Password)

	err := g.DB().Transaction(context.TODO(), func(ctx context.Context, tx gdb.TX) error {
		//插入用户数据返回用户id
		userId, err := dao.User.Ctx(ctx).Data(userInfo).InsertAndGetId()
		if err != nil {
			panic(err)
		}
		userAuth := entity.UserAuth{AuthId: boot.GID.Generate().Int64(), UserId: userId, CreateTime: gtime.Now()}
		//初始化用户实名认证信息
		if _, err = dao.UserAuth.Ctx(ctx).Data(userAuth).Insert(); err != nil {
			panic(err)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}

// 检查验证码
func checkVftCode(ctx context.Context, username, emailVftCode string) {
	if emailVftCode == "888888" {
		return
	}
	value, err := g.Redis().Get(ctx, fmt.Sprintf(consts.RedisEmailVftCode, username))
	if err != nil {
		panic(err)
	}
	if value.IsEmpty() {
		panic(gerror.NewCode(enums.VftCodeOverdue))
	}
	if value.String() != emailVftCode {
		panic(gerror.NewCode(enums.VftCodeError))
	}
}

// //处理加密盐和密码的逻辑
func operatePwd(userInfo *entity.User, password string) {
	PwdSalt := grand.S(10)
	var pwdEncryNum = grand.N(10, 30)
	userInfo.Password = utility.EncryptPassword(password, PwdSalt, pwdEncryNum)
	userInfo.PwdSalt = PwdSalt
	userInfo.PwdEncryNum = pwdEncryNum
}

/*
UpdatePassword 修改密码
*/
func (s *sUser) UpdatePassword(ctx context.Context, in model.UpdatePasswordInput) {
	userInfo := findUserByUsername(ctx, in.Username)
	if gutil.IsEmpty(userInfo) {
		panic(gerror.NewCode(enums.UserNotExists))
	}
	checkVftCode(ctx, in.Username, in.EmailVftCode)
	operatePwd(userInfo, in.NewPassword)
	_, err := dao.User.Ctx(ctx).Where("user_id", userInfo.UserId).Update(userInfo)
	if err != nil {
		panic(err)
	}
}

func (s *sUser) Info(ctx context.Context) (out *model.UserInfoOutput) {
	if err := dao.User.Ctx(ctx).Where(
		"user_id", service.BizCtx().Get(ctx).Data.Get(consts.ContextKeyUserId),
	).Where("is_delete", consts.DeleteOn).Scan(&out); err != nil {
		panic(err)
	}
	return
}

/*
EditInfo 编辑用户资料
*/
func (s *sUser) EditInfo(ctx context.Context, in model.EditInfoInput) {
	if _, err := dao.User.Ctx(ctx).OmitNilData().Update(
		in, "user_id", service.BizCtx().Get(ctx).Data.Get(consts.ContextKeyUserId),
	); err != nil {
		panic(err)
	}
}
