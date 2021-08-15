package global

import (
	"CommonConfig/config"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GORM   *gorm.DB
	CONFIG config.Server
	VIPER  *viper.Viper
	LOG    *zap.Logger
	REDIS  *redis.Client
)
