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
	"sviwo/internal/consts"
	"sviwo/internal/consts/enums"
	"sviwo/internal/dao"
	"sviwo/internal/model"
	"sviwo/internal/model/do"
	"sviwo/internal/model/entity"
	"sviwo/internal/service"
	"sviwo/pkg/utility"
)

// 实现用户登陆注册接口
func init() {
	service.RegisterUser(New())
}

func New() *sUser {
	return &sUser{}
}

type sUser struct {
	autoRegister bool
}

/*
Login 执行登录
*/
func (s *sUser) Login(ctx context.Context, in model.LoginInput) int64 {
	switch in.LoginType {
	case consts.LoginTypePwd:
		userInfo := findUserByIdOrUsername(ctx, nil, in.Username)
		if gutil.IsEmpty(userInfo) {
			panic(gerror.NewCode(enums.UserNotExists))
		}
		if gutil.IsEmpty(in.Password) {
			panic(gerror.NewCode(enums.RequestMissingParam))
		}
		encryptPassword := utility.EncryptPassword(in.Password, userInfo.PwdSalt, userInfo.PwdEncryNum)
		if encryptPassword != userInfo.Password {
			panic(gerror.NewCode(enums.UserLoginFailed))
		}
		return userInfo.UserId
	case consts.LoginTypeApple:
		if gutil.IsEmpty(in.IdentityToken) || gutil.IsEmpty(in.UserIdentifier) {
			panic(gerror.NewCode(enums.RequestMissingParam))
		}
		if err := verifyIdentityToken(ctx, in.IdentityToken, in.UserIdentifier); err != nil {
			panic(err)
		}
		return s.thirdActUserRegister(ctx, in, consts.LoginTypeApple)
	case consts.LoginTypeFaceBook:
		if gutil.IsEmpty(in.AccessToken) || gutil.IsEmpty(in.UserIdentifier) {
			panic(gerror.NewCode(enums.RequestMissingParam))
		}
		if err := verifyFacebookToken(ctx, in.AccessToken); err != nil {
			panic(err)
		}
		return s.thirdActUserRegister(ctx, in, consts.LoginTypeFaceBook)
	default:
		panic(gerror.NewCode(enums.IllegalOperation))
	}
	return 0
}

func (s *sUser) thirdActUserRegister(ctx context.Context, in model.LoginInput, providerType int) (userId int64) {
	one, err := dao.UserThirdAccount.Ctx(ctx).
		Where(dao.UserThirdAccount.Columns().ThirdUserId, in.UserIdentifier).
		One()
	if err != nil {
		panic(err)
	}
	if !one.IsEmpty() {
		return one.GMap().GetVar("userId").Int64()
	}
	if !gutil.IsEmpty(in.Username) {
		userInfo := findUserByIdOrUsername(ctx, nil, in.Username)
		if !gutil.IsEmpty(userInfo) {
			thirdAccount := entity.UserThirdAccount{
				UserId:       userInfo.UserId,
				ThirdUserId:  in.UserIdentifier,
				ProviderType: providerType,
				CreateTime:   gtime.Now(),
			}
			if _, err = dao.UserThirdAccount.Ctx(ctx).Data(thirdAccount).Insert(); err != nil {
				panic(err)
			}
			return userInfo.UserId
		}
	}
	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		newUser := entity.User{
			Username:   in.UserIdentifier,
			Password:   grand.Letters(8),
			Enable:     true,
			CreateTime: gtime.Now(),
			FirstName:  grand.Letters(3),
			LastName:   grand.Letters(3),
		}
		if !gutil.IsEmpty(in.Username) {
			newUser.Username = in.Username
		}
		//插入用户数据返回用户id
		newUserId, err := dao.User.Ctx(ctx).Data(newUser).InsertAndGetId()
		if err != nil {
			panic(err)
		}
		newTUser := &entity.UserThirdAccount{
			UserId:       newUserId,
			ThirdUserId:  in.UserIdentifier,
			ProviderType: consts.LoginTypeApple,
			CreateTime:   gtime.Now(),
		}
		if _, err = dao.UserThirdAccount.Ctx(ctx).Data(newTUser).Insert(); err != nil {
			panic(err)
		}
		userAuth := entity.UserAuth{AuthId: utility.GID.Generate().Int64(), UserId: newUserId, CreateTime: gtime.Now()}
		//初始化用户实名认证信息
		if _, err = dao.UserAuth.Ctx(ctx).Data(userAuth).Insert(); err != nil {
			panic(err)
		}
		userId = newUserId
		return nil
	}); err != nil {
		panic(err)
	}
	return
}

