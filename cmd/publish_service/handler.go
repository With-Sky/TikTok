package main

import (
	"context"
	"errors"
	"fmt"
	"tiktok/cmd/publish_service/global"
	"tiktok/cmd/publish_service/pack"
	"tiktok/cmd/publish_service/service"
	publish "tiktok/kitex_gen/publish"
	"tiktok/pkg/utils"
)

// PublishServiceImpl implements the last service interface defined in the IDL.
type PublishServiceImpl struct{}

// Publish implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) Publish(ctx context.Context, req *publish.PublishReq) (*publish.PublishRes, error) {
	resp := new(publish.PublishRes)
	//解析token
	j := utils.NewJWT(global.Config)
	userId, err := j.GetIdByToken(req.Token, global.Config)
	fmt.Println("token", req.Token)
	fmt.Println(userId)
	if err != nil {
		global.LOG.Error(err.Error())
		err := errors.New("解析token失败")
		resp.BaseResp = pack.BuildBaseResp(err, 400)
		return resp, nil
	}
	//判断标题是否为空
	if req.Title == "" {
		err := errors.New("标题为空")
		global.LOG.Error(err.Error())
		resp.BaseResp = pack.BuildBaseResp(err, 400)
		return resp, nil
	}
	////判断视频流是否为空
	err = utils.Verify(req.Data, utils.EmptyAppVerify)
	if req.Data == nil {
		err := errors.New("视频为空")
		global.LOG.Error(err.Error())
		resp.BaseResp = pack.BuildBaseResp(err, 400)
		return resp, nil
	}
	//判断文件类型
	err = service.PublishAction(ctx, userId, req.Data, req.Title)
	if err != nil {
		global.LOG.Error(err.Error())
		err := errors.New("发布失败")
		resp.BaseResp = pack.BuildBaseResp(err, 400)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(err, 0)
	return resp, nil
}

// PublishList implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishList(ctx context.Context, req *publish.PublishListReq) (*publish.PublishListRes, error) {
	resp := new(publish.PublishListRes)

	//解析token
	//j := utils.NewJWT(global.Config)
	//c, err := j.ParseToken(req.Token, global.Config)
	//if err != nil {
	//	resp.BaseResp = pack.BuildBaseResp(err, "未登录或者登录过期", 404)
	//	return resp, nil
	//}
	//判断用户ID是否为空
	if req.UserId == 0 {
		err := errors.New("未登录或者登录过期")
		global.LOG.Error(err.Error())
		resp.BaseResp = pack.BuildBaseResp(err, 400)
		return resp, nil
	}
	videoList, err := service.PublishList(ctx, req.UserId, req.UserId)
	if err != nil {
		global.LOG.Error(err.Error())
		err := errors.New("获取视频列表失败")
		resp.BaseResp = pack.BuildBaseResp(err, 400)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(err, 0)
	resp.VideoList = videoList
	return resp, nil
}
