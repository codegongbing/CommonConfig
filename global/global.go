package global

import (
	"CommonConfig/config"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	GORM       *gorm.DB
	CONFIG config.Server
	GVA_VP     *viper.Viper
)
