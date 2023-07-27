package model

type (
	// Video 视频表
	Video struct {
		Model
		AuthorID      int64      `json:"-" gorm:"index"`
		Author        User       `json:"author" gorm:"comment:视频作者信息"`
		PlayUrl       string     `json:"play_url" gorm:"comment:视频播放地址"`
		CoverUrl      string     `json:"cover_url" gorm:"comment:视频封面地址"`
		FavoriteCount int64      `json:"favorite_count" gorm:"comment:视频的点赞总数"`
		CommentCount  int64      `json:"comment_count" gorm:"comment:视频的评论总数"`
		PlayCount     int64      `json:"play_count" gorm:"comment:视频的播放量"`
		IsFavorite    bool       `json:"is_favorite" gorm:"-"` // 是否点赞
		Title         string     `json:"title" gorm:"comment:视频标题"`
		Desc          string     `json:"desc" gorm:"comment:简介"`
		Comment       []*Comment `json:"comment,omitempty"` // 评论列表
		// 自建字段
		CoAuthor []*CoAuthor `json:"authors,omitempty" gorm:"foreignKey:AuthorID;comment:视频作者信息"`
	}
	// CoAuthor 联合作者
	CoAuthor struct {
		VideoID  int64  `json:"video_id,omitempty"`
		Video    Video  `json:"-" gorm:"primaryKey"`
		AuthorID int64  `json:"author_id"`
		Author   User   `json:"-" gorm:"primaryKey"`
		Type     string `json:"type" gorm:"comment:创作者类型"` //参演，剪辑，录像，道具，编剧，打酱油
	}
)

func init() {
	addMigrate(&Video{}, &CoAuthor{})
}
