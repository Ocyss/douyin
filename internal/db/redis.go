package db

import (
	"context"
	"github.com/Ocyss/douyin/internal/model"
	"github.com/redis/go-redis/v9"
	"strconv"
	"strings"
	"sync"
	"time"
)

var playCountUpdates = make(chan string, 35)

// getFavoriteCount 视频的点赞总数
func getFavoriteCount(wg *sync.WaitGroup, vid int64, val *int64) {
	defer wg.Done()
	var builder strings.Builder
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	builder.Grow(50)
	builder.WriteString("video:favorite_count/")
	builder.WriteString(strconv.FormatInt(vid, 10))
	favoriteCount, err := rdb.Get(ctx, builder.String()).Int64()
	if err == redis.Nil {
		db.Table("user_favorite").Where("video_id = ?", vid).Count(&favoriteCount)
		_ = rdb.Set(ctx, builder.String(), favoriteCount, 300*time.Second)
	}
	*val = favoriteCount
}

// getCommentCount 视频的评论总数
func getCommentCount(wg *sync.WaitGroup, vid int64, val *int64) {
	defer wg.Done()
	var builder strings.Builder
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	builder.Grow(50)
	builder.WriteString("video:comment_count/")
	builder.WriteString(strconv.FormatInt(vid, 10))
	CommentCount, err := rdb.Get(ctx, builder.String()).Int64()
	if err == redis.Nil {
		db.Model(&model.Comment{}).Where("video_id = ?", vid).Count(&CommentCount)
		_ = rdb.Set(ctx, builder.String(), CommentCount, 300*time.Second)
	}
	*val = CommentCount
}

// setPlayCount 视频的播放量
func setPlayCount(wg *sync.WaitGroup, ip string, vid int64, val *int64) {
	defer wg.Done()
	var builder strings.Builder
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	builder.Grow(50)
	builder.WriteString("video:play_count/")
	builder.WriteString(strconv.FormatInt(vid, 10))
	rdb.PFAdd(ctx, builder.String(), ip)
	*val, _ = rdb.PFCount(ctx, builder.String()).Result()
	go func() {
		playCountUpdates <- builder.String()
		if len(playCountUpdates) >= 30 {
			for i := 0; i < 30; i++ {
				v := <-playCountUpdates
				val, _ := rdb.PFCount(ctx, v).Result()
				db.Model(&model.Video{}).Where("id = ?", v[17:]).Update("play_count", val)
			}
		}
	}()
}

// getIsFavorite 视频是否点赞
func getIsFavorite(wg *sync.WaitGroup, uid, vid int64, val *bool) {
	defer wg.Done()
	result := map[string]any{}
	*val = db.Table("user_favorite").Where("user_id = ? AND video_id = ?", uid, vid).Take(&result).RowsAffected == 1
	//data[i].IsFavorite = db.Raw("SELECT * FROM user_favorite WHERE user_id = ? AND video_id = ?", uid, data[i].ID).Scan(&result).RowsAffected == 1
}
