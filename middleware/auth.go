package middleware

import (
	"errors"
	"giligili/auth"
	"giligili/model"
	"giligili/serializer"
	"github.com/gin-gonic/gin"
)

// Auth 验证登陆
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Token")
		if token == "" {
			c.JSON(200, serializer.NoLogin("", errors.New("没有获取到Token")))
			c.Abort()
			return
		}

		id, err := auth.CheckToken(token, auth.TypeToken)
		if err != nil {
			msg := ""
			if err.Error() == "token过期" {
				msg = "身份已过期，请登录"
			}
			c.JSON(401, serializer.NoLogin(msg, err))
			c.Abort()
			return
		}

		user := model.GetUser(id)
		c.Set("user", &user)
		c.Next()
	}
}
