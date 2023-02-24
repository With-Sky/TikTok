package controllor

import (
	"context"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"tiktok/cmd/api/global"
	"tiktok/cmd/api/rpc"
	"tiktok/kitex_gen/feed"
)

// FeedAction 视频流
func FeedAction(con echo.Context) error {
	// 从上下文获取请求
	fmt.Println("Feed请求")
	var feedRequestData FeedParam

	fmt.Println(feedRequestData)
	if err := con.Bind(&feedRequestData); err != nil {
		FailWithMessage("获取请求失败", con)
		global.LOG.Error("获取请求失败")
		return err
	}

	lastTime, err := strconv.ParseInt(feedRequestData.LatestTime, 10, 64)
	if err != nil {
		FailWithMessage("获取latest_time失败", con)
		global.LOG.Error("获取latest_time失败")
		return err
	}

	//视频流服务的请求参数
	feedReq := feed.FeedReq{
		LatestTime: lastTime,
		Token:      feedRequestData.Token,
	}
	fmt.Println(feedReq)
	// 获得服务的返回
	res, err := rpc.Feed(context.Background(), &feedReq)
	if res == nil {
		FailWithMessage("服务请求失败", con)
		global.LOG.Error("服务请求失败")
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
		global.LOG.Error("响应失败")
		return err
	}
	return nil
}
