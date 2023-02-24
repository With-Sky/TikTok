package controllor

import (
	"context"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"tiktok/cmd/api/global"
	"tiktok/cmd/api/rpc"
	"tiktok/kitex_gen/publish"
)

// PublishAction 投稿操作
func PublishAction(con echo.Context) error {
	//从上下文获取请求
	header, err := con.FormFile("data")
	if err != nil {
		FailWithMessage("获取文件失败", con)
		global.LOG.Error("获取文件失败")
		return err
	}
	var fileByte = make([]byte, header.Size)
	file, err := header.Open()
	file.Read(fileByte)

	// 投稿服务的请求参数
	publishReq := publish.PublishReq{
		Token: con.FormValue("token"),
		Title: con.FormValue("title"),
		Data:  fileByte,
	}
	//fmt.Println(publishReq)
	// 获得服务的返回
	res, err := rpc.Publish(context.Background(), &publishReq)
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
	return nil
}

// PublishList 发布列表
func PublishList(con echo.Context) error {
	// 从上下文获取请求
	var publishListReqData PublishListParam
	publishListReqData.Token = con.FormValue("token")
	userID := con.QueryParam("user_id")
	//j:=new(utils.JWT)
	//userId,err:=j.GetIdByToken(publishListReqData.Token,global.Config)
	uID, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		FailWithMessage("获取对方用户ID失败", con)
		global.LOG.Error("获取对方用户ID失败")
		return err
	}
	fmt.Println(publishListReqData)
	publishListReq := publish.PublishListReq{
		UserId: uID,
		Token:  publishListReqData.Token,
	}
	res, err := rpc.PublishList(context.Background(), &publishListReq)
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
