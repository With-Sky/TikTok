package global

import (
	"go.uber.org/zap"
	utils "tiktok/pkg/utils"
)

var (
	Config = utils.ConfigInit("TIKTOK_FEED", "feedConfig")
	LOG    *zap.Logger
)
