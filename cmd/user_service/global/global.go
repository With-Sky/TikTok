package global

import (
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	utils "tiktok/pkg/utils"
)

var (
	Config = utils.ConfigInit("TIKTOK_USER", "userConfig")
	LOG    *zap.Logger
	Redis  *redis.Client
)
