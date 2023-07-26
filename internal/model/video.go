package model

type Video struct {
	Model
	AuthorID      int64  `json:"-"`
	Author        User   `json:"author" gorm:"comment:视频作者信息,index"`
	PlayUrl       string `json:"play_url" gorm:"comment:视频播放地址"`
	CoverUrl      string `json:"cover_url" gorm:"comment:视频封面地址"`
	FavoriteCount int64  `json:"favorite_count" gorm:"comment:视频的点赞总数"`
	CommentCount  int64  `json:"comment_count" gorm:"comment:视频的评论总数"`
	IsFavorite    bool   `json:"is_favorite" gorm:"-"` // 是否点赞
	Title         string `json:"title" gorm:"comment:视频标题"`
}
