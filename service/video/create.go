package video

import (
	"giligili/model"
	"giligili/serializer"
	"giligili/service/common"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"mime/multipart"
	"strings"
)

type CreateService struct {
	Title string                `form:"title" binding:"required,min=2,max=30"`
	Info  string                `form:"info" binding:"max=255"`
	Video *multipart.FileHeader `form:"video" binding:"required"`
	Cover *multipart.FileHeader `form:"cover" binding:"required"`
}

func (s *CreateService) Create(c *gin.Context) serializer.Response {
	if ct := s.Cover.Header.Get("Content-Type"); !common.AllowImageContType[ct] {
		return serializer.CliParErr("封面类型 "+ct+" 不合法", nil)
	}
	if ct := s.Video.Header.Get("Content-Type"); ct != "video/mp4" {
		return serializer.CliParErr("视频类型 "+ct+" 不合法", nil)
	}

	if float64(s.Cover.Size)/1024/1024 > 2 {
		return serializer.CliParErr("封面大小不能超过2MB", nil)
	}
	if float64(s.Video.Size)/1024/1024 > 1024 {
		return serializer.CliParErr("视频大小不能超过1GB", nil)
	}

	filename := uuid.NewV4().String()
	cn := filename + s.Cover.Filename[strings.LastIndex(s.Cover.Filename, "."):]
	vn := filename + s.Video.Filename[strings.LastIndex(s.Video.Filename, "."):]

	if err := c.SaveUploadedFile(s.Cover, "./resource/cover/"+cn); err != nil {
		return serializer.SerErr(err)
	}
	if err := c.SaveUploadedFile(s.Video, "./resource/video/"+vn); err != nil {
		return serializer.SerErr(err)
	}

	video := model.Video{
		Title: s.Title,
		Info:  s.Info,
		Video: vn,
		Cover: cn,
		Uid:   common.GetCurrentUser(c).ID,
	}
	if err := model.DB.Create(&video).Error; err != nil {
		return serializer.SerDbErr(err)
	}
	return serializer.BuildVideoResponse(video)
}
