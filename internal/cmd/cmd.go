package cmd

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/os/gcmd"
	"os"
	"os/signal"
	"sviwo/internal/boot"
	"sviwo/internal/consts"
	"sviwo/internal/logic/tdengine"
	"syscall"
	"time"
)

var (
	Main = gcmd.Command{
		Name:  consts.ProjectName,
		Usage: consts.ProjectUsage,
		Brief: consts.ProjectBrief,
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			var signalChannel = make(chan os.Signal, 1)

			RunServer(ctx, signalChannel)

			boot.Boot(ctx)

			signal.Notify(signalChannel, os.Interrupt, os.Kill, syscall.SIGTERM)
			fmt.Println("收到关闭服务信号:", <-signalChannel)
			time.Sleep(time.Second * 3)
			tdengine.Close()
			fmt.Println("成功关闭服务器")

			return
		},
	}
)
