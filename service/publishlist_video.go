package service

type VideoList struct {
}
type AuthorInfoForVideoList struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	FollowCount   int32  `json:"follow_count"`
	FollowerCount int32  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}
type VideoInfoForVideoList struct {
}
