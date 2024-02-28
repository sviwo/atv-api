package user

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
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

type sUser struct{}

/*
Login 执行登录
*/
func (s *sUser) Login(ctx context.Context, in model.LoginInput) (err error, userId *uint64) {
	userInfo := findUserByUsername(ctx, in.Username)
	if gutil.IsEmpty(userInfo) {
		panic(gerror.NewCode(rcode.UserNotExists))
	}
	if !userInfo.Enable {
		panic(gerror.NewCode(rcode.UserAcctFrozen))
	}
	//第三方登陆
	if consts.LOGIN_TYPE_THIRD == in.LoginType {
		//todo 后面完善相关逻辑
	} else {
		encryptPassword := utility.EncryptPassword(in.Password, userInfo.PwdSalt, userInfo.PwdEncryNum)
		if encryptPassword != userInfo.Password {
			return gerror.NewCode(rcode.UserLoginFailed), nil
		}
	}
	return nil, &userInfo.UserId
}

func findUserByUsername(ctx context.Context, username string) (user *entity.User) {
	err := dao.User.Ctx(ctx).Where("username", username).Where("is_delete", 0).Scan(&user)
	if err != nil {
		panic(err)
	}
	return
}

/*
Register 用户注册
*/
func (s *sUser) Register(ctx context.Context, in model.RegisterInput) error {
	if !gutil.IsEmpty(findUserByUsername(ctx, in.Username)) {
		return gerror.NewCode(rcode.UserExists)
	}
	checkVftCode(ctx, in.Username, in.EmailVftCode)
	userInfo := entity.User{Username: in.Username, Enable: true, CreateTime: gtime.Now(), IsDelete: false}
	operatePwd(&userInfo, in.Password)
	//插入用户数据返回用户id
	_, err := dao.User.Ctx(ctx).Data(userInfo).InsertAndGetId()
	if err != nil {
		panic(err)
	}
	return nil
}

// 检查验证码
func checkVftCode(ctx context.Context, username, emailVftCode string) {
	value, err := g.Redis().Get(ctx, fmt.Sprintf(consts.RedisEmailVftCode, username))
	if err != nil {
		panic(err)
	}
	if value.IsEmpty() {
		panic(gerror.NewCode(rcode.VftCodeOverdue))
	}
	if value.String() != emailVftCode {
		panic(gerror.NewCode(rcode.IllegalArgument))
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
func (s *sUser) UpdatePassword(ctx context.Context, in model.UpdatePasswordInput) error {
	userInfo := findUserByUsername(ctx, in.Username)
	if gutil.IsEmpty(userInfo) {
		return gerror.NewCode(rcode.UserNotExists)
	}
	checkVftCode(ctx, in.Username, in.EmailVftCode)
	operatePwd(userInfo, in.NewPassword)
	_, err := dao.User.Ctx(ctx).Where("user_id", userInfo.UserId).Update(userInfo)
	if err != nil {
		panic(err)
	}
	return nil
}

func (s *sUser) Info(ctx context.Context, in model.UserInfoInput) (out *model.UserInfoOutput) {
	err := dao.User.Ctx(ctx).Where("user_id", in.UserId).Scan(&out)
	if err != nil {
		panic(err)
	}
	return
}

/*
EditInfo 编辑用户资料
*/
func (s *sUser) EditInfo(ctx context.Context, in model.EditInfoInput) error {
	_, err := dao.User.Ctx(ctx).Where("user_id", in.UserId).Update(in)
	if err != nil {
		panic(err)
	}
	return nil
}
