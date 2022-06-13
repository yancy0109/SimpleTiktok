package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yancy0109/SimpleTiktok/middleware"
	"github.com/yancy0109/SimpleTiktok/repository"
	"github.com/yancy0109/SimpleTiktok/service"
)

type FollowListResponse struct {
	StatusCode int64               `json:"status_code"`
	StatusMsg  string              `json:"status_msg"`
	UserList   []repository.Author `json:"user_list"`
}

func FollowerList(context *gin.Context) {
	var token string
	var userId int64
	var tokenUserId int64
	var exist bool
	var err error
	if token, exist = context.GetQuery("token"); !exist {
		context.JSON(http.StatusOK, FollowListResponse{
			StatusCode: -1,
			StatusMsg:  "缺少token",
		})
		return
	}
	if tokenUserId, err = middleware.ParseToken(token); err != nil {
		context.JSON(http.StatusOK, FollowListResponse{
			StatusCode: -1,
			StatusMsg:  "token无效",
		})
		return
	}
	if userId, err = strconv.ParseInt(context.Query("user_id"), 10, 64); err != nil {
		context.JSON(http.StatusOK, FollowListResponse{
			StatusCode: -1,
			StatusMsg:  "参数错误",
		})
		return
	}
	var followerList []repository.Author
	if followerList, err = service.GetFollowerList(tokenUserId, userId); err != nil {
		context.JSON(http.StatusOK, FollowListResponse{
			StatusCode: -1,
			StatusMsg:  "获取粉丝列表失败",
		})
		return
	}
	context.JSON(http.StatusOK, FollowListResponse{
		StatusCode: 0,
		StatusMsg:  "获取粉丝列表成功",
		UserList:   followerList,
	})
}

func FollowList(context *gin.Context) {
	var token string
	var userId int64
	var tokenUserId int64
	var exist bool
	var err error
	if token, exist = context.GetQuery("token"); !exist {
		context.JSON(http.StatusOK, FollowListResponse{
			StatusCode: -1,
			StatusMsg:  "缺少token",
		})
		return
	}
	if tokenUserId, err = middleware.ParseToken(token); err != nil {
		context.JSON(http.StatusOK, FollowListResponse{
			StatusCode: -1,
			StatusMsg:  "token无效",
		})
		return
	}
	if userId, err = strconv.ParseInt(context.Query("user_id"), 10, 64); err != nil {
		context.JSON(http.StatusOK, FollowListResponse{
			StatusCode: -1,
			StatusMsg:  "参数错误",
		})
		return
	}
	var followList []repository.Author
	if followList, err = service.GetFollowList(tokenUserId, userId); err != nil {
		context.JSON(http.StatusOK, FollowListResponse{
			StatusCode: -1,
			StatusMsg:  "获取关注列表失败",
		})
		return
	}
	context.JSON(http.StatusOK, FollowListResponse{
		StatusCode: 0,
		StatusMsg:  "获取关注列表成功",
		UserList:   followList,
	})
}

type RelationActionResponse struct {
	StatusCode int    `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
}

func RelationAction(context *gin.Context) {
	token := context.Query("token")
	//检查token
	userId, err := middleware.ParseToken(token)
	if err != nil {
		context.JSON(http.StatusOK, RelationActionResponse{
			StatusCode: 1,
			StatusMsg:  "token无效",
		})
		return
	}
	follow := context.Query("to_user_id")
	followId, errorOne := strconv.ParseInt(follow, 10, 64)

	actionType := context.Query("action_type")
	action, errorTwo := strconv.Atoi(actionType)

	if followId == userId {
		context.JSON(http.StatusOK, RelationActionResponse{
			StatusCode: 1,
			StatusMsg:  "无法关注自己",
		})
		return
	}
	if errorOne != nil || errorTwo != nil || (action != 1 && action != 2) {
		context.JSON(http.StatusOK, RelationActionResponse{
			StatusCode: 1,
			StatusMsg:  "参数错误",
		})
		return
	}

	var followActionService service.FollowActionService
	msg, error, code := followActionService.UpdateFollowStatus(userId, followId, action)
	if error != nil {
		context.JSON(http.StatusOK, RelationActionResponse{
			StatusCode: 1,
			StatusMsg:  "内部错误",
		})
		return
	}
	context.JSON(http.StatusOK, RelationActionResponse{
		StatusCode: code,
		StatusMsg:  *msg,
	})
	return
}
