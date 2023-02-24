package service

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"tiktok/cmd/user_service/dal/db"
	"tiktok/cmd/user_service/global"
	"tiktok/pkg/utils"
)

func UserRegister(ctx context.Context, userName string, password string) (token string, userID int64, err error) {
	global.LOG.Info("用户注册服务")
	password, err = utils.Encrypt(password, []byte(global.Config.Viper.GetString("Password.key")))
	userModel := &db.User{
		Model:         gorm.Model{},
		Username:      userName,
		Password:      password,
		FollowCount:   0,
		FollowerCount: 0,
		LikeTotal:     0,
	}
	//查找该用户是否存在
	_, err = db.UserNameQuery(ctx, userModel)
	//没有查找到则创建用户
	if err == gorm.ErrRecordNotFound {
		err = db.CreateUser(ctx, userModel)
		if err != nil {
			return
		}
		userID = int64(userModel.ID)
		j := utils.NewJWT(global.Config)
		token, err = j.GenToken(userID, global.Config)
		if err != nil {
			global.LOG.Error("token生成失败")
			return "", 0, err
		}
		global.LOG.Error("注册成功")
		return token, userID, nil
	}
	//找到了用户说明已经注册过了
	if err == nil {
		err = errors.New("用户名已存在，请登录或使用其它用户名注册")
		global.LOG.Error("用户名已存在")
		return "", 0, err
	}
	//其它错误说明查找时出错了
	return "", 0, err
}
