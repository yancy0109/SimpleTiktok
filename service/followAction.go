package service

import (
	"github.com/yancy0109/SimpleTiktok/repository"
)

var followActionAction = repository.NewFollowActionDaoInstance()

type FollowActionService struct {
}

// UpdateFollowStatus 更新数据 action 1-关注，2-取消关注,首先查询用户是否存在，不存在则返回一个error
func (*FollowActionService) UpdateFollowStatus(follow int64, beFollow int64, action int) (*string, error, int) {
	var msg = "success"
	code := 0
	_, userErr := repository.FindUserById(beFollow)
	if userErr != nil {
		msg = "用户不存在"
		return &msg, nil, 1
	}
	followStatus, err := followActionAction.GetFollowState(follow, beFollow)
	if err != nil {
		return nil, err, 1
	}
	//没有信息
	if followStatus == nil {
		if action == 2 {
			msg = "未关注"
			code = 1
		} else {
			err := followActionAction.CreateFollow(follow, beFollow)
			if err != nil {
				return nil, err, code
			}
		}
	} else {
		//有信息校验是否删除了
		if followStatus.IsDel == 1 {
			//已经删除了这个关系
			if action == 1 {
				err := followActionAction.UpdateStatus(follow, beFollow, followStatus.Id, 0)
				if err != nil {
					return nil, err, 1
				}
			} else {
				msg = "未关注"
			}
		} else {
			//还没有删除这个关系
			if action == 2 {
				err := followActionAction.UpdateStatus(follow, beFollow, followStatus.Id, 1)
				if err != nil {
					return nil, err, 1
				}
			} else {
				msg = "已关注"
			}
		}

	}
	return &msg, nil, code
}
