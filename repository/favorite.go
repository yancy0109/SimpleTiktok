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
func (*FavoriteDao) GetList(favoriteRecord FavoriteRecord) []FavoriteRecord {
	var results []FavoriteRecord
	db.Table("video_favorite").Where("video_id = ? AND Status = ?", favoriteRecord.Userid, 1).Find(&results) // (*sql.Row)
	return results
}
func (*FavoriteDao) GetCount(favoriteRecord FavoriteRecord) int64 {
	var count int64
	db.Model(&favoriteRecord).Where("video_id = ? AND Status = ? ", favoriteRecord.Videoid, 1).Count(&count)

	return count
}

func (*FavoriteDao) UpdateStatus(favoriteRecord FavoriteRecord) error {
	var id int64
	//查询id是否存在
	db.Raw("SELECT id FROM video_favorite WHERE user_id = ? AND video_id = ?", favoriteRecord.Userid, favoriteRecord.Videoid).Scan(&id)
	if id == 0 {
		favoriteRecord.Createdate = time.Now()
		err := db.Model(&favoriteRecord).Create(&favoriteRecord)
		if err != nil {
			return err.Error
		}
	}
	favoriteRecord.Id = id
	err := favoriteDao.ChangeStatus(favoriteRecord)
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
