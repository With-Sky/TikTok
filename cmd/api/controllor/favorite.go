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
	"tiktok/kitex_gen/favorite"
)

// FavoriteAction 赞操作
func FavoriteAction(con echo.Context) error {
	// 从上下文获取请求
	var favoriteRequestData FavoriteParam
	favoriteRequestData.ActionType = con.QueryParam("action_type")
	favoriteRequestData.Token = con.QueryParam("token")
	favoriteRequestData.VideoId = con.QueryParam("video_id")
	fmt.Println(favoriteRequestData)
	//if err := con.Bind(&favoriteRequestData); err != nil {
	//	FailWithMessage("获取请求失败", con)
	//	//global.LOG.Error("获取请求失败")
	//	return err
	//}
	vID, err := strconv.ParseInt(favoriteRequestData.VideoId, 10, 64)
	if err != nil {
		FailWithMessage("获取视频ID失败", con)
		global.LOG.Error("获取视频ID失败")
		return err
	}

	// 判断是点赞（1），还是取消点赞（2）
	if favoriteRequestData.ActionType == "1" {
		fmt.Println("点赞")
		// 点赞服务的请求参数
		favoriteReq := favorite.FavoriteReq{
			Token:   favoriteRequestData.Token,
			VideoId: vID,
		}

		// 获得服务的返回
		res, err := rpc.Favorite(context.Background(), &favoriteReq)
		if res == nil {
			FailWithMessage("服务请求失败", con)
			global.LOG.Error("服务请求失败")
			return err
		}
		if err != nil {
			FailWithMessage(res.BaseResp.StatusMessage, con)
			global.LOG.Error(res.BaseResp.StatusMessage)
			return err
		}
		// 返回响应
		if err := con.JSON(http.StatusOK, res); err != nil {
			FailWithMessage("响应失败", con)
			global.LOG.Error("响应失败")
			return err
		}
	} else if favoriteRequestData.ActionType == "2" {
		fmt.Println("取消点赞")

		// 取消点赞服务的请求参数
		deletefavoriteReq := favorite.DeleteFavoriteReq{
			Token:   favoriteRequestData.Token,
			VideoId: vID,
		}

		// 获得服务的返回
		res, err := rpc.CancelFavorite(context.Background(), &deletefavoriteReq)
		if res == nil {
			FailWithMessage("服务请求失败", con)
			global.LOG.Error("服务请求失败")
			return err
		}
		if err != nil {
			FailWithMessage(res.BaseResp.StatusMessage, con)
			global.LOG.Error(res.BaseResp.StatusMessage)
			return err
		}
		// 返回响应
		if err := con.JSON(http.StatusOK, res); err != nil {
			FailWithMessage("响应失败", con)
			global.LOG.Error("响应失败")
			return err
		}
	} else {
		FailWithMessage("操作必须为点赞或取消", con)
		global.LOG.Error("操作必须为点赞或取消")
		return errors.New("Action must be 1 or 2\n")
	}
	return nil
}

// FavoriteList 喜欢列表
func FavoriteList(con echo.Context) error {
	// 从上下文获取请求
	var favoriteListReqData FavoriteListParam
	if err := con.Bind(&favoriteListReqData); err != nil {
		return err
	}
	uID, err := strconv.ParseInt(favoriteListReqData.UserID, 10, 64)
	if err != nil {
		FailWithMessage("获取用户ID失败", con)
		global.LOG.Error("获取用户ID失败")
		return err
	}
	// 喜欢列表请求服务的参数
	favoriteListReq := favorite.FavoriteListReq{
		UserId: uID,
		Token:  favoriteListReqData.Token,
	}
	res, err := rpc.FavoriteList(context.Background(), &favoriteListReq)
	if res == nil {
		FailWithMessage("服务请求失败", con)
		global.LOG.Error("服务请求失败")
		return err
	}
	if err != nil {
		FailWithMessage(res.BaseResp.StatusMessage, con)
		global.LOG.Error(res.BaseResp.StatusMessage)
		return err
	}

	if err := con.JSON(http.StatusOK, res); err != nil {
		FailWithMessage("响应失败", con)
		global.LOG.Error("响应失败")
		return err
	}
	return nil
}
