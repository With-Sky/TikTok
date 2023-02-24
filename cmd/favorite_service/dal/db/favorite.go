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

// CreateFavorite 在表中创建喜欢
func CreateFavorite(ctx context.Context, favorite *UserFavorite) error {
	//使用事务
	//global.LOG.Error("数据库")
	//return nil
	return TIK_DB.Transaction(func(tx *gorm.DB) error {
		//创建喜欢
		if err := tx.WithContext(ctx).Model(&UserFavorite{}).Create(favorite).Error; err != nil {
			//global.LOG.Error("数据库创建喜欢失败")
			return err
		}
		video := new(Video)
		//获得视频点赞数
		if err := tx.WithContext(ctx).Model(&Video{}).First(video, favorite.VideoID).Error; err != nil {
			//global.LOG.Error("数据库获得点赞数失败")
			return err
		}
		video.FavoriteCount++
		//点赞数加1后保存数据库
		if err := tx.WithContext(ctx).Model(&Video{}).Where("id = ?", video.ID).Save(video).Error; err != nil {
			//global.LOG.Error("数据库保存失败")
			return err
		}
		//发布视频的用户获赞数加一
		if err := tx.WithContext(ctx).Model(&User{}).
			Exec("update users set like_total=like_total+1 WHERE users.id = ?", video.UserID).Error; err != nil {
			//global.LOG.Error("数据库获赞数增加失败")
			return err
		}
		return nil
	})
}

// DeleteFavorite 在表中删除喜欢
func DeleteFavorite(ctx context.Context, favorite *UserFavorite) error {
	//使用事务
	return TIK_DB.Transaction(func(tx *gorm.DB) error {
		//删除喜欢
		if err := tx.WithContext(ctx).Model(&UserFavorite{}).Where("user_id = ?", favorite.UserID).
			Where("video_id = ?", favorite.VideoID).Delete(&UserFavorite{}).Error; err != nil {
			//global.LOG.Error("数据库删除喜欢失败")
			return err
		}
		video := new(Video)
		//获得视频点赞数
		if err := tx.WithContext(ctx).Model(&Video{}).First(video, favorite.VideoID).Error; err != nil {
			//global.LOG.Error("数据库获得点赞数失败")
			return err
		}
		if video.FavoriteCount > 0 {
			video.FavoriteCount--
		}
		//点赞数减1后保存数据库
		if err := tx.WithContext(ctx).Model(&video).Save(video).Error; err != nil {
			//global.LOG.Error("数据库保存失败")
			return err
		}
		if err := tx.Exec("update users set like_total=CASE WHEN like_total > 0 THEN "+
			"like_total-1 ELSE 0 END WHERE users.id = ?", video.UserID).Error; err != nil {
			//global.LOG.Error("数据库获赞数减少失败")
			return err
		}
		return nil
	})
}

// GetFavoriteList 查找表中用户ID为userID的点赞
func GetFavoriteList(ctx context.Context, userID int64) ([]UserFavorite, error) {
	var favorites []UserFavorite
	err := TIK_DB.WithContext(ctx).Model(&UserFavorite{}).Where("user_id = ?", userID).Find(&favorites).Error
	//global.LOG.Error("数据库获取视频评论数失败")
	return favorites, err
}

// GetVideoList 查找视频列表
func GetVideoList(ctx context.Context, videoIDs []int64) ([]Video, error) {
	var videos []Video
	err := TIK_DB.WithContext(ctx).Model(&Video{}).Find(&videos, videoIDs).Error
	//global.LOG.Error("数据库获取视频评论数失败")
	return videos, err
}

// GetUserList 查找用户列表
func GetUserList(ctx context.Context, userIDs []int64) ([]User, error) {
	var users []User
	err := TIK_DB.WithContext(ctx).Model(&User{}).Find(&users, userIDs).Error
	//global.LOG.Error("数据库获取视频评论数失败")
	return users, err
}

// GetRelationList 查找用户的关系列表
func GetRelationList(ctx context.Context, userID int64) ([]UserFocusOn, error) {
	var userRelations []UserFocusOn
	err := TIK_DB.WithContext(ctx).Model(&UserFocusOn{}).
		Where("user_id = ?", userID).Or("to_user_id = ?", userID).Find(&userRelations).Error
	//global.LOG.Error("数据库获取视频评论数失败")
	return userRelations, err
}
