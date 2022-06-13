package service

import "github.com/yancy0109/SimpleTiktok/repository"

type VideoListModal struct {
	StatusCode int64   `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string `json:"status_msg"`  // 返回状态描述
	VideoList  []Video `json:"video_list"`  // 用户发布的视频列表
}
type VideoListService struct {
}

var videoListDao = repository.NewVideoListInstance()

func (*VideoListService) GetVideoList(tokenId int64, userId int64) *VideoListModal {
	var videoListModal VideoListModal
	//查询他名下的所有视频列表
	videos, _ := videoListDao.VideoListForUserId(userId)
	//查询这些视频的信息
	//关系信息应当是与登录的账号的而不是和查询的账号的
	videoList, code := GetVideoInformation(videos, tokenId)
	videoListModal.StatusCode = code
	videoListModal.VideoList = videoList
	var msg string
	if tokenId == -1 {
		msg = "用户未登录"
	} else {
		msg = "正常"
	}
	videoListModal.StatusMsg = &msg
	return &videoListModal
}
