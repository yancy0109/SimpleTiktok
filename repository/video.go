package repository

import (
	"errors"
	"gorm.io/gorm"
	"sync"
	"time"
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
