package controllor

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"tiktok/cmd/api/rpc"
	"tiktok/kitex_gen/relation"
)

// FollowAction 关注操作
func FollowAction(con echo.Context) error {
	// 从上下文获取请求
	var followRequestData FollowParam
	followRequestData.ActionType = con.QueryParam("action_type")
	followRequestData.Token = con.QueryParam("token")
	followRequestData.ToUserId = con.QueryParam("to_user_id")
	fmt.Println(followRequestData)
	//if err := con.Bind(&followRequestData); err != nil {
	//	FailWithMessage("获取请求失败", con)
	//	//global.LOG.Error("获取请求失败")
	//	return err
	//}
	// 获取用户ID
	toUID, err := strconv.ParseInt(followRequestData.ToUserId, 10, 64)
	if err != nil {
		FailWithMessage("获取对方用户ID失败", con)
		//global.LOG.Error("获取对方用户ID失败")
		return err
	}
	//if err := utils.Verify(toUID, utils.EmptyAppVerify); err != nil {
	//	FailWithMessage("对方用户ID为空", con)
	//	//global.LOG.Error("对方用户ID为空")
	//	return err
	//}
	// 判断是关注（1），还是取消关注（2）
	if followRequestData.ActionType == "1" {
		fmt.Println("关注")
		// 关注服务的请求参数
		relationReq := relation.FollowReq{
			Token:    followRequestData.Token,
			ToUserId: toUID,
		}
		// 获得服务的返回
		res, err := rpc.Follow(context.Background(), &relationReq)
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
	} else if followRequestData.ActionType == "2" {
		fmt.Println("取消关注")

		// 关注服务的请求参数
		cancelFallowReq := relation.CancelFollowReq{
			Token:    followRequestData.Token,
			ToUserId: toUID,
		}
		// 获得服务的返回
		res, err := rpc.CancelFollow(context.Background(), &cancelFallowReq)
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
		FailWithMessage("操作必须为关注或取消", con)
		//global.LOG.Error("操作必须为关注或取消")
		return errors.New("Action must be 1 or 2\n")
	}
	return nil
}

// FollowList 关注列表
func FollowList(con echo.Context) error {
	// 从上下文获取请求
	var followListReqData FollowListParam
	if err := con.Bind(&followListReqData); err != nil {
		FailWithMessage("获取请求失败", con)
		//global.LOG.Error("获取请求失败")
		return err
	}
	fmt.Println(followListReqData)
	//if err := utils.Verify(followListReqData.UserID, utils.EmptyAppVerify); err != nil {
	//	FailWithMessage("用户ID为空", con)
	//	//global.LOG.Error("用户ID为空")
	//	return err
	//}
	uID, err := strconv.ParseInt(followListReqData.UserID, 10, 64)
	if err != nil {
		FailWithMessage("获取用户ID失败", con)
		//global.LOG.Error("获取用户ID失败")
		return err
	}
	followListReq := relation.FollowListReq{
		UserId: uID,
		Token:  followListReqData.Token,
	}
	res, err := rpc.FollowList(context.Background(), &followListReq)
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
	if err := con.JSON(http.StatusOK, res); err != nil {
		FailWithMessage("响应失败", con)
		//global.LOG.Error("响应失败")
		return err
	}
	return nil
}

// FollowerList 粉丝列表
func FollowerList(con echo.Context) error {
	// 从上下文获取请求
	var followerListReqData FollowerListParam
	if err := con.Bind(&followerListReqData); err != nil {
		FailWithMessage("获取请求失败", con)
		//global.LOG.Error("获取请求失败")
		return err
	}
	fmt.Println(followerListReqData)
	//if err := utils.Verify(followerListReqData.UserID, utils.EmptyAppVerify); err != nil {
	//	FailWithMessage("用户ID为空", con)
	//	//global.LOG.Error("用户ID为空")
	//	return err
	//}
	uID, err := strconv.ParseInt(followerListReqData.UserID, 10, 64)
	if err != nil {
		FailWithMessage("获取用户ID失败", con)
		//global.LOG.Error("获取用户ID失败")
		return err
	}
	//粉丝列表服务的请求参数
	followerListReq := relation.FollowerListReq{
		UserId: uID,
		Token:  followerListReqData.Token,
	}
	res, err := rpc.FollowerList(context.Background(), &followerListReq)
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
	if err := con.JSON(http.StatusOK, res); err != nil {
		FailWithMessage("响应失败", con)
		//global.LOG.Error("响应失败")
		return err
	}
	return nil
}

// FriendList 好友列表
//func FriendList(con echo.Context) error {
//	// 从上下文获取请求
//	var friendListReqData FriendListParam
//	if err := con.Bind(&friendListReqData); err != nil {
//		FailWithMessage("获取请求失败", con)
//		//global.LOG.Error("获取请求失败")
//		return err
//	}
//	fmt.Println(friendListReqData)
//	//if err := utils.Verify(friendListReqData.UserID, utils.EmptyAppVerify); err != nil {
//	//	FailWithMessage("用户ID为空", con)
//	//	//global.LOG.Error("用户ID为空")
//	//	return err
//	//}
//	uID, err := strconv.ParseInt(friendListReqData.UserID, 10, 64)
//	if err != nil {
//		FailWithMessage("获取用户ID失败", con)
//		//global.LOG.Error("获取用户ID失败")
//		return err
//	}
//
//	friendListReq := relation.FriendListReq{
//		Token:  friendListReqData.Token,
//		UserId: uID,
//	}
//	res, err := rpc.FriendList(context.Background(), &friendListReq)
//	if res == nil {
//		FailWithMessage("服务请求失败", con)
//		//global.LOG.Error("服务请求失败")
//		return err
//	}
//	if err != nil {
//		FailWithMessage(res.BaseResp.StatusMessage, con)
//		//global.LOG.Error(res.BaseResp.StatusMessage)
//		return err
//	}
//	if err := con.JSON(http.StatusOK, res); err != nil {
//		FailWithMessage("响应失败", con)
//		//global.LOG.Error("响应失败")
//		return err
//	}
//	return nil
//}
