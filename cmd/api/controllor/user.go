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
	"tiktok/kitex_gen/user"
	"tiktok/pkg/utils"
)

type OnlineUser struct {
	Id            int64  `json:"id" `           // 用户id
	LoginTime     int64  `json:"loginTime"`     // 登录时间
	LoginLocation string `json:"loginLocation"` // 归属地
	Ip            string `json:"ip"`            // ip地址
	Token         string `json:"key"`           // token
}

// Register 注册操作
func Register(con echo.Context) error {
	// 从上下文获取请求
	var registerRequestData RegisterParam
	registerRequestData.UserName = con.QueryParam("username")
	registerRequestData.PassWord = con.QueryParam("password")
	global.LOG.Info("请求注册")

	if registerRequestData.UserName == "" {
		global.LOG.Error("用户名为空")
		FailWithMessage("用户名为空", con)
	}
	if registerRequestData.PassWord == "" {
		global.LOG.Error("密码为空")
		FailWithMessage("密码为空", con)
	}
	if len(registerRequestData.UserName) < 6 && len(registerRequestData.UserName) > 32 {
		FailWithMessage("用户名不能超过32个字符,且不能小于6", con)
		return errors.New("用户名不能超过32个字符，且不能小于6")
	}
	if len(registerRequestData.PassWord) < 6 && len(registerRequestData.PassWord) > 32 {
		FailWithMessage("密码不能超过32个字符,且不能小于6", con)
		return errors.New("密码能超过32个字符,且不能小于6")
	}
	// 注册服务的请求参数
	registerReq := user.RegisterReq{
		Username: registerRequestData.UserName,
		Password: registerRequestData.PassWord,
	}

	// 获得服务的返回
	res, err := rpc.Register(context.Background(), &registerReq)
	fmt.Println(res)
	if res.BaseResp.StatusCode != 0 {
		FailWithMessage(res.BaseResp.StatusMessage, con)
		global.LOG.Error("服务请求失败")
		return err
	}
	if err != nil {
		FailWithMessage(res.BaseResp.StatusMessage, con)
		global.LOG.Error(err.Error())
		return err
	}
	// 返回响应
	if err := con.JSON(http.StatusOK, res); err != nil {
		global.LOG.Error(err.Error())
		FailWithMessage("响应失败", con)
		return err
	}

	return nil
}

// Login 登录操作
func Login(con echo.Context) error {
	// 从上下文获取请求
	var LoginReqData LoginParam
	LoginReqData.UserName = con.QueryParam("username")
	LoginReqData.PassWord = con.QueryParam("password")
	if LoginReqData.UserName == "" {
		global.LOG.Error("用户名为空")
		FailWithMessage("用户名为空", con)
	}
	if LoginReqData.PassWord == "" {
		global.LOG.Error("密码为空")
		FailWithMessage("密码为空", con)
	}
	if len(LoginReqData.UserName) < 6 && len(LoginReqData.UserName) > 32 {
		FailWithMessage("用户名不能超过32个字符,且不能小于6", con)
		return errors.New("用户名不能超过32个字符，且不能小于6")
	}
	if len(LoginReqData.PassWord) < 6 && len(LoginReqData.PassWord) > 32 {
		FailWithMessage("密码不能超过32个字符,且不能小于6", con)
		return errors.New("密码能超过32个字符,且不能小于6")
	}
	// 登陆服务的请求参数
	LoginReq := &user.LoginReq{
		Username: LoginReqData.UserName,
		Password: LoginReqData.PassWord,
	}
	fmt.Println(LoginReqData.UserName)
	res, err := rpc.Login(context.Background(), LoginReq)

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
	//r:=utils.Redis(global.Config,global.LOG)
	var onlineUser = OnlineUser{
		Id:            res.UserId,
		LoginTime:     utils.NowUnix(),
		LoginLocation: utils.GetLocation(con.RealIP()),
		Ip:            con.RealIP(),
		Token:         res.Token,
	}
	fmt.Println(onlineUser)
	//s:=strconv.FormatInt(res.UserId,10)
	//r.Set(context.Background(),s,onlineUser,time.Duration(global.Config.Viper.GetInt("JWT.ExpiresTime"))*time.Second*19)
	if err := con.JSON(http.StatusOK, res); err != nil {
		FailWithMessage("响应失败", con)
		return err
	}
	return nil
}

// UserInfo 用户信息
func UserInfo(con echo.Context) error {
	// 从上下文获取请求
	userInfoReqData := new(UserInfoParam)
	if err := con.Bind(userInfoReqData); err != nil {
		FailWithMessage("获取请求失败", con)
		return err
	}
	if userInfoReqData.UserId == "" {
		FailWithMessage("用户ID为空", con)
		global.LOG.Error("用户ID为空")
	}
	fmt.Println(userInfoReqData.UserId)
	fmt.Println(userInfoReqData.Token)
	uID, err := strconv.ParseInt(userInfoReqData.UserId, 10, 64)
	if err != nil {
		FailWithMessage("获取用户ID失败", con)
		return err
	}
	// 请求服务的参数
	userInfoReq := user.UserInfoReq{
		UserId: uID,
		Token:  userInfoReqData.Token,
	}
	res, err := rpc.UserInfo(context.Background(), &userInfoReq)
	//if res == nil {
	//	FailWithMessage("服务请求失败", con)
	//	global.LOG.Error("服务请求失败")
	//	return err
	//}
	//if err != nil {
	//	FailWithMessage(res.BaseResp.StatusMessage, con)
	//	global.LOG.Error(res.BaseResp.StatusMessage)
	//	return err
	//}
	//
	if err := con.JSON(http.StatusOK, res); err != nil {
		FailWithMessage("响应失败", con)
		global.LOG.Error("响应失败")
		return err
	}
	return nil
}
