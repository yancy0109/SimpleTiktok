package service

import "github.com/yancy0109/SimpleTiktok/repository"

type GetFollowListFlow struct {
	UserId     int64               `json:"user_id"`
	FollowList []repository.Author `json:"user_list"`
}

func GetFollowList(userId int64) ([]repository.Author, error) {
	return NewGetFollowList(userId).Do()
}

func NewGetFollowList(userId int64) *GetFollowListFlow {
	return &GetFollowListFlow{
		UserId: userId,
	}
}

func (f *GetFollowListFlow) Do() ([]repository.Author, error) {
	if err := f.CheckUserId(); err != nil {
		return []repository.Author{}, err
	}
	if err := f.GetList(); err != nil {
		return []repository.Author{}, err
	}
	return f.FollowList, nil
}

//检查用户Id是否存在
func (f *GetFollowListFlow) CheckUserId() error {
	return nil
}

//先获取Id的粉丝列表，再根据列表返回粉丝信息
func (f *GetFollowListFlow) GetList() error {
	followIdList := repository.NewFollowListDaoInstance().GetFollowIdList(f.UserId)
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
