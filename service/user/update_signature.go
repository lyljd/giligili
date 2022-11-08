package user

import (
	"giligili/model"
	"giligili/serializer"
	"giligili/service/common"
	"github.com/gin-gonic/gin"
)

type UpdateSignatureService struct {
	NewSignature string `form:"new_signature" json:"new_signature" binding:"max=50"`
}

func (s *UpdateSignatureService) UpdateSignature(c *gin.Context) serializer.Response {
	user := common.GetCurrentUser(c)
	if user.Signature == s.NewSignature {
		return serializer.CliParErr("新签名与原签名相同", nil)
	}

	user.Signature = s.NewSignature
	if err := model.DB.Save(&user).Error; err != nil {
		return serializer.SerDbErr(err)
	}

	return serializer.BuildUserResponse(*user)
}
