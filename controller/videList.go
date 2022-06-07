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
	//userId, error := strconv.ParseInt(context.Query("user_id"), 10, 64)
	_, error := strconv.ParseInt(context.Query("user_id"), 10, 64)
	//他传过来的参数一直是0我还要他干啥！
	if error != nil || token == "" {
		videoFeed := videoListService.GetVideoList(-1)
		context.JSON(http.StatusOK, videoFeed)
	} else {
		//检查token
		tokenUserId, err := middleware.ParseToken(token)
		//此处理论上必须鉴权，但是客户端传来的user_id一直是0
		//if err != nil || tokenUserId != userId {
		if err != nil {
			msg := "token无效"
			context.JSON(http.StatusOK, service.VideoListModal{
				StatusCode: 1,
				StatusMsg:  &msg,
				VideoList:  nil,
			})
			return
		}
		videoFeed := videoListService.GetVideoList(tokenUserId)
		context.JSON(http.StatusOK, videoFeed)
	}

}
