package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

func Publish(context *gin.Context) {
	context.JSON(http.StatusOK,
		&Response{
			StatusCode: 1,
			StatusMsg:  "拿捏了",
		})
}
