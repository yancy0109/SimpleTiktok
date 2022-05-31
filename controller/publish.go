package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yancy0109/SimpleTiktok/common"
	"net/http"
)

func Publish(context *gin.Context) {
	context.JSON(http.StatusOK,
		&common.Response{
			StatusCode: 1,
			StatusMsg:  "拿捏了",
		})
}
