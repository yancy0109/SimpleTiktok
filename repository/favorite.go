package repository

import (
	"sync"
	"time"
)

type FavoriteRecord struct {
	//gorm.Model
	Id         int64     `gorm:"column:id" `
	Videoid    int64     `gorm:"column:video_id" `
	Userid     int64     `gorm:"column:user_id" `
	Status     int64     `gorm:"column:status" `
	Createdate time.Time `gorm:"column:create_date" `
}
type FavoriteDao struct {
}

var favoriteDao *FavoriteDao
var favoriteOnce sync.Once

func NewfavoriteDaoInstance() *FavoriteDao {
	favoriteOnce.Do(
		func() {
			favoriteDao = &FavoriteDao{}
		})
	return favoriteDao
}
func (FavoriteRecord) TableName() string {
	return "video_favorite"
}

func (*FavoriteDao) GetList(UserId int64) ([]Video, error) {
	var results []Video
	result := db.Table("video v").Select("v.id , v.author_id ,v.title ,v.play_url ,v.cover_url , v.create_date").Joins("left join video_favorite f on f.user_id = ? where v.id = f.video_id AND v.status <> 0", UserId).Scan(&results)
	if result.Error != nil {
		return nil, result.Error
	}
	//返回
	return results, nil
}

// Favorite 返回是否点赞
func (*FavoriteDao) Favorite(UserId int64, VideoId int64) (bool, error) {
	var id int64
	//查询id是否存在
	result := db.Raw("SELECT id FROM video_favorite WHERE user_id = ? AND video_id = ? AND Status = 1 ", UserId, VideoId).Scan(&id)
	// == 0 表示不存在该记录
	if result.Error != nil {
		return false, result.Error
	}
	return id != 0, nil
}

func (*FavoriteDao) GetCount(VideoId int64) (int64, error) {
	var count int64
	result := db.Table("video_favorite").Where("video_id = ? AND Status = ? ", VideoId, 1).Count(&count)
	if result.Error != nil {
		return -1, result.Error
	}
	return count, nil
}

func (*FavoriteDao) UpdateStatus(favoriteRecord FavoriteRecord) error {
	var id int64
	//查询记录是否存在
	db.Raw("SELECT id FROM video_favorite WHERE user_id = ? AND video_id = ?", favoriteRecord.Userid, favoriteRecord.Videoid).Scan(&id)
	if id == 0 {
		favoriteRecord.Createdate = time.Now()
		err := db.Model(&favoriteRecord).Create(&favoriteRecord) //记录不存在-创建
		if err != nil {
			return err.Error
		}
	}
	favoriteRecord.Id = id
	err := favoriteDao.ChangeStatus(favoriteRecord) //记录存在-更新
	if err != nil {
		return err
	}
	return nil
}
func (*FavoriteDao) ChangeStatus(favoriteRecord FavoriteRecord) error {
	if err := db.Model(&favoriteRecord).Where("user_id = ? AND video_id = ?", favoriteRecord.Userid, favoriteRecord.Videoid).Update("Status", favoriteRecord.Status); err != nil {
		return err.Error
	}
	return nil
}
