package service

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"tiktok/cmd/comment_service/dal/db"
	"tiktok/kitex_gen/comment"
)

func CommentAction(ctx context.Context, userID int64, videoID int64, commentText string) (*comment.User, int64, error) {
	userRes := new(comment.User)
	commentModel := &db.Comment{
		Model:       gorm.Model{},
		VideoID:     videoID,
		CommentText: commentText,
		UserID:      userID,
	}
	fmt.Println("视频ID", videoID)
	cID, err := db.CreateComment(ctx, commentModel)
	if err != nil {
		return userRes, 0, err
	}
	userDB, err := db.GetUser(ctx, userID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return userRes, 0, err
	}
	userRes = &comment.User{
		Id:            userID,
		Name:          userDB.Username,
		FollowCount:   userDB.FollowCount,
		FollowerCount: userDB.FollowerCount,
		IsFollow:      false,
	}
	return userRes, cID, nil
}
