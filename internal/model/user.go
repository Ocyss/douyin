package model

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/Ocyss/douyin/utils"
	"gorm.io/gorm"
)

type (
	// User 用户信息表
	User struct {
		Model
		Name            string     `json:"name" gorm:"index:,unique;size:32;comment:用户名称"`
		Pawd            string     `json:"-" gorm:"size:128;comment:用户密码"`
		FollowCount     int64      `json:"follow_count" gorm:"-"`   // 关注总数
		FollowerCount   int64      `json:"follower_count" gorm:"-"` // 粉丝总数
		IsFollow        bool       `json:"is_follow" gorm:"-"`      // 是否关注
		Avatar          string     `json:"avatar" gorm:"comment:用户头像"`
		BackgroundImage string     `json:"background_image" gorm:"comment:用户个人页顶部大图"`
		Signature       string     `json:"signature" gorm:"default:此人巨懒;comment:个人简介"`
		WorkCount       int64      `json:"work_count" gorm:"default:0;comment:作品数量"`
		TotalFavorited  int64      `json:"total_favorited" gorm:"-"` // TODO: 获赞数量
		FavoriteCount   int64      `json:"favorite_count" gorm:"-"`  // TODO: 点赞数量
		Follow          []*User    `json:"follow,omitempty" gorm:"many2many:UserFollow;comment:关注列表"`
		Follower        []*User    `json:"follower,omitempty" gorm:"-"` // 粉丝列表
		Friend          []*User    `json:"friend,omitempty" gorm:"-"`   // 好友列表
		Favorite        []*Video   `json:"like_list,omitempty" gorm:"many2many:UserFavorite;comment:喜欢列表"`
		Videos          []*Video   `json:"video_list,omitempty" gorm:"many2many:UserCreation;comment:作品列表"`
		Comment         []*Comment `json:"comment_list,omitempty" gorm:"comment:评论列表"`
	}
	// FriendUser 好友结构体
	FriendUser struct {
		User
		Message string `json:"message" gorm:"comment:和该好友的最新聊天消息"`
		MsgType bool   `json:"msg_type,number" gorm:"comment:消息类型"` // 0 => 当前请求用户接收的消息， 1 => 当前请求用户发送的消息
	}
	// UserCreation 联合作者
	UserCreation struct {
		VideoID   int64          `json:"video_id,omitempty" gorm:"primaryKey"`
		UserID    int64          `json:"author_id" gorm:"primaryKey"`
		Type      string         `json:"type" gorm:"comment:创作者类型"` // Up主, 参演，剪辑，录像，道具，编剧，打酱油
		CreatedAt time.Time      `json:"created_at"`
		DeletedAt gorm.DeletedAt `json:"-"`
	}
)

var (
	userFollowCountKey   = make([]byte, 0, 50)
	userFollowerCountKey = make([]byte, 0, 50)
)

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == 0 {
		u.ID = utils.GetId(2, 114514)
	}
	if u.Avatar == "" {
		url := make([]byte, 0, 88)
		url = append(url, "https://api.multiavatar.com/"...)
		url = append(url, u.Name...)
		url = append(url, ".png"...)
		u.Avatar = string(url)
	}
	if u.BackgroundImage == "" {
		u.BackgroundImage = "https://api.paugram.com/wallpaper/"
	}
	return
}

func (u *User) AfterFind(tx *gorm.DB) (err error) {
	if uid, ok := tx.Get("user_id"); ok || u.ID != 0 {
		result := map[string]any{}
		u.IsFollow = tx.Table("user_follow").Where("follow_id = ? AND user_id = ?", u.ID, uid).Take(&result).RowsAffected == 1
	}
	// tx.Table("user_follow").Where("user_id = ?", u.ID).Count(&u.FollowCount)
	// tx.Table("user_follow").Where("follow_id = ?", u.ID).Count(&u.FollowerCount)
	u.FollowCount = getUserFollowCount(tx, u.ID)
	u.FollowerCount = getUserFollowerCount(tx, u.ID)
	return
}

// getUserFollowCount 获取关注数
func getUserFollowCount(tx *gorm.DB, uid int64) int64 {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	key := getKey(uid, userFollowCountKey)
	FollowCount, err := rdb.Get(ctx, key).Int64()
	if err == redis.Nil {
		tx.Table("user_follow").Where("user_id = ?", uid).Count(&FollowCount)
		_ = rdb.Set(ctx, key, FollowCount, 3*time.Second)
	}
	return FollowCount
}

// getUserFollowerCount 获取粉丝数
func getUserFollowerCount(tx *gorm.DB, uid int64) int64 {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	key := getKey(uid, userFollowerCountKey)
	FollowerCount, err := rdb.Get(ctx, key).Int64()
	if err == redis.Nil {
		tx.Table("user_follow").Where("follow_id = ?", uid).Count(&FollowerCount)
		_ = rdb.Set(ctx, key, FollowerCount, 3*time.Second)
	}
	return FollowerCount
}

func init() {
	addMigrate(&User{}, &UserCreation{})
	userFollowCountKey = append(userFollowCountKey, "user:follow_count/"...)
	userFollowerCountKey = append(userFollowerCountKey, "user:follower_count/"...)
}
