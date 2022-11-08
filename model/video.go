package model

import (
	"giligili/cache"
	"gorm.io/gorm"
	"strconv"
)

// Video 视频模型
type Video struct {
	gorm.Model
	Title string
	Info  string
	Video string
	Cover string
	Uid   uint
}

// GetView 获取视频点击数
func (v *Video) GetView() uint64 {
	count, _ := cache.RedisClient.Get(cache.GetVideoViewKey(v.ID)).Uint64()
	return count
}

// View 浏览视频
func (v *Video) View() {
	//增加视频点击数
	cache.RedisClient.Incr(cache.GetVideoViewKey(v.ID))
	//在每日排行中增加该视频的点击数
	cache.RedisClient.ZIncrBy(cache.DailyRankKey, 1, strconv.Itoa(int(v.ID)))
}
