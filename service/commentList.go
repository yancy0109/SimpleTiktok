package service

import (
	"github.com/yancy0109/SimpleTiktok/repository"
)

type CommentInfo struct {
	Id         int64             `json:"id"`
	User       repository.Author `json:"user"`
	Content    string            `json:"content"`
	CreateDate string            `json:"create_date"`
}

func GetCommentList(videoId int64, userId int64) ([]CommentInfo, error) {
	return NewGetCommentList(videoId, userId).Do()
}

func NewGetCommentList(videoId int64, userId int64) *GetCommentListFlow {
	return &GetCommentListFlow{VideoId: videoId, UserId: userId}
}

type GetCommentListFlow struct {
	VideoId         int64
	UserId          int64
	CommentInfoList []CommentInfo
}

func (f *GetCommentListFlow) Do() ([]CommentInfo, error) {
	if err := f.CheckVideoId(); err != nil {
		return []CommentInfo{}, err
	}
	if err := f.GetList(); err != nil {
		return []CommentInfo{}, err
	}
	return f.CommentInfoList, nil
}

func (f *GetCommentListFlow) CheckVideoId() error {
	if err := repository.NewVideoDaoInstance().IsExistVideo(f.VideoId); err != nil {
		return err
	}
	return nil
}

func (f *GetCommentListFlow) GetList() error {
	var commentList []repository.CommentList
	commentList = repository.NewCommentListDaoInstance().GetCommentListById(f.VideoId)
	for _, comment := range commentList {
		var user *repository.Author
		var commentInfo CommentInfo
		var err error
		if user, err = repository.NewVideoDaoInstance().AuthorInformation(comment.AuthorId, f.UserId, 1); err != nil {
			return err
		}
		commentInfo.User = *user
		commentInfo.Content = comment.Content
		commentInfo.CreateDate = comment.CreateDate.Format("01-02")
		commentInfo.Id = comment.Id
		f.CommentInfoList = append(f.CommentInfoList, commentInfo)
	}
	return nil
}
