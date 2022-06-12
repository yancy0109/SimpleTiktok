package repository

import "sync"

type FollowListDao struct {
}

var followListDao *FollowListDao
var followListOnce sync.Once

func NewFollowListDaoInstance() *FollowListDao {
	followListOnce.Do(func() {
		followListDao = &FollowListDao{}
	})
	return followListDao
}

//根据用户id返回其粉丝的id数组
func (*FollowListDao) GetFollowIdList(userId int64) []int64 {
	var followIdList []int64
	var followList []Follow
	db.Select("follow").Find(&followList, "be_follow = ? and is_del = ?", userId, 0)
	for _, followInfo := range followList {
		followIdList = append(followIdList, followInfo.Follow)
	}
	return followIdList
}

//根据用户id返回其关注的id数组
func (*FollowListDao) GetBeFollowIdList(userId int64) []int64 {
	var followIdList []int64
	var followList []Follow
	db.Select("be_follow").Find(&followList, "follow = ? and is_del = ?", userId, 0)
	for _, followInfo := range followList {
		followIdList = append(followIdList, followInfo.Follow)
	}
	return followIdList
}