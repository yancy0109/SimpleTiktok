package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yancy0109/SimpleTiktok/service"
	"net/http"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

func FavoriteAction(context *gin.Context) {
	token := context.Query("token")

	//检查token
	if token != "" {
		userid := "1"
		videoid := context.Query("video_id")
		actionType := context.Query("action_type")
		var favoriteService service.FavoriteService
		err := favoriteService.Update(userid, videoid, actionType)
		if err != nil {
			context.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
		} else {
			context.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "successfully"})
		}
	} else {
		context.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}