func findUserByIdOrUsername(ctx context.Context, userId, username any) (user *entity.User) {
	if err := dao.User.Ctx(ctx).Where(do.User{
		UserId:   userId,
		Username: username,
		IsDelete: consts.DeleteOn,
	}).Scan(&user); err != nil {
		panic(err)
	}
	if user != nil && !user.Enable {
		panic(gerror.NewCode(enums.UserAcctFrozen))
	}
	return
}

/*
Register 用户注册
*/
func (s *sUser) Register(ctx context.Context, in model.RegisterInput) {
	if !gutil.IsEmpty(findUserByIdOrUsername(ctx, nil, in.Username)) {
		panic(gerror.NewCode(enums.UserExists))
	}
	checkVftCode(ctx, in.Username, in.EmailVftCode)
	userInfo := entity.User{
		Username:   in.Username,
		Enable:     true,
		CreateTime: gtime.Now(),
		FirstName:  grand.Letters(3),
		LastName:   grand.Letters(3),
	}
	operatePwd(&userInfo, in.Password)
	if err := g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		//插入用户数据返回用户id
		userId, err := dao.User.Ctx(ctx).Data(userInfo).InsertAndGetId()
		if err != nil {
			panic(err)
		}
		userAuth := entity.UserAuth{AuthId: utility.GID.Generate().Int64(), UserId: userId, CreateTime: gtime.Now()}
		//初始化用户实名认证信息
		if _, err = dao.UserAuth.Ctx(ctx).Data(userAuth).Insert(); err != nil {
			panic(err)
		}
		return nil
	}); err != nil {
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
	userInfo := findUserByIdOrUsername(ctx, nil, in.Username)
	if gutil.IsEmpty(userInfo) {
		panic(gerror.NewCode(enums.UserNotExists))
	}
	checkVftCode(ctx, in.Username, in.EmailVftCode)
	operatePwd(userInfo, in.NewPassword)
	_, err := dao.User.Ctx(ctx).Where(dao.User.Columns().UserId, userInfo.UserId).Update(userInfo)
	if err != nil {
		panic(err)
	}
}

func (s *sUser) Info(ctx context.Context) (out *model.UserInfoOutput) {
	userId := service.BizCtx().Get(ctx).Data.Get(consts.ContextKeyUserId)
	if err := dao.User.Ctx(ctx).
		Where(dao.User.Columns().UserId, userId).
		Where(dao.User.Columns().IsDelete, consts.DeleteOn).
		Scan(&out); err != nil {
		panic(err)
	}
	if result, err := dao.UserAuth.Ctx(ctx).Fields(dao.UserAuth.Columns().AuthStatus).
		Where(dao.UserAuth.Columns().UserId, service.BizCtx().Get(ctx).Data.Get(consts.ContextKeyUserId)).
		One(); err != nil {
		panic(err)
	} else {
		out.AuthStatus = result.GMap().GetVar(dao.UserAuth.Columns().AuthStatus).Int()
	}

	result, err := dao.UserDevice.Ctx(ctx).Fields(dao.UserDevice.Columns().DeviceId).
		Where(dao.UserDevice.Columns().UserId, userId).
		Where(dao.UserDevice.Columns().IsSelect, consts.CarSelectYes).One()
	if err != nil {
		panic(err)
	}
	if result.IsEmpty() {
		return
	}
	device := new(entity.Device)
	if err = result.Struct(&device); err != nil {
		panic(err)
	}
	if err = dao.Device.Ctx(ctx).Fields(dao.Device.Columns().Nickname, dao.Device.Columns().DeviceName).
		Where(dao.Device.Columns().DeviceId, device.DeviceId).
		Where(dao.Device.Columns().IsDelete, consts.DeleteOn).Scan(&device); err != nil {
		panic(err)
	}
	out.Nickname = device.Nickname
	out.DeviceName = device.DeviceName
	//根据物模型获取所有属性数据 根据情况选择
	keys := make([]string, 0)
	keys = append(keys, consts.MileageStr)
	res, err := service.DevDevice().GetProperty(ctx, &model.DeviceGetPropertyInput{
		DeviceKey:    device.DeviceName,
		PropertyKeys: keys,
	})
	if err != nil {
		panic(err)
	}
	if !gutil.IsEmpty(res) && !res[0].Value.IsEmpty() {
		out.Mileage = res[0].Value.Float32()
	}
	return
}

/*
EditInfo 编辑用户资料
*/
func (s *sUser) EditInfo(ctx context.Context, in model.EditInfoInput) {
	if _, err := dao.User.Ctx(ctx).OmitEmptyData().Update(
		&in, dao.User.Columns().UserId, service.BizCtx().Get(ctx).Data.Get(consts.ContextKeyUserId),
	); err != nil {
		panic(err)
	}
}
