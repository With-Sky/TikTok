package main

import (
	"context"
	"errors"
	"fmt"
	"tiktok/cmd/message_service/global"
	"tiktok/cmd/message_service/pack"
	"tiktok/cmd/message_service/service"
	"tiktok/kitex_gen/message"
	"tiktok/pkg/utils"
)

// MassageServiceImpl implements the last service interface defined in the IDL.
type MassageServiceImpl struct{}

// FriendList implements the MassageServiceImpl interface.
func (s *MassageServiceImpl) FriendList(ctx context.Context, req *message.FriendListReq) (*message.FriendListRes, error) {
	var resp = new(message.FriendListRes)
	friendUser, err := service.FriendList(ctx, req.UserId)
	fmt.Println(friendUser)
	if err != nil {
		global.LOG.Error(err.Error())
		err = errors.New("获取错误")
		resp.BaseResp = pack.BuildBaseResp(err, 400)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(err, 400)
	resp.UserList = friendUser
	return resp, nil
}

// SendMessage implements the MassageServiceImpl interface.
func (s *MassageServiceImpl) SendMessage(ctx context.Context, req *message.MessageActionReq) (*message.MessageActionRes, error) {
	resp := new(message.MessageActionRes)
	j := new(utils.JWT)
	userId, err := j.GetIdByToken(req.Token, global.Config)
	if err != nil {
		global.LOG.Error(err.Error())
		err = errors.New("解析token错误")
		resp.BaseResp = pack.BuildBaseResp(err, 400)
		return resp, nil
	}
	err = service.SendMessage(ctx, req.Content, req.ToUserId, userId)
	resp.BaseResp = pack.BuildBaseResp(err, 400)
	return resp, nil
}

// MessageList implements the MassageServiceImpl interface.
func (s *MassageServiceImpl) MessageList(ctx context.Context, req *message.MessageChatReq) (*message.MessageChatRes, error) {
	resp := new(message.MessageChatRes)
	messageList, err := service.MessageList(ctx, req)
	if err != nil {
		global.LOG.Error(err.Error())
		err = errors.New("解析token错误")
		resp.BaseResp = pack.BuildBaseResp(err, 400)
		return resp, nil
	}
	resp.MessageList = messageList
	resp.BaseResp = pack.BuildBaseResp(err, 400)
	return resp, nil
}
