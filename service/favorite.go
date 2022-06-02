package service

import (
	"github.com/yancy0109/SimpleTiktok/repository"
	"strconv"
)

type FavoriteService struct {
}

var favoriteDao = repository.NewfavoriteDaoInstance()

func (f FavoriteService) Update(userid string, videoid string, actionType string) error {
	userId, err := strconv.ParseInt(userid, 10, 64)
	if err != nil {
		return err
	}
	videoId, err := strconv.ParseInt(videoid, 10, 64)
	if err != nil {
		return err
	}
	status, err := strconv.ParseInt(actionType, 10, 64)
	if err != nil {
		return err
	}
	favoriteRecord := repository.FavoriteRecord{
		Userid:  userId,
		Videoid: videoId,
		Status:  status,
	}
	return favoriteDao.UpdateStatus(favoriteRecord)
}
