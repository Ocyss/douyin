package utils

import (
	"math/rand"
	"time"
)

var r *rand.Rand

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// GetId 获取ID
func GetId(a, b int64) int64 {
	// 使用纳秒时间戳，保证递增
	now := time.Now()
	return now.UnixNano()*a - b
}

// RandString 获取随机字符串
func RandString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[r.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

func RandVid(all []int64, n int) (res []int64) {
	if len(all) <= n {
		return all
	}
	set := make(map[int64]struct{}, n)
	for len(set) < n {
		if index := r.Intn(len(all)); index >= 0 {
			set[all[index]] = struct{}{}
		}
	}
	for k := range set {
		res = append(res, k)
	}
	return
}

func RandShuffle(n int, f func(int, int)) {
	r.Shuffle(n, f)
}

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}
