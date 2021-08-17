package main

import (
	"CommonConfig/core"
	"CommonConfig/global"
)

// @title CommonConfig
// @version v1
// @description This is a common web project config
// @license.name Apache 2.0
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	global.VIPER = core.Viper("config.toml")
	global.LOG = core.Zap()
	core.Server()
}
