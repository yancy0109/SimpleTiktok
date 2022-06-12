package service

import "github.com/yancy0109/SimpleTiktok/repository"

type GetFollowerListFlow struct {
	UserId       int64               `json:"user_id"`
	FollowerList []repository.Author `json:"user_list"`
}

func GetFollowerList(userId int64) ([]repository.Author, error) {
	return NewGetFollowerList(userId).Do()
}

func NewGetFollowerList(userId int64) *GetFollowerListFlow {
	return &GetFollowerListFlow{
		UserId: userId,
	}
}

func (f *GetFollowerListFlow) Do() ([]repository.Author, error) {
	if err := f.CheckUserId(); err != nil {
		return []repository.Author{}, err
	}
	if err := f.GetList(); err != nil {
		return []repository.Author{}, err
	}
	return f.FollowerList, nil
}

//检查用户Id是否存在
func (f *GetFollowerListFlow) CheckUserId() error {
	if _, err := repository.FindUserById(f.UserId); err != nil {
		return err
	}
	return nil
}

//先获取Id的粉丝列表，再根据列表返回粉丝信息
func (f *GetFollowerListFlow) GetList() error {
	followIdList := repository.NewFollowListDaoInstance().GetFollowIdList(f.UserId)
	for _, followId := range followIdList {
		var user *repository.Author
		var err error
		if user, err = repository.NewVideoDaoInstance().AuthorInformation(followId, f.UserId); err != nil {
			return err
		}
		f.FollowerList = append(f.FollowerList, *user)
	}
	return nil
}
