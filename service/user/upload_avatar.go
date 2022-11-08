package user

import (
	"giligili/model"
	"giligili/serializer"
	"giligili/service/common"
	"giligili/util"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"mime/multipart"
	"os"
	"strings"
)

type UploadAvatarService struct {
	Avatar *multipart.FileHeader `form:"avatar" binding:"required"`
}

func (s *UploadAvatarService) UploadAvatar(c *gin.Context) serializer.Response {
	if ct := s.Avatar.Header.Get("Content-Type"); !common.AllowImageContType[ct] {
		return serializer.CliParErr("头像类型 "+ct+" 不合法", nil)
	}

	if float64(s.Avatar.Size)/1024/1024 > 2 {
		return serializer.CliParErr("头像大小不能超过2MB", nil)
	}

	user := common.GetCurrentUser(c)
	if user.Avatar == "" {
		an := uuid.NewV4().String() + s.Avatar.Filename[strings.LastIndex(s.Avatar.Filename, "."):]
		if err := c.SaveUploadedFile(s.Avatar, "./resource/avatar/"+an); err != nil {
			return serializer.SerErr(err)
		}
		user.Avatar = an
		if err := model.DB.Save(&user).Error; err != nil {
			if err := os.Remove("./resource/avatar/" + an); err != nil {
				util.Log().Info("删除头像 " + an + " 失败，请手动删除")
			}
			return serializer.SerDbErr(err)
		}
	} else {
		an := user.Avatar
		if err := c.SaveUploadedFile(s.Avatar, "./resource/avatar/"+an); err != nil {
			return serializer.SerErr(err)
		}
	}

	return serializer.BuildUserResponse(*user)
}
