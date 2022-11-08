package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	gorm.Model
	Username   string
	Nickname   string
	Password   string
	Avatar     string
	Signature  string
	IpLocation string
}

const (
	// PassWordCost 密码加密难度(加密次数，相当于对加密的密码加密)
	PassWordCost = 12
)

func GetUser(id uint) (u User) {
	DB.First(&u, id)
	return
}

func (u *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
