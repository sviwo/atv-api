package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	_ "github.com/taosdata/driver-go/v3/taosRestful"
	_ "github.com/taosdata/driver-go/v3/taosWS"
	"sviwo/internal/boot"
	"sviwo/internal/cmd"
	_ "sviwo/internal/logic"
	_ "sviwo/internal/packed"
)

func main() {
	ctx := gctx.GetInitCtx()
	boot.AllSystemInit(ctx)
	cmd.Main.Run(gctx.New())
}
