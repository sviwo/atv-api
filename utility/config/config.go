package config

import (
	"github.com/bwmarrin/snowflake"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"gopkg.in/gomail.v2"
)

var (
	Node      *snowflake.Node
	NewDialer *gomail.Dialer
)

func init() {
	Node, _ = snowflake.NewNode(1)
	NewDialer = initSendEmail()
}

func initSendEmail() *gomail.Dialer {
	ctx := gctx.GetInitCtx()
	return gomail.NewDialer(
		g.Cfg().MustGet(ctx, "email.host").String(),
		g.Cfg().MustGet(ctx, "email.port").Int(),
		g.Cfg().MustGet(ctx, "email.username").String(),
		g.Cfg().MustGet(ctx, "email.password").String(),
	)
}
