package common

import (
	"encoding/json"
	"giligili/model"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

// AllowImageContType 图片上传时所允许的类型
var AllowImageContType = map[string]bool{
	"image/png":  true,
	"image/jpg":  true,
	"image/jpeg": true,
}

// GetCurrentUser 获取当前用户(通过了Auth的路由才能用)
func GetCurrentUser(c *gin.Context) *model.User {
	user, _ := c.Get("user")
	return user.(*model.User)
}

// GetIpLocation 获取Ip所在位置
func GetIpLocation(ip string) (location string) {
	location = "未知"

	res, err := http.Get("https://ip.useragentinfo.com/json?ip=" + ip)
	if err != nil {
		return
	}

	body, err := io.ReadAll(res.Body)
	_ = res.Body.Close()
	if err != nil {
		return
	}

	il := struct {
		Country  string `json:"country"`
		Province string `json:"province"`
	}{}
	_ = json.Unmarshal(body, &il)
	if il.Country == "中国" {
		if il.Province != "" {
			location = il.Province[:len(il.Province)-3]
		}
	} else if il.Country != "保留地址" && il.Country != "" {
		location = il.Country
	}
	return
}

// SetUserIpLocation 设置用户Ip所在位置
func SetUserIpLocation(id uint, ip string) {
	user := model.GetUser(id)
	model.DB.Model(&user).Update("IpLocation", GetIpLocation(ip))
}
