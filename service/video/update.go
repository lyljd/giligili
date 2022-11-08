package video

import (
	"giligili/model"
	"giligili/serializer"
	"giligili/service/common"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"strconv"
)

type UpdateService struct {
	Title string                `form:"title" binding:"required,min=2,max=30"`
	Info  string                `form:"info" binding:"max=255"`
	Video *multipart.FileHeader `form:"video"`
	Cover *multipart.FileHeader `form:"cover"`
}

func (s *UpdateService) Update(id string, c *gin.Context) serializer.Response {
	if _, err := strconv.Atoi(id); err != nil {
		return serializer.CliParErr("id不合法", err)
	}

	var video model.Video
	if err := model.DB.First(&video, id).Error; err != nil {
		return serializer.NotFound("视频不存在", err)
	}

	if common.GetCurrentUser(c).ID != video.Uid {
		return serializer.NoPower("该视频不属于你", nil)
	}

	if s.Cover != nil {
		if ct := s.Cover.Header.Get("Content-Type"); !common.AllowImageContType[ct] {
			return serializer.CliParErr("封面类型 "+ct+" 不合法", nil)
		}
		if float64(s.Cover.Size)/1024/1024 > 2 {
			return serializer.CliParErr("封面大小不能超过2MB", nil)
		}
	}
	if s.Video != nil {
		if ct := s.Video.Header.Get("Content-Type"); ct != "video/mp4" {
			return serializer.CliParErr("视频类型 "+ct+" 不合法", nil)
		}
		if float64(s.Video.Size)/1024/1024 > 1024 {
			return serializer.CliParErr("视频大小不能超过1GB", nil)
		}
	}

	if s.Cover != nil {
		if err := c.SaveUploadedFile(s.Cover, "./resource/cover/"+video.Cover); err != nil {
			return serializer.SerErr(err)
		}
	}
	if s.Video != nil {
		if err := c.SaveUploadedFile(s.Video, "./resource/video/"+video.Video); err != nil {
			return serializer.SerErr(err)
		}
	}

	video.Title, video.Info = s.Title, s.Info
	if err := model.DB.Save(&video).Error; err != nil {
		return serializer.SerDbErr(err)
	}
	return serializer.BuildVideoResponse(video)
}
