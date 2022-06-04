package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yancy0109/SimpleTiktok/service"
	"net/http"
	"strconv"
)

type PublishListResponse struct {
	StatusCode int32             `json:"status_code"`
	StatusMsg  string            `json:"status_msg"`
	VideoList  service.VideoList `json:"video_list"`
}

func PublishList(context *gin.Context) {
	var user_id string
	var userId int64
	var exist bool
	var err error
	if user_id, exist = context.GetQuery("user_id"); !exist {
		context.JSON(http.StatusOK, PublishListResponse{
			StatusCode: -1,
			StatusMsg:  "缺少user_id",
		})
	}
	if userId, err = strconv.ParseInt(user_id, 10, 64); err != nil {
		context.JSON(http.StatusOK, PublishListResponse{
			StatusCode: -1,
			StatusMsg:  "user_id解析错误",
		})
	}
	fmt.Println(userId)
}
