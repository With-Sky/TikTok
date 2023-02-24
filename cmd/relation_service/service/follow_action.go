package service

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"tiktok/cmd/relation_service/dal/db"
	"tiktok/cmd/relation_service/global"
)

func FollowAction(ctx context.Context, userID int64, toUserID int64) error {
	if userID == toUserID {
		return errors.New("不能对自己进行关注和取消关注操作")
	}
	//先查询关注状态
	userFocusDB, err := db.QueryFocus(ctx, userID, toUserID)
	//没有找到则创建一个关注
	if err == gorm.ErrRecordNotFound {
		followModel := &db.UserFocusOn{
			UserID:   userID,
			ToUserID: toUserID,
			IsFollow: 1,
		}
		return db.CreateFocus(ctx, followModel)
	}
	//不属于未查找的其他错误直接返回
	if err != nil {
		return err
	}
	//为3说明已经互关
	if userFocusDB.IsFollow == 3 {
		return errors.New("已关注")
	}
	if userFocusDB.UserID == userID && userFocusDB.ToUserID == toUserID { //表中保存的是userID用户对toUserID用户的状态
		//为1说明userID用户对toUserID用户已经关注
		if userFocusDB.IsFollow == 1 {
			return errors.New("已关注")
		}
		//为2说明userID用户被toUserID用户关注
		if userFocusDB.IsFollow == 2 {
			userFocusDB.IsFollow = 3
			if err = SendMessage(ctx, "我们已经成为好友了，开始聊天吧", toUserID, userID); err != nil {
				return err
			}
			return db.IncreaseFocus(ctx, userFocusDB, userID, toUserID)
		}
		return errors.New("关注失败")
	} else if userFocusDB.UserID == toUserID && userFocusDB.ToUserID == userID { //表中保存的是toUserID用户对userID用户的状态
		//为2说明userID用户对toUserID用户已经关注
		if userFocusDB.IsFollow == 2 {
			return errors.New("已关注")
		}
		//为1说明userID用户被toUserID用户关注
		if userFocusDB.IsFollow == 1 {
			userFocusDB.IsFollow = 3
			if err = SendMessage(ctx, "我们已经成为好友了，开始聊天吧", toUserID, userID); err != nil {
				return err
			}
			return db.IncreaseFocus(ctx, userFocusDB, toUserID, userID)
		}
		return errors.New("关注失败")
	}
	return err
}

func SendMessage(ctx context.Context, content string, toUserId int64, userId int64) error {
	err := db.CreateMessage(ctx, content, toUserId, userId)
	if err != nil {
		err = errors.New("发送信息失败")
		global.LOG.Error("发送信息失败")
		return err
	}
	return nil
}
