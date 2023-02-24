package main

import (
	"context"
	"errors"
	"fmt"
	"tiktok/cmd/user_service/global"
	"tiktok/cmd/user_service/pack"
	"tiktok/cmd/user_service/service"
	user "tiktok/kitex_gen/user"
	"tiktok/pkg/utils"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginReq) (*user.LoginRes, error) {
	global.LOG.Info("用户登录")
	resp := new(user.LoginRes)
	if req.Username == "" {
		err := errors.New("用户名为空")
		global.LOG.Error("用户名为空")
		resp.BaseResp = pack.BuildBaseResp(err, 400)
		return resp, nil
	}
	if req.Password == "" {
		err := errors.New("密码为空")
		global.LOG.Error("密码为空")
		resp.BaseResp = pack.BuildBaseResp(err, 400)
		return resp, nil
	}
	if len(req.Username) < 6 && len(req.Username) > 32 {
		err := errors.New("用户名不能超过32个字符,且不能小于6")
		global.LOG.Error("用户名不能超过32个字符,且不能小于6")
		resp.BaseResp = pack.BuildBaseResp(err, 400)
		return resp, nil
	}
	if len(req.Password) < 6 && len(req.Password) > 32 {
		err := errors.New("密码不能超过32个字符,且不能小于6")
		global.LOG.Error("密码不能超过32个字符,且不能小于6")
		resp.BaseResp = pack.BuildBaseResp(err, 400)
		return resp, nil
	}
	token, userID, err := service.UserLogin(ctx, req.Username, req.Password)
	if err == nil {
		global.LOG.Info("用户登录服务成功")

		resp.Token = token
		resp.UserId = userID
		resp.BaseResp = pack.BuildBaseResp(err, 0)
	} else {
		global.LOG.Error("用户登录服务错误")
		resp.BaseResp = pack.BuildBaseResp(err, 400)
	}
	return resp, nil
}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterReq) (*user.RegisterRes, error) {
	global.LOG.Info("用户注册")
	resp := new(user.RegisterRes)
	if req.Username == "" {
		err := errors.New("用户名为空")
		global.LOG.Error("用户名为空")
		resp.BaseResp = pack.BuildBaseResp(err, 400)
		return resp, nil
	}
	if req.Password == "" {
		err := errors.New("密码为空")
		global.LOG.Error("密码为空")
		resp.BaseResp = pack.BuildBaseResp(err, 400)
		return resp, nil
	}
	if len(req.Username) < 6 && len(req.Username) > 32 {
		err := errors.New("用户名不能超过32个字符,且不能小于6")
		global.LOG.Error("用户名不能超过32个字符,且不能小于6")
		resp.BaseResp = pack.BuildBaseResp(err, 400)
		return resp, nil
	}
	if len(req.Password) < 6 && len(req.Password) > 32 {
		err := errors.New("密码不能超过32个字符,且不能小于6")
		global.LOG.Error("密码不能超过32个字符,且不能小于6")
		resp.BaseResp = pack.BuildBaseResp(err, 400)
		return resp, nil
	}
	token, userID, err := service.UserRegister(ctx, req.Username, req.Password)
	if err == nil {
		global.LOG.Info("用户注册服务成功")
		resp.Token = token
		resp.UserId = userID
		resp.BaseResp = pack.BuildBaseResp(err, 0)
	} else {
		global.LOG.Error("用户注册服务错误")
		resp.BaseResp = pack.BuildBaseResp(err, 400)
	}
	return resp, nil
}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(ctx context.Context, req *user.UserInfoReq) (*user.UserInfoRes, error) {
	global.LOG.Info("用户信息")
	resp := new(user.UserInfoRes)
	fmt.Println(req.UserId)
	if req.UserId == 0 {
		err := errors.New("参数错误")
		global.LOG.Error("用户ID为空")
		resp.BaseResp = pack.BuildBaseResp(err, 0)
		return resp, nil
	}
	j := utils.NewJWT(global.Config)
	c, err := j.ParseToken(req.Token, global.Config)
	global.Config.Viper.GetInt("JWT.ExpiresTime")
	if err != nil {
		err := errors.New("token解析失败")
		resp.BaseResp = pack.BuildBaseResp(err, 400)
		return resp, err
	}
	userInfo, err := service.UserInfo(ctx, c.UserID, req.UserId)
	if err == nil {
		global.LOG.Info("用户信息服务成功")
		resp.User = userInfo
		resp.BaseResp = pack.BuildBaseResp(err, 0)
		return resp, nil
	} else {
		global.LOG.Error("用户信息服务错误")
		resp.BaseResp = pack.BuildBaseResp(err, 400)
		return resp, nil
	}
	return resp, nil
}
