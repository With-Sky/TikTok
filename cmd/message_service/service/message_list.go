package service

import (
	"context"
	"errors"
	"fmt"
	"tiktok/cmd/message_service/dal/db"
	"tiktok/cmd/message_service/global"
	"tiktok/kitex_gen/message"
	"tiktok/pkg/utils"
)

func MessageList(ctx context.Context, req *message.MessageChatReq) ([]*message.Message, error) {
	fmt.Println(req.Token)
	j := new(utils.JWT)
	userId, err := j.GetIdByToken(req.Token, global.Config)
	if err != nil {
		var msg []*message.Message
		err = errors.New("token解析错误，请重新登录")
		global.LOG.Error(err.Error())
		return msg, err
	}
	msg, err := db.GetMessageList(ctx, req.PreMsgTime, req.ToUserId, userId)
	if err != nil {
		var msg []*message.Message
		err = errors.New("获取数据库数据失败")
		global.LOG.Error(err.Error())
		return msg, err
	}
	var messages = make([]*message.Message, len(msg))
	for i := 0; i < len(msg); i++ {
		var m message.Message
		if msg[i].MsgType == 1 {

			m = message.Message{
				Id:         int64(msg[i].ID),
				ToUserId:   msg[i].ToUserId,
				FromUserId: msg[i].FromUserId,
				Content:    msg[i].Content,
				CreateTime: utils.TimeToFormatData(msg[i].CreatedAt),
			}

		} else if msg[i].MsgType == 0 {
			m = message.Message{
				Id:         int64(msg[i].ID),
				ToUserId:   msg[i].FromUserId,
				FromUserId: msg[i].ToUserId,
				Content:    msg[i].Content,
				CreateTime: utils.TimeToFormatData(msg[i].CreatedAt),
			}
		}
		messages = append(messages, &m)

	}

	fmt.Println(messages)
	return messages, err
}
