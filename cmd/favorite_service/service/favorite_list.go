package service

import (
	"context"
	"tiktok/cmd/favorite_service/dal/db"
	"tiktok/kitex_gen/favorite"
)

func FavoriteList(ctx context.Context, thisUserID int64, userID int64) (videoList []*favorite.Video, err error) {
	//通过用户ID得到喜欢的视频列表
	favoriteDBList, err := db.GetFavoriteList(ctx, userID)
	if err != nil {
		videoList = nil
		return
	}
	//没有点赞的视频时直接返回
	favoriteLen := len(favoriteDBList)
	if favoriteLen == 0 {
		return
	}
	//通过点赞列表得到被点赞的视频ID列表
	videoIDs := make([]int64, favoriteLen)
	for i := 0; i < favoriteLen; i++ {
		videoIDs[i] = favoriteDBList[i].VideoID
	}
	//通过视频ID列表得到视频列表
	videoDBList, err := db.GetVideoList(ctx, videoIDs)
	if err != nil {
		videoList = nil
		return
	}
	//通过视频列表得到发布视频的用户ID列表
	userIDs := make([]int64, favoriteLen)
	for i := 0; i < favoriteLen; i++ {
		userIDs[i] = videoDBList[i].UserID
	}
	//通过发布视频的用户ID列表得到用户列表
	userDBList, err := db.GetUserList(ctx, userIDs)
	if err != nil {
		videoList = nil
		return
	}

	//查询本人的关系列表
	myFollowDBList, err := db.GetRelationList(ctx, thisUserID)
	if err != nil {
		videoList = nil
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

	//通过本人ID得到本人点赞的视频ID列表
	myFavoriteList, err := db.GetFavoriteList(ctx, thisUserID)
	if err != nil {
		videoList = nil
		return
	}
	//将喜欢的视频ID放入集合
	favoriteMap := make(map[int64]bool)
	for _, i := range myFavoriteList {
		favoriteMap[i.VideoID] = true
	}

	videoList = make([]*favorite.Video, favoriteLen)
	for i := 0; i < favoriteLen; i++ {
		videoList[i] = &favorite.Video{
			Id: videoIDs[i],
			Author: &favorite.User{
				Id:            int64(userDBList[i].ID),
				Name:          userDBList[i].Username,
				FollowCount:   userDBList[i].FollowCount,
				FollowerCount: userDBList[i].FollowerCount,
				IsFollow:      myFollowMap[int64(userDBList[i].ID)],
			},
			PlayUrl:       videoDBList[i].PlayUrl,
			CoverUrl:      videoDBList[i].CoverUrl,
			FavoriteCount: videoDBList[i].FavoriteCount,
			CommentCount:  videoDBList[i].CommentCount,
			IsFavorite:    favoriteMap[videoIDs[i]],
			Title:         videoDBList[i].Title,
		}
	}
	return
}
