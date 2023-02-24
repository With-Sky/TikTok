package main

import (
	"context"
	"errors"
	"tiktok/cmd/favorite_service/global"
	"tiktok/cmd/favorite_service/pack"
	"tiktok/cmd/favorite_service/service"
	"tiktok/kitex_gen/favorite"
	"tiktok/pkg/utils"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// Favorite implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) Favorite(ctx context.Context, req *favorite.FavoriteReq) (*favorite.FavoriteRes, error) {
	global.LOG.Info("点赞服务")
	resp := new(favorite.FavoriteRes)
	//解析token，获得用户ID
	j := utils.NewJWT(global.Config)
	userID, err := j.GetIdByToken(req.Token, global.Config)
	if err != nil {
		err := errors.New("token解析失败")
		global.LOG.Error("token解析失败")
		resp.BaseResp = pack.BuildBaseResp(err, 400)
		return resp, nil
	}
	//判断视频ID是否为空
	if req.VideoId == 0 {
		err := errors.New("视频ID为空")
		global.LOG.Error("视频ID为空")
		resp.BaseResp = pack.BuildBaseResp(err, 400)
		return resp, nil
	}
	global.LOG.Info("开始调用服务")
	err = service.FavoriteAction(ctx, userID, req.VideoId)
	if err == nil {
		global.LOG.Info("点赞服务成功")
		resp.BaseResp = pack.BuildBaseResp(err, 0)
	} else {
		global.LOG.Error("点赞服务失败")
		resp.BaseResp = pack.BuildBaseResp(err, 400)
	}
	return resp, nil
}

// DeleteFavorite implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) DeleteFavorite(ctx context.Context, req *favorite.DeleteFavoriteReq) (*favorite.DeleteFavoriteRes, error) {
	global.LOG.Info("取消点赞服务")
	resp := new(favorite.DeleteFavoriteRes)
	//解析token，获得用户ID
	j := utils.NewJWT(global.Config)
	userID, err := j.GetIdByToken(req.Token, global.Config)
	if err != nil {
		err := errors.New("token解析失败")
		global.LOG.Error("token解析失败")
		resp.BaseResp = pack.BuildBaseResp(err, 400)
		return resp, nil
	}
	//判断视频ID是否为空
	if req.VideoId == 0 {
		err := errors.New("视频ID为空")
		global.LOG.Error("视频ID为空")
		resp.BaseResp = pack.BuildBaseResp(err, 400)
		return resp, nil
	}
	err = service.CancelFavorite(ctx, userID, req.VideoId)
	if err == nil {
		global.LOG.Info("取消点赞服务成功")
		resp.BaseResp = pack.BuildBaseResp(err, 0)
	} else {
		global.LOG.Error("取消点赞服务失败")
		resp.BaseResp = pack.BuildBaseResp(err, 400)
	}
	return resp, nil
}

// FavoriteList implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteList(ctx context.Context, req *favorite.FavoriteListReq) (*favorite.FavoriteListRes, error) {
	global.LOG.Info("喜欢列表")
	resp := new(favorite.FavoriteListRes)
	//解析token
	j := utils.NewJWT(global.Config)
	UserID, err := j.GetIdByToken(req.Token, global.Config)
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
	videoList, err := service.FavoriteList(ctx, UserID, req.UserId)
	if err == nil {
		resp.VideoList = videoList
		global.LOG.Info("取消点赞服务成功")
		resp.BaseResp = pack.BuildBaseResp(err, 0)
	} else {
		global.LOG.Error("取消点赞服务失败")
		resp.BaseResp = pack.BuildBaseResp(err, 400)
	}
	return resp, nil
}
