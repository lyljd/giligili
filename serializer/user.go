package serializer

import (
	"giligili/auth"
	"giligili/model"
	"giligili/util"
)

// User 用户序列化器
type User struct {
	ID         uint   `json:"id"`
	CreatedAt  int64  `json:"created_at"`
	Nickname   string `json:"nickname"`
	Avatar     string `json:"avatar"`
	Signature  string `json:"signature"`
	IpLocation string `json:"ip_location"`
}

// BuildUser 序列化用户
func BuildUser(user model.User) User {
	return User{
		ID:         user.ID,
		CreatedAt:  user.CreatedAt.Unix(),
		Nickname:   user.Nickname,
		Avatar:     util.SignatureResource("avatar", user.Avatar, ""),
		Signature:  user.Signature,
		IpLocation: user.IpLocation,
	}
}

// BuildUserResponse 序列化用户响应
func BuildUserResponse(user model.User) Response {
	return Response{
		Data: BuildUser(user),
	}
}

// BuildUserLoginResponse 序列化用户登陆响应
func BuildUserLoginResponse(user model.User) Response {
	return Response{
		Data: struct {
			User
			Token        string `json:"token"`
			RefreshToken string `json:"refresh_token"`
		}{BuildUser(user),
			auth.NewToken(user.ID, auth.TypeToken),
			auth.NewToken(user.ID, auth.TypeRefreshToken)}}
}
