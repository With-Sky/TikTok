package global

import (
	"go.uber.org/zap"
	utils "tiktok/pkg/utils"
)

var (
	Config = utils.ConfigInit("TIKTOK_RELATION", "relationConfig")
	LOG    *zap.Logger
)
