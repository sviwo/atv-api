package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"sviwo/internal/cmd"
	_ "sviwo/internal/logic"
	_ "sviwo/internal/packed"
)

func main() {
	cmd.Main.Run(gctx.New())
}
