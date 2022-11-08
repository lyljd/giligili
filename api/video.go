package api

import (
	"giligili/serializer"
	"giligili/service/video"
	"github.com/gin-gonic/gin"
)

func CreateVideo(c *gin.Context) {
	var s video.CreateService
	if err := c.ShouldBind(&s); err == nil {
		res := s.Create(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.CliParErr("", err))
	}
}

func ShowVideo(c *gin.Context) {
	var s video.ShowService
	if err := c.ShouldBind(&s); err == nil {
		res := s.Show(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.CliParErr("", err))
	}
}

func ListVideo(c *gin.Context) {
	var s video.ListService
	if err := c.ShouldBind(&s); err == nil {
		res := s.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.CliParErr("", err))
	}
}

func UpdateVideo(c *gin.Context) {
	var s video.UpdateService
	if err := c.ShouldBind(&s); err == nil {
		res := s.Update(c.Param("id"), c)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.CliParErr("", err))
	}
}

func DeleteVideo(c *gin.Context) {
	res := video.Delete(c.Param("id"), c)
	c.JSON(200, res)
}

func GetVideoDailyRank(c *gin.Context) {
	var s video.DailyRank
	if err := c.ShouldBind(&s); err == nil {
		res := s.Get()
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.CliParErr("", err))
	}
}

func Search(c *gin.Context) {
	var s video.SearchService
	if err := c.ShouldBind(&s); err == nil {
		res := s.Search()
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.CliParErr("", err))
	}
}
