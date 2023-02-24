package global

import (
	"go.uber.org/zap"
	"tiktok/pkg/utils"
)

var (
	Config = utils.ConfigInit("TIKTOK_COMMENT", "commentConfig")
	LOG    *zap.Logger
)
