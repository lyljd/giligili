package common

import (
	"giligili/auth"
	"giligili/serializer"
	"github.com/gin-gonic/gin"
)

type RefreshTokenService struct {
	Token string `form:"token" json:"token" binding:"required"`
}

func (s *RefreshTokenService) RefreshToken(c *gin.Context) serializer.Response {
	id, err := auth.CheckToken(s.Token, auth.TypeRefreshToken)
	if err != nil {
		msg := ""
		if err.Error() == "token过期" {
			msg = "身份已过期，请登录"
		}
		return serializer.RefreshTokenErr(msg, err)
	}

	SetUserIpLocation(id, c.ClientIP())

	return serializer.Response{
		Data: struct {
			Token string `json:"token"`
		}{auth.NewToken(id, auth.TypeToken)},
	}
}
