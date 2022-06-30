package main

import (
	_ "my_goframe_monorepo/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"my_goframe_monorepo/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
