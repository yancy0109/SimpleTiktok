package service

import (
	"github.com/yancy0109/SimpleTiktok/repository"
	"strconv"
)

type FavoriteService struct {
}

var favoriteDao = repository.NewfavoriteDaoInstance()

func (f FavoriteService) Update(userid int64, videoid string, actionType string) error {
	videoId, err := strconv.ParseInt(videoid, 10, 64)
	if err != nil {
		return err
	}
	status, err := strconv.ParseInt(actionType, 10, 64)
	if err != nil {
		return err
	}
	favoriteRecord := repository.FavoriteRecord{
		Userid:  userid,
		Videoid: videoId,
		Status:  status,
	}
	return favoriteDao.UpdateStatus(favoriteRecord)
}
