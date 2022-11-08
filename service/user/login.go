package user

import (
	"giligili/model"
	"giligili/serializer"
	"giligili/service/common"
	"github.com/gin-gonic/gin"
)

type LoginService struct {
	Username string `form:"username" json:"username" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

func (s *LoginService) Login(c *gin.Context) serializer.Response {
	var user model.User

	if err := model.DB.Where("username = ?", s.Username).First(&user).Error; err != nil {
		return serializer.CliParErr("用户名或密码错误", err)
	}

	if user.CheckPassword(s.Password) == false {
		return serializer.CliParErr("用户名或密码错误", nil)
	}

	common.SetUserIpLocation(user.ID, c.ClientIP())

	return serializer.BuildUserLoginResponse(user)
}
