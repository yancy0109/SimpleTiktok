package service

import (
	"github.com/yancy0109/SimpleTiktok/repository"
	"strconv"
)

type CommentService struct {
}

var commentDao = repository.NewCommentDaoInstance()

func (f CommentService) INSERT(authorid int64, videoid string, actionType string, content string) (int64, error) {
	videoId, _ := strconv.ParseInt(videoid, 10, 64)
	status, _ := strconv.ParseInt(actionType, 10, 64)
	commentRecord := repository.CommentRecord{
		AuthorId: authorid,
		VideoId:  videoId,
		Status:   status,
		Content:  content,
	}
	return commentDao.UpdateStatus(commentRecord), nil
}

func (f CommentService) UPDATE(authorid int64, videoid string, actionType string, comment_id string) error {
	videoId, err := strconv.ParseInt(videoid, 10, 64)
	if err != nil {
		return err
	}
	status, err := strconv.ParseInt(actionType, 10, 64)
	if err != nil {
		return err
	}
	comment_Id, err := strconv.ParseInt(comment_id, 10, 64)
	if err != nil {
		return err
	}
	commentRecord := repository.CommentRecord{
		Id:       comment_Id,
		AuthorId: authorid,
		VideoId:  videoId,
		Status:   status,
	}
	return commentDao.ChangeStatus(commentRecord)
}
