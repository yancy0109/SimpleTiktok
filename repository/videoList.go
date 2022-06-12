package repository

import (
	"errors"
	"sync"

	"gorm.io/gorm"
)

type VideoListDao struct {
}

// Author 视频作者信息
// User
type Author struct {
	Id            int64  `gorm:"id" json:"id"`
	UserName      string `gorm:"user_name" json:"name"`
	FollowCount   int64  `gorm:"Follow_Count" json:"follow_count"`
	FollowerCount int64  `gorm:"Follower_Count" json:"follower_count"`
	IsFollow      bool   `gorm:"Is_Follow" json:"is_follow"`
}

//单例模式
var videoListDao *VideoListDao
var videoListOnce sync.Once

func NewVideoListInstance() *VideoListDao {
	videoListOnce.Do(
		func() {
			videoListDao = &VideoListDao{}
		})
	return videoListDao
}

func (*VideoListDao) VideoList(latestTime int64) ([]Video, error) {

	listSize := 10
	//latestTime *= 1000
	//根据latestTime查询符合条件的视频及其信息
	var videoList []Video
	//这里要将数据库里面的时间格式转化为时间戳格式再进行比较
	result := db.Table("video").Where("UNIX_TIMESTAMP(create_date) < ? and status <> 0", latestTime*1000).Order("create_date desc").Limit(listSize).Find(&videoList)
	if result.Error != nil {
		return nil, result.Error
	}
	//返回
	return videoList, nil
}
func (*VideoListDao) VideoListForUserId(userId int64) ([]Video, error) {
	var videoList []Video
	//通过userId进行查询，查询结果按照时间倒序
	result := db.Table("video").Where("author_id = ? and status <> 0", userId).Order("create_date desc").Find(&videoList)
	if result.Error != nil {
		return nil, result.Error
	}
	//返回
	return videoList, nil
}
func (*VideoListDao) AuthorInformation(AuthorId int64, userId int64) (*Author, error) {
	var author *Author
	author = new(Author)
	author.Id = AuthorId
	//根据authorId查询作者的名称信息
	result := db.Table("user").Select("user_name").Where("id = ?", userId).Limit(1).Find(&author.UserName)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			author.UserName = "用户已注销"
			author.FollowCount = 0
			author.FollowerCount = 0
			author.IsFollow = false
			return author, nil
		}
		return nil, result.Error
	}
	//根据userId和authorId查询是否关注了
	resultIsFollow := db.Table("follow").Select("is_del <> 1").Where("follow = ? and be_follow = ?", userId, AuthorId).Limit(1).Find(&author.IsFollow)
	if resultIsFollow.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			author.IsFollow = false
		} else {
			return author, resultIsFollow.Error
		}
	}
	//查询author的粉丝数
	resultFollowerCount := db.Table("follow").Select("count(*)").Where("is_del <> 1 and be_follow = ?", AuthorId).Limit(1).Find(&author.FollowerCount)
	if resultFollowerCount.Error != nil {
		return author, resultFollowerCount.Error
	}
	//查询author的关注数量
	resultFollowCount := db.Table("follow").Select("count(*)").Where("is_del <> 1 and follow = ?", AuthorId).Limit(1).Find(&author.FollowCount)
	if resultFollowCount.Error != nil {
		return author, resultFollowCount.Error
	}

	//返回
	return author, nil
}

// FavoriteStatus 查询是否建立了喜欢
func (*VideoListDao) FavoriteStatus(videoId int64, userId int64) (bool, error) {
	status := 0
	result := db.Table("video_favorite").Select("status").Where("video_id = ? and user_id = ?", videoId, userId).Limit(1).Find(&status)
	if result.Error != nil {
		return false, result.Error
	}
	//返回
	return status != 0, nil
}

// VideoFavoriteCount 视频的喜欢数
func (*VideoListDao) VideoFavoriteCount(videoId int64) (int, error) {
	var count int
	result := db.Table("video_favorite").Select("count(*)").Where("video_id = ? and status <> 0", videoId).Limit(1).Find(&count)
	if result.Error != nil {
		return 0, nil
	}
	return count, nil
}

// VideoCommentCount 视频的评论数
func (*VideoListDao) VideoCommentCount(videoId int64) (int, error) {
	var count int
	result := db.Table("comment").Select("count(*)").Where("video_id = ? and status <> 0", videoId).Limit(1).Find(&count)
	if result.Error != nil {
		return 0, nil
	}
	return count, nil
}
