package service

import "github.com/yancy0109/SimpleTiktok/repository"

func GetBeFollow(userId int64) ([]repository.Author, error) {
	return NewGetFollowList(userId).Get()
}

func (f *GetFollowListFlow) Get() ([]repository.Author, error) {
	if err := f.CheckUserId(); err != nil {
		return []repository.Author{}, err
	}
	if err := f.GetBeFollowList(); err != nil {
		return []repository.Author{}, err
	}
	return f.FollowList, nil
}

//先获取Id的关注列表，再根据列表返回关注信息
func (f *GetFollowListFlow) GetBeFollowList() error {
	followIdList := repository.NewFollowListDaoInstance().GetBeFollowIdList(f.UserId)
	for _, followId := range followIdList {
		var user *repository.Author
		var err error
		if user, err = repository.NewVideoDaoInstance().AuthorInformation(followId, f.UserId); err != nil {
			return err
		}
		f.FollowList = append(f.FollowList, *user)
	}
	return nil
}
