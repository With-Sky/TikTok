package service

import (
	"context"
	"errors"
	"fmt"
	"tiktok/cmd/message_service/dal/db"
	"tiktok/cmd/message_service/global"
	"tiktok/kitex_gen/message"
)

func FriendList(ctx context.Context, userId int64) ([]*message.FriendUser, error) {

	friend, err := db.GetFriendList(ctx, userId)
	var friendUser = make([]*message.FriendUser, len(friend))
	if err != nil {
		global.LOG.Error(err.Error())
		err := errors.New("获取朋友列表错误")
		return friendUser, err
	}
	var friendId = make([]int64, len(friend))
	for i := 0; i < len(friend); i++ {
		if friend[i].UserID != userId {
			friendId[i] = friend[i].UserID
		} else {
			friendId[i] = friend[i].ToUserID
		}
	}
	user, err := db.GetUserById(ctx, friendId, userId)
	fmt.Println(user)
	for i := 0; i < len(user); i++ {
		fmt.Println(i)
		friends := message.FriendUser{
			Id:            user[i].Id,
			Name:          user[i].Username,
			FollowCount:   user[i].FollowCount,
			FollowerCount: user[i].FollowerCount,
			IsFollow:      true,
			Message:       user[i].Content,
			MsgType:       int64(user[i].MsgType),
		}
		//friendUser[i].MsgType = int64(user[i].MsgType)
		//friendUser[i].Message = user[i].Content
		//friendUser[i].Id = user[i].Id
		//friendUser[i].FollowCount = user[i].FollowCount
		//friendUser[i].FollowerCount = user[i].FollowerCount
		//friendUser[i].IsFollow = true
		//friendUser[i].Name = user[i].Username
		friendUser = append(friendUser, &friends)
	}
	if err != nil {
		global.LOG.Error(err.Error())
		err := errors.New("获取朋友列表错误")
		return friendUser, err
	}
	fmt.Println("++++++++++++++++")
	return friendUser, nil
}
