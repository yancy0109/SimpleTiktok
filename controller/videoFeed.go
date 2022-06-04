package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yancy0109/SimpleTiktok/middleware"
	"github.com/yancy0109/SimpleTiktok/service"
	"net/http"
	"strconv"
	"time"
)

func VideoFeed(context *gin.Context) {
	token := context.Query("token")
	var videoService service.VideoService
	//处理时间，分为携带和不携带
	latestTimeStr := context.Query("latest_time")
	var latestTime int64
	if latestTimeStr == "" {
		latestTime = time.Now().Unix() / 1000
	} else {
		latestTime, _ = strconv.ParseInt(latestTimeStr, 10, 64)
	}

	//未登录
	if token == "" {
		videoFeed := videoService.GetVideoFeed(latestTime, -1)
		context.JSON(http.StatusOK, videoFeed)
	} else {
		//检查token
		userId, err := middleware.ParseToken(token)
		if err != nil {
			context.JSON(http.StatusOK, Response{
				StatusCode: 1,
				StatusMsg:  "token无效",
			})
			return
		}
		videoFeed := videoService.GetVideoFeed(latestTime, userId)
		context.JSON(http.StatusOK, videoFeed)
	}

}
