package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yancy0109/SimpleTiktok/middleware"
	"github.com/yancy0109/SimpleTiktok/repository"
	"github.com/yancy0109/SimpleTiktok/service"
	"net/http"
)

type FollowListResponse struct {
	StatusCode int64               `json:"status_code"`
	StatusMsg  string              `json:"status_msg"`
	FollowList []repository.Author `json:"user_list"`
}

func FollowerList(context *gin.Context) {
	var token string
	var userId int64
	var exist bool
	var err error
	if token, exist = context.GetQuery("token"); !exist {
		context.JSON(http.StatusOK, FollowListResponse{
			StatusCode: -1,
			StatusMsg:  "缺少token",
		})
		return
	}
	if userId, err = middleware.ParseToken(token); err != nil {
		context.JSON(http.StatusOK, FollowListResponse{
			StatusCode: -1,
			StatusMsg:  "token无效",
		})
		return
	}
	var followList []repository.Author
	if followList, err = service.GetFollowerList(userId); err != nil {
		context.JSON(http.StatusOK, FollowListResponse{
			StatusCode: -1,
			StatusMsg:  "获取粉丝列表失败",
		})
		return
	}
	context.JSON(http.StatusOK, FollowListResponse{
		StatusCode: 0,
		StatusMsg:  "获取粉丝列表成功",
		FollowList: followList,
	})
}
