package user

import (
	"giligili/model"
	"giligili/serializer"
	"giligili/service/common"
	"github.com/gin-gonic/gin"
)

type UpdateNicknameService struct {
	NewNickname string `form:"new_nickname" json:"new_nickname" binding:"required,min=2,max=30"`
}

func (s *UpdateNicknameService) UpdateNickname(c *gin.Context) serializer.Response {
	user := common.GetCurrentUser(c)
	if user.Nickname == s.NewNickname {
		return serializer.CliParErr("新昵称与原昵称相同", nil)
	}

	count := int64(0)
	model.DB.Model(&model.User{}).Where("nickname = ?", s.NewNickname).Count(&count)
	if count > 0 {
		return serializer.CliParErr("昵称已被使用", nil)
	}

	user.Nickname = s.NewNickname
	if err := model.DB.Save(&user).Error; err != nil {
		return serializer.SerDbErr(err)
	}

	return serializer.BuildUserResponse(*user)
}
