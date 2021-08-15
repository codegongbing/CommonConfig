package core

import (
	"CommonConfig/global"
	"CommonConfig/initialize"
	"fmt"
	"go.uber.org/zap"
	"os"
	"strings"
)

func Server() {
	r := initialize.Router()
	initialize.Redis()
	global.GORM = initialize.Gorm()

	port := global.CONFIG.Project.Port
	global.LOG.Info("server run success on ", zap.String("address", port))
	{
		getWd, _ := os.Getwd()
		split := strings.Split(getWd, "/")
		fmt.Printf("欢迎使用%s\n", split[len(split)-1])
	}
	global.LOG.Error(r.Run(":" + port).Error())
}
