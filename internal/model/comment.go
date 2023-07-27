package model

type (
	// Comment 评论表
	Comment struct {
		Model
		UserID  int64  `json:"-" gorm:"index"`
		VideoID int64  `json:"-" gorm:"index"`
		User    User   `json:"user" gorm:"comment:评论用户信息"`
		Video   Video  `json:"video" gorm:"comment:评论视频信息"`
		Content string `json:"content" gorm:"comment:评论内容"`
		//create_date string // 评论发布日期，格式 mm-dd
		// 自建字段
		ReplyID int64 `json:"reply_id" gorm:"index;comment:回复ID"`
	}
)

func init() {
	addMigrate(&Comment{})
}
