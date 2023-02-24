package db

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	ToUserId   int64  `json:"to_user_id" gorm:"comment:该消息接收者的id"`   // 该消息接收者的id
	FromUserId int64  `json:"from_user_id" gorm:"comment:该消息发送者的id"` // 该消息发送者的id
	Content    string `json:"content" gorm:"comment:消息内容"`           // 消息内容
	MsgType    int32  `json:"msg_type" gorm:"comment:消息状态"`          //0 => 当前请求用户接收的消息， 1 => 当前请求用户发送的消息
	IfLast     bool   `json:"if_last" gorm:"comment:是否为最新"`
}

func (fa *Message) TableName() string {
	return "message"
}

type UserFocusOn struct {
	gorm.Model
	UserID   int64 `json:"user_id" gorm:"comment:用户id"`      // 用户ID
	ToUserID int64 `json:"to_user_id" gorm:"comment:对方用户id"` // 对方用户id
	IsFollow int32 `json:"is_follow"  gorm:"comment:双方关注状态,1表示user关注to_user,2表示to_user关注user,3表示互关"`
}
type Friend struct {
	UserID   int64 `json:"user_id" gorm:"comment:用户id"`      // 用户ID
	ToUserID int64 `json:"to_user_id" gorm:"comment:对方用户id"` // 对方用户id
}

type User struct {
	gorm.Model
	Username      string `json:"username" gorm:"comment:用户登录名"`      // 用户登录名
	Password      string `json:"-"  gorm:"comment:用户登录密码"`           // 用户登录密码
	FollowCount   int64  `json:"follow_count" gorm:"comment:关注总数"`   // 关注总数
	FollowerCount int64  `json:"follower_count" gorm:"comment:粉丝总数"` // 粉丝总数
	LikeTotal     int64  `json:"like_total" gorm:"comment:获赞总数"`     // 获赞总数
}
type FriendUser struct {
	Id            int64  `json:"id" gorm:"comment:用户登录名"`
	Username      string `json:"username" gorm:"comment:用户登录名"`      // 用户登录名
	Password      string `json:"-"  gorm:"comment:用户登录密码"`           // 用户登录密码
	FollowCount   int64  `json:"follow_count" gorm:"comment:关注总数"`   // 关注总数
	FollowerCount int64  `json:"follower_count" gorm:"comment:粉丝总数"` // 粉丝总数
	LikeTotal     int64  `json:"like_total" gorm:"comment:获赞总数"`     // 获赞总数
	Content       string `json:"content" gorm:"comment:消息内容"`        // 消息内容
	MsgType       int32  `json:"msg_type" gorm:"comment:消息状态"`       //0 => 当前请求用户接收的消息， 1 => 当前请求用户发送的消息
}

func GetMessageList(ctx context.Context, PreMsgTime int64, toUserId int64, userId int64) ([]Message, error) {

	var message []Message
	err := TIK_DB.WithContext(ctx).Model(&Message{}).
		Raw("select * from message where (from_user_id= ? and to_user_id= ? ) or (to_user_id= ?  and from_user_id= ? ) ORDER BY created_at desc limit 10", toUserId, userId, toUserId, userId).
		Scan(&message).Error
	fmt.Println(message)
	return message, err
}
func CreateMessage(ctx context.Context, content string, toUserId int64, userId int64) error {
	messageDB := &Message{
		ToUserId:   toUserId,
		FromUserId: userId,
		Content:    content,
		MsgType:    1,
		IfLast:     true,
	}
	return TIK_DB.WithContext(ctx).Model(&Message{}).Create(messageDB).Error
	//err := TIK_DB.WithContext(ctx).Model(&Message{}).Create(&Message{
	//	ToUserId:   toUserId,
	//	FromUserId: userId,
	//	Content:    content,
	//	IfLast:     true,
	//}).Error
	//return err
}
func GetFriendList(ctx context.Context, userId int64) ([]Friend, error) {
	var friend []Friend
	err := TIK_DB.WithContext(ctx).Model(&UserFocusOn{}).
		Raw("select user_id,to_user_id from user_focus_ons where user_id=? or to_user_id=? and is_follow=?", userId, userId, 3).
		Scan(&friend).Error
	fmt.Println(friend)
	return friend, err
}
func GetUserById(ctx context.Context, friendId []int64, userId int64) ([]FriendUser, error) {
	fmt.Println(friendId)
	fmt.Println(userId)
	var friendUser []FriendUser
	err := TIK_DB.WithContext(ctx).Model(&Message{}).
		Raw("select u.id,username,follow_count,follower_count,content,msg_type from message as m  LEFT JOIN  users as u  on ( m.from_user_id=u.id or m.to_user_id=u.id) where (m.to_user_id= ?  and m.from_user_id in ?) or (m.from_user_id=? and m.to_user_id in ?) and m.if_last=1 ", userId, friendId, userId, friendId).
		Scan(&friendUser).Error
	fmt.Println(friendUser)
	return friendUser, err
}
