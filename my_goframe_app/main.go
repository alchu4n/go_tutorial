package main

import (
	_ "my_goframe_app/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"my_goframe_app/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
