package service

import (
	"context"
	"gorm.io/gorm"
	"tiktok/cmd/user_service/dal/db"
	"tiktok/cmd/user_service/global"
	"tiktok/kitex_gen/user"
)

func UserInfo(ctx context.Context, thisUserID int64, userID int64) (*user.User, error) {
	global.LOG.Info("用户信息服务")
	userInfo := new(user.User)
	var err error
	//查询目标用户是否存在
	userModel, err := db.GetUser(ctx, userID)
	if err != nil {
		global.LOG.Error("用户查询失败")
		return userInfo, err
	}
	//查询关系
	relationDB, err := db.QueryFocus(ctx, thisUserID, userID)
	if err != gorm.ErrRecordNotFound && err != nil {
		global.LOG.Error("用户关系查询失败")
		return userInfo, err
	}
	var isFollow bool = false
	if err == gorm.ErrRecordNotFound {
		isFollow = false
		err = nil
	} else if relationDB.IsFollow == 3 {
		isFollow = true
	} else if (relationDB.UserID == thisUserID && relationDB.IsFollow == 1) ||
		(relationDB.ToUserID == thisUserID && relationDB.IsFollow == 2) {
		isFollow = true
	}
	userInfo = &user.User{
		Id:            userID,
		Name:          userModel.Username,
		FollowCount:   userModel.FollowCount,
		FollowerCount: userModel.FollowerCount,
		IsFollow:      isFollow,
	}
	global.LOG.Info("用户信息查询成功")
	return userInfo, nil
}
