package user

import (
	"errors"
	"giligili/model"
	"giligili/serializer"
	"giligili/service/common"
	"github.com/gin-gonic/gin"
)

type UpdatePasswordService struct {
	NewPassword        string `form:"new_password" json:"new_password" binding:"required,min=8,max=40"`
	NewPasswordConfirm string `form:"new_password_confirm" json:"new_password_confirm" binding:"required,min=8,max=40"`
}

func (s *UpdatePasswordService) UpdatePassword(c *gin.Context) serializer.Response {
	if s.NewPasswordConfirm != s.NewPassword {
		return serializer.CliParErr("两次密码输入不一致", nil)
	}

	user := common.GetCurrentUser(c)
	if user.CheckPassword(s.NewPassword) {
		return serializer.CliParErr("新密码与原密码相同", nil)
	}

	if err := user.SetPassword(s.NewPassword); err != nil {
		return serializer.SerErr(errors.New("密码加密失败"))
	}

	if err := model.DB.Save(&user).Error; err != nil {
		return serializer.SerDbErr(err)
	}

	return serializer.BuildUserResponse(*user)
}
