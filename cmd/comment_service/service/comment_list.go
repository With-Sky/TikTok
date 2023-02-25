package service

import (
	"context"
	"tiktok/cmd/comment_service/dal/db"
	"tiktok/kitex_gen/comment"
	"tiktok/pkg/utils"
)

func CommentList(ctx context.Context, thisUserID int64, videoID int64) (commentList []*comment.Comment, err error) {
	//通过视频ID得到评论列表
	commentDBList, err := db.GetCommentList(ctx, videoID)
	if err != nil {
		commentList = nil
		return
	}
	//没有评论时直接返回
	commentLen := len(commentDBList)
	if commentLen == 0 {
		return
	}
	//通过评论列表得到发布评论的用户ID列表
	userIDs := make([]int64, commentLen)
	for i := 0; i < commentLen; i++ {
		userIDs[i] = commentDBList[i].UserID
	}
	//通过ID列表得到发布评论的用户列表
	userList, err := db.GetUserList(ctx, userIDs)
	if err != nil {
		commentList = nil
		return
	}
	userMap := make(map[int64]db.User)
	for _, j := range userList {
		userMap[int64(j.ID)] = j
	}
	//查询本人的关系列表
	myFollowDBList, err := db.GetRelationList(ctx, thisUserID)
	if err != nil {
		commentList = nil
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

	commentList = make([]*comment.Comment, commentLen)
	for i := 0; i < commentLen; i++ {
		commentList[i] = &comment.Comment{
			Id: int64(commentDBList[i].ID),
			User: &comment.User{
				Id:            userIDs[i],
				Name:          userMap[userIDs[i]].Username,
				FollowCount:   userMap[userIDs[i]].FollowCount,
				FollowerCount: userMap[userIDs[i]].FollowerCount,
				IsFollow:      myFollowMap[userIDs[i]],
			},
			Content:    commentDBList[i].CommentText,
			CreateDate: utils.TimeToFormatData(commentDBList[i].CreatedAt),
		}
	}
	return
}
