package utils

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"os"
)

func Redis(config Config, LOG *zap.Logger) *redis.Client {
	//redisCfg := global.TiK_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     config.Viper.GetString("redis.Addr"),
		Password: config.Viper.GetString("redis.Password"), // no password set
		DB:       config.Viper.GetInt("redis.DB"),          // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		LOG.Error("redis connect ping failed, err:", zap.Any("err", err))
		os.Exit(0)
		return nil
	} else {
		LOG.Info("redis connect ping response:", zap.String("pong", pong))
		return client
	}
}

//func get()
