package db

import (
	"context"
	"gorm.io/gorm"
)

// Video 视频数据库模型
type Video struct {
	gorm.Model
	Title         string `json:"title" gorm:"comment:视频标题"`             // 视频标题
	PlayUrl       string `json:"play_url" gorm:"comment:视频播放地址"`        // 视频播放地址
	CoverUrl      string `json:"cover_url" gorm:"comment:视频封面地址"`       // 视频封面地址
	FavoriteCount int64  `json:"favorite_count" gorm:"comment:视频的点赞总数"` // 视频的点赞总数
	CommentCount  int64  `json:"comment_count" gorm:"comment:视频的评论总数"`  // 视频的评论总数
	Name          string `json:"name" gorm:"comment:文件名"`               // 文件名
	Tag           string `json:"tag" gorm:"comment:文件标签"`               // 文件标签
	Key           string `json:"key" gorm:"comment:编号"`                 //编号
	UserID        int64  `json:"user_id" gorm:"comment:发布者的ID"`         // 发布者的ID
}

func (v *Video) TableName() string {
	return "videos"
}

// Comment 用户评论数据库模型
type Comment struct {
	gorm.Model
	UserID      int64  `json:"user_id" gorm:"comment:发表评论的用户id"` // 发表评论的用户ID
	VideoID     int64  `json:"video_id" gorm:"comment:视频id"`     // 视频ID
	CommentText string `json:"comment_text" gorm:"comment:评论内容"` // 评论内容
}

func (c *Comment) TableName() string {
	return "user_comment_videos"
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

// CreateComment 在表中创建评论
func CreateComment(ctx context.Context, comment *Comment) (int64, error) {
	var commentID int64
	//使用事务
	err := TIK_DB.Transaction(func(tx *gorm.DB) error {
		//创建评论
		result := tx.WithContext(ctx).Model(&Comment{}).Create(comment)
		if result.Error != nil {
			return result.Error
		}
		commentID = int64(comment.ID)
		video := new(Video)
		//获得视频评论数
		if err := tx.WithContext(ctx).Model(&Video{}).First(video, comment.VideoID).Error; err != nil {
			return err
		}
		video.CommentCount++
		//评论数加1后保存数据库
		if err := tx.WithContext(ctx).Model(&Video{}).Where("id = ?", video.ID).Save(video).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return 0, err
	}
	return commentID, nil
}

// DeleteComment 在表中删除评论
func DeleteComment(ctx context.Context, comment *Comment) error {
	//使用事务
	return TIK_DB.Transaction(func(tx *gorm.DB) error {
		//找到要删除的评论
		if err := tx.WithContext(ctx).Model(&Video{}).First(comment, comment.ID).Error; err != nil {
			return err
		}
		//删除该评论
		if err := tx.WithContext(ctx).Model(&Comment{}).Delete(&Comment{}, comment.ID).Error; err != nil {
			return err
		}
		video := new(Video)
		//获得视频评论数
		if err := tx.WithContext(ctx).Model(&Video{}).First(video, comment.VideoID).Error; err != nil {
			return err
		}
		if video.CommentCount > 0 {
			video.CommentCount--
		}
		//评论数减1后保存数据库
		if err := tx.WithContext(ctx).Model(&video).Save(video).Error; err != nil {
			return err
		}
		return nil
	})
}

// GetCommentList 查找表中视频ID为videoID的评论
func GetCommentList(ctx context.Context, videoID int64) ([]Comment, error) {
	var comments []Comment
	err := TIK_DB.WithContext(ctx).Model(&Comment{}).Where("video_id = ?", videoID).Find(&comments).Error
	return comments, err
}

// GetUser 查找用户
func GetUser(ctx context.Context, userIDs int64) (*User, error) {
	users := new(User)
	err := TIK_DB.WithContext(ctx).Model(&User{}).First(users, userIDs).Error
	return users, err
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
