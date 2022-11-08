package video

import (
	"fmt"
	"giligili/cache"
	"giligili/model"
	"giligili/serializer"
	"strings"
)

type DailyRank struct {
}

func (s *DailyRank) Get() serializer.Response {
	var videoList []model.Video

	//从redis读出今日排行前10视频的id
	videos, err := cache.RedisClient.ZRevRange(cache.DailyRankKey, 0, 9).Result()
	if err != nil {
		return serializer.SerDbErr(err)
	}

	if len(videos) > 0 {
		order := fmt.Sprintf("FIELD(id, %s)", strings.Join(videos, ","))
		if err = model.DB.Where("id in (?)", videos).Order(order).Find(&videoList).Error; err != nil {
			return serializer.SerDbErr(err)
		}
	}

	return serializer.BuildVideoListResponse(videoList)
}
