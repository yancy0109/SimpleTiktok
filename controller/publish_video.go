package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yancy0109/SimpleTiktok/service"
	"net/http"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

func Publish(context *gin.Context) {
	//解析context中 data token title
	title := context.PostForm("title")
	token := context.PostForm("token")
	//检验token获取user
	var authorId int64
	if token != "" {
		authorId = 2123123
	}
	data, err := context.FormFile("data")
	//接收文件失败返回
	if err != nil {
		context.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	//调取service层储存
	if err := service.PublishSave(authorId, title, data, context); err != nil {
		fmt.Println(err)
		context.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "保存失败",
		})
	}
	context.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  "uploaded successfully",
	})
}
