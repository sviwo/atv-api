package cmd

import (
	"context"
	"fmt"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gutil"
	"net/http"
	"sviwo/internal/consts"
	"sviwo/internal/consts/enums"
	"sviwo/internal/model"
	"sviwo/internal/service"
	"sviwo/pkg/utility/encrypt"
	"sviwo/pkg/utility/response"
)

// 登录gtoken
func StartGToken(ctx context.Context) *gtoken.GfToken {
	gToken := new(gtoken.GfToken)
	err := g.Cfg().MustGet(ctx, "gfToken").Struct(gToken)
	if err != nil {
		panic(err)
	}
	gToken.LoginBeforeFunc = loginBeforeFunc
	gToken.LoginAfterFunc = loginAfterFunc
	gToken.LogoutAfterFunc = logoutAfterFunc
	gToken.AuthAfterFunc = AuthAfterFunc
	return gToken
}

// 自定义登录验证方法
func loginBeforeFunc(r *ghttp.Request) (string, interface{}) {
	username := r.Get("username").String()
	loginType := r.Get("loginType").Int()
	g.Log().Infof(r.GetCtx(), "====login username======%s;", username)
	if gutil.IsEmpty(username) || gutil.IsEmpty(loginType) {
		response.JsonExit(r, enums.RequestMissingParam, nil)
	}
	return service.User().Login(
		r.GetCtx(),
		model.LoginInput{
			LoginType:      loginType,
			Username:       username,
			Password:       r.Get("password").String(),
			IdentityToken:  r.Get("identityToken").String(),
			UserIdentifier: r.Get("userIdentifier").String(),
			AccessToken:    r.Get("accessToken").String(),
		},
	), nil
}

// 自定义的登录返回方法
func loginAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	if !respData.Success() {
		response.JsonExit(r, enums.UserLoginFailed, nil)
	} else {
		userKey := respData.GetString(gtoken.KeyUserKey)
		token := respData.GetString(gtoken.KeyToken)
		m := gmap.New()
		m.Set("Authorization", token)

		eccKey, err := ecc.GenerateEccKeyHex()
		if err != nil {
			panic(err)
		}
		m.Set("PublicKey", eccKey.PublicKey)
		_, err = g.Redis().Set(r.GetCtx(), fmt.Sprintf(consts.RedisEccPrivateKey, userKey), eccKey.PrivateKey)
		if err != nil {
			panic(err)
		}
		response.SuccessMsg(r, m)
	}
}

// 自定义的登出返回方法
func logoutAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	if respData.Success() {
		response.SuccessMsg(r, nil)
	} else {
		response.FailMsg(r)
	}
}

// 自定义认证返回方法
func AuthAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	if respData.Success() {
		r.Middleware.Next()
	} else {
		r.Response.Status = http.StatusUnauthorized
		response.JsonExit(r, enums.LoginOverdue, nil)
	}
}
