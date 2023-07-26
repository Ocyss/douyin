package model

type Comment struct {
	Model
	UserID  int64  `json:"-"`
	User    User   `json:"user" gorm:"comment:评论用户信息"`
	Content string `json:"content" gorm:"comment:评论内容"`
	//create_date string // 评论发布日期，格式 mm-dd
}
