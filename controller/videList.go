package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yancy0109/SimpleTiktok/middleware"
	"github.com/yancy0109/SimpleTiktok/service"
	"net/http"
	"strconv"
)

func PublishList(context *gin.Context) {
	token := context.Query("token")
	var videoListService service.VideoListService
	//处理传来的参数userId,并查看它是否合法
	//userId的情况交给下层处理
	userId, error := strconv.ParseInt(context.Query("user_id"), 10, 64)
	if error != nil || token == "" {
		videoFeed := videoListService.GetVideoList(-1)
		context.JSON(http.StatusOK, videoFeed)
	} else {
		//检查token
		_, err := middleware.ParseToken(token)
		if err != nil {
			msg := "token无效"
			context.JSON(http.StatusOK, service.VideoListModal{
				StatusCode: 1,
				StatusMsg:  &msg,
				VideoList:  nil,
			})
			return
		}
		videoFeed := videoListService.GetVideoList(userId)
		context.JSON(http.StatusOK, videoFeed)
	}

}
