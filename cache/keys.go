package cache

import (
	"fmt"
	"strconv"
)

const (
	DailyRankKey = "rank:daily" //每日排行
)

// GetVideoViewKey 根据视频id获取在redis中的键名
func GetVideoViewKey(id uint) string {
	return fmt.Sprintf("view:video:%s", strconv.Itoa(int(id)))
}
