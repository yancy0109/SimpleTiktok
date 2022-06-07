package repository

import (
	"sync"
	"time"
)

type CommentList struct {
	Id         int64     `gorm:"id"`
	VideoId    int64     `gorm:"video_id"`
	AuthorId   int64     `gorm:"author_id"`
	Content    string    `gorm:"content"`
	CreateDate time.Time `gorm:"create_date"`
	Status     bool      `gorm:"status"`
}

func (CommentList) TableName() string {
	return "comment"
}

type CommentListDao struct {
}

var commentListDao *CommentListDao
var commentListOnce sync.Once

func NewCommentListDaoInstance() *CommentListDao {
	commentListOnce.Do(
		func() {
			commentListDao = &CommentListDao{}
		})
	return commentListDao
}

//根据视频id获取评论列表 按照时间降序排列
func (*CommentListDao) GetCommentListById(videoId int64) []CommentList {
	var commentList []CommentList
	db.Order("create_date desc").Find(&commentList, "video_id = ? and status = ?", videoId, 1)
	return commentList
}
