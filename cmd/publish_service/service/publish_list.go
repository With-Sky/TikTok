package service

import (
	"context"
	"gorm.io/gorm"
	"tiktok/cmd/publish_service/dal/db"
	"tiktok/kitex_gen/publish"
)

func PublishList(ctx context.Context, thisUserID int64, userID int64) (videoList []*publish.Video, err error) {
	//通过用户ID得到发布的视频列表
	videoDBList, err := db.GetVideoList(ctx, userID)
	if err != nil {
		videoList = nil
		return
	}

	//没有发布的视频时直接返回
	publishLen := len(videoDBList)
	if publishLen == 0 {
		return
	}
	//通过发布视频的用户ID得到用户
	userDB, err := db.GetUser(ctx, userID)
	if err != nil {
		videoList = nil
		return
	}

	//查询本人的关系列表
	myFollow, err := db.QueryFocus(ctx, thisUserID, userID)
	if err != nil && err != gorm.ErrRecordNotFound {
		videoList = nil
		return
	}
	var is_follow bool
	if myFollow.UserID == thisUserID && (myFollow.IsFollow == 1 || myFollow.IsFollow == 3) {
		is_follow = true
	} else if myFollow.ToUserID == thisUserID && (myFollow.IsFollow == 2 || myFollow.IsFollow == 3) {
		is_follow = true
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

	videoList = make([]*publish.Video, publishLen)
	for i := 0; i < publishLen; i++ {
		videoList[i] = &publish.Video{
			Id: int64(videoDBList[i].ID),
			Author: &publish.User{
				Id:            int64(userDB.ID),
				Name:          userDB.Username,
				FollowCount:   userDB.FollowCount,
				FollowerCount: userDB.FollowerCount,
				IsFollow:      is_follow,
			},
			PlayUrl:       videoDBList[i].PlayUrl,
			CoverUrl:      videoDBList[i].CoverUrl,
			FavoriteCount: videoDBList[i].FavoriteCount,
			CommentCount:  videoDBList[i].CommentCount,
			IsFavorite:    favoriteMap[int64(videoDBList[i].ID)],
			Title:         videoDBList[i].Title,
		}
	}
	return
}
