package repository

import (
	"sync"
	"time"
)

type Video struct {
	Id         int64     `gorm:"primarykey"json:"id"`
	AuthorId   int64     `gorm:"author_id"json:"authorId"`
	Title      string    `gorm:"title"json:"title"`
	PlayUrl    string    `gorm:"play_url"json:"playUrl"`
	CoverUrl   string    `gorm:"cover_url"json:"coverUrl"`
	CreateDate time.Time `gorm:"create_date"json:"CreateDate"`
	Status     bool      `gorm:"status"json:"status"`
}

func (Video) TableName() string {
	return "Video"
}

type VideoDao struct {
}

var videoDao *VideoDao
var videoOnce sync.Once

func NewVedioDaoInstance() *VideoDao {
	videoOnce.Do(
		func() {
			videoDao = &VideoDao{}
		})
	return videoDao
}

func (*VideoDao) PublishVideo(video *Video) error {
	video.CreateDate = time.Now()
	video.Status = true
	if err := db.Create(&video); err != nil {
		return err.Error
	}
	return nil
}
