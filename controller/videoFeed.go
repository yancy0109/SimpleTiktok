package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yancy0109/SimpleTiktok/middleware"
	"github.com/yancy0109/SimpleTiktok/service"
	"net/http"
	"strconv"
)

func VideoFeed(context *gin.Context) {
	token := context.Query("token")
	//检查token
	userId, err := middleware.ParseToken(token)
	if err != nil {
		context.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "token无效",
		})
		return
	}
	latestTime, _ := strconv.Atoi(context.Query("latest_time"))
	var videoService service.VideoService
	videoFeed := videoService.GetVideoFeed(latestTime, userId)
	context.JSON(http.StatusOK, videoFeed)
}
