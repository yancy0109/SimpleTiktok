package service

import "github.com/yancy0109/SimpleTiktok/repository"

func GetFollowList(tokenUserId int64, userId int64) ([]repository.Author, error) {
	return NewGetFollowerList(tokenUserId, userId).Get()
}

func (f *GetFollowerListFlow) Get() ([]repository.Author, error) {
	if err := f.CheckUserId(); err != nil {
		return []repository.Author{}, err
	}
	if err := f.GetFollowList(); err != nil {
		return []repository.Author{}, err
	}
	return f.FollowerList, nil
}

//先获取Id的关注列表，再根据列表返回关注信息
func (f *GetFollowerListFlow) GetFollowList() error {
	followIdList := repository.NewFollowListDaoInstance().GetFollowIdList(f.UserId)
	for _, followId := range followIdList {
		var user *repository.Author
		var err error
		if user, err = repository.NewVideoDaoInstance().AuthorInformation(followId, f.TokenUserId, 0); err != nil {
			return err
		}
		f.FollowerList = append(f.FollowerList, *user)
	}
	return nil
}
