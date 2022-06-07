package service

type UserService struct {
}

func (*UserService) QueryUserInfo(userId int64, tokenId int64) (User, error) {
	user, err := videoListDao.AuthorInformation(userId, tokenId)
	return User{
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		ID:            user.Id,
		IsFollow:      user.IsFollow,
		Name:          user.UserName,
	}, err
}
