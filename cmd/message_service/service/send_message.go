package service

import (
	"context"
	"errors"
	"tiktok/cmd/message_service/dal/db"
	"tiktok/cmd/message_service/global"
)

func SendMessage(ctx context.Context, content string, toUserId int64, userId int64) error {
	err := db.CreateMessage(ctx, content, toUserId, userId)
	if err != nil {
		err = errors.New("发送信息失败")
		global.LOG.Error(err.Error())
		return err
	}
	return nil
}
