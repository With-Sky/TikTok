package controllor

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"tiktok/cmd/api/global"
	"tiktok/cmd/api/rpc"
	"tiktok/kitex_gen/comment"
)

// CommentAction 评论操作
func CommentAction(con echo.Context) error {
	// 从上下文获取请求
	var commentRequestData CommentParam
	commentRequestData.CommentId = con.QueryParam("comment_id")
	commentRequestData.CommentText = con.QueryParam("comment_text")
	commentRequestData.ActionType = con.QueryParam("action_type")
	commentRequestData.Token = con.QueryParam("token")
	commentRequestData.VideoId = con.QueryParam("video_id")
	fmt.Println("视频ID：", commentRequestData.VideoId)
	//fmt.Println(commentRequestData)

	vID, err := strconv.ParseInt(commentRequestData.VideoId, 10, 64)
	if err != nil {
		FailWithMessage("获取视频ID失败", con)
		global.LOG.Error("获取视频ID失败")
		return err
	}
	// 判断是发表评论（1），还是删除评论（2）
	if commentRequestData.ActionType == "1" {
		//fmt.Println("评论")
		// 发表评论
		// 评论服务的请求参数
		//if utils.Verify(commentRequestData.CommentText, utils.EmptyAppVerify) != nil {
		//	FailWithMessage("评论内容为空", con)
		//	//global.LOG.Error("评论内容为空")
		//	return errors.New("评论内容为空")
		//}
		commentReq := comment.CommentReq{
			Token:       commentRequestData.Token,
			CommentText: commentRequestData.CommentText,
			VideoId:     vID,
		}
		// 获得服务的返回
		res, err := rpc.Comment(context.Background(), &commentReq)
		if res == nil {
			FailWithMessage("服务请求失败", con)
			//global.LOG.Error("服务请求失败")
			return err
		}
		if err != nil {
			FailWithMessage(res.BaseResp.StatusMessage, con)
			//global.LOG.Error(res.BaseResp.StatusMessage)
			return err
		}
		// 返回响应
		if err := con.JSON(http.StatusOK, res); err != nil {
			FailWithMessage("响应失败", con)
			//global.LOG.Error("响应失败")
			return err
		}
	} else if commentRequestData.ActionType == "2" {
		//fmt.Println("删除评论")
		// 删除评论
		// 评论服务的请求参数
		cID, err := strconv.ParseInt(commentRequestData.CommentId, 10, 64)
		if err != nil {
			FailWithMessage("获取评论ID失败", con)
			//global.LOG.Error("获取评论ID失败")
			return err
		}
		//if err := utils.Verify(cID, utils.EmptyAppVerify); err != nil {
		//	FailWithMessage("评论ID为空", con)
		//	global.LOG.Error("评论ID为空")
		//	return err
		//}
		deleteCommentReq := comment.DeleteCommentReq{
			Token:     commentRequestData.Token,
			CommentId: cID,
			VideoId:   vID,
		}

		// 获得服务的返回
		res, err := rpc.DeleteComment(context.Background(), &deleteCommentReq)
		if res == nil {
			FailWithMessage("服务请求失败", con)
			//global.LOG.Error("服务请求失败")
			return err
		}
		if err != nil {
			FailWithMessage(res.BaseResp.StatusMessage, con)
			//global.LOG.Error(res.BaseResp.StatusMessage)
			return err
		}
		// 返回响应
		if err := con.JSON(http.StatusOK, res); err != nil {
			FailWithMessage("响应失败", con)
			//global.LOG.Error("响应失败")
			return err
		}
	} else {
		FailWithMessage("操作必须为评论或删除", con)
		//global.LOG.Error("操作必须为评论或删除")
		return errors.New("Action must be 1 or 2\n")
	}
	return nil
}

// CommentList 评论列表
func CommentList(con echo.Context) error {
	// 从上下文获取请求
	var commentListReqData CommentListParam
	if err := con.Bind(&commentListReqData); err != nil {
		FailWithMessage("获取请求失败", con)
		//global.LOG.Error("获取请求失败")
		return err
	}
	fmt.Println(commentListReqData)
	vID, err := strconv.ParseInt(commentListReqData.VideoId, 10, 64)
	if err != nil {
		FailWithMessage("获取视频ID失败", con)
		//global.LOG.Error("获取视频ID失败")
		return err
	}
	//if err := utils.Verify(vID, utils.EmptyAppVerify); err != nil {
	//	FailWithMessage("视频ID为空", con)
	//	//global.LOG.Error("视频ID为空")
	//	return err
	//}
	commentListReq := comment.CommentListReq{
		VideoId: vID,
		Token:   commentListReqData.Token,
	}
	res, err := rpc.CommentList(context.Background(), &commentListReq)
	if res == nil {
		FailWithMessage("服务请求失败", con)
		//global.LOG.Error("服务请求失败")
		return err
	}
	if err != nil {
		FailWithMessage(res.BaseResp.StatusMessage, con)
		global.LOG.Error(res.BaseResp.StatusMessage)
		return err
	}
	if err := con.JSON(http.StatusOK, res); err != nil {
		FailWithMessage("响应失败", con)
		//global.LOG.Error("响应失败")
		return err
	}
	return nil
}
