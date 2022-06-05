package service

import (
	"errors"
	"github.com/yancy0109/SimpleTiktok/repository"
	"gorm.io/gorm"
	"strconv"
	"sync"
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
	videoList := make([]Video, len(videos))
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(videos))
	for index, indexVideo := range videos {
		//多线程查询video的信息们
		indexVideo := indexVideo
		//此处需要对indexVideo进行阴影处理，否则会产生一个数据竞争
		go func(video *repository.Video, userId int64, index int) {
			defer waitGroup.Done()
			author, err1 := videoListDao.AuthorInformation(video.AuthorId, userId)
			if err1 != nil {
				if errors.Is(err1, gorm.ErrRecordNotFound) {
					author.UserName = "用户已注销"
				} else {
					favorite.StatusCode++
				}
			}
			countCount, err2 := videoListDao.VideoCommentCount(video.Id)
			if err2 != nil {
				if errors.Is(err2, gorm.ErrRecordNotFound) {
					countCount = 0
				} else {
					favorite.StatusCode++
				}
			}
			IsFavorite, err3 := favoriteDao.Favorite(video.Id, userId)
			if err3 != nil {
				if errors.Is(err3, gorm.ErrRecordNotFound) {
					IsFavorite = false
				} else {
					favorite.StatusCode++
				}
			}
			favoriteCount, err4 := favoriteDao.GetCount(video.Id)
			if err4 != nil {
				if errors.Is(err4, gorm.ErrRecordNotFound) {
					favoriteCount = 0
				} else {
					favorite.StatusCode++
				}
			}
			videoList[index] = Video{
				Author: User{
					FollowCount:   author.FollowCount,
					FollowerCount: author.FollowerCount,
					ID:            author.Id,
					IsFollow:      author.IsFollow,
					Name:          author.UserName,
				},
				CommentCount:  int64(countCount),
				CoverURL:      video.CoverUrl,
				FavoriteCount: favoriteCount,
				ID:            video.Id,
				IsFavorite:    IsFavorite,
				PlayURL:       video.PlayUrl,
				Title:         video.Title,
			}
		}(&indexVideo, userId, index)
	}
	waitGroup.Wait()
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
