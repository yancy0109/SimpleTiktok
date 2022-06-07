package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yancy0109/SimpleTiktok/middleware"
	"github.com/yancy0109/SimpleTiktok/service"
	"net/http"
	"strconv"
)

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
	beFollow := context.Query("to_user_id")
	beFollowId, errorOne := strconv.ParseInt(beFollow, 10, 64)

	actionType := context.Query("action_type")
	action, errorTwo := strconv.Atoi(actionType)

	if beFollowId == userId {
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
	msg, error, code := followActionService.UpdateFollowStatus(userId, beFollowId, action)
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
