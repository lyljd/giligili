package serializer

import (
	"giligili/model"
	"giligili/util"
)

// Video 视频序列化器
type Video struct {
	ID        uint   `json:"id"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	Title     string `json:"title"`
	Info      string `json:"info"`
	Video     string `json:"video"`
	Cover     string `json:"cover"`
	View      uint64 `json:"view"`

	UID        uint   `json:"u_id"`
	UNickname  string `json:"u_nickname"`
	UAvatar    string `json:"u_avatar"`
	USignature string `json:"u_signature"`
}

// BuildVideo 序列化视频
func BuildVideo(video model.Video) Video {
	user := model.GetUser(video.Uid)

	v := Video{
		ID:        video.ID,
		CreatedAt: video.CreatedAt.Unix(),
		Title:     video.Title,
		Info:      video.Info,
		Video:     util.SignatureResource("video", video.Video, ""),
		Cover:     util.SignatureResource("cover", video.Cover, ""),
		View:      video.GetView(),

		UID:        user.ID,
		UNickname:  user.Nickname,
		UAvatar:    util.SignatureResource("avatar", user.Avatar, ""),
		USignature: user.Signature,
	}
	if ua := video.UpdatedAt.Unix(); ua > v.CreatedAt {
		v.UpdatedAt = ua
	}
	return v
}

// BuildVideoResponse 序列化视频响应
func BuildVideoResponse(video model.Video) Response {
	return Response{
		Data: BuildVideo(video),
	}
}

// BuildVideoListResponse 序列化视频列表响应
func BuildVideoListResponse(videoList []model.Video) Response {
	var videoListResp []Video
	for _, video := range videoList {
		videoListResp = append(videoListResp, BuildVideo(video))
	}
	return Response{
		Data: videoListResp,
	}
}
