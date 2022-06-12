package repository

import (
	"errors"
	"sync"
	"time"

	"gorm.io/gorm"
)

type Video struct {
	Id         int64     `gorm:"primarykey" json:"id"`
	AuthorId   int64     `gorm:"author_id"  json:"authorId"`
	Title      string    `gorm:"title"      json:"title"`
	PlayUrl    string    `gorm:"play_url"   json:"playUrl"`
	CoverUrl   string    `gorm:"cover_url"  json:"coverUrl"`
	CreateDate time.Time `gorm:"create_date" json:"CreateDate"`
	Status     bool      `gorm:"status"     json:"status"`
}

func (Video) TableName() string {
	return "video"
}

type VideoDao struct {
}

var videoDao *VideoDao
var videoOnce sync.Once

func NewVideoDaoInstance() *VideoDao {
	videoOnce.Do(
		func() {
			videoDao = &VideoDao{}
		})
	return videoDao
}

//发布视频
func (*VideoDao) PublishVideo(video *Video) error {
	video.CreateDate = time.Now()
	video.Status = true
	if err := db.Create(&video); err != nil {
		return err.Error
	}
	return nil
}

//检验是否存在视频
func (*VideoDao) IsExistVideo(videoId int64) error {
	result := db.First(&Video{}, "id = ?", videoId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	return nil
}

//查询user1的信息  并返回是否关注user2
func (*VideoDao) AuthorInformation(userId1 int64, userId2 int64) (*Author, error) {
	var author *Author
	author = new(Author)
	author.Id = userId1
	//根据userId1查询作者的名称信息
	result := db.Table("user").Select("user_name").Where("id = ?", userId1).Limit(1).Find(&author.UserName)
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
	//根据userId1和userId2查询是否关注了
	resultIsFollow := db.Table("follow").Select("is_del <> 1").Where("follow = ? and be_follow = ?", userId1, userId2).Limit(1).Find(&author.IsFollow)
	if resultIsFollow.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			author.IsFollow = false
		} else {
			return author, resultIsFollow.Error
		}
	}
	//查询userId1的粉丝数
	db.Table("follow").Select("count(*)").Where("is_del <> 1 and be_follow = ?", userId1).Limit(1).Find(&author.FollowerCount)
	//查询userId1的关注数量
	db.Table("follow").Select("count(*)").Where("is_del <> 1 and follow = ?", userId1).Limit(1).Find(&author.FollowCount)
	//返回
	return author, nil
}

//查询user信息
func (*VideoDao) GetUserInformation(userId int64) (*Author, error) {
	var author *Author
	author = new(Author)
	author.Id = userId
	//根据userId1查询作者的名称信息
	result := db.Table("user").Select("user_name").Where("id = ?", userId).Limit(1).Find(&author.UserName)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			author.UserName = "用户已注销"
			author.FollowCount = 0
			author.FollowerCount = 0
			author.IsFollow = false
			return author, nil
		}
		return nil, result.Error
	}
	author.IsFollow = false
	//查询userId的粉丝数
	db.Table("follow").Select("count(*)").Where("is_del <> 1 and be_follow = ?", userId).Limit(1).Find(&author.FollowerCount)
	//查询userId1的关注数量
	db.Table("follow").Select("count(*)").Where("is_del <> 1 and follow = ?", userId).Limit(1).Find(&author.FollowCount)
	//返回
	return author, nil
}
