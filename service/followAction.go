package service

import (
	"github.com/yancy0109/SimpleTiktok/repository"
)

var followActionAction = repository.NewFollowActionDaoInstance()

type FollowActionService struct {
}

// UpdateFollowStatus 更新数据 action 1-关注，2-取消关注,首先查询用户是否存在，不存在则返回一个error
func (*FollowActionService) UpdateFollowStatus(follower int64, follow int64, action int) (*string, error, int) {
	var msg = "success"
	code := 0
	_, userErr := repository.FindUserById(follow)
	if userErr != nil {
		msg = "用户不存在"
		return &msg, nil, 1
	}
	followStatus, err := followActionAction.GetFollowState(follower, follow)
	if err != nil {
		return nil, err, 1
	}
	//没有信息
	if followStatus == nil {
		if action == 2 {
			msg = "未关注"
			code = 1
		} else {
			err := followActionAction.CreateFollow(follower, follow)
			if err != nil {
				return nil, err, code
			}
		}
	} else {
		//有信息 校验是否删除了
		isDel := action - 1
		if followStatus.IsDel == isDel {
			code = 1
			if isDel == 0 {
				msg = "已关注"
			} else {
				msg = "未关注"
			}
		} else {
			err := followActionAction.UpdateStatus(follower, follow, followStatus.Id, isDel)
			if err != nil {
				return nil, err, 1
			}
		}
	}
	return &msg, nil, code
}
