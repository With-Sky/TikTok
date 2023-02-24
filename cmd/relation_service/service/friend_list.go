package service

import (
	"context"
	"tiktok/cmd/relation_service/dal/db"
	"tiktok/kitex_gen/relation"
)

func FriendList(ctx context.Context, thisUserID int64, UserID int64) (friendList []*relation.User, err error) {
	//通过用户ID得到其关系的列表
	relationDBList, err := db.GetRelationList(ctx, UserID)
	if err != nil {
		relationDBList = nil
		return
	}
	friendLen := len(relationDBList)
	//没有任何关系直接返回
	if friendLen == 0 {
		return
	}
	//得到互关对象的ID列表
	var followIDs []int64
	for i := 0; i < friendLen; i++ {
		followStatus := relationDBList[i].IsFollow
		if relationDBList[i].UserID == UserID && followStatus == 3 {
			followIDs = append(followIDs, relationDBList[i].ToUserID)
		} else if relationDBList[i].ToUserID == UserID && followStatus == 3 {
			followIDs = append(followIDs, relationDBList[i].UserID)
		}
	}
	friendLen = len(followIDs)

	//通过用户ID列表得到用户信息
	userDBList, err := db.GetUserList(ctx, followIDs)
	if err != nil {
		friendList = nil
		return
	}

	//查询本人的关系列表
	myFollowDBList, err := db.GetRelationList(ctx, thisUserID)
	if err != nil {
		friendList = nil
		return
	}
	//将本人关注的用户ID放入集合
	myFollowMap := make(map[int64]bool)
	for _, i := range myFollowDBList {
		followStatus := i.IsFollow
		if i.UserID == thisUserID && (followStatus == 1 || followStatus == 3) {
			myFollowMap[i.ToUserID] = true
		} else if i.ToUserID == thisUserID && (followStatus == 2 || followStatus == 3) {
			myFollowMap[i.UserID] = true
		}
	}

	friendList = make([]*relation.User, friendLen)
	for i := 0; i < friendLen; i++ {
		friendList[i] = &relation.User{
			Id:            int64(userDBList[i].ID),
			Name:          userDBList[i].Username,
			FollowCount:   userDBList[i].FollowCount,
			FollowerCount: userDBList[i].FollowerCount,
			IsFollow:      myFollowMap[int64(userDBList[i].ID)],
		}
	}
	return
}
