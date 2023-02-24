package db

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

// Video 视频数据库模型
type Video struct {
	gorm.Model
	UserID        int64  `json:"user_id" gorm:"comment:发布者的ID"`         // 发布者的ID
	Title         string `json:"title" gorm:"comment:视频标题"`             // 视频标题
	PlayUrl       string `json:"play_url" gorm:"comment:视频播放地址"`        // 视频播放地址
	CoverUrl      string `json:"cover_url" gorm:"comment:视频封面地址"`       // 视频封面地址
	FavoriteCount int64  `json:"favorite_count" gorm:"comment:视频的点赞总数"` // 视频的点赞总数
	CommentCount  int64  `json:"comment_count" gorm:"comment:视频的评论总数"`  // 视频的评论总数
	Name          string `json:"name" gorm:"comment:文件名"`               // 文件名
	Tag           string `json:"tag" gorm:"comment:文件标签"`               // 文件标签
	Key           string `json:"key" gorm:"comment:编号"`
}

func (v *Video) TableName() string {
	return "videos"
}

// User 用户数据库模型
type User struct {
	gorm.Model
	Username      string `json:"username" gorm:"comment:用户登录名"`      // 用户登录名
	Password      string `json:"-"  gorm:"comment:用户登录密码"`           // 用户登录密码
	FollowCount   int64  `json:"follow_count" gorm:"comment:关注总数"`   // 关注总数
	FollowerCount int64  `json:"follower_count" gorm:"comment:粉丝总数"` // 粉丝总数
	LikeTotal     int64  `json:"like_total" gorm:"comment:获赞总数"`     // 获赞总数
}

func (u *User) TableName() string {
	return "users"
}

// UserFocusOn 用户关注数据库模型
type UserFocusOn struct {
	gorm.Model
	UserID   int64 `json:"user_id" gorm:"comment:用户id"`      // 用户ID
	ToUserID int64 `json:"to_user_id" gorm:"comment:对方用户id"` // 对方用户id
	IsFollow int32 `json:"is_follow"  gorm:"comment:双方关注状态,1表示user关注to_user,2表示to_user关注user,3表示互关"`
}
type UserReleaseVideo struct {
	gorm.Model
	UserID  int64 `json:"user_id" gorm:"comment:用户id"`  // 用户ID
	VideoID int64 `json:"video_id" gorm:"comment:视频id"` // 视频ID
}

func (fo *UserFocusOn) TableName() string {
	return "user_focus_ons"
}

// UserFavorite 用户点赞数据库模型
type UserFavorite struct {
	gorm.Model
	VideoID    int64 `json:"video_id" gorm:"comment:点赞的视频ID"`
	UserID     int64 `json:"user_id" gorm:"comment:点赞的用户ID"`
	IsFavorite bool  `json:"is_favorite" gorm:"comment:点赞状态"` // 点赞状态
}

func (fa *UserFavorite) TableName() string {
	return "user_like_videos"
}

// GetVideos 查找表中视频ID为videoID的评论
func GetVideos(ctx context.Context, videoID int64) ([]Video, error) {
	var videos []Video
	err := TIK_DB.WithContext(ctx).Model(&Video{}).Where("video_id = ?", videoID).Find(&videos).Error
	return videos, err
}

func FindVideos(ctx context.Context) ([]Video, error) {
	var videoList []Video
	if err := TIK_DB.WithContext(ctx).Model(&Video{}).
		Find(&videoList).Limit(30).Error; err != nil {
		return nil, err
	}
	fmt.Println(videoList)
	return videoList, nil
}
func FindUserByVideoId(videoIdList []int64) ([]UserReleaseVideo, error) {
	var userReleaseVideo []UserReleaseVideo
	if err := TIK_DB.Model(&UserReleaseVideo{}).
		Where("video_id in ?", videoIdList).
		Find(&userReleaseVideo).Error; err != nil {
		return nil, err
	}
	return userReleaseVideo, nil
}

func FindUser(userIdList []int64) ([]User, error) {
	var users []User
	if err := TIK_DB.Model(&User{}).
		Where("id in ?", userIdList).
		Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func UserLikeVideo(userId int64, videoIdList []int64) ([]int64, error) {
	VideoId := make([]int64, len(videoIdList))
	if err := TIK_DB.Model(&UserFavorite{}).Select("video_id").
		Where("user_id = ? and video_id in ?", userId, videoIdList).
		Find(&VideoId).Error; err != nil {
		return nil, err
	}
	return VideoId, nil
}

// GetFavoriteList 查找表中用户ID为userID的点赞
func GetFavoriteList(ctx context.Context, userID int64) ([]UserFavorite, error) {
	var favorites []UserFavorite
	err := TIK_DB.WithContext(ctx).Model(&UserFavorite{}).Where("user_id = ?", userID).Find(&favorites).Error
	return favorites, err
}

// GetVideoList 查找视频列表
func GetVideoList(ctx context.Context, videoIDs []int64) ([]Video, error) {
	var videos []Video
	err := TIK_DB.WithContext(ctx).Model(&Video{}).Find(&videos, videoIDs).Error
	return videos, err
}

// GetUserList 查找用户列表
func GetUserList(ctx context.Context, userIDs []int64) ([]User, error) {
	var users []User
	err := TIK_DB.WithContext(ctx).Model(&User{}).Find(&users, userIDs).Error
	return users, err
}

// GetRelationList 查找用户的关系列表
func GetRelationList(ctx context.Context, userID int64) ([]UserFocusOn, error) {
	var userRelations []UserFocusOn
	err := TIK_DB.WithContext(ctx).Model(&UserFocusOn{}).
		Where("user_id = ?", userID).Or("to_user_id = ?", userID).Find(&userRelations).Error
	return userRelations, err
}
