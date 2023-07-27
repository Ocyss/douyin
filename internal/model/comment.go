package model

type (
	// Comment 评论表
	Comment struct {
		Model
		UserID  int64  `json:"-" gorm:"index:idx_uvid;comment:评论用户信息"`
		VideoID int64  `json:"-" gorm:"index:idx_uvid;comment:评论视频信息"`
		User    User   `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
		Video   Video  `json:"video" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
		Content string `json:"content" gorm:"comment:评论内容"`
		//create_date string // 评论发布日期，格式 mm-dd
		// 自建字段
		ReplyID int64 `json:"reply_id" gorm:"index;comment:回复ID"`
	}
)

func init() {
	addMigrate(&Comment{})
}
