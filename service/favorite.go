package service

import (
	"github.com/yancy0109/SimpleTiktok/repository"
	"strconv"
)

type FavoriteService struct {
}
type FavoriteList struct {
	StatusCode int64   `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string `json:"status_msg"`  // 返回状态描述
	VideoList  []Video `json:"video_list"`  // 视频列表
}

var favoriteDao = repository.NewfavoriteDaoInstance()

func (f FavoriteService) Update(userid int64, videoid string, actionType string) error {
	videoId, err := strconv.ParseInt(videoid, 10, 64)
	if err != nil {
		return err
	}
	status, err := strconv.ParseInt(actionType, 10, 64)
	if err != nil {
		return err
	}
	favoriteRecord := repository.FavoriteRecord{
		Userid:  userid,
		Videoid: videoId,
		Status:  status,
	}
	return favoriteDao.UpdateStatus(favoriteRecord)
}
func (f FavoriteService) FavoriteList(userId int64) *FavoriteList {
	var favorite FavoriteList
	videos, _ := favoriteDao.GetList(userId)
	//查询视频的信息
	videoList, code := GetVideoInformation(videos, userId)
	favorite.StatusCode = code
	favorite.VideoList = videoList
	var msg string
	if userId == -1 {
		msg = "用户未登录"
	} else {
		msg = "正常"
	}
	favorite.StatusMsg = &msg
	return &favorite
}
