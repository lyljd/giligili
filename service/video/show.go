package video

import (
	"giligili/model"
	"giligili/serializer"
	"strconv"
)

type ShowService struct {
}

func (s *ShowService) Show(id string) serializer.Response {
	if _, err := strconv.Atoi(id); err != nil {
		return serializer.CliParErr("id不合法", err)
	}
	var video model.Video
	if err := model.DB.First(&video, id).Error; err != nil {
		return serializer.NotFound("视频不存在", err)
	}

	video.View() //增加点击数

	return serializer.BuildVideoResponse(video)
}
