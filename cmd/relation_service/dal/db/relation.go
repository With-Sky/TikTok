package db

import (
	"context"
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

// Message 消息数据库模型
type Message struct {
	gorm.Model
	ToUserId   int64  `json:"to_user_id" gorm:"comment:该消息接收者的id"`   // 该消息接收者的id
	FromUserId int64  `json:"from_user_id" gorm:"comment:该消息发送者的id"` // 该消息发送者的id
	Content    string `json:"content" gorm:"comment:消息内容"`           // 消息内容
	MsgType    int32  `json:"msg_type" gorm:"comment:消息状态"`          //0 => 当前请求用户接收的消息， 1 => 当前请求用户发送的消息
}

func (m *Message) TableName() string {
	return "message"
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

// CreateFocus 在表中创建关注
func CreateFocus(ctx context.Context, userFocus *UserFocusOn) error {
	//使用事务
	return TIK_DB.Transaction(func(tx *gorm.DB) error {
		//创建关注
		if err := tx.WithContext(ctx).Model(&UserFocusOn{}).Create(userFocus).Error; err != nil {
			return err
			//global.LOG.Error("数据库创建关注失败")
		}
		//本人的关注数加一
		if err := tx.WithContext(ctx).Model(&User{}).
			Exec("update users set follow_count=follow_count+1 WHERE users.id = ?", userFocus.UserID).Error; err != nil {
			//global.LOG.Error("数据库保存关注增加失败")
			return err
		}
		//对方的粉丝数加一
		if err := tx.WithContext(ctx).Model(&User{}).
			Exec("update users set follower_count=follower_count+1 WHERE users.id = ?", userFocus.ToUserID).Error; err != nil {
			//global.LOG.Error("数据库保存粉丝增加失败")
			return err
		}
		return nil
	})
}

// DeleteFocusByID 通过主键删除关注
func DeleteFocusByID(ctx context.Context, userFocus *UserFocusOn, userID int64, toUserID int64) error {
	//使用事务
	return TIK_DB.Transaction(func(tx *gorm.DB) error {
		//删除关注
		if err := tx.WithContext(ctx).Model(&UserFocusOn{}).Delete(&UserFocusOn{}, userFocus.ID).Error; err != nil {
			//global.LOG.Error("数据库删除关注失败")
			return err
		}
		//本人的关注数减一
		if err := tx.WithContext(ctx).Model(&User{}).
			Exec("update users set follow_count=CASE WHEN follow_count > 0 THEN "+
				"follow_count-1 ELSE 0 END WHERE users.id = ?", userID).Error; err != nil {
			//global.LOG.Error("数据库保存关注减少失败")
			return err
		}
		//对方的粉丝数减一
		if err := tx.WithContext(ctx).Model(&User{}).
			Exec("update users set follower_count=CASE WHEN follower_count > 0 THEN "+
				"follower_count-1 ELSE 0 END WHERE users.id = ?", toUserID).Error; err != nil {
			//global.LOG.Error("数据库保存粉丝减少失败")
			return err
		}
		return nil
	})
}

// IncreaseFocus 在表中保存关注，并增加userID用户的关注数，增加toUserID用户的粉丝数
func IncreaseFocus(ctx context.Context, userFocus *UserFocusOn, userID int64, toUserID int64) error {
	//使用事务
	return TIK_DB.Transaction(func(tx *gorm.DB) error {
		//保存关注
		if err := tx.WithContext(ctx).Model(&UserFocusOn{}).Save(userFocus).Error; err != nil {
			//global.LOG.Error("数据库保存关注失败")
			return err
		}
		//本人的关注数加一
		if err := tx.WithContext(ctx).Model(&User{}).
			Exec("update users set follow_count=follow_count+1 WHERE users.id = ?", userID).Error; err != nil {
			//global.LOG.Error("数据库保存关注增加失败")
			return err
		}
		//对方的粉丝数加一
		if err := tx.WithContext(ctx).Model(&User{}).
			Exec("update users set follower_count=follower_count+1 WHERE users.id = ?", toUserID).Error; err != nil {
			return err
		}
		return nil
	})
}

// DecreaseFocus 在表中保存关注，并减少userID用户的关注数，减少toUserID用户的粉丝数
func DecreaseFocus(ctx context.Context, userFocus *UserFocusOn, userID int64, toUserID int64) error {
	//使用事务
	return TIK_DB.Transaction(func(tx *gorm.DB) error {
		//保存关注
		if err := tx.WithContext(ctx).Model(&UserFocusOn{}).Save(userFocus).Error; err != nil {
			return err
		}
		//本人的关注数减一
		if err := tx.WithContext(ctx).Model(&User{}).
			Exec("update users set follow_count=CASE WHEN follow_count > 0 THEN "+
				"follow_count-1 ELSE 0 END WHERE users.id = ?", userID).Error; err != nil {
			return err
		}
		//对方的粉丝数减一
		if err := tx.WithContext(ctx).Model(&User{}).
			Exec("update users set follower_count=CASE WHEN follower_count > 0 THEN "+
				"follower_count-1 ELSE 0 END WHERE users.id = ?", toUserID).Error; err != nil {
			return err
		}
		return nil
	})
}

// GetRelationList 查找用户的关系列表
func GetRelationList(ctx context.Context, userID int64) ([]UserFocusOn, error) {
	var users []UserFocusOn
	err := TIK_DB.WithContext(ctx).Model(&UserFocusOn{}).
		Where("user_id = ?", userID).Or("to_user_id = ?", userID).Find(&users).Error
	return users, err
}

// GetUserList 查找用户列表
func GetUserList(ctx context.Context, userIDs []int64) ([]User, error) {
	var users []User
	err := TIK_DB.WithContext(ctx).Model(&User{}).Find(&users, userIDs).Error
	return users, err
}

func CreateMessage(ctx context.Context, content string, toUserId int64, userId int64) error {
	err := TIK_DB.WithContext(ctx).Create(&Message{
		Model:      gorm.Model{},
		ToUserId:   toUserId,
		FromUserId: userId,
		Content:    content,
	}).Error
	return err
}
