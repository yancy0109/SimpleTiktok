package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yancy0109/SimpleTiktok/middleware"
	"github.com/yancy0109/SimpleTiktok/service"
	"net/http"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

var favoriteService service.FavoriteService

func FavoriteAction(context *gin.Context) {
	token := context.Query("token")
	//检查token
	userid, err := middleware.ParseToken(token)
	if err != nil {
		context.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "token无效",
		})
		return
	}
	videoid := context.Query("video_id")
	actionType := context.Query("action_type")
	err = favoriteService.Update(userid, videoid, actionType)
	if err != nil {
		context.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
	} else {
		context.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "successfully"})
	}
}
func FavoriteList(context *gin.Context) {
	token := context.Query("token")
	//检查token
	userid, err := middleware.ParseToken(token)
	if err != nil {
		context.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "token无效",
		})
		return
	}
	favoriteList := favoriteService.FavoriteList(userid)
	context.JSON(http.StatusOK, favoriteList)
}
