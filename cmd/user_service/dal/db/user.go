package db

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

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

// CreateUser 在表中创建新用户
func CreateUser(ctx context.Context, user *User) error {
	return TIK_DB.WithContext(ctx).Model(&User{}).Create(user).Error
}

// UserQuery 在表中查询用户账号密码是否正确
func UserQuery(ctx context.Context, user *User) error {
	return TIK_DB.WithContext(ctx).Model(&User{}).
		Where("user_name = ?", user.Username).
		Where("password = ?", user.Password).First(user).Error
}

// UserNameQuery 在表中查询用户名
func UserNameQuery(ctx context.Context, user *User) (*User, error) {
	fmt.Println("__________")
	userRes := new(User)
	err := TIK_DB.WithContext(ctx).Model(&User{}).
		Where("username = ?", user.Username).First(userRes).Error
	return userRes, err
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

// GetUser 查找用户
func GetUser(ctx context.Context, userID int64) (*User, error) {
	user := new(User)
	err := TIK_DB.WithContext(ctx).Model(&User{}).First(user, userID).Error
	return user, err
}
