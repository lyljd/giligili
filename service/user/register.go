package user

import (
	"errors"
	"giligili/model"
	"giligili/serializer"
)

type RegisterService struct {
	Nickname        string `form:"nickname" json:"nickname" binding:"required,min=2,max=30"`
	Username        string `form:"username" json:"username" binding:"required,min=5,max=30"`
	Password        string `form:"password" json:"password" binding:"required,min=8,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=8,max=40"`
}

func (s *RegisterService) Register() serializer.Response {
	//验证输入合法性
	if err := s.valid(); err != nil {
		return *err
	}

	user := model.User{
		Nickname: s.Nickname,
		Username: s.Username,
	}

	//加密密码
	if err := user.SetPassword(s.Password); err != nil {
		return serializer.SerErr(errors.New("密码加密失败"))
	}

	//创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.SerDbErr(err)
	}

	return serializer.BuildUserResponse(user)
}

// 验证输入合法性
func (s *RegisterService) valid() *serializer.Response {
	if s.PasswordConfirm != s.Password {
		ret := serializer.CliParErr("两次密码输入不一致", nil)
		return &ret
	}

	count := int64(0)
	model.DB.Model(&model.User{}).Where("nickname = ?", s.Nickname).Count(&count)
	if count > 0 {
		ret := serializer.CliParErr("昵称已被使用", nil)
		return &ret
	}

	count = 0
	model.DB.Model(&model.User{}).Where("username = ?", s.Username).Count(&count)
	if count > 0 {
		ret := serializer.CliParErr("用户名已被注册", nil)
		return &ret
	}

	return nil
}
