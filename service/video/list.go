package video

import (
	"giligili/model"
	"giligili/serializer"
)

type ListService struct {
}

func (s *ListService) List() serializer.Response {
	var videoList []model.Video
	if err := model.DB.Find(&videoList).Error; err != nil {
		return serializer.SerDbErr(err)
	}
	return serializer.BuildVideoListResponse(videoList)
}
