package video

import (
	"giligili/model"
	"giligili/serializer"
)

type SearchService struct {
	Keyword string `form:"keyword" json:"keyword" binding:"required,max=255"`
}

func (s *SearchService) Search() serializer.Response {
	var v []model.Video
	model.DB.Where("title like ? or info like ?", "%"+s.Keyword+"%", s.Keyword).Find(&v)
	return serializer.BuildVideoListResponse(v)
}
