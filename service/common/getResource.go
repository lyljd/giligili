package common

import (
	"fmt"
	"giligili/model"
	"giligili/serializer"
	"giligili/util"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"time"
)

func GetResource(c *gin.Context) {
	typ, name := c.Param("type"), c.Param("name")
	expiresString, signature := c.Query("expires"), c.Query("signature")

	expires, err := strconv.ParseInt(expiresString, 10, 64)
	if err != nil {
		c.JSON(200, serializer.CliParErr("资源链接无效", err))
		return
	}
	now := time.Now().Unix()
	if expires < now {
		c.JSON(200, serializer.ResourceExpireErr(nil))
		return
	}

	trueSignature := util.SignatureResource(typ, name, expiresString)
	tsi := strings.LastIndex(trueSignature, "signature=") + 10
	trueSignature = trueSignature[tsi:]
	if signature != trueSignature {
		c.JSON(200, serializer.CliParErr("资源链接无效", nil))
		return
	}

	var obj any
	switch typ {
	case "video":
		obj = new(model.Video)
	case "cover":
		obj = new(model.Video)
	case "avatar":
		obj = new(model.User)
	}
	if err := model.DB.Where(typ+" = ?", name).First(obj).Error; err != nil {
		c.JSON(200, serializer.NotFound("资源已被删除", err))
		return
	}

	c.File(fmt.Sprintf("./resource/%s/%s", typ, name))
}
