package tasks

import "giligili/cache"

// ClearDailyRank 清除redis中的每日排行
func ClearDailyRank() error {
	return cache.RedisClient.Unlink(cache.DailyRankKey).Err()
}
