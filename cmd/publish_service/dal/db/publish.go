package db

import (
	"context"
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
type UserReleaseVideo struct {
	gorm.Model
	UserID  int64 `json:"user_id" gorm:"comment:用户id"`  // 用户ID
	VideoID int64 `json:"video_id" gorm:"comment:视频id"` // 视频ID
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

func (fo *UserFocusOn) TableName() string {
	return "user_focus_ons"
}

// UserFavorite 用户点赞数据库模型
type UserFavorite struct {
	gorm.Model
	VideoID int64 `json:"video_id" gorm:"comment:点赞的视频ID"`
	UserID  int64 `json:"user_id" gorm:"comment:点赞的用户ID"`
}

func (fa *UserFavorite) TableName() string {
	return "user_like_videos"
}

// CreateVideo 在表中创建视频
func CreateVideo(ctx context.Context, video Video) error {
	err := TIK_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Model(&Video{}).Create(&video).Error; err != nil {
			return err
		}
		if err := tx.WithContext(ctx).Create(&UserReleaseVideo{
			UserID:  video.UserID,
			VideoID: int64(video.ID),
		}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteVideo 在表中删除视频
func DeleteVideo(ctx context.Context, videoID int64) error {
	return TIK_DB.WithContext(ctx).Model(&Video{}).Delete(&Video{}, videoID).Error
}

// GetVideoList 查找表中发布者ID为userID的视频
func GetVideoList(ctx context.Context, userID int64) ([]Video, error) {
	var videos []Video
	err := TIK_DB.WithContext(ctx).Model(&Video{}).Where("user_id = ?", userID).Find(&videos).Error
	return videos, err
}

// GetUserList 查找用户列表
func GetUserList(ctx context.Context, userIDs []int64) ([]User, error) {
	var users []User
	err := TIK_DB.WithContext(ctx).Model(&User{}).Find(&users, userIDs).Error
	return users, err
}

// GetFavoriteList 查找表中用户ID为userID的点赞
func GetFavoriteList(ctx context.Context, userID int64) ([]UserFavorite, error) {
	var favorites []UserFavorite
	err := TIK_DB.WithContext(ctx).Model(&UserFavorite{}).Where("user_id = ?", userID).Find(&favorites).Error
	return favorites, err
}

// GetRelationList 查找用户的关系列表
func GetRelationList(ctx context.Context, userID int64) ([]UserFocusOn, error) {
	var userRelations []UserFocusOn
	err := TIK_DB.WithContext(ctx).Model(&UserFocusOn{}).
		Where("user_id = ?", userID).Or("to_user_id = ?", userID).Find(&userRelations).Error
	return userRelations, err
}
func CreateUserReleaseVideo(ctx context.Context, userId int64, videoId int64) error {
	if err := TIK_DB.WithContext(ctx).Create(&UserReleaseVideo{
		UserID:  userId,
		VideoID: videoId,
	}).Error; err != nil {
		return err
	}
	return nil
}

// GetUser 查找用户
func GetUser(ctx context.Context, userID int64) (*User, error) {
	user := new(User)
	err := TIK_DB.WithContext(ctx).Model(&User{}).First(user, userID).Error
	return user, err
}

// QueryFocus 在表中查询两个人当前的关注状态
func QueryFocus(ctx context.Context, userID int64, toUserID int64) (*UserFocusOn, error) {
	userList1 := []int64{userID, toUserID}
	userList2 := []int64{toUserID, userID}
	userFocus := new(UserFocusOn)
	err := TIK_DB.WithContext(ctx).Model(&UserFocusOn{}).Where("user_id in ?", userList1).
		Where("to_user_id in ?", userList2).First(userFocus).Error
	return userFocus, err
}
