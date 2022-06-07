package repository

import (
	"sync"
	"time"
)

type CommentRecord struct {
	//gorm.Model
	Id         int64     `gorm:"column:id" `
	VideoId    int64     `gorm:"column:video_id" `
	AuthorId   int64     `gorm:"column:author_id" `
	Content    string    `gorm:"column:content" `
	Status     int64     `gorm:"column:status" `
	CreateDate time.Time `gorm:"column:create_date" `
}
type CommentDao struct {
}

var commentDao *CommentDao
var commentOnce sync.Once

func NewCommentDaoInstance() *CommentDao {
	commentOnce.Do(
		func() {
			commentDao = &CommentDao{}
		})
	return commentDao
}
func (CommentRecord) TableName() string {
	return "comment"
}

func (*CommentDao) GetList(commentRecord CommentRecord) []CommentRecord {
	var results []CommentRecord
	db.Table("comment").Where("video_id = ? AND Status = ?", commentRecord.AuthorId, 1).Find(&results) // (*sql.Row)
	return results
}

func (*CommentDao) GetCount(commentRecord CommentRecord) int64 {
	var count int64
	db.Model(&commentRecord).Where("video_id = ? AND Status = ? ", commentRecord.VideoId, 1).Count(&count)
	return count
}

func (*CommentDao) UpdateStatus(commentRecord CommentRecord) int64 {
	commentRecord.CreateDate = time.Now()
	db.Model(&commentRecord).Create(&commentRecord) //记录-创建
	return commentRecord.Id
}

func (*CommentDao) ChangeStatus(commentRecord CommentRecord) error {
	if err := db.Model(&commentRecord).Where("author_id = ? AND video_id = ?", commentRecord.AuthorId, commentRecord.VideoId).Update("Status", commentRecord.Status); err != nil {
		return err.Error
	}
	return nil
}
