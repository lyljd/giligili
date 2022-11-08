package api

import (
	"giligili/serializer"
	"giligili/service/common"
	"github.com/gin-gonic/gin"
)

// Ping 状态检查页面
func Ping(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Msg: "Pong",
	})
}

func RefreshToken(c *gin.Context) {
	var s common.RefreshTokenService
	if err := c.ShouldBind(&s); err == nil {
		res := s.RefreshToken(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.RefreshTokenErr("", err))
	}
}
