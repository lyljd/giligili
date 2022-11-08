package video

import (
	"giligili/cache"
	"giligili/model"
	"giligili/serializer"
	"giligili/service/common"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Delete(id string, c *gin.Context) serializer.Response {
	var iid int
	var err error
	if iid, err = strconv.Atoi(id); err != nil {
		return serializer.CliParErr("id不合法", err)
	}

	var video model.Video
	if err = model.DB.First(&video, id).Error; err != nil {
		return serializer.NotFound("视频不存在", err)
	}

	if common.GetCurrentUser(c).ID != video.Uid {
		return serializer.NoPower("该视频不属于你", nil)
	}

	if err = model.DB.Delete(&video).Error; err != nil {
		return serializer.SerDbErr(err)
	}

	cache.RedisClient.Unlink(cache.GetVideoViewKey(uint(iid)))
	cache.RedisClient.ZRem(cache.DailyRankKey, strconv.Itoa(iid))

	// gorm数据库删除是软删除，为了后续可以恢复，所以不用删文件
	/*if err = os.Remove("./resource/video/" + video.Video); err != nil {
		util.Log().Info("删除视频 " + video.Video + " 失败，请手动删除")
	}
	if err = os.Remove("./resource/cover/" + video.Cover); err != nil {
		util.Log().Info("删除封面 " + video.Cover + " 失败，请手动删除")
	}*/

	return serializer.Response{}
}
