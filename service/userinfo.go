package service

import "github.com/yancy0109/SimpleTiktok/repository"

func GetUserInfo(userId int64) (repository.Author, error) {

	return NewGetUserInfoFlow(userId).Do()
}

type GetUserInfoFlow struct {
	userId int64
	User   *repository.Author
}

func NewGetUserInfoFlow(userId int64) *GetUserInfoFlow {
	return &GetUserInfoFlow{userId: userId}
}

func (f *GetUserInfoFlow) Do() (repository.Author, error) {
	if err := f.CheckUserId(); err != nil {
		return repository.Author{}, err
	}
	if err := f.GetUserInfo(); err != nil {
		return repository.Author{}, err
	}
	return *f.User, nil
}
func (f *GetUserInfoFlow) CheckUserId() error {
	if _, err := repository.FindUserById(f.userId); err != nil {
		return err
	}
	return nil
}
func (f *GetUserInfoFlow) GetUserInfo() error {
	var err error
	if f.User, err = repository.NewVideoDaoInstance().GetUserInformation(f.userId); err != nil {
		return err
	}
	return nil
}
