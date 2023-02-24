package controllor

//定义http请求的结构体

// FeedParam 视频流请求参数
type FeedParam struct {
	LatestTime string `query:"latest_time"` // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
	Token      string `query:"token"`       // 可选参数，登录用户设置
}

// RegisterParam 用户注册请求参数
type RegisterParam struct {
	UserName string `query:"username"` // 用户名
	PassWord string `query:"password"` // 用户密码
}

// LoginParam 用户登录请求参数
type LoginParam struct {
	UserName string `query:"username"` // 用户名
	PassWord string `query:"password"` // 用户密码
}

//type Login struct {
//	Username string `json:"username" xml:"username" form:"username" query:"username"`
//	Password string `json:"password" xml:"password" form:"password" query:"password"`
//}

// UserInfoParam 用户信息请求参数
type UserInfoParam struct {
	UserId string `query:"user_id"` // 用户id
	Token  string `query:"token"`   // 用户鉴权token
}

// PublishParam 发布视频操作请求参数
type PublishParam struct {
	Data  []byte `query:"data"`  // 视频数据
	Token string `query:"token"` // 用户鉴权token
	Title string `query:"title"` // 视频标题
}

// PublishListParam 发布列表请求参数
type PublishListParam struct {
	Token  string `query:"token"`   // 用户鉴权token
	UserID int64  `query:"user_id"` // 用户id
}

// FavoriteParam 点赞操作请求参数
type FavoriteParam struct {
	Token      string `query:"token"`       // 用户鉴权token
	VideoId    string `query:"video_id"`    // 视频id
	ActionType string `query:"action_type"` // 1-点赞，2-取消点赞
}

// FavoriteListParam 喜欢列表
type FavoriteListParam struct {
	UserID string `query:"user_id"` // 用户id
	Token  string `query:"token"`   // 用户鉴权token
}

// CommentParam 评论操作 请求参数
type CommentParam struct {
	Token       string `query:"token"`        // 用户鉴权token
	VideoId     string `query:"video_id"`     // 视频id
	ActionType  string `query:"action_type"`  // 1-发布评论，2-删除评论
	CommentText string `query:"comment_text"` // 用户填写的评论内容，在action_type=1的时候使用
	CommentId   string `query:"comment_id"`   // 要删除的评论id，在action_type=2的时候使用
}

// CommentListParam 获取评论列表请求参数
type CommentListParam struct {
	Token   string `query:"token"`    // 用户鉴权token
	VideoId string `query:"video_id"` // 视频id
}

// FollowParam 关注操作请求参数
type FollowParam struct {
	Token      string `query:"token"`       // 用户鉴权token
	ToUserId   string `query:"to_user_id"`  // 对方用户id
	ActionType string `query:"action_type"` // 1-关注，2-取消关注
}

// FollowListParam 关注列表请求参数
type FollowListParam struct {
	UserID string `query:"user_id"` // 用户id
	Token  string `query:"token"`   // 用户鉴权token
}

// FollowerListParam 粉丝列表请求参数
type FollowerListParam struct {
	UserID string `query:"user_id"` // 用户id
	Token  string `query:"token"`   // 用户鉴权token
}

// FriendListParam 好友列表请求参数
type FriendListParam struct {
	UserID string `query:"user_id"` // 用户id
	Token  string `query:"token"`   // 用户鉴权token
}

// MessageParam 发送消息请求参数
type MessageParam struct {
	Token      string `query:"token"`       // 用户鉴权token
	ToUserID   string `query:"to_user_id"`  // 对方的用户id
	ActionType string `query:"action_type"` // 1-发送消息
	Content    string `query:"content"`     //消息内容
}

// MessageChatParam 聊天记录请求参数
type MessageChatParam struct {
	Token      string `query:"token"`        // 用户鉴权token
	ToUserId   int64  `query:"to_user_id"`   // 对方用户id
	PreMsgTime int64  `query:"pre_msg_time"` //上次最新消息的时间（新增字段-apk更新中）
}
