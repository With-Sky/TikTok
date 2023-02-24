package service

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"tiktok/cmd/user_service/dal/db"
	"tiktok/cmd/user_service/global"
	"tiktok/pkg/utils"
)

func UserLogin(ctx context.Context, userName string, password string) (token string, userID int64, err error) {
	global.LOG.Info("用户登录服务")
	password, err = utils.Encrypt(password, []byte(global.Config.Viper.GetString("Password.key")))
	if err != nil {
		global.LOG.Error("密码错误")
		err = errors.New("用户名或密码错误，请检查您输入的用户名或密码")
		return "", 0, err
	}
	//密码加密，未完成
	userModel := &db.User{
		Model:    gorm.Model{},
		Username: userName,
	}
	fmt.Println(password)
	//查找该用户是否存在
	user, err := db.UserNameQuery(ctx, userModel)

	//没找到说明没注册过
	if err == gorm.ErrRecordNotFound {
		global.LOG.Error("未注册")
		err = errors.New("用户名或密码错误，请检查您输入的用户名或密码")
		return "", 0, err
	}

	//不属于未注册的其它错误直接返回
	if err != nil {
		global.LOG.Error("登录错误")
		return "", 0, err
	}
	//找到了用户且没有错误就开始验证密码
	if user.Password != password {
		global.LOG.Error("用户名或密码错误")
		err = errors.New("用户名或密码错误，请检查您输入的用户名或密码")
	}
	userID = int64(user.ID)
	j := utils.NewJWT(global.Config)
	token, err = j.GenToken(userID, global.Config)
	global.LOG.Info(token)
	if err != nil {
		return "", 0, err
	}
	return token, int64(user.ID), nil
}
