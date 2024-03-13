package boot

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"gopkg.in/gomail.v2"
)

var (
	NewDialer *gomail.Dialer
)

func InitSendEmail(ctx context.Context) error {
	NewDialer = gomail.NewDialer(
		g.Cfg().MustGet(ctx, "email.host").String(),
		g.Cfg().MustGet(ctx, "email.port").Int(),
		g.Cfg().MustGet(ctx, "email.username").String(),
		g.Cfg().MustGet(ctx, "email.password").String(),
	)
	return nil
}
