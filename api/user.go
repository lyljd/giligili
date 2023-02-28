package api

import (
	"errors"
	"giligili/serializer"
	"giligili/service/common"
	"giligili/service/user"
	"github.com/gin-gonic/gin"
	"os"
)

func Register(c *gin.Context) {
	if os.Getenv("OPENAPI_REGISTER") != "true" {
		c.JSON(200, serializer.NoPower("服务器已关闭用户注册接口", nil))
		return
	}
	var s user.RegisterService
	if err := c.ShouldBind(&s); err == nil {
		res := s.Register()
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.CliParErr("", err))
	}
}

func Login(c *gin.Context) {
	var s user.LoginService
	if err := c.ShouldBind(&s); err == nil {
		res := s.Login(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.CliParErr("", err))
	}
}

func Me(c *gin.Context) {
	user := common.GetCurrentUser(c)
	if user != nil {
		c.JSON(200, serializer.BuildUserResponse(*user))
		return
	}
	c.JSON(200, serializer.SerErr(errors.New("获取当前用户失败，指针为空")))
}

func ShowUser(c *gin.Context) {
	var s user.ShowService
	if err := c.ShouldBind(&s); err == nil {
		res := s.Show(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.CliParErr("", err))
	}
}

func UpdateNickname(c *gin.Context) {
	var s user.UpdateNicknameService
	if err := c.ShouldBind(&s); err == nil {
		res := s.UpdateNickname(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, err)
	}
}

func UpdateSignature(c *gin.Context) {
	var s user.UpdateSignatureService
	if err := c.ShouldBind(&s); err == nil {
		res := s.UpdateSignature(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.CliParErr("", err))
	}
}

func UpdatePassword(c *gin.Context) {
	var s user.UpdatePasswordService
	if err := c.ShouldBind(&s); err == nil {
		res := s.UpdatePassword(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.CliParErr("", err))
	}
}

func UploadAvatar(c *gin.Context) {
	var u user.UploadAvatarService
	if err := c.ShouldBind(&u); err == nil {
		res := u.UploadAvatar(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.CliParErr("", err))
	}
}
