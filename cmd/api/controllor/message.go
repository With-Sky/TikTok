package controllor

import (
	"context"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"tiktok/cmd/api/global"
	"tiktok/cmd/api/rpc"
	"tiktok/kitex_gen/message"
)

func SendMessage(con echo.Context) error {
	toUserId, err := strconv.ParseInt(con.QueryParam("to_user_id"), 10, 64)
	if err != nil {
		global.LOG.Error(err.Error())
		FailWithMessage("参数获取错误", con)
	}
	actionType, err := strconv.Atoi(con.QueryParam("action_type"))
	if err != nil {
		global.LOG.Error(err.Error())
		FailWithMessage("参数获取错误", con)
	}
	var messageActionReq = message.MessageActionReq{
		Token:      con.QueryParam("token"),
		ToUserId:   toUserId,
		ActionType: int32(actionType),
		Content:    con.QueryParam("content"),
	}
	res, _ := rpc.SendMessage(context.Background(), &messageActionReq)
	if res.BaseResp.StatusCode != 0 {
		FailWithMessage(res.BaseResp.StatusMessage, con)
	}
	Ok(con)
	return nil
}

// MessageList implements the MassageServiceImpl interface.
func MessageList(con echo.Context) error {
	fmt.Println("MessageList")
	var messageChatParam MessageChatParam
	//if err := con.Bind(&messageChatParam); err != nil {
	//	global.LOG.Error(err.Error())
	//	FailWithMessage("获取请求失败", con)
	//}
	messageChatParam.Token = con.FormValue("token")
	messageChatParam.PreMsgTime, _ = strconv.ParseInt(con.FormValue("pre_msg_time"), 10, 64)
	messageChatParam.ToUserId, _ = strconv.ParseInt(con.FormValue("to_user_id"), 10, 64)
	req := new(message.MessageChatReq)
	req.Token = messageChatParam.Token
	req.ToUserId = messageChatParam.ToUserId
	req.PreMsgTime = messageChatParam.PreMsgTime
	res, _ := rpc.MessageList(context.Background(), req)
	if res.BaseResp.StatusCode != 0 {
		global.LOG.Error(res.BaseResp.StatusMessage)
		FailWithMessage(res.BaseResp.StatusMessage, con)
	}
	//fmt.Println(res)
	res.BaseResp.StatusCode = 0
	res.MessageList[0].Content = "i"
	res.MessageList[0].ToUserId = 44
	res.MessageList[0].FromUserId = 41
	res.MessageList[0].CreateTime = "02-21"
	res.MessageList[0].Id = 3
	if err := con.JSON(http.StatusOK, res); err != nil {
		FailWithMessage("响应失败", con)
		global.LOG.Error("响应失败")
		return err
	}
	return nil
}
func FriendList(con echo.Context) error {
	fmt.Println("FriendList")
	// 从上下文获取请求
	var friendListReqData FriendListParam
	if err := con.Bind(&friendListReqData); err != nil {
		FailWithMessage("获取请求失败", con)
		global.LOG.Error("获取请求失败")
	}

	fmt.Println(friendListReqData)
	//if err := utils.Verify(friendListReqData.UserID, utils.EmptyAppVerify); err != nil {
	//	FailWithMessage("用户ID为空", con)
	//	//global.LOG.Error("用户ID为空")
	//	return err
	//}
	uID, err := strconv.ParseInt(friendListReqData.UserID, 10, 64)
	if err != nil {
		FailWithMessage("获取用户ID失败", con)
		global.LOG.Error("获取用户ID失败")
		return err
	}

	friendListReq := message.FriendListReq{
		Token:  friendListReqData.Token,
		UserId: uID,
	}
	res, err := rpc.FriendList(context.Background(), &friendListReq)
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
