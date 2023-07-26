package model

type (
	User struct {
		Model
		Name            string  `json:"name" gorm:"comment:用户名称,index"`
		FollowCount     int64   `json:"follow_count" gorm:"comment:关注总数"`
		FollowerCount   int64   `json:"follower_count" gorm:"comment:粉丝总数"`
		IsFollow        bool    `json:"is_follow" gorm:"comment:是否关注"`
		Avatar          string  `json:"avatar" gorm:"comment:用户头像"`
		BackgroundImage string  `json:"background_image" gorm:"comment:用户个人页顶部大图"`
		Signature       string  `json:"signature" gorm:"comment:个人简介"`
		TotalFavorited  int64   `json:"total_favorited" gorm:"comment:获赞数量"`
		WorkCount       int64   `json:"work_count" gorm:"comment:作品数量"`
		FavoriteCount   int64   `json:"favorite_count" gorm:"comment:点赞数量"`
		Follow          []*User `json:"follow,omitempty" gorm:"many2many:user_follow;comment:关注列表"`
		Follower        []*User `json:"follower,omitempty" gorm:"many2many:user_follower;comment:粉丝列表"`
		Friend          []*User `json:"friend,omitempty" gorm:"many2many:user_friend;comment:好友列表"`
		Favorite        []Video `json:"video_list,omitempty" gorm:"many2many:favorite;comment:喜欢列表"`
	}
	FriendUser struct {
		User
		Message string `json:"message" gorm:"comment:和该好友的最新聊天消息"`
		MsgType bool   `json:"msg_type,number" gorm:"comment:消息类型"` // 0 => 当前请求用户接收的消息， 1 => 当前请求用户发送的消息
	}
)
