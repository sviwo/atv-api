package user

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"sviwo/internal/consts/enums"
)

// Facebook /debug_token 结构体
type FacebookData struct {
	Data struct {
		AppID       string `json:"app_id"`
		Type        string `json:"type"`
		Application string `json:"application"`
		ExpiresAt   int64  `json:"expires_at"`
		IsValid     bool   `json:"is_valid"`
		IssuedAt    int    `json:"issued_at"`
		Metadata    struct {
			AuthType string `json:"auth_type"`
		} `json:"metadata"`
		Scopes []string `json:"scopes"`
		UserID string   `json:"user_id"`
	} `json:"data"`
}

// verifyFacebookToken 校验 Facebook 的用户访问令牌
func verifyFacebookToken(ctx context.Context, userAccessToken string) error {
	// 构建应用访问令牌
	appToken := fmt.Sprintf(
		"%s|%s",
		g.Cfg().MustGet(ctx, "facebookLogin.appID").String(),
		g.Cfg().MustGet(ctx, "facebookLogin.appSecret").String(),
	)
	// 验证用户访问令牌的 URL
	fbTokenUrl := fmt.Sprintf(
		g.Cfg().MustGet(ctx, "facebookLogin.keyReqUrl").String()+"?input_token=%s&access_token=%s",
		userAccessToken,
		appToken,
	)
	// 使用 gclient 发送请求
	client := g.Client()
	clientProxy := g.Cfg().MustGet(ctx, "clientProxy")
	if !clientProxy.IsEmpty() {
		client.Proxy(clientProxy.String())
	}
	resp, err := client.Get(ctx, fbTokenUrl)
	if err != nil {
		return err
	}
	defer func(resp *gclient.Response) {
		err = resp.Close()
		if err != nil {
			panic(err)
		}
	}(resp)
	// 解析 JSON 响应
	data := new(FacebookData)
	if err = gconv.Scan(resp.ReadAllString(), &data); err != nil {
		return err
	}
	if !data.Data.IsValid || data.Data.ExpiresAt < gtime.Now().Timestamp() {
		return gerror.NewCode(enums.LoginOverdue)
	}
	return nil
}
