package main

import (
	"CommonConfig/core"
	"CommonConfig/global"
	"fmt"
)

func main() {
	global.VIPER = core.Viper("config.toml")
	global.LOG = core.Zap()
	fmt.Println(global.CONFIG.Zap)
	core.Server()
}
