package service

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"tiktok/cmd/relation_service/dal/db"
)

func CancelFollow(ctx context.Context, userID int64, toUserID int64) error {
	if userID == toUserID {
		return errors.New("不能对自己进行关注和取消关注操作")
	}
	//先查询关注状态
	userFocusDB, err := db.QueryFocus(ctx, userID, toUserID)
	//没有找到则返回错误
	if err == gorm.ErrRecordNotFound {
		return errors.New("尚未关注")
	}
	//不属于未查找的其他错误直接返回
	if err != nil {
		return err
	}
	//为3说明已经互关
	if userFocusDB.IsFollow == 3 {
		if userFocusDB.UserID == userID && userFocusDB.ToUserID == toUserID { //表中保存的是userID用户对toUserID用户的状态
			userFocusDB.IsFollow = 2
			return db.DecreaseFocus(ctx, userFocusDB, userID, toUserID)
		} else if userFocusDB.UserID == toUserID && userFocusDB.ToUserID == userID { //表中保存的是userID用户对toUserID用户的状态
			userFocusDB.IsFollow = 1
			return db.DecreaseFocus(ctx, userFocusDB, toUserID, userID)
		}
		return nil
	}
	if userFocusDB.UserID == userID && userFocusDB.ToUserID == toUserID { //表中保存的是userID用户对toUserID用户的状态
		//为2说明userID用户对toUserID用户尚未关注
		if userFocusDB.IsFollow == 2 {
			return errors.New("尚未关注")
		}
		//为1说明userID用户被toUserID用户关注，取关需要执行删除
		if userFocusDB.IsFollow == 1 {
			return db.DeleteFocusByID(ctx, userFocusDB, userID, toUserID)
		}
		return errors.New("关注状态异常")
	} else if userFocusDB.UserID == toUserID && userFocusDB.ToUserID == userID { //表中保存的是userID用户对toUserID用户的状态
		//为1说明userID用户对toUserID用户尚未关注
		if userFocusDB.IsFollow == 1 {
			return errors.New("尚未关注")
		}
		//为2说明userID用户被toUserID用户关注，取关需要执行删除
		if userFocusDB.IsFollow == 2 {
			return db.DeleteFocusByID(ctx, userFocusDB, toUserID, userID)
		}
		return errors.New("关注状态异常")
	}
	return err
}
