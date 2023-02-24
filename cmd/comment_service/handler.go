package main

import (
	"context"
	"errors"
	"fmt"
	"tiktok/cmd/comment_service/global"
	"tiktok/cmd/comment_service/pack"
	"tiktok/cmd/comment_service/service"
	comment "tiktok/kitex_gen/comment"
	"tiktok/pkg/utils"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// CommentAction implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentAction(ctx context.Context, req *comment.CommentReq) (*comment.CommentRes, error) {
	fmt.Println("视频ID：", req.VideoId)
	global.LOG.Info("评论服务")
	resp := new(comment.CommentRes)
	//解析token，获得用户ID
	j := utils.NewJWT(global.Config)
	userID, err := j.GetIdByToken(req.Token, global.Config)
	if err != nil {
		err := errors.New("token解析失败")
		global.LOG.Error("token解析失败")
		resp.BaseResp = pack.BuildBaseResp(err, 400)
		return resp, nil
	}
	//判断评论内容是否为空
	if req.CommentText == "" {
		err := errors.New("评论内容为空")
		global.LOG.Error("评论内容为空")
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
	user, commentID, err := service.CommentAction(ctx, userID, req.VideoId, req.CommentText)
	if err == nil {
		resp.Comment = &comment.Comment{
			Id:         commentID,
			User:       user,
			Content:    req.CommentText,
			CreateDate: utils.NowDate(),
		}
		global.LOG.Info("评论服务成功")
		resp.BaseResp = pack.BuildBaseResp(err, 0)
	} else {
		global.LOG.Error("评论服务失败")
		resp.BaseResp = pack.BuildBaseResp(err, 400)
	}
	return resp, nil
}

// DeleteComment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) DeleteComment(ctx context.Context, req *comment.DeleteCommentReq) (*comment.DeleteCommentRes, error) {
	global.LOG.Info("删除评论服务")
	resp := new(comment.DeleteCommentRes)
	//解析token，获得用户ID
	j := utils.NewJWT(global.Config)
	_, err := j.GetIdByToken(req.Token, global.Config)
	if err != nil {
		err := errors.New("token解析失败")
		global.LOG.Error("token解析失败")
		resp.BaseResp = pack.BuildBaseResp(err, 400)
		return resp, nil
	}
	//判断评论ID是否为空
	if req.CommentId == 0 {
		err := errors.New("评论ID为空")
		global.LOG.Error("评论ID为空")
		resp.BaseResp = pack.BuildBaseResp(err, 400)
		return resp, nil
	}
	err = service.DeleteComment(ctx, req.CommentId)
	if err == nil {
		global.LOG.Info("删除评论服务成功")
		resp.BaseResp = pack.BuildBaseResp(err, 0)
	} else {
		global.LOG.Error("删除评论服务失败")
		resp.BaseResp = pack.BuildBaseResp(err, 400)
	}
	return resp, nil
}

// CommentList implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentList(ctx context.Context, req *comment.CommentListReq) (*comment.CommentListRes, error) {
	global.LOG.Info("评论列表服务")
	resp := new(comment.CommentListRes)
	//解析token,获得用户ID
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
	commentList, err := service.CommentList(ctx, userID, req.VideoId)
	if err == nil {
		global.LOG.Info("评论列表服务成功")
		resp.CommentList = commentList
		resp.BaseResp = pack.BuildBaseResp(err, 0)
	} else {
		global.LOG.Error("评论列表服务失败")
		resp.BaseResp = pack.BuildBaseResp(err, 400)
	}
	return resp, nil
}
