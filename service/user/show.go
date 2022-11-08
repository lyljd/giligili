package user

import (
	"giligili/model"
	"giligili/serializer"
	"strconv"
)

type ShowService struct {
}

func (s *ShowService) Show(id string) serializer.Response {
	if _, err := strconv.Atoi(id); err != nil {
		return serializer.CliParErr("id不合法", err)
	}
	var user model.User
	if err := model.DB.First(&user, id).Error; err != nil {
		return serializer.NotFound("用户不存在", err)
	}

	return serializer.BuildUserResponse(user)
}
