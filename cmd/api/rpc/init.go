package rpc

import (
	"tiktok/cmd/api/global"
)

func InitRpc() {
	InitCommentRpc()
	InitFavoriteRpc()
	InitFeedRpc()
	InitUserRpc()
	InitPublishRpc()
	InitRelationRpc()
	InitMessageRpc()
	global.LOG.Info("成功初始化rpc")
}
