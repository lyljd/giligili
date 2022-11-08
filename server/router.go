package server

import (
	"giligili/api"
	"giligili/middleware"
	"giligili/service/common"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors())

	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", api.Ping)

		v1.POST("refresh-token", api.RefreshToken)

		v1.GET("resource/:type/:name", common.GetResource)

		v1.POST("register", api.Register)

		v1.POST("login", api.Login)

		v1.GET("user/:id", api.ShowUser) //用户详情

		v1.GET("video/:id", api.ShowVideo) //视频详情

		v1.GET("videos", api.ListVideo) //视频列表

		v1.GET("videos/rank/daily", api.GetVideoDailyRank) //视频每日排行

		v1.POST("videos/search", api.Search) //视频搜索

		// 需要登陆
		auth := v1.Group("")
		auth.Use(middleware.Auth())
		{
			auth.GET("me", api.Me)

			auth.POST("videos", api.CreateVideo) //视频投稿

			auth.PUT("video/:id", api.UpdateVideo) //视频更新

			auth.DELETE("video/:id", api.DeleteVideo) //视频删除

			auth.PUT("user/nickname", api.UpdateNickname)

			auth.PUT("user/signature", api.UpdateSignature)

			auth.PUT("user/password", api.UpdatePassword)

			auth.POST("user/avatar", api.UploadAvatar)
		}
	}

	return r
}
