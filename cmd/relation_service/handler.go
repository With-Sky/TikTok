package main

import (
	"context"
	"errors"
	"fmt"
	"tiktok/cmd/relation_service/global"
	"tiktok/cmd/relation_service/pack"
	service "tiktok/cmd/relation_service/service"
	relation "tiktok/kitex_gen/relation"
	"tiktok/pkg/utils"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// Follow implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) Follow(ctx context.Context, req *relation.FollowReq) (*relation.FollowRes, error) {
	global.LOG.Info("关注服务")
	resp := new(relation.FollowRes)
	//解析token
	j := utils.NewJWT(global.Config)
	userID, err := j.GetIdByToken(req.Token, global.Config)
	if err != nil {
		err := errors.New("token解析失败")
		global.LOG.Error("token解析失败")
		resp.BaseResp = pack.BuildBaseResp(err, 400)
		return resp, nil
	}
	//判断用户ID是否为空
	if req.ToUserId == 0 {
		err := errors.New("用户ID为空")
		global.LOG.Error("用户ID为空")
		resp.BaseResp = pack.BuildBaseResp(err, 400)
		return resp, nil
	}
	fmt.Println("uid", userID, " touid", req.ToUserId)
	err = service.FollowAction(ctx, userID, req.ToUserId)
	if err == nil {
		global.LOG.Info("关注服务成功")
		resp.BaseResp = pack.BuildBaseResp(err, 0)
	} else {
		global.LOG.Error("关注服务失败")
		resp.BaseResp = pack.BuildBaseResp(err, 400)
	}
	return resp, nil
}

// CancelFollow implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) CancelFollow(ctx context.Context, req *relation.CancelFollowReq) (*relation.CancelFollowRes, error) {
	global.LOG.Info("取消关注服务")
	resp := new(relation.CancelFollowRes)
	//解析token
	j := utils.NewJWT(global.Config)
	userID, err := j.GetIdByToken(req.Token, global.Config)
	if err != nil {
		err := errors.New("token解析失败")
		global.LOG.Error("token解析失败")
		resp.BaseResp = pack.BuildBaseResp(err, 400)
		return resp, nil
	}
	fmt.Println("uid", userID, " touid", req.ToUserId)
	//判断用户ID是否为空
	if req.ToUserId == 0 {
		err := errors.New("用户ID为空")
		global.LOG.Error("用户ID为空")
		resp.BaseResp = pack.BuildBaseResp(err, 400)
		return resp, nil
	}
	err = service.CancelFollow(ctx, userID, req.ToUserId)
	if err == nil {
		global.LOG.Info("取消关注服务成功")
		resp.BaseResp = pack.BuildBaseResp(err, 0)
	} else {
		global.LOG.Error("取消关注服务失败")
		resp.BaseResp = pack.BuildBaseResp(err, 400)
	}
	return resp, nil
}

// FollowList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) FollowList(ctx context.Context, req *relation.FollowListReq) (*relation.FollowListRes, error) {
	global.LOG.Info("关注列表服务")
	resp := new(relation.FollowListRes)
	//解析token
	j := utils.NewJWT(global.Config)
	userID, err := j.GetIdByToken(req.Token, global.Config)
	if err != nil {
		err := errors.New("token解析失败")
		global.LOG.Error("token解析失败")
		resp.BaseResp = pack.BuildBaseResp(err, 400)
		return resp, nil
	}
	//判断用户ID是否为空
	if req.UserId == 0 {
		err := errors.New("用户ID为空")
		global.LOG.Error("用户ID为空")
		resp.BaseResp = pack.BuildBaseResp(err, 400)
		return resp, nil
	}
	userList, err := service.FollowList(ctx, userID, req.UserId)
	if err == nil {
		resp.UserList = userList
		global.LOG.Info("关注列表服务成功")
		resp.BaseResp = pack.BuildBaseResp(err, 0)
	} else {
		global.LOG.Error("关注列表服务失败")
		resp.BaseResp = pack.BuildBaseResp(err, 400)
	}
	return resp, nil
}

// FollowerList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) FollowerList(ctx context.Context, req *relation.FollowerListReq) (*relation.FollowerListRes, error) {
	global.LOG.Info("粉丝列表服务")
	resp := new(relation.FollowerListRes)
	//解析token
	j := utils.NewJWT(global.Config)
	userID, err := j.GetIdByToken(req.Token, global.Config)
	if err != nil {
		err := errors.New("token解析失败")
		global.LOG.Error("token解析失败")
		resp.BaseResp = pack.BuildBaseResp(err, 400)
		return resp, nil
	}
	//判断用户ID是否为空
	if req.UserId == 0 {
		err := errors.New("用户ID为空")
		global.LOG.Error("用户ID为空")
		resp.BaseResp = pack.BuildBaseResp(err, 400)
		return resp, nil
	}
	userList, err := service.FollowerList(ctx, userID, req.UserId)
	if err == nil {
		resp.UserList = userList
		global.LOG.Info("粉丝列表服务成功")
		resp.BaseResp = pack.BuildBaseResp(err, 0)
	} else {
		global.LOG.Error("粉丝列表服务失败")
		resp.BaseResp = pack.BuildBaseResp(err, 400)
	}
	return resp, nil
}

// FriendList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) FriendList(ctx context.Context, req *relation.FriendListReq) (*relation.FriendListRes, error) {
	global.LOG.Info("好友列表服务")
	resp := new(relation.FriendListRes)
	//解析token
	j := utils.NewJWT(global.Config)
	userID, err := j.GetIdByToken(req.Token, global.Config)
	if err != nil {
		err := errors.New("token解析失败")
		global.LOG.Error("token解析失败")
		resp.BaseResp = pack.BuildBaseResp(err, 400)
		return resp, nil
	}
	//判断用户ID是否为空
	if req.UserId == 0 {
		err := errors.New("用户ID为空")
		global.LOG.Error("用户ID为空")
		resp.BaseResp = pack.BuildBaseResp(err, 400)
		return resp, nil
	}
	userList, err := service.FriendList(ctx, userID, req.UserId)
	if err == nil {
		resp.UserList = userList
		global.LOG.Info("好友列表服务成功")
		resp.BaseResp = pack.BuildBaseResp(err, 0)
	} else {
		global.LOG.Error("好友列表服务失败")
		resp.BaseResp = pack.BuildBaseResp(err, 400)
	}
	return resp, nil
}
