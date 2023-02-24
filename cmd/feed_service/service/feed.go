package service

import (
	"context"
	"errors"
	"fmt"
	"tiktok/cmd/feed_service/dal/db"
	"tiktok/cmd/feed_service/global"
	"tiktok/kitex_gen/feed"
	"tiktok/pkg/utils"
	"time"
)

type Video struct {
	Author        db.User `json:"author"`         // 视频作者信息
	CommentCount  int64   `json:"comment_count"`  // 视频的评论总数
	CoverUrl      string  `json:"cover_url"`      // 视频封面地址
	FavoriteCount int64   `json:"favorite_count"` // 视频的点赞总数
	ID            int64   `json:"id"`             // 视频唯一标识
	IsFavorite    bool    `json:"is_favorite"`    // true-已点赞，false-未点赞
	PlayUrl       string  `json:"play_url"`       // 视频播放地址
	Title         string  `json:"title"`          // 视频标题
}

func GetFeed(ctx context.Context, latestTime int64, token string) ([]*feed.Video, time.Time, error) {
	var resp []*feed.Video
	var nextTime time.Time
	var err error
	var videoList []db.Video
	//var userReleaseVideo []db.UserReleaseVideo

	videoList, err = db.FindVideos(ctx)
	videoIdList := make([]int64, len(videoList))
	//userIdList := make([]int64, len(videoList))

	if err != nil {
		global.LOG.Error(err.Error())
		err = errors.New("获取视频失败")
		return nil, nextTime, err
	}
	fmt.Println(len(videoList) + 10)
	nextTime = videoList[0].CreatedAt
	for i := 0; i < len(videoList); i++ {
		videoIdList = append(videoIdList, int64(videoList[i].ID))
	}
	//userReleaseVideo, err = db.FindUserByVideoId(videoIdList)
	//
	//if err != nil {
	//	global.LOG.Error(err.Error())
	//	err = errors.New("获取视频失败")
	//	return nil, nextTime, err
	//}
	//for i := 0; i < len(userReleaseVideo); i++ {
	//	userIdList = append(userIdList, userReleaseVideo[i].UserID)
	//}
	//userList := make([]feed.User, len(videoIdList))
	//var users []db.User
	//users, err = db.FindUser(userIdList)
	//if err != nil {
	//	global.LOG.Error(err.Error())
	//	err = errors.New("获取视频失败")
	//	return nil, nextTime, err
	//}
	//for i := 0; i < len(users); i++ {
	//	userList[i].Name = users[i].Username
	//	userList[i].FollowerCount = users[i].FollowerCount
	//	userList[i].Id = int64(users[i].ID)
	//	userList[i].FollowCount = users[i].FollowCount
	//}
	//IsFavorite := make([]bool, len(videoIdList))
	//if token != "" {
	//	j := utils.NewJWT(global.Config)
	//	userId, err := j.GetIdByToken(token, global.Config)
	//	if err != nil {
	//		global.LOG.Error(err.Error())
	//		err = errors.New("获取视频失败")
	//		return nil, nextTime, err
	//	}
	//	var FavoriteVideo []int64
	//	FavoriteVideo, err = db.UserLikeVideo(userId, videoIdList)
	//	fmt.Println(FavoriteVideo)
	//	for i := 0; i < len(videoIdList); i++ {
	//		if len(FavoriteVideo) != 0 && videoIdList[i] == FavoriteVideo[i] {
	//			IsFavorite[i] = true
	//		} else {
	//			IsFavorite[i] = false
	//		}
	//	}
	//	if err != nil {
	//		global.LOG.Error(err.Error())
	//		err = errors.New("获取视频失败")
	//		return nil, nextTime, err
	//	}
	//
	//	myFollowDBList, err := db.GetRelationList(ctx, userId)
	//	if err != nil {
	//		global.LOG.Error(err.Error())
	//		err = errors.New("获取视频失败")
	//		return nil, nextTime, err
	//	}
	//	if len(myFollowDBList) != 0 {
	//		myFollowMap := make(map[int64]bool)
	//		for _, i := range myFollowDBList {
	//			followStatus := i.IsFollow
	//			if i.UserID == userId && (followStatus == 1 || followStatus == 3) {
	//				myFollowMap[i.ToUserID] = true
	//			} else if i.ToUserID == userId && (followStatus == 2 || followStatus == 3) {
	//				myFollowMap[i.UserID] = true
	//			}
	//		}
	//	}
	//}
	//countUser := feed.User{}
	//fmt.Println(len(videoList))
	//for i, j := 0, 0; i < len(videoList); i++ {
	//	if j < len(userList) && userReleaseVideo[i].UserID == userList[j].Id { // 优胜略汰
	//		countUser = userList[j]
	//		j++ // 一起走
	//	}
	//	fmt.Println(countUser)
	//	videosList := feed.Video{
	//		Author:        &countUser,
	//		CommentCount:  videoList[i].CommentCount,
	//		CoverUrl:      videoList[i].CoverUrl,
	//		FavoriteCount: videoList[i].FavoriteCount,
	//		Id:            int64(videoList[i].ID),
	//		IsFavorite:    IsFavorite[i],
	//		PlayUrl:       videoList[i].PlayUrl,
	//		Title:         videoList[i].Title,
	//	}
	//	resp = append(resp, &videosList)
	//}
	//通过视频ID列表得到视频列表
	videoDBList, err := db.GetVideoList(ctx, videoIdList)
	if err != nil {
		videoList = nil
		return nil, time.Now(), err
	}
	videoLen := len(videoDBList)
	//通过视频列表得到发布视频的用户ID列表
	userIDs := make([]int64, videoLen)
	for i := 0; i < videoLen; i++ {
		userIDs[i] = videoDBList[i].UserID
	}
	//通过发布视频的用户ID列表得到用户列表
	userDBList, err := db.GetUserList(ctx, userIDs)
	if err != nil {
		videoList = nil
		return nil, time.Now(), err
	}
	userMap := make(map[int64]db.User)
	for _, j := range userDBList {
		userMap[int64(j.ID)] = j
	}
	for i := 0; i < videoLen; i++ {
		userIDs[i] = videoDBList[i].UserID
	}
	global.LOG.Info("视频长度")
	fmt.Println(len(videoDBList))
	fmt.Println(len(userDBList))
	myFollowMap := make(map[int64]bool)
	favoriteMap := make(map[int64]bool)
	if token != "" {
		j := new(utils.JWT)
		thisUserID, err := j.GetIdByToken(token, global.Config)
		if err != nil {
			return nil, time.Now(), err
		}
		//查询本人的关系列表
		myFollowDBList, err := db.GetRelationList(ctx, thisUserID)
		if err != nil {
			videoList = nil
			return nil, time.Now(), err
		}
		//将本人关注的用户ID放入集合
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
			return nil, time.Now(), err
		}
		//将喜欢的视频ID放入集合

		for _, i := range myFavoriteList {
			favoriteMap[i.VideoID] = true
		}
	}
	resp = make([]*feed.Video, videoLen)
	for i := 0; i < videoLen; i++ {
		resp[i] = &feed.Video{
			Id: videoIdList[i],
			Author: &feed.User{
				Id:            userIDs[i],
				Name:          userMap[userIDs[i]].Username,
				FollowCount:   userMap[userIDs[i]].FollowCount,
				FollowerCount: userMap[userIDs[i]].FollowerCount,
				IsFollow:      myFollowMap[userIDs[i]],
			},
			PlayUrl:       videoDBList[i].PlayUrl,
			CoverUrl:      videoDBList[i].CoverUrl,
			FavoriteCount: videoDBList[i].FavoriteCount,
			CommentCount:  videoDBList[i].CommentCount,
			IsFavorite:    favoriteMap[videoIdList[i]],
			Title:         videoDBList[i].Title,
		}
	}
	return resp, nextTime, err
}
