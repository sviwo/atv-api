package user

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/gogf/gf/v2/util/gutil"
	"sviwo/internal/consts"
	"sviwo/internal/dao"
	"sviwo/internal/logic/biz/enums"
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

type sUser struct {
	LoginType uint8
}

/*
Login 执行登录
*/
func (s *sUser) Login(ctx context.Context, in model.LoginInput) error {
	userInfo := s.findUserByUsername(ctx, in.Username)
	//验证码登陆
	if consts.LOGIN_TYPE_VTF == in.LoginType {
		//value, err := g.Redis().Get(ctx, fmt.Sprintf(consts.MethodUserServiceRedisKey, in.Username))
		//if err != nil {
		//	panic(err)
		//}
		//if gutil.IsEmpty(value) || gconv.String(value) != in.VftCode {
		//	return gerror.NewCode(rcode.ErrVftCode)
		//}
		if "888888" != in.VftCode {
			return gerror.NewCode(rcode.ErrVftCode)
		}
		//检查此账号是否已注册,没有注册则执行注册操作
		if gutil.IsEmpty(userInfo) {
			err := s.Register(ctx, model.RegisterInput{Username: in.Username})
			if err != nil {
				panic(err)
			}
		}

	} else {
		encryptPassword := utility.EncryptPassword(in.Password, userInfo.PwdSalt, userInfo.PwdEncryNum)
		if encryptPassword != userInfo.Password {
			return gerror.NewCode(rcode.UserLoginFailed)
		}
	}
	return nil
}

func (s *sUser) findUserByUsername(ctx context.Context, username string) (user *entity.User) {
	err := dao.User.Ctx(ctx).Where("username", username).Scan(&user)
	if err != nil {
		panic(err)
	}
	return
}

/*
Register 用户注册
*/
func (s *sUser) Register(ctx context.Context, in model.RegisterInput) error {
	if !gutil.IsEmpty(s.findUserByUsername(ctx, in.Username)) {
		return gerror.NewCode(rcode.UserExists)
	}
	userInfo := entity.User{Enable: true, CreateTime: gtime.Now(), IsDelete: false}
	if consts.LOGIN_TYPE_VTF == s.LoginType {
		in.Password = grand.S(10)
	} else {
		//处理加密盐和密码的逻辑
		UserSalt := grand.S(10)
		var pwdEncryNum = 10
		in.Password = utility.EncryptPassword(in.Password, UserSalt, pwdEncryNum)
		userInfo.PwdSalt = UserSalt
		userInfo.PwdEncryNum = pwdEncryNum
	}

	//插入用户数据返回用户id
	if err := gconv.Struct(in, &userInfo); err != nil {
		panic(err)
	}
	dao.User.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		userId, err := dao.User.Ctx(ctx).Data(userInfo).InsertAndGetId()
		if err != nil {
			panic(err)
		}
		//初始化用户账户
		userAccount := entity.UserAccount{UserId: userId, CreateTime: gtime.Now(), IsDelete: false}
		_, err = dao.UserAccount.Ctx(ctx).Data(userAccount).Insert()
		if err != nil {
			panic(err)
		}
		return nil
	})
	return nil
}

/*
UpdatePassword 修改密码
*/
func (s *sUser) UpdatePassword(ctx context.Context, in model.UpdatePasswordInput) error {
	var user *entity.User
	err := dao.User.Ctx(ctx).Where("user_id", in.UserId).Scan(&user)
	if err != nil {
		panic(err)
	}
	if user == nil {
		return gerror.NewCode(rcode.UserNotExists)
	}
	//处理加密盐和密码的逻辑
	PwdSalt := grand.S(10)
	var pwdEncryNum = 10
	newPassword := utility.EncryptPassword(in.Password, PwdSalt, pwdEncryNum)

	user.Password = newPassword
	user.PwdSalt = PwdSalt
	user.PwdEncryNum = pwdEncryNum
	_, err = dao.User.Ctx(ctx).Where("user_id", user.UserId).Update(user)
	if err != nil {
		panic(err)
	}
	return nil
}

func (s *sUser) Info(ctx context.Context, in model.UserInfoInput) (out *model.UserInfoOutput) {
	err := dao.User.Ctx(ctx).Where("username", in.Username).Scan(&out)
	if err != nil {
		panic(err)
	}
	return
}

/*
EditInfo 编辑用户资料
*/
func (s *sUser) EditInfo(ctx context.Context, in model.EditInfoInput) error {
	user := entity.User{}
	err := gconv.Struct(in, &user)
	if err != nil {
		panic(err)
	}
	_, err = dao.User.Ctx(ctx).Update(user)
	if err != nil {
		panic(err)
	}
	return nil
}
