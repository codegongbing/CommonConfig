package main

import (
	"CommonConfig/core"
	"CommonConfig/global"
)



// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	global.VIPER = core.Viper("config.toml")
	global.LOG = core.Zap()
	core.Server()
}
