package service

import (
	"github.com/yancy0109/SimpleTiktok/repository"
	"sync"
)

type videoFeed struct {
	NextTime   *int64  `json:"next_time"`   // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
	StatusCode int64   `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string `json:"status_msg"`  // 返回状态描述
	VideoList  []Video `json:"video_list"`  // 视频列表
}

// Video
type Video struct {
	Author        User   `json:"author"`         // 视频作者信息
	CommentCount  int64  `json:"comment_count"`  // 视频的评论总数
	CoverURL      string `json:"cover_url"`      // 视频封面地址
	FavoriteCount int64  `json:"favorite_count"` // 视频的点赞总数
	ID            int64  `json:"id"`             // 视频唯一标识
	IsFavorite    bool   `json:"is_favorite"`    // true-已点赞，false-未点赞
	PlayURL       string `json:"play_url"`       // 视频播放地址
	Title         string `json:"title"`          // 视频标题
}

// 视频作者信息
//
// User
type User struct {
	FollowCount   int64  `json:"follow_count"`   // 关注总数
	FollowerCount int64  `json:"follower_count"` // 粉丝总数
	ID            int64  `json:"id"`             // 用户id
	IsFollow      bool   `json:"is_follow"`      // true-已关注，false-未关注
	Name          string `json:"name"`           // 用户名称
}
type VideoService struct {
}

var videoListDao = repository.NewVideoListInstance()

func (*VideoService) GetVideoFeed(latestTime int, userId int64) *videoFeed {
	var videoFeed videoFeed
	videos, _ := videoListDao.VideoList(latestTime)
	videoList := make([]Video, len(videos))
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(videos))
	for index, video := range videos {
		//多线程查询video的信息们
		go func(video *repository.Video, userId int64, index int) {
			author, _ := videoListDao.AuthorInformation(video.AuthorId, userId)
			countCount, _ := videoListDao.VideoCommentCount(video.Id)
			IsFavorite, _ := videoListDao.FavoriteStatus(video.Id, userId)
			favoriteCount, _ := videoListDao.VideoFavoriteCount(video.Id)
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
				FavoriteCount: int64(favoriteCount),
				ID:            video.Id,
				IsFavorite:    IsFavorite,
				PlayURL:       video.PlayUrl,
				Title:         video.Title,
			}
		}(&video, userId, index)
	}
	waitGroup.Wait()
	return &videoFeed
}
