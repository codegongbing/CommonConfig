package initialize

import (
	"CommonConfig/global"
	"context"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func Redis() {
	redisConfig := global.CONFIG.Redis
	client := redis.NewClient(
		&redis.Options{
			Addr:     redisConfig.Addr,
			Password: redisConfig.Password, // no password set
			DB:       redisConfig.DB,       // use default DB
		},
	)
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.LOG.Error("redis connect ping failed, err:", zap.Any("err", err))
	} else {
		global.LOG.Info("redis connect ping response:", zap.String("pong", pong))
		global.REDIS = client
	}
}
